package models

type Modelo struct {
	Id              int64  `xorm:"'id' bigint pk autoincr not null"`
	Nome            string `xorm:"'nome_modelo' varchar(40) not null"`
	IdMontadora     int64  `xorm:"'id_monta' bigint"`
	CodigoVeiculoEA int    `xorm:"'codvei_ea' int"`
}

func (p *Modelo) TableName() string {
	return "modelo"
}
