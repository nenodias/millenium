package historico

import (
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/mux"
	core "github.com/nenodias/millenium/core/domain"
	"github.com/nenodias/millenium/core/domain/utils"
	"github.com/nenodias/millenium/handlers"
	"github.com/rs/zerolog/log"

	domain "github.com/nenodias/millenium/core/domain/historico"
)

type HistoricoController struct {
	handlers.Controller[domain.Historico, domain.HistoricoFilter]
}

func NewController(service *domain.HistoricoService) *HistoricoController {
	return &HistoricoController{
		Controller: handlers.Controller[domain.Historico, domain.HistoricoFilter]{
			Service:    *service,
			GetFilters: GetFilters,
			SetModelId: func(id int64, model *domain.Historico) {
				model.Id = id
			},
		},
	}
}

func (hc *HistoricoController) GetReport(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	txtId := params["id"]
	id := utils.StringToInt64(txtId, 0)
	model, err := hc.Service.FindOne(id)
	if err != nil {
		log.Error().Msg(err.Error())
		w.WriteHeader(500)
		return
	}
	if model != nil {
		//TODO Generate Report
		utils.WriteJson(model, w, 200, 500)
	} else {
		w.WriteHeader(204)
	}
}

func GetFilters(query url.Values) domain.HistoricoFilter {
	page := utils.StringToInt(query.Get("page"), utils.DEFAULT_PAGE)
	size := utils.StringToInt(query.Get("size"), utils.DEFAULT_SIZE)
	sortColumn := utils.StringNormalized(query.Get("sortColumn"), "id")
	sortDirection := core.GetSortDirection(query.Get("sortDirection"))
	data := utils.StringToDate(query.Get("texto"), time.Time{})
	numeroOrdem := utils.StringToInt(query.Get("numeroOrdem"), 0)
	idTecnico := utils.StringToInt64(query.Get("idTecnico"), 0)
	idCliente := utils.StringToInt64(query.Get("idCliente"), 0)
	idVeiculo := utils.StringToInt64(query.Get("idVeiculo"), 0)
	var tipo *domain.TipoHistorico
	if query.Get("tipo") != "" {
		newTipo, err := domain.GetTipoHistorico(query.Get("tipo"))
		if err != nil {
			tipo = nil
		} else {
			tipo = &newTipo
		}
	}
	return domain.HistoricoFilter{
		NumeroOrdem: numeroOrdem,
		IdTecnico:   idTecnico,
		IdCliente:   idCliente,
		IdVeiculo:   idVeiculo,
		Tipo:        tipo,
		Data:        data,
		Pageable: core.Pageable{
			PageSize: size, PageNumber: page,
			Sort: core.SortRequest{
				SortColumn: sortColumn, SortDirection: sortDirection,
			},
		},
	}
}
