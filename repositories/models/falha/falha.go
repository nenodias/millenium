package falha

import (
	"context"
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
	return "falha"
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

func hasWhere(ctx context.Context, filter *domain.FalhaFilter) bool {
	return filter.Descricao != "" && strings.TrimSpace(filter.Descricao) != ""
}

func doWhere(ctx context.Context, query *xorm.Session, filter *domain.FalhaFilter) *xorm.Session {
	where := []string{"descricao ILIKE ?", "%" + filter.Descricao + "%"}
	return query.Where(where[0], where[1])
}

func MapperToEntity(ctx context.Context, dto *domain.Falha) *Falha {
	entity := new(Falha)
	copyToEntity(ctx, dto, entity)
	return entity
}

func MapperToDTO(ctx context.Context, entity *Falha) *domain.Falha {
	dto := new(domain.Falha)
	copyToDto(ctx, entity, dto)
	return dto
}

func copyToEntity(ctx context.Context, source *domain.Falha, destiny *Falha) {
	destiny.Id = source.Id
	destiny.Descricao = source.Descricao
}

func copyToDto(ctx context.Context, source *Falha, destiny *domain.Falha) {
	destiny.Id = source.Id
	destiny.Descricao = source.Descricao
}
