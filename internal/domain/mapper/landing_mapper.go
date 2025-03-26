package mapper

import (
	"context"
	"fmt"
	"slices"
	"time"

	"github.com/guregu/null/v6"
	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driven/repository"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type LandingMapper struct {
	imageMapper ImageMapper
}

func NewLandingMapper(imageMapper ImageMapper) LandingMapper {
	return LandingMapper{imageMapper}
}

func (m LandingMapper) mapPackageItemEntityToResponse(
	ctx context.Context,
	packageRepository repository.PackageRepository,
	imageRepository repository.ImageRepository,
	packageSessionRepository repository.PackageSessionRepository,
	itineraryRepository repository.ItineraryRepository,
	itineraryDayRepository repository.ItineraryDayRepository,
	flightRouteRepository repository.FlightRouteRepository,
	flightRepository repository.FlightRepository,
	airlineRepository repository.AirlineRepository,
	landingPackageItem entity.LandingPackageItem,
) (dto.LandingPackageItemResponse, error) {
	packageEntity, err := packageRepository.FindById(ctx, landingPackageItem.PackageId)
	if err != nil {
		return dto.LandingPackageItemResponse{}, err
	}

	packageThumbnailResponse := null.NewValue(dto.ImageResponse{}, false)
	if packageEntity.ThumbnailId.Valid {
		packageThumbnailEntity, err := imageRepository.FindById(ctx, packageEntity.ThumbnailId.Int64)
		if err != nil {
			return dto.LandingPackageItemResponse{}, err
		}

		packageThumbnailResponse = null.ValueFrom(m.imageMapper.MapEntityToResponse(ctx, packageThumbnailEntity))
	}

	packageSessionEntities, err := packageSessionRepository.FindAll(ctx, repository.FindAllOptions{
		Where: map[string]any{
			"package_id": packageEntity.Id,
		},
	})
	if err != nil {
		return dto.LandingPackageItemResponse{}, err
	}

	packageSessionQuadEntity := packageSessionEntities[0]
	for _, packageSessionEntity := range packageSessionEntities[1:] {
		if (packageSessionEntity.QuadFinalPrice.Valid && packageSessionQuadEntity.QuadFinalPrice.Valid && packageSessionEntity.QuadFinalPrice.Float64 < packageSessionQuadEntity.QuadFinalPrice.Float64) ||
			(packageSessionEntity.QuadPrice < packageSessionQuadEntity.QuadPrice) {
			packageSessionQuadEntity = packageSessionEntity
		}
	}

	itineraryDayEntities := []entity.ItineraryDay{}
	for itineraryId := null.NewInt(packageSessionQuadEntity.ItineraryId, true); itineraryId.Valid; {
		itineraryEntity, err := itineraryRepository.FindById(ctx, itineraryId.Int64)
		if err != nil {
			return dto.LandingPackageItemResponse{}, err
		}

		for itineraryDayId := null.NewInt(itineraryEntity.DayId, true); itineraryDayId.Valid; {
			itineraryDayEntity, err := itineraryDayRepository.FindById(ctx, itineraryDayId.Int64)
			if err != nil {
				return dto.LandingPackageItemResponse{}, err
			}
			itineraryDayEntities = append(itineraryDayEntities, itineraryDayEntity)

			itineraryDayId = itineraryDayEntity.NextId
		}

		itineraryId = itineraryEntity.NextId
	}

	tags := []dto.LandingPackageItemTagResponse{
		{
			Icon:  "clock",
			Label: fmt.Sprintf("%d Hari", len(itineraryDayEntities)),
		},
		{
			Icon:  "la-kaaba",
			Label: string(packageEntity.Type),
		},
	}

	departureFlightRouteEntity, err := flightRouteRepository.FindById(ctx, packageSessionQuadEntity.DepartureFlightRouteId)
	if err != nil {
		return dto.LandingPackageItemResponse{}, err
	}
	if departureFlightRouteEntity.NextId.Valid {
		tags = append(tags, dto.LandingPackageItemTagResponse{
			Icon:  "plane",
			Label: "Transit",
		})
	} else {
		tags = append(tags, dto.LandingPackageItemTagResponse{
			Icon:  "plane",
			Label: "Langsung",
		})
	}

	if packageEntity.FastTrain {
		tags = append(tags, dto.LandingPackageItemTagResponse{
			Icon:  "train",
			Label: "Kereta Cepat",
		})
	}

	departureFlightEntity, err := flightRepository.FindById(ctx, departureFlightRouteEntity.FlightId)
	if err != nil {
		return dto.LandingPackageItemResponse{}, err
	}

	departureAirlineEntity, err := airlineRepository.FindById(ctx, departureFlightEntity.AirlineId)
	if err != nil {
		return dto.LandingPackageItemResponse{}, err
	}

	departureDate := []dto.LandingPackageItemDepartureDateResponse{}
	for _, packageSessionEntity := range packageSessionEntities {
		if !slices.ContainsFunc(departureDate, func(item dto.LandingPackageItemDepartureDateResponse) bool {
			return item.Date.Equal(packageSessionEntity.DepartureDate)
		}) {
			status := "active"
			if packageSessionEntity.DepartureDate.Before(time.Now()) {
				status = "expired"
			}

			departureDate = append(departureDate, dto.LandingPackageItemDepartureDateResponse{
				Date:   packageSessionEntity.DepartureDate,
				Status: status,
			})
		}
	}

	return dto.LandingPackageItemResponse{
		IsEnabled:     landingPackageItem.IsEnabled,
		IsMobile:      landingPackageItem.IsMobile,
		IsDesktop:     landingPackageItem.IsDesktop,
		Id:            packageEntity.Id,
		Thumbnail:     packageThumbnailResponse,
		Tags:          tags,
		Title:         packageEntity.Name,
		DepartureDate: departureDate,
		Details: []dto.LandingPackageItemDetailResponse{
			{
				Icon:    "airline",
				Label:   "Maskapai",
				Value:   departureAirlineEntity.Name,
				AltText: "Airline",
				Rating:  departureAirlineEntity.SkytraxRating,
			},
			{
				Icon:    "hotel",
				Label:   "Madinah",
				Value:   "",
				AltText: "Hotel",
				Rating:  0,
			},
			{
				Icon:    "hotel",
				Label:   "Makkah",
				Value:   "",
				AltText: "Hotel",
				Rating:  0,
			},
		},
		Price: dto.LandingPackageItemPriceResponse{
			DoublePrice:      packageSessionQuadEntity.DoublePrice,
			DoubleFinalPrice: packageSessionQuadEntity.DoubleFinalPrice,
			TriplePrice:      packageSessionQuadEntity.TriplePrice,
			TripleFinalPrice: packageSessionQuadEntity.TripleFinalPrice,
			QuadPrice:        packageSessionQuadEntity.QuadPrice,
			QuadFinalPrice:   packageSessionQuadEntity.QuadFinalPrice,
			InfantPrice:      packageSessionQuadEntity.InfantPrice,
			InfantFinalPrice: packageSessionQuadEntity.InfantFinalPrice,
		},
		ButtonLabel: landingPackageItem.ButtonLabel,
		Category:    packageEntity.Category,
	}, nil
}

