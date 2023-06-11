package main

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/Minnakhmetov/soa-practice-2/mafia"
)

type gamePhase interface {
	OnStart()
	VoteAgainst(player string, target string)
	Shoot(player string, target string)
	EndTurn(player string)
	Check(player string, target string)
	PublishCheckResult(player string)
	GetGameSession() *gameSession
	OnFinish()
	GetName() string
	GetType() mafia.GamePhaseType
	GetPhaseDuration() time.Duration
	GetFinishChannel() chan struct{}
	GetRedisChannel(player string) string
}

type baseGamePhase struct {
	name               string
	phaseType          mafia.GamePhaseType
	duration           time.Duration
	gameSession        *gameSession
	finish             chan struct{}
	redisChannelByRole map[mafia.Role]string
	groupEventSender
}

func (t *baseGamePhase) GetRedisChannel(player string) string {
	return t.redisChannelByRole[t.GetGameSession().GetPlayerRole(player)]
}

func (t *baseGamePhase) GetType() mafia.GamePhaseType {
	return t.phaseType
}

func (t *baseGamePhase) GetGameSession() *gameSession {
	return t.gameSession
}

func (t *baseGamePhase) GetName() string {
	return t.name
}

func (t *baseGamePhase) GetPhaseDuration() time.Duration {
	return t.duration
}

func (t *baseGamePhase) GetFinishChannel() chan struct{} {
	return t.finish
}

func RunGamePhase(phase gamePhase) {
	phase.GetGameSession().sendAllPhaseChange(phase.GetType())
	for _, player := range phase.GetGameSession().GetAlivePlayers() {
		phase.GetGameSession().sendNewRedisChannel(player, phase.GetRedisChannel(player))
	}

	phase.OnStart()

	sendLeftTimeMsg := func(timeLeft time.Duration) {
		phase.GetGameSession().sendAllMsgByServer(
			fmt.Sprintf("The %s will end in %d seconds.", phase.GetName(), int(math.Round(timeLeft.Seconds()))),
		)
	}

	phaseEnd := time.Now().Add(phase.GetPhaseDuration())

	ticker := time.NewTicker(time.Second * 20)

	sendLeftTimeMsg(phase.GetPhaseDuration())

	for {
		select {
		case currentTime := <-ticker.C:
			timeLeft := phaseEnd.Sub(currentTime)

			if timeLeft < time.Millisecond {
				ticker.Stop()
				phase.OnFinish()
				return
			}

			sendLeftTimeMsg(timeLeft)

		case <-phase.GetFinishChannel():
			ticker.Stop()
			phase.OnFinish()
			return
		}
	}
}

type gamePhaseDay struct {
	baseGamePhase
	votesAgainstCount map[string]int
	playerVote        map[string]string
	endedTurn         map[string]struct{}
}

func (t *gamePhaseDay) OnStart() {
	t.sendAllMsgByServer("The city is waking up. Discuss and vote!")
}

func (t *gamePhaseDay) playerEndedTurn(player string) bool {
	_, ok := t.endedTurn[player]
	return ok
}

func (t *gamePhaseDay) VoteAgainst(player string, target string) {
	if t.playerEndedTurn(player) {
		t.sendMsgFromServer(player, "You can't vote after you ended your turn.")
		return
	}
	if currentTarget, ok := t.playerVote[player]; ok {
		t.votesAgainstCount[currentTarget]--
	}
	t.playerVote[player] = target
	t.votesAgainstCount[target]++
	t.sendMsgFromServer(player, fmt.Sprintf("Your vote is against %s now.", target))
}

func (t *gamePhaseDay) Shoot(player string, target string) {
	t.sendMsgFromServer(player, "You can't shoot during the day")
}

func (t *gamePhaseDay) EndTurn(player string) {
	if _, ok := t.endedTurn[player]; ok {
		t.sendMsgFromServer(player, "You already finished your turn.")
		return
	}
	t.endedTurn[player] = struct{}{}
	if len(t.endedTurn) == t.GetGameSession().GetAlivePlayerCount() {
		t.GetFinishChannel() <- struct{}{}
	}
	t.sendAllMsgByServer(fmt.Sprintf("%s finished their turn.", player))
}

func (t *gamePhaseDay) OnFinish() {
	t.sendAllMsgByServer("The day is over.")

	voteResults := make([]string, t.GetGameSession().GetAlivePlayerCount())
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

	addLineToResult("\n")
	addLineToResult("Poll result:")
	addLineToResult("player\tvotes")

	for player, votes := range t.votesAgainstCount {
		addLineToResult(fmt.Sprintf("%s:\t%d", player, votes))
	}

	addLineToResult("\n")

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

	if maxVoteCount == 0 {
		addLineToResult("No one voted today so no one will die today.")
	} else {
		t.sendAllMsgByServer(strings.Join(resultTableLines, "\n"))
		if multipleWinners {
			t.sendAllMsgByServer("There are multiple players with max number of votes so no one will die today.")
		} else {
			t.sendAllMsgByServer(fmt.Sprintf("%[1]s got max number of votes. It's time to go, %[1]s.", mostVoted))
			t.GetGameSession().kill(mostVoted)
		}
	}

	t.GetGameSession().enqueuePhase(MakeGamePhaseNight(t.GetGameSession()))
}

