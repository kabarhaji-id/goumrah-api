package repository

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
)

type ItineraryWidgetRepository interface {
	Create(ctx context.Context, itineraryWidget entity.ItineraryWidget) (entity.ItineraryWidget, error)
	FindById(ctx context.Context, id int64) (entity.ItineraryWidget, error)
	FindAll(ctx context.Context, opt FindAllOptions) ([]entity.ItineraryWidget, error)
	Update(ctx context.Context, id int64, itineraryWidget entity.ItineraryWidget) (entity.ItineraryWidget, error)
	Delete(ctx context.Context, id int64) (entity.ItineraryWidget, error)
}

type ItineraryWidgetActivityRepository interface {
	Create(ctx context.Context, itineraryWidgetActivity entity.ItineraryWidgetActivity) (entity.ItineraryWidgetActivity, error)
	FindById(ctx context.Context, id int64) (entity.ItineraryWidgetActivity, error)
	FindAll(ctx context.Context, opt FindAllOptions) ([]entity.ItineraryWidgetActivity, error)
	Update(ctx context.Context, id int64, itineraryWidgetActivity entity.ItineraryWidgetActivity) (entity.ItineraryWidgetActivity, error)
	Delete(ctx context.Context, id int64) (entity.ItineraryWidgetActivity, error)

	AttachImages(ctx context.Context, id int64, imageIds []int64) ([]int64, error)
	FindImages(ctx context.Context, id int64) ([]entity.Image, error)
	FindImageIds(ctx context.Context, id int64) ([]int64, error)
	DetachImages(ctx context.Context, id int64) ([]int64, error)
}

type ItineraryWidgetHotelRepository interface {
	Create(ctx context.Context, itineraryWidgetHotel entity.ItineraryWidgetHotel) (entity.ItineraryWidgetHotel, error)
	FindById(ctx context.Context, id int64) (entity.ItineraryWidgetHotel, error)
	FindAll(ctx context.Context, opt FindAllOptions) ([]entity.ItineraryWidgetHotel, error)
	Update(ctx context.Context, id int64, itineraryWidgetHotel entity.ItineraryWidgetHotel) (entity.ItineraryWidgetHotel, error)
	Delete(ctx context.Context, id int64) (entity.ItineraryWidgetHotel, error)
}

type ItineraryWidgetInformationRepository interface {
	Create(ctx context.Context, itineraryWidgetInformation entity.ItineraryWidgetInformation) (entity.ItineraryWidgetInformation, error)
	FindById(ctx context.Context, id int64) (entity.ItineraryWidgetInformation, error)
	FindAll(ctx context.Context, opt FindAllOptions) ([]entity.ItineraryWidgetInformation, error)
	Update(ctx context.Context, id int64, itineraryWidgetInformation entity.ItineraryWidgetInformation) (entity.ItineraryWidgetInformation, error)
	Delete(ctx context.Context, id int64) (entity.ItineraryWidgetInformation, error)
}

type ItineraryWidgetTransportRepository interface {
	Create(ctx context.Context, itineraryWidgetTransport entity.ItineraryWidgetTransport) (entity.ItineraryWidgetTransport, error)
	FindById(ctx context.Context, id int64) (entity.ItineraryWidgetTransport, error)
	FindAll(ctx context.Context, opt FindAllOptions) ([]entity.ItineraryWidgetTransport, error)
	Update(ctx context.Context, id int64, itineraryWidgetTransport entity.ItineraryWidgetTransport) (entity.ItineraryWidgetTransport, error)
	Delete(ctx context.Context, id int64) (entity.ItineraryWidgetTransport, error)
}

type ItineraryWidgetRecommendationRepository interface {
	Create(ctx context.Context, itineraryWidgetRecommendation entity.ItineraryWidgetRecommendation) (entity.ItineraryWidgetRecommendation, error)
	FindById(ctx context.Context, id int64) (entity.ItineraryWidgetRecommendation, error)
	FindAll(ctx context.Context, opt FindAllOptions) ([]entity.ItineraryWidgetRecommendation, error)
	Update(ctx context.Context, id int64, itineraryWidgetRecommendation entity.ItineraryWidgetRecommendation) (entity.ItineraryWidgetRecommendation, error)
	Delete(ctx context.Context, id int64) (entity.ItineraryWidgetRecommendation, error)

	AttachImages(ctx context.Context, id int64, imageIds []int64) ([]int64, error)
	FindImages(ctx context.Context, id int64) ([]entity.Image, error)
	FindImageIds(ctx context.Context, id int64) ([]int64, error)
	DetachImages(ctx context.Context, id int64) ([]int64, error)
}
