package servico

import (
	"net/url"

	core "github.com/nenodias/millenium/core/domain"
	"github.com/nenodias/millenium/core/domain/utils"
	"github.com/nenodias/millenium/handlers"

	domain "github.com/nenodias/millenium/core/domain/servico"
)

type ServicoController struct {
	handlers.Controller[domain.Servico, domain.ServicoFilter]
}

func NewController(service *domain.ServicoService) *ServicoController {
	return &ServicoController{
		Controller: handlers.Controller[domain.Servico, domain.ServicoFilter]{
			Service:    *service,
			GetFilters: GetFilters,
			SetModelId: func(id int64, model *domain.Servico) {
				model.Id = id
			},
		},
	}
}

func GetFilters(query url.Values) domain.ServicoFilter {
	page := utils.StringToInt(query.Get("page"), utils.DEFAULT_PAGE)
	size := utils.StringToInt(query.Get("size"), utils.DEFAULT_SIZE)
	sortColumn := utils.StringNormalized(query.Get("sortColumn"), "id")
	sortDirection := core.GetSortDirection(query.Get("sortDirection"))
	descricao := utils.StringNormalized(query.Get("descricao"), "")
	return domain.ServicoFilter{
		Descricao: descricao,
		Pageable: core.Pageable{
			PageSize: size, PageNumber: page,
			Sort: core.SortRequest{
				SortColumn: sortColumn, SortDirection: sortDirection,
			},
		},
	}
}
