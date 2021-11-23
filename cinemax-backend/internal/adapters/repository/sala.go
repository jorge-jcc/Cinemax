package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/domain"
)

func (r *repository) GetSalas(ctx context.Context) ([]domain.Sala, error) {
	query := `
		SELECT "S"."SALA_ID", "S"."NOMBRE", "S"."DESCRIPCION",
			"TS"."TIPO_SALA_ID" AS "TIPO_SALA.TIPO_SALA_ID", 
			"TS"."CLAVE" AS "TIPO_SALA.CLAVE",
			"TS"."DESCRIPCION" AS "TIPO_SALA.DESCRIPCION"
		FROM "SALA" AS "S", "TIPO_SALA" AS "TS"
		WHERE "S"."TIPO_SALA_ID" = "TS"."TIPO_SALA_ID"
	`
	var salas []domain.Sala
	err := r.db.SelectContext(ctx, &salas, query)
	if err != nil {
		return nil, domain.NewInternal()
	}
	return salas, nil
}

func (r *repository) GetSalaById(ctx context.Context, id string) (*domain.Sala, error) {
	query := `
		SELECT "S"."SALA_ID", "S"."NOMBRE", "S"."DESCRIPCION",
			"TS"."TIPO_SALA_ID" AS "TIPO_SALA.TIPO_SALA_ID", 
			"TS"."CLAVE" AS "TIPO_SALA.CLAVE",
			"TS"."DESCRIPCION" AS "TIPO_SALA.DESCRIPCION"
		FROM "SALA" AS "S", "TIPO_SALA" AS "TS"
		WHERE "S"."TIPO_SALA_ID" = "TS"."TIPO_SALA_ID"
			AND "S"."SALA_ID" = $1
	`
	sala := &domain.Sala{}
	err := r.db.GetContext(ctx, sala, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("Error sala not found in database: %v\n", err)
			return nil, domain.NewNotFound("sala_id", id)
		}
		return nil, domain.NewInternal()
	}
	return sala, nil
}
