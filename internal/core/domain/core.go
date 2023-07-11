package domain

import "context"

type PagebleContent[T any] struct {
	Content       []T      `json:"content"`
	Pageable      Pageable `json:"pageable"`
	TotalPages    int64    `json:"totalPages"`
	TotalElements int64    `json:"totalElements"`
	Size          int      `json:"size"`
	Number        int      `json:"number"`
}

type Identifiable interface {
	GetId() int64
}

type SortDirection string

const (
	ASC  SortDirection = "ASC"
	DESC SortDirection = "DESC"
)

func GetSortDirection(value string) SortDirection {
	switch value {
	case string(ASC):
		return ASC
	case string(DESC):
		return DESC
	default:
		return ASC
	}
}

type SortRequest struct {
	SortColumn    string        `json:"sortColumn"`
	SortDirection SortDirection `json:"sortDirection"`
}

type PageableFilter interface {
	GetSort() SortRequest
	GetPageSize() int
	GetPageNumber() int
}

type Pageable struct {
	Sort       SortRequest `json:"sort"`
	PageSize   int         `json:"pageSize"`
	PageNumber int         `json:"pageNumber"`
}

func (p Pageable) GetSort() SortRequest {
	return p.Sort
}

func (p Pageable) GetPageSize() int {
	return p.PageSize
}

func (p Pageable) GetPageNumber() int {
	return p.PageNumber
}

type Service[T any, F any] interface {
	FindMany(context.Context, F) (PagebleContent[T], error)
	FindOne(context.Context, int64) (T, error)
	DeleteOne(context.Context, int64) (bool, error)
	Save(context.Context, T) (bool, error)
}
