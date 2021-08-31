-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

drop index idx_instruction_instruction_id;
drop index idx_instruction_instruction_id_deleted;

CREATE TABLE instruction_tmp (
 id SERIAL NOT NULL PRIMARY KEY,
 instruction_id int,
 classroom_id int,
 text text,
 prev_id int,
 deleted boolean default false not null
) partition by HASH (id);

create unique index idx_instruction_instruction_id on instruction_tmp (id, instruction_id);
create index idx_instruction_instruction_id_deleted on instruction_tmp (instruction_id, deleted);

create table "instruction_1" partition of "instruction_tmp" FOR VALUES WITH (MODULUS 2, REMAINDER 0);
create table "instruction_2" partition of "instruction_tmp" FOR VALUES WITH (MODULUS 2, REMAINDER 1);

insert into instruction_tmp select * from instruction;
drop table instruction;
alter table instruction_tmp rename to instruction;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
