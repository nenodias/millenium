package Veiculo

import (
	"net/url"

	core "github.com/nenodias/millenium/core/domain"
	"github.com/nenodias/millenium/core/domain/utils"
	"github.com/nenodias/millenium/handlers"

	domain "github.com/nenodias/millenium/core/domain/veiculo"
)

type VeiculoController struct {
	handlers.Controller[domain.Veiculo, domain.VeiculoFilter]
}

func NewVeiculoController(service *domain.VeiculoService) *VeiculoController {
	return &VeiculoController{
		Controller: handlers.Controller[domain.Veiculo, domain.VeiculoFilter]{
			Service:    *service,
			GetFilters: GetFilters,
			SetModelId: func(id int64, model *domain.Veiculo) {
				model.Id = id
			},
		},
	}
}

func GetFilters(query url.Values) domain.VeiculoFilter {
	page := utils.StringToInt(query.Get("page"), utils.DEFAULT_PAGE)
	size := utils.StringToInt(query.Get("size"), utils.DEFAULT_SIZE)
	sortColumn := utils.StringNormalized(query.Get("sortColumn"), "codveiculo")
	sortDirection := core.GetSortDirection(query.Get("sortDirection"))
	placa := utils.StringNormalized(query.Get("placa"), "")
	idCliente := utils.StringToInt64(query.Get("idCliente"), 0)
	idModelo := utils.StringToInt64(query.Get("idModelo"), 0)
	return domain.VeiculoFilter{
		Placa:     placa,
		IdCliente: idCliente,
		IdModelo:  idModelo,
		Pageable: core.Pageable{
			PageSize: size, PageNumber: page,
			Sort: core.SortRequest{
				SortColumn: sortColumn, SortDirection: sortDirection,
			},
		},
	}
}
