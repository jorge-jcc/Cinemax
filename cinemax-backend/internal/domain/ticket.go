package domain

import "time"

type Ticket struct {
	Id            string `db:"TICKET_ID"`
	Monto         float32
	FechaCompra   time.Time
	FuncionId     string
	EmpleadoId    string
	Boletos       []Boleto
	TransaccionId string
}
