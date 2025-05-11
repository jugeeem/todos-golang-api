CREATE TABLE IF NOT EXISTS todos (
	id		serial 				primary key

	,title		varchar(32) 			not null
	,description	varchar(255)			not null
	,completed	boolean				not null default false

	,user_id	integer				not null

	,created_at	timestamp with time zone	not null default current_timestamp
	,updated_at	timestamp with time zone	not null default current_timestamp
	,delete_flag	boolean				not null default false

	,CONSTRAINT fk_todos_user
		FOREIGN KEY (user_id)
		REFERENCES users(id)
		ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_todos_user_id ON todos(user_id);
