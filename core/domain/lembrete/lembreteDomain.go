package lembrete

import (
	"time"

	"github.com/nenodias/millenium/core/domain"
)

type Lembrete struct {
	Id        int64     `json:"id"`
	Texto     string    `json:"texto"`
	IdCliente int64     `json:"idCliente"`
	IdVeiculo int64     `json:"idVeiculo"`
	Data      time.Time `json:"dataNotificacao"`
}

func (t Lembrete) GetId() int64 {
	return t.Id
}

type LembreteFilter struct {
	domain.Pageable
	Texto     string
	IdCliente int64
	IdVeiculo int64
}

type LembreteService interface {
	domain.Service[*Lembrete, *LembreteFilter]
}
