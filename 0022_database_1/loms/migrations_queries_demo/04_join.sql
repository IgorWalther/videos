SELECT o.id, o.user_id, o.status, oi.sku, oi.amount
FROM loms.orders o
         INNER JOIN loms.order_info oi ON o.id = oi.order_id;
