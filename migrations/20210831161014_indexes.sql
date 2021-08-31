-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
create unique index idx_instruction_instruction_id on instruction (instruction_id);
create index idx_instruction_instruction_id_deleted on instruction (instruction_id, deleted);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop index idx_instruction_instruction_id_deleted;
-- +goose StatementEnd
