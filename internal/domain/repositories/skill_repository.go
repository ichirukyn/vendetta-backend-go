package repositories

import "vendetta/internal/domain/entities"

type SkillRepository interface {
	Create(*entities.Skill) error
	GetAll(*entities.Filter, *entities.WhereQueryFilter) ([]*entities.Skill, error)
	GetByID(string) (*entities.Skill, error)
	Update(*entities.Skill) error
	Delete(*entities.Skill) error
}

//type SkillEffectRepository interface {
//	GetAll(*entities.Filter, *entities.WhereQueryFilter) ([]*entities.SkillEffect, error)
//	GetBySkillID(string) ([]*entities.SkillEffect, error)
//}
