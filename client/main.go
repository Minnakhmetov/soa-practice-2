package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"

	"github.com/Minnakhmetov/soa-practice-2/mafia"
	pb "github.com/Minnakhmetov/soa-practice-2/mafia"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

// type mafiaClient struct {
// 	pbMafiaClient pb.MafiaClient
// }

type commandInfo struct {
	argsCount   int
	example     string
	description string
}

var commandInfoByName = map[string]commandInfo{
	"/endturn": {argsCount: 0, example: "/endturn", description: "end your turn after voting during the day "},
	"/check":   {argsCount: 1, example: "/check target", description: "if you are a commissar, check if another player is mafioso during the night"},
	"/publish": {argsCount: 0, example: "/publish", description: "if you are a commissar, publish check results during the day"},
	"/shoot":   {argsCount: 1, example: "/shoot target", description: "if you are a mafioso, shoot another player during the night"},
	"/vote":    {argsCount: 1, example: "/vote target", description: "vote against another player during the day"},
	"/help":    {argsCount: 0, example: "/help", description: "show help list"},
}

func getContextWithUsername(username string) context.Context {
	return metadata.AppendToOutgoingContext(context.Background(), "username", username)
}

func getCommandUsage(command string) string {
	info := commandInfoByName[command]
	return fmt.Sprintf("%s   \t%s\n\t\texample: %s", command, info.description, info.example)
}

func getHelpList() string {
	var lines []string

	addLine := func(newLine string) {
		lines = append(lines, newLine)
	}

	addLine("")
	addLine("type message for other players or one of commands below and press enter:")

	for command := range commandInfoByName {
		addLine(getCommandUsage(command))
	}

	return strings.Join(lines, "\n\n")
}

type mafiaClient struct {
	username         string
	protoMafiaClient pb.MafiaClient
}

func (t *mafiaClient) Shoot(target string) error {
	_, err := t.protoMafiaClient.Shoot(getContextWithUsername(t.username), &pb.ShootRequest{Target: target})
	return err
}

func (t *mafiaClient) VoteAgainst(target string) error {
	_, err := t.protoMafiaClient.VoteAgainst(getContextWithUsername(t.username), &pb.VoteAgainstRequest{Target: target})
	return err
}

func (t *mafiaClient) Check(target string) error {
	_, err := t.protoMafiaClient.Check(getContextWithUsername(t.username), &pb.CheckRequest{Target: target})
	return err
}

func (t *mafiaClient) EndTurn() error {
	_, err := t.protoMafiaClient.EndTurn(getContextWithUsername(t.username), &pb.EndTurnRequest{})
	return err
}

func (t *mafiaClient) PublishCheckResult() error {
	_, err := t.protoMafiaClient.PublishCheckResult(getContextWithUsername(t.username), &pb.PublishCheckResultRequest{})
	return err
}

func (t *mafiaClient) GetAlivePlayers() ([]string, error) {
	response, err := t.protoMafiaClient.GetAlivePlayers(getContextWithUsername(t.username), &pb.GetAlivePlayersRequest{})
	if err != nil {
		return nil, err
	}
	return response.AlivePlayers, nil
}

func (t *mafiaClient) Login() (pb.Mafia_LoginClient, error) {
	return t.protoMafiaClient.Login(getContextWithUsername(t.username), &pb.LoginRequest{})
}

func main() {
	var host string
	var port int
	var username string
	var autoMode bool
	var redisPassword string

	flag.StringVar(&host, "host", "localhost", "server host")
	flag.IntVar(&port, "port", 65434, "server port")
	flag.StringVar(&username, "username", "", "username")
	flag.StringVar(&redisPassword, "rpass", "changeme", "redis password")
	flag.BoolVar(&autoMode, "auto", false, "run a bot")
	flag.Parse()

	if username == "" {
		panic("username can't be empty.")
	}

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", host, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := &mafiaClient{username: username, protoMafiaClient: pb.NewMafiaClient(conn)}

	var role mafia.Role

	doRandomAction := func(phase mafia.GamePhaseType) {
		alivePlayers, err := client.GetAlivePlayers()
		if err != nil {
			panic(err)
		}
		println("doing random action")
		if phase == mafia.GamePhaseTypeDay {
			if role == mafia.RoleCommisar {
				client.PublishCheckResult()
			}
			client.VoteAgainst(alivePlayers[rand.Intn(len(alivePlayers))])
			client.EndTurn()
		} else {
			if role == mafia.RoleMafia {
				client.Shoot(alivePlayers[rand.Intn(len(alivePlayers))])
			} else if role == mafia.RoleCommisar {
				client.Check(alivePlayers[rand.Intn(len(alivePlayers))])
			}
		}
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: redisPassword,
		DB:       0,
	})

	newRedisChannel := make(chan string)
	var currentRedisChannel string = "lobby"
	var pubsub *redis.PubSub = redisClient.Subscribe(context.Background(), currentRedisChannel)

	go func() {
		listenCurrentChannel := func() {
			for {
				select {
				case newChannel := <-newRedisChannel:
					currentRedisChannel = newChannel
					pubsub.Close()
					pubsub = redisClient.Subscribe(context.Background(), newChannel)
					return

				case msg := <-pubsub.Channel():
					println(msg.Payload)
				}
			}
		}
		for {
			listenCurrentChannel()
		}
	}()

	go func() {
		stream, err := client.Login()
		if err != nil {
			panic(err)
		}
		for {
			for {
				event, err := stream.Recv()
				if err == io.EOF {
					break
				}
				if err != nil {
					log.Fatal(err)
				}
				switch e := event.GetEvent().(type) {
				case *pb.LoginResponse_NewMessage_:
					println(e.NewMessage.Text)
				case *pb.LoginResponse_PhaseChange_:
					println("got phase change")
					if autoMode {
						doRandomAction(pb.GamePhaseType(e.PhaseChange.NewPhase))
					}
				case *pb.LoginResponse_RoleAssignment_:
					role = pb.Role(e.RoleAssignment.GetRole())
				case *pb.LoginResponse_NewBrokerChannel_:
					newRedisChannel <- e.NewBrokerChannel.Name
				}
			}
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			println("type /help for info")
			continue
		}

		if line[0] == '/' {
			tokens := strings.Split(line, " ")

			info, ok := commandInfoByName[tokens[0]]
			if !ok {
				println("command not found, type /help for info")
			}
			if len(tokens) != info.argsCount+1 {
				println("too many args, type /help for info")
			}

			var err error

			switch tokens[0] {
			case "/endturn":
				err = client.EndTurn()
			case "/check":
				err = client.Check(tokens[1])
			case "/vote":
				err = client.VoteAgainst(tokens[1])
			case "/shoot":
				err = client.Shoot(tokens[1])
			case "/publish":
				err = client.PublishCheckResult()
			case "/help":
				println(getHelpList())
			}

			if err != nil {
				panic(err)
			}
		} else {
			if currentRedisChannel != "" {
				redisClient.Publish(context.Background(), currentRedisChannel, fmt.Sprintf("[%s] %s", username, line))
			}
		}
	}

}
