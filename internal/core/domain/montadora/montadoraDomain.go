package montadora

import (
	"github.com/nenodias/millenium/internal/core/domain"
)

type Montadora struct {
	Id                int64  `json:"id"`
	Nome              string `json:"nome"`
	Origem            string `json:"origem"`
	CodigoMontadoraEA int    `json:"codigoMontadoraEA"`
}

func (t Montadora) GetId() int64 {
	return t.Id
}

type MontadoraFilter struct {
	domain.Pageable
	Nome string
}

type MontadoraService interface {
	domain.Service[*Montadora, *MontadoraFilter]
}
