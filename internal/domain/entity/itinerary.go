package entity

import (
	"time"

	"github.com/guregu/null/v5"
)

type ItineraryWidgetActivity struct {
	Id          int64
	Title       string
	Description string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}

func (ItineraryWidgetActivity) Type() string {
	return "activity"
}

type ItineraryWidgetHotel struct {
	Id      int64
	HotelId int64

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}

func (ItineraryWidgetHotel) Type() string {
	return "hotel"
}

type ItineraryWidgetInformation struct {
	Id          int64
	Description string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}

func (ItineraryWidgetInformation) Type() string {
	return "information"
}

type ItineraryWidgetTransport struct {
	Id             int64
	Transportation string
	From           string
	To             string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}

func (ItineraryWidgetTransport) Type() string {
	return "transport"
}

type ItineraryWidgetRecommendation struct {
	Id          int64
	Description string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}

func (ItineraryWidgetRecommendation) Type() string {
	return "recommendation"
}

type ItineraryWidgetInterface interface {
	Type() string
}

type ItineraryWidget struct {
	Id               int64
	ActivityId       null.Int64
	HotelId          null.Int64
	InformationId    null.Int64
	TransportId      null.Int64
	RecommendationId null.Int64
	NextId           null.Int64

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}

type ItineraryDay struct {
	Id          int64
	Title       string
	Description string
	WidgetId    null.Int64
	NextId      null.Int64

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}

type Itinerary struct {
	Id     int64
	City   string
	DayId  int64
	NextId null.Int64

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}
