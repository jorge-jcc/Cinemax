package domain

type Sala struct {
	ID          string   `json:"id" db:"SALA_ID"`
	Nombre      string   `json:"nombre" db:"NOMBRE"`
	Descripcion string   `json:"descripcion" db:"DESCRIPCION"`
	TipoSala    TipoSala `json:"tipoSala" db:"TIPO_SALA"`
}

type TipoSala struct {
	ID          string `json:"id" db:"TIPO_SALA_ID"`
	Clave       string `json:"clave" db:"CLAVE"`
	Descripcion string `json:"descripcion" db:"DESCRIPCION"`
}
