
-- +goose Up
CREATE TABLE examples
(
    id uuid PRIMARY KEY,
    cpu_units numeric,
    memory_mb integer,
    name text,
    price numeric,
    storage_mb integer,
    user_id uuid  REFERENCES users(id),
    created_by uuid NOT NULL,
    bound boolean NOT NULL DEFAULT false,
    deleted boolean NOT NULL DEFAULT false,
    deleted_at TIMESTAMP,
    deleted_by uuid,
    isolated_entity_id uuid NOT NULL
);


-- +goose Down
DROP TABLE IF EXISTS examples;