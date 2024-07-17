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

