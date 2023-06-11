package mafia

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
	RoleCitizen:  2,
}

var PlayersInGame = 4

var PhaseDurationInSecs = map[GamePhaseType]int{
	GamePhaseTypeDay:   15,
	GamePhaseTypeNight: 10,
}
