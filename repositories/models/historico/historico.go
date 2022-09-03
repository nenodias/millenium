package historico

import (
	"fmt"
	"time"

	domain "github.com/nenodias/millenium/core/domain/historico"
	models "github.com/nenodias/millenium/repositories/models"
	clienteModel "github.com/nenodias/millenium/repositories/models/cliente"
	modeloModel "github.com/nenodias/millenium/repositories/models/modelo"
	montadoraModel "github.com/nenodias/millenium/repositories/models/montadora"
	tecnicoModel "github.com/nenodias/millenium/repositories/models/tecnico"
	veiculoModel "github.com/nenodias/millenium/repositories/models/veiculo"
	"github.com/rs/zerolog/log"
	"xorm.io/xorm"
)

type Historico struct {
	Id          int64                `xorm:"'id' bigint pk autoincr not null"`
	IdCliente   int64                `xorm:"'id_cliente' bigint not null"`
	IdVeiculo   int64                `xorm:"'id_veiculo' bigint not null"`
	IdTecnico   *int64               `xorm:"'id_tecnico' bigint"`
	NumeroOrdem int64                `xorm:"'numero' int"`
	Placa       string               `xorm:"'placa' varchar(8)"`
	Sistema     int                  `xorm:"'sistema' int"`
	Data        time.Time            `xorm:"'data' timestamp"`
	Tipo        domain.TipoHistorico `xorm:"'tipo' varchar(4)"`
	ValorTotal  float64              `xorm:"'valor_total' double"`
	Observacao  string               `xorm:"'observacao' varchar(500)"`
	Items       []HistoricoItem      `xorm:"-"`
	Vistoria    HistoricoVistoria    `xorm:"-"`
}

func (p *Historico) TableName() string {
	return "historico"
}

type HistoricoItem struct {
	Id          int64                    `xorm:"'id' bigint pk autoincr not null"`
	IdHistorico int64                    `xorm:"'id_historico' bigint not null"`
	Ordem       int64                    `xorm:"'ordem' bigint not null"`
	Tipo        domain.TipoHistoricoItem `xorm:"'tipo' varchar(1)"`
	Descricao   string                   `xorm:"'historico' varchar(75)"`
	Quantidade  int                      `xorm:"'quantidade' int"`
	Valor       float64                  `xorm:"'valor' double"`
}

func (p *HistoricoItem) TableName() string {
	return "historico_item"
}

type HistoricoVistoria struct {
	Id               int64   `xorm:"'id_historico' bigint pk autoincr not null"`
	IdVeiculo        int64   `xorm:"'id_veiculo' bigint"`
	NivelCombustivel int64   `xorm:"'nivelcomb' int"`
	Kilometragem     float64 `xorm:"'km' double"`
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
	Observacao       string  `xorm:"'observacao' varchar(500)"`
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
			MapperToDTO:    MapperToDTO,
			MapperToEntity: MapperToEntity,
			CopyToDto:      copyToDto,
			HasWhere:       hasWhere,
			DoWhere:        doWhere,
			AfterFind:      AfterFind,
			AfterSave:      AfterSave,
			AfterUpdate:    AfterSave,
		},
	}
	return domain.HistoricoService(&repository)
}

