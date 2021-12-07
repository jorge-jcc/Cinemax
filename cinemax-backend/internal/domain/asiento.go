package domain

import "time"

type Asiento struct {
	ID    string `json:"id" db:"ASIENTO_ID"`
	Clave string `json:"clave" db:"CLAVE"`
	Sala  *Sala  `json:"sala" db:"SALA_ID"`
}

type AsignacionAsiento struct {
	ID            string `json:"id" db:"ASIGNACION_ASIENTO_ID"`
	Clave         string `json:"clave" db:"CLAVE"`
	FuncionID     string `json:"funcionId" db:"FUNCION_ID"`
	TransaccionId string `json:"transaccionId" db:"TRANSACCION_ID"`
	BoletoId      *string
	StatusAsiento string    `db:"STATUS"`
	UpdatedAt     time.Time `db:"UPDATED_AT"`
}

type StatusAsiento string

const (
	StatusDisponible   StatusAsiento = "DISPONIBLE"
	StatusSeleccionado StatusAsiento = "SELECCIONADO"
	StatusEnCompra     StatusAsiento = "EN_COMPRA"
	StatusAsignado     StatusAsiento = "ASIGNADO"
	StatusDesconocido  StatusAsiento = "DESCONOCIDO"
)
