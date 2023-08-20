CREATE TABLE IF NOT EXISTS "user" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL UNIQUE,
  "email" varchar NOT NULL UNIQUE,
  "full_name" varchar NOT NULL,
  "birth_date" timestamptz NOT NULL,
  "password" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "last_login" timestamptz NOT NULL DEFAULT (now())
);


CREATE INDEX ON "user" ("id");

CREATE INDEX ON "user" ("username");
