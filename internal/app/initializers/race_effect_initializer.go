package initializers

import (
	"github.com/gin-gonic/gin"
	"vendetta/internal/adapters/http/rest"
	"vendetta/internal/app/config"
	"vendetta/internal/domain/store"
	"vendetta/internal/usecases"
	"vendetta/pkg/utils"
)

type RaceEffectInitializer struct {
	c *config.Config
	l *utils.Logger

	r *gin.Engine

	store store.Store

	RaceEffectService *usecases.RaceEffectService
	RaceEffectHandler *rest.RaceEffectHandler
}

func (i *RaceEffectInitializer) Init() {
	i.RaceEffectService = usecases.NewRaceEffectService(i.c, i.l, i.store.RaceEffect())
	i.RaceEffectHandler = rest.NewRaceEffectHandler(i.RaceEffectService)

	ig := i.r.Group("/races/effects")
	{
		ig.GET("/", i.RaceEffectHandler.GetAll)
		ig.GET("/:id", i.RaceEffectHandler.GetByRaceID)
	}
}

func NewRaceEffectInitializer(c *config.Config, logger *utils.Logger, r *gin.Engine, store store.Store) *RaceEffectInitializer {
	return &RaceEffectInitializer{
		c: c,
		l: logger,

		r: r,

		store: store,
	}
}
