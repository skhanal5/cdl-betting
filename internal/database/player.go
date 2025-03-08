package database

type Player struct {
	ID uint
	Name string
	SeasonStat PlayerSeasonStat
	StageStats []PlayerStageStat
}

type PlayerSeasonStat struct {
	ID uint
	SeasonRank int
	KDRatio int
}

type PlayerStageStat struct {
	ID uint
	Name string
	KDRatio int
	MapsPlayed []PlayerMapStat
}

type PlayerMapStat struct {
	ID uint
	Name string
	Mode string
	KDRatio int
	AverageKills int
	KillsPerMinute int
	DamagePerMinute int
}