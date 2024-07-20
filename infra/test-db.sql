CREATE DATABASE test_snippetbox character
SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE USER 'test_runner' @'localhost';

GRANT CREATE, DROP, ALTER, INDEX, SELECT, INSERT, UPDATE, DELETE ON test_snippetbox.* TO 'test_runner' @'localhost';

ALTER USER 'test_runner' @'localhost' IDENTIFIED BY
    'pass';

