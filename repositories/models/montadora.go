package models

type Montadora struct {
	Id                int64  `xorm:"'id' bigint pk autoincr not null"`
	Nome              string `xorm:"'nome_montadora' varchar(20) not null"`
	Origem            string `xorm:"'origem' varchar(1) not null"`
	CodigoMontadoraEA int    `xorm:"'codmon_ea' int"`
}

func (p *Montadora) TableName() string {
	return "montadora"
}
