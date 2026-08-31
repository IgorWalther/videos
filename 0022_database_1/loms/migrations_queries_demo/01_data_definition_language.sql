CREATE TABLE demo_users
(
    id   BIGINT PRIMARY KEY,
    name TEXT NOT NULL
);

ALTER TABLE demo_users
    ADD COLUMN email TEXT;

ALTER TABLE demo_users
    RENAME COLUMN name TO full_name;

ALTER TABLE demo_users
    RENAME COLUMN name TO full_name;
