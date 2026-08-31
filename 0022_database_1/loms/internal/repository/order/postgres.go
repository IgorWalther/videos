package order

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	entitypkg "github.com/igoroutine-courses/microservices.ecommerce.loms/internal/entity"
	sqlcorder "github.com/igoroutine-courses/microservices.ecommerce.loms/internal/repository/order/sqlc"
)

type postgresRepository struct {
	queries *sqlcorder.Queries
}

func NewPostgresRepository(qdb sqlcorder.DBTX) *postgresRepository {
	return &postgresRepository{
		queries: sqlcorder.New(qdb),
	}
}

// TODO: transactions!

func (r *postgresRepository) CreateOrder(
	ctx context.Context,
	o *entitypkg.Order,
) (int64, error) {
	orderRow, err := r.queries.InsertOrder(ctx, sqlcorder.InsertOrderParams{
		ID:     o.ID,
		UserID: o.UserID,
		Status: toSQLCStatus(o.Status),
	})

	if err != nil {
		return 0, err
	}

	for _, item := range o.Items {
		err = r.queries.InsertOrderItem(ctx, sqlcorder.InsertOrderItemParams{
			OrderID: orderRow.ID,
			Sku:     int32(item.SKU),
			Amount:  int32(item.Count),
		})

		if err != nil {
			return 0, err
		}
	}

	return orderRow.ID, nil
}

func (r *postgresRepository) GetOrder(ctx context.Context, orderID int64) (*entitypkg.Order, error) {
	orderRow, err := r.queries.GetOrderByID(ctx, orderID)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, entitypkg.ErrOrderNotFound
		}

		return nil, err
	}

	itemRows, err := r.queries.ListOrderItemsByOrderID(ctx, orderID)

	if err != nil {
		return nil, err
	}

	items := make([]entitypkg.OrderItem, 0, len(itemRows))
	for _, item := range itemRows {
		items = append(items, entitypkg.OrderItem{
			SKU:   uint32(item.Sku),
			Count: uint32(item.Amount),
		})
	}

	order := &entitypkg.Order{
		ID:        orderRow.ID,
		UserID:    orderRow.UserID,
		Status:    toEntityStatus(orderRow.Status),
		Items:     items,
		CreatedAt: orderRow.CreatedAt.Time,
		UpdatedAt: orderRow.UpdatedAt.Time,
	}

	return order, nil
}

func (r *postgresRepository) SetOrderStatus(
	ctx context.Context,
	orderID int64,
	status entitypkg.OrderStatus,
) error {
	return r.queries.SetOrderStatus(ctx, sqlcorder.SetOrderStatusParams{
		ID:        orderID,
		Status:    toSQLCStatus(status),
		UpdatedAt: pgTimestamptz(time.Now()),
	})
}

func pgInt8(v int64) pgtype.Int8 {
	return pgtype.Int8{
		Int64: v,
		Valid: true,
	}
}

func pgInt4(v int32) pgtype.Int4 {
	return pgtype.Int4{
		Int32: v,
		Valid: true,
	}
}

func pgTimestamptz(t time.Time) pgtype.Timestamptz {
	return pgtype.Timestamptz{
		Time:  t,
		Valid: true,
	}
}

func toSQLCStatus(status entitypkg.OrderStatus) sqlcorder.LomsOrderStatus {
	switch status {
	case entitypkg.OrderStatusNew:
		return sqlcorder.LomsOrderStatusNew
	case entitypkg.OrderStatusAwaitingPayment:
		return sqlcorder.LomsOrderStatusAwaitingpayment
	case entitypkg.OrderStatusFailed:
		return sqlcorder.LomsOrderStatusFailed
	case entitypkg.OrderStatusPaid:
		return sqlcorder.LomsOrderStatusPaid
	case entitypkg.OrderStatusCancelled:
		return sqlcorder.LomsOrderStatusCancelled
	default:
		return sqlcorder.LomsOrderStatusNew
	}
}

func toEntityStatus(status sqlcorder.LomsOrderStatus) entitypkg.OrderStatus {
	switch status {
	case sqlcorder.LomsOrderStatusNew:
		return entitypkg.OrderStatusNew
	case sqlcorder.LomsOrderStatusAwaitingpayment:
		return entitypkg.OrderStatusAwaitingPayment
	case sqlcorder.LomsOrderStatusFailed:
		return entitypkg.OrderStatusFailed
	case sqlcorder.LomsOrderStatusPaid:
		return entitypkg.OrderStatusPaid
	case sqlcorder.LomsOrderStatusCancelled:
		return entitypkg.OrderStatusCancelled
	default:
		return entitypkg.OrderStatusNew
	}
}
