package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Minnakhmetov/soa-practice-2/mafia"
)

type player struct {
	username string
	role     mafia.Role
	isDead   bool
}

type GameSession struct {
	players            []player
	isFinished         bool
	aliveMafiaCount    int
	aliveNonMafiaCount int
	msgs               chan string
	finish             chan struct{}
	phase              mafia.GamePhase
	phaseChange        chan struct{}
}

func (t *GameSession) mafiaWon() bool {
	return t.aliveMafiaCount >= t.aliveNonMafiaCount
}

func (t *GameSession) citizensWon() bool {
	return t.aliveMafiaCount == 0
}

func MakeGameSession(usernames []string) *GameSession {
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

	msgs := make(chan string)
	finish := make(chan struct{})

	session := GameSession{players: players, msgs: msgs, finish: finish}

	return &session
}

func (t *GameSession) Start() {
	go func() {
		for _ = range t.phaseChange {
			// if t.mafiaWon() {
			// 	session
			// }

			phaseDuration := mafia.PhaseDurationInSecs[t.phase]
			phaseEnd := time.Now().Add(time.Duration(phaseDuration) * time.Second)

			ticker := time.NewTicker(time.Second * 5)

			for currentTime := range ticker.C {
				timeLeft := phaseEnd.Sub(currentTime)
				if timeLeft < time.Millisecond {

				}
				t.msgs <- fmt.Sprintf("Phase will end in %d secs", int(timeLeft.Seconds()+1))
			}

		}
	}()
}

func (t *GameSession) Execute(executor string, executee string) error {

}

func (t *GameSession) Kill(killer string, victim string) error {

}
