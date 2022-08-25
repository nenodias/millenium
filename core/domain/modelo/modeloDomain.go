package modelo

import (
	"github.com/nenodias/millenium/core/domain"
)

type Modelo struct {
	Id              int64  `json:"id"`
	Nome            string `json:"nome"`
	IdMontadora     int64  `json:"idMontadora"`
	CodigoVeiculoEA int    `json:"codigoVeiculoEA"`
}

func (t Modelo) GetId() int64 {
	return t.Id
}

type ModeloFilter struct {
	domain.Pageable
	Nome     string
	IdModelo int64
}

type ModeloService interface {
	domain.Service[*Modelo, *ModeloFilter]
}
