-- +goose Up  
;
-- +goose Down

DROP TABLE IF EXISTS deploys;
DROP TABLE IF EXISTS examples;
DROP TABLE IF EXISTS transactions;
DROP TABLE IF EXISTS users;