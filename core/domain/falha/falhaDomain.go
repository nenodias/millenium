package falha

import (
	"github.com/nenodias/millenium/core/domain"
)

type Falha struct {
	Id        int64  `json:"id"`
	Descricao string `json:"descricao"`
}

func (t Falha) GetId() int64 {
	return t.Id
}

type FalhaFilter struct {
	domain.Pageable
	Descricao string
}

type VeiculoService interface {
	domain.Service[*Falha, *FalhaFilter]
}
