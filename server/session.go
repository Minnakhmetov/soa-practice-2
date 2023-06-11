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
	playerInfo         map[string]*playerInfo
	isFinished         bool
	aliveMafiaCount    int
	aliveNonMafiaCount int
	finish             chan struct{}
	nextPhase          chan gamePhase
	currentPhase       gamePhase
	uncoveredMafia     []string
	groupMsgSender
}

// func (t *gameSession) send(sender string, receiver string, msg string) {
// 	t.playerInfo[receiver].msgs <- fmt.Sprintf("[%s] %s\n", sender, msg)
// }

func (t *gameSession) enqueuePhase(phase gamePhase) {
	t.nextPhase <- phase
}

func (t *gameSession) kill(username string) {
	t.playerInfo[username].isDead = true
	t.sendAllServerMessage(fmt.Sprintf("%s died.", username))
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

func (t *gameSession) playerExists(username string) bool {
	_, ok := t.playerInfo[username]
	return ok
}

func (t *gameSession) GetPlayerRole(username string) mafia.Role {
	return t.playerInfo[username].role
}

func (t *gameSession) IsDead(username string) bool {
	return t.playerInfo[username].isDead
}

func (t *gameSession) AddUncoveredMafia(mafia string) {
	for _, value := range t.uncoveredMafia {
		if value == mafia {
			return
		}
	}
	t.uncoveredMafia = append(t.uncoveredMafia, mafia)
}

func MakeGameSession(usernames []string, sender msgSender) *gameSession {
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
		groupMsgSender: groupMsgSender{
			members:   usernames,
			msgSender: sender,
		},
	}

	for i := 0; i < len(roles); i++ {
		session.playerInfo[usernames[i]] = &playerInfo{
			role:   roles[i],
			isDead: false,
		}
	}

	return &session
}

func (t *gameSession) Run() {
	for player, playerInfo := range t.playerInfo {
		t.sendServerMessage(player, fmt.Sprintf("The game begins. Your role: %s", playerInfo.role))
	}
	go func() {
		for phase := range t.nextPhase {
			t.currentPhase = phase
			RunGamePhase(phase)
		}
	}()
	t.enqueuePhase(MakeGamePhaseDay(t))
}

func (t *gameSession) doPhaseAction(player string, action func()) {
	if t.IsDead(player) {
		t.sendServerMessage(player, "You can't do it when you're dead.")
		return
	}
	action()
}

func (t *gameSession) VoteAgainst(player string, target string) {
	t.doPhaseAction(player, func() {
		t.currentPhase.VoteAgainst(player, target)
	})

}

func (t *gameSession) Shoot(player string, target string) {
	t.doPhaseAction(player, func() {
		t.currentPhase.Shoot(player, target)
	})
}

func (t *gameSession) Check(player string, target string) {
	t.doPhaseAction(player, func() {
		t.currentPhase.Check(player, target)
	})
}

func (t *gameSession) PublishCheckResult(player string) {
	t.doPhaseAction(player, func() {
		t.currentPhase.PublishCheckResult(player)
	})
}

func (t *gameSession) EndTurn(player string) {
	t.doPhaseAction(player, func() {
		t.currentPhase.EndTurn(player)
	})
}
