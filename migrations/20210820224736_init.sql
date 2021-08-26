-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE instruction (
    id SERIAL NOT NULL PRIMARY KEY,
    instruction_id int,
    classroom_id int,
    text text,
    prev_id int
);
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop table instruction;
-- +goose StatementEnd
