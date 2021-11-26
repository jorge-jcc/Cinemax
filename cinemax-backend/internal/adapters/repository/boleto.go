package repository

import (
	"context"

	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/domain"
)

func (r *repository) GetPreciosBoletos(ctx context.Context) {
	query := `
	SELECT "CLAVE", "PRECIO", "SALA_ID", 
		CASE
			WHEN "AS"."STATUS_ASIENTO_ID" = 1 OR (EXTRACT(EPOCH FROM NOW() - "T"."UPDATED_AT")/60 >= 1 AND "AS"."STATUS_ASIENTO_ID" IN (2, 3)) THEN 'DISPONIBLE'
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
	r.db.SelectContext(ctx, &asientos, query)
}
