package domain

type Boleto struct {
	ID           string `db:"BOLETO_ID"`
	TipoBoletoId string `db:"TIPO_BOLETO_ID"`
	AsientoId    string `db:"BOLETO_ASIENTO_ID"`
}

type TipoBoleto struct {
	ID          string `json:"id" db:"FUNCION_ID"`
	Clave       string `json:"clave" db:"FUNCION_ID"`
	Descripcion string `json:"descripcion" db:"FUNCION_ID"`
}

type PrecioBoleto struct {
	ID            string  `json:"id" db:"TIPO_BOLETO_ID"`
	Clave         string  `json:"clave" db:"CLAVE"`
	Precio        float32 `json:"precio" db:"PRECIO"`
	TipoFuncionID string  `json:"tipoFuncionIs" db:"TIPO_FUNCION_ID"`
}
