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

const (
	MaxUsernameLength = 15
)

type connection struct {
	ch    chan *pb.LoginResponse
	close chan struct{}
}

type mafiaServer struct {
	pb.UnimplementedMafiaServer

	playerToSession   map[string]*gameSession
	playerToConnecton map[string]*connection

	waitingListMutex sync.Mutex
	waitingList      []string

	eventSender
}

func (t *mafiaServer) init() {
	// TO DO: run garbage collector
}

func (t *mafiaServer) disconnect(username string) {
	delete(t.playerToConnecton, username)
}

func (t *mafiaServer) send(receiver string, event *pb.LoginResponse) {
	// log.Printf("sending \"%s\" from %s to %s\n", msg, sender, receiver)
	conn, ok := t.playerToConnecton[receiver]
	if !ok {
		panic(fmt.Sprintf("user %s not connected.", receiver))
	}
	select {
	case conn.ch <- event:
		// noop
	default:
		log.Printf("%s user msg buffer is full", receiver)
		t.disconnect(receiver)
	}
}

func (t *mafiaServer) sendPhaseChange(receiver string, newPhase mafia.GamePhaseType) {
	log.Printf("sending phase change \"%s\" to %s\n", newPhase, receiver)
	t.send(receiver, &pb.LoginResponse{
		Event: &pb.LoginResponse_PhaseChange_{
			PhaseChange: &pb.LoginResponse_PhaseChange{NewPhase: string(newPhase)},
		},
	})
}

func (t *mafiaServer) sendRoleAssignment(receiver string, role mafia.Role) {
	log.Printf("sending role assignment \"%s\" to %s\n", role, receiver)
	t.send(receiver, &pb.LoginResponse{
		Event: &pb.LoginResponse_RoleAssignment_{
			RoleAssignment: &pb.LoginResponse_RoleAssignment{Role: string(role)},
		},
	})
}

func (t *mafiaServer) sendMsg(sender string, receiver string, msg string) {
	log.Printf("sending msg \"%s\" from %s to %s\n", msg, sender, receiver)
	text := fmt.Sprintf("[%s] %s", sender, msg)
	t.send(receiver, &pb.LoginResponse{
		Event: &pb.LoginResponse_NewMessage_{
			NewMessage: &pb.LoginResponse_NewMessage{Text: text},
		},
	})
}

func (t *mafiaServer) sendMsgFromServer(receiver string, msg string) {
	t.sendMsg("server", receiver, msg)
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
		t.waitingList = nil
	} else if len(t.waitingList) < mafia.PlayersInGame {
		needPlayers := mafia.PlayersInGame - len(t.waitingList)

		log.Printf("need %d more players for game", needPlayers)

		sender := makeGroupMsgSender(t, t.waitingList)
		sender.sendAllMsgByServer(fmt.Sprintf("%s connected", username))
		sender.sendAllMsgByServer(
			fmt.Sprintf("Waiting for players. %d more players are needed.", needPlayers),
		)

	} else {
		panic("waiting list too long")
	}
}

func (t *mafiaServer) Login(request *mafia.LoginRequest, serv mafia.Mafia_LoginServer) error {
	username := extractUsername(serv.Context())

	if len(username) > MaxUsernameLength {
		return status.Errorf(codes.AlreadyExists, "Username is too long. Max length is %d.", MaxUsernameLength)
	}

	if _, ok := t.playerToConnecton[username]; ok {
		return status.Errorf(codes.AlreadyExists, "User with this name already logged in.")
	}

	close := make(chan struct{})
	ch := make(chan *pb.LoginResponse, 100)

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
			case resp := <-ch:
				err := serv.Send(resp)

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

	t.sendMsgFromServer(username, "You are connected to server.")

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

func (t *mafiaServer) GetAlivePlayers(ctx context.Context, request *pb.GetAlivePlayersRequest) (*pb.GetAlivePlayersResponse, error) {
	username := extractUsername(ctx)
	session, err := t.GetGameSession(username)
	if err != nil {
		return nil, err
	}
	return &pb.GetAlivePlayersResponse{AlivePlayers: session.GetAlivePlayers()}, nil
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

type eventSender interface {
	send(receiver string, event *pb.LoginResponse)
	sendPhaseChange(receiver string, newPhase mafia.GamePhaseType)
	sendRoleAssignment(receiver string, role mafia.Role)
	sendMsg(sender string, receiver string, msg string)
	sendMsgFromServer(receiver string, msg string)
}

type groupEventSender struct {
	eventSender
	members []string
}

func (t *groupEventSender) sendMsgAll(sender string, msg string) {
	for _, player := range t.members {
		t.sendMsg(sender, player, msg)
	}
}

func (t *groupEventSender) sendAllPhaseChange(newPhase mafia.GamePhaseType) {
	for _, player := range t.members {
		t.sendPhaseChange(player, newPhase)
	}
}

func (t *groupEventSender) sendAllMsgByServer(msg string) {
	t.sendMsgAll("server", msg)
}

func makeGroupMsgSender(sender eventSender, members []string) *groupEventSender {
	return &groupEventSender{eventSender: sender, members: members}
}

func main() {
	var host string
	var port int
	flag.StringVar(&host, "host", "", "host")
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
