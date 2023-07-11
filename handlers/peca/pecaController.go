package peca

import (
	"context"
	"net/url"

	core "github.com/nenodias/millenium/core/domain"
	"github.com/nenodias/millenium/core/domain/utils"
	"github.com/nenodias/millenium/handlers"

	domain "github.com/nenodias/millenium/core/domain/peca"
)

type PecaController struct {
	handlers.Controller[domain.Peca, domain.PecaFilter]
}

func NewController(service *domain.PecaService) *PecaController {
	return &PecaController{
		Controller: handlers.Controller[domain.Peca, domain.PecaFilter]{
			Service:    *service,
			GetFilters: GetFilters,
			SetModelId: func(id int64, model *domain.Peca) {
				model.Id = id
			},
		},
	}
}

func GetFilters(ctx context.Context, query url.Values) domain.PecaFilter {
	page := utils.StringToInt(query.Get("page"), utils.DEFAULT_PAGE)
	size := utils.StringToInt(query.Get("size"), utils.DEFAULT_SIZE)
	sortColumn := utils.StringNormalized(query.Get("sortColumn"), "id")
	sortDirection := core.GetSortDirection(query.Get("sortDirection"))
	descricao := utils.StringNormalized(query.Get("descricao"), "")
	return domain.PecaFilter{
		Descricao: descricao,
		Pageable: core.Pageable{
			PageSize: size, PageNumber: page,
			Sort: core.SortRequest{
				SortColumn: sortColumn, SortDirection: sortDirection,
			},
		},
	}
}
