package main

import (
	"fmt"
	"time"
)

type GamePhaseType string

const (
	GamePhaseTypeDay   GamePhaseType = "day"
	GamePhaseTypeNight GamePhaseType = "night"
)

type gamePhase interface {
	VoteAgainst(voter string, target string) error
	Kill(mafia string, victim string) error
	EndTurn(player string) error
	Check(commissar string, suspect string) (bool, error)
	PublishCheckResult(commissar string) error
	GetGameSession() *gameSession
	OnFinish()
	GetName() string
	GetPhaseDurationInSeconds() int
}

type baseGamePhase struct {
	name        string
	duration    int
	gameSession *gameSession
}

func (t *baseGamePhase) GetGameSession() *gameSession {
	return t.gameSession
}

func (t *baseGamePhase) GetPhaseName() string {
	return t.name
}

func (t *baseGamePhase) GetPhaseDurationInSeconds() int {
	return t.duration
}

func RunGamePhase(phase gamePhase) {
	go func() {
		phaseEnd := time.Now().Add(time.Duration(phase.GetPhaseDurationInSeconds()) * time.Second)

		ticker := time.NewTicker(time.Second * 5)

		for currentTime := range ticker.C {
			timeLeft := phaseEnd.Sub(currentTime)

			if timeLeft < time.Millisecond {
				ticker.Stop()
				phase.OnFinish()
				break
			}

			phase.GetGameSession().broadcast(
				fmt.Sprintf("%s will end in %d secs", phase.GetName(), int(timeLeft.Seconds()+1)),
			)
		}
	}()
}

type gamePhaseDay struct {
	votesAgainstCount map[string]int
	playerTarget        map[string]string
}

func (t *gamePhaseDay) VoteAgainst(voter string, target string) error {
	if currentTarget, ok := t.playerTarget[voter]; ok {
		t.votesAgainstCount[currentTarget]--
	}
	t.playerTarget[voter] = target
	t.votesAgainstCount[target]++
	
}

// Kill(mafia string, victim string) error
// EndTurn(player string) error
// Check(commissar string, suspect string) (bool, error)
// PublishCheckResult(commissar string) error
// GetGameSession() *gameSession
// OnFinish()
// GetName() string
// GetPhaseDurationInSeconds() int

func makeGamePhaseDay() *gamePhaseDay {

}

type gamePhaseNight struct {
}

func makeGamePhaseNight() *gamePhaseNigh {

}

func MakeGamePhase(gamePhaseType GamePhaseType) gamePhase {
	switch gamePhaseType {
	case GamePhaseTypeDay:
		return makeGamePhaseDay()
	case GamePhaseTypeNight:
		return makeGamePhaseNight()
	}
}
