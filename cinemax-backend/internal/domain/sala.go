package domain

type Sala struct {
	ID        string `json:"id" db:"SALA_ID"`
	Clave     string `json:"clave" db:"CLAVE"`
	Nombre    string `json:"nombre" db:"NOMBRE"`
	Ubicacion string `json:"ubicacion" db:"UBICACION"`
}
