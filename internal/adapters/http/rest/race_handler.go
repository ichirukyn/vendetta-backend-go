package rest

import (
	"github.com/gin-gonic/gin"
	"vendetta/internal/adapters/http/rest/dto"
	"vendetta/internal/domain"
	"vendetta/internal/usecases"
	"vendetta/pkg/utils"
)

// RaceHandler хэндлер для работы с пользователями
type RaceHandler struct {
	RaceService *usecases.RaceService
}

// GetAll хэндлер для получения пользователей (entities.Race)
func (h *RaceHandler) GetAll(ctx *gin.Context) {
	filter := utils.GetDefaultsFilterFromQuery(ctx)

	var payload dto.FindRacesDTO
	if err := ctx.ShouldBindQuery(&payload); err != nil {
		utils.NewBadRequestError(ctx, domain.NewBadRequestError("bad request data"))
		return
	}

	query, err := utils.StructToMap(&payload)
	if err != nil {
		utils.NewBadRequestError(ctx, domain.NewBadRequestError("bad request data"))
		return
	}

	races, dErr := h.RaceService.GetAll(filter, &query)
	if dErr != nil {
		utils.NewBadRequestError(ctx, dErr)
		return
	}

	utils.ResponseSuccessHandler(ctx, races)
}

// GetByID хэндлер для получения пользователя по id (entities.Race)
func (h *RaceHandler) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")

	races, dErr := h.RaceService.GetByID(id)
	if dErr != nil {
		utils.NewBadRequestError(ctx, dErr)
		return
	}

	utils.ResponseSuccessHandler(ctx, races)
}

// NewRaceHandler инициализатор UserHandler
func NewRaceHandler(raceService *usecases.RaceService) *RaceHandler {
	return &RaceHandler{
		RaceService: raceService,
	}
}
