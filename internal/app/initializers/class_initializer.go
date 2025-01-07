package initializers

import (
	"github.com/gin-gonic/gin"
	"vendetta/internal/adapters/http/rest"
	"vendetta/internal/app/config"
	"vendetta/internal/domain/store"
	"vendetta/internal/usecases"
	"vendetta/pkg/utils"
)

type ClassInitializer struct {
	c *config.Config
	l *utils.Logger

	r *gin.Engine

	store store.Store

	ClassService *usecases.ClassService
	ClassHandler *rest.ClassHandler
}

func (i *ClassInitializer) Init() {
	i.ClassService = usecases.NewClassService(i.c, i.l, i.store.Class())
	i.ClassHandler = rest.NewClassHandler(i.ClassService)

	ig := i.r.Group("/classes")
	{
		ig.POST("/", i.ClassHandler.Create)
		ig.GET("/", i.ClassHandler.GetAll)
		ig.GET("/:id", i.ClassHandler.GetByID)
		ig.PUT("/:id", i.ClassHandler.Update)
		ig.DELETE("/:id", i.ClassHandler.Delete)
	}
}

func NewClassInitializer(c *config.Config, logger *utils.Logger, r *gin.Engine, store store.Store) *ClassInitializer {
	return &ClassInitializer{
		c: c,
		l: logger,

		r: r,

		store: store,
	}
}
