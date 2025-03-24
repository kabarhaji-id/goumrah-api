package service

import (
	"context"

	"github.com/guregu/null/v6"
	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
	"github.com/kabarhaji-id/goumrah-api/internal/domain/mapper"
	"github.com/kabarhaji-id/goumrah-api/internal/domain/validator"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driven/repository"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
	serviceport "github.com/kabarhaji-id/goumrah-api/internal/port/driving/service"
)

type landingServiceImpl struct {
	landingHeroContentRepository                repository.LandingHeroContentRepository
	landingSectionHeaderRepository              repository.LandingSectionHeaderRepository
	landingPackageItemRepository                repository.LandingPackageItemRepository
	landingSinglePackageContentRepository       repository.LandingSinglePackageContentRepository
	landingPackageDetailRepository              repository.LandingPackageDetailRepository
	landingPackageDetailItemRepository          repository.LandingPackageDetailItemRepository
	landingPackagesContentRepository            repository.LandingPackagesContentRepository
	landingFeaturesContentRepository            repository.LandingFeaturesContentRepository
	landingFeaturesContentBenefitRepository     repository.LandingFeaturesContentBenefitRepository
	landingMomentsContentRepository             repository.LandingMomentsContentRepository
	landingMomentsContentImageRepository        repository.LandingMomentsContentImageRepository
	landingAffiliatesContentRepository          repository.LandingAffiliatesContentRepository
	landingAffiliatesContentAffiliateRepository repository.LandingAffiliatesContentAffiliateRepository
	landingFaqContentRepository                 repository.LandingFaqContentRepository
	landingFaqContentFaqRepository              repository.LandingFaqContentFaqRepository
	landingMenuRepository                       repository.LandingMenuRepository

	imageRepository          repository.ImageRepository
	packageRepository        repository.PackageRepository
	packageSessionRepository repository.PackageSessionRepository
	itineraryRepository      repository.ItineraryRepository
	itineraryDayRepository   repository.ItineraryDayRepository
	flightRouteRepository    repository.FlightRouteRepository
	flightRepository         repository.FlightRepository
	airlineRepository        repository.AirlineRepository

	landingValidator validator.LandingValidator
	landingMapper    mapper.LandingMapper

	unitOfWork repository.UnitOfWork
}

func NewLandingService(
	landingHeroContentRepository repository.LandingHeroContentRepository,
	landingSectionHeaderRepository repository.LandingSectionHeaderRepository,
	landingPackageItemRepository repository.LandingPackageItemRepository,
	landingSinglePackageContentRepository repository.LandingSinglePackageContentRepository,
	landingPackageDetailRepository repository.LandingPackageDetailRepository,
	landingPackageDetailItemRepository repository.LandingPackageDetailItemRepository,
	landingPackagesContentRepository repository.LandingPackagesContentRepository,
	landingFeaturesContentRepository repository.LandingFeaturesContentRepository,
	landingFeaturesContentBenefitRepository repository.LandingFeaturesContentBenefitRepository,
	landingMomentsContentRepository repository.LandingMomentsContentRepository,
	landingMomentsContentImageRepository repository.LandingMomentsContentImageRepository,
	landingAffiliatesContentRepository repository.LandingAffiliatesContentRepository,
	landingAffiliatesContentAffiliateRepository repository.LandingAffiliatesContentAffiliateRepository,
	landingFaqContentRepository repository.LandingFaqContentRepository,
	landingFaqContentFaqRepository repository.LandingFaqContentFaqRepository,
	landingMenuRepository repository.LandingMenuRepository,
	imageRepository repository.ImageRepository,
	packageRepository repository.PackageRepository,
	packageSessionRepository repository.PackageSessionRepository,
	itineraryRepository repository.ItineraryRepository,
	itineraryDayRepository repository.ItineraryDayRepository,
	flightRouteRepository repository.FlightRouteRepository,
	flightRepository repository.FlightRepository,
	airlineRepository repository.AirlineRepository,
	landingValidator validator.LandingValidator,
	landingMapper mapper.LandingMapper,
	unitOfWork repository.UnitOfWork,
) serviceport.LandingService {
	return landingServiceImpl{
		landingHeroContentRepository,
		landingSectionHeaderRepository,
		landingPackageItemRepository,
		landingSinglePackageContentRepository,
		landingPackageDetailRepository,
		landingPackageDetailItemRepository,
		landingPackagesContentRepository,
		landingFeaturesContentRepository,
		landingFeaturesContentBenefitRepository,
		landingMomentsContentRepository,
		landingMomentsContentImageRepository,
		landingAffiliatesContentRepository,
		landingAffiliatesContentAffiliateRepository,
		landingFaqContentRepository,
		landingFaqContentFaqRepository,
		landingMenuRepository,
		imageRepository,
		packageRepository,
		packageSessionRepository,
		itineraryRepository,
		itineraryDayRepository,
		flightRouteRepository,
		flightRepository,
		airlineRepository,
		landingValidator,
		landingMapper,
		unitOfWork,
	}
}

