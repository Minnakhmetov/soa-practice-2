package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"net"
// 	"sync"

// 	"github.com/Minnakhmetov/soa-practice-2/mafia"
// 	pb "github.com/Minnakhmetov/soa-practice-2/mafia"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/codes"
// 	"google.golang.org/grpc/status"
// )

// type connection struct {
// 	ch    chan string
// 	close chan struct{}
// }

// type mafiaServer struct {
// 	pb.UnimplementedMafiaServer

// 	playerToSession   map[string]*GameSession
// 	playerToConnecton map[string]*connection

// 	waitingPlayersMutex sync.Mutex
// 	waitingPlayers      []string
// }

// func (t *mafiaServer) init() {
// 	// TO DO: run garbage collector
// }

// func (t *mafiaServer) broadcast(usernames []string, msg string) {
// 	for _, username := range usernames {
// 		select {
// 		case t.playerToConnecton[username].ch <- msg:
// 			// noop
// 		default:
// 			log.Printf("Couldn't send msg to %s\n", username)
// 		}
// 	}
// }

// func (t *mafiaServer) disconnect(username string) {
// 	delete(t.playerToConnecton, username)
// }

// func (t *mafiaServer) startGameSession(usernames []string) {
// 	session := MakeGameSession(usernames)

// 	for _, username := range usernames {
// 		t.playerToSession[username] = session
// 	}

// 	// go func() {
// 	// 	for {
// 	// 		select {
// 	// 		case msg := <-session.msgs:
// 	// 			t.broadcast(usernames, msg)
// 	// 		}
// 	// 	}
// 	// }()

// 	session.Start()
// }

// func (t *mafiaServer) inActiveGameSession(username string) bool {
// 	_, ok := t.playerToSession[username]
// 	return ok
// }

// func (t *mafiaServer) addWaitingPlayer(username string) {
// 	t.waitingPlayersMutex.Lock()
// 	defer t.waitingPlayersMutex.Unlock()

// 	t.waitingPlayers = append(t.waitingPlayers, username)
// 	if len(t.waitingPlayers) == mafia.PlayersRequired {
// 		t.startGameSession(t.waitingPlayers)
// 	}
// }

// func (t *mafiaServer) Login(request *mafia.LoginRequest, serv mafia.Mafia_LoginServer) error {
// 	username := extractUsername(serv.Context())

// 	if _, ok := t.playerToConnecton[username]; ok {
// 		return status.Errorf(codes.AlreadyExists, "User with this name already logged in")
// 	}

// 	close := make(chan struct{})
// 	ch := make(chan string, 50)

// 	t.playerToConnecton[username] = &connection{ch: ch, close: close}

// 	wg := sync.WaitGroup{}
// 	defer wg.Wait()
// 	wg.Add(1)

// 	go func() {
// 		select {
// 		case <-close:
// 			t.disconnect(username)
// 		case msg := <-ch:
// 			err := serv.Send(&pb.Event{Message: msg})
// 			if err != nil {
// 				log.Printf("error when sending message to %s: %s\n", username, err.Error())
// 				t.disconnect(username)
// 			}
// 		}
// 		wg.Done()
// 	}()

// 	// TO DO: check if player already in active game session (reconnected)

// 	t.addWaitingPlayer(username)
// }

// func (t *mafiaServer) ExecutePlayer(ctx context.Context, request *pb.ExecutePlayerRequest) (*pb.ExecutePlayerResponse, error) {
// 	username := extractUsername(ctx)
// 	if _, ok := t.playerToConnecton[username]; ok {
// 		return nil, status.Errorf(codes.Unauthenticated, "User is not logged in")
// 	}
// 	session, ok := t.playerToSession[username]
// 	if ok {
// 		return nil, status.Errorf(codes.NotFound, "User is not in game session")
// 	}

// 	if session.phase != mafia.GamePhaseDay {
// 		return nil, status.Errorf(codes.FailedPrecondition, "Execution can be performed only during the day")
// 	}

// }

// func extractUsername(context context.Context) string {
// 	return context.Value(context).(string)
// }

// // func (t *mafiaServer) Login(context context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
// // 	t.waitingPlayersMutex.Lock()
// // 	defer t.waitingPlayersMutex.Unlock()
// // 	t.playerToConnecton[request.GetUsername()] = connection{

// // 	}
// // 	// if len(t.waitingPlayers) ==
// // }

// // func (t *server) Do(context context.Context, request *pb.Request) (*pb.Response, error) {
// // 	return &pb.Response{Message: fmt.Sprintf("hi %s", request.GetMessage()), Z: request.GetX() * request.GetY()}, nil
// // }

// // const port = 65434

// // func main() {
// // 	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
// // 	if err != nil {
// // 		panic(err)
// // 	}
// // 	srv := grpc.NewServer()
// // 	pb.RegisterReverseServer(srv, &mafiaServer{})
// // 	srv.Serve(listener)
// // }
