-- # 1 column
-- # row 1
-- ## 269
CREATE TABLE public.schema_migration (
	version VARCHAR(14) NOT NULL,
	rowid INT8 NOT VISIBLE NOT NULL DEFAULT unique_rowid(),
	CONSTRAINT "primary" PRIMARY KEY (rowid ASC),
	UNIQUE INDEX schema_migration_version_idx (version ASC),
	FAMILY "primary" (version, rowid)
);
-- # row 2
-- ## 210
CREATE TABLE public.e2e_authors (
	id UUID NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	CONSTRAINT "primary" PRIMARY KEY (id ASC),
	FAMILY "primary" (id, created_at, updated_at)
);
-- # row 3
-- ## 352
CREATE TABLE public.e2e_user_posts (
	id UUID NOT NULL,
	content VARCHAR(255) NOT NULL DEFAULT '':::STRING,
	slug VARCHAR(32) NOT NULL,
	user_id UUID NOT NULL,
	CONSTRAINT "primary" PRIMARY KEY (id ASC),
	UNIQUE INDEX e2e_user_notes_slug_idx (slug ASC),
	INDEX e2e_user_notes_user_id_idx (user_id ASC),
	FAMILY "primary" (id, content, slug, user_id)
);
-- # row 4
-- ## 154
ALTER TABLE public.e2e_user_posts ADD CONSTRAINT e2e_user_notes_e2e_users_id_fk FOREIGN KEY (user_id) REFERENCES public.e2e_authors(id) ON DELETE CASCADE;
-- # row 5
-- ## 115
-- Validate foreign key constraints. These can fail if there was unvalidated data during the SHOW CREATE ALL TABLES
-- # row 6
-- ## 85
ALTER TABLE public.e2e_user_posts VALIDATE CONSTRAINT e2e_user_notes_e2e_users_id_fk;
-- # 6 rows
