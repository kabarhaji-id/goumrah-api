package mapper

import (
	"context"
	"time"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driven/repository"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type FlightMapper struct {
	airlineMapper AirlineMapper
	airportMapper AirportMapper
}

func NewFlightMapper(airlineMapper AirlineMapper, airportMapper AirportMapper) FlightMapper {
	return FlightMapper{
		airlineMapper, airportMapper,
	}
}

func (FlightMapper) MapRequestToEntity(ctx context.Context, request dto.FlightRequest) entity.Flight {
	departureAt, _ := time.Parse("15:04", request.DepartureAt)
	arrivalAt, _ := time.Parse("15:04", request.ArrivalAt)

	return entity.Flight{
		AirlineId:          request.Airline,
		Aircraft:           request.Aircraft,
		Baggage:            request.Baggage,
		CabinBaggage:       request.CabinBaggage,
		DepartureAirportId: request.DepartureAirport,
		DepartureTerminal:  request.DepartureTerminal,
		DepartureAt:        departureAt,
		ArrivalAirportId:   request.ArrivalAirport,
		ArrivalTerminal:    request.ArrivalTerminal,
		ArrivalAt:          arrivalAt,
		Code:               request.Code,
		SeatLayout:         request.SeatLayout,
		Class:              request.Class,
	}
}

func (m FlightMapper) MapEntityToResponse(
	ctx context.Context,
	imageRepository repository.ImageRepository,
	airlineRepository repository.AirlineRepository,
	airportRepository repository.AirportRepository,
	flightEntity entity.Flight,
) (dto.FlightResponse, error) {
	airline, err := airlineRepository.FindById(ctx, flightEntity.AirlineId)
	if err != nil {
		return dto.FlightResponse{}, err
	}
	airlineResponse, err := m.airlineMapper.MapEntityToResponse(ctx, imageRepository, airline)
	if err != nil {
		return dto.FlightResponse{}, err
	}

	departureAirport, err := airportRepository.FindById(ctx, flightEntity.DepartureAirportId)
	if err != nil {
		return dto.FlightResponse{}, err
	}
	departureAirportResponse := m.airportMapper.MapEntityToResponse(ctx, departureAirport)

	arrivalAirport, err := airportRepository.FindById(ctx, flightEntity.ArrivalAirportId)
	if err != nil {
		return dto.FlightResponse{}, err
	}
	arrivalAirportResponse := m.airportMapper.MapEntityToResponse(ctx, arrivalAirport)

	var duration time.Duration
	if flightEntity.DepartureAt.Before(flightEntity.ArrivalAt) {
		departureAtMinute := flightEntity.DepartureAt.Hour()*60 + flightEntity.DepartureAt.Minute()
		arrivalAtMinute := flightEntity.ArrivalAt.Hour()*60 + flightEntity.ArrivalAt.Minute()
		duration = time.Minute * time.Duration(arrivalAtMinute-departureAtMinute)
	} else {
		departureAtMinute := flightEntity.DepartureAt.Hour()*60 + flightEntity.DepartureAt.Minute()
		arrivalAtMinute := flightEntity.ArrivalAt.Hour()*60 + flightEntity.ArrivalAt.Minute()
		duration = time.Minute * time.Duration((arrivalAtMinute+1440)-departureAtMinute)
	}

	return dto.FlightResponse{
		Id:                flightEntity.Id,
		Airline:           airlineResponse,
		Aircraft:          flightEntity.Aircraft,
		Baggage:           flightEntity.Baggage,
		CabinBaggage:      flightEntity.CabinBaggage,
		DepartureAirport:  departureAirportResponse,
		DepartureTerminal: flightEntity.DepartureTerminal,
		DepartureAt:       flightEntity.DepartureAt,
		ArrivalAirport:    arrivalAirportResponse,
		ArrivalTerminal:   flightEntity.ArrivalTerminal,
		ArrivalAt:         flightEntity.ArrivalAt,
		Duration:          duration,
		Code:              flightEntity.Code,
		SeatLayout:        flightEntity.SeatLayout,
		Class:             flightEntity.Class,
		CreatedAt:         flightEntity.CreatedAt,
		UpdatedAt:         flightEntity.UpdatedAt,
		DeletedAt:         flightEntity.DeletedAt,
	}, nil
}

func (m FlightMapper) MapEntitiesToResponses(
	ctx context.Context,
	imageRepository repository.ImageRepository,
	airlineRepository repository.AirlineRepository,
	airportRepository repository.AirportRepository,
	flightEntities []entity.Flight,
) ([]dto.FlightResponse, error) {
	flightResponses := make([]dto.FlightResponse, len(flightEntities))
	var err error

	for i, flightEntity := range flightEntities {
		flightResponses[i], err = m.MapEntityToResponse(
			ctx,
			imageRepository,
			airlineRepository,
			airportRepository,
			flightEntity,
		)
		if err != nil {
			return nil, err
		}
	}

	return flightResponses, nil
}
