package domain

type Cartelera struct {
	Pelicula
	Horarios []string
}

type Horario struct {
	Id   string
	Hora string
}
