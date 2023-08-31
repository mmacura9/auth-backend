ALTER TABLE "user"
ADD "deleted_at" timestamptz DEFAULT NULL;

CREATE TABLE IF NOT EXISTS "sessions" (
    "id" VARCHAR(100) PRIMARY KEY,
    "username" varchar NOT NULL,
    "refresh_token" VARCHAR(511) NOT NULL,
    "user_agent" VARCHAR(255) NOT NULL,
    "client_ip" VARCHAR(45) NOT NULL,
    "is_blocked" BOOLEAN NOT NULL DEFAULT false,
    "expires_at" timestamptz NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);
