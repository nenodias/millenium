package peca

import (
	"github.com/nenodias/millenium/core/domain"
)

type Peca struct {
	Id        int64   `json:"id"`
	Valor     float64 `json:"valor"`
	Descricao string  `json:"descricao"`
}

func (t Peca) GetId() int64 {
	return t.Id
}

type PecaFilter struct {
	domain.Pageable
	Descricao string
}

type VeiculoService interface {
	domain.Service[*Peca, *PecaFilter]
}
