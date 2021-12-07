package application

import (
	"context"

	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/domain"
	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/ports"
)

func (s *service) CreateTicket(ctx context.Context, ticket *domain.Ticket) error {
	return s.r.Transaction(ctx, func(c context.Context, r ports.Repository) error {
		asientos := make([]*domain.AsignacionAsiento, len(ticket.Boletos))
		var err error
		// Se obtiene los asientos
		for i := range ticket.Boletos {
			asientos[i], err = r.GetAsientoByID(c, ticket.Boletos[i].AsientoId)
			if err != nil {
				return err
			}
		}
		// Se valida que la transaccion exista y aún este vigente
		err = r.ValidarTransaccion(c, ticket.TransaccionId)
		if err != nil {
			if e, ok := err.(*domain.Error); ok && e.Type != domain.NotFound {
				return err
			}
			// Si la transaccion no esta vigente, se valida que no se hayan modificado
			// los asientos para intentar completar la transaccion
			for i := range asientos {
				if asientos[i].StatusAsiento != string(domain.StatusDisponible) {
					return &domain.Error{
						Type:    domain.Conflict,
						Message: "La transaccion venció, y los asientos ya fueron seleccionados",
					}
				}
			}
		}
		// Validar que no se hayan modificado los asientos
		for i := range asientos {
			if asientos[i].StatusAsiento != string(domain.StatusDisponible) && asientos[i].StatusAsiento != string(domain.StatusEnCompra) {
				return &domain.Error{
					Type:    domain.Conflict,
					Message: "Los asientos ya no se encuntran disponibles",
				}
			}
		}
		// Update Transaccion
		err = s.r.UpdateTimeTransaction(c, ticket.TransaccionId)
		if err != nil {
			return err
		}
		// Create Ticket
		err = r.CreateTicket(c, ticket)
		if err != nil {
			return err
		}
		// Create Boleto
		for i := range ticket.Boletos {
			boletoId, err := r.CreateBoleto(c, ticket.Id, ticket.Boletos[i].TipoBoletoId)
			if err != nil {
				return err
			}
			asientos[i].BoletoId = &boletoId
		}
		// Actualizar Asignacion Asiento
		for i := range asientos {
			asientos[i].StatusAsiento = "4"
			err := r.UpdateBoletoIdAsiento(c, asientos[i])
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (s *service) IniciarCompra(ctx context.Context, transaccionId string, a []string) error {
	return s.r.Transaction(ctx, func(c context.Context, r ports.Repository) error {
		asientos := make([]*domain.AsignacionAsiento, len(a))
		var err error
		// Se obtiene los asientos
		for i := range a {
			asientos[i], err = r.GetAsientoByID(c, a[i])
			if err != nil {
				return err
			}
		}
		// Se valida que la transaccion exista y aún este vigente
		err = r.ValidarTransaccion(ctx, transaccionId)
		if err != nil {
			if e, _ := err.(*domain.Error); e.Type != domain.NotFound {
				return err
			}
			// Validar que no se hayan modificado los asientos
			for i := range a {
				if asientos[i].StatusAsiento != string(domain.StatusDisponible) {
					return &domain.Error{
						Type:    domain.Conflict,
						Message: "La transaccion venció, y los asientos ya fueron seleccionados",
					}
				}
			}
		}
		// Validar que no se hayan modificado los asientos
		for i := range asientos {
			if asientos[i].TransaccionId != transaccionId {
				if asientos[i].StatusAsiento != string(domain.StatusDisponible) {
					return &domain.Error{
						Type:    domain.Conflict,
						Message: "Los asientos ya no se encuntran disponibles",
					}
				}
			}
		}
		// Update Status AsignacionAsiento
		for i := range asientos {
			asientos[i].StatusAsiento = "3"
			err := r.UpdateStatusAsiento(c, asientos[i])
			if err != nil {
				return err
			}
		}
		// Update Transaccion
		return s.r.UpdateTimeTransaction(c, transaccionId)
	})
}
