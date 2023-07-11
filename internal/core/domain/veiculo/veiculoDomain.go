package veiculo

import (
	"github.com/nenodias/millenium/internal/core/domain"
)

type Combustivel string

const (
	GASOLINA Combustivel = "GASOLINA"
	ALCOOL   Combustivel = "ALCOOL"
	DIESEL   Combustivel = "DIESEL"
	FLEX     Combustivel = "FLEX"
	ELETRICO Combustivel = "ELETRICO"
	GAS      Combustivel = "GAS"
	OUTRO    Combustivel = "OUTRO"
)

type Veiculo struct {
	Id          int64       `json:"id"`
	IdCliente   int64       `json:"idCliente"`
	Placa       string      `json:"placa"`
	Pais        string      `json:"pais"`
	Cor         string      `json:"cor"`
	Combustivel Combustivel `json:"combustivel"`
	Renavam     string      `json:"renavam"`
	Chassi      string      `json:"chassi"`
	Ano         string      `json:"ano"`
	IdModelo    int64       `json:"idModelo"`
}

func (t Veiculo) GetId() int64 {
	return t.Id
}

type VeiculoFilter struct {
	domain.Pageable
	Placa     string
	IdCliente int64
	IdModelo  int64
}

type VeiculoService interface {
	domain.Service[*Veiculo, *VeiculoFilter]
}
