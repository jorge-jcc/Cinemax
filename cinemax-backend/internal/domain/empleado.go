package domain

type Empleado struct {
	Id        string `db:"EMPLEADO_ID"`
	Nombre    string `db:"NOMBRE"`
	ApPaterno string `db:"AP_PATERNO"`
	ApMaterno string `db:"AP_MATERNO"`
	Rfc       string `db:"RFC"`
	Edad      int8   `db:"EDAD"`
	Email     string `db:"EMAIL"`
	Direccion string `db:"DIRECCION"`
	Telefono  string `db:"TELEFONO"`
	Password  string `db:"PASSWORD"`
}

func NewEmpleado(nombre, apPaterno, ApMaterno, rfc, email, direccion, telefono,
	password string, edad int8,
) *Empleado {
	return &Empleado{
		Nombre:    nombre,
		ApPaterno: apPaterno,
		ApMaterno: ApMaterno,
		Rfc:       rfc,
		Edad:      edad,
		Email:     email,
		Direccion: direccion,
		Telefono:  telefono,
		Password:  password,
	}
}
