CREATE TABLE users(
	id SERIAL PRIMARY KEY,
	name VARCHAR (100) NOT NULL,
	email VARCHAR (255) UNIQUE NOT NULL,
	oauth_token JSONB,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	deleted_at TIMESTAMP
);

INSERT INTO users(name, email, created_at, updated_at)
VALUES
 ('Miguel Moll', 'me@example.com', NOW(), NOW());
