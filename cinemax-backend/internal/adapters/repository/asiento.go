package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/domain"
	"github.com/lib/pq"
)

func (r *repository) InicializarAsientos(ctx context.Context, f *domain.Funcion) error {
	query := `
		INSERT INTO "ASIGNACION_ASIENTO" ("ASIENTO_ID", "FUNCION_ID", "STATUS_ASIENTO_ID")
			SELECT "ASIENTO_ID", $1 AS "FUNCION_ID", 1 AS "STATUS_ASIENTO_ID"
			FROM "ASIENTO" WHERE "SALA_ID" = $2
	`
	_, err := r.db.ExecContext(ctx, query, f.ID, f.Sala.ID)
	if pqErr, ok := err.(*pq.Error); ok {
		if pqErr.Code.Name() == "foreign_key_violation" {
			return domain.NewNotFound("sala_id", f.Sala.ID)
		}
		return domain.NewInternal()
	}
	return nil
}

func (r *repository) GetAsientosByFuncion(ctx context.Context, funcionId string) ([]domain.AsignacionAsiento, error) {
	query := `
		SELECT "ASIGNACION_ASIENTO_ID", "CLAVE", "FUNCION_ID", 
			CASE
				WHEN "AS"."STATUS_ASIENTO_ID" = 1 OR (TO_NUMBER(TO_CHAR(NOW() - "T"."UPDATED_AT", 'MI'), '99') >= 1 AND "AS"."STATUS_ASIENTO_ID" IN (2, 3)) THEN 'DISPONIBLE'
				WHEN "AS"."STATUS_ASIENTO_ID" IN (2, 3) THEN 'EN PROCESO'
				WHEN "AS"."STATUS_ASIENTO_ID" = 4 THEN 'ASIGNADO'
				ELSE 'DESCONOCIDO'
			END
			AS "STATUS"
		FROM "ASIGNACION_ASIENTO" AS "AS" 
			JOIN "ASIENTO" AS "A" ON "AS"."ASIENTO_ID" = "A"."ASIENTO_ID" 
			LEFT JOIN "TRANSACCION" AS "T" ON "AS"."TRANSACCION_ID" = "T"."TRANSACCION_ID"
		WHERE "FUNCION_ID" = $1
		ORDER BY SUBSTRING("CLAVE", 1, 1), SUBSTRING("CLAVE" FROM '([0-9]+)$')::INT 
	`
	var asientos []domain.AsignacionAsiento
	err := r.db.SelectContext(ctx, &asientos, query, funcionId)
	fmt.Println(err, len(asientos))
	if err != nil {
		return nil, domain.NewInternal()
	}
	return asientos, nil
}

func (r *repository) GetAsientoByID(ctx context.Context, asientoId string) (*domain.AsignacionAsiento, error) {
	query := `
		SELECT "ASIGNACION_ASIENTO_ID", "FUNCION_ID", "UPDATED_AT"
		FROM "ASIGNACION_ASIENTO"
		WHERE "ASIGNACION_ASIENTO_ID" = $1
	`
	a := &domain.AsignacionAsiento{}
	err := r.db.GetContext(ctx, a, query, asientoId)
	fmt.Println(err)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("Error asiento not found in database: %v\n", err)
			return nil, domain.NewNotFound("asignacion_asiento_id", asientoId)
		}
		return nil, domain.NewInternal()
	}
	return a, nil
}

func (r *repository) DisponibilidadAsiento(ctx context.Context, a *domain.AsignacionAsiento) bool {
	query := `
		SELECT COUNT(*) VALUE FROM "ASIGNACION_ASIENTO" AS "AS"
			LEFT JOIN "TRANSACCION" AS "T" ON "AS"."TRANSACCION_ID" = "T"."TRANSACCION_ID"
		WHERE "ASIGNACION_ASIENTO_ID" = $1
			AND "FUNCION_ID" = $2
			AND (
				"STATUS_ASIENTO_ID" = 1 OR (
					TO_NUMBER(TO_CHAR(NOW() - "T"."UPDATED_AT", 'MI'), '99') >= 1 AND "STATUS_ASIENTO_ID" IN (2, 3)
				)
			)
	`
	var result int
	err := r.db.GetContext(ctx, &result, query, a.ID, a.FuncionID)
	if err != nil {
		return false
	}
	return result == 1
}

// Actualiza el status de un asiento
func (r *repository) UpdateStatusAsiento(ctx context.Context, a *domain.AsignacionAsiento) error {
	query := `
		UPDATE "ASIGNACION_ASIENTO" SET 
				"STATUS_ASIENTO_ID" = $1, 
				"TRANSACCION_ID" = $2, 
				"UPDATED_AT" = NOW()
		WHERE "ASIGNACION_ASIENTO_ID" = $3 AND "UPDATED_AT" = $4
	`
	result, err := r.db.ExecContext(ctx, query, a.StatusID, a.TransaccionId, a.ID, a.UpdatedAt)
	fmt.Println(err)
	if err != nil {
		return domain.NewInternal()
	}
	if rows, e := result.RowsAffected(); rows != 1 {
		fmt.Println("hOLA", rows, e)
		return domain.NewNotFound("asignacion_asiento_id", a.ID)
	}
	return nil
}

// UpdateTimeTransaction actualiza de la ultima modificacion de una transaccion
func (r *repository) UpdateTimeTransaction(ctx context.Context, transaccionId string) error {
	query := `
		UPDATE "TRANSACCION" SET "UPDATED_AT" = NOW()
		WHERE "TRANSACCION_ID" = $1
	`
	result, err := r.db.ExecContext(ctx, query, transaccionId)
	fmt.Println(result, err)
	if err != nil {
		return domain.NewInternal()
	}
	if rows, _ := result.RowsAffected(); rows != 1 {
		return domain.NewNotFound("transaccion_id", transaccionId)
	}
	return nil
}

// Devuleve el ID de una nueva transaccion
func (r *repository) GetNewTransactionID(ctx context.Context) (string, error) {
	query := `INSERT INTO "TRANSACCION" ("CREATED_AT") VALUES(NOW()) RETURNING "TRANSACCION_ID"`
	var result string
	err := r.db.GetContext(ctx, &result, query)
	if err != nil {
		return "", err
	}
	return result, nil
}

// ValidarTransaccion devulve un error si la transacción no existe o ya expiro
func (r *repository) ValidarTransaccion(ctx context.Context, transaccionId string) error {
	query := `
		SELECT TO_NUMBER(TO_CHAR(NOW() - "UPDATED_AT", 'MI'), '99') 
		FROM "TRANSACCION" 
		WHERE "TRANSACCION_ID" = $1
	`
	var result int
	err := r.db.GetContext(ctx, &result, query, transaccionId)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.NewNotFound("transaccionId", transaccionId)
		}
		return domain.NewInternal()
	}
	if result >= 1 {
		return &domain.Error{
			Type:    domain.Conflict,
			Message: fmt.Sprintf("la transacción %v dejo de ser válida", transaccionId),
		}
	}
	return nil
}
