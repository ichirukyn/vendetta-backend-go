package rest

import (
	"github.com/gin-gonic/gin"
	"vendetta/internal/adapters/factory"
	"vendetta/internal/adapters/http/rest/dto"
	"vendetta/internal/domain"
	"vendetta/internal/usecases"
	"vendetta/pkg/utils"
)

type SkillHandler struct {
	SkillService *usecases.SkillService
}

func (h *SkillHandler) Create(ctx *gin.Context) {
	var body dto.CreateSkillDTO
	if err := ctx.ShouldBindJSON(&body); err != nil {
		utils.NewBadRequestError(ctx, domain.NewBadRequestError("bad request data"))
		return
	}

	skillFactory := factory.NewSkillFactory()
	skill := skillFactory.CreateSkillFromCreateDTO(body)
	if err := h.SkillService.Create(skill); err != nil {
		utils.NewBadRequestError(ctx, err)
		return
	}

	utils.ResponseSuccessHandler(ctx, skill)
}

func (h *SkillHandler) GetAll(ctx *gin.Context) {
	filter := utils.GetDefaultsFilterFromQuery(ctx)

	var payload dto.FindSkillsDTO
	if err := ctx.ShouldBindQuery(&payload); err != nil {
		utils.NewBadRequestError(ctx, domain.NewBadRequestError("bad request data"))
		return
	}

	query, err := utils.StructToMap(&payload)
	if err != nil {
		utils.NewBadRequestError(ctx, domain.NewBadRequestError("bad request data"))
		return
	}

	skills, dErr := h.SkillService.GetAll(filter, &query)
	if dErr != nil {
		utils.NewBadRequestError(ctx, dErr)
		return
	}

	utils.ResponseSuccessHandler(ctx, skills)
}

func (h *SkillHandler) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")

	skills, dErr := h.SkillService.GetByID(id)
	if dErr != nil {
		utils.NewBadRequestError(ctx, dErr)
		return
	}

	utils.ResponseSuccessHandler(ctx, skills)
}

func (h *SkillHandler) Update(ctx *gin.Context) {
	id := ctx.Param("id")

	var body dto.UpdateSkillDTO
	if err := ctx.ShouldBindJSON(&body); err != nil {
		utils.NewBadRequestError(ctx, domain.NewBadRequestError("bad request data"))
		return
	}

	skillFactory := factory.NewSkillFactory()
	skill := skillFactory.CreateSkillFromUpdateDTO(id, body)
	if err := h.SkillService.Update(skill); err != nil {
		utils.NewBadRequestError(ctx, err)
		return
	}

	utils.ResponseSuccessHandler(ctx, skill)
}

func (h *SkillHandler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := h.SkillService.Delete(id); err != nil {
		utils.NewBadRequestError(ctx, err)
		return
	}

	utils.ResponseSuccessHandler(ctx, true)
}

// NewSkillHandler инициализатор SkillHandler
func NewSkillHandler(skillService *usecases.SkillService) *SkillHandler {
	return &SkillHandler{
		SkillService: skillService,
	}
}
