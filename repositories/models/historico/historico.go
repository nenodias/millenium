package historico

import (
	"time"

	domain "github.com/nenodias/millenium/core/domain/historico"
	models "github.com/nenodias/millenium/repositories/models"
	"github.com/rs/zerolog/log"
	"xorm.io/xorm"
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
	Items       []HistoricoItem      `xorm:"-"`
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

type HistoricoRepository struct {
	models.GenericRepository[domain.Historico, domain.HistoricoFilter, Historico]
}

func NewService(engine *xorm.Engine) domain.HistoricoService {
	repository := HistoricoRepository{
		GenericRepository: models.GenericRepository[domain.Historico, domain.HistoricoFilter, Historico]{
			DB:             engine,
			MapperToDTO:    mapperToDTO,
			MapperToEntity: mapperToEntity,
			CopyToDto:      copyToDto,
			HasWhere:       hasWhere,
			DoWhere:        doWhere,
			AfterFind:      AfterFind,
		},
	}
	return domain.HistoricoService(&repository)
}

func AfterFind(gr *models.GenericRepository[domain.Historico, domain.HistoricoFilter, Historico], m *Historico) {
	if m.Id != 0 {
		model := new(HistoricoItem)
		rows, err := gr.DB.Where("sequencia = ?", m.Id).Rows(model)
		if err != nil {
			log.Error().Msgf("Error searching items: %s", err.Error())
		}
		defer rows.Close()
		m.Items = make([]HistoricoItem, 0)
		for rows.Next() {
			err = rows.Scan(model)
			if err != nil {
				log.Error().Msgf("Error scanning item: %s", err.Error())
			}
			m.Items = append(m.Items, *model)
		}
	}
}

func hasWhere(filter *domain.HistoricoFilter) bool {
	hasData := !filter.Data.IsZero()
	hasTipo := filter.Tipo != nil
	hasNumeroOrdem := filter.IdCliente != int64(0)
	hasIdCliente := filter.IdCliente != int64(0)
	hasIdVeiculo := filter.IdVeiculo != int64(0)
	hasIdTecnico := filter.IdTecnico != int64(0)
	return hasData || hasTipo || hasNumeroOrdem || hasIdCliente || hasIdVeiculo || hasIdTecnico
}

func doWhere(query *xorm.Session, filter *domain.HistoricoFilter) *xorm.Session {
	hasData := !filter.Data.IsZero()
	hasTipo := filter.Tipo != nil
	hasNumeroOrdem := filter.IdCliente != int64(0)
	hasIdCliente := filter.IdCliente != int64(0)
	hasIdVeiculo := filter.IdVeiculo != int64(0)
	hasIdTecnico := filter.IdTecnico != int64(0)
	where := query.Where("1 = ?", 1)
	if hasData {
		where = where.And("data = ?", filter.Data)
	}
	if hasTipo {
		where = where.And("tipo = ?", filter.Tipo)
	}
	if hasNumeroOrdem {
		where = where.And("nr_ordem = ?", filter.NumeroOrdem)
	}
	if hasIdCliente {
		where = where.And("codigo_cliente = ?", filter.IdCliente)
	}
	if hasIdVeiculo {
		where = where.And("codveiculo = ?", filter.IdVeiculo)
	}
	if hasIdTecnico {
		where = where.And("tecnico = ?", filter.IdTecnico)
	}
	return where
}

func mapperToEntity(dto *domain.Historico) *Historico {
	entity := new(Historico)
	copyToEntity(dto, entity)
	return entity
}

func mapperToDTO(entity *Historico) *domain.Historico {
	dto := new(domain.Historico)
	copyToDto(entity, dto)
	return dto
}

func mapperToEntityItem(dtos []domain.HistoricoItem) []HistoricoItem {
	items := make([]HistoricoItem, 0)
	for _, dto := range dtos {
		item := new(HistoricoItem)
		copyToEntityItem(&dto, item)
		items = append(items, *item)
	}
	return items
}

func mapperToDTOItem(dtos []HistoricoItem) []domain.HistoricoItem {
	items := make([]domain.HistoricoItem, 0)
	for _, dto := range dtos {
		item := new(domain.HistoricoItem)
		copyToDtoItem(&dto, item)
		items = append(items, *item)
	}
	return items
}

func copyToEntity(source *domain.Historico, destiny *Historico) {
	destiny.Id = source.Id
	destiny.IdCliente = source.IdCliente
	destiny.IdVeiculo = source.IdVeiculo
	destiny.IdTecnico = source.IdTecnico
	destiny.NumeroOrdem = source.NumeroOrdem
	destiny.Placa = source.Placa
	destiny.Sistema = source.Sistema
	destiny.Data = source.Data
	destiny.Tipo = source.Tipo
	destiny.ValorTotal = source.ValorTotal
	destiny.Observacao = source.Observacao
	destiny.Items = mapperToEntityItem(source.Items)
}

func copyToDto(source *Historico, destiny *domain.Historico) {
	destiny.Id = source.Id
	destiny.IdCliente = source.IdCliente
	destiny.IdVeiculo = source.IdVeiculo
	destiny.IdTecnico = source.IdTecnico
	destiny.NumeroOrdem = source.NumeroOrdem
	destiny.Placa = source.Placa
	destiny.Sistema = source.Sistema
	destiny.Data = source.Data
	destiny.Tipo = source.Tipo
	destiny.ValorTotal = source.ValorTotal
	destiny.Observacao = source.Observacao
	destiny.Items = mapperToDTOItem(source.Items)
}

func copyToEntityItem(source *domain.HistoricoItem, destiny *HistoricoItem) {
	destiny.Id = source.Id
	destiny.IdHistorico = source.IdHistorico
	destiny.Ordem = source.Ordem
	destiny.Tipo = source.Tipo
	destiny.Descricao = source.Descricao
	destiny.Quantidade = source.Quantidade
	destiny.Valor = source.Valor
}

func copyToDtoItem(source *HistoricoItem, destiny *domain.HistoricoItem) {
	destiny.Id = source.Id
	destiny.IdHistorico = source.IdHistorico
	destiny.Ordem = source.Ordem
	destiny.Tipo = source.Tipo
	destiny.Descricao = source.Descricao
	destiny.Quantidade = source.Quantidade
	destiny.Valor = source.Valor
}
