package main

import (
	"fmt"
	"math/rand"

	"github.com/Minnakhmetov/soa-practice-2/mafia"
)

type playerInfo struct {
	role   mafia.Role
	isDead bool
	msgs   chan string
}

type gameSession struct {
	playerInfo         map[string]*playerInfo
	isFinished         bool
	aliveMafiaCount    int
	aliveNonMafiaCount int
	finish             chan struct{}
	nextPhase          chan gamePhase
}

func (t *gameSession) broadcast(msg string) {
	// TO DO: send msg to all players
}

func (t *gameSession) sendAll(sender string, msg string) {
	t.broadcast(fmt.Sprintf("[%s] %s", sender, msg))
}

func (t *gameSession) kill(username string) {
	t.playerInfo[username].isDead = true
	t.broadcast(fmt.Sprintf("%s died.", username))
}

func (t *gameSession) mafiaWon() bool {
	return t.aliveMafiaCount >= t.aliveNonMafiaCount
}

func (t *gameSession) citizensWon() bool {
	return t.aliveMafiaCount == 0
}

func (t *gameSession) GetPlayerCount() int {
	return len(t.playerInfo)
}

func (t *gameSession) GetPlayerRole(username string) mafia.Role {
	return t.playerInfo[username].role
}

func MakeGameSession(usernames []string) *gameSession {
	var roles []mafia.Role

	for role, count := range mafia.RoleToCount {
		for i := 0; i < count; i++ {
			roles = append(roles, role)
		}
	}

	if len(roles) != len(usernames) {
		panic("Number of roles should be equal to number of players in session")
	}

	rand.Shuffle(len(roles), func(i int, j int) {
		roles[i], roles[j] = roles[j], roles[i]
	})

	session := gameSession{
		playerInfo:         map[string]*playerInfo{},
		isFinished:         false,
		aliveMafiaCount:    mafia.RoleToCount[mafia.RoleMafia],
		aliveNonMafiaCount: mafia.PlayersInGame - mafia.RoleToCount[mafia.RoleMafia],
		finish:             make(chan struct{}),
	}

	var players []playerInfo

	for i := 0; i < len(roles); i++ {
		players = append(players, playerInfo{
			username: usernames[i],
			role:     roles[i],
			isDead:   false,
		})
	}

	return &session
}

func (t *gameSession) Run() {
	go func() {
		for phase := range t.nextPhase {
			RunGamePhase(phase)
		}
	}()
	// t.nextPhase <- makeGamePhaseDay()
}

// func (t *GameSession) Execute(executor string, executee string) error {

// }

// func (t *GameSession) Kill(killer string, victim string) error {

// }
