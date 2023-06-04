package mafia

type GamePhase string

const (
	GamePhaseDay   GamePhase = "day"
	GamePhaseNight GamePhase = "night"
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

var PlayersRequired = 10

var PhaseDurationInSecs = map[GamePhase]int{
	GamePhaseDay:   15,
	GamePhaseNight: 10,
}
