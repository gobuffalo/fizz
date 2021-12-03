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
-- ## 405
CREATE TABLE public.e2e_user_posts (
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
-- # row 4
-- ## 119
CREATE TABLE public.e2e_flow (
	id UUID NOT NULL,
	CONSTRAINT "primary" PRIMARY KEY (id ASC),
	FAMILY "primary" (id)
);
-- # row 5
-- ## 122
CREATE TABLE public.e2e_address (
	id UUID NOT NULL,
	CONSTRAINT "primary" PRIMARY KEY (id ASC),
	FAMILY "primary" (id)
);
-- # row 6
-- ## 517
CREATE TABLE public.e2e_token (
	id UUID NOT NULL,
	token VARCHAR(64) NOT NULL,
	e2e_address_id UUID NOT NULL,
	issued_at TIMESTAMP NOT NULL DEFAULT '2000-01-01 00:00:00':::TIMESTAMP,
	e2e_flow_id UUID NULL,
	flow_id UUID NULL,
	expires_at TIMESTAMP NOT NULL DEFAULT '2001-01-01 00:00:00':::TIMESTAMP,
	CONSTRAINT "primary" PRIMARY KEY (id ASC),
	UNIQUE INDEX e2e_token_uq_idx (token ASC),
	INDEX e2e_token_idx (token ASC),
	FAMILY "primary" (id, token, e2e_address_id, issued_at, e2e_flow_id, flow_id, expires_at)
);
-- # row 7
-- ## 156
ALTER TABLE public.e2e_user_posts ADD CONSTRAINT e2e_user_notes_e2e_users_id_fk FOREIGN KEY (author_id) REFERENCES public.e2e_authors(id) ON DELETE CASCADE;
-- # row 8
-- ## 153
ALTER TABLE public.e2e_token ADD CONSTRAINT e2e_token_e2e_address_id_fk FOREIGN KEY (e2e_address_id) REFERENCES public.e2e_address(id) ON DELETE CASCADE;
-- # row 9
-- ## 144
ALTER TABLE public.e2e_token ADD CONSTRAINT e2e_token_e2e_flow_id_fk FOREIGN KEY (e2e_flow_id) REFERENCES public.e2e_flow(id) ON DELETE CASCADE;
-- # row 10
-- ## 137
ALTER TABLE public.e2e_token ADD CONSTRAINT e2e_token_flow_id_fk FOREIGN KEY (flow_id) REFERENCES public.e2e_flow(id) ON DELETE RESTRICT;
-- # row 11
-- ## 115
-- Validate foreign key constraints. These can fail if there was unvalidated data during the SHOW CREATE ALL TABLES
-- # row 12
-- ## 85
ALTER TABLE public.e2e_user_posts VALIDATE CONSTRAINT e2e_user_notes_e2e_users_id_fk;
-- # row 13
-- ## 77
ALTER TABLE public.e2e_token VALIDATE CONSTRAINT e2e_token_e2e_address_id_fk;
-- # row 14
-- ## 74
ALTER TABLE public.e2e_token VALIDATE CONSTRAINT e2e_token_e2e_flow_id_fk;
-- # row 15
-- ## 70
ALTER TABLE public.e2e_token VALIDATE CONSTRAINT e2e_token_flow_id_fk;
-- # 15 rows
