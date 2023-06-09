package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/Minnakhmetov/soa-practice-2/mafia"
)

type gamePhase interface {
	VoteAgainst(player string, target string) error
	Shoot(player string, target string) error
	EndTurn(player string) error
	Check(player string, target string) (bool, error)
	PublishCheckResult(player string) error
	GetGameSession() *gameSession
	OnFinish()
	GetName() string
	GetPhaseDurationInSeconds() int
	GetFinishChannel() chan struct{}
}

type baseGamePhase struct {
	name        string
	duration    int
	gameSession *gameSession
	finish      chan struct{}
}

func (t *baseGamePhase) GetGameSession() *gameSession {
	return t.gameSession
}

func (t *baseGamePhase) GetName() string {
	return t.name
}

func (t *baseGamePhase) GetPhaseDurationInSeconds() int {
	return t.duration
}

func (t *baseGamePhase) GetFinishChannel() chan struct{} {
	return t.finish
}

func RunGamePhase(phase gamePhase) {
	go func() {
		phaseEnd := time.Now().Add(time.Duration(phase.GetPhaseDurationInSeconds()) * time.Second)

		ticker := time.NewTicker(time.Second * 5)

		select {
		case currentTime := <-ticker.C:
			timeLeft := phaseEnd.Sub(currentTime)

			if timeLeft < time.Millisecond {
				ticker.Stop()
				phase.OnFinish()
				break
			}

			phase.GetGameSession().broadcast(
				fmt.Sprintf("%s will end in %d secs", phase.GetName(), int(timeLeft.Seconds()+1)),
			)
		case <-phase.GetFinishChannel():
			phase.OnFinish()
		}
	}()
}

type gamePhaseDay struct {
	baseGamePhase
	votesAgainstCount map[string]int
	playerVote      map[string]string
	endedTurn         map[string]struct{}
	uncoveredMafia    []string
}

func (t *gamePhaseDay) VoteAgainst(voter string, target string) error {
	if voter == target {
		return errors.New("can't vote against yourself")
	}
	if currentTarget, ok := t.playerVote[voter]; ok {
		t.votesAgainstCount[currentTarget]--
	}
	t.playerVote[voter] = target
	t.votesAgainstCount[target]++
	return nil
}

func (t *gamePhaseDay) Shoot(string, string) error {
	return errors.New("can't shoot during the day")
}

func (t *gamePhaseDay) EndTurn(player string) {
	t.endedTurn[player] = struct{}{}
	if len(t.endedTurn) == t.GetGameSession().GetPlayerCount() {
		t.GetFinishChannel() <- struct{}{}
	}
}

func (t *gamePhaseDay) OnFinish() {
	voteResults := make([]string, t.GetGameSession().GetPlayerCount())
	maxVoteCount := 0

	for player, votes := range t.votesAgainstCount {
		voteResults = append(voteResults, fmt.Sprintf("%s: %d", player, votes))
		if maxVoteCount < votes {
			maxVoteCount = votes
		}
	}

	var resultTableLines []string

	addLineToResult := func(line string) {
		resultTableLines = append(resultTableLines, line)
	}

	addLineToResult("Poll result:")
	addLineToResult("username\tvotes against")

	for player, votes := range t.votesAgainstCount {
		addLineToResult(fmt.Sprintf("%s: %d", player, votes))
	}

	mostVoted := ""
	multipleWinners := false

	for player, votes := range t.votesAgainstCount {
		if votes == maxVoteCount {
			if mostVoted != "" {
				multipleWinners = true
				break
			}
			mostVoted = player
		}
	}

	if multipleWinners {
		addLineToResult("There are multiple players with max number of votes so no one will die today.")
	} else {
		addLineToResult(fmt.Sprintf("%[1]s got max number of votes. It's time to go, %[1]s.", mostVoted))
		t.GetGameSession().kill(mostVoted)
	}
}

func (t *gamePhaseDay) Check(string, string) (bool, error) {
	return false, errors.New("Can't check during the day.")
}

func (t *gamePhaseDay) PublishCheckResult(player string) error {
	if t.GetGameSession().GetPlayerRole(player) != mafia.RoleCommisar {
		return errors.New("Only commissar can publish check results.")
	}
	t.GetGameSession().sendAll("[commissar]", fmt.Sprintf("uncovered mafia: %s"))
	return nil
}

// Kill(mafia string, victim string) error
// EndTurn(player string) error
// Check(commissar string, suspect string) (bool, error)
// PublishCheckResult(commissar string) error
// GetGameSession() *gameSession
// OnFinish()
// GetName() string
// GetPhaseDurationInSeconds() int

func makeGamePhaseDay(session *gameSession) *gamePhaseDay {
	return &gamePhaseDay{
		baseGamePhase: baseGamePhase{
			name:        "day",
			duration:    mafia.PhaseDurationInSecs[mafia.GamePhaseTypeDay],
			gameSession: session,
			finish:      make(chan struct{}),
		},
		votesAgainstCount: map[string]int{},
		playerVote:      map[string]string{},
		endedTurn:         map[string]struct{}{},
	}
}

type gamePhaseNight struct {
	
}

// func makeGamePhaseNight() *gamePhaseNigh {

// }

// func MakeGamePhase(gamePhaseType GamePhaseType) gamePhase {
// 	switch gamePhaseType {
// 	case GamePhaseTypeDay:
// 		return makeGamePhaseDay()
// 	case GamePhaseTypeNight:
// 		return makeGamePhaseNight()
// 	}
// }
