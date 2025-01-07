package initializers

import (
	"github.com/gin-gonic/gin"
	"vendetta/internal/adapters/http/rest"
	"vendetta/internal/app/config"
	"vendetta/internal/domain/store"
	"vendetta/internal/usecases"
	"vendetta/pkg/utils"
)

type SkillInitializer struct {
	c *config.Config
	l *utils.Logger

	r *gin.Engine

	store store.Store

	SkillService *usecases.SkillService
	SkillHandler *rest.SkillHandler
}

func (i *SkillInitializer) Init() {
	i.SkillService = usecases.NewSkillService(i.c, i.l, i.store.Skill())
	i.SkillHandler = rest.NewSkillHandler(i.SkillService)

	ig := i.r.Group("/skills")
	{
		ig.POST("/", i.SkillHandler.Create)
		ig.GET("/", i.SkillHandler.GetAll)
		ig.GET("/:id", i.SkillHandler.GetByID)
		ig.PUT("/:id", i.SkillHandler.Update)
		ig.DELETE("/:id", i.SkillHandler.Delete)
	}
}

func NewSkillInitializer(c *config.Config, logger *utils.Logger, r *gin.Engine, store store.Store) *SkillInitializer {
	return &SkillInitializer{
		c: c,
		l: logger,

		r: r,

		store: store,
	}
}
