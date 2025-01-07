package entities

type Hero struct {
	ID       string  `json:"id" gorm:"primary_key;autoIncrement"`
	UserID   string  `json:"user_id"`
	RaceID   string  `json:"race_id"`
	ClassID  string  `json:"class_id"`
	Name     string  `json:"name"`
	Rank     int     `json:"rank"`
	IsHidden bool    `json:"is_hidden"`
	Level    int     `json:"level"`
	Exp      float64 `json:"exp"`
	//ClanID   string  `json:"clan_id"`

	User User `json:"user,omitempty"`
}

type HeroSpec struct {
	ID             string  `json:"id"  gorm:"primary_key;autoIncrement"`
	HeroID         string  `json:"hero_id"`
	Accuracy       int     `json:"accuracy"`
	Strength       int     `json:"strength"`
	Health         int     `json:"health"`
	Speed          int     `json:"speed"`
	Dexterity      int     `json:"dexterity"`
	Soul           int     `json:"soul"`
	Intelligence   int     `json:"intelligence"`
	Submissions    int     `json:"submissions"`
	CriticalRate   float64 `json:"critical_rate"`
	CriticalDamage float64 `json:"critical_damage"`
	Resistance     float64 `json:"resistance"`
	TotalSpec      int     `json:"total_spec"`
	FreeSpec       int     `json:"free_spec"`
}

type HeroSkill struct {
	ID      string  `json:"id"  gorm:"primary_key;autoIncrement"`
	HeroID  string  `json:"hero_id"`
	SpellID string  `json:"spell_id"`
	Level   int     `json:"level"`
	Exp     float64 `json:"exp"`

	Hero Hero `json:"hero,omitempty"`
}

// PAUSE
type HeroStatistics struct {
	ID                  string `json:"id"`
	HeroID              string `json:"hero_id"`
	Damage              int    `json:"damage"`
	DamageMax           int    `json:"damage_max"`
	DamageTaken         int    `json:"damage_taken"`
	DamageTakenMax      int    `json:"damage_taken_max"`
	BlockDamage         int    `json:"block_damage"`
	BlockCount          int    `json:"block_count"`
	Healing             int    `json:"healing"`
	HealingMax          int    `json:"healing_max"`
	HitsCount           int    `json:"hits_count"`
	MissCount           int    `json:"miss_count"`
	MoneyAll            int    `json:"money_all"`
	MoneyWasted         int    `json:"money_wasted"`
	EvasionCount        int    `json:"evasion_count"`
	EvasionSuccessCount int    `json:"evasion_success_count"`
	CounterStrikeCount  int    `json:"counter_strike_count"`
	CounterStrikeDamage int    `json:"counter_strike_damage"`
	PassCount           int    `json:"pass_count"`
	WinOneToOne         int    `json:"win_one_to_one"`
	WinTeamToTeam       int    `json:"win_team_to_team"`
	LoseOneToOne        int    `json:"lose_one_to_one"`
	LoseTeamToTeam      int    `json:"lose_team_to_team"`
	KillEnemy           int    `json:"kill_enemy"`
	KillHero            int    `json:"kill_hero"`
	Death               int    `json:"death"`
	EscapeCount         int    `json:"escape_count"`
	CriticalCount       int    `json:"critical_count"`
	CountOneToOne       int    `json:"count_one_to_one"`
	CountTeamToTeam     int    `json:"count_team_to_team"`

	Hero Hero `json:"hero,omitempty"`
}

// PAUSE
type HeroTeam struct {
	ID       string `json:"id"`
	HeroID   string `json:"hero_id"`
	TeamID   string `json:"team_id"`
	IsLeader bool   `json:"is_leader"`
	Prefix   string `json:"prefix"`

	Hero Hero `json:"hero,omitempty"`
}

// PAUSE
type HeroWeapon struct {
	ID       string `json:"id"`
	HeroID   string `json:"hero_id"`
	WeaponID string `json:"weapon_id"`
	Level    int    `json:"level"`

	Hero Hero `json:"hero,omitempty"`
}

// PAUSE
type HeroStorage struct {
	ID        string `json:"id"`
	HeroID    string `json:"hero_id"`
	StorageID string `json:"storage_id"`

	Hero    Hero    `json:"hero,omitempty"`
	Storage Storage `json:"storage,omitempty"`
}
