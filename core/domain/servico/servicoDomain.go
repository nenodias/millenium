package servico

import (
	"github.com/nenodias/millenium/core/domain"
)

type Servico struct {
	Id        int64   `json:"id"`
	Valor     float64 `json:"valor"`
	Descricao string  `json:"descricao"`
}

func (t Servico) GetId() int64 {
	return t.Id
}

type ServicoFilter struct {
	domain.Pageable
	Descricao string
}

type VeiculoService interface {
	domain.Service[*Servico, *ServicoFilter]
}
