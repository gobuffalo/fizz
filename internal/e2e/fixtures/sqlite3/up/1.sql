CREATE TABLE IF NOT EXISTS "schema_migration" (
"version" TEXT NOT NULL
);
CREATE UNIQUE INDEX "schema_migration_version_idx" ON "schema_migration" (version);
CREATE TABLE IF NOT EXISTS "e2e_users" (
"id" TEXT PRIMARY KEY,
"username" TEXT,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE TABLE IF NOT EXISTS "e2e_user_notes" (
"id" TEXT PRIMARY KEY,
"notes" TEXT,
"title" TEXT NOT NULL DEFAULT '',
"user_id" char(36) NOT NULL,
FOREIGN KEY (user_id) REFERENCES e2e_users (id) ON DELETE cascade
);
CREATE INDEX "e2e_user_notes_user_id_idx" ON "e2e_user_notes" (user_id);
CREATE INDEX "e2e_user_notes_title_idx" ON "e2e_user_notes" (title);
