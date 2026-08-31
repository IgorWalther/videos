package stocks

import (
	"context"

	sqlcstocks "github.com/igoroutine-courses/microservices.ecommerce.loms/internal/repository/stocks/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type postgresRepository struct {
	queries *sqlcstocks.Queries
}

func NewPostgresRepository(db sqlcstocks.DBTX) *postgresRepository {
	return &postgresRepository{
		queries: sqlcstocks.New(db),
	}
}

// NOTE: специально разный нейминг

func (r *postgresRepository) ReserveStock(ctx context.Context, sku uint32, count uint64) error {
	_, err := r.queries.DecrementAvailableStock(ctx, sqlcstocks.DecrementAvailableStockParams{
		Sku:    int32(sku),
		Amount: int32(count),
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *postgresRepository) ReleaseStock(ctx context.Context, sku uint32, count uint64) error {
	return r.queries.AddToAvailableStock(ctx, sqlcstocks.AddToAvailableStockParams{
		Sku:    int32(sku),
		Amount: int32(count),
	})
}

func (r *postgresRepository) SetStock(ctx context.Context, sku uint32, count uint64) error {
	return r.queries.UpsertAvailableStock(ctx, sqlcstocks.UpsertAvailableStockParams{
		Sku:    int32(sku),
		Amount: int32(count),
	})
}

// NOTE: специально разный нейминг

func (r *postgresRepository) GetStock(ctx context.Context, sku uint32) (uint64, error) {
	amount, err := r.queries.GetAvailableStockAmount(ctx, int32(sku))

	if err != nil {
		return 0, err
	}

	return uint64(amount), nil
}

func pgInt4(v int32) pgtype.Int4 {
	return pgtype.Int4{
		Int32: v,
		Valid: true,
	}
}
