package tecnico

import (
	"context"
	"strings"

	domain "github.com/nenodias/millenium/core/domain/tecnico"
	models "github.com/nenodias/millenium/repositories/models"
	"xorm.io/xorm"
)

type Tecnico struct {
	Id   int64  `xorm:"'id' bigint pk autoincr not null"`
	Nome string `xorm:"'nome' varchar(60) not null"`
}

func (p *Tecnico) TableName() string {
	return "tecnico"
}

type TecnicoRepository struct {
	models.GenericRepository[domain.Tecnico, domain.TecnicoFilter, Tecnico]
}

func NewService(engine *xorm.Engine) domain.TecnicoService {
	repository := TecnicoRepository{
		GenericRepository: models.GenericRepository[domain.Tecnico, domain.TecnicoFilter, Tecnico]{
			DB:             engine,
			MapperToDTO:    MapperToDTO,
			MapperToEntity: MapperToEntity,
			CopyToDto:      copyToDto,
			HasWhere:       hasWhere,
			DoWhere:        doWhere,
		},
	}
	return domain.TecnicoService(&repository)
}

func hasWhere(ctx context.Context, filter *domain.TecnicoFilter) bool {
	return filter.Nome != "" && strings.TrimSpace(filter.Nome) != ""
}

func doWhere(ctx context.Context, query *xorm.Session, filter *domain.TecnicoFilter) *xorm.Session {
	where := []string{"nome ILIKE ?", "%" + filter.Nome + "%"}
	return query.Where(where[0], where[1])
}

func MapperToEntity(ctx context.Context, dto *domain.Tecnico) *Tecnico {
	entity := new(Tecnico)
	copyToEntity(ctx, dto, entity)
	return entity
}

func MapperToDTO(ctx context.Context, entity *Tecnico) *domain.Tecnico {
	dto := new(domain.Tecnico)
	copyToDto(ctx, entity, dto)
	return dto
}

func copyToEntity(ctx context.Context, source *domain.Tecnico, destiny *Tecnico) {
	destiny.Id = source.Id
	destiny.Nome = source.Nome
}

func copyToDto(ctx context.Context, source *Tecnico, destiny *domain.Tecnico) {
	destiny.Id = source.Id
	destiny.Nome = source.Nome
}
