drop table if exists tasks;
create table tasks (
  id serial primary key,
  name varchar(255),
  status integer
);

insert into tasks (id, name, status) values
  (1, 'タスク1', 0),
  (2, 'タスク2', 0),
  (3, 'タスク3', 0),
  (4, 'タスク4', 0),
  (5, 'タスク5', 0),
  (6, 'タスク6', 0),
  (7, 'タスク7', 0),
  (8, 'タスク8', 0),
  (9, 'タスク9', 0),
  (10, 'タスク10', 0)
;
