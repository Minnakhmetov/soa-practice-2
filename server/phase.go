package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/Minnakhmetov/soa-practice-2/mafia"
)

type gamePhase interface {
	VoteAgainst(player string, target string)
	Shoot(player string, target string)
	EndTurn(player string)
	Check(player string, target string)
	PublishCheckResult(player string)
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

			phase.GetGameSession().sendAllServerMessage(
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
	playerVote        map[string]string
	endedTurn         map[string]struct{}
}

func (t *gamePhaseDay) playerEndedTurn(player string) bool {
	_, ok := t.endedTurn[player]
	return ok
}

func (t *gamePhaseDay) VoteAgainst(player string, target string) {
	if t.playerEndedTurn(player) {
		t.GetGameSession().sendServerMessage(player, "You can't vote after you ended your turn.")
		return
	}
	if player == target {
		t.GetGameSession().sendServerMessage(player, "You can't vote against yourself.")
		return
	}
	if currentTarget, ok := t.playerVote[player]; ok {
		t.votesAgainstCount[currentTarget]--
	}
	t.playerVote[player] = target
	t.votesAgainstCount[target]++
	t.GetGameSession().sendServerMessage(player, fmt.Sprintf("Your vote is against %s now.", target))
}

func (t *gamePhaseDay) Shoot(player string, target string) {
	t.GetGameSession().sendServerMessage(player, "You can't shoot during the day")
}

func (t *gamePhaseDay) EndTurn(player string) {
	if _, ok := t.endedTurn[player]; ok {
		t.GetGameSession().sendServerMessage(player, "You already finished your turn.")
		return
	}
	t.endedTurn[player] = struct{}{}
	if len(t.endedTurn) == t.GetGameSession().GetPlayerCount() {
		t.GetFinishChannel() <- struct{}{}
	}
	t.GetGameSession().sendAllServerMessage(fmt.Sprintf("%s finished their turn.", player))
}

func (t *gamePhaseDay) OnFinish() {
	t.GetGameSession().sendAllServerMessage("The day is over.")

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

	addLineToResult("")
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
	}

	t.GetGameSession().sendAllServerMessage(strings.Join(resultTableLines, "\n"))

	if !multipleWinners {
		t.GetGameSession().kill(mostVoted)
	}

	t.GetGameSession().enqueuePhase(MakeGamePhaseNight(t.GetGameSession()))
}

func (t *gamePhaseDay) Check(player string, target string) {
	t.GetGameSession().sendServerMessage(player, "You can't check during the day.")
}

func (t *gamePhaseDay) PublishCheckResult(player string) {
	if t.GetGameSession().GetPlayerRole(player) != mafia.RoleCommisar {
		t.GetGameSession().sendServerMessage(player, "Only commissar can publish check results.")
		return
	}
	msg := fmt.Sprintf("For now I know that following users are mafiosi: %s", strings.Join(t.GetGameSession().uncoveredMafia, ", "))
	t.GetGameSession().sendAll("commissar", msg)
}

func MakeGamePhaseDay(session *gameSession) *gamePhaseDay {
	return &gamePhaseDay{
		baseGamePhase: baseGamePhase{
			name:        "day",
			duration:    mafia.PhaseDurationInSecs[mafia.GamePhaseTypeDay],
			gameSession: session,
			finish:      make(chan struct{}),
		},
		votesAgainstCount: map[string]int{},
		playerVote:        map[string]string{},
		endedTurn:         map[string]struct{}{},
	}
}

type gamePhaseNight struct {
	baseGamePhase
	checkUsed   bool
	shotByMafia string
}

func (t *gamePhaseNight) VoteAgainst(player string, target string) {
	t.GetGameSession().sendServerMessage(player, "You can't vote during the night.")
}

func (t *gamePhaseNight) EndTurn(player string) {
	t.GetGameSession().sendServerMessage(player, "You can't end turn during the night.")
}

func (t *gamePhaseNight) Check(player string, target string) {
	if t.GetGameSession().GetPlayerRole(player) != mafia.RoleCommisar {
		t.GetGameSession().sendServerMessage(player, "Only commissar can do a check.")
		return
	}

	if t.checkUsed {
		t.GetGameSession().sendServerMessage(player, "You already did a check tonight.")
		return
	}

	t.checkUsed = true

	if t.GetGameSession().GetPlayerRole(target) == mafia.RoleMafia {
		t.GetGameSession().AddUncoveredMafia(target)
		t.GetGameSession().sendServerMessage(player, fmt.Sprintf("%s is a mafioso. Good job!", target))
	} else {
		t.GetGameSession().sendServerMessage(player, fmt.Sprintf("%s is not a mafioso. Keep trying!", target))
	}
}

func (t *gamePhaseNight) nightRolesDone() bool {
	return t.checkUsed && t.shotByMafia != ""
}

func (t *gamePhaseNight) Shoot(player string, target string) {
	if t.GetGameSession().GetPlayerRole(player) != mafia.RoleMafia {
		t.GetGameSession().sendServerMessage(player, "Only mafiosi can shoot.")
		return
	}
	if t.shotByMafia != "" {
		t.GetGameSession().sendServerMessage(player, "You already shot someone tonight.")
		return
	}
	if t.GetGameSession().IsDead(target) {
		t.GetGameSession().sendServerMessage(player, "You can't shot dead people.")
		return
	}
	t.shotByMafia = target
	if t.nightRolesDone() {
		t.GetFinishChannel() <- struct{}{}
	}
}

func (t *gamePhaseNight) PublishCheckResult(player string) {
	t.GetGameSession().sendServerMessage(player, "You can't publish check result during the night.")
}

func (t *gamePhaseNight) OnFinish() {
	t.GetGameSession().sendAllServerMessage("The night is over.")
	if t.shotByMafia != "" {
		t.GetGameSession().sendAllServerMessage("Mafia didn't shoot anyone today.")
	} else {
		t.GetGameSession().sendAllServerMessage(fmt.Sprintf("Today %s was shot by mafiosi.", t.shotByMafia))
	}
	t.GetGameSession().enqueuePhase(MakeGamePhaseDay(t.GetGameSession()))
}

func MakeGamePhaseNight(session *gameSession) *gamePhaseNight {
	return &gamePhaseNight{
		baseGamePhase: baseGamePhase{
			name:        "night",
			duration:    mafia.PhaseDurationInSecs[mafia.GamePhaseTypeNight],
			gameSession: session,
			finish:      make(chan struct{}),
		},
		checkUsed:   false,
		shotByMafia: "",
	}
}
