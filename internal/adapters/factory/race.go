package factory

// RaceFactory фабрика для создания сущностей пользователя (entities.Race)
type RaceFactory struct {
}

// NewRaceFactory инициализатор RaceFactory
func NewRaceFactory() *RaceFactory {
	return &RaceFactory{}
}
