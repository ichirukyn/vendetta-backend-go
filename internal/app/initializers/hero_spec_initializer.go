package initializers

import (
	"github.com/gin-gonic/gin"
	"vendetta/internal/adapters/http/rest"
	"vendetta/internal/app/config"
	"vendetta/internal/domain/store"
	"vendetta/internal/usecases"
	"vendetta/pkg/utils"
)

type HeroSpecInitializer struct {
	c *config.Config
	l *utils.Logger

	r *gin.Engine

	store store.Store

	HeroSpecService *usecases.HeroSpecService
	HeroSpecHandler *rest.HeroSpecHandler
}

func (i *HeroSpecInitializer) Init() {
	i.HeroSpecService = usecases.NewHeroSpecService(i.c, i.l, i.store.HeroSpec())
	i.HeroSpecHandler = rest.NewHeroSpecHandler(i.HeroSpecService)

	ig := i.r.Group("/heroes/specs")
	{
		ig.POST("/", i.HeroSpecHandler.Create)
		ig.GET("/:id", i.HeroSpecHandler.GetByHeroID)
		ig.PUT("/:id", i.HeroSpecHandler.Update)
		ig.DELETE("/:id", i.HeroSpecHandler.Delete)
	}
}

func NewHeroSpecInitializer(c *config.Config, logger *utils.Logger, r *gin.Engine, store store.Store) *HeroSpecInitializer {
	return &HeroSpecInitializer{
		c: c,
		l: logger,

		r: r,

		store: store,
	}
}
