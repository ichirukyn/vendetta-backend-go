package rest

import (
	"github.com/gin-gonic/gin"
	"vendetta/internal/adapters/http/rest/dto"
	"vendetta/internal/domain"
	"vendetta/internal/usecases"
	"vendetta/pkg/utils"
)

// ClassEffectHandler хэндлер для работы с пользователями
type ClassEffectHandler struct {
	ClassEffectService *usecases.ClassEffectService
}

// GetAll
//
//	@Tags		classes
//
//	@Success	200	{array}	[]entities.ClassEffect
//	@Router		/classes/effects [get]
func (h *ClassEffectHandler) GetAll(ctx *gin.Context) {
	filter := utils.GetDefaultsFilterFromQuery(ctx)

	var payload dto.FindClassesEffectsDTO
	if err := ctx.ShouldBindQuery(&payload); err != nil {
		utils.NewBadRequestError(ctx, domain.NewBadRequestError("bad request data"))
		return
	}

	query, err := utils.StructToMap(&payload)
	if err != nil {
		utils.NewBadRequestError(ctx, domain.NewBadRequestError("bad request data"))
		return
	}

	classes, dErr := h.ClassEffectService.GetAll(filter, &query)
	if dErr != nil {
		utils.NewBadRequestError(ctx, dErr)
		return
	}

	utils.ResponseSuccessHandler(ctx, classes)
}

// GetByClassID
//
//	@Tags		classes
//
//	@Success	200	{array}	entities.ClassEffect
//	@Router		/classes/:id/effects [get]
func (h *ClassEffectHandler) GetByClassID(ctx *gin.Context) {
	id := ctx.Param("id")

	classes, dErr := h.ClassEffectService.GetByID(id)
	if dErr != nil {
		utils.NewBadRequestError(ctx, dErr)
		return
	}

	utils.ResponseSuccessHandler(ctx, classes)
}

// NewClassEffectHandler инициализатор UserHandler
func NewClassEffectHandler(raceService *usecases.ClassEffectService) *ClassEffectHandler {
	return &ClassEffectHandler{
		ClassEffectService: raceService,
	}
}