func (m LandingMapper) mapPackageItemEntitiesToResponses(
	ctx context.Context,
	packageRepository repository.PackageRepository,
	imageRepository repository.ImageRepository,
	packageSessionRepository repository.PackageSessionRepository,
	itineraryRepository repository.ItineraryRepository,
	itineraryDayRepository repository.ItineraryDayRepository,
	flightRouteRepository repository.FlightRouteRepository,
	flightRepository repository.FlightRepository,
	airlineRepository repository.AirlineRepository,
	landingPackageItems []entity.LandingPackageItem,
) ([]dto.LandingPackageItemResponse, error) {
	landingPackageItemResponses := make([]dto.LandingPackageItemResponse, len(landingPackageItems))
	var err error

	for i, landingPackageItem := range landingPackageItems {
		landingPackageItemResponses[i], err = m.mapPackageItemEntityToResponse(
			ctx,
			packageRepository,
			imageRepository,
			packageSessionRepository,
			itineraryRepository,
			itineraryDayRepository,
			flightRouteRepository,
			flightRepository,
			airlineRepository,
			landingPackageItem,
		)
		if err != nil {
			return nil, err
		}
	}

	return landingPackageItemResponses, nil
}

func (m LandingMapper) mapPackageDetailEntityToResponse(
	ctx context.Context,
	packageRepository repository.PackageRepository,
	imageRepository repository.ImageRepository,
	packageSessionRepository repository.PackageSessionRepository,
	itineraryRepository repository.ItineraryRepository,
	itineraryDayRepository repository.ItineraryDayRepository,
	flightRouteRepository repository.FlightRouteRepository,
	flightRepository repository.FlightRepository,
	airlineRepository repository.AirlineRepository,
	landingSectionHeaderRepository repository.LandingSectionHeaderRepository,
	landingPackageDetailItemRepository repository.LandingPackageDetailItemRepository,
	landingPackageItemRepository repository.LandingPackageItemRepository,
	landingPackageDetail entity.LandingPackageDetail,
) (dto.LandingPackageDetailResponse, error) {
	landingSectionHeaderEntity, err := landingSectionHeaderRepository.FindById(ctx, landingPackageDetail.LandingSectionHeaderId)
	if err != nil {
		return dto.LandingPackageDetailResponse{}, err
	}

	landingPackageDetailItemEntities, err := landingPackageDetailItemRepository.FindAll(ctx, repository.FindAllOptions{})
	if err != nil {
		return dto.LandingPackageDetailResponse{}, err
	}
	landingPackageItemEntities := make([]entity.LandingPackageItem, len(landingPackageDetailItemEntities))
	for index, landingPackageDetailItemEntity := range landingPackageDetailItemEntities {
		landingPackageItemEntity, err := landingPackageItemRepository.FindById(ctx, landingPackageDetailItemEntity.LandingPackageItemId)
		if err != nil {
			return dto.LandingPackageDetailResponse{}, err
		}

		landingPackageItemEntities[index] = landingPackageItemEntity
	}

	landingPackageItemResponses, err := m.mapPackageItemEntitiesToResponses(
		ctx,
		packageRepository,
		imageRepository,
		packageSessionRepository,
		itineraryRepository,
		itineraryDayRepository,
		flightRouteRepository,
		flightRepository,
		airlineRepository,
		landingPackageItemEntities,
	)
	if err != nil {
		return dto.LandingPackageDetailResponse{}, err
	}

	return dto.LandingPackageDetailResponse{
		IsEnabled: landingPackageDetail.IsEnabled,
		IsMobile:  landingPackageDetail.IsMobile,
		IsDesktop: landingPackageDetail.IsDesktop,
		Header: dto.LandingSectionHeaderResponse{
			IsEnabled: landingSectionHeaderEntity.IsEnabled,
			IsMobile:  landingSectionHeaderEntity.IsMobile,
			IsDesktop: landingSectionHeaderEntity.IsDesktop,
			Title:     landingSectionHeaderEntity.Title,
			Subtitle:  landingSectionHeaderEntity.Subtitle,
			TagsLine:  landingSectionHeaderEntity.TagsLine,
		},
		Packages: landingPackageItemResponses,
	}, nil
}

