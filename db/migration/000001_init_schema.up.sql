CREATE TABLE IF NOT EXISTS "user" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "email" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "birth_date" timestamptz NOT NULL,
  "password" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "last_login" timestamptz NOT NULL DEFAULT (now())
);


CREATE INDEX ON "user" ("id");

CREATE INDEX ON "user" ("username");

CREATE TABLE IF NOT EXISTS "sessions" (
    "id" VARCHAR(100) PRIMARY KEY,
    "username" varchar NOT NULL,
    "refresh_token" VARCHAR(511) NOT NULL,
    "user_agent" VARCHAR(255) NOT NULL,
    "client_ip" VARCHAR(45) NOT NULL,
    "is_blocked" BOOLEAN NOT NULL DEFAULT false,
    "expires_at" TIMESTAMP NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
