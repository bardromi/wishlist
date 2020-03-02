drop table users;

create table users
(
    id         uuid primary key,
    name       text      not null,
    email      text unique,
    password   text      not null,
    created_at timestamp not null
);

create table wishes
(
    id SERIAL  primary key,
    owner_id uuid REFERENCES users(id),
    title text,
    price numeric,
    created_at timestamp not null
)

-- psql wishlist < C:\Users\Bar\go\src\github.com\bardromi\wishlist\internal\setup.sql

-- If using posershell
-- cmd /c 'psql wishlist < C:\Users\Bar\go\src\github.com\bardromi\wishlist\internal\setup.sql'