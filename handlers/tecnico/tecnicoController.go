package tecnico

import (
	"context"
	"net/url"

	core "github.com/nenodias/millenium/core/domain"
	"github.com/nenodias/millenium/core/domain/utils"
	"github.com/nenodias/millenium/handlers"

	domain "github.com/nenodias/millenium/core/domain/tecnico"
)

type TecnicoController struct {
	handlers.Controller[domain.Tecnico, domain.TecnicoFilter]
}

func NewController(service *domain.TecnicoService) *TecnicoController {
	return &TecnicoController{
		Controller: handlers.Controller[domain.Tecnico, domain.TecnicoFilter]{
			Service:    *service,
			GetFilters: GetFilters,
			SetModelId: func(id int64, model *domain.Tecnico) {
				model.Id = id
			},
		},
	}
}

func GetFilters(ctx context.Context, query url.Values) domain.TecnicoFilter {
	page := utils.StringToInt(query.Get("page"), utils.DEFAULT_PAGE)
	size := utils.StringToInt(query.Get("size"), utils.DEFAULT_SIZE)
	sortColumn := utils.StringNormalized(query.Get("sortColumn"), "id")
	sortDirection := core.GetSortDirection(query.Get("sortDirection"))
	nome := utils.StringNormalized(query.Get("nome"), "")
	return domain.TecnicoFilter{
		Nome: nome,
		Pageable: core.Pageable{
			PageSize: size, PageNumber: page,
			Sort: core.SortRequest{
				SortColumn: sortColumn, SortDirection: sortDirection,
			},
		},
	}
}
