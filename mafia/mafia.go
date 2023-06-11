package mafia

import "time"

type GamePhaseType string

const (
	GamePhaseTypeDay   GamePhaseType = "day"
	GamePhaseTypeNight GamePhaseType = "night"
)

type Role string

const (
	RoleMafia    Role = "mafia"
	RoleCitizen  Role = "citizen"
	RoleCommisar Role = "commissar"
)

var RoleToCount = map[Role]int{
	RoleMafia:    1,
	RoleCommisar: 1,
	RoleCitizen:  1,
}

var PlayersInGame = 3

var PhaseDuration = map[GamePhaseType]time.Duration{
	GamePhaseTypeDay:   time.Second * 60,
	GamePhaseTypeNight: time.Second * 60,
}
