package domain

import (
	"time"
)

type Pelicula struct {
	ID                 string        `json:"id,omitempty" db:"PELICULA_ID"`
	Nombre             string        `json:"nombre" db:"NOMBRE"`
	Director           string        `json:"director" db:"DIRECTOR"`
	Descripcion        string        `json:"descripcion" db:"DESCRIPCION"`
	DuracionMinutos    int16         `json:"duracionMinutos" db:"DURACION_MINUTOS"`
	Anio               string        `json:"anio" db:"ANIO"`
	FechaDisponiblidad time.Time     `json:"fechaDisponibilidad" db:"FECHA_DISPONIBILIDAD"`
	Imagen             string        `json:"imagen,omitempty" db:"IMAGEN"`
	Resena             string        `json:"resena" db:"RESENA"`
	Clasificacion      Clasificacion `json:"clasificacion" db:"CLASIFICACION"`
	Idioma             Idioma        `json:"idioma" db:"IDIOMA"`
	Subtitulo          Idioma        `json:"subtitulos" db:"SUBTITULO"`
	Genero             Genero        `json:"genero" db:"GENERO"`
}

type Clasificacion struct {
	ID          string `json:"id,omitempty" db:"CLASIFICACION_ID"`
	Clave       string `json:"clave" db:"CLAVE"`
	Descripcion string `json:"descripcion" db:"DESCRIPCION"`
}

type Idioma struct {
	ID     string `json:"id,omitempty" db:"IDIOMA_ID"`
	Nombre string `json:"nombre" db:"NOMBRE"`
}

type Genero struct {
	ID     string `json:"id,omitempty" db:"GENERO_ID"`
	Nombre string `json:"nombre" db:"NOMBRE"`
}

func NewPelicula(nombre, director, descripcion string, duracionMinutos int16,
	anio string, fechaDisponibilidad time.Time, resena string,
	clasificacionId, idiomaId, subtituloId, generoId string,
) *Pelicula {
	clasificacion := Clasificacion{ID: clasificacionId}
	idioma := Idioma{ID: idiomaId}
	subtitulo := Idioma{ID: subtituloId}
	genero := Genero{ID: generoId}

	pelicula := &Pelicula{
		Nombre:             nombre,
		Director:           director,
		Descripcion:        descripcion,
		DuracionMinutos:    duracionMinutos,
		Anio:               anio,
		FechaDisponiblidad: fechaDisponibilidad,
		Resena:             resena,
		Clasificacion:      clasificacion,
		Idioma:             idioma,
		Subtitulo:          subtitulo,
		Genero:             genero,
	}
	return pelicula
}
