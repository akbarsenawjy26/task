CREATE TABLE "users" (
    "GUID" SERIAL,
    "email" VARCHAR(255) NOT NULL UNIQUE ,
    "username" VARCHAR(255) NOT NULL UNIQUE,
    "password" VARCHAR NOT NULL UNIQUE,
    "created_at" TIMESTAMP NOT NULL ,
    "updated_at" TIMESTAMP NOT NULL ,

    PRIMARY KEY ("GUID")
);