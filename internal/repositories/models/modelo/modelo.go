package modelo

import (
	"context"

	domain "github.com/nenodias/millenium/internal/core/domain/modelo"
	"github.com/nenodias/millenium/internal/core/domain/utils"
	"github.com/nenodias/millenium/internal/repositories"
	models "github.com/nenodias/millenium/internal/repositories/models"
	"xorm.io/xorm"
)

type Modelo struct {
	Id              int64  `xorm:"'id' bigint pk autoincr not null"`
	Nome            string `xorm:"'nome' varchar(40) not null"`
	IdMontadora     int64  `xorm:"'id_montadora' bigint"`
	CodigoVeiculoEA int    `xorm:"'codvei_ea' int"`
}

func (p *Modelo) TableName() string {
	return "modelo"
}

type ModeloRepository struct {
	models.GenericRepository[domain.Modelo, domain.ModeloFilter, Modelo]
}

func NewService(engine *repositories.DatabaseEngine) domain.ModeloService {
	repository := ModeloRepository{
		GenericRepository: models.GenericRepository[domain.Modelo, domain.ModeloFilter, Modelo]{
			DB:             engine,
			MapperToDTO:    MapperToDTO,
			MapperToEntity: MapperToEntity,
			CopyToDto:      copyToDto,
			HasWhere:       hasWhere,
			DoWhere:        doWhere,
		},
	}
	return domain.ModeloService(&repository)
}

func hasWhere(ctx context.Context, filter *domain.ModeloFilter) bool {
	hasNome := utils.HasValue(filter.Nome)
	hasIdModelo := utils.HasValueInt64(filter.IdModelo)
	return hasNome || hasIdModelo
}

func doWhere(ctx context.Context, query *xorm.Session, filter *domain.ModeloFilter) *xorm.Session {
	hasNome := utils.HasValue(filter.Nome)
	hasIdModelo := utils.HasValueInt64(filter.IdModelo)
	if hasNome && hasIdModelo {
		return query.Where("nome ILIKE ?", "%"+filter.Nome+"%").And("id_montadora = ?", filter.IdModelo)
	} else if hasNome {
		return query.Where("nome ILIKE ?", "%"+filter.Nome+"%")
	} else {
		return query.Where("id_montadora = ?", filter.IdModelo)
	}
}

func MapperToEntity(ctx context.Context, dto *domain.Modelo) *Modelo {
	entity := new(Modelo)
	copyToEntity(ctx, dto, entity)
	return entity
}

func MapperToDTO(ctx context.Context, entity *Modelo) *domain.Modelo {
	dto := new(domain.Modelo)
	copyToDto(ctx, entity, dto)
	return dto
}

func copyToEntity(ctx context.Context, source *domain.Modelo, destiny *Modelo) {
	destiny.Id = source.Id
	destiny.Nome = source.Nome
	destiny.IdMontadora = source.IdMontadora
	destiny.CodigoVeiculoEA = source.CodigoVeiculoEA
}

func copyToDto(ctx context.Context, source *Modelo, destiny *domain.Modelo) {
	destiny.Id = source.Id
	destiny.Nome = source.Nome
	destiny.IdMontadora = source.IdMontadora
	destiny.CodigoVeiculoEA = source.CodigoVeiculoEA
}
