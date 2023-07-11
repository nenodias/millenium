package models

import (
	"context"
	"fmt"
	"math"

	core "github.com/nenodias/millenium/core/domain"
	"github.com/rs/zerolog/log"
	"xorm.io/xorm"
)

type GenericRepository[T core.Identifiable, F core.PageableFilter, MODEL any] struct {
	MapperToDTO    func(context.Context, *MODEL) *T
	MapperToEntity func(context.Context, *T) *MODEL
	CopyToDto      func(context.Context, *MODEL, *T)
	HasWhere       func(context.Context, *F) bool
	DoWhere        func(context.Context, *xorm.Session, *F) *xorm.Session
	DB             *xorm.Engine
	AfterFind      func(context.Context, *GenericRepository[T, F, MODEL], *MODEL)
	AfterSave      func(context.Context, *GenericRepository[T, F, MODEL], *xorm.Session, *MODEL) bool
	AfterUpdate    func(context.Context, *GenericRepository[T, F, MODEL], *xorm.Session, *MODEL) bool
	AfterDelete    func(context.Context, *GenericRepository[T, F, MODEL], *xorm.Session, int64) bool
}

func (gr *GenericRepository[T, F, MODEL]) FindOne(ctx context.Context, id int64) (*T, error) {
	p := new(MODEL)
	exists, err := gr.DB.ID(id).Get(p)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}
	if !exists {
		return nil, nil
	}
	if gr.AfterFind != nil {
		gr.AfterFind(ctx, gr, p)
	}
	return gr.MapperToDTO(ctx, p), nil
}

func (gr *GenericRepository[T, F, MODEL]) Save(ctx context.Context, dto *T) (bool, error) {
	entity := gr.MapperToEntity(ctx, dto)
	model := new(MODEL)
	id := (*dto).GetId()
	session := gr.DB.NewSession()
	err := session.Begin()
	if err != nil {
		log.Error().Msg(err.Error())
		session.Rollback()
		return false, err
	}
	exists, err := session.ID(id).Exist(model)
	if err != nil {
		log.Error().Msg(err.Error())
		session.Rollback()
		return false, err
	}
	if !exists {
		rowsAffected, err := gr.DB.InsertOne(entity)
		if err != nil {
			log.Error().Msg(err.Error())
			session.Rollback()
			return false, err
		}
		if gr.AfterSave != nil {
			if !gr.AfterSave(ctx, gr, session, entity) {
				session.Rollback()
				return false, fmt.Errorf("error on after save")
			}
		}
		gr.CopyToDto(ctx, entity, dto)
		err = session.Commit()
		if err != nil {
			log.Error().Msg(err.Error())
			return false, err
		}
		return rowsAffected == 1, nil
	}
	rowsAffected, err := gr.DB.ID(id).Update(entity)
	if gr.AfterUpdate != nil {
		if !gr.AfterUpdate(ctx, gr, session, entity) {
			session.Rollback()
			return false, fmt.Errorf("error on after update")
		}
	}
	if err != nil {
		log.Error().Msg(err.Error())
		session.Rollback()
		return false, err
	}
	gr.CopyToDto(ctx, entity, dto)
	err = session.Commit()
	if err != nil {
		log.Error().Msg(err.Error())
		return false, err
	}
	return rowsAffected == 1, nil
}

func (gr *GenericRepository[T, F, MODEL]) DeleteOne(ctx context.Context, id int64) (bool, error) {
	model := new(MODEL)
	session := gr.DB.NewSession()
	err := session.Begin()
	if err != nil {
		log.Error().Msg(err.Error())
		session.Rollback()
		return false, err
	}
	exists, err := session.ID(id).Exist(model)
	if err != nil {
		log.Error().Msg(err.Error())
		session.Rollback()
		return false, err
	}
	if !exists {
		session.Rollback()
		return true, nil
	} else {
		rowsAffected, err := session.ID(id).Delete(model)
		if err != nil {
			log.Error().Msg(err.Error())
			session.Rollback()
			return false, err
		}
		if gr.AfterDelete != nil {
			if !gr.AfterDelete(ctx, gr, session, id) {
				session.Rollback()
				return false, fmt.Errorf("error on after delete")
			}
		}
		err = session.Commit()
		if err != nil {
			log.Error().Msg(err.Error())
			return false, err
		}
		return rowsAffected == 1, nil
	}
}

func (gr *GenericRepository[T, F, MODEL]) FindMany(ctx context.Context, filter *F) (core.PagebleContent[*T], error) {
	response := core.PagebleContent[*T]{
		Number: (*filter).GetPageNumber(),

		Pageable: core.Pageable{
			PageNumber: (*filter).GetPageNumber(),
			PageSize:   (*filter).GetPageSize(),
			Sort:       (*filter).GetSort(),
		},
	}

	orderBy := fmt.Sprintf("%s %s", (*filter).GetSort().SortColumn, (*filter).GetSort().SortDirection)
	query := gr.DB.OrderBy(orderBy)

	model := new(MODEL)
	var total int64
	hasWhere := gr.HasWhere(ctx, filter)

	if hasWhere {
		total, _ = gr.DoWhere(ctx, query, filter).Count(model)
	} else {
		total, _ = query.Count(model)
	}
	response.TotalElements = total
	response.TotalPages = int64(math.Ceil(float64(total) / float64((*filter).GetPageSize())))

	var rows *xorm.Rows
	var err error
	offset := (*filter).GetPageNumber() * (*filter).GetPageSize()
	if hasWhere {
		rows, err = gr.DoWhere(ctx, query, filter).OrderBy(orderBy).Limit(
			(*filter).GetPageSize(), offset,
		).Rows(model)
	} else {
		rows, err = query.OrderBy(orderBy).Limit(
			(*filter).GetPageSize(), offset,
		).Rows(model)
	}
	if err != nil {
		log.Error().Msg(err.Error())
		return response, err
	}
	defer rows.Close()
	response.Content = make([]*T, 0)
	count := 0
	for rows.Next() {
		err = rows.Scan(model)
		if err != nil {
			log.Error().Msg(err.Error())
		}
		count++
		if gr.AfterFind != nil {
			gr.AfterFind(ctx, gr, model)
		}
		response.Content = append(response.Content, gr.MapperToDTO(ctx, model))

	}
	response.Size = count
	return response, nil
}