CREATE TABLE e2e_authors (
	id UUID NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	CONSTRAINT "primary" PRIMARY KEY (id ASC),
	FAMILY "primary" (id, created_at, updated_at)
);

CREATE TABLE e2e_user_posts (
	id UUID NOT NULL,
	content VARCHAR(255) NOT NULL DEFAULT '':::STRING,
	slug VARCHAR(32) NOT NULL,
	published BOOL NOT NULL DEFAULT false,
	author_id UUID NULL,
	CONSTRAINT "primary" PRIMARY KEY (id ASC),
	UNIQUE INDEX e2e_user_notes_slug_idx (slug ASC),
	INDEX e2e_user_notes_user_id_idx (author_id ASC),
	FAMILY "primary" (id, content, slug, published, author_id)
);

CREATE TABLE schema_migration (
	version VARCHAR(14) NOT NULL,
	UNIQUE INDEX schema_migration_version_idx (version ASC),
	FAMILY "primary" (version, rowid)
);

ALTER TABLE e2e_user_posts ADD CONSTRAINT e2e_user_notes_e2e_users_id_fk FOREIGN KEY (author_id) REFERENCES e2e_authors(id) ON DELETE CASCADE;

-- Validate foreign key constraints. These can fail if there was unvalidated data during the dump.
ALTER TABLE e2e_user_posts VALIDATE CONSTRAINT e2e_user_notes_e2e_users_id_fk;
