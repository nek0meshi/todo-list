drop table if exists tasks;
create table tasks (
  id serial primary key,
  name varchar(255),
  status integer
);

insert into tasks (id, name, status) values (1, 'タスク1', 0);
