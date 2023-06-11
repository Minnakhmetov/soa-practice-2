package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	pb "github.com/Minnakhmetov/soa-practice-2/mafia"
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
	for command := range commandInfoByName {
		lines = append(lines, getCommandUsage(command))
	}
	return strings.Join(lines, "\n\n")
}

func main() {
	var host string
	var port int
	var username string

	flag.StringVar(&host, "host", "localhost", "server host")
	flag.IntVar(&port, "port", 65434, "server port")
	flag.StringVar(&username, "username", "", "username")
	flag.Parse()

	if username == "" {
		panic("username can't be empty.")
	}

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", host, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewMafiaClient(conn)

	// stream, err := client.Login(metadata.AppendToOutgoingContext(context.Background(), "username", username), &pb.LoginRequest{})
	// event, err := stream.Recv()
	// if err != nil {
	// 	print(err.Error())
	// }
	// print(event.GetMessage())
	// return

	go func() {
		stream, err := client.Login(getContextWithUsername(username), &pb.LoginRequest{})
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
				println(event.GetMessage())
			}
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {

		} else if line[0] == '/' {
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
				_, err = client.EndTurn(getContextWithUsername(username), &pb.EndTurnRequest{})
			case "/check":
				_, err = client.Check(getContextWithUsername(username), &pb.CheckRequest{Target: tokens[1]})
			case "/vote":
				_, err = client.VoteAgainst(getContextWithUsername(username), &pb.VoteAgainstRequest{Target: tokens[1]})
			case "/shoot":
				_, err = client.Shoot(getContextWithUsername(username), &pb.ShootRequest{Target: tokens[1]})
			case "/publish":
				_, err = client.PublishCheckResult(getContextWithUsername(username), &pb.PublishCheckResultRequest{})
			case "/help":
				println(getHelpList())
			}

			if err != nil {
				panic(err)
			}
		}
	}

}
