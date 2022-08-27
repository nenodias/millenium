package lembrete

import (
	"strings"
	"time"

	domain "github.com/nenodias/millenium/core/domain/lembrete"
	models "github.com/nenodias/millenium/repositories/models"
	"xorm.io/xorm"
)

type Lembrete struct {
	Id        int64     `xorm:"'id' bigint pk autoincr not null"`
	Texto     string    `xorm:"'texto' varchar(5000)"`
	IdCliente int64     `xorm:"'id_cliente' bigint"`
	IdVeiculo int64     `xorm:"'id_veiculo' bigint"`
	Data      time.Time `xorm:"'data_notificacao' timestamp"`
}

func (p *Lembrete) TableName() string {
	return "lembretes"
}

type LembreteRepository struct {
	models.GenericRepository[domain.Lembrete, domain.LembreteFilter, Lembrete]
}

func NewService(engine *xorm.Engine) domain.LembreteService {
	repository := LembreteRepository{
		GenericRepository: models.GenericRepository[domain.Lembrete, domain.LembreteFilter, Lembrete]{
			DB:             engine,
			MapperToDTO:    mapperToDTO,
			MapperToEntity: mapperToEntity,
			CopyToDto:      copyToDto,
			HasWhere:       hasWhere,
			DoWhere:        doWhere,
		},
	}
	return domain.LembreteService(&repository)
}

func hasWhere(filter *domain.LembreteFilter) bool {
	hasTexto := filter.Texto != "" && strings.TrimSpace(filter.Texto) != ""
	hasIdCliente := filter.IdCliente != int64(0)
	hasIdVeiculo := filter.IdVeiculo != int64(0)
	return hasTexto || hasIdCliente || hasIdVeiculo
}

func doWhere(query *xorm.Session, filter *domain.LembreteFilter) *xorm.Session {
	hasTexto := filter.Texto != "" && strings.TrimSpace(filter.Texto) != ""
	hasIdCliente := filter.IdCliente != int64(0)
	hasIdVeiculo := filter.IdVeiculo != int64(0)
	where := make([]interface{}, 0)
	if hasTexto {
		where = append(where, "texto Like ?", "%"+filter.Texto+"%")
	}
	if hasIdCliente {
		where = append(where, "id_cliente = ?", filter.IdCliente)
	}
	if hasIdVeiculo {
		where = append(where, "id_veiculo = ?", filter.IdVeiculo)
	}
	if len(where) == 2 {
		return query.Where(where[0], where[1])
	} else if len(where) == 4 {
		return query.Where(where[0], where[1]).And(where[2], where[3])
	} else {
		return query.Where(where[0], where[1]).And(where[2], where[3]).And(where[4], where[5])
	}
}

func mapperToEntity(dto *domain.Lembrete) *Lembrete {
	entity := new(Lembrete)
	copyToEntity(dto, entity)
	return entity
}

func mapperToDTO(entity *Lembrete) *domain.Lembrete {
	dto := new(domain.Lembrete)
	copyToDto(entity, dto)
	return dto
}

func copyToEntity(source *domain.Lembrete, destiny *Lembrete) {
	destiny.Id = source.Id
	destiny.Texto = source.Texto
	destiny.IdCliente = source.IdCliente
	destiny.IdVeiculo = source.IdVeiculo
	destiny.Data = source.Data
}

func copyToDto(source *Lembrete, destiny *domain.Lembrete) {
	destiny.Id = source.Id
	destiny.Texto = source.Texto
	destiny.IdCliente = source.IdCliente
	destiny.IdVeiculo = source.IdVeiculo
	destiny.Data = source.Data
}