func (s landingServiceImpl) CreateLanding(ctx context.Context, request dto.LandingRequest) (dto.LandingResponse, error) {
	// Validate request
	if err := s.landingValidator.ValidateRequest(request); err != nil {
		return dto.LandingResponse{}, err
	}

	// Create response
	response := dto.LandingResponse{}

	// Start transaction with unit of work
	err := s.unitOfWork.Do(ctx, func(ctx context.Context, factory repository.Factory) error {
		// Create landing repositories
		landingHeroContentRepository := factory.NewLandingHeroContentRepository()
		landingSectionHeaderRepository := factory.NewLandingSectionHeaderRepository()
		landingPackageItemRepository := factory.NewLandingPackageItemRepository()
		landingSinglePackageContentRepository := factory.NewLandingSinglePackageContentRepository()
		landingPackageDetailRepository := factory.NewLandingPackageDetailRepository()
		landingPackageDetailItemRepository := factory.NewLandingPackageDetailItemRepository()
		landingPackagesContentRepository := factory.NewLandingPackagesContentRepository()
		landingFeaturesContentRepository := factory.NewLandingFeaturesContentRepository()
		landingFeaturesContentBenefitRepository := factory.NewLandingFeaturesContentBenefitRepository()
		landingMomentsContentRepository := factory.NewLandingMomentsContentRepository()
		landingMomentsContentImageRepository := factory.NewLandingMomentsContentImageRepository()
		landingAffiliatesContentRepository := factory.NewLandingAffiliatesContentRepository()
		landingAffiliatesContentAffiliateRepository := factory.NewLandingAffiliatesContentAffiliateRepository()
		landingFaqContentRepository := factory.NewLandingFaqContentRepository()
		landingFaqContentFaqRepository := factory.NewLandingFaqContentFaqRepository()
		landingMenuRepository := factory.NewLandingMenuRepository()

		// Create image repository
		imageRepository := factory.NewImageRepository()

		// Create package repository
		packageRepository := factory.NewPackageRepository()

		// Create package session repository
		packageSessionRepository := factory.NewPackageSessionRepository()

		// Create itinerary repository
		itineraryRepository := factory.NewItineraryRepository()

		// Create itinerary day repository
		itineraryDayRepository := factory.NewItineraryDayRepository()

		// Create flight route repository
		flightRouteRepository := factory.NewFlightRouteRepository()

		// Create flight repository
		flightRepository := factory.NewFlightRepository()

		// Create airline repository
		airlineRepository := factory.NewAirlineRepository()

		// Create landing hero content with repository
		landingHeroContent, err := landingHeroContentRepository.Create(
			ctx,
			entity.LandingHeroContent{
				IsEnabled:   request.HeroContent.IsEnabled,
				Title:       request.HeroContent.Title,
				Description: request.HeroContent.Description,
				TagsLine:    request.HeroContent.TagsLine,
				ButtonLabel: request.HeroContent.ButtonLabel,
				ButtonUrl:   request.HeroContent.ButtonUrl,
				ImageId:     request.HeroContent.Image,
			},
		)
		if err != nil {
			return err
		}

		// Create landing single package content with repository
		landingSinglePackageContentHeader, err := landingSectionHeaderRepository.Create(
			ctx,
			entity.LandingSectionHeader{
				IsEnabled: request.SinglePackageContent.Header.IsEnabled,
				Title:     request.SinglePackageContent.Header.Title,
				Subtitle:  request.SinglePackageContent.Header.Subtitle,
				TagsLine:  request.SinglePackageContent.Header.TagsLine,
			},
		)
		if err != nil {
			return err
		}

		landingSinglePackageContentSilverLandingPackageItem, err := landingPackageItemRepository.Create(
			ctx,
			entity.LandingPackageItem{
				IsEnabled:   request.SinglePackageContent.Silver.IsEnabled,
				ButtonLabel: request.SinglePackageContent.Silver.ButtonLabel,
				PackageId:   request.SinglePackageContent.Silver.Package,
			},
		)
		if err != nil {
			return err
		}

		landingSinglePackageContentGoldLandingPackageItem, err := landingPackageItemRepository.Create(
			ctx,
			entity.LandingPackageItem{
				IsEnabled:   request.SinglePackageContent.Gold.IsEnabled,
				ButtonLabel: request.SinglePackageContent.Gold.ButtonLabel,
				PackageId:   request.SinglePackageContent.Gold.Package,
			},
		)
		if err != nil {
			return err
		}

		landingSinglePackageContentPlatinumLandingPackageItem, err := landingPackageItemRepository.Create(
			ctx,
			entity.LandingPackageItem{
				IsEnabled:   request.SinglePackageContent.Platinum.IsEnabled,
				PackageId:   request.SinglePackageContent.Platinum.Package,
				ButtonLabel: request.SinglePackageContent.Platinum.ButtonLabel,
			},
		)
		if err != nil {
			return err
		}

		landingSinglePackageContent, err := landingSinglePackageContentRepository.Create(
			ctx,
			entity.LandingSinglePackageContent{
				IsEnabled:                    request.SinglePackageContent.IsEnabled,
				LandingSectionHeaderId:       landingSinglePackageContentHeader.Id,
				SilverLandingPackageItemId:   null.IntFrom(landingSinglePackageContentSilverLandingPackageItem.Id),
				GoldLandingPackageItemId:     null.IntFrom(landingSinglePackageContentGoldLandingPackageItem.Id),
				PlatinumLandingPackageItemId: null.IntFrom(landingSinglePackageContentPlatinumLandingPackageItem.Id),
			},
		)
		if err != nil {
			return err
		}

		// Create landing packages content with repository
		landingPackagesContentSilverLandingPackageDetailLandingSectionHeader, err := landingSectionHeaderRepository.Create(
			ctx,
			entity.LandingSectionHeader{
				IsEnabled: request.PackagesContent.Silver.Header.IsEnabled,
				Title:     request.PackagesContent.Silver.Header.Title,
				Subtitle:  request.PackagesContent.Silver.Header.Subtitle,
				TagsLine:  request.PackagesContent.Silver.Header.TagsLine,
			},
		)
		if err != nil {
			return err
		}

		landingPackagesContentSilverLandingPackageDetail, err := landingPackageDetailRepository.Create(
			ctx,
			entity.LandingPackageDetail{
				IsEnabled:              request.PackagesContent.Silver.IsEnabled,
				LandingSectionHeaderId: landingPackagesContentSilverLandingPackageDetailLandingSectionHeader.Id,
			},
		)
		if err != nil {
			return err
		}

		landingPackagesContentSilverLandingPackageDetailItems := make(
			[]entity.LandingPackageDetailItem, len(request.PackagesContent.Silver.Packages),
		)
		for i, packageItem := range request.PackagesContent.Silver.Packages {
			landingPackagesContentSilverLandingPackageItem, err := landingPackageItemRepository.Create(
				ctx,
				entity.LandingPackageItem{
					IsEnabled:   packageItem.IsEnabled,
					PackageId:   packageItem.Package,
					ButtonLabel: packageItem.ButtonLabel,
				},
			)
			if err != nil {
				return err
			}

			landingPackagesContentSilverLandingPackageDetailItems[i] = entity.LandingPackageDetailItem{
				IsEnabled:              packageItem.IsEnabled,
				LandingPackageDetailId: landingPackagesContentSilverLandingPackageDetail.Id,
				LandingPackageItemId:   landingPackagesContentSilverLandingPackageItem.Id,
			}
		}
		landingPackagesContentSilverLandingPackageDetailItems, err = landingPackageDetailItemRepository.CreateMany(
			ctx, landingPackagesContentSilverLandingPackageDetailItems,
		)
		if err != nil {
			return err
		}

		landingPackagesContentGoldLandingPackageDetailLandingSectionHeader, err := landingSectionHeaderRepository.Create(
			ctx,
			entity.LandingSectionHeader{
				IsEnabled: request.PackagesContent.Gold.Header.IsEnabled,
				Title:     request.PackagesContent.Gold.Header.Title,
				Subtitle:  request.PackagesContent.Gold.Header.Subtitle,
				TagsLine:  request.PackagesContent.Gold.Header.TagsLine,
			},
		)
		if err != nil {
			return err
		}

		landingPackagesContentGoldLandingPackageDetail, err := landingPackageDetailRepository.Create(
			ctx,
			entity.LandingPackageDetail{
				IsEnabled:              request.PackagesContent.Gold.IsEnabled,
				LandingSectionHeaderId: landingPackagesContentGoldLandingPackageDetailLandingSectionHeader.Id,
			},
		)
		if err != nil {
			return err
		}

		landingPackagesContentGoldLandingPackageDetailItems := make(
			[]entity.LandingPackageDetailItem, len(request.PackagesContent.Gold.Packages),
		)
		for i, packageItem := range request.PackagesContent.Gold.Packages {
			landingPackagesContentGoldLandingPackageItem, err := landingPackageItemRepository.Create(
				ctx,
				entity.LandingPackageItem{
					IsEnabled:   packageItem.IsEnabled,
					PackageId:   packageItem.Package,
					ButtonLabel: packageItem.ButtonLabel,
				},
			)
			if err != nil {
				return err
			}

			landingPackagesContentGoldLandingPackageDetailItems[i] = entity.LandingPackageDetailItem{
				IsEnabled:              packageItem.IsEnabled,
				LandingPackageDetailId: landingPackagesContentGoldLandingPackageDetail.Id,
				LandingPackageItemId:   landingPackagesContentGoldLandingPackageItem.Id,
			}
		}
		landingPackagesContentGoldLandingPackageDetailItems, err = landingPackageDetailItemRepository.CreateMany(
			ctx, landingPackagesContentGoldLandingPackageDetailItems,
		)
		if err != nil {
			return err
		}

		landingPackagesContentPlatinumLandingPackageDetailLandingSectionHeader, err := landingSectionHeaderRepository.Create(
			ctx,
			entity.LandingSectionHeader{
				IsEnabled: request.PackagesContent.Platinum.Header.IsEnabled,
				Title:     request.PackagesContent.Platinum.Header.Title,
				Subtitle:  request.PackagesContent.Platinum.Header.Subtitle,
				TagsLine:  request.PackagesContent.Platinum.Header.TagsLine,
			},
		)
		if err != nil {
			return err
		}

		landingPackagesContentPlatinumLandingPackageDetail, err := landingPackageDetailRepository.Create(
			ctx,
			entity.LandingPackageDetail{
				IsEnabled:              request.PackagesContent.Platinum.IsEnabled,
				LandingSectionHeaderId: landingPackagesContentPlatinumLandingPackageDetailLandingSectionHeader.Id,
			},
		)
		if err != nil {
			return err
		}

		landingPackagesContentPlatinumLandingPackageDetailItems := make(
			[]entity.LandingPackageDetailItem, len(request.PackagesContent.Platinum.Packages),
		)
		for i, packageItem := range request.PackagesContent.Platinum.Packages {
			landingPackagesContentPlatinumLandingPackageItem, err := landingPackageItemRepository.Create(
				ctx,
				entity.LandingPackageItem{
					IsEnabled:   packageItem.IsEnabled,
					PackageId:   packageItem.Package,
					ButtonLabel: packageItem.ButtonLabel,
				},
			)
			if err != nil {
				return err
			}

			landingPackagesContentPlatinumLandingPackageDetailItems[i] = entity.LandingPackageDetailItem{
				IsEnabled:              packageItem.IsEnabled,
				LandingPackageDetailId: landingPackagesContentPlatinumLandingPackageDetail.Id,
				LandingPackageItemId:   landingPackagesContentPlatinumLandingPackageItem.Id,
			}
		}
		landingPackagesContentPlatinumLandingPackageDetailItems, err = landingPackageDetailItemRepository.CreateMany(
			ctx, landingPackagesContentPlatinumLandingPackageDetailItems,
		)
		if err != nil {
			return err
		}

		landingPackagesContent, err := landingPackagesContentRepository.Create(
			ctx,
			entity.LandingPackagesContent{
				IsEnabled:                      request.PackagesContent.IsEnabled,
				SilverLandingPackageDetailId:   null.IntFrom(landingPackagesContentSilverLandingPackageDetail.Id),
				GoldLandingPackageDetailId:     null.IntFrom(landingPackagesContentGoldLandingPackageDetail.Id),
				PlatinumLandingPackageDetailId: null.IntFrom(landingPackagesContentPlatinumLandingPackageDetail.Id),
			},
		)
		if err != nil {
			return err
		}

		// Create landing features content with repository
		landingFeaturesContentHeader, err := landingSectionHeaderRepository.Create(
			ctx,
			entity.LandingSectionHeader{
				IsEnabled: request.FeaturesContent.Header.IsEnabled,
				Title:     request.FeaturesContent.Header.Title,
				Subtitle:  request.FeaturesContent.Header.Subtitle,
				TagsLine:  request.FeaturesContent.Header.TagsLine,
			},
		)
		if err != nil {
			return err
		}

		benefits := make([]entity.LandingFeaturesContentBenefit, len(request.FeaturesContent.Benefits))
		for index, benefit := range request.FeaturesContent.Benefits {
			benefits[index] = entity.LandingFeaturesContentBenefit{
				IsEnabled: benefit.IsEnabled,
				Title:     benefit.Title,
				Subtitle:  benefit.Subtitle,
				LogoId:    benefit.Logo,
			}
		}
		if _, err = landingFeaturesContentBenefitRepository.CreateMany(ctx, benefits); err != nil {
			return err
		}

		landingFeaturesContent, err := landingFeaturesContentRepository.Create(
			ctx,
			entity.LandingFeaturesContent{
				IsEnabled:              request.FeaturesContent.IsEnabled,
				LandingSectionHeaderId: landingFeaturesContentHeader.Id,
				FooterTitle:            request.FeaturesContent.FooterTitle,
				ButtonAbout:            request.FeaturesContent.ButtonAbout,
				ButtonPackage:          request.FeaturesContent.ButtonPackage,
			},
		)
		if err != nil {
			return err
		}

		// Create landing moments content with repository
		landingMomentsContentHeader, err := landingSectionHeaderRepository.Create(
			ctx,
			entity.LandingSectionHeader{
				IsEnabled: request.MomentsContent.Header.IsEnabled,
				Title:     request.MomentsContent.Header.Title,
				Subtitle:  request.MomentsContent.Header.Subtitle,
				TagsLine:  request.MomentsContent.Header.TagsLine,
			},
		)
		if err != nil {
			return err
		}

		landingMomentsContentImages := make([]entity.LandingMomentsContentImage, len(request.MomentsContent.Images))
		for index, image := range request.MomentsContent.Images {
			landingMomentsContentImages[index] = entity.LandingMomentsContentImage{
				IsEnabled: image.IsEnabled,
				ImageId:   image.Image,
			}
		}
		if _, err = landingMomentsContentImageRepository.CreateMany(
			ctx,
			landingMomentsContentImages,
		); err != nil {
			return err
		}

		landingMomentsContent, err := landingMomentsContentRepository.Create(
			ctx,
			entity.LandingMomentsContent{
				IsEnabled:              request.MomentsContent.IsEnabled,
				LandingSectionHeaderId: landingMomentsContentHeader.Id,
			},
		)
		if err != nil {
			return err
		}

		// Create landing affiliates content with repository
		landingAffiliatesContentHeader, err := landingSectionHeaderRepository.Create(
			ctx,
			entity.LandingSectionHeader{
				IsEnabled: request.AffiliatesContent.Header.IsEnabled,
				Title:     request.AffiliatesContent.Header.Title,
				Subtitle:  request.AffiliatesContent.Header.Subtitle,
				TagsLine:  request.AffiliatesContent.Header.TagsLine,
			},
		)
		if err != nil {
			return err
		}

		affiliates := make([]entity.LandingAffiliatesContentAffiliate, len(request.AffiliatesContent.Affiliates))
		for index, affiliate := range request.AffiliatesContent.Affiliates {
			affiliates[index] = entity.LandingAffiliatesContentAffiliate{
				IsEnabled: affiliate.IsEnabled,
				Name:      affiliate.Name,
				LogoId:    affiliate.Logo,
				Width:     affiliate.Width,
				Height:    affiliate.Height,
			}
		}
		if _, err = landingAffiliatesContentAffiliateRepository.CreateMany(ctx, affiliates); err != nil {
			return err
		}

		landingAffiliatesContent, err := landingAffiliatesContentRepository.Create(
			ctx,
			entity.LandingAffiliatesContent{
				IsEnabled:              request.AffiliatesContent.IsEnabled,
				LandingSectionHeaderId: landingAffiliatesContentHeader.Id,
			},
		)
		if err != nil {
			return err
		}

		// Create landing faq content with repository
		landingFaqContentHeader, err := landingSectionHeaderRepository.Create(
			ctx,
			entity.LandingSectionHeader{
				IsEnabled: request.FaqContent.Header.IsEnabled,
				Title:     request.FaqContent.Header.Title,
				Subtitle:  request.FaqContent.Header.Subtitle,
				TagsLine:  request.FaqContent.Header.TagsLine,
			},
		)
		if err != nil {
			return err
		}

		faqs := make([]entity.LandingFaqContentFaq, len(request.FaqContent.Faqs))
		for index, faq := range request.FaqContent.Faqs {
			faqs[index] = entity.LandingFaqContentFaq{
				IsEnabled: faq.IsEnabled,
				Question:  faq.Question,
				Answer:    faq.Answer,
			}
		}
		if _, err = landingFaqContentFaqRepository.CreateMany(ctx, faqs); err != nil {
			return err
		}

		landingFaqContent, err := landingFaqContentRepository.Create(
			ctx,
			entity.LandingFaqContent{
				IsEnabled:              request.FaqContent.IsEnabled,
				LandingSectionHeaderId: landingFaqContentHeader.Id,
			},
		)
		if err != nil {
			return err
		}

		// Create landing menu with repository
		landingMenus := make([]entity.LandingMenu, len(request.Menus))
		for index, menu := range request.Menus {
			landingMenus[index] = entity.LandingMenu{
				IsEnabled: menu.IsEnabled,
				Icon:      menu.Icon,
				Label:     menu.Label,
				Path:      menu.Path,
			}
		}
		if _, err = landingMenuRepository.CreateMany(ctx, landingMenus); err != nil {
			return err
		}

		// Map entity into response
		response, err = s.landingMapper.MapEntityToResponse(
			ctx,
			landingSectionHeaderRepository,
			landingPackageItemRepository,
			landingPackageDetailRepository,
			landingPackageDetailItemRepository,
			landingFeaturesContentBenefitRepository,
			landingMomentsContentImageRepository,
			landingAffiliatesContentAffiliateRepository,
			landingFaqContentFaqRepository,
			imageRepository,
			packageRepository,
			packageSessionRepository,
			itineraryRepository,
			itineraryDayRepository,
			flightRouteRepository,
			flightRepository,
			airlineRepository,
			landingHeroContent,
			landingSinglePackageContent,
			landingPackagesContent,
			landingFeaturesContent,
			landingMomentsContent,
			landingAffiliatesContent,
			landingFaqContent,
			landingMenus,
		)

		return err
	})

	return response, err
}

