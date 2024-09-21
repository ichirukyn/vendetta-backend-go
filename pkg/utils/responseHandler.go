package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vendetta/internal/app/config"
	"vendetta/internal/domain"
)

const (
	statusSuccess = "success"
	statusFail    = "failed"
)

func ResponseSuccessHandler(ctx *gin.Context, data interface{}) {
	ResponseHandler(ctx, http.StatusOK, data)
}

func ResponseHandler(ctx *gin.Context, code int, data interface{}) {
	ctx.AbortWithStatusJSON(code, gin.H{
		"status": statusSuccess,
		"data":   data,
	})
}

func ErrorResponseHandler(ctx *gin.Context, err *domain.AppError) {
	ctx.AbortWithStatusJSON(err.Code, gin.H{
		"status": statusFail,
		"code":   err.Code,
		"error":  err.Message,
	})
}

func NewNotFoundError(ctx *gin.Context, err *domain.AppError) {
	if err == nil {
		err = domain.NewNotFoundError(config.ErrNotFound)
	}

	ErrorResponseHandler(ctx, err)
}

func NewUnexpectedError(ctx *gin.Context, err *domain.AppError) {
	if err == nil {
		err = domain.NewUnexpectedError(config.ErrInternalServerError)
	}

	ErrorResponseHandler(ctx, err)
}

func NewForbiddenError(ctx *gin.Context, err *domain.AppError) {
	if err == nil {
		err = domain.NewForbiddenError(config.ErrForbiddenAccess)
	}

	ErrorResponseHandler(ctx, err)
}

func NewBadRequestError(ctx *gin.Context, err *domain.AppError) {
	if err == nil {
		err = domain.NewBadRequestError(config.ErrBadRequest)
	}

	ErrorResponseHandler(ctx, err)
}

func NewAlreadyExistError(ctx *gin.Context, err *domain.AppError) {
	if err == nil {
		err = domain.NewAlreadyExistError(config.ErrConflict)
	}

	ErrorResponseHandler(ctx, err)
}

func NewValidationError(ctx *gin.Context, err *domain.AppError) {
	if err == nil {
		err = domain.NewValidationError(config.ErrBadRequest)
	}

	ErrorResponseHandler(ctx, err)
}
