package application

import (
	"context"

	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/domain"
	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/ports"
)

func (s *service) GetAsientosByFuncion(ctx context.Context, funcionId string) ([]domain.AsignacionAsiento, error) {
	return s.r.GetAsientosByFuncion(ctx, funcionId)
}

func (s *service) SeleccionarAsiento(ctx context.Context, asientoId string, transaccionId *string) error {
	return s.r.Transaction(ctx, func(c context.Context, r ports.Repository) error {
		// Se verifica que el asiento exista
		a, err := r.GetAsientoByID(ctx, asientoId)
		if err != nil {
			return err
		}
		// Se verifica que el asiento este disponible
		ok := r.DisponibilidadAsiento(ctx, a)
		if !ok {
			// revisar el formato del error
			return &domain.Error{
				Type:    domain.Conflict,
				Message: "El asiento ya no se encuentra disponible",
			}
		}
		// si transaccionId esta vacio, se crea una nueva transaccion, sino,
		// se verifica que la transaccion exista y además aún este vigente
		if *transaccionId == "" {
			*transaccionId, err = r.GetNewTransactionID(ctx)
			if err != nil {
				return err
			}
		} else {
			err = r.ValidarTransaccion(ctx, *transaccionId)
			if err != nil {
				return err
			}
		}

		// Se actualiza y transaccionId del asiento
		a.StatusID = "2" // 'SELECCIONADO'
		a.TransaccionId = *transaccionId

		// Se actualiza el estado del asiento en la BD
		err = r.UpdateStatusAsiento(ctx, a)
		if err != nil {
			return err
		}
		// Se actualiza el tiempo de la ultima modificación de la transacción
		return r.UpdateTimeTransaction(ctx, *transaccionId)
	})
}

func (s *service) DeseleccionarAsiento(ctx context.Context, asientoId, transaccionId string) error {
	// TODO Validar que la transaccion corresponde al usuario
	return s.r.Transaction(ctx, func(c context.Context, r ports.Repository) error {
		err := s.r.ValidarTransaccion(ctx, transaccionId)
		if err != nil {
			return err
		}
		a, err := r.GetAsientoByID(ctx, asientoId)
		if err != nil {
			return err
		}

		if transaccionId != a.TransaccionId {
			return &domain.Error{
				Type:    domain.NotFound,
				Message: "La transacción no corresponde asiento",
			}
		}

		a.StatusID = "1" // Disponible
		return r.UpdateStatusAsiento(ctx, a)
	})
}

func (s *service) DeshacerTransaccion(ctx context.Context, transaccionId string) error {
	return s.r.Transaction(ctx, func(c context.Context, r ports.Repository) error {
		return r.DeshacerTransaccion(ctx, transaccionId)
	})
}
