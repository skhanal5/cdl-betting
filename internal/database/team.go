package database

type Team struct {
	ID uint 
	Name string
	Players []Player 
	Season TeamSeasonStat
	Stages []TeamStageStat
}


type TeamSeasonStat struct {
	ID uint
	Wins int 
	Losses int
}

type TeamStageStat struct {
	ID uint
	Name string
	Wins int
	Losses int
	MapsPlayed []Map
	Matches []Match
}

type Map struct {
	ID uint
	Map string
	Mode string
	Wins int
	Losses int
	Date string
}

type Match struct {
	ID uint
	Map string
	Mode string
	Result string;
	Opponent Team;
	Date string;
}