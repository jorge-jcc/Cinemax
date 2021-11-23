package application

import (
	"context"
	"fmt"

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
			fmt.Println("Aqui", err)
			return err
		}
		// Se verifica que el asiento este disponible
		ok := r.DisponibilidadAsiento(ctx, a)
		if !ok {
			// revisar el formato del error
			return domain.NewConflict("asignacion_asiento_id", asientoId)
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
