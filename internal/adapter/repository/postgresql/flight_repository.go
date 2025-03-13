package postgresql

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driven/repository"
	"github.com/kabarhaji-id/goumrah-api/pkg/sqlbuilder"
)

type flightRepositoryPostgresql struct {
	db DB
}

func NewFlightRepository(db DB) repository.FlightRepository {
	return flightRepositoryPostgresql{db}
}

func (r flightRepositoryPostgresql) Create(ctx context.Context, flight entity.Flight) (entity.Flight, error) {
	builder := sqlbuilder.New().
		S(`INSERT INTO "flights" ("airline_id", "aircraft", "baggage", "cabin_baggage", "departure_airport_id", "departure_terminal", "departure_at", "arrival_airport_id", "arrival_terminal", "arrival_at", "code", "seat_layout", "class", "created_at", "updated_at", "deleted_at")`).
		S(
			`VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, NOW(), NOW(), NULL)`,
			flight.AirlineId, flight.Aircraft, flight.Baggage, flight.CabinBaggage, flight.DepartureAirportId, flight.DepartureTerminal, flight.DepartureAt, flight.ArrivalAirportId, flight.ArrivalTerminal, flight.ArrivalAt, flight.Code, flight.SeatLayout, flight.Class,
		).
		S(`RETURNING "id", "airline_id", "aircraft", "baggage", "cabin_baggage", "departure_airport_id", "departure_terminal", "departure_at", "arrival_airport_id", "arrival_terminal", "arrival_at", "code", "seat_layout", "class", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&flight.Id, &flight.AirlineId, &flight.Aircraft, &flight.Baggage, &flight.CabinBaggage, &flight.DepartureAirportId, &flight.DepartureTerminal, &flight.DepartureAt, &flight.ArrivalAirportId, &flight.ArrivalTerminal, &flight.ArrivalAt, &flight.Code, &flight.SeatLayout, &flight.Class,
		&flight.CreatedAt, &flight.UpdatedAt, &flight.DeletedAt,
	)

	return flight, err
}

func (r flightRepositoryPostgresql) FindById(ctx context.Context, id int64) (entity.Flight, error) {
	flight := entity.Flight{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "airline_id", "aircraft", "baggage", "cabin_baggage", "departure_airport_id", "departure_terminal", "departure_at", "arrival_airport_id", "arrival_terminal", "arrival_at", "code", "seat_layout", "class", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "flights" WHERE "id" = $1 AND "deleted_at" IS NULL`, id)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&flight.Id, &flight.AirlineId, &flight.Aircraft, &flight.Baggage, &flight.CabinBaggage, &flight.DepartureAirportId, &flight.DepartureTerminal, &flight.DepartureAt, &flight.ArrivalAirportId, &flight.ArrivalTerminal, &flight.ArrivalAt, &flight.Code, &flight.SeatLayout, &flight.Class,
		&flight.CreatedAt, &flight.UpdatedAt, &flight.DeletedAt,
	)

	return flight, err
}

func (r flightRepositoryPostgresql) FindAll(ctx context.Context, opt repository.FindAllOptions) ([]entity.Flight, error) {
	flights := []entity.Flight{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "airline_id", "aircraft", "baggage", "cabin_baggage", "departure_airport_id", "departure_terminal", "departure_at", "arrival_airport_id", "arrival_terminal", "arrival_at", "code", "seat_layout", "class", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "flights" WHERE "deleted_at" IS NULL`).
		S(`ORDER BY "id" ASC`)
	if opt.Limit.Valid {
		builder.SA(`LIMIT ?`, opt.Limit)
	}
	if opt.Offset.Valid {
		builder.SA(`OFFSET ?`, opt.Offset)
	}

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return flights, err
	}

	for rows.Next() {
		flight := entity.Flight{}
		err = rows.Scan(
			&flight.Id, &flight.AirlineId, &flight.Aircraft, &flight.Baggage, &flight.CabinBaggage, &flight.DepartureAirportId, &flight.DepartureTerminal, &flight.DepartureAt, &flight.ArrivalAirportId, &flight.ArrivalTerminal, &flight.ArrivalAt, &flight.Code, &flight.SeatLayout, &flight.Class,
			&flight.CreatedAt, &flight.UpdatedAt, &flight.DeletedAt,
		)
		if err != nil {
			return flights, err
		}

		flights = append(flights, flight)
	}

	return flights, nil
}

func (r flightRepositoryPostgresql) Update(ctx context.Context, id int64, flight entity.Flight) (entity.Flight, error) {
	builder := sqlbuilder.New().
		S(
			`UPDATE "flights" SET "airline_id" = $1, "aircraft" = $2, "baggage" = $3, "cabin_baggage" = $4, "departure_airport_id" = $5, "departure_terminal" = $6, "departure_at" = $7, "arrival_airport_id" = $8, "arrival_terminal" = $9, "arrival_at" = $10, "code" = $11, "seat_layout" = $12, "class" = $13, "updated_at" = NOW()`,
			flight.AirlineId, flight.Aircraft, flight.Baggage, flight.CabinBaggage, flight.DepartureAirportId, flight.DepartureTerminal, flight.DepartureAt, flight.ArrivalAirportId, flight.ArrivalTerminal, flight.ArrivalAt, flight.Code, flight.SeatLayout, flight.Class,
		).
		S(`WHERE "id" = $14 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "airline_id", "aircraft", "baggage", "cabin_baggage", "departure_airport_id", "departure_terminal", "departure_at", "arrival_airport_id", "arrival_terminal", "arrival_at", "code", "seat_layout", "class", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&flight.Id, &flight.AirlineId, &flight.Aircraft, &flight.Baggage, &flight.CabinBaggage, &flight.DepartureAirportId, &flight.DepartureTerminal, &flight.DepartureAt, &flight.ArrivalAirportId, &flight.ArrivalTerminal, &flight.ArrivalAt, &flight.Code, &flight.SeatLayout, &flight.Class,
		&flight.CreatedAt, &flight.UpdatedAt, &flight.DeletedAt,
	)

	return flight, err
}

func (r flightRepositoryPostgresql) Delete(ctx context.Context, id int64) (entity.Flight, error) {
	flight := entity.Flight{}

	builder := sqlbuilder.New().
		S(`UPDATE "flights" SET "deleted_at" = NOW()`).
		S(`WHERE "id" = $1 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "airline_id", "aircraft", "baggage", "cabin_baggage", "departure_airport_id", "departure_terminal", "departure_at", "arrival_airport_id", "arrival_terminal", "arrival_at", "code", "seat_layout", "class", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&flight.Id, &flight.AirlineId, &flight.Aircraft, &flight.Baggage, &flight.CabinBaggage, &flight.DepartureAirportId, &flight.DepartureTerminal, &flight.DepartureAt, &flight.ArrivalAirportId, &flight.ArrivalTerminal, &flight.ArrivalAt, &flight.Code, &flight.SeatLayout, &flight.Class,
		&flight.CreatedAt, &flight.UpdatedAt, &flight.DeletedAt,
	)

	return flight, err
}
