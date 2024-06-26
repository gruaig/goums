CREATE DATABASE goums;


CREATE TABLE "user" (
    "id" SERIAL PRIMARY KEY,
    "Name" VARCHAR(255),
    "LastName" VARCHAR(255),
    "Email" VARCHAR(255) UNIQUE,
    "UserName" VARCHAR(255) UNIQUE,
    "Password" VARCHAR(255),
    "CreatedAt" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "UpdatedAt" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "Enabled" BOOLEAN DEFAULT TRUE
);

