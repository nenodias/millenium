package modelo

import (
	"strings"

	domain "github.com/nenodias/millenium/core/domain/modelo"
	models "github.com/nenodias/millenium/repositories/models"
	"xorm.io/xorm"
)

type Modelo struct {
	Id              int64  `xorm:"'id' bigint pk autoincr not null"`
	Nome            string `xorm:"'nome_modelo' varchar(40) not null"`
	IdMontadora     int64  `xorm:"'id_monta' bigint"`
	CodigoVeiculoEA int    `xorm:"'codvei_ea' int"`
}

func (p *Modelo) TableName() string {
	return "modelo"
}

type ModeloRepository struct {
	models.GenericRepository[domain.Modelo, domain.ModeloFilter, Modelo]
}

func NewService(engine *xorm.Engine) domain.ModeloService {
	repository := ModeloRepository{
		GenericRepository: models.GenericRepository[domain.Modelo, domain.ModeloFilter, Modelo]{
			DB:             engine,
			MapperToDTO:    mapperToDTO,
			MapperToEntity: mapperToEntity,
			CopyToDto:      copyToDto,
			HasWhere:       hasWhere,
			DoWhere:        doWhere,
		},
	}
	return domain.ModeloService(&repository)
}

func hasWhere(filter *domain.ModeloFilter) bool {
	hasNome := filter.Nome != "" && strings.TrimSpace(filter.Nome) != ""
	hasIdModelo := filter.IdModelo != int64(0)
	return hasNome || hasIdModelo
}

func doWhere(query *xorm.Session, filter *domain.ModeloFilter) *xorm.Session {
	hasNome := filter.Nome != "" && strings.TrimSpace(filter.Nome) != ""
	hasIdModelo := filter.IdModelo != int64(0)
	if hasNome && hasIdModelo {
		return query.Where("nome_modelo Like ?", "%"+filter.Nome+"%").And("id_monta = ?", filter.IdModelo)
	} else if hasNome {
		return query.Where("nome_modelo Like ?", "%"+filter.Nome+"%")
	} else {
		return query.Where("id_monta = ?", filter.IdModelo)
	}
}

func mapperToEntity(dto *domain.Modelo) *Modelo {
	entity := new(Modelo)
	copyToEntity(dto, entity)
	return entity
}

func mapperToDTO(entity *Modelo) *domain.Modelo {
	dto := new(domain.Modelo)
	copyToDto(entity, dto)
	return dto
}

func copyToEntity(source *domain.Modelo, destiny *Modelo) {
	destiny.Id = source.Id
	destiny.Nome = source.Nome
	destiny.IdMontadora = source.IdMontadora
	destiny.CodigoVeiculoEA = source.CodigoVeiculoEA
}

func copyToDto(source *Modelo, destiny *domain.Modelo) {
	destiny.Id = source.Id
	destiny.Nome = source.Nome
	destiny.IdMontadora = source.IdMontadora
	destiny.CodigoVeiculoEA = source.CodigoVeiculoEA
}