func (m LandingMapper) MapEntityToResponse(
	ctx context.Context,
	landingSectionHeaderRepository repository.LandingSectionHeaderRepository,
	landingPackageItemRepository repository.LandingPackageItemRepository,
	landingPackageDetailRepository repository.LandingPackageDetailRepository,
	landingPackageDetailItemRepository repository.LandingPackageDetailItemRepository,
	landingFeaturesContentBenefitRepository repository.LandingFeaturesContentBenefitRepository,
	landingMomentsContentImageRepository repository.LandingMomentsContentImageRepository,
	landingAffiliatesContentAffiliateRepository repository.LandingAffiliatesContentAffiliateRepository,
	LandingFaqContentFaqRepository repository.LandingFaqContentFaqRepository,
	imageRepository repository.ImageRepository,
	packageRepository repository.PackageRepository,
	packageSessionRepository repository.PackageSessionRepository,
	itineraryRepository repository.ItineraryRepository,
	itineraryDayRepository repository.ItineraryDayRepository,
	flightRouteRepository repository.FlightRouteRepository,
	flightRepository repository.FlightRepository,
	airlineRepository repository.AirlineRepository,
	landingHeroContentEntity entity.LandingHeroContent,
	landingSinglePackageContentEntity entity.LandingSinglePackageContent,
	landingPackagesContentEntity entity.LandingPackagesContent,
	landingFeaturesContentEntity entity.LandingFeaturesContent,
	landingMomentsContentEntity entity.LandingMomentsContent,
	landingAffiliatesContentEntity entity.LandingAffiliatesContent,
	landingFaqContentEntity entity.LandingFaqContent,
	landingMenuEntities []entity.LandingMenu,
) (dto.LandingResponse, error) {
	landingHeroContentImageResponse := null.NewValue(dto.ImageResponse{}, false)
	if landingHeroContentEntity.ImageId.Valid {
		landingHeroContentImageEntity, err := imageRepository.FindById(ctx, landingHeroContentEntity.ImageId.Int64)
		if err != nil {
			return dto.LandingResponse{}, err
		}

		landingHeroContentImageResponse = null.ValueFrom(m.imageMapper.MapEntityToResponse(ctx, landingHeroContentImageEntity))
	}

	landingSinglePackageContentHeaderEntity, err := landingSectionHeaderRepository.FindById(ctx, landingSinglePackageContentEntity.LandingSectionHeaderId)
	if err != nil {
		return dto.LandingResponse{}, err
	}

	landingSinglePackageContentSilverResponse := null.NewValue(dto.LandingPackageItemResponse{}, false)
	if landingSinglePackageContentEntity.SilverLandingPackageItemId.Valid {
		landingSinglePackageContentSilverEntity, err := landingPackageItemRepository.FindById(ctx, landingSinglePackageContentEntity.SilverLandingPackageItemId.Int64)
		if err != nil {
			return dto.LandingResponse{}, err
		}

		landingSinglePackageContentSilverResponseValue, err := m.mapPackageItemEntityToResponse(
			ctx,
			packageRepository,
			imageRepository,
			packageSessionRepository,
			itineraryRepository,
			itineraryDayRepository,
			flightRouteRepository,
			flightRepository,
			airlineRepository,
			landingSinglePackageContentSilverEntity,
		)
		if err != nil {
			return dto.LandingResponse{}, err
		}

		landingSinglePackageContentSilverResponse = null.ValueFrom(landingSinglePackageContentSilverResponseValue)
	}

	landingSinglePackageContentGoldResponse := null.NewValue(dto.LandingPackageItemResponse{}, false)
	if landingSinglePackageContentEntity.GoldLandingPackageItemId.Valid {
		landingSinglePackageContentGoldEntity, err := landingPackageItemRepository.FindById(ctx, landingSinglePackageContentEntity.GoldLandingPackageItemId.Int64)
		if err != nil {
			return dto.LandingResponse{}, err
		}

		landingSinglePackageContentGoldResponseValue, err := m.mapPackageItemEntityToResponse(
			ctx,
			packageRepository,
			imageRepository,
			packageSessionRepository,
			itineraryRepository,
			itineraryDayRepository,
			flightRouteRepository,
			flightRepository,
			airlineRepository,
			landingSinglePackageContentGoldEntity,
		)
		if err != nil {
			return dto.LandingResponse{}, err
		}

		landingSinglePackageContentGoldResponse = null.ValueFrom(landingSinglePackageContentGoldResponseValue)
	}

	landingSinglePackageContentPlatinumResponse := null.NewValue(dto.LandingPackageItemResponse{}, false)
	if landingSinglePackageContentEntity.PlatinumLandingPackageItemId.Valid {
		landingSinglePackageContentPlatinumEntity, err := landingPackageItemRepository.FindById(ctx, landingSinglePackageContentEntity.PlatinumLandingPackageItemId.Int64)
		if err != nil {
			return dto.LandingResponse{}, err
		}

		landingSinglePackageContentPlatinumResponseValue, err := m.mapPackageItemEntityToResponse(
			ctx,
			packageRepository,
			imageRepository,
			packageSessionRepository,
			itineraryRepository,
			itineraryDayRepository,
			flightRouteRepository,
			flightRepository,
			airlineRepository,
			landingSinglePackageContentPlatinumEntity,
		)
		if err != nil {
			return dto.LandingResponse{}, err
		}

		landingSinglePackageContentPlatinumResponse = null.ValueFrom(landingSinglePackageContentPlatinumResponseValue)
	}

	landingPackagesContentSilverEntity, err := landingPackageDetailRepository.FindById(ctx, landingPackagesContentEntity.SilverLandingPackageDetailId)
	if err != nil {
		return dto.LandingResponse{}, err
	}
	landingPackagesContentSilverResponse, err := m.mapPackageDetailEntityToResponse(
		ctx,
		packageRepository,
		imageRepository,
		packageSessionRepository,
		itineraryRepository,
		itineraryDayRepository,
		flightRouteRepository,
		flightRepository,
		airlineRepository,
		landingSectionHeaderRepository,
		landingPackageDetailItemRepository,
		landingPackageItemRepository,
		landingPackagesContentSilverEntity,
	)
	if err != nil {
		return dto.LandingResponse{}, err
	}

	landingPackagesContentGoldEntity, err := landingPackageDetailRepository.FindById(ctx, landingPackagesContentEntity.GoldLandingPackageDetailId)
	if err != nil {
		return dto.LandingResponse{}, err
	}
	landingPackagesContentGoldResponse, err := m.mapPackageDetailEntityToResponse(
		ctx,
		packageRepository,
		imageRepository,
		packageSessionRepository,
		itineraryRepository,
		itineraryDayRepository,
		flightRouteRepository,
		flightRepository,
		airlineRepository,
		landingSectionHeaderRepository,
		landingPackageDetailItemRepository,
		landingPackageItemRepository,
		landingPackagesContentGoldEntity,
	)
	if err != nil {
		return dto.LandingResponse{}, err
	}

	landingPackagesContentPlatinumEntity, err := landingPackageDetailRepository.FindById(ctx, landingPackagesContentEntity.PlatinumLandingPackageDetailId)
	if err != nil {
		return dto.LandingResponse{}, err
	}
	landingPackagesContentPlatinumResponse, err := m.mapPackageDetailEntityToResponse(
		ctx,
		packageRepository,
		imageRepository,
		packageSessionRepository,
		itineraryRepository,
		itineraryDayRepository,
		flightRouteRepository,
		flightRepository,
		airlineRepository,
		landingSectionHeaderRepository,
		landingPackageDetailItemRepository,
		landingPackageItemRepository,
		landingPackagesContentPlatinumEntity,
	)
	if err != nil {
		return dto.LandingResponse{}, err
	}

	landingFeaturesContentHeaderEntity, err := landingSectionHeaderRepository.FindById(ctx, landingFeaturesContentEntity.LandingSectionHeaderId)
	if err != nil {
		return dto.LandingResponse{}, err
	}

	landingFeaturesContentBenefitEntities, err := landingFeaturesContentBenefitRepository.FindAll(ctx, repository.FindAllOptions{})
	if err != nil {
		return dto.LandingResponse{}, err
	}
	landingFeaturesContentBenefitResponses := make([]dto.LandingFeaturesContentBenefitResponse, len(landingFeaturesContentBenefitEntities))
	for i, landingFeaturesContentBenefitEntity := range landingFeaturesContentBenefitEntities {
		logoResponse := null.NewValue(dto.ImageResponse{}, false)
		if landingFeaturesContentBenefitEntity.LogoId.Valid {
			logoEntity, err := imageRepository.FindById(ctx, landingFeaturesContentBenefitEntity.LogoId.Int64)
			if err != nil {
				return dto.LandingResponse{}, err
			}

			logoResponse = null.ValueFrom(m.imageMapper.MapEntityToResponse(ctx, logoEntity))
		}

		landingFeaturesContentBenefitResponses[i] = dto.LandingFeaturesContentBenefitResponse{
			IsEnabled: landingFeaturesContentBenefitEntity.IsEnabled,
			IsMobile:  landingFeaturesContentBenefitEntity.IsMobile,
			IsDesktop: landingFeaturesContentBenefitEntity.IsDesktop,
			Title:     landingFeaturesContentBenefitEntity.Title,
			Subtitle:  landingFeaturesContentBenefitEntity.Subtitle,
			Logo:      logoResponse,
		}
	}

	landingMomentsContentHeaderEntity, err := landingSectionHeaderRepository.FindById(ctx, landingMomentsContentEntity.LandingSectionHeaderId)
	if err != nil {
		return dto.LandingResponse{}, err
	}

	landingMomentsContentImageEntities, err := landingMomentsContentImageRepository.FindAll(ctx, repository.FindAllOptions{})
	if err != nil {
		return dto.LandingResponse{}, err
	}
	landingMomentsContentImageResponses := make([]dto.LandingMomentsContentImageResponse, len(landingMomentsContentImageEntities))
	for i, landingMomentsContentImageEntity := range landingMomentsContentImageEntities {
		imageEntity, err := imageRepository.FindById(ctx, landingMomentsContentImageEntity.ImageId)
		if err != nil {
			return dto.LandingResponse{}, err
		}

		landingMomentsContentImageResponses[i] = dto.LandingMomentsContentImageResponse{
			IsEnabled: landingMomentsContentImageEntity.IsEnabled,
			IsMobile:  landingMomentsContentImageEntity.IsMobile,
			IsDesktop: landingMomentsContentImageEntity.IsDesktop,
			Image:     m.imageMapper.MapEntityToResponse(ctx, imageEntity),
		}
	}

	landingAffiliatesContentHeaderEntity, err := landingSectionHeaderRepository.FindById(ctx, landingAffiliatesContentEntity.LandingSectionHeaderId)
	if err != nil {
		return dto.LandingResponse{}, err
	}

	landingAffiliatesContentAffiliateEntities, err := landingAffiliatesContentAffiliateRepository.FindAll(ctx, repository.FindAllOptions{})
	if err != nil {
		return dto.LandingResponse{}, err
	}
	landingAffiliatesContentAffiliateResponses := make([]dto.LandingAffiliatesContentAffiliateResponse, len(landingAffiliatesContentAffiliateEntities))
	for i, landingAffiliatesContentAffiliateEntity := range landingAffiliatesContentAffiliateEntities {
		logoResponse := null.NewValue(dto.ImageResponse{}, false)
		if landingAffiliatesContentAffiliateEntity.LogoId.Valid {
			logoEntity, err := imageRepository.FindById(ctx, landingAffiliatesContentAffiliateEntity.LogoId.Int64)
			if err != nil {
				return dto.LandingResponse{}, err
			}

			logoResponse = null.ValueFrom(m.imageMapper.MapEntityToResponse(ctx, logoEntity))
		}

		landingAffiliatesContentAffiliateResponses[i] = dto.LandingAffiliatesContentAffiliateResponse{
			IsEnabled: landingAffiliatesContentAffiliateEntity.IsEnabled,
			IsMobile:  landingAffiliatesContentAffiliateEntity.IsMobile,
			IsDesktop: landingAffiliatesContentAffiliateEntity.IsDesktop,
			Name:      landingAffiliatesContentAffiliateEntity.Name,
			Logo:      logoResponse,
			Width:     landingAffiliatesContentAffiliateEntity.Width,
			Height:    landingAffiliatesContentAffiliateEntity.Height,
		}
	}

	landingFaqContentHeaderEntity, err := landingSectionHeaderRepository.FindById(ctx, landingFaqContentEntity.LandingSectionHeaderId)
	if err != nil {
		return dto.LandingResponse{}, err
	}

	landingFaqContentFaqEntities, err := LandingFaqContentFaqRepository.FindAll(ctx, repository.FindAllOptions{})
	if err != nil {
		return dto.LandingResponse{}, err
	}
	landingFaqContentFaqResponses := make([]dto.LandingFaqContentFaqResponse, len(landingFaqContentFaqEntities))
	for i, landingFaqContentFaqEntity := range landingFaqContentFaqEntities {
		landingFaqContentFaqResponses[i] = dto.LandingFaqContentFaqResponse{
			IsEnabled: landingFaqContentFaqEntity.IsEnabled,
			IsMobile:  landingFaqContentFaqEntity.IsMobile,
			IsDesktop: landingFaqContentFaqEntity.IsDesktop,
			Question:  landingFaqContentFaqEntity.Question,
			Answer:    landingFaqContentFaqEntity.Answer,
		}
	}

	landingMenuResponses := make([]dto.LandingMenuResponse, len(landingMenuEntities))
	for i, landingMenuEntity := range landingMenuEntities {
		landingMenuResponses[i] = dto.LandingMenuResponse{
			IsEnabled: landingMenuEntity.IsEnabled,
			IsMobile:  landingMenuEntity.IsMobile,
			IsDesktop: landingMenuEntity.IsDesktop,
			Icon:      landingMenuEntity.Icon,
			Label:     landingMenuEntity.Label,
			Path:      landingMenuEntity.Path,
		}
	}

	return dto.LandingResponse{
		HeroContent: dto.LandingHeroContentResponse{
			IsEnabled:   landingHeroContentEntity.IsEnabled,
			IsMobile:    landingHeroContentEntity.IsMobile,
			IsDesktop:   landingHeroContentEntity.IsDesktop,
			Title:       landingHeroContentEntity.Title,
			Description: landingHeroContentEntity.Description,
			TagsLine:    landingHeroContentEntity.TagsLine,
			ButtonLabel: landingHeroContentEntity.ButtonLabel,
			ButtonUrl:   landingHeroContentEntity.ButtonUrl,
			Image:       landingHeroContentImageResponse,
		},
		SinglePackageContent: dto.LandingSinglePackageContentResponse{
			IsEnabled: landingSinglePackageContentEntity.IsEnabled,
			IsMobile:  landingSinglePackageContentEntity.IsMobile,
			IsDesktop: landingSinglePackageContentEntity.IsDesktop,
			Header: dto.LandingSectionHeaderResponse{
				IsEnabled: landingSinglePackageContentHeaderEntity.IsEnabled,
				IsMobile:  landingSinglePackageContentHeaderEntity.IsMobile,
				IsDesktop: landingSinglePackageContentHeaderEntity.IsDesktop,
				Title:     landingSinglePackageContentHeaderEntity.Title,
				Subtitle:  landingSinglePackageContentHeaderEntity.Subtitle,
				TagsLine:  landingSinglePackageContentHeaderEntity.TagsLine,
			},
			Silver:   landingSinglePackageContentSilverResponse,
			Gold:     landingSinglePackageContentGoldResponse,
			Platinum: landingSinglePackageContentPlatinumResponse,
		},
		PackagesContent: dto.LandingPackagesContentResponse{
			IsEnabled: landingPackagesContentEntity.IsEnabled,
			IsMobile:  landingPackagesContentEntity.IsMobile,
			IsDesktop: landingPackagesContentEntity.IsDesktop,
			Silver:    landingPackagesContentSilverResponse,
			Gold:      landingPackagesContentGoldResponse,
			Platinum:  landingPackagesContentPlatinumResponse,
		},
		FeaturesContent: dto.LandingFeaturesContentResponse{
			IsEnabled: landingFeaturesContentEntity.IsEnabled,
			IsMobile:  landingFeaturesContentEntity.IsMobile,
			IsDesktop: landingFeaturesContentEntity.IsDesktop,
			Header: dto.LandingSectionHeaderResponse{
				IsEnabled: landingFeaturesContentHeaderEntity.IsEnabled,
				IsMobile:  landingFeaturesContentHeaderEntity.IsMobile,
				IsDesktop: landingFeaturesContentHeaderEntity.IsDesktop,
				Title:     landingFeaturesContentHeaderEntity.Title,
				Subtitle:  landingFeaturesContentHeaderEntity.Subtitle,
				TagsLine:  landingFeaturesContentHeaderEntity.TagsLine,
			},
			Benefits:      landingFeaturesContentBenefitResponses,
			FooterTitle:   landingFeaturesContentEntity.FooterTitle,
			ButtonAbout:   landingFeaturesContentEntity.ButtonAbout,
			ButtonPackage: landingFeaturesContentEntity.ButtonPackage,
		},
		MomentsContent: dto.LandingMomentsContentResponse{
			IsEnabled: landingMomentsContentEntity.IsEnabled,
			IsMobile:  landingMomentsContentEntity.IsMobile,
			IsDesktop: landingMomentsContentEntity.IsDesktop,
			Header: dto.LandingSectionHeaderResponse{
				IsEnabled: landingMomentsContentHeaderEntity.IsEnabled,
				IsMobile:  landingMomentsContentHeaderEntity.IsMobile,
				IsDesktop: landingMomentsContentHeaderEntity.IsDesktop,
				Title:     landingMomentsContentHeaderEntity.Title,
				Subtitle:  landingMomentsContentHeaderEntity.Subtitle,
				TagsLine:  landingMomentsContentHeaderEntity.TagsLine,
			},
			Images: landingMomentsContentImageResponses,
		},
		AffiliatesContent: dto.LandingAffiliatesContentResponse{
			IsEnabled: landingAffiliatesContentEntity.IsEnabled,
			IsMobile:  landingAffiliatesContentEntity.IsMobile,
			IsDesktop: landingAffiliatesContentEntity.IsDesktop,
			Header: dto.LandingSectionHeaderResponse{
				IsEnabled: landingAffiliatesContentHeaderEntity.IsEnabled,
				IsMobile:  landingAffiliatesContentHeaderEntity.IsMobile,
				IsDesktop: landingAffiliatesContentHeaderEntity.IsDesktop,
				Title:     landingAffiliatesContentHeaderEntity.Title,
				Subtitle:  landingAffiliatesContentHeaderEntity.Subtitle,
				TagsLine:  landingAffiliatesContentHeaderEntity.TagsLine,
			},
			Affiliates: landingAffiliatesContentAffiliateResponses,
		},
		FaqContent: dto.LandingFaqContentResponse{
			IsEnabled: landingFaqContentEntity.IsEnabled,
			IsMobile:  landingFaqContentEntity.IsMobile,
			IsDesktop: landingFaqContentEntity.IsDesktop,
			Header: dto.LandingSectionHeaderResponse{
				IsEnabled: landingFaqContentHeaderEntity.IsEnabled,
				IsMobile:  landingFaqContentHeaderEntity.IsMobile,
				IsDesktop: landingFaqContentHeaderEntity.IsDesktop,
				Title:     landingFaqContentHeaderEntity.Title,
				Subtitle:  landingFaqContentHeaderEntity.Subtitle,
				TagsLine:  landingFaqContentHeaderEntity.TagsLine,
			},
			Faqs: landingFaqContentFaqResponses,
		},
		Menus: landingMenuResponses,
	}, nil
}
