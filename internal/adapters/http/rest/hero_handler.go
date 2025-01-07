package rest

import (
	"github.com/gin-gonic/gin"
	"vendetta/internal/adapters/factory"
	"vendetta/internal/adapters/http/rest/dto"
	"vendetta/internal/domain"
	_default "vendetta/internal/domain/default"
	"vendetta/internal/usecases"
	"vendetta/pkg/utils"
)

// HeroHandler хэндлер для работы с пользователями
type HeroHandler struct {
	HeroService     *usecases.HeroService
	HeroSpecService *usecases.HeroSpecService
}

func (h *HeroHandler) Create(ctx *gin.Context) {
	var body dto.CreateHeroDTO
	if err := ctx.ShouldBindJSON(&body); err != nil {
		utils.NewBadRequestError(ctx, domain.NewBadRequestError("bad request data"))
		return
	}

	heroFactory := factory.NewHeroFactory()
	hero := heroFactory.CreateHeroFromCreateDTO(body)
	if err := h.HeroService.Create(hero); err != nil {
		utils.NewBadRequestError(ctx, err)
		return
	}

	specBody := _default.NewCreateHeroSpecDTO(hero.ID)

	specFactory := factory.NewHeroSpecFactory()
	spec := specFactory.CreateHeroSpecFromCreateDTO(specBody)

	if err := h.HeroSpecService.Create(spec); err != nil {
		utils.NewBadRequestError(ctx, err)
		return
	}

	utils.ResponseSuccessHandler(ctx, hero)
}

func (h *HeroHandler) GetAll(ctx *gin.Context) {
	filter := utils.GetDefaultsFilterFromQuery(ctx)

	var payload dto.FindHeroesDTO
	if err := ctx.ShouldBindQuery(&payload); err != nil {
		utils.NewBadRequestError(ctx, domain.NewBadRequestError("bad request data"))
		return
	}

	query, err := utils.StructToMap(&payload)
	if err != nil {
		utils.NewBadRequestError(ctx, domain.NewBadRequestError("bad request data"))
		return
	}

	heroes, dErr := h.HeroService.GetAll(filter, &query)
	if dErr != nil {
		utils.NewBadRequestError(ctx, dErr)
		return
	}

	utils.ResponseSuccessHandler(ctx, heroes)
}

func (h *HeroHandler) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")

	hero, dErr := h.HeroService.GetByID(id)
	if dErr != nil {
		utils.NewBadRequestError(ctx, dErr)
		return
	}

	utils.ResponseSuccessHandler(ctx, hero)
}

func (h *HeroHandler) GetByUserID(ctx *gin.Context) {
	id := ctx.Param("id")

	races, dErr := h.HeroService.GetByUserID(id)
	if dErr != nil {
		utils.NewBadRequestError(ctx, dErr)
		return
	}

	utils.ResponseSuccessHandler(ctx, races)
}

func (h *HeroHandler) Update(ctx *gin.Context) {
	id := ctx.Param("id")

	var body dto.UpdateHeroDTO
	if err := ctx.ShouldBindJSON(&body); err != nil {
		utils.NewBadRequestError(ctx, domain.NewBadRequestError("bad request data"))
		return
	}

	heroFactory := factory.NewHeroFactory()
	hero := heroFactory.CreateHeroFromUpdateDTO(id, body)
	if err := h.HeroService.Update(hero); err != nil {
		utils.NewBadRequestError(ctx, err)
		return
	}

	utils.ResponseSuccessHandler(ctx, hero)
}

func (h *HeroHandler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := h.HeroService.Delete(id); err != nil {
		utils.NewBadRequestError(ctx, err)
		return
	}

	utils.ResponseSuccessHandler(ctx, true)
}

func NewHeroHandler(heroService *usecases.HeroService, heroSpecService *usecases.HeroSpecService) *HeroHandler {
	return &HeroHandler{
		HeroService:     heroService,
		HeroSpecService: heroSpecService,
	}
}
