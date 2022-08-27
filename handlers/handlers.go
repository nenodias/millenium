package handlers

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/nenodias/millenium/core/domain"
	core "github.com/nenodias/millenium/core/domain"
	"github.com/nenodias/millenium/core/domain/utils"
	"github.com/rs/zerolog/log"
)

type CrudAPI interface {
	FindMany(w http.ResponseWriter, r *http.Request)
	Save(w http.ResponseWriter, r *http.Request)
	FindOne(w http.ResponseWriter, r *http.Request)
	DeleteOne(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}

type Controller[T domain.Identifiable, F domain.PageableFilter] struct {
	Service    core.Service[*T, *F]
	GetFilters func(url.Values) F
	SetModelId func(int64, *T)
}

func (t *Controller[T, F]) FindMany(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	filter := t.GetFilters(query)
	response, err := t.Service.FindMany(&filter)
	if err != nil {
		log.Error().Msg(err.Error())
		w.WriteHeader(500)
	} else {
		utils.WriteJson(response, w, 200, 500)
	}
}

func (t *Controller[T, F]) Save(w http.ResponseWriter, r *http.Request) {
	model := new(T)
	err := json.NewDecoder(r.Body).Decode(model)
	if err != nil {
		log.Error().Msg(err.Error())
		w.WriteHeader(400)
		return
	}
	success, err := t.Service.Save(model)
	if err != nil {
		log.Error().Msg(err.Error())
		w.WriteHeader(500)
		return
	}
	if success {
		utils.WriteJson(model, w, 200, 500)
	} else {
		w.WriteHeader(400)
	}
}

func (t *Controller[T, F]) FindOne(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	txtId := params["id"]
	id := utils.StringToInt64(txtId, 0)
	model, err := t.Service.FindOne(id)
	if err != nil {
		log.Error().Msg(err.Error())
		w.WriteHeader(500)
		return
	}
	if model != nil {
		utils.WriteJson(model, w, 200, 500)
	} else {
		w.WriteHeader(204)
	}
}

func (t *Controller[T, F]) DeleteOne(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	txtId := params["id"]
	id := utils.StringToInt64(txtId, 0)

	_, err := t.Service.DeleteOne(id)
	if err != nil {
		log.Error().Msg(err.Error())
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(204)
}

func (t *Controller[T, F]) Update(w http.ResponseWriter, r *http.Request) {
	model := new(T)
	params := mux.Vars(r)
	txtId := params["id"]
	id := utils.StringToInt64(txtId, 0)
	record, err := t.Service.FindOne(id)
	if err != nil {
		log.Error().Msg(err.Error())
		w.WriteHeader(500)
		return
	}
	if record != nil {
		err := json.NewDecoder(r.Body).Decode(model)
		if err != nil {
			log.Error().Msg(err.Error())
			w.WriteHeader(400)
			return
		}
		t.SetModelId(id, model)
		success, err := t.Service.Save(model)
		if err != nil {
			log.Error().Msg(err.Error())
			w.WriteHeader(500)
			return
		}
		if success {
			utils.WriteJson(model, w, 200, 500)
		} else {
			w.WriteHeader(400)
		}
	} else {
		w.WriteHeader(404)
	}
}
