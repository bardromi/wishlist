drop table wishes;
drop table users;

CREATE TABLE users
(
    id         UUID,
    name       TEXT      not null,
    email      TEXT     unique,
    password   TEXT      not null,

    date_created TIMESTAMP NOT NULL,
    date_updated TIMESTAMP,

    PRIMARY KEY (id)
);

CREATE TABLE wishes
(
    id SERIAL  ,
    owner_id UUID,
    title TEXT,
    price NUMERIC,

    date_created TIMESTAMP NOT NULL,
    date_updated TIMESTAMP,

    PRIMARY KEY (id),
    FOREIGN KEY (owner_id) REFERENCES users(id) ON DELETE CASCADE
);

-- psql wishlist < C:\Users\Bar\go\src\github.com\bardromi\wishlist\internal\setup.sql

-- If using powershell
-- cmd /c 'psql wishlist < C:\Users\Bar\go\src\github.com\bardromi\wishlist\internal\setup.sql'