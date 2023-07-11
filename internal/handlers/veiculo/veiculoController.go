package Veiculo

import (
	"context"
	"net/url"

	core "github.com/nenodias/millenium/internal/core/domain"
	"github.com/nenodias/millenium/internal/core/domain/utils"
	"github.com/nenodias/millenium/internal/handlers"

	domain "github.com/nenodias/millenium/internal/core/domain/veiculo"
)

type VeiculoController struct {
	handlers.Controller[domain.Veiculo, domain.VeiculoFilter]
}

func NewController(service *domain.VeiculoService) *VeiculoController {
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

func GetFilters(ctx context.Context, query url.Values) domain.VeiculoFilter {
	page := utils.StringToInt(query.Get("page"), utils.DEFAULT_PAGE)
	size := utils.StringToInt(query.Get("size"), utils.DEFAULT_SIZE)
	sortColumn := utils.StringNormalized(query.Get("sortColumn"), "id")
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
