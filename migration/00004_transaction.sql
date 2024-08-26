
-- +goose Up
CREATE TABLE transactions
(
    id uuid PRIMARY KEY,
    amount numeric,
    created_at timestamp,
    info text,
    type text,
    user_id uuid  REFERENCES users(id),
    created_by uuid NOT NULL,
    bound boolean NOT NULL DEFAULT false,
    deleted boolean NOT NULL DEFAULT false,
    deleted_at TIMESTAMP,
    deleted_by uuid,
    isolated_entity_id uuid NOT NULL
);


-- +goose Down
DROP TABLE IF EXISTS transactions;