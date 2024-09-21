package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"vendetta/internal/domain"
	"vendetta/internal/domain/entities"
)

func GetDefaultsFilter(filter *entities.Filter, prefix ...string) *entities.Filter {
	if filter.Limit == 0 {
		filter.Limit = 100
	}
	if filter.Order == "" {
		if len(prefix) > 0 {
			filter.Order = prefix[0] + ".id desc"
		} else {
			filter.Order = "id desc"
		}
	}

	return &entities.Filter{
		Limit:  filter.Limit,
		Offset: filter.Offset,
		Order:  filter.Order,
	}
}

func GetDefaultsFilterFromQuery(ctx *gin.Context, prefix ...string) *entities.Filter {
	var limitInt int
	var offsetInt int

	limit := ctx.Query("limit")
	if limit != "" {
		li, err := strconv.Atoi(limit)
		if err != nil {
			ErrorResponseHandler(ctx, domain.NewUnexpectedError("internal server error"))
			return nil
		}
		limitInt = li
	}
	offset := ctx.Query("offset")
	if offset != "" {
		oi, err := strconv.Atoi(offset)
		if err != nil {
			ErrorResponseHandler(ctx, domain.NewUnexpectedError("internal server error"))
			return nil
		}
		offsetInt = oi
	}

	order := ctx.Query("order")
	if order == "" {
		order = "id desc"
	}
	if len(prefix) > 0 {
		order = prefix[0] + "." + order
	}

	f := &entities.Filter{
		Limit:  limitInt,
		Offset: offsetInt,
		Order:  order,
	}

	return GetDefaultsFilter(f)
}
