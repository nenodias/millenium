package models

import (
	"strings"

	domain "github.com/nenodias/millenium/core/domain/tecnico"
	"xorm.io/xorm"
)

type Tecnico struct {
	Id   int64  `xorm:"'codigo_tecnico' bigint pk autoincr not null"`
	Nome string `xorm:"'nome' varchar(60) not null"`
}

func (p *Tecnico) TableName() string {
	return "tecnico"
}

type TecnicoRepository struct {
	GenericRepository[domain.Tecnico, domain.TecnicoFilter, Tecnico]
}

func NewTecnicoService(engine *xorm.Engine) domain.TecnicoService {
	repository := TecnicoRepository{
		GenericRepository: GenericRepository[domain.Tecnico, domain.TecnicoFilter, Tecnico]{
			DB:             engine,
			MapperToDTO:    mapperToDTO,
			MapperToEntity: mapperToEntity,
			CopyToDto:      copyToDto,
			HasWhere:       hasWhere,
			DoWhere:        doWhere,
		},
	}
	return domain.TecnicoService(&repository)
}

func hasWhere(filter *domain.TecnicoFilter) bool {
	return filter.Nome != "" && strings.TrimSpace(filter.Nome) != ""
}

func doWhere(query *xorm.Session, filter *domain.TecnicoFilter) *xorm.Session {
	where := []string{"nome Like ?", "%" + filter.Nome + "%"}
	return query.Where(where[0], where[1])
}

func mapperToEntity(dto *domain.Tecnico) *Tecnico {
	entity := new(Tecnico)
	copyToEntity(dto, entity)
	return entity
}

func mapperToDTO(entity *Tecnico) *domain.Tecnico {
	dto := new(domain.Tecnico)
	copyToDto(entity, dto)
	return dto
}

func copyToEntity(source *domain.Tecnico, destiny *Tecnico) {
	destiny.Id = source.Id
	destiny.Nome = source.Nome
}

func copyToDto(source *Tecnico, destiny *domain.Tecnico) {
	destiny.Id = source.Id
	destiny.Nome = source.Nome
}
