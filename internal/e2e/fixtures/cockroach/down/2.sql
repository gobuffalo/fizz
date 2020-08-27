CREATE TABLE e2e_users (
	id UUID NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	username VARCHAR(255) NULL,
	CONSTRAINT "primary" PRIMARY KEY (id ASC),
	FAMILY "primary" (id, created_at, updated_at, username)
);

CREATE TABLE e2e_user_notes (
	id UUID NOT NULL,
	user_id UUID NOT NULL,
	notes VARCHAR(255) NULL,
	title VARCHAR(64) NOT NULL DEFAULT '':::STRING,
	CONSTRAINT "primary" PRIMARY KEY (id ASC),
	INDEX e2e_user_notes_auto_index_e2e_user_notes_e2e_users_id_fk (user_id ASC),
	INDEX e2e_user_notes_user_id_idx (user_id ASC),
	INDEX e2e_user_notes_title_idx (title ASC),
	FAMILY "primary" (id, user_id, notes, title)
);

CREATE TABLE schema_migration (
	version VARCHAR(14) NOT NULL,
	UNIQUE INDEX schema_migration_version_idx (version ASC),
	FAMILY "primary" (version, rowid)
);

ALTER TABLE e2e_user_notes ADD CONSTRAINT e2e_user_notes_e2e_users_id_fk FOREIGN KEY (user_id) REFERENCES e2e_users(id) ON DELETE CASCADE;

-- Validate foreign key constraints. These can fail if there was unvalidated data during the dump.
ALTER TABLE e2e_user_notes VALIDATE CONSTRAINT e2e_user_notes_e2e_users_id_fk;
