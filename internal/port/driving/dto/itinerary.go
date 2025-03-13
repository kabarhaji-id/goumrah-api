package dto

import (
	"time"

	"github.com/guregu/null/v5"
)

type ItineraryWidgetActivityRequest struct {
	Title       string
	Description string
	Images      []int64
}

func (r ItineraryWidgetActivityRequest) Type() string {
	return "Activity"
}

type ItineraryWidgetHotelRequest struct {
	Hotel int64
}

func (r ItineraryWidgetHotelRequest) Type() string {
	return "Hotel"
}

type ItineraryWidgetInformationRequest struct {
	Description string
}

func (r ItineraryWidgetInformationRequest) Type() string {
	return "Information"
}

type ItineraryWidgetTransportRequest struct {
	Transportation string
	From           string
	To             string
}

func (r ItineraryWidgetTransportRequest) Type() string {
	return "Transport"
}

type ItineraryWidgetRecommendationRequest struct {
	Description string
	Images      []int64
}

func (r ItineraryWidgetRecommendationRequest) Type() string {
	return "Recommendation"
}

type ItineraryWidgetRequest interface {
	Type() string
}

type ItineraryDayRequest struct {
	Title       string
	Description string
	Widgets     []ItineraryWidgetRequest
}

type ItineraryRequest struct {
	City   string
	Images []int64
	Days   []ItineraryDayRequest
}

type GetAllItineraryRequest struct {
	Page    int
	PerPage int
}

type ItineraryWidgetActivityResponse struct {
	Title       string
	Description string
	Images      []ImageResponse
}

func (r ItineraryWidgetActivityResponse) Type() string {
	return "Activity"
}

type ItineraryWidgetHotelResponse struct {
	Hotel HotelResponse
}

func (r ItineraryWidgetHotelResponse) Type() string {
	return "Hotel"
}

type ItineraryWidgetInformationResponse struct {
	Description string
}

func (r ItineraryWidgetInformationResponse) Type() string {
	return "Information"
}

type ItineraryWidgetTransportResponse struct {
	Transportation string
	From           string
	To             string
}

func (r ItineraryWidgetTransportResponse) Type() string {
	return "Transport"
}

type ItineraryWidgetRecommendationResponse struct {
	Description string
	Images      []ImageResponse
}

func (r ItineraryWidgetRecommendationResponse) Type() string {
	return "Recommendation"
}

type ItineraryWidgetResponse interface {
	Type() string
}

type ItineraryDayResponse struct {
	Title       string
	Description string
	Widgets     []ItineraryWidgetResponse
}

type ItineraryResponse struct {
	Id     int64
	City   string
	Images []ImageResponse
	Days   []ItineraryDayResponse

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}
