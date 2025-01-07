package rest

import (
	"github.com/gin-gonic/gin"
	"vendetta/internal/adapters/factory"
	"vendetta/internal/adapters/http/rest/dto"
	"vendetta/internal/domain"
	"vendetta/internal/usecases"
	"vendetta/pkg/utils"
)

// HeroSpecHandler хэндлер для работы с пользователями
type HeroSpecHandler struct {
	HeroSpecService *usecases.HeroSpecService
}

func (h *HeroSpecHandler) Create(ctx *gin.Context) {
	var body dto.CreateHeroSpecDTO
	if err := ctx.ShouldBindJSON(&body); err != nil {
		utils.NewBadRequestError(ctx, domain.NewBadRequestError("bad request data"))
		return
	}

	heroSpecFactory := factory.NewHeroSpecFactory()
	heroSpec := heroSpecFactory.CreateHeroSpecFromCreateDTO(body)
	if err := h.HeroSpecService.Create(heroSpec); err != nil {
		utils.NewBadRequestError(ctx, err)
		return
	}

	utils.ResponseSuccessHandler(ctx, heroSpec)
}

func (h *HeroSpecHandler) GetByHeroID(ctx *gin.Context) {
	id := ctx.Param("id")

	heroSpecs, dErr := h.HeroSpecService.GetByHeroID(id)
	if dErr != nil {
		utils.NewBadRequestError(ctx, dErr)
		return
	}

	utils.ResponseSuccessHandler(ctx, heroSpecs)
}

func (h *HeroSpecHandler) Update(ctx *gin.Context) {
	id := ctx.Param("id")

	var body dto.UpdateHeroSpecDTO
	if err := ctx.ShouldBindJSON(&body); err != nil {
		utils.NewBadRequestError(ctx, domain.NewBadRequestError("bad request data"))
		return
	}

	heroSpecFactory := factory.NewHeroSpecFactory()
	heroSpec := heroSpecFactory.CreateHeroSpecFromUpdateDTO(id, body)
	if err := h.HeroSpecService.Update(heroSpec); err != nil {
		utils.NewBadRequestError(ctx, err)
		return
	}

	utils.ResponseSuccessHandler(ctx, heroSpec)
}

func (h *HeroSpecHandler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := h.HeroSpecService.Delete(id); err != nil {
		utils.NewBadRequestError(ctx, err)
		return
	}

	utils.ResponseSuccessHandler(ctx, true)
}

func NewHeroSpecHandler(heroSpecService *usecases.HeroSpecService) *HeroSpecHandler {
	return &HeroSpecHandler{
		HeroSpecService: heroSpecService,
	}
}
