package models

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
	Id          int64       `xorm:"'codveiculo' bigint pk autoincr not null"`
	IdCliente   int64       `xorm:"'codigo_cliente' bigint"`
	Placa       string      `xorm:"'placa' varchar(8) not null"`
	Pais        string      `xorm:"'pais' varchar(20)"`
	Cor         string      `xorm:"'cor' varchar(20)"`
	Combustivel Combustivel `xorm:"'combustivel' varchar(10)"`
	Renavan     string      `xorm:"'renavan' varchar(40)"`
	Chassi      string      `xorm:"'chassi' varchar(40)"`
	Ano         string      `xorm:"'ano' varchar(4)"`
	IdModelo    int64       `xorm:"'id_modelo' bigint"`
}

func (p *Veiculo) TableName() string {
	return "veiculo"
}
