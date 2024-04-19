package servico

import (
	"context"

	domain "github.com/nenodias/millenium/internal/core/domain/servico"
	"github.com/nenodias/millenium/internal/core/domain/utils"
	"github.com/nenodias/millenium/internal/repositories"
	models "github.com/nenodias/millenium/internal/repositories/models"
	"xorm.io/xorm"
)

type Servico struct {
	Id        int64   `xorm:"'id' bigint pk autoincr not null"`
	Valor     float64 `xorm:"'valor' double"`
	Descricao string  `xorm:"'descricao' varchar(60) not null"`
}

func (p *Servico) TableName() string {
	return "servico"
}

type ServicoRepository struct {
	models.GenericRepository[domain.Servico, domain.ServicoFilter, Servico]
}

func NewService(engine *repositories.DatabaseEngine) domain.ServicoService {
	repository := ServicoRepository{
		GenericRepository: models.GenericRepository[domain.Servico, domain.ServicoFilter, Servico]{
			DB:             engine,
			MapperToDTO:    MapperToDTO,
			MapperToEntity: MapperToEntity,
			CopyToDto:      copyToDto,
			HasWhere:       hasWhere,
			DoWhere:        doWhere,
		},
	}
	return domain.ServicoService(&repository)
}

func hasWhere(ctx context.Context, filter *domain.ServicoFilter) bool {
	return utils.HasValue(filter.Descricao)
}

func doWhere(ctx context.Context, query *xorm.Session, filter *domain.ServicoFilter) *xorm.Session {
	where := []string{"descricao ILIKE ?", "%" + filter.Descricao + "%"}
	return query.Where(where[0], where[1])
}

func MapperToEntity(ctx context.Context, dto *domain.Servico) *Servico {
	entity := new(Servico)
	copyToEntity(ctx, dto, entity)
	return entity
}

func MapperToDTO(ctx context.Context, entity *Servico) *domain.Servico {
	dto := new(domain.Servico)
	copyToDto(ctx, entity, dto)
	return dto
}

func copyToEntity(ctx context.Context, source *domain.Servico, destiny *Servico) {
	destiny.Id = source.Id
	destiny.Descricao = source.Descricao
}

func copyToDto(ctx context.Context, source *Servico, destiny *domain.Servico) {
	destiny.Id = source.Id
	destiny.Descricao = source.Descricao
}
