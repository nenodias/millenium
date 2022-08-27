package historico

import (
	"time"

	domain "github.com/nenodias/millenium/core/domain/historico"
)

type Historico struct {
	Id          int64                `xorm:"'sequencia' bigint pk autoincr not null"`
	IdCliente   int64                `xorm:"'codigo_cliente' bigint not null"`
	IdVeiculo   int64                `xorm:"'codveiculo' bigint not null"`
	IdTecnico   int64                `xorm:"'tecnico' bigint"`
	NumeroOrdem int                  `xorm:"'nr_ordem' int"`
	Placa       string               `xorm:"'placa' varchar(8)"`
	Sistema     int                  `xorm:"'sistema' int"`
	Data        time.Time            `xorm:"'data' timestamp"`
	Tipo        domain.TipoHistorico `xorm:"'tipo' varchar(4)"`
	ValorTotal  float64              `xorm:"'valor_total' double"`
	Observacao  string               `xorm:"'obs' varchar(500)"`
}

func (p *Historico) TableName() string {
	return "historico"
}

type HistoricoItem struct {
	Id          int64                    `xorm:"'id' bigint pk autoincr not null"`
	IdHistorico int64                    `xorm:"'sequencia' bigint not null"`
	Ordem       int64                    `xorm:"'item' bigint not null"`
	Tipo        domain.TipoHistoricoItem `xorm:"'tipo' varchar(1)"`
	Descricao   string                   `xorm:"'historico' varchar(75)"`
	Quantidade  int                      `xorm:"'qtd' int"`
	Valor       float64                  `xorm:"'valor' double"`
}

func (p *HistoricoItem) TableName() string {
	return "historico_item"
}

type HistoricoVistoria struct {
	Id               int64   `xorm:"'sequencia' bigint pk autoincr not null"`
	IdVeiculo        int     `xorm:"'carrovistoria' int"`
	NivelCombustivel int64   `xorm:"'nivelcomb' int"`
	Kilometragem     float64 `xorm:"'kilometragem' double"`
	TocaFitas        uint8   `xorm:"'tocafitas' smallint"`
	Cd               uint8   `xorm:"'cd' smallint"`
	Disqueteira      uint8   `xorm:"'disqueteira' smallint"`
	Antena           uint8   `xorm:"'antena' smallint"`
	Calotas          uint8   `xorm:"'calotas' smallint"`
	Triangulo        uint8   `xorm:"'triangulo' smallint"`
	Macaco           uint8   `xorm:"'macaco' smallint"`
	Estepe           uint8   `xorm:"'estepe' smallint"`
	Outro1           uint8   `xorm:"'outro1' smallint"`
	DescricaoOutro   string  `xorm:"'outro1descr' varchar(20)"`
	Outro2           uint8   `xorm:"'outro2' smallint"`
	DescricaoOutro2  string  `xorm:"'outro2descr' varchar(20)"`
	Observacao       string  `xorm:"'obs' varchar(500)"`
}

func (p *HistoricoVistoria) TableName() string {
	return "vistoria"
}
