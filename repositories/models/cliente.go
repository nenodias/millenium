package models

import "time"

type Cliente struct {
	Id                int64     `xorm:"'codigo_cliente' bigint pk autoincr not null"`
	Nome              string    `xorm:"'nome_cliente' varchar(60) not null"`
	RG                string    `xorm:"'ie_rg' varchar(16)"`
	CPF               string    `xorm:"'cgc' varchar(19)"`
	Endereco          string    `xorm:"'endereco' varchar(50)"`
	Complemento       string    `xorm:"'complemento' varchar(30)"`
	Bairro            string    `xorm:"'bairro' varchar(30)"`
	Cidade            string    `xorm:"'cidade' varchar(30)"`
	CEP               string    `xorm:"'cep' varchar(9)"`
	Estado            string    `xorm:"'estado' varchar(2)"`
	Pais              string    `xorm:"'pais' varchar(20)"`
	Telefone          string    `xorm:"'tel_res' varchar(20)"`
	Fax               string    `xorm:"'fax_res' varchar(20)"`
	Celular           string    `xorm:"'celular' varchar(20)"`
	TelefoneComercial string    `xorm:"'tel_com' varchar(20)"`
	FaxComercial      string    `xorm:"'fax_com' varchar(20)"`
	Email             string    `xorm:"'e_mail' varchar(40)"`
	BIP               string    `xorm:"'bip_cod' varchar(30)"`
	DataNascimento    time.Time `xorm:"'dtnasc' timestamp"`
	Mes               int       `xorm:"'mes' int"`
}

func (p *Cliente) TableName() string {
	return "clientes"
}
