CREATE TABLE IF NOT EXISTS users (
	id		serial 				primary key

	,username	varchar(32) 			not null
	,password	varchar(255)			not null
	,email		varchar(255)			not null

	,created_at	timestamp with time zone	not null default current_timestamp
	,updated_at	timestamp with time zone	not null default current_timestamp
	,delete_flag	boolean				not null default false

	,CONSTRAINT uq_users_username
		UNIQUE (username)
	,CONSTRAINT uq_users_email
		UNIQUE (email)	
);

CREATE INDEX IF NOT EXISTS idx_users_username ON users(username);
CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
