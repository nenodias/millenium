package domain

type PagebleContent[T any] struct {
	Content       []T      `json:"content"`
	Pageable      Pageable `json:"pageable"`
	TotalPages    int64    `json:"totalPages"`
	TotalElements int64    `json:"totalElements"`
	Size          int      `json:"size"`
	Number        int      `json:"number"`
}

type SortDirection string

const (
	ASC  SortDirection = "ASC"
	DESC SortDirection = "DESC"
)

type SortRequest struct {
	SortColumn    string        `json:"sortColumn"`
	SortDirection SortDirection `json:"sortDirection"`
}

type Pageable struct {
	Sort       SortRequest `json:"sort"`
	PageSize   int         `json:"pageSize"`
	PageNumber int         `json:"pageNumber"`
}

type Service[T any, F any] interface {
	FindMany(F) (PagebleContent[T], error)
	FindOne(int64) (T, error)
	DeleteOne(int64) (bool, error)
	Save(T) (bool, error)
}
