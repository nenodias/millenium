package tecnico

import (
	"github.com/nenodias/millenium/core/domain"
)

type Tecnico struct {
	Id   int64
	Nome string
}

func (t Tecnico) GetId() int64 {
	return t.Id
}

type TecnicoFilter struct {
	domain.Pageable
	Nome string
}

type TecnicoService interface {
	domain.Service[*Tecnico, *TecnicoFilter]
}
