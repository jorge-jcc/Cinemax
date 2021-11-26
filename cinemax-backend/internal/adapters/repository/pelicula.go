package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/domain"
	"github.com/lib/pq"
)

func (r *repository) CreatePelicula(ctx context.Context, p *domain.Pelicula) error {
	query := `
	INSERT INTO "PELICULA" ("NOMBRE", "DIRECTOR", "DESCRIPCION", "DURACION_MINUTOS", 
		"ANIO", "FECHA_DISPONIBILIDAD", "RESENA", "CLASIFICACION_ID", 
		"IDIOMA_ID", "SUBTITULO_ID", "GENERO_ID") 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`
	_, err := r.db.ExecContext(ctx, query,
		p.Nombre, p.Director, p.Descripcion, p.DuracionMinutos,
		p.Anio, p.FechaDisponiblidad, p.Resena, 1, 1, 1, 1,
	)
	if err != nil {
		if _, ok := err.(*pq.Error); ok {
			return err
		}
		return domain.NewInternal()
	}
	return nil
}

func (r *repository) UpdateImage(ctx context.Context, id, imagen string) error {
	query := `UPDATE "PELICULA" SET "IMAGEN" = $1 WHERE "PELICULA_ID" = $2`
	result, err := r.db.ExecContext(ctx, query, imagen, id)
	if err != nil {
		log.Printf("Error updating image_url in database: %v\n", err)
		return domain.NewInternal()
	}
	if rows, _ := result.RowsAffected(); rows == 0 {
		log.Printf("Error, imagen_id not found in database: %v\n", err)
		return domain.NewNotFound("imagenId", id)
	}
	return nil
}

func (r *repository) GetPeliculaById(ctx context.Context, id string) (*domain.Pelicula, error) {
	query := `
		SELECT "PELICULA_ID", "NOMBRE", "DIRECTOR", "DESCRIPCION", "DURACION_MINUTOS", "ANIO", 
			"FECHA_DISPONIBILIDAD", "RESENA",
		CASE
			WHEN "IMAGEN" is NULL THEN ''
			ELSE "IMAGEN"
		END AS "IMAGEN"
		FROM "PELICULA" WHERE "PELICULA_ID" = $1
	`
	p := &domain.Pelicula{}
	err := r.db.GetContext(ctx, p, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("Error pelicula not found in database: %v\n", err)
			return nil, domain.NewNotFound("pelicula_id", id)
		}
		return nil, domain.NewInternal()
	}
	return p, nil
}

func (r *repository) GetPeliculasByNombre(ctx context.Context, nombre string, limit, offset int16) ([]domain.Pelicula, error) {
	query := `
		SELECT "P"."NOMBRE", "P"."DIRECTOR", "P"."DESCRIPCION", "P"."DURACION_MINUTOS",
			"P"."ANIO", "P"."FECHA_DISPONIBILIDAD", "P"."RESENA",
			"I"."NOMBRE" AS "IDIOMA.NOMBRE", "S"."NOMBRE" AS "SUBTITULO.NOMBRE",
			"G"."NOMBRE" AS "GENERO.NOMBRE", "C"."CLAVE" AS "CLASIFICACION.CLAVE",
			"C"."DESCRIPCION" AS "CLASIFICACION.DESCRIPCION"
		FROM "PELICULA" AS "P", "CLASIFICACION" AS "C", "IDIOMA" AS "I", "IDIOMA" AS "S",
			"GENERO" AS "G"
		WHERE "P"."CLASIFICACION_ID" = "C"."CLASIFICACION_ID"
			AND "P"."IDIOMA_ID" = "I"."IDIOMA_ID"
			AND "P"."SUBTITULO_ID" = "S"."IDIOMA_ID"
			AND "P"."GENERO_ID" = "G"."GENERO_ID"
			AND LOWER("P"."NOMBRE") LIKE LOWER($1) LIMIT $2 OFFSET $3
	`
	var peliculas []domain.Pelicula
	err := r.db.SelectContext(ctx, &peliculas, query, "%"+nombre+"%", 20, 0)
	if err != nil {
		return nil, domain.NewInternal()
	}
	return peliculas, nil
}

func (r *repository) GetPeliculasEnCartelera(ctx context.Context) ([]domain.Pelicula, error) {
	// TO_CHAR("F"."FECHA_INICIO", 'YYYY-MM-DD') = TO_CHAR(NOW(), 'YYYY-MM-DD')
	query := `
		SELECT DISTINCT "P"."PELICULA_ID", "P"."NOMBRE" 
		FROM "PELICULA" AS "P"
			JOIN "FUNCION" AS "F" ON "P"."PELICULA_ID" = "F"."PELICULA_ID"
		WHERE TO_CHAR("F"."FECHA_INICIO", 'YYYY-MM-DD') = TO_CHAR(NOW(), 'YYYY-MM-DD')
	`
	var peliculas []domain.Pelicula
	err := r.db.SelectContext(ctx, &peliculas, query)
	if err != nil {
		return nil, domain.NewInternal()
	}
	return peliculas, nil
}

func (r *repository) GetClasificaciones(ctx context.Context) ([]domain.Clasificacion, error) {
	query := `SELECT "CLASIFICACION_ID", "CLAVE", "DESCRIPCION" FROM "CLASIFICACION"`
	var clasificaciones []domain.Clasificacion
	err := r.db.SelectContext(ctx, &clasificaciones, query)
	if err != nil {
		return nil, domain.NewInternal()
	}
	return clasificaciones, nil
}

func (r *repository) GetIdiomas(ctx context.Context) ([]domain.Idioma, error) {
	query := `SELECT "IDIOMA_ID", "NOMBRE" FROM "IDIOMA"`
	var idiomas []domain.Idioma
	err := r.db.SelectContext(ctx, &idiomas, query)
	if err != nil {
		return nil, domain.NewInternal()
	}
	return idiomas, nil
}

func (r *repository) GetGeneros(ctx context.Context) ([]domain.Genero, error) {
	query := `SELECT "GENERO_ID", "NOMBRE" FROM "GENERO"`
	var generos []domain.Genero
	err := r.db.SelectContext(ctx, &generos, query)
	if err != nil {
		return nil, domain.NewInternal()
	}
	return generos, nil
}
