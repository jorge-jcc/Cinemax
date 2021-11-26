package domain

type Sala struct {
	ID       string   `json:"id" db:"SALA_ID"`
	Clave    string   `json:"clave" db:"CLAVE"`
	Nombre   string   `json:"nombre" db:"NOMBRE"`
	TipoSala TipoSala `json:"tipoSala" db:"TIPO_SALA"`
}

type TipoSala struct {
	ID          string `json:"id" db:"TIPO_SALA_ID"`
	Clave       string `json:"clave" db:"CLAVE"`
	Descripcion string `json:"descripcion" db:"DESCRIPCION"`
}
