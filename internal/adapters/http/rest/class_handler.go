package rest

import (
	"github.com/gin-gonic/gin"
	"vendetta/internal/adapters/factory"
	"vendetta/internal/adapters/http/rest/dto"
	"vendetta/internal/domain"
	"vendetta/internal/usecases"
	"vendetta/pkg/utils"
)

// ClassHandler хэндлер для работы с пользователями
type ClassHandler struct {
	ClassService *usecases.ClassService
}

// Create
//
//	@Tags		classes
//	@Param		classes	body	dto.CreateClassDTO	true	"field"
//
//	@Success	200		{array}	entities.Class
//	@Router		/classes [post]
func (h *ClassHandler) Create(ctx *gin.Context) {
	var body dto.CreateClassDTO
	if err := ctx.ShouldBindJSON(&body); err != nil {
		utils.NewBadRequestError(ctx, domain.NewBadRequestError("bad request data"))
		return
	}

	classFactory := factory.NewClassFactory()
	class := classFactory.CreateClassFromCreateDTO(body)
	if err := h.ClassService.Create(class); err != nil {
		utils.NewBadRequestError(ctx, err)
		return
	}

	utils.ResponseSuccessHandler(ctx, class)
}

// GetAll
//
//	@Tags		classes
//
//	@Success	200	{array}	[]entities.Class
//	@Router		/classes [get]
func (h *ClassHandler) GetAll(ctx *gin.Context) {
	filter := utils.GetDefaultsFilterFromQuery(ctx)

	var payload dto.FindClassesDTO
	if err := ctx.ShouldBindQuery(&payload); err != nil {
		utils.NewBadRequestError(ctx, domain.NewBadRequestError("bad request data"))
		return
	}

	query, err := utils.StructToMap(&payload)
	if err != nil {
		utils.NewBadRequestError(ctx, domain.NewBadRequestError("bad request data"))
		return
	}

	classs, dErr := h.ClassService.GetAll(filter, &query)
	if dErr != nil {
		utils.NewBadRequestError(ctx, dErr)
		return
	}

	utils.ResponseSuccessHandler(ctx, classs)
}

// GetByID
//
//	@Tags		classes
//
//	@Success	200	{array}	entities.Class
//	@Router		/classes/:id [get]
func (h *ClassHandler) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")

	classs, dErr := h.ClassService.GetByID(id)
	if dErr != nil {
		utils.NewBadRequestError(ctx, dErr)
		return
	}

	utils.ResponseSuccessHandler(ctx, classs)
}

// Update
//
//	@Tags		classes
//	@Param		classes	body	dto.UpdateClassDTO	true	"field"
//
//	@Success	200		{array}	entities.Class
//	@Router		/classes/:id [put]
func (h *ClassHandler) Update(ctx *gin.Context) {
	id := ctx.Param("id")

	var body dto.UpdateClassDTO
	if err := ctx.ShouldBindJSON(&body); err != nil {
		utils.NewBadRequestError(ctx, domain.NewBadRequestError("bad request data"))
		return
	}

	classFactory := factory.NewClassFactory()
	class := classFactory.CreateClassFromUpdateDTO(id, body)
	if err := h.ClassService.Update(class); err != nil {
		utils.NewBadRequestError(ctx, err)
		return
	}

	utils.ResponseSuccessHandler(ctx, class)
}

// Delete
//
//	@Tags		classes
//
//	@Success	200
//	@Router		/classes/:id [delete]
func (h *ClassHandler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := h.ClassService.Delete(id); err != nil {
		utils.NewBadRequestError(ctx, err)
		return
	}

	utils.ResponseSuccessHandler(ctx, true)
}

// NewClassHandler инициализатор ClassHandler
func NewClassHandler(classService *usecases.ClassService) *ClassHandler {
	return &ClassHandler{
		ClassService: classService,
	}
}
