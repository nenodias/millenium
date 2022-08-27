package lembrete

import "time"

type Lembrete struct {
	Id        int64     `xorm:"'id' bigint pk autoincr not null"`
	Texto     string    `xorm:"'texto' varchar(5000)"`
	IdCliente int64     `xorm:"'id_cliente' bigint"`
	IdVeiculo int64     `xorm:"'id_veiculo' bigint"`
	Data      time.Time `xorm:"'data_notificacao' timestamp"`
}

func (p *Lembrete) TableName() string {
	return "lembretes"
}
