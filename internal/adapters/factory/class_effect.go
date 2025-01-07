package factory

// ClassFactory фабрика для создания сущностей пользователя (entities.ClassEffect)
type ClassEffectFactory struct {
}

// NewClassEffectFactory инициализатор ClassEffectFactory
func NewClassEffectFactory() *ClassEffectFactory {
	return &ClassEffectFactory{}
}
