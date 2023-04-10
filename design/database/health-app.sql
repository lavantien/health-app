CREATE TABLE "users" (
  "id" integer PRIMARY KEY,
  "name" varchar,
  "password" varchar,
  "birthday" varchar
);

CREATE TABLE "meals" (
  "id" integer PRIMARY KEY,
  "user_id" integer,
  "occasion" integer,
  "name" varchar,
  "date" varchar
);

CREATE TABLE "records" (
  "id" integer PRIMARY KEY,
  "user_id" integer,
  "date" varchar,
  "weight" integer,
  "body_fat" integer
);

CREATE TABLE "exercises" (
  "id" integer PRIMARY KEY,
  "user_id" integer,
  "duration" integer,
  "calories" integer,
  "date" varchar
);

CREATE TABLE "diaries" (
  "id" integer PRIMARY KEY,
  "user_id" integer,
  "date" varchar,
  "time" varchar,
  "title" varchar,
  "content" varchar
);

ALTER TABLE "meals" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "records" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "exercises" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "diaries" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
