package models

import (
	"fmt"
	"math"
	"strings"

	"github.com/rs/zerolog/log"

	core "github.com/nenodias/millenium/core/domain"
	domain "github.com/nenodias/millenium/core/domain/tecnico"
	"xorm.io/xorm"
)

type Tecnico struct {
	Id   int64  `xorm:"'codigo_tecnico' bigint pk autoincr not null"`
	Nome string `xorm:"'nome' varchar(60) not null"`
}

func (p *Tecnico) TableName() string {
	return "tecnico"
}

type TecnicoRepository struct {
	GenericRepository[domain.Tecnico, domain.TecnicoFilter, Tecnico]
}

func NewTecnicoService(engine *xorm.Engine) domain.TecnicoService {
	repository := TecnicoRepository{
		GenericRepository: GenericRepository[domain.Tecnico, domain.TecnicoFilter, Tecnico]{
			DB:             engine,
			MapperToDTO:    mapperToDTO,
			MapperToEntity: mapperToEntity,
			CopyToDto:      copyToDto,
		},
	}
	return domain.TecnicoService(&repository)
}

func (tf *TecnicoRepository) FindMany(filter *domain.TecnicoFilter) (core.PagebleContent[*domain.Tecnico], error) {
	response := core.PagebleContent[*domain.Tecnico]{
		Number: filter.PageNumber,

		Pageable: core.Pageable{
			PageNumber: filter.PageNumber,
			PageSize:   filter.PageSize,
			Sort:       filter.Sort,
		},
	}

	orderBy := fmt.Sprintf("%s %s", filter.Sort.SortColumn, filter.Sort.SortDirection)
	query := tf.DB.OrderBy(orderBy)

	model := new(Tecnico)
	var total int64
	where := []string{}
	if filter.Nome != "" && strings.TrimSpace(filter.Nome) != "" {
		where = []string{"nome Like ?", "%" + filter.Nome + "%"}
	}

	if len(where) > 0 {
		total, _ = query.Where(where[0], where[1]).Count(model)
	} else {
		total, _ = query.Count(model)
	}
	response.TotalElements = total
	response.TotalPages = int64(math.Ceil(float64(total) / float64(filter.PageSize)))

	var rows *xorm.Rows
	var err error
	if len(where) > 0 {
		rows, err = query.Where(where[0], where[1]).OrderBy(orderBy).Limit(
			filter.PageSize, filter.PageNumber*filter.PageSize,
		).Rows(model)
	} else {
		rows, err = query.OrderBy(orderBy).Limit(
			filter.PageSize, filter.PageNumber*filter.PageSize,
		).Rows(model)
	}
	if err != nil {
		return response, err
	}
	defer rows.Close()
	response.Content = make([]*domain.Tecnico, 0)
	count := 0
	for rows.Next() {
		err = rows.Scan(model)
		if err != nil {
			log.Error().Msg(err.Error())
		}
		count++
		response.Content = append(response.Content, mapperToDTO(model))

	}
	response.Size = count
	return response, nil
}

func mapperToEntity(dto *domain.Tecnico) *Tecnico {
	entity := new(Tecnico)
	copyToEntity(dto, entity)
	return entity
}

func mapperToDTO(entity *Tecnico) *domain.Tecnico {
	dto := new(domain.Tecnico)
	copyToDto(entity, dto)
	return dto
}

func copyToEntity(source *domain.Tecnico, destiny *Tecnico) {
	destiny.Id = source.Id
	destiny.Nome = source.Nome
}

func copyToDto(source *Tecnico, destiny *domain.Tecnico) {
	destiny.Id = source.Id
	destiny.Nome = source.Nome
}
