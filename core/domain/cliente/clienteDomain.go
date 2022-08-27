package cliente

import (
	"time"

	"github.com/nenodias/millenium/core/domain"
)

type Cliente struct {
	Id                int64     `json:"id"`
	Nome              string    `json:"nome"`
	RG                string    `json:"rg"`
	CPF               string    `json:"cpf"`
	Endereco          string    `json:"endereco"`
	Complemento       string    `json:"complemento"`
	Bairro            string    `json:"bairro"`
	Cidade            string    `json:"cidade"`
	CEP               string    `json:"cep"`
	Estado            string    `json:"estado"`
	Pais              string    `json:"pais"`
	Telefone          string    `json:"telefoneResidencial"`
	Fax               string    `json:"faxResidencial"`
	Celular           string    `json:"celular"`
	TelefoneComercial string    `json:"telefoneComercial"`
	FaxComercial      string    `json:"faxComercial"`
	Email             string    `json:"email"`
	BIP               string    `json:"bip"`
	DataNascimento    time.Time `json:"dataNascimento"`
	Mes               int       `json:"mes"`
}

func (t Cliente) GetId() int64 {
	return t.Id
}

type ClienteFilter struct {
	domain.Pageable
	Nome     string
	Telefone string
	Celular  string
}

type VeiculoService interface {
	domain.Service[*Cliente, *ClienteFilter]
}
