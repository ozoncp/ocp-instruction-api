-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE instruction (
    id int NOT NULL,
    instruction_id int,
    classroom_id int,
    text text,
    prev_id int,
    PRIMARY KEY(id)
);
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop table instruction;
-- +goose StatementEnd
