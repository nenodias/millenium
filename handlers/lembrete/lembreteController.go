package lembrete

import (
	"context"
	"net/url"

	core "github.com/nenodias/millenium/core/domain"
	"github.com/nenodias/millenium/core/domain/utils"
	"github.com/nenodias/millenium/handlers"

	domain "github.com/nenodias/millenium/core/domain/lembrete"
)

type LembreteController struct {
	handlers.Controller[domain.Lembrete, domain.LembreteFilter]
}

func NewController(service *domain.LembreteService) *LembreteController {
	return &LembreteController{
		Controller: handlers.Controller[domain.Lembrete, domain.LembreteFilter]{
			Service:    *service,
			GetFilters: GetFilters,
			SetModelId: func(id int64, model *domain.Lembrete) {
				model.Id = id
			},
		},
	}
}

func GetFilters(ctx context.Context, query url.Values) domain.LembreteFilter {
	page := utils.StringToInt(query.Get("page"), utils.DEFAULT_PAGE)
	size := utils.StringToInt(query.Get("size"), utils.DEFAULT_SIZE)
	sortColumn := utils.StringNormalized(query.Get("sortColumn"), "id")
	sortDirection := core.GetSortDirection(query.Get("sortDirection"))
	texto := utils.StringNormalized(query.Get("texto"), "")
	idCliente := utils.StringToInt64(query.Get("idCliente"), 0)
	idVeiculo := utils.StringToInt64(query.Get("idVeiculo"), 0)
	return domain.LembreteFilter{
		Texto:     texto,
		IdCliente: idCliente,
		IdVeiculo: idVeiculo,
		Pageable: core.Pageable{
			PageSize: size, PageNumber: page,
			Sort: core.SortRequest{
				SortColumn: sortColumn, SortDirection: sortDirection,
			},
		},
	}
}
