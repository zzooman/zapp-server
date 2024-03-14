CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar(255) UNIQUE NOT NULL,
  "password" varchar(255) NOT NULL,
  "email" varchar(255) UNIQUE NOT NULL,
  "phone" varchar(11),
  "location" varchar(255) NOT NULL,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY,
  "owner" varchar NOT NULL,
  "balance" bigint NOT NULL,
  "currency" varchar NOT NULL,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "entries" (
  "id" bigserial PRIMARY KEY,
  "account_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "transfers" (
  "id" bigserial PRIMARY KEY,
  "from_account_id" bigint NOT NULL,
  "to_account_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "posts" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "product_id" bigint,
  "title" varchar(255) NOT NULL,
  "content" text NOT NULL,
  "media" varchar[],
  "location" varchar(255),
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "products" (
  "id" bigserial PRIMARY KEY,
  "seller_id" bigint NOT NULL,
  "name" varchar(255) NOT NULL,
  "description" text,
  "price" bigint NOT NULL,
  "stock" bigint NOT NULL,
  "images" varchar[]
);

CREATE TABLE "orders" (
  "id" bigserial PRIMARY KEY,
  "product_id" bigserial NOT NULL,
  "buyer_id" bigserial NOT NULL,
  "quantity" bigint NOT NULL,
  "price_at_order" bigint NOT NULL,
  "status" varchar(31),
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "like_with_post" (
  "user_id" bigint NOT NULL,
  "post_id" bigint NOT NULL
);

CREATE TABLE "wish_with_product" (
  "user_id" bigint NOT NULL,
  "product_id" bigint NOT NULL
);

CREATE TABLE "reviews" (
  "id" bigserial PRIMARY KEY,
  "product_id" bigint NOT NULL,
  "reviewer_id" bigint NOT NULL,
  "rating" int NOT NULL,
  "content" text,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "comments" (
  "id" bigserial PRIMARY KEY,
  "post_id" bigint NOT NULL,
  "commentor_id" bigint NOT NULL,
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

CREATE INDEX ON "posts" ("user_id");

CREATE INDEX ON "posts" ("product_id");

CREATE INDEX ON "posts" ("created_at");

CREATE INDEX ON "posts" ("location");

CREATE INDEX ON "products" ("seller_id");

CREATE INDEX ON "products" ("name");

CREATE INDEX ON "products" ("price");

CREATE INDEX ON "orders" ("buyer_id");

CREATE INDEX ON "orders" ("product_id");

CREATE INDEX ON "orders" ("created_at");

CREATE INDEX ON "orders" ("status");

CREATE INDEX ON "orders" ("buyer_id", "product_id");

CREATE INDEX ON "like_with_post" ("user_id");

CREATE INDEX ON "like_with_post" ("post_id");

CREATE INDEX ON "wish_with_product" ("user_id");

CREATE INDEX ON "wish_with_product" ("product_id");

CREATE INDEX ON "reviews" ("product_id");

CREATE INDEX ON "reviews" ("reviewer_id");

CREATE INDEX ON "reviews" ("created_at");

CREATE INDEX ON "comments" ("post_id");

CREATE INDEX ON "comments" ("commentor_id");

CREATE INDEX ON "comments" ("created_at");

COMMENT ON COLUMN "entries"."amount" IS 'cab be negative or positive';

COMMENT ON COLUMN "transfers"."amount" IS 'must be positive';

ALTER TABLE "accounts" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");

ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id");

ALTER TABLE "posts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "posts" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "products" ADD FOREIGN KEY ("seller_id") REFERENCES "users" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("buyer_id") REFERENCES "users" ("id");

ALTER TABLE "like_with_post" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "like_with_post" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id");

ALTER TABLE "wish_with_product" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "wish_with_product" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "reviews" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "reviews" ADD FOREIGN KEY ("reviewer_id") REFERENCES "users" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("commentor_id") REFERENCES "users" ("id");
