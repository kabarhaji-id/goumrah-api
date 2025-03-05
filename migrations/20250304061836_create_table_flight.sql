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
    "next_id" BIGINT NULL DEFAULT NULL,

    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "deleted_at" TIMESTAMP NULL DEFAULT NULL,

    CONSTRAINT "flight_routes_id_pkey" PRIMARY KEY ("id"),
    CONSTRAINT "flight_routes_flight_id_fkey" FOREIGN KEY ("flight_id") REFERENCES "flights" ("id"),
    CONSTRAINT "flight_routes_next_id_fkey" FOREIGN KEY ("next_id") REFERENCES "flight_routes" ("id")
);

ALTER TABLE IF EXISTS "package_sessions"
ADD COLUMN "departure_flight_route_id" BIGINT NULL DEFAULT NULL,
ADD COLUMN "return_flight_route_id" BIGINT NULL DEFAULT NULL,
ADD CONSTRAINT "package_sessions_departure_flight_route_id_fkey" FOREIGN KEY ("departure_flight_route_id") REFERENCES "flight_routes" ("id"),
ADD CONSTRAINT "package_sessions_return_flight_route_id_fkey" FOREIGN KEY ("return_flight_route_id") REFERENCES "flight_routes" ("id");

CREATE OR REPLACE FUNCTION delete_flight_on_airline_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL THEN
        UPDATE flights SET deleted_at = NOW() WHERE airline_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION prevent_insert_flight_if_airline_is_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF (SELECT deleted_at FROM airlines WHERE id = NEW.airline_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert flight with soft deleted airline'
            USING ERRCODE = '23503', CONSTRAINT = 'flight_routes_flight_id_fkey';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER delete_flight_on_airline_soft_deleted
BEFORE UPDATE ON airlines
FOR EACH ROW
WHEN (OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL)
EXECUTE FUNCTION delete_flight_on_airline_soft_deleted();

CREATE TRIGGER prevent_insert_flight_if_airline_is_soft_deleted
BEFORE INSERT OR UPDATE ON flights
FOR EACH ROW
EXECUTE FUNCTION prevent_insert_flight_if_airline_is_soft_deleted();

CREATE OR REPLACE FUNCTION delete_flight_on_departure_airport_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL THEN
        UPDATE flights SET deleted_at = NOW() WHERE departure_airport_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION prevent_insert_flight_if_departure_airport_is_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF (SELECT deleted_at FROM airports WHERE id = NEW.departure_airport_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert flight with soft deleted departure airport'
            USING ERRCODE = '23503', CONSTRAINT = 'flights_departure_airport_id_fkey';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER delete_flight_on_departure_airport_soft_deleted
BEFORE UPDATE ON airports
FOR EACH ROW
WHEN (OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL)
EXECUTE FUNCTION delete_flight_on_departure_airport_soft_deleted();

CREATE TRIGGER prevent_insert_flight_if_departure_airport_is_soft_deleted
BEFORE INSERT OR UPDATE ON flights
FOR EACH ROW
EXECUTE FUNCTION prevent_insert_flight_if_departure_airport_is_soft_deleted();

CREATE OR REPLACE FUNCTION delete_flight_on_arrival_airport_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL THEN
        UPDATE flights SET deleted_at = NOW() WHERE arrival_airport_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION prevent_insert_flight_if_arrival_airport_is_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF (SELECT deleted_at FROM airports WHERE id = NEW.arrival_airport_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert flight with soft deleted arrival airport'
            USING ERRCODE = '23503', CONSTRAINT = 'flights_arrival_airport_id_fkey';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER delete_flight_on_arrival_airport_soft_deleted
BEFORE UPDATE ON airports
FOR EACH ROW
WHEN (OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL)
EXECUTE FUNCTION delete_flight_on_arrival_airport_soft_deleted();

CREATE TRIGGER prevent_insert_flight_if_arrival_airport_is_soft_deleted
BEFORE INSERT OR UPDATE ON flights
FOR EACH ROW
EXECUTE FUNCTION prevent_insert_flight_if_arrival_airport_is_soft_deleted();

CREATE OR REPLACE FUNCTION delete_flight_route_on_flight_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL THEN
        UPDATE flight_routes SET deleted_at = NOW() WHERE flight_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION prevent_insert_flight_route_if_flight_is_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF (SELECT deleted_at FROM flights WHERE id = NEW.flight_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert flight route with soft deleted flight'
            USING ERRCODE = '23503', CONSTRAINT = 'flight_routes_flight_id_fkey';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER delete_flight_route_on_flight_soft_deleted
BEFORE UPDATE ON flights
FOR EACH ROW
WHEN (OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL)
EXECUTE FUNCTION delete_flight_route_on_flight_soft_deleted();

CREATE TRIGGER prevent_insert_flight_route_if_flight_is_soft_deleted
BEFORE INSERT OR UPDATE ON flight_routes
FOR EACH ROW
EXECUTE FUNCTION prevent_insert_flight_route_if_flight_is_soft_deleted();

CREATE OR REPLACE FUNCTION set_next_id_null_on_flight_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.deleted_at IS NOT NULL THEN
        UPDATE flight_routes SET next_id = NULL WHERE next_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION prevent_insert_flight_route_if_next_is_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.next_id IS NOT NULL THEN
        IF (SELECT deleted_at FROM flight_routes WHERE id = NEW.next_id) IS NOT NULL THEN
            RAISE EXCEPTION 'Cannot insert flight route with soft deleted next'
                USING ERRCODE = '23503', CONSTRAINT = 'flights_return_flight_route_id_fkey';
        END IF;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_next_id_null_on_flight_soft_deleted
BEFORE UPDATE ON flight_routes
FOR EACH ROW
WHEN (OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL)
EXECUTE FUNCTION set_next_id_null_on_flight_soft_deleted();

CREATE TRIGGER prevent_insert_flight_route_if_next_is_soft_deleted
BEFORE INSERT OR UPDATE ON flight_routes
FOR EACH ROW
EXECUTE FUNCTION prevent_insert_flight_route_if_next_is_soft_deleted();

CREATE OR REPLACE FUNCTION set_package_session_departure_flight_route_id_null_on_flight_route_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.deleted_at IS NOT NULL THEN
        UPDATE package_sessions SET departure_flight_route_id = NULL WHERE departure_flight_route_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION prevent_insert_package_session_if_departure_flight_route_is_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.departure_flight_route_id IS NOT NULL THEN
        IF (SELECT deleted_at FROM flight_routes WHERE id = NEW.departure_flight_route_id) IS NOT NULL THEN
            RAISE EXCEPTION 'Cannot insert package session with soft deleted departure flight route'
                USING ERRCODE = '23503', CONSTRAINT = 'package_sessions_departure_flight_route_id_fkey';
        END IF;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_package_session_departure_flight_route_id_null_on_flight_route_soft_deleted
BEFORE UPDATE ON flight_routes
FOR EACH ROW
WHEN (OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL)
EXECUTE FUNCTION set_package_session_departure_flight_route_id_null_on_flight_route_soft_deleted();

CREATE TRIGGER prevent_insert_package_session_if_departure_flight_route_is_soft_deleted
BEFORE INSERT OR UPDATE ON package_sessions
FOR EACH ROW
EXECUTE FUNCTION prevent_insert_package_session_if_departure_flight_route_is_soft_deleted();

CREATE OR REPLACE FUNCTION set_package_session_return_flight_route_id_null_on_flight_route_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.deleted_at IS NOT NULL THEN
        UPDATE package_sessions SET return_flight_route_id = NULL WHERE return_flight_route_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION prevent_insert_package_session_if_return_flight_route_is_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.return_flight_route_id IS NOT NULL THEN
        IF (SELECT deleted_at FROM flight_routes WHERE id = NEW.return_flight_route_id) IS NOT NULL THEN
            RAISE EXCEPTION 'Cannot insert package session with soft deleted return flight route'
                USING ERRCODE = '23503', CONSTRAINT = 'package_sessions_return_flight_route_id_fkey';
        END IF;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_package_session_return_flight_route_id_null_on_flight_route_soft_deleted
BEFORE UPDATE ON flight_routes
FOR EACH ROW
WHEN (OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL)
EXECUTE FUNCTION set_package_session_return_flight_route_id_null_on_flight_route_soft_deleted();

CREATE TRIGGER prevent_insert_package_session_if_return_flight_route_is_soft_deleted
BEFORE INSERT OR UPDATE ON package_sessions
FOR EACH ROW
EXECUTE FUNCTION prevent_insert_package_session_if_return_flight_route_is_soft_deleted();

-- migrate:down

DROP TRIGGER IF EXISTS "prevent_insert_package_session_if_return_flight_route_is_soft_deleted" ON "package_sessions";
DROP TRIGGER IF EXISTS "set_package_session_return_flight_route_id_null_on_flight_route_soft_deleted" ON "flight_routes";
DROP TRIGGER IF EXISTS "prevent_insert_package_session_if_departure_flight_route_is_soft_deleted" ON "package_sessions";
DROP TRIGGER IF EXISTS "set_package_session_departure_flight_route_id_null_on_flight_route_soft_deleted" ON "flight_routes";
DROP TRIGGER IF EXISTS "prevent_insert_flight_route_if_next_is_soft_deleted" ON "flight_routes";
DROP TRIGGER IF EXISTS "set_next_id_null_on_flight_soft_deleted" ON "flight_routes";
DROP TRIGGER IF EXISTS "prevent_insert_flight_route_if_flight_is_soft_deleted" ON "flight_routes";
DROP TRIGGER IF EXISTS "delete_flight_route_on_flight_soft_deleted" ON "flights";
DROP TRIGGER IF EXISTS "prevent_insert_flight_if_arrival_airport_is_soft_deleted" ON "flights";
DROP TRIGGER IF EXISTS "delete_flight_on_arrival_airport_soft_deleted" ON "airports";
DROP TRIGGER IF EXISTS "prevent_insert_flight_if_departure_airport_is_soft_deleted" ON "flights";
DROP TRIGGER IF EXISTS "delete_flight_on_departure_airport_soft_deleted" ON "airports";
DROP TRIGGER IF EXISTS "prevent_insert_flight_if_airline_is_soft_deleted" ON "flights";
DROP TRIGGER IF EXISTS "delete_flight_on_airline_soft_deleted" ON "airlines";

DROP FUNCTION IF EXISTS "prevent_insert_package_session_if_return_flight_route_is_soft_deleted";
DROP FUNCTION IF EXISTS "set_package_session_return_flight_route_id_null_on_flight_route_soft_deleted";
DROP FUNCTION IF EXISTS "prevent_insert_package_session_if_departure_flight_route_is_soft_deleted";
DROP FUNCTION IF EXISTS "set_package_session_departure_flight_route_id_null_on_flight_route_soft_deleted";
DROP FUNCTION IF EXISTS "prevent_insert_flight_route_if_next_is_soft_deleted";
DROP FUNCTION IF EXISTS "set_next_id_null_on_flight_soft_deleted";
DROP FUNCTION IF EXISTS "prevent_insert_flight_route_if_flight_is_soft_deleted";
DROP FUNCTION IF EXISTS "delete_flight_route_on_flight_soft_deleted";
DROP FUNCTION IF EXISTS "prevent_insert_flight_if_arrival_airport_is_soft_deleted";
DROP FUNCTION IF EXISTS "delete_flight_on_arrival_airport_soft_deleted";
DROP FUNCTION IF EXISTS "prevent_insert_flight_if_departure_airport_is_soft_deleted";
DROP FUNCTION IF EXISTS "delete_flight_on_departure_airport_soft_deleted";
DROP FUNCTION IF EXISTS "prevent_insert_flight_if_airline_is_soft_deleted";
DROP FUNCTION IF EXISTS "delete_flight_on_airline_soft_deleted";

ALTER TABLE IF EXISTS "package_sessions"
DROP COLUMN IF EXISTS "departure_flight_route_id",
DROP COLUMN IF EXISTS "return_flight_route_id";

DROP TABLE IF EXISTS "flight_routes";

DROP TABLE IF EXISTS "flights";

DROP TYPE IF EXISTS "flight_class";
