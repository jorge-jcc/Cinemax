package domain

type Boleto struct {
	ID         string     `json:"id" db:"BOLETO_ID"`
	Ticket     string     `json:"ticket" db:""`
	TipoBoleto TipoBoleto `json:"tipoBoleto" db:"FUNCION_ID"`
}

type TipoBoleto struct {
	ID          string `json:"id" db:"FUNCION_ID"`
	Clave       string `json:"clave" db:"FUNCION_ID"`
	Descripcion string `json:"descripcion" db:"FUNCION_ID"`
}

type PrecioBoleto struct {
	Clave       string
	Precio      float32
	TipoFuncion string
}
