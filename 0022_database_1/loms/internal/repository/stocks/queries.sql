-- name: DecrementAvailableStock :execrows
UPDATE loms.available_stocks
SET amount = amount - $2
WHERE sku = $1
  AND amount >= $2;

-- name: AddToAvailableStock :exec
INSERT INTO loms.available_stocks AS a (sku, amount)
VALUES ($1, $2)
ON CONFLICT (sku) DO UPDATE
    SET amount = a.amount + EXCLUDED.amount;

-- name: GetAvailableStockAmount :one
SELECT COALESCE(
               (SELECT amount
                FROM loms.available_stocks
                WHERE sku = $1),
               0
       )::integer AS amount;

-- name: UpsertAvailableStock :exec
INSERT INTO loms.available_stocks (sku, amount)
VALUES ($1, $2)
ON CONFLICT (sku) DO UPDATE SET amount = EXCLUDED.amount;
