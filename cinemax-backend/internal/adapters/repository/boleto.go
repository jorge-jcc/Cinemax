package repository

import (
	"context"
	"fmt"

	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/domain"
)

func (r *repository) GetPreciosBoletos(ctx context.Context) ([]domain.PrecioBoleto, error) {
	query := `
	SELECT "T"."TIPO_BOLETO_ID", "CLAVE", "PRECIO", "TIPO_FUNCION_ID"
	FROM "TIPO_BOLETO" AS "T"
		JOIN "PRECIO_BOLETO" AS "P" ON "T"."TIPO_BOLETO_ID" = "P"."TIPO_BOLETO_ID"
`
	var precios []domain.PrecioBoleto
	err := r.db.SelectContext(ctx, &precios, query)
	if err != nil {
		return nil, domain.NewInternal()
	}
	return precios, nil
}

func (r *repository) CreateBoleto(ctx context.Context, ticketId, tipoBoletoId string) (string, error) {
	query := `
		INSERT INTO "BOLETO" ("TICKET_ID", "TIPO_BOLETO_ID") VALUES
		($1, $2) RETURNING "BOLETO_ID"
	`
	var boletoId string
	fmt.Println(ticketId, tipoBoletoId)
	err := r.db.GetContext(ctx, &boletoId, query, ticketId, tipoBoletoId)
	fmt.Println(err)
	if err != nil {
		return "", domain.NewInternal()
	}
	return boletoId, err
}
