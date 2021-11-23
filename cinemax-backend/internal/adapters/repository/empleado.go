package repository

import (
	"context"
	"time"

	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/domain"
	"github.com/lib/pq"
)

type empleado struct {
	domain.Empleado
	Create_at time.Time `db:"CREATED_AT"`
	Update_at time.Time `db:"UPDATED_AT"`
}

// CreateEmpleado crea un empleado en la base de datos, y le asigna su id
func (r *repository) CreateEmpleado(ctx context.Context, e *domain.Empleado) error {
	query := `
		INSERT INTO "EMPLEADO"("NOMBRE", "AP_PATERNO", "AP_MATERNO", "RFC", "EDAD", "EMAIL", 
			"DIRECCION", "TELEFONO", "PASSWORD")
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING "EMPLEADO_ID"
	`
	err := r.db.GetContext(ctx, e, query,
		e.Nombre, e.ApPaterno, e.ApMaterno, e.Rfc, e.Edad, e.Email, e.Direccion,
		e.Telefono, e.Password,
	)
	if pqErr, ok := err.(*pq.Error); ok {
		if pqErr.Code.Name() == "unique_violation" {
			switch pqErr.Constraint {
			case "EMPLEADO_EMAIL_UK":
				return domain.NewConflict("email", e.Email)
			case "EMPLEADO_TELEFONO_UK":
				return domain.NewConflict("telefono", e.Telefono)
			}
		}
		return domain.NewInternal()
	}
	return nil
}

func (r *repository) FindEmpleadoByEmail(ctx context.Context, email string) (*domain.Empleado, error) {
	query := `SELECT * FROM "EMPLEADO" WHERE "EMAIL" = $1`
	e := &empleado{}
	err := r.db.GetContext(ctx, e, query, email)
	if err != nil {
		return nil, domain.NewNotFound("email", email)
	}
	return &e.Empleado, nil
}
