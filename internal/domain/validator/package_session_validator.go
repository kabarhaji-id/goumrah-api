package validator

import (
	"context"
	"fmt"
	"time"

	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type PackageSessionValidator struct {
}

func NewPackageSessionValidator() PackageSessionValidator {
	return PackageSessionValidator{}
}

func (v PackageSessionValidator) ValidatePackageId(ctx context.Context, packageId int64) error {
	if packageId < 1 {
		return newError("PackageId", mustBeGte(1))
	}

	return nil
}

func (v PackageSessionValidator) ValidateRequest(ctx context.Context, request dto.PackageSessionRequest) error {
	if request.Embarkation < 1 {
		return newError("Embarkation", mustBeGte(1))
	}

	if len(request.DepartureDate) < 1 {
		return newError("DepartureDate", mustBeNotEmpty)
	}
	if _, err := time.Parse("02/01/2006", request.DepartureDate); err != nil {
		return newError("DepartureDate", invalidDate("DD/MM/YYYY"))
	}

	if request.Quota < 1 {
		return newError("Quota", mustBeGte(1))
	}

	if request.DoublePrice < 1 {
		return newError("DoublePrice", mustBeGte(1))
	}

	if request.DoubleFinalPrice.Valid {
		if request.DoubleFinalPrice.Float64 < 1 {
			return newError("DoubleFinalPrice", mustBeGte(1))
		}
	}

	if request.TriplePrice < 1 {
		return newError("TriplePrice", mustBeGte(1))
	}

	if request.TripleFinalPrice.Valid {
		if request.TripleFinalPrice.Float64 < 1 {
			return newError("TripleFinalPrice", mustBeGte(1))
		}
	}

	if request.QuadPrice < 1 {
		return newError("QuadPrice", mustBeGte(1))
	}

	if request.QuadFinalPrice.Valid {
		if request.QuadFinalPrice.Float64 < 1 {
			return newError("QuadFinalPrice", mustBeGte(1))
		}
	}

	if request.InfantPrice.Valid {
		if request.InfantPrice.Float64 < 1 {
			return newError("InfantPrice", mustBeGte(1))
		}
	}

	if request.InfantFinalPrice.Valid {
		if request.InfantFinalPrice.Float64 < 1 {
			return newError("InfantFinalPrice", mustBeGte(1))
		}
	}

	for i, departureFlight := range request.DepartureFlights {
		if departureFlight < 1 {
			return newError(fmt.Sprintf("DepartureFlights.%d", i), mustBeGte(1))
		}
	}

	if len(request.ReturnFlights) < 1 {
		return newError("ReturnFlights", mustBeNotEmpty)
	}
	for i, returnFlight := range request.ReturnFlights {
		if returnFlight < 1 {
			return newError(fmt.Sprintf("ReturnFlights.%d", i), mustBeGte(1))
		}
	}

	if len(request.DepartureFlights) < 1 {
		return newError("DepartureFlights", mustBeNotEmpty)
	}
	for i, guide := range request.Guides {
		if guide < 1 {
			return newError(fmt.Sprintf("Guides.%d", i), mustBeGte(1))
		}
	}

	if request.Bus < 1 {
		return newError("Bus", mustBeGte(1))
	}

	if len(request.Itineraries) < 1 {
		return newError("Itineraries", mustBeNotEmpty)
	}
	for itineraryIndex, itinerary := range request.Itineraries {
		itineraryCityLength := len(itinerary.City)
		if itineraryCityLength < 1 {
			return newError(fmt.Sprintf("Itineraries.%d.City", itineraryIndex), mustBeNotEmpty)
		}
		if itineraryCityLength > 100 {
			return newError(fmt.Sprintf("Itineraries.%d.City", itineraryIndex), maxChars(100))
		}

		for itineraryImageIndex, itineraryImage := range itinerary.Images {
			if itineraryImage < 1 {
				return newError(fmt.Sprintf("Itineraries.%d.Images.%d", itineraryIndex, itineraryImageIndex), mustBeGte(1))
			}
		}

		if len(itinerary.Days) < 1 {
			return newError(fmt.Sprintf("Itineraries.%d.Days", itineraryIndex), mustBeNotEmpty)
		}
		for itineraryDayIndex, itineraryDay := range itinerary.Days {
			itineraryDayTitleLength := len(itineraryDay.Title)
			if itineraryDayTitleLength < 1 {
				return newError(fmt.Sprintf("Itineraries.%d.Days.%d.Title", itineraryIndex, itineraryDayIndex), mustBeNotEmpty)
			}
			if itineraryDayTitleLength > 100 {
				return newError(fmt.Sprintf("Itineraries.%d.Days.%d.Title", itineraryIndex, itineraryDayIndex), maxChars(100))
			}

			itineraryDayDescriptionLength := len(itineraryDay.Description)
			if itineraryDayDescriptionLength < 1 {
				return newError(fmt.Sprintf("Itineraries.%d.Days.%d.Description", itineraryIndex, itineraryDayIndex), mustBeNotEmpty)
			}
			if itineraryDayDescriptionLength > 500 {
				return newError(fmt.Sprintf("Itineraries.%d.Days.%d.Description", itineraryIndex, itineraryDayIndex), maxChars(500))
			}

			for itineraryWidgetIndex, itineraryWidget := range itineraryDay.Widgets {
				switch itineraryWidget.(type) {
				case dto.ItineraryWidgetActivityRequest:
					activityWidget := itineraryWidget.(dto.ItineraryWidgetActivityRequest)

					widgetTitleLength := len(activityWidget.Title)
					if widgetTitleLength < 1 {
						return newError(fmt.Sprintf("Itineraries.%d.Days.%d.Widgets.%d.Title", itineraryIndex, itineraryDayIndex, itineraryWidgetIndex), mustBeNotEmpty)
					}
					if widgetTitleLength > 100 {
						return newError(fmt.Sprintf("Itineraries.%d.Days.%d.Widgets.%d.Title", itineraryIndex, itineraryDayIndex, itineraryWidgetIndex), maxChars(100))
					}

					widgetDescriptionLength := len(activityWidget.Description)
					if widgetDescriptionLength < 1 {
						return newError(fmt.Sprintf("Itineraries.%d.Days.%d.Widgets.%d.Description", itineraryIndex, itineraryDayIndex, itineraryWidgetIndex), mustBeNotEmpty)
					}
					if widgetDescriptionLength > 500 {
						return newError(fmt.Sprintf("Itineraries.%d.Days.%d.Widgets.%d.Description", itineraryIndex, itineraryDayIndex, itineraryWidgetIndex), maxChars(500))
					}

					for imageIndex, image := range activityWidget.Images {
						if image < 1 {
							return newError(fmt.Sprintf("Itineraries.%d.Days.%d.Widgets.%d.Images.%d", itineraryIndex, itineraryDayIndex, itineraryWidgetIndex, imageIndex), mustBeGte(1))
						}
					}
				case dto.ItineraryWidgetHotelRequest:
					hotelWidget := itineraryWidget.(dto.ItineraryWidgetHotelRequest)

					if hotelWidget.Hotel < 1 {
						return newError(fmt.Sprintf("Itineraries.%d.Days.%d.Widgets.%d.Hotel", itineraryIndex, itineraryDayIndex, itineraryWidgetIndex), mustBeGte(1))
					}
				case dto.ItineraryWidgetInformationRequest:
					informationWidget := itineraryWidget.(dto.ItineraryWidgetInformationRequest)

					widgetDescriptionLength := len(informationWidget.Description)
					if widgetDescriptionLength < 1 {
						return newError(fmt.Sprintf("Itineraries.%d.Days.%d.Widgets.%d.Description", itineraryIndex, itineraryDayIndex, itineraryWidgetIndex), mustBeNotEmpty)
					}
					if widgetDescriptionLength > 500 {
						return newError(fmt.Sprintf("Itineraries.%d.Days.%d.Widgets.%d.Description", itineraryIndex, itineraryDayIndex, itineraryWidgetIndex), maxChars(500))
					}
				case dto.ItineraryWidgetTransportRequest:
					transportWidget := itineraryWidget.(dto.ItineraryWidgetTransportRequest)

					widgetTransportationLength := len(transportWidget.Transportation)
					if widgetTransportationLength < 1 {
						return newError(fmt.Sprintf("Itineraries.%d.Days.%d.Widgets.%d.Transportation", itineraryIndex, itineraryDayIndex, itineraryWidgetIndex), mustBeNotEmpty)
					}
					if widgetTransportationLength > 100 {
						return newError(fmt.Sprintf("Itineraries.%d.Days.%d.Widgets.%d.Transportation", itineraryIndex, itineraryDayIndex, itineraryWidgetIndex), maxChars(100))
					}

					widgetFromLength := len(transportWidget.From)
					if widgetFromLength < 1 {
						return newError(fmt.Sprintf("Itineraries.%d.Days.%d.Widgets.%d.From", itineraryIndex, itineraryDayIndex, itineraryWidgetIndex), mustBeNotEmpty)
					}
					if widgetFromLength > 100 {
						return newError(fmt.Sprintf("Itineraries.%d.Days.%d.Widgets.%d.From", itineraryIndex, itineraryDayIndex, itineraryWidgetIndex), maxChars(100))
					}

					widgetToLength := len(transportWidget.To)
					if widgetToLength < 1 {
						return newError(fmt.Sprintf("Itineraries.%d.Days.%d.Widgets.%d.To", itineraryIndex, itineraryDayIndex, itineraryWidgetIndex), mustBeNotEmpty)
					}
					if widgetToLength > 100 {
						return newError(fmt.Sprintf("Itineraries.%d.Days.%d.Widgets.%d.To", itineraryIndex, itineraryDayIndex, itineraryWidgetIndex), maxChars(100))
					}
				case dto.ItineraryWidgetRecommendationRequest:
					recommendationWidget := itineraryWidget.(dto.ItineraryWidgetRecommendationRequest)

					widgetDescriptionLength := len(recommendationWidget.Description)
					if widgetDescriptionLength < 1 {
						return newError(fmt.Sprintf("Itineraries.%d.Days.%d.Widgets.%d.Description", itineraryIndex, itineraryDayIndex, itineraryWidgetIndex), mustBeNotEmpty)
					}
					if widgetDescriptionLength > 500 {
						return newError(fmt.Sprintf("Itineraries.%d.Days.%d.Widgets.%d.Description", itineraryIndex, itineraryDayIndex, itineraryWidgetIndex), maxChars(500))
					}

					for imageIndex, image := range recommendationWidget.Images {
						if image < 1 {
							return newError(fmt.Sprintf("Itineraries.%d.Days.%d.Widgets.%d.Images.%d", itineraryIndex, itineraryDayIndex, itineraryWidgetIndex, imageIndex), mustBeGte(1))
						}
					}
				}
			}
		}
	}

	return nil
}

func (v PackageSessionValidator) ValidateId(ctx context.Context, id int64) error {
	if id < 1 {
		return newError("id", mustBeGte(1))
	}

	return nil
}

func (v PackageSessionValidator) ValidateGetAllRequest(ctx context.Context, request dto.GetAllPackageSessionRequest) error {
	if request.Package.Valid {
		if request.Package.Int64 < 1 {
			return newError("Package", mustBeGte(1))
		}
	}

	if request.Page < 1 {
		return newError("Page", mustBeGte(1))
	}

	if request.PerPage < 1 {
		return newError("PerPage", mustBeGte(1))
	}

	return nil
}
