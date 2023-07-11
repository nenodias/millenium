package falha

import (
	"context"
	"net/url"

	core "github.com/nenodias/millenium/core/domain"
	"github.com/nenodias/millenium/core/domain/utils"
	"github.com/nenodias/millenium/handlers"

	domain "github.com/nenodias/millenium/core/domain/falha"
)

type FalhaController struct {
	handlers.Controller[domain.Falha, domain.FalhaFilter]
}

func NewController(service *domain.FalhaService) *FalhaController {
	return &FalhaController{
		Controller: handlers.Controller[domain.Falha, domain.FalhaFilter]{
			Service:    *service,
			GetFilters: GetFilters,
			SetModelId: func(id int64, model *domain.Falha) {
				model.Id = id
			},
		},
	}
}

func GetFilters(ctx context.Context, query url.Values) domain.FalhaFilter {
	page := utils.StringToInt(query.Get("page"), utils.DEFAULT_PAGE)
	size := utils.StringToInt(query.Get("size"), utils.DEFAULT_SIZE)
	sortColumn := utils.StringNormalized(query.Get("sortColumn"), "id")
	sortDirection := core.GetSortDirection(query.Get("sortDirection"))
	descricao := utils.StringNormalized(query.Get("descricao"), "")
	return domain.FalhaFilter{
		Descricao: descricao,
		Pageable: core.Pageable{
			PageSize: size, PageNumber: page,
			Sort: core.SortRequest{
				SortColumn: sortColumn, SortDirection: sortDirection,
			},
		},
	}
}
