package domain

import "time"

type Asiento struct {
	ID    string `json:"id" db:"ASIENTO_ID"`
	Clave string `json:"clave" db:"CLAVE"`
	Sala  *Sala  `json:"sala" db:"SALA_ID"`
}

type AsignacionAsiento struct {
	ID            string    `json:"id" db:"ASIGNACION_ASIENTO_ID"`
	Clave         string    `json:"clave" db:"CLAVE"`
	FuncionID     string    `json:"funcionId" db:"FUNCION_ID"`
	StatusID      string    `json:"statusId" db:"STATUS"`
	TransaccionId string    `json:"transaccionId" db:"TRANSACCION_ID"`
	UpdatedAt     time.Time `db:"UPDATED_AT"`
}
