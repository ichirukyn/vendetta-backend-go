package initializers

import (
	"github.com/gin-gonic/gin"
	"vendetta/internal/adapters/http/rest"
	"vendetta/internal/app/config"
	"vendetta/internal/domain/store"
	"vendetta/internal/usecases"
	"vendetta/pkg/utils"
)

type RaceInitializer struct {
	c *config.Config
	l *utils.Logger

	r *gin.Engine

	store store.Store

	RaceService *usecases.RaceService
	RaceHandler *rest.RaceHandler
}

func (i *RaceInitializer) Init() {
	i.RaceService = usecases.NewRaceService(i.c, i.l, i.store.Race())
	i.RaceHandler = rest.NewRaceHandler(i.RaceService)

	ig := i.r.Group("/races")
	{
		ig.GET("/", i.RaceHandler.GetAll)
		ig.GET("/:id", i.RaceHandler.GetByID)
	}
}

func NewRaceInitializer(c *config.Config, logger *utils.Logger, r *gin.Engine, store store.Store) *RaceInitializer {
	return &RaceInitializer{
		c: c,
		l: logger,

		r: r,

		store: store,
	}
}
