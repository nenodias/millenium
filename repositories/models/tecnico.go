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
	DB *xorm.Engine
}

func NewTecnicoService(engine *xorm.Engine) domain.TecnicoService {
	return &TecnicoRepository{
		DB: engine,
	}
}

func (tf *TecnicoRepository) FindMany(filter *domain.TecnicoFilter) (core.PagebleContent[*domain.Tecnico], error) {
	response := core.PagebleContent[*domain.Tecnico]{}
	response.Number = filter.PageNumber
	response.Pageable.PageNumber = filter.PageNumber
	response.Pageable.PageSize = filter.PageSize
	response.Pageable.Sort.SortColumn = filter.Sort.SortColumn
	response.Pageable.Sort.SortDirection = filter.Sort.SortDirection
	orderBy := fmt.Sprintf("%s %s", filter.Sort.SortColumn, filter.Sort.SortDirection)
	query := tf.DB.OrderBy(orderBy)
	model := new(Tecnico)
	if filter.Nome != "" && strings.TrimSpace(filter.Nome) != "" {
		query = query.Where("nome Like ?", "%"+filter.Nome+"%")
	}
	total, _ := query.Count(model)
	response.TotalElements = total
	response.TotalPages = int64(math.Ceil(float64(total) / float64(filter.PageSize)))
	rows, err := query.Limit(filter.PageSize, filter.PageNumber*filter.PageSize).Rows(model)
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

func (tf *TecnicoRepository) FindOne(id int64) (*domain.Tecnico, error) {
	p := new(Tecnico)
	exists, err := tf.DB.ID(id).Get(p)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, nil
	}
	return mapperToDTO(p), nil
}

func (tf *TecnicoRepository) DeleteOne(id int64) (bool, error) {
	model := new(Tecnico)
	exists, err := tf.DB.ID(id).Exist(model)
	if err != nil {
		return false, err
	}
	if !exists {
		return true, nil
	} else {
		rowsAffected, err := tf.DB.ID(id).Delete(model)
		if err != nil {
			return false, err
		}
		return rowsAffected == 1, nil
	}
}

func (tf *TecnicoRepository) Save(dto *domain.Tecnico) (bool, error) {
	entity := mapperToEntity(dto)
	exists, err := tf.DB.ID(dto.Id).Exist(new(Tecnico))
	if err != nil {
		return false, err
	}
	if !exists {
		rowsAffected, err := tf.DB.InsertOne(entity)
		if err != nil {
			return false, err
		}
		copyToDto(entity, dto)
		return rowsAffected == 1, nil
	}
	rowsAffected, err := tf.DB.ID(entity.Id).Update(entity)
	if err != nil {
		return false, err
	}
	copyToDto(entity, dto)
	return rowsAffected == 1, nil
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
