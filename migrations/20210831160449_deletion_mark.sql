-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE "instruction" ADD IF NOT EXISTS deleted boolean default false not null;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE "instruction" DROP IF EXISTS deleted;
-- +goose StatementEnd
