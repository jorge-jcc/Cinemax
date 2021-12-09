package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/domain"
)

func (r *repository) CreateFunction(ctx context.Context, f *domain.Funcion) error {
	query := `
		INSERT INTO "FUNCION" ("FECHA_INICIO", "FECHA_FIN", "PELICULA_ID", "SALA_ID", "TIPO_FUNCION_ID")
		VALUES ($1, $2, $3, $4, $5)
		RETURNING "FUNCION_ID"
	`
	err := r.db.GetContext(ctx, f, query,
		f.FechaInicio, f.FechaFin, f.Pelicula.ID, f.Sala.ID, f.TipoFuncion,
	)
	if err != nil {
		return domain.NewInternal()
	}
	return nil
}

func (r *repository) DisponibilidadFuncion(ctx context.Context, f *domain.Funcion) bool {
	query := `
		SELECT COUNT(*) VALUE FROM "FUNCION"
		WHERE	"SALA_ID" = $1
			AND "FECHA_INICIO" < $3
			AND "FECHA_FIN" > $2
	`
	var result int
	err := r.db.GetContext(ctx, &result, query, f.Sala.ID, f.FechaInicio, f.FechaFin)
	fmt.Println(result)
	if err != nil {
		return false
	}
	return result == 0
}

func (r *repository) GetFuncionesByPeliculaAndFechaInicio(
	ctx context.Context, peliculaId string, fecha time.Time,
) ([]domain.Funcion, error) {
	query := `
		select "FUNCION_ID", "FECHA_INICIO", "FECHA_FIN", 
			"PELICULA_ID" as "PELICULA.PELICULA_ID", 
			"CLAVE" AS "SALA.CLAVE",
			"TIPO_FUNCION_ID"
		from "FUNCION" "F"
			JOIN "SALA" "S" ON "F"."SALA_ID" = "S"."SALA_ID"
		where "PELICULA_ID" = $1
			and to_char("FECHA_INICIO", 'YYYY-MM-DD') = to_char(NOW(), 'YYYY-MM-DD')
		ORDER BY "FECHA_INICIO"
	`
	var funciones []domain.Funcion
	err := r.db.SelectContext(ctx, &funciones, query, peliculaId)
	if err != nil {
		return nil, domain.NewInternal()
	}
	return funciones, nil
}
