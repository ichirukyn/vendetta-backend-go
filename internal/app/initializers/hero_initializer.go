package initializers

import (
	"github.com/gin-gonic/gin"
	"vendetta/internal/adapters/http/rest"
	"vendetta/internal/app/config"
	"vendetta/internal/domain/store"
	"vendetta/internal/usecases"
	"vendetta/pkg/utils"
)

type HeroInitializer struct {
	c *config.Config
	l *utils.Logger

	r *gin.Engine

	store store.Store

	HeroService     *usecases.HeroService
	HeroSpecService *usecases.HeroSpecService
	HeroHandler     *rest.HeroHandler
}

func (i *HeroInitializer) Init() {
	i.HeroService = usecases.NewHeroService(i.c, i.l, i.store.Hero())
	i.HeroSpecService = usecases.NewHeroSpecService(i.c, i.l, i.store.HeroSpec())
	i.HeroHandler = rest.NewHeroHandler(i.HeroService, i.HeroSpecService)

	ig := i.r.Group("/heroes")
	{
		ig.POST("/", i.HeroHandler.Create)
		ig.GET("/", i.HeroHandler.GetAll)
		ig.GET("/:id", i.HeroHandler.GetByID)
		ig.GET("/user/:id", i.HeroHandler.GetByUserID)
		ig.PUT("/:id", i.HeroHandler.Update)
		ig.DELETE("/:id", i.HeroHandler.Delete)
	}
}

func NewHeroInitializer(c *config.Config, logger *utils.Logger, r *gin.Engine, store store.Store) *HeroInitializer {
	return &HeroInitializer{
		c: c,
		l: logger,

		r: r,

		store: store,
	}
}
