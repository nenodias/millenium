package models

type Servico struct {
	Id        int64   `xorm:"'id' bigint pk autoincr not null"`
	Valor     float64 `xorm:"'valor' double"`
	Descricao string  `xorm:"'descricao' varchar(60) not null"`
}

func (p *Servico) TableName() string {
	return "servicos"
}
