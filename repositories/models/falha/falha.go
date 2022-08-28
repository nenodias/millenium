package falha

import (
	"strings"

	domain "github.com/nenodias/millenium/core/domain/falha"
	models "github.com/nenodias/millenium/repositories/models"
	"xorm.io/xorm"
)

type Falha struct {
	Id        int64  `xorm:"'id' bigint pk autoincr not null"`
	Descricao string `xorm:"'descricao' varchar(60) not null"`
}

func (p *Falha) TableName() string {
	return "falhas"
}

type FalhaRepository struct {
	models.GenericRepository[domain.Falha, domain.FalhaFilter, Falha]
}

func NewService(engine *xorm.Engine) domain.FalhaService {
	repository := FalhaRepository{
		GenericRepository: models.GenericRepository[domain.Falha, domain.FalhaFilter, Falha]{
			DB:             engine,
			MapperToDTO:    MapperToDTO,
			MapperToEntity: MapperToEntity,
			CopyToDto:      copyToDto,
			HasWhere:       hasWhere,
			DoWhere:        doWhere,
		},
	}
	return domain.FalhaService(&repository)
}

func hasWhere(filter *domain.FalhaFilter) bool {
	return filter.Descricao != "" && strings.TrimSpace(filter.Descricao) != ""
}

func doWhere(query *xorm.Session, filter *domain.FalhaFilter) *xorm.Session {
	where := []string{"descricao Like ?", "%" + filter.Descricao + "%"}
	return query.Where(where[0], where[1])
}

func MapperToEntity(dto *domain.Falha) *Falha {
	entity := new(Falha)
	copyToEntity(dto, entity)
	return entity
}

func MapperToDTO(entity *Falha) *domain.Falha {
	dto := new(domain.Falha)
	copyToDto(entity, dto)
	return dto
}

func copyToEntity(source *domain.Falha, destiny *Falha) {
	destiny.Id = source.Id
	destiny.Descricao = source.Descricao
}

func copyToDto(source *Falha, destiny *domain.Falha) {
	destiny.Id = source.Id
	destiny.Descricao = source.Descricao
}
