package entities

type Enemy struct {
	ID        string `json:"id"  gorm:"primary_key;autoIncrement"`
	SpeciesID string `json:"species_id"`
	TypesID   string `json:"types_id"`
	RankID    string `json:"rank_id"`
	Name      string `json:"name"`
	Level     int    `json:"level"`
	Affixes   Affix  `json:"affixes"`
}

type Affix struct {
	ID          string `json:"id"  gorm:"primary_key;autoIncrement"`
	Name        string `json:"name"`
	Desc        string `json:"desc"`
	TagID       string `json:"tag_id"`
	Probability int    `json:"probability"`
	Effects     Effect `json:"effects"`
}

type Effect struct {
}

type EnemySpec struct {
	ID             string  `json:"id"  gorm:"primary_key;autoIncrement"`
	EnemyID        string  `json:"hero_id"`
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

	Enemy Enemy `json:"hero,omitempty"`
}

type EnemySkill struct {
	ID      string  `json:"id"  gorm:"primary_key;autoIncrement"`
	EnemyID string  `json:"hero_id"`
	SpellID string  `json:"spell_id"`
	Level   int     `json:"level"`
	Exp     float64 `json:"exp"`

	Enemy Enemy `json:"hero,omitempty"`
}

// PAUSE
type EnemyWeapon struct {
	ID       string `json:"id"`
	EnemyID  string `json:"hero_id"`
	WeaponID string `json:"weapon_id"`
	Level    int    `json:"level"`

	Enemy Enemy `json:"hero,omitempty"`
}

// PAUSE
type EnemyStorage struct {
	ID        string `json:"id"`
	EnemyID   string `json:"hero_id"`
	StorageID string `json:"storage_id"`

	Enemy   Enemy   `json:"hero,omitempty"`
	Storage Storage `json:"storage,omitempty"`
}
