package initializers

import (
	"github.com/gin-gonic/gin"
	"vendetta/internal/adapters/http/rest"
	"vendetta/internal/app/config"
	"vendetta/internal/domain/store"
	"vendetta/internal/usecases"
	"vendetta/pkg/utils"
)

type ClassEffectInitializer struct {
	c *config.Config
	l *utils.Logger

	r *gin.Engine

	store store.Store

	ClassEffectService *usecases.ClassEffectService
	ClassEffectHandler *rest.ClassEffectHandler
}

func (i *ClassEffectInitializer) Init() {
	i.ClassEffectService = usecases.NewClassEffectService(i.c, i.l, i.store.ClassEffect())
	i.ClassEffectHandler = rest.NewClassEffectHandler(i.ClassEffectService)

	ig := i.r.Group("/classes")
	{
		ig.GET("/effects", i.ClassEffectHandler.GetAll)
		ig.GET("/:id/effects", i.ClassEffectHandler.GetByClassID)
	}
}

func NewClassEffectInitializer(c *config.Config, logger *utils.Logger, r *gin.Engine, store store.Store) *ClassEffectInitializer {
	return &ClassEffectInitializer{
		c: c,
		l: logger,

		r: r,

		store: store,
	}
}