func (t *gamePhaseDay) Check(player string, target string) {
	t.sendMsgFromServer(player, "You can't check during the day.")
}

func (t *gamePhaseDay) PublishCheckResult(player string) {
	if t.GetGameSession().GetPlayerRole(player) != mafia.RoleCommisar {
		t.sendMsgFromServer(player, "Only commissar can publish check results.")
		return
	}
	msg := fmt.Sprintf("For now I know that following users are mafiosi: %s", strings.Join(t.GetGameSession().uncoveredMafia, ", "))
	t.sendMsgAll("commissar", msg)
}

func MakeGamePhaseDay(session *gameSession) *gamePhaseDay {
	commonChat := generateRandomString(10)
	redisChannelByRole := map[mafia.Role]string{
		mafia.RoleMafia:    commonChat,
		mafia.RoleCommisar: commonChat,
		mafia.RoleCitizen:  commonChat,
	}
	return &gamePhaseDay{
		baseGamePhase: baseGamePhase{
			redisChannelByRole: redisChannelByRole,
			name:               "day",
			phaseType:          mafia.GamePhaseTypeDay,
			duration:           mafia.PhaseDuration[mafia.GamePhaseTypeDay],
			gameSession:        session,
			finish:             make(chan struct{}),
			groupEventSender:   session.groupEventSender,
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

func (t *gamePhaseNight) OnStart() {
	t.sendAllMsgByServer("The night is falling. It's time for mafiosi and a commissar to make their moves.")
}

func (t *gamePhaseNight) VoteAgainst(player string, target string) {
	t.sendMsgFromServer(player, "You can't vote during the night.")
}

func (t *gamePhaseNight) EndTurn(player string) {
	t.sendMsgFromServer(player, "You can't end turn during the night.")
}

func (t *gamePhaseNight) Check(player string, target string) {
	if t.GetGameSession().GetPlayerRole(player) != mafia.RoleCommisar {
		t.sendMsgFromServer(player, "Only commissar can do a check.")
		return
	}

	if t.checkUsed {
		t.sendMsgFromServer(player, "You already did a check tonight.")
		return
	}

	t.checkUsed = true

	if t.GetGameSession().GetPlayerRole(target) == mafia.RoleMafia {
		t.GetGameSession().AddUncoveredMafia(target)
		t.sendMsgFromServer(player, fmt.Sprintf("%s is a mafioso. Good job!", target))
	} else {
		t.sendMsgFromServer(player, fmt.Sprintf("%s is not a mafioso. Keep trying!", target))
	}
	if t.nightRolesDone() {
		t.GetFinishChannel() <- struct{}{}
	}
}

func (t *gamePhaseNight) nightRolesDone() bool {
	return t.checkUsed && t.shotByMafia != ""
}

func (t *gamePhaseNight) Shoot(player string, target string) {
	if t.GetGameSession().GetPlayerRole(player) != mafia.RoleMafia {
		t.sendMsgFromServer(player, "Only mafiosi can shoot.")
		return
	}
	if t.shotByMafia != "" {
		t.sendMsgFromServer(player, "You already shot someone tonight.")
		return
	}
	if t.GetGameSession().IsDead(target) {
		t.sendMsgFromServer(player, "You can't shot dead people.")
		return
	}
	t.shotByMafia = target
	t.sendMsgFromServer(player, fmt.Sprintf("You brutally killed %s.", target))
	if t.nightRolesDone() {
		t.GetFinishChannel() <- struct{}{}
	}
}

func (t *gamePhaseNight) PublishCheckResult(player string) {
	t.sendMsgFromServer(player, "You can't publish check result during the night.")
}

func (t *gamePhaseNight) OnFinish() {
	t.sendAllMsgByServer("The night is over.")
	if t.shotByMafia == "" {
		t.sendAllMsgByServer("Mafia didn't shoot anyone today.")
	} else {
		t.sendAllMsgByServer(fmt.Sprintf("Today %s was shot by mafiosi.", t.shotByMafia))
		t.GetGameSession().kill(t.shotByMafia)
	}
	t.GetGameSession().enqueuePhase(MakeGamePhaseDay(t.GetGameSession()))
}

func MakeGamePhaseNight(session *gameSession) *gamePhaseNight {
	redisChannelByRole := map[mafia.Role]string{
		mafia.RoleMafia:    generateRandomString(10),
		mafia.RoleCommisar: generateRandomString(10),
		mafia.RoleCitizen:  "",
	}
	return &gamePhaseNight{
		baseGamePhase: baseGamePhase{
			redisChannelByRole: redisChannelByRole,
			name:               "night",
			phaseType:          mafia.GamePhaseTypeNight,
			duration:           mafia.PhaseDuration[mafia.GamePhaseTypeNight],
			gameSession:        session,
			finish:             make(chan struct{}),
			groupEventSender:   session.groupEventSender,
		},
		checkUsed:   false,
		shotByMafia: "",
	}
}
