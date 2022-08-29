package cliente

import (
	"strings"
	"time"

	domain "github.com/nenodias/millenium/core/domain/cliente"
	models "github.com/nenodias/millenium/repositories/models"
	"xorm.io/xorm"
)

type Cliente struct {
	Id                int64     `xorm:"'id' bigint pk autoincr not null"`
	Nome              string    `xorm:"'nome' varchar(60) not null"`
	RG                string    `xorm:"'ie_rg' varchar(16)"`
	CPF               string    `xorm:"'cpf' varchar(19)"`
	Endereco          string    `xorm:"'endereco' varchar(50)"`
	Complemento       string    `xorm:"'complemento' varchar(30)"`
	Bairro            string    `xorm:"'bairro' varchar(30)"`
	Cidade            string    `xorm:"'cidade' varchar(30)"`
	CEP               string    `xorm:"'cep' varchar(9)"`
	Estado            string    `xorm:"'estado' varchar(2)"`
	Pais              string    `xorm:"'pais' varchar(20)"`
	Telefone          string    `xorm:"'telefone' varchar(20)"`
	Fax               string    `xorm:"'fax' varchar(20)"`
	Celular           string    `xorm:"'celular' varchar(20)"`
	TelefoneComercial string    `xorm:"'tel_comercial' varchar(20)"`
	FaxComercial      string    `xorm:"'fax_comercial' varchar(20)"`
	Email             string    `xorm:"'email' varchar(40)"`
	BIP               string    `xorm:"'bip' varchar(30)"`
	DataNascimento    time.Time `xorm:"'data_nascimento' timestamp"`
	Mes               int       `xorm:"'mes' int"`
}

func (p *Cliente) TableName() string {
	return "cliente"
}

type ClienteRepository struct {
	models.GenericRepository[domain.Cliente, domain.ClienteFilter, Cliente]
}

func NewService(engine *xorm.Engine) domain.ClienteService {
	repository := ClienteRepository{
		GenericRepository: models.GenericRepository[domain.Cliente, domain.ClienteFilter, Cliente]{
			DB:             engine,
			MapperToDTO:    MapperToDTO,
			MapperToEntity: MapperToEntity,
			CopyToDto:      copyToDto,
			HasWhere:       hasWhere,
			DoWhere:        doWhere,
		},
	}
	return domain.ClienteService(&repository)
}

func hasWhere(filter *domain.ClienteFilter) bool {
	hasNome := filter.Nome != "" && strings.TrimSpace(filter.Nome) != ""
	hasTelefone := filter.Telefone != "" && strings.TrimSpace(filter.Telefone) != ""
	hasCelular := filter.Celular != "" && strings.TrimSpace(filter.Celular) != ""
	return hasNome || hasTelefone || hasCelular
}

func doWhere(query *xorm.Session, filter *domain.ClienteFilter) *xorm.Session {
	hasNome := filter.Nome != "" && strings.TrimSpace(filter.Nome) != ""
	hasTelefone := filter.Telefone != "" && strings.TrimSpace(filter.Telefone) != ""
	hasCelular := filter.Celular != "" && strings.TrimSpace(filter.Celular) != ""
	where := make([]interface{}, 0)
	if hasNome {
		where = append(where, "nome Like ?", "%"+filter.Nome+"%")
	}
	if hasTelefone {
		where = append(where, "telefone Like ?", "%"+filter.Telefone+"%")
	}
	if hasCelular {
		where = append(where, "celular Like ?", "%"+filter.Celular+"%")
	}
	if len(where) == 2 {
		return query.Where(where[0], where[1])
	} else if len(where) == 4 {
		return query.Where(where[0], where[1]).Or(where[2], where[3])
	} else {
		return query.Where(where[0], where[1]).Or(where[2], where[3]).Or(where[4], where[5])
	}
}

func MapperToEntity(dto *domain.Cliente) *Cliente {
	entity := new(Cliente)
	copyToEntity(dto, entity)
	return entity
}

func MapperToDTO(entity *Cliente) *domain.Cliente {
	dto := new(domain.Cliente)
	copyToDto(entity, dto)
	return dto
}

func copyToEntity(source *domain.Cliente, destiny *Cliente) {
	destiny.Id = source.Id
	destiny.Nome = source.Nome
	destiny.RG = source.RG
	destiny.CPF = source.CPF
	destiny.Endereco = source.Endereco
	destiny.Complemento = source.Complemento
	destiny.Bairro = source.Bairro
	destiny.Cidade = source.Cidade
	destiny.CEP = source.CEP
	destiny.Estado = source.Estado
	destiny.Pais = source.Pais
	destiny.Telefone = source.Telefone
	destiny.Fax = source.Fax
	destiny.Celular = source.Celular
	destiny.TelefoneComercial = source.TelefoneComercial
	destiny.FaxComercial = source.FaxComercial
	destiny.Email = source.Email
	destiny.BIP = source.BIP
	destiny.DataNascimento = source.DataNascimento
	destiny.Mes = source.Mes
}

func copyToDto(source *Cliente, destiny *domain.Cliente) {
	destiny.Id = source.Id
	destiny.Nome = source.Nome
	destiny.RG = source.RG
	destiny.CPF = source.CPF
	destiny.Endereco = source.Endereco
	destiny.Complemento = source.Complemento
	destiny.Bairro = source.Bairro
	destiny.Cidade = source.Cidade
	destiny.CEP = source.CEP
	destiny.Estado = source.Estado
	destiny.Pais = source.Pais
	destiny.Telefone = source.Telefone
	destiny.Fax = source.Fax
	destiny.Celular = source.Celular
	destiny.TelefoneComercial = source.TelefoneComercial
	destiny.FaxComercial = source.FaxComercial
	destiny.Email = source.Email
	destiny.BIP = source.BIP
	destiny.DataNascimento = source.DataNascimento
	destiny.Mes = source.Mes
}
