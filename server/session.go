package main

import (
	"fmt"
	"math/rand"

	"github.com/Minnakhmetov/soa-practice-2/mafia"
)

type playerInfo struct {
	role   mafia.Role
	isDead bool
}

type gameSession struct {
	playerInfoByUsername map[string]*playerInfo
	isFinished           bool
	finish               chan struct{}
	nextPhase            chan gamePhase
	currentPhase         gamePhase
	uncoveredMafia       []string
	groupEventSender
}

func (t *gameSession) enqueuePhase(phase gamePhase) {
	t.nextPhase <- phase
}

func (t *gameSession) kill(username string) {
	t.playerInfoByUsername[username].isDead = true
	t.sendAllMsgByServer(fmt.Sprintf("%s died.", username))
}

func (t *gameSession) getAliveMafiaCount() int {
	aliveMafia := 0
	for _, playerInfo := range t.playerInfoByUsername {
		if playerInfo.role == mafia.RoleMafia && !playerInfo.isDead {
			aliveMafia++
		}
	}
	return aliveMafia
}

func (t *gameSession) getAliveNonMafiaCount() int {
	aliveNonMafia := 0
	for _, playerInfo := range t.playerInfoByUsername {
		if playerInfo.role != mafia.RoleMafia && !playerInfo.isDead {
			aliveNonMafia++
		}
	}
	return aliveNonMafia
}

func (t *gameSession) mafiaWon() bool {
	return t.getAliveMafiaCount() >= t.getAliveNonMafiaCount()
}

func (t *gameSession) citizensWon() bool {
	return t.getAliveMafiaCount() == 0
}

func (t *gameSession) GetAlivePlayerCount() int {
	alive := 0
	for player := range t.playerInfoByUsername {
		if !t.IsDead(player) {
			alive++
		}
	}
	return alive
}

func (t *gameSession) playerExists(username string) bool {
	_, ok := t.playerInfoByUsername[username]
	return ok
}

func (t *gameSession) GetPlayerRole(username string) mafia.Role {
	return t.playerInfoByUsername[username].role
}

func (t *gameSession) IsDead(username string) bool {
	return t.playerInfoByUsername[username].isDead
}

func (t *gameSession) AddUncoveredMafia(mafia string) {
	for _, value := range t.uncoveredMafia {
		if value == mafia {
			return
		}
	}
	t.uncoveredMafia = append(t.uncoveredMafia, mafia)
}

func MakeGameSession(usernames []string, sender eventSender) *gameSession {
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
		playerInfoByUsername: map[string]*playerInfo{},
		isFinished:           false,
		finish:               make(chan struct{}),
		groupEventSender: groupEventSender{
			members:     usernames,
			eventSender: sender,
		},
		nextPhase: make(chan gamePhase, 1),
	}

	for i := 0; i < len(roles); i++ {
		session.playerInfoByUsername[usernames[i]] = &playerInfo{
			role:   roles[i],
			isDead: false,
		}
	}

	return &session
}

func (t *gameSession) Run() {
	for player, playerInfo := range t.playerInfoByUsername {
		t.sendMsgFromServer(player, fmt.Sprintf("The game begins. Your role: %s", playerInfo.role))
		t.sendRoleAssignment(player, playerInfo.role)	
	}

	go func() {
		for phase := range t.nextPhase {
			t.currentPhase = phase
			RunGamePhase(phase)
			if t.mafiaWon() {
				t.sendAllMsgByServer("Game over. Mafia won.")
				return
			} else if t.citizensWon() {
				t.sendAllMsgByServer("Game over. Citizen won.")
				return
			}
		}
	}()

	t.enqueuePhase(MakeGamePhaseDay(t))
}

func (t *gameSession) doPhaseAction(player string, action func(string)) {
	if t.IsDead(player) {
		t.sendMsgFromServer(player, "You can't do it when you're dead.")
		return
	}
	action(player)
}

func (t *gameSession) doPhaseActionWithTarget(player string, target string, action func(string, string)) {
	if t.IsDead(player) {
		t.sendMsgFromServer(player, "You can't do it when you're dead.")
		return
	}
	if !t.playerExists(target) {
		t.sendMsgFromServer(player, fmt.Sprintf("User %s not found.", target))
		return
	}
	action(player, target)
}

func (t *gameSession) VoteAgainst(player string, target string) {
	t.doPhaseActionWithTarget(player, target, t.currentPhase.VoteAgainst)
}

func (t *gameSession) Shoot(player string, target string) {
	t.doPhaseActionWithTarget(player, target, t.currentPhase.Shoot)
}

func (t *gameSession) Check(player string, target string) {
	t.doPhaseActionWithTarget(player, target, t.currentPhase.Check)
}

func (t *gameSession) PublishCheckResult(player string) {
	t.doPhaseAction(player, t.currentPhase.PublishCheckResult)
}

func (t *gameSession) EndTurn(player string) {
	t.doPhaseAction(player, t.currentPhase.EndTurn)
}

func (t *gameSession) GetAlivePlayers() []string {
	var alivePlayers []string
	for player, playerInfo := range t.playerInfoByUsername {
		if !playerInfo.isDead {
			alivePlayers = append(alivePlayers, player)
		}
	}
	return alivePlayers
}

func (t *gameSession) GetRedisChannel(player string) string {
	return t.currentPhase.GetRedisChannel(player)
}
