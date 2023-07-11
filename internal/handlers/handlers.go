package handlers

import (
	"context"
	"net/http"
	"net/url"

	"github.com/go-chi/chi/v5"
	"github.com/goccy/go-json"

	core "github.com/nenodias/millenium/internal/core/domain"
	"github.com/nenodias/millenium/internal/core/domain/utils"
	"github.com/rs/zerolog/log"
)

type CORSHandler struct {
	Inner  http.Handler
	Origin string
}

func (ch CORSHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", ch.Origin)
	w.Header().Add("Access-Control-Max-Age", "86400")
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
	} else {
		ch.Inner.ServeHTTP(w, r)
	}
}

type CrudAPI interface {
	FindMany(w http.ResponseWriter, r *http.Request)
	Save(w http.ResponseWriter, r *http.Request)
	FindOne(w http.ResponseWriter, r *http.Request)
	DeleteOne(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}

type Controller[T core.Identifiable, F core.PageableFilter] struct {
	Service    core.Service[*T, *F]
	GetFilters func(context.Context, url.Values) F
	SetModelId func(int64, *T)
}

func (t *Controller[T, F]) FindMany(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	filter := t.GetFilters(r.Context(), query)
	response, err := t.Service.FindMany(r.Context(), &filter)
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
	success, err := t.Service.Save(r.Context(), model)
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
	txtId := chi.URLParam(r, "id")
	id := utils.StringToInt64(txtId, 0)
	model, err := t.Service.FindOne(r.Context(), id)
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
	txtId := chi.URLParam(r, "id")
	id := utils.StringToInt64(txtId, 0)

	_, err := t.Service.DeleteOne(r.Context(), id)
	if err != nil {
		log.Error().Msg(err.Error())
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(204)
}

func (t *Controller[T, F]) Update(w http.ResponseWriter, r *http.Request) {
	model := new(T)
	txtId := chi.URLParam(r, "id")
	id := utils.StringToInt64(txtId, 0)
	record, err := t.Service.FindOne(r.Context(), id)
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
		success, err := t.Service.Save(r.Context(), model)
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
