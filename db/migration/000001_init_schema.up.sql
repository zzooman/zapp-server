CREATE TABLE "users" (
  "username" varchar(255) PRIMARY KEY,
  "password" varchar(255) NOT NULL,
  "email" varchar(255) UNIQUE NOT NULL,
  "phone" varchar(11),
  "location" varchar(255),
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "accounts" (
  "id" BIGSERIAL PRIMARY KEY,
  "owner" varchar NOT NULL,
  "balance" bigint NOT NULL,
  "currency" varchar NOT NULL,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "entries" (
  "id" BIGSERIAL PRIMARY KEY,
  "account_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "transfers" (
  "id" BIGSERIAL PRIMARY KEY,
  "from_account_id" bigint NOT NULL,
  "to_account_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "posts" (
  "id" BIGSERIAL PRIMARY KEY,
  "author" varchar(255) NOT NULL,
  "product_id" bigint,
  "title" varchar(255) NOT NULL,
  "content" text NOT NULL,
  "media" varchar[],
  "location" varchar(255),
  "created_at" timestamptz DEFAULT (now()),
  "views" bigint DEFAULT 0
);

CREATE TABLE "products" (
  "id" BIGSERIAL PRIMARY KEY,
  "seller" varchar(255) NOT NULL,
  "name" varchar(255) NOT NULL,
  "description" text,
  "price" bigint NOT NULL,
  "stock" bigint NOT NULL,
  "medias" varchar[]
);

CREATE TABLE "orders" (
  "id" BIGSERIAL PRIMARY KEY,
  "product_id" BIGSERIAL NOT NULL,
  "buyer" varchar(255) NOT NULL,
  "quantity" bigint NOT NULL,
  "price_at_order" bigint NOT NULL,
  "status" varchar(31),
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "like_with_post" (
  "username" varchar(255) NOT NULL,
  "post_id" bigint NOT NULL
);

CREATE TABLE "wish_with_product" (
  "username" varchar(255) NOT NULL,
  "product_id" bigint NOT NULL
);

CREATE TABLE "reviews" (
  "id" BIGSERIAL PRIMARY KEY,
  "product_id" bigint NOT NULL,
  "reviewer" varchar(255) NOT NULL,
  "rating" int NOT NULL,
  "medias" varchar[],
  "content" text,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "comments" (
  "id" BIGSERIAL PRIMARY KEY,
  "post_id" bigint NOT NULL,
  "commentor" varchar(255) NOT NULL,
  "content" text NOT NULL,
  "created_at" timestamptz DEFAULT (now())
);

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "users" ("location");

CREATE INDEX ON "accounts" ("owner");

CREATE UNIQUE INDEX ON "accounts" ("owner", "currency");

CREATE INDEX ON "entries" ("account_id");

CREATE INDEX ON "transfers" ("from_account_id");

CREATE INDEX ON "transfers" ("to_account_id");

CREATE INDEX ON "transfers" ("from_account_id", "to_account_id");

CREATE INDEX ON "posts" ("author");

CREATE INDEX ON "posts" ("product_id");

CREATE INDEX ON "posts" ("created_at");

CREATE INDEX ON "posts" ("location");

CREATE INDEX ON "products" ("seller");

CREATE INDEX ON "products" ("name");

CREATE INDEX ON "products" ("price");

CREATE INDEX ON "orders" ("buyer");

CREATE INDEX ON "orders" ("product_id");

CREATE INDEX ON "orders" ("created_at");

CREATE INDEX ON "orders" ("status");

CREATE INDEX ON "orders" ("buyer", "product_id");

CREATE INDEX ON "like_with_post" ("username");

CREATE INDEX ON "like_with_post" ("post_id");

CREATE INDEX ON "wish_with_product" ("username");

CREATE INDEX ON "wish_with_product" ("product_id");

CREATE INDEX ON "reviews" ("product_id");

CREATE INDEX ON "reviews" ("reviewer");

CREATE INDEX ON "reviews" ("created_at");

CREATE INDEX ON "comments" ("post_id");

CREATE INDEX ON "comments" ("commentor");

CREATE INDEX ON "comments" ("created_at");

COMMENT ON COLUMN "entries"."amount" IS 'cab be negative or positive';

COMMENT ON COLUMN "transfers"."amount" IS 'must be positive';

ALTER TABLE "accounts" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");

ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id");

ALTER TABLE "posts" ADD FOREIGN KEY ("author") REFERENCES "users" ("username");

ALTER TABLE "posts" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "products" ADD FOREIGN KEY ("seller") REFERENCES "users" ("username");

ALTER TABLE "orders" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("buyer") REFERENCES "users" ("username");

ALTER TABLE "like_with_post" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "like_with_post" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id");

ALTER TABLE "wish_with_product" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "wish_with_product" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "reviews" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "reviews" ADD FOREIGN KEY ("reviewer") REFERENCES "users" ("username");

ALTER TABLE "comments" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("commentor") REFERENCES "users" ("username");
