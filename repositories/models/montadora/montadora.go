package montadora

import (
	"strings"

	domain "github.com/nenodias/millenium/core/domain/montadora"
	models "github.com/nenodias/millenium/repositories/models"
	"xorm.io/xorm"
)

type Montadora struct {
	Id                int64  `xorm:"'id' bigint pk autoincr not null"`
	Nome              string `xorm:"'nome_montadora' varchar(20) not null"`
	Origem            string `xorm:"'origem' varchar(1) not null"`
	CodigoMontadoraEA int    `xorm:"'codmon_ea' int"`
}

func (p *Montadora) TableName() string {
	return "montadora"
}

type MontadoraRepository struct {
	models.GenericRepository[domain.Montadora, domain.MontadoraFilter, Montadora]
}

func NewMontadoraService(engine *xorm.Engine) domain.MontadoraService {
	repository := MontadoraRepository{
		GenericRepository: models.GenericRepository[domain.Montadora, domain.MontadoraFilter, Montadora]{
			DB:             engine,
			MapperToDTO:    mapperToDTO,
			MapperToEntity: mapperToEntity,
			CopyToDto:      copyToDto,
			HasWhere:       hasWhere,
			DoWhere:        doWhere,
		},
	}
	return domain.MontadoraService(&repository)
}

func hasWhere(filter *domain.MontadoraFilter) bool {
	return filter.Nome != "" && strings.TrimSpace(filter.Nome) != ""
}

func doWhere(query *xorm.Session, filter *domain.MontadoraFilter) *xorm.Session {
	where := []string{"nome_montadora Like ?", "%" + filter.Nome + "%"}
	return query.Where(where[0], where[1])
}

func mapperToEntity(dto *domain.Montadora) *Montadora {
	entity := new(Montadora)
	copyToEntity(dto, entity)
	return entity
}

func mapperToDTO(entity *Montadora) *domain.Montadora {
	dto := new(domain.Montadora)
	copyToDto(entity, dto)
	return dto
}

func copyToEntity(source *domain.Montadora, destiny *Montadora) {
	destiny.Id = source.Id
	destiny.Nome = source.Nome
	destiny.Origem = source.Origem
	destiny.CodigoMontadoraEA = source.CodigoMontadoraEA
}

func copyToDto(source *Montadora, destiny *domain.Montadora) {
	destiny.Id = source.Id
	destiny.Nome = source.Nome
	destiny.Origem = source.Origem
	destiny.CodigoMontadoraEA = source.CodigoMontadoraEA
}
