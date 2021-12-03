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
-- # 1 row
