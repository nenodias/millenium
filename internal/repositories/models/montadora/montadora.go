package montadora

import (
	"context"

	domain "github.com/nenodias/millenium/internal/core/domain/montadora"
	"github.com/nenodias/millenium/internal/core/domain/utils"
	"github.com/nenodias/millenium/internal/repositories"
	models "github.com/nenodias/millenium/internal/repositories/models"
	"xorm.io/xorm"
)

type Montadora struct {
	Id                int64  `xorm:"'id' bigint pk autoincr not null"`
	Nome              string `xorm:"'nome' varchar(20) not null"`
	Origem            string `xorm:"'origem' varchar(1) not null"`
	CodigoMontadoraEA int    `xorm:"'codmon_ea' int"`
}

func (p *Montadora) TableName() string {
	return "montadora"
}

type MontadoraRepository struct {
	models.GenericRepository[domain.Montadora, domain.MontadoraFilter, Montadora]
}

func NewService(engine *repositories.DatabaseEngine) domain.MontadoraService {
	repository := MontadoraRepository{
		GenericRepository: models.GenericRepository[domain.Montadora, domain.MontadoraFilter, Montadora]{
			DB:             engine,
			MapperToDTO:    MapperToDTO,
			MapperToEntity: MapperToEntity,
			CopyToDto:      copyToDto,
			HasWhere:       hasWhere,
			DoWhere:        doWhere,
		},
	}
	return domain.MontadoraService(&repository)
}

func hasWhere(ctx context.Context, filter *domain.MontadoraFilter) bool {
	return utils.HasValue(filter.Nome)
}

func doWhere(ctx context.Context, query *xorm.Session, filter *domain.MontadoraFilter) *xorm.Session {
	where := []string{"nome ILIKE ?", "%" + filter.Nome + "%"}
	return query.Where(where[0], where[1])
}

func MapperToEntity(ctx context.Context, dto *domain.Montadora) *Montadora {
	entity := new(Montadora)
	copyToEntity(ctx, dto, entity)
	return entity
}

func MapperToDTO(ctx context.Context, entity *Montadora) *domain.Montadora {
	dto := new(domain.Montadora)
	copyToDto(ctx, entity, dto)
	return dto
}

func copyToEntity(ctx context.Context, source *domain.Montadora, destiny *Montadora) {
	destiny.Id = source.Id
	destiny.Nome = source.Nome
	destiny.Origem = source.Origem
	destiny.CodigoMontadoraEA = source.CodigoMontadoraEA
}

func copyToDto(ctx context.Context, source *Montadora, destiny *domain.Montadora) {
	destiny.Id = source.Id
	destiny.Nome = source.Nome
	destiny.Origem = source.Origem
	destiny.CodigoMontadoraEA = source.CodigoMontadoraEA
}
