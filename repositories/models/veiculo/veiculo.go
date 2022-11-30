package veiculo

import (
	"strings"

	domain "github.com/nenodias/millenium/core/domain/veiculo"
	models "github.com/nenodias/millenium/repositories/models"
	"xorm.io/xorm"
)

type Veiculo struct {
	Id          int64              `xorm:"'id' bigint pk autoincr not null"`
	IdCliente   int64              `xorm:"'id_cliente' bigint"`
	Placa       string             `xorm:"'placa' varchar(8) not null"`
	Pais        string             `xorm:"'pais' varchar(20)"`
	Cor         string             `xorm:"'cor' varchar(20)"`
	Combustivel domain.Combustivel `xorm:"'combustivel' varchar(10)"`
	Renavam     string             `xorm:"'renavam' varchar(40)"`
	Chassi      string             `xorm:"'chassi' varchar(40)"`
	Ano         string             `xorm:"'ano' varchar(4)"`
	IdModelo    int64              `xorm:"'id_modelo' bigint"`
}

func (p *Veiculo) TableName() string {
	return "veiculo"
}

type VeiculoRepository struct {
	models.GenericRepository[domain.Veiculo, domain.VeiculoFilter, Veiculo]
}

func NewService(engine *xorm.Engine) domain.VeiculoService {
	repository := VeiculoRepository{
		GenericRepository: models.GenericRepository[domain.Veiculo, domain.VeiculoFilter, Veiculo]{
			DB:             engine,
			MapperToDTO:    MapperToDTO,
			MapperToEntity: MapperToEntity,
			CopyToDto:      copyToDto,
			HasWhere:       hasWhere,
			DoWhere:        doWhere,
		},
	}
	return domain.VeiculoService(&repository)
}

func hasWhere(filter *domain.VeiculoFilter) bool {
	hasPlaca := filter.Placa != "" && strings.TrimSpace(filter.Placa) != ""
	hasIdCliente := filter.IdCliente != int64(0)
	hasIdModelo := filter.IdModelo != int64(0)
	return hasPlaca || hasIdCliente || hasIdModelo
}

func doWhere(query *xorm.Session, filter *domain.VeiculoFilter) *xorm.Session {
	hasPlaca := filter.Placa != "" && strings.TrimSpace(filter.Placa) != ""
	hasIdCliente := filter.IdCliente != int64(0)
	hasIdModelo := filter.IdModelo != int64(0)
	where := make([]interface{}, 0)
	if hasPlaca {
		where = append(where, "placa ILIKE ?", "%"+filter.Placa+"%")
	}
	if hasIdCliente {
		where = append(where, "id_cliente = ?", filter.IdCliente)
	}
	if hasIdModelo {
		where = append(where, "id_modelo = ?", filter.IdModelo)
	}
	if len(where) == 2 {
		return query.Where(where[0], where[1])
	} else if len(where) == 4 {
		return query.Where(where[0], where[1]).And(where[2], where[3])
	} else {
		return query.Where(where[0], where[1]).And(where[2], where[3]).And(where[4], where[5])
	}
}

func MapperToEntity(dto *domain.Veiculo) *Veiculo {
	entity := new(Veiculo)
	copyToEntity(dto, entity)
	return entity
}

func MapperToDTO(entity *Veiculo) *domain.Veiculo {
	dto := new(domain.Veiculo)
	copyToDto(entity, dto)
	return dto
}

func copyToEntity(source *domain.Veiculo, destiny *Veiculo) {
	destiny.Id = source.Id
	destiny.IdCliente = source.IdCliente
	destiny.Placa = source.Placa
	destiny.Pais = source.Pais
	destiny.Cor = source.Cor
	destiny.Combustivel = source.Combustivel
	destiny.Renavam = source.Renavam
	destiny.Chassi = source.Chassi
	destiny.Ano = source.Ano
	destiny.IdModelo = source.IdModelo
}

func copyToDto(source *Veiculo, destiny *domain.Veiculo) {
	destiny.Id = source.Id
	destiny.IdCliente = source.IdCliente
	destiny.Placa = source.Placa
	destiny.Pais = source.Pais
	destiny.Cor = source.Cor
	destiny.Combustivel = source.Combustivel
	destiny.Renavam = source.Renavam
	destiny.Chassi = source.Chassi
	destiny.Ano = source.Ano
	destiny.IdModelo = source.IdModelo
}
