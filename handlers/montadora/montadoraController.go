package montadora

import (
	"context"
	"net/url"

	core "github.com/nenodias/millenium/core/domain"
	"github.com/nenodias/millenium/core/domain/utils"
	"github.com/nenodias/millenium/handlers"

	domain "github.com/nenodias/millenium/core/domain/montadora"
)

type MontadoraController struct {
	handlers.Controller[domain.Montadora, domain.MontadoraFilter]
}

func NewController(service *domain.MontadoraService) *MontadoraController {
	return &MontadoraController{
		Controller: handlers.Controller[domain.Montadora, domain.MontadoraFilter]{
			Service:    *service,
			GetFilters: GetFilters,
			SetModelId: func(id int64, model *domain.Montadora) {
				model.Id = id
			},
		},
	}
}

func GetFilters(ctx context.Context, query url.Values) domain.MontadoraFilter {
	page := utils.StringToInt(query.Get("page"), utils.DEFAULT_PAGE)
	size := utils.StringToInt(query.Get("size"), utils.DEFAULT_SIZE)
	sortColumn := utils.StringNormalized(query.Get("sortColumn"), "id")
	sortDirection := core.GetSortDirection(query.Get("sortDirection"))
	nome := utils.StringNormalized(query.Get("nome"), "")
	return domain.MontadoraFilter{
		Nome: nome,
		Pageable: core.Pageable{
			PageSize: size, PageNumber: page,
			Sort: core.SortRequest{
				SortColumn: sortColumn, SortDirection: sortDirection,
			},
		},
	}
}
