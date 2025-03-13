package schema

import (
	"time"

	"github.com/guregu/null/v5"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type ItineraryWidgetRequest struct {
	Type           string  `json:"type"`
	Title          string  `json:"title"`
	Description    string  `json:"description"`
	Hotel          int64   `json:"hotel"`
	Transportation string  `json:"transportation"`
	From           string  `json:"from"`
	To             string  `json:"to"`
	Images         []int64 `json:"images"`
}

func (r ItineraryWidgetRequest) ToDtoRequest() dto.ItineraryWidgetRequest {
	switch r.Type {
	case "Activity":
		return dto.ItineraryWidgetActivityRequest{
			Title:       r.Title,
			Description: r.Description,
			Images:      r.Images,
		}
	case "Hotel":
		return dto.ItineraryWidgetHotelRequest{
			Hotel: r.Hotel,
		}
	case "Information":
		return dto.ItineraryWidgetInformationRequest{
			Description: r.Description,
		}
	case "Transport":
		return dto.ItineraryWidgetTransportRequest{
			Transportation: r.Transportation,
			From:           r.From,
			To:             r.To,
		}
	case "Recommendation":
		return dto.ItineraryWidgetRecommendationRequest{
			Description: r.Description,
			Images:      r.Images,
		}
	default:
		return nil
	}
}

type ItineraryDayRequest struct {
	Title       string                   `json:"title"`
	Description string                   `json:"description"`
	Widgets     []ItineraryWidgetRequest `json:"widgets"`
}

func (r ItineraryDayRequest) ToDtoRequest() dto.ItineraryDayRequest {
	widgets := make([]dto.ItineraryWidgetRequest, len(r.Widgets))
	for i, widget := range r.Widgets {
		widgets[i] = widget.ToDtoRequest()
	}

	return dto.ItineraryDayRequest{
		Title:       r.Title,
		Description: r.Description,
		Widgets:     widgets,
	}
}

type ItineraryRequest struct {
	City   string                `json:"city"`
	Images []int64               `json:"images"`
	Days   []ItineraryDayRequest `json:"days"`
}

func (r ItineraryRequest) ToDtoRequest() dto.ItineraryRequest {
	days := make([]dto.ItineraryDayRequest, len(r.Days))
	for i, day := range r.Days {
		days[i] = day.ToDtoRequest()
	}

	return dto.ItineraryRequest{
		City:   r.City,
		Images: r.Images,
		Days:   days,
	}
}

type ItineraryWidgetActivityResponse struct {
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Images      []ImageResponse `json:"images"`
}

func (r ItineraryWidgetActivityResponse) Type() string {
	return "Activity"
}

type ItineraryWidgetHotelResponse struct {
	Hotel HotelResponse `json:"hotel"`
}

func (r ItineraryWidgetHotelResponse) Type() string {
	return "Hotel"
}

type ItineraryWidgetInformationResponse struct {
	Description string `json:"description"`
}

func (r ItineraryWidgetInformationResponse) Type() string {
	return "Information"
}

type ItineraryWidgetTransportResponse struct {
	Transportation string `json:"transportation"`
	From           string `json:"from"`
	To             string `json:"to"`
}

func (r ItineraryWidgetTransportResponse) Type() string {
	return "Transport"
}

type ItineraryWidgetRecommendationResponse struct {
	Description string          `json:"description"`
	Images      []ImageResponse `json:"images"`
}

func (r ItineraryWidgetRecommendationResponse) Type() string {
	return "Recommendation"
}

type ItineraryWidgetResponse interface {
	Type() string
}

func NewItineraryWidgetResponse(dtoResponse dto.ItineraryWidgetResponse) ItineraryWidgetResponse {
	switch dtoResponse.Type() {
	case "Activity":
		activityDtoResponse := dtoResponse.(dto.ItineraryWidgetActivityResponse)
		return ItineraryWidgetActivityResponse{
			Title:       activityDtoResponse.Title,
			Description: activityDtoResponse.Description,
			Images:      NewImageResponses(activityDtoResponse.Images),
		}
	case "Hotel":
		hotelDtoResponse := dtoResponse.(dto.ItineraryWidgetHotelResponse)
		return ItineraryWidgetHotelResponse{
			Hotel: NewHotelResponse(hotelDtoResponse.Hotel),
		}
	case "Information":
		informationDtoResponse := dtoResponse.(dto.ItineraryWidgetInformationResponse)
		return ItineraryWidgetInformationResponse{
			Description: informationDtoResponse.Description,
		}
	case "Transport":
		transporDtoResponse := dtoResponse.(dto.ItineraryWidgetTransportResponse)
		return ItineraryWidgetTransportResponse{
			Transportation: transporDtoResponse.Transportation,
			From:           transporDtoResponse.From,
			To:             transporDtoResponse.To,
		}
	case "Recommendation":
		recommendationDtoResponse := dtoResponse.(dto.ItineraryWidgetRecommendationResponse)
		return ItineraryWidgetRecommendationResponse{
			Description: recommendationDtoResponse.Description,
			Images:      NewImageResponses(recommendationDtoResponse.Images),
		}
	default:
		return nil
	}
}

func NewItineraryWidgetResponses(dtoResponses []dto.ItineraryWidgetResponse) []ItineraryWidgetResponse {
	widgetResponses := make([]ItineraryWidgetResponse, len(dtoResponses))

	for i, dtoResponse := range dtoResponses {
		widgetResponses[i] = NewItineraryWidgetResponse(dtoResponse)
	}

	return widgetResponses
}

type ItineraryDayResponse struct {
	Title       string                    `json:"title"`
	Description string                    `json:"description"`
	Widgets     []ItineraryWidgetResponse `json:"widgets"`
}

func NewItineraryDayResponse(dtoResponse dto.ItineraryDayResponse) ItineraryDayResponse {
	return ItineraryDayResponse{
		Title:       dtoResponse.Title,
		Description: dtoResponse.Description,
		Widgets:     NewItineraryWidgetResponses(dtoResponse.Widgets),
	}
}

func NewItineraryDayResponses(dtoResponses []dto.ItineraryDayResponse) []ItineraryDayResponse {
	dayResponses := make([]ItineraryDayResponse, len(dtoResponses))

	for i, dtoResponse := range dtoResponses {
		dayResponses[i] = NewItineraryDayResponse(dtoResponse)
	}

	return dayResponses
}

type ItineraryResponse struct {
	Id     int64                  `json:"id"`
	City   string                 `json:"city"`
	Images []ImageResponse        `json:"images"`
	Days   []ItineraryDayResponse `json:"days"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt null.Time `json:"deleted_at"`
}

func NewItineraryResponse(dtoResponse dto.ItineraryResponse) ItineraryResponse {
	return ItineraryResponse{
		Id:        dtoResponse.Id,
		City:      dtoResponse.City,
		Images:    NewImageResponses(dtoResponse.Images),
		Days:      NewItineraryDayResponses(dtoResponse.Days),
		CreatedAt: dtoResponse.CreatedAt,
		UpdatedAt: dtoResponse.UpdatedAt,
		DeletedAt: dtoResponse.DeletedAt,
	}
}

func NewItineraryResponses(dtoResponses []dto.ItineraryResponse) []ItineraryResponse {
	itineraryResponses := make([]ItineraryResponse, len(dtoResponses))

	for i, dtoResponse := range dtoResponses {
		itineraryResponses[i] = NewItineraryResponse(dtoResponse)
	}

	return itineraryResponses
}
