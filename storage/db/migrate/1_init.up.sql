CREATE TABLE users(
	id SERIAL PRIMARY KEY,
	name VARCHAR (50) UNIQUE NOT NULL,
	email VARCHAR (355) UNIQUE NOT NULL,
	created_at TIMESTAMP NOT NULL,
	update_at TIMESTAMP,
	deleted_at TIMESTAMP
);

CREATE TABLE tokens(
	id SERIAL PRIMARY KEY,
	user_id INTEGER,
	data JSONB,
	created_at TIMESTAMP NOT NULL,
	update_at TIMESTAMP,
	deleted_at TIMESTAMP,
	FOREIGN KEY (user_id) REFERENCES users (id)
);

INSERT INTO users(name, email, created_at)
VALUES
 ('Miguel Moll', 'me@example.com', NOW());