func (s landingServiceImpl) GetLanding(ctx context.Context) (dto.LandingResponse, error) {
	// Find landing hero content with repository
	landingHeroContent, err := s.landingHeroContentRepository.Find(ctx)
	if err != nil {
		return dto.LandingResponse{}, err
	}

	// Find landing single package content with repository
	landingSinglePackageContent, err := s.landingSinglePackageContentRepository.Find(ctx)
	if err != nil {
		return dto.LandingResponse{}, err
	}

	// Find landing packages content with repository
	landingPackagesContent, err := s.landingPackagesContentRepository.Find(ctx)
	if err != nil {
		return dto.LandingResponse{}, err
	}

	// Find landing features content with repository
	landingFeaturesContent, err := s.landingFeaturesContentRepository.Find(ctx)
	if err != nil {
		return dto.LandingResponse{}, err
	}

	// Find landing moments content with repository
	landingMomentsContent, err := s.landingMomentsContentRepository.Find(ctx)
	if err != nil {
		return dto.LandingResponse{}, err
	}

	// Find landing affiliates content with repository
	landingAffiliatesContent, err := s.landingAffiliatesContentRepository.Find(ctx)
	if err != nil {
		return dto.LandingResponse{}, err
	}

	// Find landing faq content with repository
	landingFaqContent, err := s.landingFaqContentRepository.Find(ctx)
	if err != nil {
		return dto.LandingResponse{}, err
	}

	// Find landing menu with repository
	landingMenus, err := s.landingMenuRepository.FindAll(ctx, repository.FindAllOptions{})
	if err != nil {
		return dto.LandingResponse{}, err
	}

	// Map entity into response
	response, err := s.landingMapper.MapEntityToResponse(
		ctx,
		s.landingSectionHeaderRepository,
		s.landingPackageItemRepository,
		s.landingPackageDetailRepository,
		s.landingPackageDetailItemRepository,
		s.landingFeaturesContentBenefitRepository,
		s.landingMomentsContentImageRepository,
		s.landingAffiliatesContentAffiliateRepository,
		s.landingFaqContentFaqRepository,
		s.imageRepository,
		s.packageRepository,
		s.packageSessionRepository,
		s.itineraryRepository,
		s.itineraryDayRepository,
		s.flightRouteRepository,
		s.flightRepository,
		s.airlineRepository,
		landingHeroContent,
		landingSinglePackageContent,
		landingPackagesContent,
		landingFeaturesContent,
		landingMomentsContent,
		landingAffiliatesContent,
		landingFaqContent,
		landingMenus,
	)
	if err != nil {
		return dto.LandingResponse{}, err
	}

	return response, nil
}

