-- migrate:up

CREATE TYPE "flight_class" AS ENUM ('Economy', 'Business', 'First');

CREATE TABLE IF NOT EXISTS "flights" (
    "id" BIGINT GENERATED ALWAYS AS IDENTITY NOT NULL,
    "airline_id" BIGINT NOT NULL,
    "aircraft" VARCHAR(100) NOT NULL,
    "baggage" DECIMAL(8, 2) NOT NULL,
    "cabin_baggage" DECIMAL(8, 2) NOT NULL,
    "departure_airport_id" BIGINT NOT NULL,
    "departure_terminal" VARCHAR(100) NULL,
    "departure_at" TIME NOT NULL,
    "arrival_airport_id" BIGINT NOT NULL,
    "arrival_terminal" VARCHAR(100) NULL,
    "arrival_at" TIME NOT NULL,
    "code" VARCHAR(10) NOT NULL,
    "seat_layout" VARCHAR(10) NOT NULL,
    "class" "flight_class" NOT NULL,

    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "deleted_at" TIMESTAMP NULL DEFAULT NULL,
    
    CONSTRAINT "flights_id_pkey" PRIMARY KEY ("id"),
    CONSTRAINT "flights_airline_id_fkey" FOREIGN KEY ("airline_id") REFERENCES "airlines" ("id"),
    CONSTRAINT "flights_departure_airport_id_fkey" FOREIGN KEY ("departure_airport_id") REFERENCES "airports" ("id"),
    CONSTRAINT "flights_arrival_airport_id_fkey" FOREIGN KEY ("arrival_airport_id") REFERENCES "airports" ("id")
);

CREATE TABLE IF NOT EXISTS "flight_routes" (
    "id" BIGINT GENERATED ALWAYS AS IDENTITY NOT NULL,
    "flight_id" BIGINT NOT NULL,
    "next_flight_id" BIGINT NULL DEFAULT NULL,

    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "deleted_at" TIMESTAMP NULL DEFAULT NULL,

    CONSTRAINT "flight_routes_id_pkey" PRIMARY KEY ("id"),
    CONSTRAINT "flight_routes_flight_id_fkey" FOREIGN KEY ("flight_id") REFERENCES "flights" ("id"),
    CONSTRAINT "flight_routes_next_flight_id_fkey" FOREIGN KEY ("next_flight_id") REFERENCES "flights" ("id")
);

ALTER TABLE IF EXISTS "flights"
ADD COLUMN "departure_flight_route_id" BIGINT NULL DEFAULT NULL,
ADD COLUMN "return_flight_route_id" BIGINT NULL DEFAULT NULL,
ADD CONSTRAINT "flights_departure_flight_route_id_fkey" FOREIGN KEY ("departure_flight_route_id") REFERENCES "flight_routes" ("id"),
ADD CONSTRAINT "flights_return_flight_route_id_fkey" FOREIGN KEY ("return_flight_route_id") REFERENCES "flight_routes" ("id");

-- migrate:down

ALTER TABLE IF EXISTS "flights"
DROP COLUMN IF EXISTS "departure_flight_route_id",
DROP COLUMN IF EXISTS "return_flight_route_id";

DROP TABLE IF EXISTS "flight_routes";

DROP TABLE IF EXISTS "flights";

DROP TYPE IF EXISTS "flight_class";
