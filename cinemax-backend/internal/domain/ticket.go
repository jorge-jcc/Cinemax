package domain

type Ticket struct {
	Id            string `db:"TICKET_ID"`
	Monto         float32
	FechaCompra   string
	FuncionId     string
	EmpleadoId    string
	Boletos       []Boleto
	TransaccionId string
}
