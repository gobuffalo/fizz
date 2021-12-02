CREATE TABLE IF NOT EXISTS "schema_migration" (
"version" TEXT NOT NULL
);
CREATE UNIQUE INDEX "schema_migration_version_idx" ON "schema_migration" (version);
CREATE TABLE IF NOT EXISTS "e2e_authors" (
"id" TEXT PRIMARY KEY,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE TABLE IF NOT EXISTS "e2e_user_posts" (
"id" TEXT PRIMARY KEY,
"author_id" char(36),
"slug" TEXT NOT NULL,
"content" TEXT NOT NULL DEFAULT '',
"published" bool NOT NULL DEFAULT FALSE,
FOREIGN KEY (author_id) REFERENCES e2e_authors (id) ON UPDATE NO ACTION ON DELETE CASCADE
);
CREATE INDEX "e2e_user_notes_user_id_idx" ON "e2e_user_posts" (author_id);
CREATE UNIQUE INDEX "e2e_user_notes_slug_idx" ON "e2e_user_posts" (slug);
CREATE TABLE IF NOT EXISTS "e2e_flow" (
"id" TEXT PRIMARY KEY
);
CREATE TABLE IF NOT EXISTS "e2e_address" (
"id" TEXT PRIMARY KEY
);
CREATE TABLE IF NOT EXISTS "e2e_token" (
"id" TEXT PRIMARY KEY,
"token" TEXT NOT NULL,
"e2e_flow_id" char(36),
"e2e_address_id" char(36) NOT NULL,
"expires_at" DATETIME NOT NULL DEFAULT '2000-01-01 00:00:00',
"issued_at" DATETIME NOT NULL DEFAULT '2000-01-01 00:00:00',
FOREIGN KEY (e2e_address_id) REFERENCES e2e_address (id) ON UPDATE NO ACTION ON DELETE CASCADE,
FOREIGN KEY (e2e_flow_id) REFERENCES e2e_flow (id) ON UPDATE NO ACTION ON DELETE CASCADE
);
CREATE INDEX "e2e_token_idx" ON "e2e_token" (token);
CREATE UNIQUE INDEX "e2e_token_uq_idx" ON "e2e_token" (token);
