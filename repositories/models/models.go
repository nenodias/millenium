package models

import (
	core "github.com/nenodias/millenium/core/domain"
	"xorm.io/xorm"
)

type GenericRepository[T core.Identifiable, F any, MODEL any] struct {
	MapperToDTO    func(*MODEL) *T
	MapperToEntity func(*T) *MODEL
	CopyToDto      func(*MODEL, *T)
	DB             *xorm.Engine
}

func (gr *GenericRepository[T, F, MODEL]) FindOne(id int64) (*T, error) {
	p := new(MODEL)
	exists, err := gr.DB.ID(id).Get(p)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, nil
	}
	return gr.MapperToDTO(p), nil
}

func (gr *GenericRepository[T, F, MODEL]) Save(dto *T) (bool, error) {
	entity := gr.MapperToEntity(dto)
	model := new(MODEL)
	id := (*dto).GetId()
	exists, err := gr.DB.ID(id).Exist(model)
	if err != nil {
		return false, err
	}
	if !exists {
		rowsAffected, err := gr.DB.InsertOne(entity)
		if err != nil {
			return false, err
		}
		gr.CopyToDto(entity, dto)
		return rowsAffected == 1, nil
	}
	rowsAffected, err := gr.DB.ID(id).Update(entity)
	if err != nil {
		return false, err
	}
	gr.CopyToDto(entity, dto)
	return rowsAffected == 1, nil
}

func (gr *GenericRepository[T, F, MODEL]) DeleteOne(id int64) (bool, error) {
	model := new(MODEL)
	exists, err := gr.DB.ID(id).Exist(model)
	if err != nil {
		return false, err
	}
	if !exists {
		return true, nil
	} else {
		rowsAffected, err := gr.DB.ID(id).Delete(model)
		if err != nil {
			return false, err
		}
		return rowsAffected == 1, nil
	}
}
