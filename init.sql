CREATE TABLE IF NOT EXISTS users (
    id serial PRIMARY KEY,
    name VARCHAR NOT NULL,
    email VARCHAR NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS events (
    id serial PRIMARY KEY,
    user_id INTEGER NOT NULL,
	name TEXT NOT NULL,
	complete BOOLEAN,
	recurring BOOLEAN,
	frequency VARCHAR,
	created TIMESTAMP,
	completed TIMESTAMP,
	updated TIMESTAMP,
	reopen TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

