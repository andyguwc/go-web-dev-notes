/*

Flyweight

Create a Betting Webpage

The acceptance criteria for a Flyweight pattern must always reduce the amount of memory
that is used, and must be focused primarily on this objective:
1. We will create a Team struct with some basic information such as the team's
name, players, historical results, and an image depicting their shield.
2. We must ensure correct team creation (note the word creation here, candidate for
a creational pattern), and not having duplicates.
3. When creating the same team twice, we must have two pointers pointing to the
same memory address.


*/



package flyweight

import "time"

const (
	TEAM_A = iota
	TEAM_B
)

type Player struct {
	Name         string
	Surname      string
	PreviousTeam uint64
	Photo        []byte
}

type HistoricalData struct {
	Year          uint8
	LeagueResults []Match
}

type Team struct {
	ID             uint64
	Name           string
	Shield         []byte
	Players        []Player
	HistoricalData []HistoricalData
}

type Match struct {
	Date          time.Time
	VisitorID     uint64
	LocalID       uint64
	LocalScore    byte
	VisitorScore  byte
	LocalShoots   uint16
	VisitorShoots uint16
}

func getTeamFactory(team int) Team {
	switch team {
	case TEAM_B:
		return Team{
			ID:   2,
			Name: "TEAM_B",
		}
	default:
		return Team{
			ID:   1,
			Name: "TEAM_A",
		}
	}
}

func NewTeamFactory() teamFlyweightFactory {
	return teamFlyweightFactory{
		createdTeams: make(map[int]*Team, 0),
	}
}

type teamFlyweightFactory struct {
	createdTeams map[int]*Team
}

func (t *teamFlyweightFactory) GetTeam(teamName int) *Team {
	if t.createdTeams[teamName] != nil {
		return t.createdTeams[teamName]
	}

	team := getTeamFactory(teamName)
	t.createdTeams[teamName] = &team

	return t.createdTeams[teamName]
}

func (t *teamFlyweightFactory) GetNumberOfObjects() int {
	return len(t.createdTeams)
}
