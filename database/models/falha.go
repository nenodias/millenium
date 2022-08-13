package models

type Falha struct {
	Id        int64  `xorm:"'id' bigint pk autoincr not null"`
	Descricao string `xorm:"'descricao' varchar(60) not null"`
}

func (p *Falha) TableName() string {
	return "falhas"
}
