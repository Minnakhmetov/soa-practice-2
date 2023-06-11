package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/Minnakhmetov/soa-practice-2/mafia"
	pb "github.com/Minnakhmetov/soa-practice-2/mafia"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type connection struct {
	ch    chan string
	close chan struct{}
}

type mafiaServer struct {
	pb.UnimplementedMafiaServer

	playerToSession   map[string]*gameSession
	playerToConnecton map[string]*connection

	waitingListMutex sync.Mutex
	waitingList      []string

	msgSender
}

func (t *mafiaServer) init() {
	// TO DO: run garbage collector
}

func (t *mafiaServer) disconnect(username string) {
	delete(t.playerToConnecton, username)
}

func (t *mafiaServer) send(sender string, receiver string, msg string) {
	conn, ok := t.playerToConnecton[receiver]
	if !ok {
		panic(fmt.Sprintf("user %s not connected.", receiver))
	}
	select {
	case conn.ch <- fmt.Sprintf("[%s] %s", sender, msg):
		// noop
	default:
		log.Printf("%s user msg buffer is full", receiver)
		t.disconnect(sender)
	}
}

func (t *mafiaServer) sendServerMessage(receiver string, msg string) {
	t.send("server", receiver, msg)
}

func (t *mafiaServer) runGameSession(usernames []string) {
	session := MakeGameSession(usernames, t)

	for _, username := range usernames {
		t.playerToSession[username] = session
	}

	session.Run()
}

func (t *mafiaServer) inActiveGameSession(username string) bool {
	_, ok := t.playerToSession[username]
	return ok
}

func (t *mafiaServer) addToWaitingList(username string) {
	t.waitingListMutex.Lock()
	defer t.waitingListMutex.Unlock()

	t.waitingList = append(t.waitingList, username)

	log.Printf("added %s to waiting list", username)

	if len(t.waitingList) == mafia.PlayersInGame {
		t.runGameSession(t.waitingList)

	} else if len(t.waitingList) < mafia.PlayersInGame {
		needPlayers := mafia.PlayersInGame - len(t.waitingList)

		log.Printf("need %d more players for game", needPlayers)

		makeGroupMsgSender(t, t.waitingList).sendAllServerMessage(
			fmt.Sprintf("Waiting for players. %d more players are needed.", needPlayers),
		)

	} else {
		panic("waiting list too long")
	}
}

func (t *mafiaServer) Login(request *mafia.LoginRequest, serv mafia.Mafia_LoginServer) error {
	username := extractUsername(serv.Context())

	// if _, ok := t.playerToConnecton[username]; ok {
	// 	return status.Errorf(codes.AlreadyExists, "User with this name already logged in.")
	// }

	close := make(chan struct{})
	ch := make(chan string, 5)

	t.playerToConnecton[username] = &connection{ch: ch, close: close}

	log.Printf("user %s connected\n", username)

	wg := sync.WaitGroup{}
	defer wg.Wait()
	wg.Add(1)

	areYouReady := make(chan struct{})
	yesCaptain := make(chan struct{})

	go func() {
		defer wg.Done()
		for {
			select {
			case <-close:
				t.disconnect(username)
				return
			case msg := <-ch:
				log.Printf("sending \"%s\" to %s\n", msg, username)
				err := serv.Send(&pb.Event{Message: msg})
				if err != nil {
					log.Printf("error when sending message to %s: %s\n", username, err.Error())
					t.disconnect(username)
					return
				}
			case <-areYouReady:
				yesCaptain <- struct{}{}
			}
		}
	}()

	areYouReady <- struct{}{}
	<-yesCaptain

	t.sendServerMessage(username, "You are connected to server.")

	// TO DO: check if player already in active game session (reconnected)

	t.addToWaitingList(username)
	wg.Wait()
	return nil
}

func (t *mafiaServer) isLoggedIn(username string) bool {
	_, ok := t.playerToConnecton[username]
	return ok
}

func (t *mafiaServer) GetGameSession(username string) (*gameSession, error) {
	if !t.isLoggedIn(username) {
		return nil, status.Error(codes.Unauthenticated, "User is not logged in.")
	}
	session, ok := t.playerToSession[username]
	if !ok {
		return nil, status.Error(codes.NotFound, "User is not in a game session.")
	}
	return session, nil
}

func handleClientAction[R any](server *mafiaServer, ctx context.Context, action func(*gameSession, string)) (*R, error) {
	username := extractUsername(ctx)
	session, err := server.GetGameSession(username)
	if err != nil {
		return nil, err
	}
	action(session, username)
	return new(R), nil
}

func (t *mafiaServer) EndTurn(ctx context.Context, request *mafia.EndTurnRequest) (*mafia.EndTurnResponse, error) {
	return handleClientAction[mafia.EndTurnResponse](t, ctx, (*gameSession).EndTurn)
}

func (t *mafiaServer) VoteAgainst(ctx context.Context, request *mafia.VoteAgainstRequest) (*mafia.VoteAgainstResponse, error) {
	return handleClientAction[mafia.VoteAgainstResponse](t, ctx,
		func(gs *gameSession, username string) {
			gs.VoteAgainst(username, request.GetTarget())
		},
	)
}

func (t *mafiaServer) Shoot(ctx context.Context, request *mafia.ShootRequest) (*mafia.ShootResponse, error) {
	return handleClientAction[mafia.ShootResponse](t, ctx,
		func(gs *gameSession, username string) {
			gs.Shoot(username, request.GetTarget())
		},
	)
}

func (t *mafiaServer) Check(ctx context.Context, request *mafia.CheckRequest) (*mafia.CheckResponse, error) {
	return handleClientAction[mafia.CheckResponse](t, ctx,
		func(gs *gameSession, username string) {
			gs.Check(username, request.GetTarget())
		},
	)
}

func (t *mafiaServer) PublishCheckResult(ctx context.Context, _ *mafia.PublishCheckResultRequest) (*mafia.PublishCheckResultResponse, error) {
	return handleClientAction[mafia.PublishCheckResultResponse](t, ctx, (*gameSession).PublishCheckResult)
}

func extractUsername(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}
	return md["username"][0]
}

func MakeMafiaServer() *mafiaServer {
	return &mafiaServer{
		playerToSession:   map[string]*gameSession{},
		playerToConnecton: map[string]*connection{},
	}
}

type msgSender interface {
	send(sender string, receiver string, msg string)
	sendServerMessage(receiver string, msg string)
}

type groupMsgSender struct {
	msgSender
	members []string
}

func (t *groupMsgSender) broadcast(sender string, msg string) {
	for _, player := range t.members {
		t.send(sender, player, msg)
	}
}

func (t *groupMsgSender) sendAll(sender string, msg string) {
	t.broadcast(sender, msg)
}

func (t *groupMsgSender) sendAllServerMessage(msg string) {
	t.sendAll("server", msg)
}

func makeGroupMsgSender(sender msgSender, members []string) *groupMsgSender {
	return &groupMsgSender{msgSender: sender, members: members}
}

func main() {
	var host string
	var port int
	flag.StringVar(&host, "host", "localhost", "host")
	flag.IntVar(&port, "port", 65434, "port")
	flag.Parse()

	log.Printf("Server running on %s:%d\n", host, port)

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	pb.RegisterMafiaServer(srv, MakeMafiaServer())
	srv.Serve(listener)
}
