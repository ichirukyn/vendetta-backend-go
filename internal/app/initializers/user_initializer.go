package initializers

import (
	"github.com/gin-gonic/gin"
	"vendetta/internal/adapters/http/rest"
	"vendetta/internal/app/config"
	"vendetta/internal/domain/store"
	"vendetta/internal/usecases"
	"vendetta/pkg/utils"
)

type UserInitializer struct {
	c *config.Config
	l *utils.Logger

	r *gin.Engine

	store store.Store

	UserService *usecases.UserService
	UserHandler *rest.UserHandler
}

func (i *UserInitializer) Init() {
	i.UserService = usecases.NewUserService(i.c, i.l, i.store.User())
	i.UserHandler = rest.NewUserHandler(i.UserService)

	ig := i.r.Group("/users")
	{
		ig.POST("/", i.UserHandler.Create)
		ig.GET("/", i.UserHandler.GetAll)
		ig.GET("/:id", i.UserHandler.GetByID)
		ig.PUT("/:id", i.UserHandler.Update)
		ig.DELETE("/:id", i.UserHandler.Delete)
	}
}

func NewUserInitializer(c *config.Config, logger *utils.Logger, r *gin.Engine, store store.Store) *UserInitializer {
	return &UserInitializer{
		c: c,
		l: logger,

		r: r,

		store: store,
	}
}
