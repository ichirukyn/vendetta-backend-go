package rest

import (
	"github.com/gin-gonic/gin"
	"vendetta/internal/adapters/factory"
	"vendetta/internal/adapters/http/rest/dto"
	"vendetta/internal/domain"
	"vendetta/internal/usecases"
	"vendetta/pkg/utils"
)

// UserHandler хэндлер для работы с пользователями
type UserHandler struct {
	UserService *usecases.UserService
}

// Create хэндлер для создания пользователя (entities.User)
func (h *UserHandler) Create(ctx *gin.Context) {
	var body dto.CreateUserDTO
	if err := ctx.ShouldBindJSON(&body); err != nil {
		utils.NewBadRequestError(ctx, domain.NewBadRequestError("bad request data"))
		return
	}

	userFactory := factory.NewUserFactory()
	user := userFactory.CreateUserFromCreateDTO(body)
	if err := h.UserService.Create(user); err != nil {
		utils.NewBadRequestError(ctx, err)
		return
	}

	utils.ResponseSuccessHandler(ctx, user)
}

// GetAll хэндлер для получения пользователей (entities.User)
func (h *UserHandler) GetAll(ctx *gin.Context) {
	chatID := ctx.Query("chat_id")
	if len(chatID) != 0 {
		user, dErr := h.UserService.GetByChatID(chatID)
		if dErr != nil {
			utils.NewBadRequestError(ctx, dErr)
			return
		}

		utils.ResponseSuccessHandler(ctx, user)
		return
	}

	filter := utils.GetDefaultsFilterFromQuery(ctx)

	var payload dto.FindUsersDTO
	if err := ctx.ShouldBindQuery(&payload); err != nil {
		utils.NewBadRequestError(ctx, domain.NewBadRequestError("bad request data"))
		return
	}

	query, err := utils.StructToMap(&payload)
	if err != nil {
		utils.NewBadRequestError(ctx, domain.NewBadRequestError("bad request data"))
		return
	}

	users, dErr := h.UserService.GetAll(filter, &query)
	if dErr != nil {
		utils.NewBadRequestError(ctx, dErr)
		return
	}

	utils.ResponseSuccessHandler(ctx, users)
}

// GetByID хэндлер для получения пользователя по id (entities.User)
func (h *UserHandler) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")

	users, dErr := h.UserService.GetByID(id)
	if dErr != nil {
		utils.NewBadRequestError(ctx, dErr)
		return
	}

	utils.ResponseSuccessHandler(ctx, users)
}

// Update хэндлер для обновления пользователя по id (entities.User)
func (h *UserHandler) Update(ctx *gin.Context) {
	id := ctx.Param("id")

	var body dto.UpdateUserDTO
	if err := ctx.ShouldBindJSON(&body); err != nil {
		utils.NewBadRequestError(ctx, domain.NewBadRequestError("bad request data"))
		return
	}

	userFactory := factory.NewUserFactory()
	user := userFactory.CreateUserFromUpdateDTO(id, body)
	if err := h.UserService.Update(user); err != nil {
		utils.NewBadRequestError(ctx, err)
		return
	}

	utils.ResponseSuccessHandler(ctx, user)
}

// Delete хэндлер для удаления пользователя по id (entities.User)
func (h *UserHandler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := h.UserService.Delete(id); err != nil {
		utils.NewBadRequestError(ctx, err)
		return
	}

	utils.ResponseSuccessHandler(ctx, true)
}

// NewUserHandler инициализатор UserHandler
func NewUserHandler(userService *usecases.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}
