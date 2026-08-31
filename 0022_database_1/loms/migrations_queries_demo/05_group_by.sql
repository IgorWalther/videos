SELECT user_id, COUNT(*) AS orders_count
FROM loms.orders
GROUP BY user_id;

SELECT sku, SUM(amount) AS total_ordered
FROM loms.order_info
GROUP BY sku;

SELECT status, COUNT(*) AS cnt
FROM loms.orders
GROUP BY status;
