-- +goose Up
CREATE SCHEMA IF NOT EXISTS loms;

CREATE TYPE loms.order_status AS ENUM ('new', 'awaiting payment', 'failed', 'paid', 'cancelled');

CREATE TABLE IF NOT EXISTS loms.orders
(
    id         BIGINT            NOT NULL PRIMARY KEY,
    user_id    BIGINT            NOT NULL,
    status     loms.order_status NOT NULL,
    created_at TIMESTAMPTZ       NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ       NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_orders_user_id ON loms.orders USING BTREE (user_id, status);

CREATE TABLE IF NOT EXISTS loms.order_info
(
    order_id BIGINT  NOT NULL REFERENCES loms.orders (id) ON DELETE CASCADE,
    sku      INTEGER NOT NULL,
    amount   INTEGER NOT NULL CHECK (amount > 0),

    PRIMARY KEY (order_id, sku)
);

CREATE INDEX IF NOT EXISTS idx_order_info_sku ON loms.order_info USING BTREE (sku);

-- +goose Down
DROP INDEX IF EXISTS loms.idx_order_info_sku;
DROP TABLE IF EXISTS loms.order_info;

DROP INDEX IF EXISTS loms.idx_orders_user_id;
DROP TABLE IF EXISTS loms.orders;
DROP TYPE IF EXISTS loms.order_status;
