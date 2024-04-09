CREATE TABLE "user"
(
    "id"              bigserial PRIMARY KEY,
    "email"           varchar unique NOT NULL,
    "hashed_password" varchar        NOT NULL,
    "created_at"      timestamptz    NOT NULL DEFAULT (now())
);

CREATE TABLE "solution"
(
    "id"         bigserial PRIMARY KEY,
    "condition"  JSONB       NOT NULL,
    "solution"   JSONB       NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);


CREATE INDEX ON "solution" ("condition");

CREATE INDEX ON "user" ("email");

-- ALTER TABLE "solution"
--     ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id") ON DELETE CASCADE;
