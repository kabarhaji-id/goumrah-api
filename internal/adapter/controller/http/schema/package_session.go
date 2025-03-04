package schema

import (
	"time"

	"github.com/guregu/null/v5"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type PackageSessionRequest struct {
	Embarkation   int64   `json:"embarkation"`
	DepartureDate string  `json:"departure_date"`
	Guides        []int64 `json:"guides"`
}

func (r PackageSessionRequest) ToDtoRequest() dto.PackageSessionRequest {
	return dto.PackageSessionRequest{
		Embarkation:   r.Embarkation,
		DepartureDate: r.DepartureDate,
		Guides:        r.Guides,
	}
}

type GetAllPackageSessionQuery struct {
	Page    null.Int `query:"page"`
	PerPage null.Int `query:"per_page"`
}

type PackageSessionParams struct {
	Id int64 `params:"id"`
}

type PackageSessionResponse struct {
	Id            int64               `json:"id"`
	Package       int64               `json:"package"`
	Embarkation   EmbarkationResponse `json:"embarkation"`
	DepartureDate time.Time           `json:"departure_date"`
	Guides        []GuideResponse     `json:"guides"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt null.Time `json:"deleted_at"`
}

func NewPackageSessionResponse(dtoResponse dto.PackageSessionResponse) PackageSessionResponse {
	embarkation := NewEmbarkationResponse(dtoResponse.Embarkation)
	guides := NewGuideResponses(dtoResponse.Guides)

	return PackageSessionResponse{
		Id:            dtoResponse.Id,
		Package:       dtoResponse.Id,
		Embarkation:   embarkation,
		DepartureDate: dtoResponse.DepartureDate,
		Guides:        guides,
		CreatedAt:     dtoResponse.CreatedAt,
		UpdatedAt:     dtoResponse.UpdatedAt,
		DeletedAt:     dtoResponse.DeletedAt,
	}
}

func NewPackageSessionResponses(dtoResponses []dto.PackageSessionResponse) []PackageSessionResponse {
	packageSessionResponses := make([]PackageSessionResponse, len(dtoResponses))

	for i, dtoResponse := range dtoResponses {
		packageSessionResponses[i] = NewPackageSessionResponse(dtoResponse)
	}

	return packageSessionResponses
}

type PackageSessionListResponse struct {
	Id            int64               `json:"id"`
	Package       int64               `json:"package"`
	Embarkation   EmbarkationResponse `json:"embarkation"`
	DepartureDate time.Time           `json:"departure_date"`
	Guides        []int64             `json:"guides"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt null.Time `json:"deleted_at"`
}

func NewPackageSessionListResponse(dtoResponse dto.PackageSessionListResponse) PackageSessionListResponse {
	embarkation := NewEmbarkationResponse(dtoResponse.Embarkation)

	return PackageSessionListResponse{
		Id:            dtoResponse.Id,
		Package:       dtoResponse.Id,
		Embarkation:   embarkation,
		DepartureDate: dtoResponse.DepartureDate,
		Guides:        dtoResponse.Guides,
		CreatedAt:     dtoResponse.CreatedAt,
		UpdatedAt:     dtoResponse.UpdatedAt,
		DeletedAt:     dtoResponse.DeletedAt,
	}
}

func NewPackageSessionListResponses(dtoResponses []dto.PackageSessionListResponse) []PackageSessionListResponse {
	packageSessionListResponses := make([]PackageSessionListResponse, len(dtoResponses))

	for i, dtoResponse := range dtoResponses {
		packageSessionListResponses[i] = NewPackageSessionListResponse(dtoResponse)
	}

	return packageSessionListResponses
}
