INSERT INTO loms.orders (id, user_id, status)
VALUES (NULL, NULL, 'new');

INSERT INTO loms.order_info (order_id, sku, amount)
VALUES (1001, 101, 2),
       (1001, 102, 1),
       (1002, 101, 3),
       (1003, 103, 5),
       (1004, 104, 1);

INSERT INTO loms.available_stocks (sku, amount)
VALUES (101, 50),
       (102, 20),
       (103, 100),
       (104, 10);

INSERT INTO loms.reserved_stocks (sku, order_id, amount)
VALUES (101, 1001, 2),
       (102, 1001, 1),
       (101, 1002, 3),
       (103, 1003, 5);
