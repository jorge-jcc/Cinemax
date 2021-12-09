package repository

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/domain"
)

func (r *repository) GetSalasDisponibles(ctx context.Context, fechaInicio, fechaFin time.Time) ([]domain.Sala, error) {
	query := `
		SELECT "S"."SALA_ID", "S"."CLAVE", "S"."NOMBRE"
		FROM "SALA" AS "S"
		EXCEPT
		SELECT "S"."SALA_ID", "S"."CLAVE", "S"."NOMBRE"
		FROM "SALA" AS "S"
			JOIN "FUNCION" AS "F" ON "S"."SALA_ID"  = "F"."SALA_ID"
		WHERE TO_CHAR("F"."FECHA_INICIO", 'DD-MM-YYYY') = TO_CHAR($1::date, 'DD-MM-YYYY')
			AND (
				($1 >= "FECHA_INICIO" AND $2 <= "FECHA_FIN")
				OR
				($2 >= "FECHA_INICIO" AND $2 <= "FECHA_FIN")
			)
		ORDER BY "CLAVE"
	`
	var salas []domain.Sala
	err := r.db.SelectContext(ctx, &salas, query, fechaInicio, fechaFin)
	if err != nil {
		return nil, domain.NewInternal()
	}
	return salas, nil
}

func (r *repository) GetSalaById(ctx context.Context, id string) (*domain.Sala, error) {
	query := `
		SELECT "SALA_ID", "CLAVE", "NOMBRE", "UBICACION"
		FROM "SALA"
		WHERE "SALA_ID" = $1
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

func (r *repository) GetSalaByFuncionId(ctx context.Context, funcionId string) (*domain.Sala, error) {
	query := `
	SELECT "SALA_ID", "CLAVE", "NOMBRE", "UBICACION"
	FROM "SALA"
	WHERE "SALA_ID" = (SELECT "SALA_ID" FROM "FUNCION" WHERE "FUNCION_ID" = $1)
	`
	sala := &domain.Sala{}
	err := r.db.GetContext(ctx, sala, query, funcionId)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("Error sala not found in database: %v\n", err)
			return nil, domain.NewNotFound("funcionId", funcionId)
		}
		return nil, domain.NewInternal()
	}
	return sala, nil
}
