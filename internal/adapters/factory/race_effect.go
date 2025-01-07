package factory

// RaceFactory фабрика для создания сущностей пользователя (entities.RaceEffect)
type RaceEffectFactory struct {
}

// NewRaceEffectFactory инициализатор RaceEffectFactory
func NewRaceEffectFactory() *RaceEffectFactory {
	return &RaceEffectFactory{}
}
