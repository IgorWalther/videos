SELECT *
FROM loms.orders;

INSERT INTO loms.orders (id, user_id, status)
VALUES (1005, 4, 'new');

UPDATE loms.orders
SET status     = 'paid',
    updated_at = NOW()
WHERE 1 = 1;

DELETE
FROM loms.orders
WHERE id = 1004;

SELECT id, user_id, status, created_at
FROM loms.orders
WHERE user_id = 1;

SELECT order_id, sku, amount
FROM loms.order_info
WHERE order_id = 1001;

UPDATE loms.orders
SET status     = 'failed',
    updated_at = NOW()
WHERE id = 1001;

UPDATE loms.order_info
SET amount = 10
WHERE order_id = 1003
  AND sku = 103;
