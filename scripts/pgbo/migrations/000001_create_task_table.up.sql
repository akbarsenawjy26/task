CREATE TABLE "task" (
    "id" SERIAL,
    "description" VARCHAR NOT NULL,
    "status" VARCHAR,
    "created_at" TIMESTAMP,
    "updated_at" TIMESTAMP,

    PRIMARY KEY ("id")
);