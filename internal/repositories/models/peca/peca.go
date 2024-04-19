package peca

import (
	"context"

	domain "github.com/nenodias/millenium/internal/core/domain/peca"
	"github.com/nenodias/millenium/internal/core/domain/utils"
	"github.com/nenodias/millenium/internal/repositories"
	models "github.com/nenodias/millenium/internal/repositories/models"
	"xorm.io/xorm"
)

type Peca struct {
	Id        int64   `xorm:"'id' bigint pk autoincr not null"`
	Valor     float64 `xorm:"'valor' double"`
	Descricao string  `xorm:"'descricao' varchar(60) not null"`
}

func (p *Peca) TableName() string {
	return "peca"
}

type PecaRepository struct {
	models.GenericRepository[domain.Peca, domain.PecaFilter, Peca]
}

func NewService(engine *repositories.DatabaseEngine) domain.PecaService {
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

func hasWhere(ctx context.Context, filter *domain.PecaFilter) bool {
	return utils.HasValue(filter.Descricao)
}

func doWhere(ctx context.Context, query *xorm.Session, filter *domain.PecaFilter) *xorm.Session {
	where := []string{"descricao ILIKE ?", "%" + filter.Descricao + "%"}
	return query.Where(where[0], where[1])
}

func MapperToEntity(ctx context.Context, dto *domain.Peca) *Peca {
	entity := new(Peca)
	copyToEntity(ctx, dto, entity)
	return entity
}

func MapperToDTO(ctx context.Context, entity *Peca) *domain.Peca {
	dto := new(domain.Peca)
	copyToDto(ctx, entity, dto)
	return dto
}

func copyToEntity(ctx context.Context, source *domain.Peca, destiny *Peca) {
	destiny.Id = source.Id
	destiny.Descricao = source.Descricao
	destiny.Valor = source.Valor
}

func copyToDto(ctx context.Context, source *Peca, destiny *domain.Peca) {
	destiny.Id = source.Id
	destiny.Descricao = source.Descricao
	destiny.Valor = source.Valor
}
