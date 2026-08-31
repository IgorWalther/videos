-- name: InsertOrder :one
INSERT INTO loms.orders (id, user_id, status)
VALUES (sqlc.arg(id), sqlc.arg(user_id), sqlc.arg(status)::loms.order_status)
RETURNING id, user_id, status, created_at, updated_at;

-- name: InsertOrderItem :exec
INSERT INTO loms.order_info (order_id, sku, amount)
VALUES ($1, $2, $3);

-- name: GetOrderByID :one
SELECT id, user_id, status, created_at, updated_at
FROM loms.orders
WHERE id = $1;

-- name: ListOrderItemsByOrderID :many
SELECT sku, amount
FROM loms.order_info
WHERE order_id = $1
ORDER BY sku;

-- name: SetOrderStatus :exec
UPDATE loms.orders
SET status     = sqlc.arg(status)::loms.order_status,
    updated_at = sqlc.arg(updated_at)
WHERE id = sqlc.arg(id);
