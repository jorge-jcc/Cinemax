package domain

import "time"

type Funcion struct {
	ID          string    `json:"id" db:"FUNCION_ID"`
	FechaInicio time.Time `json:"fechaInicio" db:"FECHA_INICIO"`
	FechaFin    time.Time `json:"fechaFin" db:"FECHA_FIN"`
	Pelicula    Pelicula  `json:"pelicula" db:"PELICULA"`
	Sala        Sala      `json:"sala" db:"SALA"`
	TipoFuncion string    `json:"tipoFuncion" db:"TIPO_FUNCION_ID"`
}