func (s landingServiceImpl) UpdateLanding(ctx context.Context, request dto.LandingRequest) (dto.LandingResponse, error) {
	// Validate request
	if err := s.landingValidator.ValidateRequest(request); err != nil {
		return dto.LandingResponse{}, err
	}

	// Create response
	response := dto.LandingResponse{}

	// Start transaction with unit of work
	err := s.unitOfWork.Do(ctx, func(ctx context.Context, factory repository.Factory) error {
		// Create landing repositories
		landingHeroContentRepository := factory.NewLandingHeroContentRepository()
		landingSectionHeaderRepository := factory.NewLandingSectionHeaderRepository()
		landingPackageItemRepository := factory.NewLandingPackageItemRepository()
		landingSinglePackageContentRepository := factory.NewLandingSinglePackageContentRepository()
		landingPackageDetailRepository := factory.NewLandingPackageDetailRepository()
		landingPackageDetailItemRepository := factory.NewLandingPackageDetailItemRepository()
		landingPackagesContentRepository := factory.NewLandingPackagesContentRepository()
		landingFeaturesContentRepository := factory.NewLandingFeaturesContentRepository()
		landingFeaturesContentBenefitRepository := factory.NewLandingFeaturesContentBenefitRepository()
		landingMomentsContentRepository := factory.NewLandingMomentsContentRepository()
		landingMomentsContentImageRepository := factory.NewLandingMomentsContentImageRepository()
		landingAffiliatesContentRepository := factory.NewLandingAffiliatesContentRepository()
		landingAffiliatesContentAffiliateRepository := factory.NewLandingAffiliatesContentAffiliateRepository()
		landingFaqContentRepository := factory.NewLandingFaqContentRepository()
		landingFaqContentFaqRepository := factory.NewLandingFaqContentFaqRepository()
		landingMenuRepository := factory.NewLandingMenuRepository()

		// Create image repository
		imageRepository := factory.NewImageRepository()

		// Create package repository
		packageRepository := factory.NewPackageRepository()

		// Create package session repository
		packageSessionRepository := factory.NewPackageSessionRepository()

		// Create itinerary repository
		itineraryRepository := factory.NewItineraryRepository()

		// Create itinerary day repository
		itineraryDayRepository := factory.NewItineraryDayRepository()

		// Create flight route repository
		flightRouteRepository := factory.NewFlightRouteRepository()

		// Create flight repository
		flightRepository := factory.NewFlightRepository()

		// Create airline repository
		airlineRepository := factory.NewAirlineRepository()

		// Update landing hero content
		landingHeroContent, err := landingHeroContentRepository.Update(
			ctx,
			entity.LandingHeroContent{
				IsEnabled:   request.HeroContent.IsEnabled,
				Title:       request.HeroContent.Title,
				Description: request.HeroContent.Description,
				TagsLine:    request.HeroContent.TagsLine,
				ButtonLabel: request.HeroContent.ButtonLabel,
				ButtonUrl:   request.HeroContent.ButtonUrl,
				ImageId:     request.HeroContent.Image,
			},
		)
		if err != nil {
			return err
		}

		// Update landing single package content
		landingSinglePackageContent, err := landingSinglePackageContentRepository.Update(
			ctx,
			entity.LandingSinglePackageContent{
				IsEnabled: request.SinglePackageContent.IsEnabled,
			},
		)
		if err != nil {
			return err
		}

		if _, err := landingSectionHeaderRepository.Update(
			ctx,
			landingSinglePackageContent.LandingSectionHeaderId,
			entity.LandingSectionHeader{
				IsEnabled: request.SinglePackageContent.Header.IsEnabled,
				Title:     request.SinglePackageContent.Header.Title,
				Subtitle:  request.SinglePackageContent.Header.Subtitle,
				TagsLine:  request.SinglePackageContent.Header.TagsLine,
			},
		); err != nil {
			return err
		}

		if _, err := landingPackageItemRepository.Update(
			ctx,
			landingSinglePackageContent.SilverLandingPackageItemId.Int64,
			entity.LandingPackageItem{
				IsEnabled:   request.SinglePackageContent.Silver.IsEnabled,
				ButtonLabel: request.SinglePackageContent.Silver.ButtonLabel,
				PackageId:   request.SinglePackageContent.Silver.Package,
			},
		); err != nil {
			return err
		}

		if _, err := landingPackageItemRepository.Update(
			ctx,
			landingSinglePackageContent.GoldLandingPackageItemId.Int64,
			entity.LandingPackageItem{
				IsEnabled:   request.SinglePackageContent.Gold.IsEnabled,
				ButtonLabel: request.SinglePackageContent.Gold.ButtonLabel,
				PackageId:   request.SinglePackageContent.Gold.Package,
			},
		); err != nil {
			return err
		}

		if _, err := landingPackageItemRepository.Update(
			ctx,
			landingSinglePackageContent.PlatinumLandingPackageItemId.Int64,
			entity.LandingPackageItem{
				IsEnabled:   request.SinglePackageContent.Platinum.IsEnabled,
				ButtonLabel: request.SinglePackageContent.Platinum.ButtonLabel,
				PackageId:   request.SinglePackageContent.Platinum.Package,
			},
		); err != nil {
			return err
		}

		// Update landing packages content
		landingPackagesContent, err := landingPackagesContentRepository.Update(
			ctx,
			entity.LandingPackagesContent{
				IsEnabled: request.PackagesContent.IsEnabled,
			},
		)
		if err != nil {
			return err
		}

		landingPackagesContentSilverLandingPackageDetail, err := landingPackageDetailRepository.Update(
			ctx,
			landingPackagesContent.SilverLandingPackageDetailId.Int64,
			entity.LandingPackageDetail{
				IsEnabled: request.PackagesContent.Silver.IsEnabled,
			},
		)
		if err != nil {
			return err
		}

		if _, err = landingSectionHeaderRepository.Update(
			ctx,
			landingPackagesContentSilverLandingPackageDetail.LandingSectionHeaderId,
			entity.LandingSectionHeader{
				IsEnabled: request.PackagesContent.Silver.Header.IsEnabled,
				Title:     request.PackagesContent.Silver.Header.Title,
				Subtitle:  request.PackagesContent.Silver.Header.Subtitle,
				TagsLine:  request.PackagesContent.Silver.Header.TagsLine,
			},
		); err != nil {
			return err
		}

		landingPackagesContentSilverLandingPackageDetailItems, err := landingPackageDetailItemRepository.DeleteMany(ctx)
		if err != nil {
			return err
		}
		for _, landingPackagesContentSilverLandingPackageDetailItem := range landingPackagesContentSilverLandingPackageDetailItems {
			if _, err := landingPackageItemRepository.Delete(
				ctx,
				landingPackagesContentSilverLandingPackageDetailItem.LandingPackageItemId,
			); err != nil {
				return err
			}
		}

		landingPackagesContentSilverLandingPackageDetailItems = make(
			[]entity.LandingPackageDetailItem, len(request.PackagesContent.Silver.Packages),
		)
		for i, packageItem := range request.PackagesContent.Silver.Packages {
			landingPackagesContentSilverLandingPackageItem, err := landingPackageItemRepository.Create(
				ctx,
				entity.LandingPackageItem{
					IsEnabled:   packageItem.IsEnabled,
					PackageId:   packageItem.Package,
					ButtonLabel: packageItem.ButtonLabel,
				},
			)
			if err != nil {
				return err
			}

			landingPackagesContentSilverLandingPackageDetailItems[i] = entity.LandingPackageDetailItem{
				IsEnabled:              packageItem.IsEnabled,
				LandingPackageDetailId: landingPackagesContentSilverLandingPackageDetail.Id,
				LandingPackageItemId:   landingPackagesContentSilverLandingPackageItem.Id,
			}
		}
		landingPackagesContentSilverLandingPackageDetailItems, err = landingPackageDetailItemRepository.CreateMany(
			ctx, landingPackagesContentSilverLandingPackageDetailItems,
		)
		if err != nil {
			return err
		}

		landingPackagesContentGoldLandingPackageDetail, err := landingPackageDetailRepository.Update(
			ctx,
			landingPackagesContent.GoldLandingPackageDetailId.Int64,
			entity.LandingPackageDetail{
				IsEnabled: request.PackagesContent.Gold.IsEnabled,
			},
		)
		if err != nil {
			return err
		}

		if _, err = landingSectionHeaderRepository.Update(
			ctx,
			landingPackagesContentGoldLandingPackageDetail.LandingSectionHeaderId,
			entity.LandingSectionHeader{
				IsEnabled: request.PackagesContent.Gold.Header.IsEnabled,
				Title:     request.PackagesContent.Gold.Header.Title,
				Subtitle:  request.PackagesContent.Gold.Header.Subtitle,
				TagsLine:  request.PackagesContent.Gold.Header.TagsLine,
			},
		); err != nil {
			return err
		}

		landingPackagesContentGoldLandingPackageDetailItems, err := landingPackageDetailItemRepository.DeleteMany(ctx)
		if err != nil {
			return err
		}
		for _, landingPackagesContentGoldLandingPackageDetailItem := range landingPackagesContentGoldLandingPackageDetailItems {
			if _, err := landingPackageItemRepository.Delete(
				ctx,
				landingPackagesContentGoldLandingPackageDetailItem.LandingPackageItemId,
			); err != nil {
				return err
			}
		}

		landingPackagesContentGoldLandingPackageDetailItems = make(
			[]entity.LandingPackageDetailItem, len(request.PackagesContent.Gold.Packages),
		)
		for i, packageItem := range request.PackagesContent.Gold.Packages {
			landingPackagesContentGoldLandingPackageItem, err := landingPackageItemRepository.Create(
				ctx,
				entity.LandingPackageItem{
					IsEnabled:   packageItem.IsEnabled,
					PackageId:   packageItem.Package,
					ButtonLabel: packageItem.ButtonLabel,
				},
			)
			if err != nil {
				return err
			}

			landingPackagesContentGoldLandingPackageDetailItems[i] = entity.LandingPackageDetailItem{
				IsEnabled:              packageItem.IsEnabled,
				LandingPackageDetailId: landingPackagesContentGoldLandingPackageDetail.Id,
				LandingPackageItemId:   landingPackagesContentGoldLandingPackageItem.Id,
			}
		}
		landingPackagesContentGoldLandingPackageDetailItems, err = landingPackageDetailItemRepository.CreateMany(
			ctx, landingPackagesContentGoldLandingPackageDetailItems,
		)
		if err != nil {
			return err
		}

		landingPackagesContentPlatinumLandingPackageDetail, err := landingPackageDetailRepository.Update(
			ctx,
			landingPackagesContent.PlatinumLandingPackageDetailId.Int64,
			entity.LandingPackageDetail{
				IsEnabled: request.PackagesContent.Platinum.IsEnabled,
			},
		)
		if err != nil {
			return err
		}

		if _, err = landingSectionHeaderRepository.Update(
			ctx,
			landingPackagesContentPlatinumLandingPackageDetail.LandingSectionHeaderId,
			entity.LandingSectionHeader{
				IsEnabled: request.PackagesContent.Platinum.Header.IsEnabled,
				Title:     request.PackagesContent.Platinum.Header.Title,
				Subtitle:  request.PackagesContent.Platinum.Header.Subtitle,
				TagsLine:  request.PackagesContent.Platinum.Header.TagsLine,
			},
		); err != nil {
			return err
		}

		landingPackagesContentPlatinumLandingPackageDetailItems, err := landingPackageDetailItemRepository.DeleteMany(ctx)
		if err != nil {
			return err
		}
		for _, landingPackagesContentPlatinumLandingPackageDetailItem := range landingPackagesContentPlatinumLandingPackageDetailItems {
			if _, err := landingPackageItemRepository.Delete(
				ctx,
				landingPackagesContentPlatinumLandingPackageDetailItem.LandingPackageItemId,
			); err != nil {
				return err
			}
		}

		landingPackagesContentPlatinumLandingPackageDetailItems = make(
			[]entity.LandingPackageDetailItem, len(request.PackagesContent.Platinum.Packages),
		)
		for index, packageItem := range request.PackagesContent.Platinum.Packages {
			landingPackagesContentPlatinumLandingPackageItem, err := landingPackageItemRepository.Create(
				ctx,
				entity.LandingPackageItem{
					IsEnabled:   packageItem.IsEnabled,
					PackageId:   packageItem.Package,
					ButtonLabel: packageItem.ButtonLabel,
				},
			)
			if err != nil {
				return err
			}

			landingPackagesContentPlatinumLandingPackageDetailItems[index] = entity.LandingPackageDetailItem{
				IsEnabled:              packageItem.IsEnabled,
				LandingPackageDetailId: landingPackagesContentPlatinumLandingPackageDetail.Id,
				LandingPackageItemId:   landingPackagesContentPlatinumLandingPackageItem.Id,
			}
		}
		landingPackagesContentPlatinumLandingPackageDetailItems, err = landingPackageDetailItemRepository.CreateMany(
			ctx, landingPackagesContentPlatinumLandingPackageDetailItems,
		)
		if err != nil {
			return err
		}

		// Update landing features content
		landingFeaturesContent, err := landingFeaturesContentRepository.Update(
			ctx,
			entity.LandingFeaturesContent{
				IsEnabled:     request.FeaturesContent.IsEnabled,
				FooterTitle:   request.FeaturesContent.FooterTitle,
				ButtonAbout:   request.FeaturesContent.ButtonAbout,
				ButtonPackage: request.FeaturesContent.ButtonPackage,
			},
		)
		if err != nil {
			return err
		}

		if _, err = landingSectionHeaderRepository.Update(
			ctx,
			landingFeaturesContent.LandingSectionHeaderId,
			entity.LandingSectionHeader{
				IsEnabled: request.FeaturesContent.Header.IsEnabled,
				Title:     request.FeaturesContent.Header.Title,
				Subtitle:  request.FeaturesContent.Header.Subtitle,
				TagsLine:  request.FeaturesContent.Header.TagsLine,
			},
		); err != nil {
			return err
		}

		if _, err = landingFeaturesContentBenefitRepository.DeleteMany(ctx); err != nil {
			return err
		}

		benefits := make([]entity.LandingFeaturesContentBenefit, len(request.FeaturesContent.Benefits))
		for index, benefit := range request.FeaturesContent.Benefits {
			benefits[index] = entity.LandingFeaturesContentBenefit{
				IsEnabled: benefit.IsEnabled,
				Title:     benefit.Title,
				Subtitle:  benefit.Subtitle,
				LogoId:    benefit.Logo,
			}
		}
		if _, err = landingFeaturesContentBenefitRepository.CreateMany(ctx, benefits); err != nil {
			return err
		}

		// Update landing moments content with repository
		landingMomentsContent, err := landingMomentsContentRepository.Update(
			ctx,
			entity.LandingMomentsContent{
				IsEnabled: request.MomentsContent.IsEnabled,
			},
		)
		if err != nil {
			return err
		}

		if _, err = landingSectionHeaderRepository.Update(
			ctx,
			landingMomentsContent.LandingSectionHeaderId,
			entity.LandingSectionHeader{
				IsEnabled: request.MomentsContent.Header.IsEnabled,
				Title:     request.MomentsContent.Header.Title,
				Subtitle:  request.MomentsContent.Header.Subtitle,
				TagsLine:  request.MomentsContent.Header.TagsLine,
			},
		); err != nil {
			return err
		}

		if _, err = landingMomentsContentImageRepository.DeleteMany(
			ctx,
		); err != nil {
			return err
		}

		landingMomentsContentImages := make([]entity.LandingMomentsContentImage, len(request.MomentsContent.Images))
		for index, image := range request.MomentsContent.Images {
			landingMomentsContentImages[index] = entity.LandingMomentsContentImage{
				IsEnabled: image.IsEnabled,
				ImageId:   image.Image,
			}
		}
		if _, err = landingMomentsContentImageRepository.CreateMany(
			ctx,
			landingMomentsContentImages,
		); err != nil {
			return err
		}

		// Update landing affiliates content with repository
		landingAffiliatesContent, err := landingAffiliatesContentRepository.Update(
			ctx,
			entity.LandingAffiliatesContent{
				IsEnabled: request.AffiliatesContent.IsEnabled,
			},
		)
		if err != nil {
			return err
		}

		if _, err = landingSectionHeaderRepository.Update(
			ctx,
			landingAffiliatesContent.LandingSectionHeaderId,
			entity.LandingSectionHeader{
				IsEnabled: request.AffiliatesContent.Header.IsEnabled,
				Title:     request.AffiliatesContent.Header.Title,
				Subtitle:  request.AffiliatesContent.Header.Subtitle,
				TagsLine:  request.AffiliatesContent.Header.TagsLine,
			},
		); err != nil {
			return err
		}

		if _, err = landingAffiliatesContentAffiliateRepository.DeleteMany(ctx); err != nil {
			return err
		}

		affiliates := make([]entity.LandingAffiliatesContentAffiliate, len(request.AffiliatesContent.Affiliates))
		for index, affiliate := range request.AffiliatesContent.Affiliates {
			affiliates[index] = entity.LandingAffiliatesContentAffiliate{
				IsEnabled: affiliate.IsEnabled,
				Name:      affiliate.Name,
				LogoId:    affiliate.Logo,
				Width:     affiliate.Width,
				Height:    affiliate.Height,
			}
		}
		if _, err = landingAffiliatesContentAffiliateRepository.CreateMany(ctx, affiliates); err != nil {
			return err
		}

		// Update landing faq content with repository
		landingFaqContent, err := landingFaqContentRepository.Update(
			ctx,
			entity.LandingFaqContent{
				IsEnabled: request.FaqContent.IsEnabled,
			},
		)
		if err != nil {
			return err
		}

		if _, err = landingSectionHeaderRepository.Update(
			ctx,
			landingFaqContent.LandingSectionHeaderId,
			entity.LandingSectionHeader{
				IsEnabled: request.FaqContent.Header.IsEnabled,
				Title:     request.FaqContent.Header.Title,
				Subtitle:  request.FaqContent.Header.Subtitle,
				TagsLine:  request.FaqContent.Header.TagsLine,
			},
		); err != nil {
			return err
		}

		if _, err = landingFaqContentFaqRepository.DeleteMany(ctx); err != nil {
			return err
		}

		faqs := make([]entity.LandingFaqContentFaq, len(request.FaqContent.Faqs))
		for index, faq := range request.FaqContent.Faqs {
			faqs[index] = entity.LandingFaqContentFaq{
				IsEnabled: faq.IsEnabled,
				Question:  faq.Question,
				Answer:    faq.Answer,
			}
		}
		if _, err = landingFaqContentFaqRepository.CreateMany(ctx, faqs); err != nil {
			return err
		}

		// Update landing menu with repository
		if _, err = landingMenuRepository.DeleteMany(ctx); err != nil {
			return err
		}

		landingMenus := make([]entity.LandingMenu, len(request.Menus))
		for index, menu := range request.Menus {
			landingMenus[index] = entity.LandingMenu{
				IsEnabled: menu.IsEnabled,
				Icon:      menu.Icon,
				Label:     menu.Label,
				Path:      menu.Path,
			}
		}
		if _, err = landingMenuRepository.CreateMany(ctx, landingMenus); err != nil {
			return err
		}

		// Map entity into response
		response, err = s.landingMapper.MapEntityToResponse(
			ctx,
			landingSectionHeaderRepository,
			landingPackageItemRepository,
			landingPackageDetailRepository,
			landingPackageDetailItemRepository,
			landingFeaturesContentBenefitRepository,
			landingMomentsContentImageRepository,
			landingAffiliatesContentAffiliateRepository,
			landingFaqContentFaqRepository,
			imageRepository,
			packageRepository,
			packageSessionRepository,
			itineraryRepository,
			itineraryDayRepository,
			flightRouteRepository,
			flightRepository,
			airlineRepository,
			landingHeroContent,
			landingSinglePackageContent,
			landingPackagesContent,
			landingFeaturesContent,
			landingMomentsContent,
			landingAffiliatesContent,
			landingFaqContent,
			landingMenus,
		)
		return err
	})

	return response, err
}
