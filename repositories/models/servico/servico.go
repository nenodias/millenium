package servico

import (
	"strings"

	domain "github.com/nenodias/millenium/core/domain/servico"
	models "github.com/nenodias/millenium/repositories/models"
	"xorm.io/xorm"
)

type Servico struct {
	Id        int64   `xorm:"'id' bigint pk autoincr not null"`
	Valor     float64 `xorm:"'valor' double"`
	Descricao string  `xorm:"'descricao' varchar(60) not null"`
}

func (p *Servico) TableName() string {
	return "servicos"
}

type ServicoRepository struct {
	models.GenericRepository[domain.Servico, domain.ServicoFilter, Servico]
}

func NewService(engine *xorm.Engine) domain.ServicoService {
	repository := ServicoRepository{
		GenericRepository: models.GenericRepository[domain.Servico, domain.ServicoFilter, Servico]{
			DB:             engine,
			MapperToDTO:    mapperToDTO,
			MapperToEntity: mapperToEntity,
			CopyToDto:      copyToDto,
			HasWhere:       hasWhere,
			DoWhere:        doWhere,
		},
	}
	return domain.ServicoService(&repository)
}

func hasWhere(filter *domain.ServicoFilter) bool {
	return filter.Descricao != "" && strings.TrimSpace(filter.Descricao) != ""
}

func doWhere(query *xorm.Session, filter *domain.ServicoFilter) *xorm.Session {
	where := []string{"descricao Like ?", "%" + filter.Descricao + "%"}
	return query.Where(where[0], where[1])
}

func mapperToEntity(dto *domain.Servico) *Servico {
	entity := new(Servico)
	copyToEntity(dto, entity)
	return entity
}

func mapperToDTO(entity *Servico) *domain.Servico {
	dto := new(domain.Servico)
	copyToDto(entity, dto)
	return dto
}

func copyToEntity(source *domain.Servico, destiny *Servico) {
	destiny.Id = source.Id
	destiny.Descricao = source.Descricao
}

func copyToDto(source *Servico, destiny *domain.Servico) {
	destiny.Id = source.Id
	destiny.Descricao = source.Descricao
}
