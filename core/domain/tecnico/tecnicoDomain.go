package tecnico

import (
	"github.com/nenodias/millenium/core/domain"
)

type Tecnico struct {
	Id   int64
	Nome string
}

type TecnicoFilter struct {
	domain.Pageable
	Nome string
}

type TecnicoService interface {
	domain.Service[*Tecnico, *TecnicoFilter]
}
