package historico

import (
	"fmt"
	"time"

	"github.com/nenodias/millenium/core/domain"
)

type TipoHistorico string

const (
	ORDEM_SERVICO TipoHistorico = "O.S."
	ORCAMENTO     TipoHistorico = "Orç."
)

func GetTipoHistorico(t string) (TipoHistorico, error) {
	switch t {
	case ORCAMENTO.String():
		return ORCAMENTO, nil
	case ORDEM_SERVICO.String():
		return ORDEM_SERVICO, nil
	default:
		return "", fmt.Errorf("tipo: %s não existe", t)
	}
}

func (t TipoHistorico) String() string {
	switch t {
	case ORCAMENTO:
		return "Orçamento"
	case ORDEM_SERVICO:
		return "Ordem de Serviço"
	default:
		return ""
	}
}

type TipoHistoricoItem string

const (
	SERVICO TipoHistoricoItem = "S"
	FALHA   TipoHistoricoItem = "F"
	PECA    TipoHistoricoItem = "P"
)

func (t TipoHistoricoItem) String() string {
	switch t {
	case SERVICO:
		return "Serviço"
	case FALHA:
		return "Falha"
	case PECA:
		return "Peça"
	default:
		return ""
	}
}

type Historico struct {
	Id           int64           `json:"id"`
	IdCliente    int64           `json:"idCliente"`
	IdVeiculo    int64           `json:"idVeiculo"`
	IdTecnico    int64           `json:"idTecnico"`
	NumeroOrdem  int             `json:"numeroOrdem"`
	Placa        string          `json:"placa"`
	Sistema      int             `json:"sistema"`
	Data         time.Time       `json:"data"`
	Tipo         TipoHistorico   `json:"tipo"`
	ValorTotal   float64         `json:"valorTotal"`
	Observacao   string          `json:"observacao"`
	Items        []HistoricoItem `json:"items"`
	Kilometragem float64         `json:"km"`
}

func (t Historico) GetId() int64 {
	return t.Id
}

type HistoricoItem struct {
	Id          int64             `json:"id"`
	IdHistorico int64             `json:"idHistorico"`
	Ordem       int64             `json:"item"`
	Tipo        TipoHistoricoItem `json:"tipo"`
	Descricao   string            `json:"descricao"`
	Quantidade  int               `json:"quantidade"`
	Valor       float64           `json:"valor"`
}

func (t HistoricoItem) GetId() int64 {
	return t.Id
}

type HistoricoFilter struct {
	domain.Pageable
	NumeroOrdem int
	IdCliente   int64
	IdVeiculo   int64
	IdTecnico   int64
	Tipo        *TipoHistorico
	Data        time.Time
}

type HistoricoService interface {
	domain.Service[*Historico, *HistoricoFilter]
	FindOneForReport(int64) (*HistoricoReport, error)
}
