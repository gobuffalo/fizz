CREATE TABLE IF NOT EXISTS "schema_migration" (
"version" TEXT NOT NULL
);
CREATE UNIQUE INDEX "schema_migration_version_idx" ON "schema_migration" (version);
CREATE TABLE IF NOT EXISTS "e2e_users" (
"id" TEXT PRIMARY KEY,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE TABLE IF NOT EXISTS "e2e_user_notes" (
"id" TEXT PRIMARY KEY,
"notes" TEXT,
"user_id" char(36) NOT NULL,
FOREIGN KEY (user_id) REFERENCES e2e_users (id) ON UPDATE NO ACTION ON DELETE CASCADE
);
