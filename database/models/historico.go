package models

import "time"

type TipoHistorico string

const (
	ORDEM_SERVICO TipoHistorico = "O.S."
	ORCAMENTO     TipoHistorico = "Orç."
)

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

type Historico struct {
	Id          int64         `xorm:"'sequencia' bigint pk autoincr not null"`
	IdCliente   int64         `xorm:"'codigo_cliente' bigint not null"`
	IdVeiculo   int64         `xorm:"'codveiculo' bigint not null"`
	IdTecnico   int64         `xorm:"'tecnico' bigint"`
	NumeroOrdem int           `xorm:"'nr_ordem' int"`
	Placa       string        `xorm:"'placa' varchar(8)"`
	Sistema     int           `xorm:"'sistema' int"`
	Data        time.Time     `xorm:"'data' timestamp"`
	Tipo        TipoHistorico `xorm:"'tipo' varchar(4)"`
	ValorTotal  float64       `xorm:"'valor_total' double"`
	Observacao  string        `xorm:"'obs' varchar(500)"`
}

func (p *Historico) TableName() string {
	return "historico"
}

type HistoricoItem struct {
	Id          int64   `xorm:"'id' bigint pk autoincr not null"`
	IdHistorico int64   `xorm:"'sequencia' bigint not null"`
	Ordem       int64   `xorm:"'item' bigint not null"`
	Tipo        string  `xorm:"'tipo' varchar(1)"`
	Descricao   string  `xorm:"'tipo' varchar(1)"`
	Quantidade  int     `xorm:"'qtd' int"`
	Valor       float64 `xorm:"'valor' double"`
}

func (p *HistoricoItem) TableName() string {
	return "historico_item"
}
