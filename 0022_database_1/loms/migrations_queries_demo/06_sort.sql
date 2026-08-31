SELECT id, user_id, status, created_at
FROM loms.orders
ORDER BY created_at DESC;

SELECT id, user_id, status
FROM loms.orders
ORDER BY user_id ASC, status ASC;
