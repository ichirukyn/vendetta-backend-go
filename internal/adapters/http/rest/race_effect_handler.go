package rest

import (
	"github.com/gin-gonic/gin"
	"vendetta/internal/adapters/http/rest/dto"
	"vendetta/internal/domain"
	"vendetta/internal/usecases"
	"vendetta/pkg/utils"
)

// RaceEffectHandler хэндлер для работы с пользователями
type RaceEffectHandler struct {
	RaceEffectService *usecases.RaceEffectService
}

// GetAll хэндлер для получения пользователей (entities.RaceEffect)
func (h *RaceEffectHandler) GetAll(ctx *gin.Context) {
	filter := utils.GetDefaultsFilterFromQuery(ctx)

	var payload dto.FindRacesEffectsDTO
	if err := ctx.ShouldBindQuery(&payload); err != nil {
		utils.NewBadRequestError(ctx, domain.NewBadRequestError("bad request data"))
		return
	}

	query, err := utils.StructToMap(&payload)
	if err != nil {
		utils.NewBadRequestError(ctx, domain.NewBadRequestError("bad request data"))
		return
	}

	races, dErr := h.RaceEffectService.GetAll(filter, &query)
	if dErr != nil {
		utils.NewBadRequestError(ctx, dErr)
		return
	}

	utils.ResponseSuccessHandler(ctx, races)
}

// GetByRaceID хэндлер для получения пользователя по id (entities.RaceEffect)
func (h *RaceEffectHandler) GetByRaceID(ctx *gin.Context) {
	id := ctx.Param("id")

	races, dErr := h.RaceEffectService.GetByID(id)
	if dErr != nil {
		utils.NewBadRequestError(ctx, dErr)
		return
	}

	utils.ResponseSuccessHandler(ctx, races)
}

// NewRaceEffectHandler инициализатор UserHandler
func NewRaceEffectHandler(raceService *usecases.RaceEffectService) *RaceEffectHandler {
	return &RaceEffectHandler{
		RaceEffectService: raceService,
	}
}
