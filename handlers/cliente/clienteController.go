package cliente

import (
	"net/url"

	core "github.com/nenodias/millenium/core/domain"
	"github.com/nenodias/millenium/core/domain/utils"
	"github.com/nenodias/millenium/handlers"

	domain "github.com/nenodias/millenium/core/domain/cliente"
)

type ClienteController struct {
	handlers.Controller[domain.Cliente, domain.ClienteFilter]
}

func NewClienteController(service *domain.ClienteService) *ClienteController {
	return &ClienteController{
		Controller: handlers.Controller[domain.Cliente, domain.ClienteFilter]{
			Service:    *service,
			GetFilters: GetFilters,
			SetModelId: func(id int64, model *domain.Cliente) {
				model.Id = id
			},
		},
	}
}

func GetFilters(query url.Values) domain.ClienteFilter {
	page := utils.StringToInt(query.Get("page"), utils.DEFAULT_PAGE)
	size := utils.StringToInt(query.Get("size"), utils.DEFAULT_SIZE)
	sortColumn := utils.StringNormalized(query.Get("sortColumn"), "id")
	sortDirection := core.GetSortDirection(query.Get("sortDirection"))
	nome := utils.StringNormalized(query.Get("nome"), "")
	telefone := utils.StringNormalized(query.Get("telefone"), "")
	celular := utils.StringNormalized(query.Get("celular"), "")
	return domain.ClienteFilter{
		Nome:     nome,
		Telefone: telefone,
		Celular:  celular,
		Pageable: core.Pageable{
			PageSize: size, PageNumber: page,
			Sort: core.SortRequest{
				SortColumn: sortColumn, SortDirection: sortDirection,
			},
		},
	}
}
