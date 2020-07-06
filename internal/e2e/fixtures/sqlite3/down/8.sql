CREATE TABLE IF NOT EXISTS "schema_migration" (
"version" TEXT NOT NULL
);
CREATE UNIQUE INDEX "schema_migration_version_idx" ON "schema_migration" (version);
CREATE TABLE IF NOT EXISTS "e2e_user_posts" (
"id" TEXT PRIMARY KEY,
"user_id" char(36) NOT NULL,
"slug" TEXT NOT NULL,
"content" TEXT NOT NULL DEFAULT '',
FOREIGN KEY (user_id) REFERENCES "_e2e_users_tmp" (id) ON UPDATE NO ACTION ON DELETE CASCADE
);
CREATE UNIQUE INDEX "e2e_user_notes_slug_idx" ON "e2e_user_posts" (slug);
CREATE TABLE IF NOT EXISTS "e2e_users" (
"id" TEXT PRIMARY KEY,
"username" TEXT,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