func (hr *HistoricoRepository) FindOneForReport(id int64) (*domain.HistoricoReport, error) {
	report := new(domain.HistoricoReport)
	model, err := hr.FindOne(id)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}
	if model != nil {
		report.Historico = *model
		cliente := new(clienteModel.Cliente)
		modelo := new(modeloModel.Modelo)
		montadora := new(montadoraModel.Montadora)
		tecnico := new(tecnicoModel.Tecnico)
		veiculo := new(veiculoModel.Veiculo)
		if model.IdCliente != 0 {
			_, err = hr.DB.ID(model.IdCliente).Get(cliente)
			if err != nil {
				log.Error().Msg(err.Error())
				return nil, err
			}
			report.Cliente = *clienteModel.MapperToDTO(cliente)
		}
		if model.IdVeiculo != 0 {
			_, err = hr.DB.ID(model.IdVeiculo).Get(veiculo)
			if err != nil {
				log.Error().Msg(err.Error())
				return nil, err
			}
			report.Veiculo = *veiculoModel.MapperToDTO(veiculo)
			if veiculo.IdModelo != 0 {
				_, err = hr.DB.ID(veiculo.IdModelo).Get(modelo)
				if err != nil {
					log.Error().Msg(err.Error())
				}
				report.Modelo = *modeloModel.MapperToDTO(modelo)
				if modelo.IdMontadora != 0 {
					_, err = hr.DB.ID(modelo.IdMontadora).Get(montadora)
					if err != nil {
						log.Error().Msg(err.Error())
					}
					report.Montadora = *montadoraModel.MapperToDTO(montadora)
				}
			}
		}

		if model.IdTecnico != nil {
			_, err = hr.DB.ID(*model.IdTecnico).Get(tecnico)
			if err != nil {
				log.Error().Msg(err.Error())
			}
			report.Tecnico = *tecnicoModel.MapperToDTO(tecnico)
		}
	} else {
		return report, fmt.Errorf("registro com id: %d nao encontrado", id)
	}
	return report, nil
}

func AfterFind(gr *models.GenericRepository[domain.Historico, domain.HistoricoFilter, Historico], m *Historico) {
	if m.Id != 0 {
		model := new(HistoricoItem)
		rows, err := gr.DB.Where("id_historico = ?", m.Id).Rows(model)
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

		_, err = gr.DB.ID(m.Id).Get(&m.Vistoria)
		if err != nil {
			log.Error().Msgf("Error searching vistoria: %s", err.Error())
		}
	}
}

func AfterSave(gr *models.GenericRepository[domain.Historico, domain.HistoricoFilter, Historico], session *xorm.Session, m *Historico) bool {
	if m.Id != 0 {
		m.NumeroOrdem = m.Id
		_, err := session.Exec("UPDATE historico SET numero = ? WHERE id = ?", m.Id, m.Id)
		if err != nil {
			log.Error().Msg(err.Error())
			session.Rollback()
			return false
		}
		_, err = session.Exec("DELETE FROM historico_item WHERE id_historico = ?", m.Id)
		if err != nil {
			log.Error().Msg(err.Error())
			session.Rollback()
			return false
		}
		for _, item := range m.Items {
			item.IdHistorico = m.Id
			_, err := session.Insert(&item)
			if err != nil {
				log.Error().Msg(err.Error())
				session.Rollback()
				return false
			}
		}

		_, err = session.Exec("DELETE FROM vistoria WHERE id_historico = ?", m.Id)
		if err != nil {
			log.Error().Msg(err.Error())
			session.Rollback()
			return false
		}
		m.Vistoria.Id = m.Id
		_, err = session.Insert(&m.Vistoria)
		if err != nil {
			log.Error().Msg(err.Error())
			session.Rollback()
			return false
		}
	}
	return true
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
		where = where.And("numero = ?", filter.NumeroOrdem)
	}
	if hasIdCliente {
		where = where.And("id_cliente = ?", filter.IdCliente)
	}
	if hasIdVeiculo {
		where = where.And("id_veiculo = ?", filter.IdVeiculo)
	}
	if hasIdTecnico {
		where = where.And("id_tecnico = ?", filter.IdTecnico)
	}
	return where
}

func MapperToEntity(dto *domain.Historico) *Historico {
	entity := new(Historico)
	copyToEntity(dto, entity)
	return entity
}

func MapperToDTO(entity *Historico) *domain.Historico {
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
	destiny.Vistoria.Kilometragem = source.Kilometragem
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
	destiny.Kilometragem = source.Vistoria.Kilometragem
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
