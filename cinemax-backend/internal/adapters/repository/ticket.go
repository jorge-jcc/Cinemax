package repository

import (
	"context"
	"fmt"

	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/domain"
)

func (r *repository) CreateTicket(ctx context.Context, ticket *domain.Ticket) error {
	query := `
		INSERT INTO "TICKET" ("FECHA_COMPRA", "MONTO", "FUNCION_ID", "EMPLEADO_ID")
		VALUES	(NOW(), $1, $2, $3)
		RETURNING "TICKET_ID"
	`
	var ticketId int64
	err := r.db.GetContext(ctx, &ticketId, query, ticket.Monto, ticket.FuncionId, ticket.EmpleadoId)
	if err != nil {
		return domain.NewInternal()
	}
	ticket.Id = fmt.Sprint(ticketId)
	return nil
}
