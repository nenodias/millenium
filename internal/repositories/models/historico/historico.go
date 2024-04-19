package historico

import (
	"context"
	"fmt"
	"time"

	domain "github.com/nenodias/millenium/internal/core/domain/historico"
	"github.com/nenodias/millenium/internal/core/domain/utils"
	"github.com/nenodias/millenium/internal/repositories"
	models "github.com/nenodias/millenium/internal/repositories/models"
	clienteModel "github.com/nenodias/millenium/internal/repositories/models/cliente"
	modeloModel "github.com/nenodias/millenium/internal/repositories/models/modelo"
	montadoraModel "github.com/nenodias/millenium/internal/repositories/models/montadora"
	tecnicoModel "github.com/nenodias/millenium/internal/repositories/models/tecnico"
	veiculoModel "github.com/nenodias/millenium/internal/repositories/models/veiculo"
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
	Descricao   string                   `xorm:"'descricao' varchar(75)"`
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

func NewService(engine *repositories.DatabaseEngine) domain.HistoricoService {
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

func (hr *HistoricoRepository) FindOneForReport(ctx context.Context, id int64) (*domain.HistoricoReport, error) {
	report := new(domain.HistoricoReport)
	model, err := hr.FindOne(ctx, id)
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
		if utils.HasValueInt64(model.IdCliente) {
			_, err = hr.DB.ID(model.IdCliente).Get(cliente)
			if err != nil {
				log.Error().Msg(err.Error())
				return nil, err
			}
			report.Cliente = *clienteModel.MapperToDTO(ctx, cliente)
		}
		if utils.HasValueInt64(model.IdVeiculo) {
			_, err = hr.DB.ID(model.IdVeiculo).Get(veiculo)
			if err != nil {
				log.Error().Msg(err.Error())
				return nil, err
			}
			report.Veiculo = *veiculoModel.MapperToDTO(ctx, veiculo)
			if utils.HasValueInt64(veiculo.IdModelo) {
				_, err = hr.DB.ID(veiculo.IdModelo).Get(modelo)
				if err != nil {
					log.Error().Msg(err.Error())
				}
				report.Modelo = *modeloModel.MapperToDTO(ctx, modelo)
				if utils.HasValueInt64(modelo.IdMontadora) {
					_, err = hr.DB.ID(modelo.IdMontadora).Get(montadora)
					if err != nil {
						log.Error().Msg(err.Error())
					}
					report.Montadora = *montadoraModel.MapperToDTO(ctx, montadora)
				}
			}
		}

		if model.IdTecnico != nil {
			_, err = hr.DB.ID(*model.IdTecnico).Get(tecnico)
			if err != nil {
				log.Error().Msg(err.Error())
			}
			report.Tecnico = *tecnicoModel.MapperToDTO(ctx, tecnico)
		}
	} else {
		return report, fmt.Errorf("registro com id: %d nao encontrado", id)
	}
	return report, nil
}

func AfterFind(ctx context.Context, gr *models.GenericRepository[domain.Historico, domain.HistoricoFilter, Historico], m *Historico) {
	if utils.HasValueInt64(m.Id) {
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

func AfterSave(ctx context.Context, gr *models.GenericRepository[domain.Historico, domain.HistoricoFilter, Historico], session *xorm.Session, m *Historico) bool {
	if utils.HasValueInt64(m.Id) {
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

func hasWhere(ctx context.Context, filter *domain.HistoricoFilter) bool {
	hasData := !filter.Data.IsZero()
	hasTipo := filter.Tipo != nil
	hasNumeroOrdem := utils.HasValueInt64(filter.IdCliente)
	hasIdCliente := utils.HasValueInt64(filter.IdCliente)
	hasIdVeiculo := utils.HasValueInt64(filter.IdVeiculo)
	hasIdTecnico := utils.HasValueInt64(filter.IdTecnico)
	return hasData || hasTipo || hasNumeroOrdem || hasIdCliente || hasIdVeiculo || hasIdTecnico
}

func doWhere(ctx context.Context, query *xorm.Session, filter *domain.HistoricoFilter) *xorm.Session {
	hasData := !filter.Data.IsZero()
	hasTipo := filter.Tipo != nil
	hasNumeroOrdem := utils.HasValueInt64(filter.IdCliente)
	hasIdCliente := utils.HasValueInt64(filter.IdCliente)
	hasIdVeiculo := utils.HasValueInt64(filter.IdVeiculo)
	hasIdTecnico := utils.HasValueInt64(filter.IdTecnico)
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

func MapperToEntity(ctx context.Context, dto *domain.Historico) *Historico {
	entity := new(Historico)
	copyToEntity(ctx, dto, entity)
	return entity
}

func MapperToDTO(ctx context.Context, entity *Historico) *domain.Historico {
	dto := new(domain.Historico)
	copyToDto(ctx, entity, dto)
	return dto
}

func mapperToEntityItem(ctx context.Context, dtos []domain.HistoricoItem) []HistoricoItem {
	items := make([]HistoricoItem, 0)
	for _, dto := range dtos {
		item := new(HistoricoItem)
		copyToEntityItem(ctx, &dto, item)
		items = append(items, *item)
	}
	return items
}

func mapperToDTOItem(ctx context.Context, dtos []HistoricoItem) []domain.HistoricoItem {
	items := make([]domain.HistoricoItem, 0)
	for _, dto := range dtos {
		item := new(domain.HistoricoItem)
		copyToDtoItem(ctx, &dto, item)
		items = append(items, *item)
	}
	return items
}

func copyToEntity(ctx context.Context, source *domain.Historico, destiny *Historico) {
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
	destiny.Items = mapperToEntityItem(ctx, source.Items)
	destiny.Vistoria.Kilometragem = source.Kilometragem
}

func copyToDto(ctx context.Context, source *Historico, destiny *domain.Historico) {
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
	destiny.Items = mapperToDTOItem(ctx, source.Items)
	destiny.Kilometragem = source.Vistoria.Kilometragem
}

func copyToEntityItem(ctx context.Context, source *domain.HistoricoItem, destiny *HistoricoItem) {
	destiny.Id = source.Id
	destiny.IdHistorico = source.IdHistorico
	destiny.Ordem = source.Ordem
	destiny.Tipo = source.Tipo
	destiny.Descricao = source.Descricao
	destiny.Quantidade = source.Quantidade
	destiny.Valor = source.Valor
}

func copyToDtoItem(ctx context.Context, source *HistoricoItem, destiny *domain.HistoricoItem) {
	destiny.Id = source.Id
	destiny.IdHistorico = source.IdHistorico
	destiny.Ordem = source.Ordem
	destiny.Tipo = source.Tipo
	destiny.Descricao = source.Descricao
	destiny.Quantidade = source.Quantidade
	destiny.Valor = source.Valor
}
