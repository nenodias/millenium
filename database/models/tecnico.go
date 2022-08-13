package models

type Tecnico struct {
	Id   int64  `xorm:"'codigo_tecnico' bigint pk autoincr not null"`
	Nome string `xorm:"'nome' varchar(60) not null"`
}

func (p *Tecnico) TableName() string {
	return "tecnico"
}
