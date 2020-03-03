drop table wishes;
drop table users;

create table users
(
    id         uuid,
    name       VARCHAR (50)      not null,
    email      VARCHAR (50) unique,
    password   VARCHAR (50)      not null,
    created_at timestamp not null,

    PRIMARY KEY (id)
);

CREATE TABLE wishes
(
    id SERIAL  ,
    owner_id uuid,
    title VARCHAR (50),
    price numeric,
    created_at timestamp not null,

    PRIMARY KEY (id),
    FOREIGN KEY (owner_id) REFERENCES users(id) ON DELETE CASCADE
);

-- psql wishlist < C:\Users\Bar\go\src\github.com\bardromi\wishlist\internal\setup.sql

-- If using powershell
-- cmd /c 'psql wishlist < C:\Users\Bar\go\src\github.com\bardromi\wishlist\internal\setup.sql'