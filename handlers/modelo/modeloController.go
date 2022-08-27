package modelo

import (
	"net/url"

	core "github.com/nenodias/millenium/core/domain"
	"github.com/nenodias/millenium/core/domain/utils"
	"github.com/nenodias/millenium/handlers"

	domain "github.com/nenodias/millenium/core/domain/modelo"
)

type ModeloController struct {
	handlers.Controller[domain.Modelo, domain.ModeloFilter]
}

func NewModeloController(service *domain.ModeloService) *ModeloController {
	return &ModeloController{
		Controller: handlers.Controller[domain.Modelo, domain.ModeloFilter]{
			Service:    *service,
			GetFilters: GetFilters,
			SetModelId: func(id int64, model *domain.Modelo) {
				model.Id = id
			},
		},
	}
}

func GetFilters(query url.Values) domain.ModeloFilter {
	page := utils.StringToInt(query.Get("page"), utils.DEFAULT_PAGE)
	size := utils.StringToInt(query.Get("size"), utils.DEFAULT_SIZE)
	sortColumn := utils.StringNormalized(query.Get("sortColumn"), "id")
	sortDirection := core.GetSortDirection(query.Get("sortDirection"))
	nome := utils.StringNormalized(query.Get("nome"), "")
	idModelo := utils.StringToInt64(query.Get("id_monta"), 0)
	return domain.ModeloFilter{
		Nome:     nome,
		IdModelo: idModelo,
		Pageable: core.Pageable{
			PageSize: size, PageNumber: page,
			Sort: core.SortRequest{
				SortColumn: sortColumn, SortDirection: sortDirection,
			},
		},
	}
}
