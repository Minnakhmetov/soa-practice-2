package main

import (
	"math/rand"

	"github.com/Minnakhmetov/soa-practice-2/mafia"
)

type player struct {
	username string
	role     mafia.Role
	isDead   bool
	msgs     chan string
}

type gameSession struct {
	players            []player
	isFinished         bool
	aliveMafiaCount    int
	aliveNonMafiaCount int
	finish             chan struct{}
	phase              mafia.GamePhase
	phaseChange        chan struct{}
}

func (t *gameSession) broadcast(msg string) {
	// TO DO: send msg to all players
}

func (t *gameSession) mafiaWon() bool {
	return t.aliveMafiaCount >= t.aliveNonMafiaCount
}

func (t *gameSession) citizensWon() bool {
	return t.aliveMafiaCount == 0
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

	var players []player

	for i := 0; i < len(roles); i++ {
		players = append(players, player{
			username: usernames[i],
			role:     roles[i],
			isDead:   false,
		})
	}

	finish := make(chan struct{})

	session := gameSession{players: players, finish: finish}

	return &session
}

func (t *gameSession) startDay() {

}

func (t *gameSession) startNight() {

}

func (t *gameSession) Start() {
	
	for {

	}
}

// func (t *GameSession) Execute(executor string, executee string) error {

// }

// func (t *GameSession) Kill(killer string, victim string) error {

// }
