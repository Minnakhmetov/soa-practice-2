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
	RoleMafia:    3,
	RoleCommisar: 1,
	RoleCitizen:  6,
}

var PlayersInGame = 10

var PhaseDurationInSecs = map[GamePhaseType]int{
	GamePhaseTypeDay:   15,
	GamePhaseTypeNight: 10,
}
