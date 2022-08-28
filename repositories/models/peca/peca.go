package peca

import (
	"strings"

	domain "github.com/nenodias/millenium/core/domain/peca"
	models "github.com/nenodias/millenium/repositories/models"
	"xorm.io/xorm"
)

type Peca struct {
	Id        int64   `xorm:"'id' bigint pk autoincr not null"`
	Valor     float64 `xorm:"'valor' double"`
	Descricao string  `xorm:"'descricao' varchar(60) not null"`
}

func (p *Peca) TableName() string {
	return "pecas"
}

type PecaRepository struct {
	models.GenericRepository[domain.Peca, domain.PecaFilter, Peca]
}

func NewService(engine *xorm.Engine) domain.PecaService {
	repository := PecaRepository{
		GenericRepository: models.GenericRepository[domain.Peca, domain.PecaFilter, Peca]{
			DB:             engine,
			MapperToDTO:    MapperToDTO,
			MapperToEntity: MapperToEntity,
			CopyToDto:      copyToDto,
			HasWhere:       hasWhere,
			DoWhere:        doWhere,
		},
	}
	return domain.PecaService(&repository)
}

func hasWhere(filter *domain.PecaFilter) bool {
	return filter.Descricao != "" && strings.TrimSpace(filter.Descricao) != ""
}

func doWhere(query *xorm.Session, filter *domain.PecaFilter) *xorm.Session {
	where := []string{"descricao Like ?", "%" + filter.Descricao + "%"}
	return query.Where(where[0], where[1])
}

func MapperToEntity(dto *domain.Peca) *Peca {
	entity := new(Peca)
	copyToEntity(dto, entity)
	return entity
}

func MapperToDTO(entity *Peca) *domain.Peca {
	dto := new(domain.Peca)
	copyToDto(entity, dto)
	return dto
}

func copyToEntity(source *domain.Peca, destiny *Peca) {
	destiny.Id = source.Id
	destiny.Descricao = source.Descricao
	destiny.Valor = source.Valor
}

func copyToDto(source *Peca, destiny *domain.Peca) {
	destiny.Id = source.Id
	destiny.Descricao = source.Descricao
	destiny.Valor = source.Valor
}
