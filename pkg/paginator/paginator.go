package paginator

import (
	"context"
	"math"
)

type PaginationParams struct {
	Limit   int
	Current int
}

type PaginationOffset struct {
	Limit  int
	Offset int
}

type PaginationInfo struct {
	TotalPages  int
	TotalItems  int
	Limit       int
	CurrentPage int
	HasNext     bool
	HasPrev     bool
}
type PaginationInfoParams struct {
	TotalItems   int
	ItemsPerPage int
	CurrentPage  int
	Limit        int
}

type LimitAndOffSetOptions func(params *PaginationParams) PaginationParams

func GetLimitAndOffSet(opts ...LimitAndOffSetOptions) PaginationOffset {

	params := &PaginationParams{}

	for _, opt := range opts {
		opt(params)
	}

	offset := (params.Current - 1) * params.Limit
	return PaginationOffset{
		Limit:  params.Limit,
		Offset: offset,
	}
}

func WithLimit(ctx context.Context, limit int) LimitAndOffSetOptions {
	return func(params *PaginationParams) PaginationParams {
		if limit == 0 {
			params.Limit = 10
		} else {
			params.Limit = limit
		}
		return *params
	}
}

func WithCurrentPage(ctx context.Context, current int) LimitAndOffSetOptions {
	return func(params *PaginationParams) PaginationParams {
		if current == 0 {
			params.Current = 1
		} else {
			params.Current = current
		}
		return *params
	}
}

func GetPaginationInfo(params PaginationInfoParams) PaginationInfo {
	totalPages := int(math.Ceil(float64(params.TotalItems) / float64(params.Limit)))
	return PaginationInfo{
		TotalItems:  params.TotalItems,
		Limit:       params.ItemsPerPage,
		CurrentPage: params.CurrentPage,
		TotalPages:  totalPages,
		HasNext:     params.CurrentPage < totalPages,
		HasPrev:     params.CurrentPage > 1,
	}
}
