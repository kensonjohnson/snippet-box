CREATE TABLE snippets(
    id integer NOT NULL PRIMARY KEY AUTO_INCREMENT,
    title varchar(100) NOT NULL,
    content text NOT NULL,
    created DATETIME NOT NULL,
    expires DATETIME NOT NULL
);

CREATE INDEX idx_snippets_created ON snippets(created);

CREATE TABLE users(
    id integer NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    hashed_password char(60) NOT NULL,
    created DATETIME NOT NULL
);

ALTER TABLE users
    ADD CONSTRAINT users_uc_email UNIQUE (email);

INSERT INTO users(name, email, hashed_password, created)
    VALUES ('Alice Jones', 'alice@example.com', '$2a$12$NuTjWXm3KKntReFwyBVHyuf/to.HEwTy.eS206TNfkGfr6HzGJSWG', '2022-01-01 09:18:24');

