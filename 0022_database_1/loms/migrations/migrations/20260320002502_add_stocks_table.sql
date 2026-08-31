-- +goose Up
CREATE TABLE loms.available_stocks
(
    sku    INTEGER NOT NULL PRIMARY KEY,
    amount INTEGER NOT NULL CHECK (amount > 0)
);

CREATE TABLE loms.reserved_stocks
(
    sku      INTEGER NOT NULL,
    order_id BIGINT  NOT NULL REFERENCES loms.orders (id) ON DELETE CASCADE,
    amount   INTEGER NOT NULL CHECK (amount > 0),

    PRIMARY KEY (sku, order_id)
);

CREATE UNIQUE INDEX idx_reserved_stocks_order_id ON loms.reserved_stocks USING BTREE (order_id);

-- +goose Down
DROP INDEX IF EXISTS loms.idx_reserved_stocks_order_id;
DROP TABLE IF EXISTS loms.reserved_stocks;
DROP TABLE IF EXISTS loms.available_stocks;
