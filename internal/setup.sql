drop table users;

create table users
(
    id         uuid primary key,
    name       text      not null,
    email      text unique,
    password   text      not null,
    created_at timestamp not null
);