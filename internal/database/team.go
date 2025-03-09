package database

type Team struct {
	Name    string `gorm:"primaryKey"`
	Players []Player
	Season  TeamSeasonStat
	Stages  []TeamStageStat
}

type TeamSeasonStat struct {
	ID     uint
	Wins   int
	Losses int
}

type TeamStageStat struct {
	ID         uint
	Name       string
	Wins       int
	Losses     int
	MapsPlayed []Map
	Matches    []Match
}

type Map struct {
	Map    string `gorm:"primaryKey"`
	Mode   string
	Wins   int
	Losses int
	Date   string
}

type Match struct {
	ID       uint
	Map      string
	Mode     string
	Result   string
	Opponent Team
	Date     string
}
