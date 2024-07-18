-- Create a `snippets` table.
CREATE TABLE IF NOT EXISTS snippets(
    id integer NOT NULL PRIMARY KEY AUTO_INCREMENT,
    title varchar(100) NOT NULL,
    content text NOT NULL,
    created DATETIME NOT NULL,
    expires DATETIME NOT NULL
);

CREATE INDEX idx_snippets_created ON snippets(created);

-- Create a `sessions` table.
CREATE TABLE IF NOT EXISTS sessions(
    token char(43) PRIMARY KEY,
    data BLOB NOT NULL,
    expiry timestamp(6) NOT NULL
);

CREATE INDEX sessions_expiry_idx ON sessions(expiry);

-- Create a `users` table
CREATE TABLE IF NOT EXISTS users(
    id integer NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    hashed_password char(60) NOT NULL,
    created DATETIME NOT NULL
);

ALTER TABLE users
    ADD CONSTRAINT users_uc_email UNIQUE (email);

