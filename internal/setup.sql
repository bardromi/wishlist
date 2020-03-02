drop table wishes;
drop table users;

create table users
(
    id         uuid,
    name       text      not null,
    email      text unique,
    password   text      not null,
    created_at timestamp not null,

    PRIMARY KEY (id)
);

CREATE TABLE wishes
(
    id SERIAL  ,
    owner_id uuid,
    title text,
    price numeric,
    created_at timestamp not null,

    PRIMARY KEY (id),
    FOREIGN KEY (owner_id) REFERENCES users(id) ON DELETE CASCADE
);

-- psql wishlist < C:\Users\Bar\go\src\github.com\bardromi\wishlist\internal\setup.sql

-- If using powershell
-- cmd /c 'psql wishlist < C:\Users\Bar\go\src\github.com\bardromi\wishlist\internal\setup.sql'