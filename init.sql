CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "password" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "subscription_tier" int NOT NULL DEFAULT 0,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "subscription_tier" (
  "id" bigserial PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL
);

CREATE TABLE "journals" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "title" varchar NOT NULL,
  "strategy" text NOT NULL,
  "initial_balance" decimal NOT NULL,
  "current_balance" decimal NOT NULL,
  "risk_tolerance" decimal NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "entries" (
  "id" bigserial PRIMARY KEY,
  "journal_id" bigint NOT NULL,
  "instrument" varchar NOT NULL,
  "position" varchar NOT NULL,
  "lot_size" decimal NOT NULL,
  "opening" decimal NOT NULL,
  "closing" decimal NOT NULL,
  "stop_loss" decimal NOT NULL,
  "take_profit" decimal NOT NULL,
  "risk_reward" decimal NOT NULL,
  "comments" text NOT NULL,
  "before_image" varchar NOT NULL,
  "after_image" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "users" ("id");

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "users" ("id", "email");

CREATE INDEX ON "journals" ("user_id");

CREATE INDEX ON "journals" ("title");

CREATE INDEX ON "entries" ("journal_id");

CREATE INDEX ON "entries" ("instrument");

CREATE INDEX ON "entries" ("position");

ALTER TABLE "users" ADD FOREIGN KEY ("subscription_tier") REFERENCES "subscription_tier" ("id");

ALTER TABLE "journals" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "entries" ADD FOREIGN KEY ("journal_id") REFERENCES "journals" ("id");
