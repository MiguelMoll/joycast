CREATE TABLE users(
	id SERIAL PRIMARY KEY,
	name VARCHAR (50) NOT NULL,
	email VARCHAR (355) UNIQUE NOT NULL,
	oauth_token JSONB,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	deleted_at TIMESTAMP
);

INSERT INTO users(name, email, created_at, update_at)
VALUES
 ('Miguel Moll', 'me@example.com', NOW(), NOW());
