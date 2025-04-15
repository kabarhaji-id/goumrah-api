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
	landingTestimonialContentRepository         repository.LandingTestimonialContentRepository
	landingTestimonialContentReviewRepository   repository.LandingTestimonialContentReviewRepository
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
	landingTestimonialContentRepository repository.LandingTestimonialContentRepository,
	landingTestimonialContentReviewRepository repository.LandingTestimonialContentReviewRepository,
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
		landingTestimonialContentRepository,
		landingTestimonialContentReviewRepository,
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
		landingTestimonialContentRepository := factory.NewLandingTestimonialContentRepository()
		landingTestimonialContentReviewRepository := factory.NewLandingTestimonialContentReviewRepository()
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
				IsMobile:    request.HeroContent.IsMobile,
				IsDesktop:   request.HeroContent.IsDesktop,
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
				IsMobile:  request.SinglePackageContent.Header.IsMobile,
				IsDesktop: request.SinglePackageContent.Header.IsDesktop,
				Title:     request.SinglePackageContent.Header.Title,
				Subtitle:  request.SinglePackageContent.Header.Subtitle,
				TagsLine:  request.SinglePackageContent.Header.TagsLine,
			},
		)
		if err != nil {
			return err
		}

		landingSinglePackageContentSilverLandingPackageItemId := null.NewInt(0, false)
		if request.SinglePackageContent.Silver.Valid {
			landingSinglePackageContentSilverLandingPackageItem, err := landingPackageItemRepository.Create(
				ctx,
				entity.LandingPackageItem{
					IsEnabled:   request.SinglePackageContent.Silver.V.IsEnabled,
					IsMobile:    request.SinglePackageContent.Silver.V.IsMobile,
					IsDesktop:   request.SinglePackageContent.Silver.V.IsDesktop,
					ButtonLabel: request.SinglePackageContent.Silver.V.ButtonLabel,
					PackageId:   request.SinglePackageContent.Silver.V.Package,
				},
			)
			if err != nil {
				return err
			}

			landingSinglePackageContentSilverLandingPackageItemId = null.IntFrom(landingSinglePackageContentSilverLandingPackageItem.Id)
		}

		landingSinglePackageContentGoldLandingPackageItemId := null.NewInt(0, false)
		if request.SinglePackageContent.Gold.Valid {
			landingSinglePackageContentGoldLandingPackageItem, err := landingPackageItemRepository.Create(
				ctx,
				entity.LandingPackageItem{
					IsEnabled:   request.SinglePackageContent.Gold.V.IsEnabled,
					IsMobile:    request.SinglePackageContent.Gold.V.IsMobile,
					IsDesktop:   request.SinglePackageContent.Gold.V.IsDesktop,
					ButtonLabel: request.SinglePackageContent.Gold.V.ButtonLabel,
					PackageId:   request.SinglePackageContent.Gold.V.Package,
				},
			)
			if err != nil {
				return err
			}

			landingSinglePackageContentGoldLandingPackageItemId = null.IntFrom(landingSinglePackageContentGoldLandingPackageItem.Id)
		}

		landingSinglePackageContentPlatinumLandingPackageItemId := null.NewInt(0, false)
		if request.SinglePackageContent.Platinum.Valid {
			landingSinglePackageContentPlatinumLandingPackageItem, err := landingPackageItemRepository.Create(
				ctx,
				entity.LandingPackageItem{
					IsEnabled:   request.SinglePackageContent.Platinum.V.IsEnabled,
					IsMobile:    request.SinglePackageContent.Platinum.V.IsMobile,
					IsDesktop:   request.SinglePackageContent.Platinum.V.IsDesktop,
					ButtonLabel: request.SinglePackageContent.Platinum.V.ButtonLabel,
					PackageId:   request.SinglePackageContent.Platinum.V.Package,
				},
			)
			if err != nil {
				return err
			}

			landingSinglePackageContentPlatinumLandingPackageItemId = null.IntFrom(landingSinglePackageContentPlatinumLandingPackageItem.Id)
		}

		landingSinglePackageContent, err := landingSinglePackageContentRepository.Create(
			ctx,
			entity.LandingSinglePackageContent{
				IsEnabled:                    request.SinglePackageContent.IsEnabled,
				IsMobile:                     request.SinglePackageContent.IsMobile,
				IsDesktop:                    request.SinglePackageContent.IsDesktop,
				LandingSectionHeaderId:       landingSinglePackageContentHeader.Id,
				SilverLandingPackageItemId:   landingSinglePackageContentSilverLandingPackageItemId,
				GoldLandingPackageItemId:     landingSinglePackageContentGoldLandingPackageItemId,
				PlatinumLandingPackageItemId: landingSinglePackageContentPlatinumLandingPackageItemId,
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
				IsMobile:  request.PackagesContent.Silver.Header.IsMobile,
				IsDesktop: request.PackagesContent.Silver.Header.IsDesktop,
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
				IsMobile:               request.PackagesContent.Silver.IsMobile,
				IsDesktop:              request.PackagesContent.Silver.IsDesktop,
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
					IsMobile:    packageItem.IsMobile,
					IsDesktop:   packageItem.IsDesktop,
					PackageId:   packageItem.Package,
					ButtonLabel: packageItem.ButtonLabel,
				},
			)
			if err != nil {
				return err
			}

			landingPackagesContentSilverLandingPackageDetailItems[i] = entity.LandingPackageDetailItem{
				IsEnabled:              packageItem.IsEnabled,
				IsMobile:               packageItem.IsMobile,
				IsDesktop:              packageItem.IsDesktop,
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
				IsMobile:  request.PackagesContent.Gold.Header.IsMobile,
				IsDesktop: request.PackagesContent.Gold.Header.IsDesktop,
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
				IsMobile:               request.PackagesContent.Gold.IsMobile,
				IsDesktop:              request.PackagesContent.Gold.IsDesktop,
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
					IsMobile:    packageItem.IsMobile,
					IsDesktop:   packageItem.IsDesktop,
					PackageId:   packageItem.Package,
					ButtonLabel: packageItem.ButtonLabel,
				},
			)
			if err != nil {
				return err
			}

			landingPackagesContentGoldLandingPackageDetailItems[i] = entity.LandingPackageDetailItem{
				IsEnabled:              packageItem.IsEnabled,
				IsMobile:               packageItem.IsMobile,
				IsDesktop:              packageItem.IsDesktop,
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
				IsMobile:  request.PackagesContent.Platinum.Header.IsMobile,
				IsDesktop: request.PackagesContent.Platinum.Header.IsDesktop,
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
				IsMobile:               request.PackagesContent.Platinum.IsMobile,
				IsDesktop:              request.PackagesContent.Platinum.IsDesktop,
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
					IsMobile:    packageItem.IsMobile,
					IsDesktop:   packageItem.IsDesktop,
					PackageId:   packageItem.Package,
					ButtonLabel: packageItem.ButtonLabel,
				},
			)
			if err != nil {
				return err
			}

			landingPackagesContentPlatinumLandingPackageDetailItems[i] = entity.LandingPackageDetailItem{
				IsEnabled:              packageItem.IsEnabled,
				IsMobile:               packageItem.IsMobile,
				IsDesktop:              packageItem.IsDesktop,
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
				IsMobile:                       request.PackagesContent.IsMobile,
				IsDesktop:                      request.PackagesContent.IsDesktop,
				SilverLandingPackageDetailId:   landingPackagesContentSilverLandingPackageDetail.Id,
				GoldLandingPackageDetailId:     landingPackagesContentGoldLandingPackageDetail.Id,
				PlatinumLandingPackageDetailId: landingPackagesContentPlatinumLandingPackageDetail.Id,
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
				IsMobile:  request.FeaturesContent.Header.IsMobile,
				IsDesktop: request.FeaturesContent.Header.IsDesktop,
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
				IsMobile:  benefit.IsMobile,
				IsDesktop: benefit.IsDesktop,
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
				IsMobile:               request.FeaturesContent.IsMobile,
				IsDesktop:              request.FeaturesContent.IsDesktop,
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
				IsMobile:  request.MomentsContent.Header.IsMobile,
				IsDesktop: request.MomentsContent.Header.IsDesktop,
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
				IsMobile:  image.IsMobile,
				IsDesktop: image.IsDesktop,
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
				IsMobile:               request.MomentsContent.IsMobile,
				IsDesktop:              request.MomentsContent.IsDesktop,
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
				IsMobile:  request.AffiliatesContent.Header.IsMobile,
				IsDesktop: request.AffiliatesContent.Header.IsDesktop,
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
				IsMobile:  affiliate.IsMobile,
				IsDesktop: affiliate.IsDesktop,
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
				IsMobile:               request.AffiliatesContent.IsMobile,
				IsDesktop:              request.AffiliatesContent.IsDesktop,
				LandingSectionHeaderId: landingAffiliatesContentHeader.Id,
			},
		)
		if err != nil {
			return err
		}

		// Create landing testimonial content with repository
		landingTestimonialContentHeader, err := landingSectionHeaderRepository.Create(
			ctx,
			entity.LandingSectionHeader{
				IsEnabled: request.TestimonialContent.Header.IsEnabled,
				IsMobile:  request.TestimonialContent.Header.IsMobile,
				IsDesktop: request.TestimonialContent.Header.IsDesktop,
				Title:     request.TestimonialContent.Header.Title,
				Subtitle:  request.TestimonialContent.Header.Subtitle,
				TagsLine:  request.TestimonialContent.Header.TagsLine,
			},
		)
		if err != nil {
			return err
		}

		reviews := make([]entity.LandingTestimonialContentReview, len(request.TestimonialContent.Reviews))
		for index, review := range request.TestimonialContent.Reviews {
			reviews[index] = entity.LandingTestimonialContentReview{
				IsEnabled: review.IsEnabled,
				IsMobile:  review.IsMobile,
				IsDesktop: review.IsDesktop,
				Reviewer:  review.Reviewer,
				Age:       review.Age,
				Address:   review.Address,
				Rating:    review.Rating,
				Review:    review.Review,
			}
		}
		if _, err = landingTestimonialContentReviewRepository.CreateMany(ctx, reviews); err != nil {
			return err
		}

		landingTestimonialContent, err := landingTestimonialContentRepository.Create(
			ctx,
			entity.LandingTestimonialContent{
				IsEnabled:              request.TestimonialContent.IsEnabled,
				IsMobile:               request.TestimonialContent.IsMobile,
				IsDesktop:              request.TestimonialContent.IsDesktop,
				LandingSectionHeaderId: landingTestimonialContentHeader.Id,
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
				IsMobile:  request.FaqContent.Header.IsMobile,
				IsDesktop: request.FaqContent.Header.IsDesktop,
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
				IsMobile:  faq.IsMobile,
				IsDesktop: faq.IsDesktop,
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
				IsMobile:               request.FaqContent.IsMobile,
				IsDesktop:              request.FaqContent.IsDesktop,
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
				IsMobile:  menu.IsMobile,
				IsDesktop: menu.IsDesktop,
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
			landingTestimonialContentReviewRepository,
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
			landingTestimonialContent,
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

	// Find landing testimonial content with repository
	landingTestimonialContent, err := s.landingTestimonialContentRepository.Find(ctx)
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
		s.landingTestimonialContentReviewRepository,
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
		landingTestimonialContent,
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
		landingTestimonialContentRepository := factory.NewLandingTestimonialContentRepository()
		landingTestimonialContentReviewRepository := factory.NewLandingTestimonialContentReviewRepository()
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
				IsMobile:    request.HeroContent.IsMobile,
				IsDesktop:   request.HeroContent.IsDesktop,
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
		landingSinglePackageContent, err := landingSinglePackageContentRepository.Find(ctx)
		if err != nil {
			return err
		}
		landingSinglePackageContent.IsEnabled = request.SinglePackageContent.IsEnabled
		landingSinglePackageContent.IsMobile = request.SinglePackageContent.IsMobile
		landingSinglePackageContent.IsDesktop = request.SinglePackageContent.IsDesktop
		if request.SinglePackageContent.Silver.Valid {
			if landingSinglePackageContent.SilverLandingPackageItemId.Valid {
				if _, err := landingPackageItemRepository.Update(
					ctx,
					landingSinglePackageContent.SilverLandingPackageItemId.Int64,
					entity.LandingPackageItem{
						IsEnabled:   request.SinglePackageContent.Silver.V.IsEnabled,
						IsMobile:    request.SinglePackageContent.Silver.V.IsMobile,
						IsDesktop:   request.SinglePackageContent.Silver.V.IsDesktop,
						ButtonLabel: request.SinglePackageContent.Silver.V.ButtonLabel,
						PackageId:   request.SinglePackageContent.Silver.V.Package,
					},
				); err != nil {
					return err
				}
			} else {
				landingSinglePackageContentSilverLandingPackageItem, err := landingPackageItemRepository.Create(
					ctx,
					entity.LandingPackageItem{
						IsEnabled:   request.SinglePackageContent.Silver.V.IsEnabled,
						IsMobile:    request.SinglePackageContent.Silver.V.IsMobile,
						IsDesktop:   request.SinglePackageContent.Silver.V.IsDesktop,
						ButtonLabel: request.SinglePackageContent.Silver.V.ButtonLabel,
						PackageId:   request.SinglePackageContent.Silver.V.Package,
					},
				)
				if err != nil {
					return err
				}

				landingSinglePackageContent.SilverLandingPackageItemId = null.IntFrom(landingSinglePackageContentSilverLandingPackageItem.Id)
			}
		} else if landingSinglePackageContent.SilverLandingPackageItemId.Valid {
			if _, err := landingPackageItemRepository.Delete(
				ctx,
				landingSinglePackageContent.SilverLandingPackageItemId.Int64,
			); err != nil {
				return err
			}
		}
		if request.SinglePackageContent.Gold.Valid {
			if landingSinglePackageContent.GoldLandingPackageItemId.Valid {
				if _, err := landingPackageItemRepository.Update(
					ctx,
					landingSinglePackageContent.GoldLandingPackageItemId.Int64,
					entity.LandingPackageItem{
						IsEnabled:   request.SinglePackageContent.Gold.V.IsEnabled,
						IsMobile:    request.SinglePackageContent.Gold.V.IsMobile,
						IsDesktop:   request.SinglePackageContent.Gold.V.IsDesktop,
						ButtonLabel: request.SinglePackageContent.Gold.V.ButtonLabel,
						PackageId:   request.SinglePackageContent.Gold.V.Package,
					},
				); err != nil {
					return err
				}
			} else {
				landingSinglePackageContentGoldLandingPackageItem, err := landingPackageItemRepository.Create(
					ctx,
					entity.LandingPackageItem{
						IsEnabled:   request.SinglePackageContent.Gold.V.IsEnabled,
						IsMobile:    request.SinglePackageContent.Gold.V.IsMobile,
						IsDesktop:   request.SinglePackageContent.Gold.V.IsDesktop,
						ButtonLabel: request.SinglePackageContent.Gold.V.ButtonLabel,
						PackageId:   request.SinglePackageContent.Gold.V.Package,
					},
				)
				if err != nil {
					return err
				}

				landingSinglePackageContent.GoldLandingPackageItemId = null.IntFrom(landingSinglePackageContentGoldLandingPackageItem.Id)
			}
		} else if landingSinglePackageContent.GoldLandingPackageItemId.Valid {
			if _, err := landingPackageItemRepository.Delete(
				ctx,
				landingSinglePackageContent.GoldLandingPackageItemId.Int64,
			); err != nil {
				return err
			}
		}
		if request.SinglePackageContent.Platinum.Valid {
			if landingSinglePackageContent.PlatinumLandingPackageItemId.Valid {
				if _, err := landingPackageItemRepository.Update(
					ctx,
					landingSinglePackageContent.PlatinumLandingPackageItemId.Int64,
					entity.LandingPackageItem{
						IsEnabled:   request.SinglePackageContent.Platinum.V.IsEnabled,
						IsMobile:    request.SinglePackageContent.Platinum.V.IsMobile,
						IsDesktop:   request.SinglePackageContent.Platinum.V.IsDesktop,
						ButtonLabel: request.SinglePackageContent.Platinum.V.ButtonLabel,
						PackageId:   request.SinglePackageContent.Platinum.V.Package,
					},
				); err != nil {
					return err
				}
			} else {
				landingSinglePackageContentPlatinumLandingPackageItem, err := landingPackageItemRepository.Create(
					ctx,
					entity.LandingPackageItem{
						IsEnabled:   request.SinglePackageContent.Platinum.V.IsEnabled,
						IsMobile:    request.SinglePackageContent.Platinum.V.IsMobile,
						IsDesktop:   request.SinglePackageContent.Platinum.V.IsDesktop,
						ButtonLabel: request.SinglePackageContent.Platinum.V.ButtonLabel,
						PackageId:   request.SinglePackageContent.Platinum.V.Package,
					},
				)
				if err != nil {
					return err
				}

				landingSinglePackageContent.PlatinumLandingPackageItemId = null.IntFrom(landingSinglePackageContentPlatinumLandingPackageItem.Id)
			}
		} else if landingSinglePackageContent.PlatinumLandingPackageItemId.Valid {
			if _, err := landingPackageItemRepository.Delete(
				ctx,
				landingSinglePackageContent.PlatinumLandingPackageItemId.Int64,
			); err != nil {
				return err
			}
		}

		landingSinglePackageContent, err = landingSinglePackageContentRepository.Update(ctx, landingSinglePackageContent)
		if err != nil {
			return err
		}

		if _, err := landingSectionHeaderRepository.Update(
			ctx,
			landingSinglePackageContent.LandingSectionHeaderId,
			entity.LandingSectionHeader{
				IsEnabled: request.SinglePackageContent.Header.IsEnabled,
				IsMobile:  request.SinglePackageContent.Header.IsMobile,
				IsDesktop: request.SinglePackageContent.Header.IsDesktop,
				Title:     request.SinglePackageContent.Header.Title,
				Subtitle:  request.SinglePackageContent.Header.Subtitle,
				TagsLine:  request.SinglePackageContent.Header.TagsLine,
			},
		); err != nil {
			return err
		}

		// Update landing packages content
		landingPackagesContent, err := landingPackagesContentRepository.Find(ctx)
		if err != nil {
			return err
		}
		landingPackagesContent.IsEnabled = request.PackagesContent.IsEnabled
		landingPackagesContent.IsMobile = request.PackagesContent.IsMobile
		landingPackagesContent.IsDesktop = request.PackagesContent.IsDesktop

		landingPackagesContent, err = landingPackagesContentRepository.Update(ctx, landingPackagesContent)
		if err != nil {
			return err
		}

		landingPackageDetailItems, err := landingPackageDetailItemRepository.DeleteMany(ctx)
		if err != nil {
			return err
		}
		for _, landingPackageDetailItem := range landingPackageDetailItems {
			if _, err := landingPackageItemRepository.Delete(
				ctx,
				landingPackageDetailItem.LandingPackageItemId,
			); err != nil {
				return err
			}
		}

		landingPackagesContentSilverLandingPackageDetail, err := landingPackageDetailRepository.FindById(
			ctx,
			landingPackagesContent.SilverLandingPackageDetailId,
		)
		if err != nil {
			return err
		}
		landingPackagesContentSilverLandingPackageDetail.IsEnabled = request.PackagesContent.Silver.IsEnabled
		landingPackagesContentSilverLandingPackageDetail.IsMobile = request.PackagesContent.Silver.IsMobile
		landingPackagesContentSilverLandingPackageDetail.IsDesktop = request.PackagesContent.Silver.IsDesktop

		landingPackagesContentSilverLandingPackageDetail, err = landingPackageDetailRepository.Update(
			ctx,
			landingPackagesContentSilverLandingPackageDetail.Id,
			landingPackagesContentSilverLandingPackageDetail,
		)
		if err != nil {
			return err
		}

		if _, err = landingSectionHeaderRepository.Update(
			ctx,
			landingPackagesContentSilverLandingPackageDetail.LandingSectionHeaderId,
			entity.LandingSectionHeader{
				IsEnabled: request.PackagesContent.Silver.Header.IsEnabled,
				IsMobile:  request.PackagesContent.Silver.Header.IsMobile,
				IsDesktop: request.PackagesContent.Silver.Header.IsDesktop,
				Title:     request.PackagesContent.Silver.Header.Title,
				Subtitle:  request.PackagesContent.Silver.Header.Subtitle,
				TagsLine:  request.PackagesContent.Silver.Header.TagsLine,
			},
		); err != nil {
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
					IsMobile:    packageItem.IsMobile,
					IsDesktop:   packageItem.IsDesktop,
					PackageId:   packageItem.Package,
					ButtonLabel: packageItem.ButtonLabel,
				},
			)
			if err != nil {
				return err
			}

			landingPackagesContentSilverLandingPackageDetailItems[i] = entity.LandingPackageDetailItem{
				IsEnabled:              packageItem.IsEnabled,
				IsMobile:               packageItem.IsMobile,
				IsDesktop:              packageItem.IsDesktop,
				LandingPackageDetailId: landingPackagesContentSilverLandingPackageDetail.Id,
				LandingPackageItemId:   landingPackagesContentSilverLandingPackageItem.Id,
			}
		}
		_, err = landingPackageDetailItemRepository.CreateMany(
			ctx, landingPackagesContentSilverLandingPackageDetailItems,
		)
		if err != nil {
			return err
		}

		landingPackagesContentGoldLandingPackageDetail, err := landingPackageDetailRepository.FindById(
			ctx,
			landingPackagesContent.GoldLandingPackageDetailId,
		)
		if err != nil {
			return err
		}
		landingPackagesContentGoldLandingPackageDetail.IsEnabled = request.PackagesContent.Gold.IsEnabled
		landingPackagesContentGoldLandingPackageDetail.IsMobile = request.PackagesContent.Gold.IsMobile
		landingPackagesContentGoldLandingPackageDetail.IsDesktop = request.PackagesContent.Gold.IsDesktop

		landingPackagesContentGoldLandingPackageDetail, err = landingPackageDetailRepository.Update(
			ctx,
			landingPackagesContent.GoldLandingPackageDetailId,
			landingPackagesContentGoldLandingPackageDetail,
		)
		if err != nil {
			return err
		}

		if _, err = landingSectionHeaderRepository.Update(
			ctx,
			landingPackagesContentGoldLandingPackageDetail.LandingSectionHeaderId,
			entity.LandingSectionHeader{
				IsEnabled: request.PackagesContent.Gold.Header.IsEnabled,
				IsMobile:  request.PackagesContent.Gold.Header.IsMobile,
				IsDesktop: request.PackagesContent.Gold.Header.IsDesktop,
				Title:     request.PackagesContent.Gold.Header.Title,
				Subtitle:  request.PackagesContent.Gold.Header.Subtitle,
				TagsLine:  request.PackagesContent.Gold.Header.TagsLine,
			},
		); err != nil {
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
					IsMobile:    packageItem.IsMobile,
					IsDesktop:   packageItem.IsDesktop,
					PackageId:   packageItem.Package,
					ButtonLabel: packageItem.ButtonLabel,
				},
			)
			if err != nil {
				return err
			}

			landingPackagesContentGoldLandingPackageDetailItems[i] = entity.LandingPackageDetailItem{
				IsEnabled:              packageItem.IsEnabled,
				IsMobile:               packageItem.IsMobile,
				IsDesktop:              packageItem.IsDesktop,
				LandingPackageDetailId: landingPackagesContentGoldLandingPackageDetail.Id,
				LandingPackageItemId:   landingPackagesContentGoldLandingPackageItem.Id,
			}
		}
		_, err = landingPackageDetailItemRepository.CreateMany(
			ctx, landingPackagesContentGoldLandingPackageDetailItems,
		)
		if err != nil {
			return err
		}

		landingPackagesContentPlatinumLandingPackageDetail, err := landingPackageDetailRepository.FindById(
			ctx,
			landingPackagesContent.PlatinumLandingPackageDetailId,
		)
		if err != nil {
			return err
		}
		landingPackagesContentPlatinumLandingPackageDetail.IsEnabled = request.PackagesContent.Platinum.IsEnabled
		landingPackagesContentPlatinumLandingPackageDetail.IsMobile = request.PackagesContent.Platinum.IsMobile
		landingPackagesContentPlatinumLandingPackageDetail.IsDesktop = request.PackagesContent.Platinum.IsDesktop

		landingPackagesContentPlatinumLandingPackageDetail, err = landingPackageDetailRepository.Update(
			ctx,
			landingPackagesContent.PlatinumLandingPackageDetailId,
			landingPackagesContentPlatinumLandingPackageDetail,
		)
		if err != nil {
			return err
		}

		if _, err = landingSectionHeaderRepository.Update(
			ctx,
			landingPackagesContentPlatinumLandingPackageDetail.LandingSectionHeaderId,
			entity.LandingSectionHeader{
				IsEnabled: request.PackagesContent.Platinum.Header.IsEnabled,
				IsMobile:  request.PackagesContent.Platinum.Header.IsMobile,
				IsDesktop: request.PackagesContent.Platinum.Header.IsDesktop,
				Title:     request.PackagesContent.Platinum.Header.Title,
				Subtitle:  request.PackagesContent.Platinum.Header.Subtitle,
				TagsLine:  request.PackagesContent.Platinum.Header.TagsLine,
			},
		); err != nil {
			return err
		}

		landingPackagesContentPlatinumLandingPackageDetailItems := make(
			[]entity.LandingPackageDetailItem, len(request.PackagesContent.Platinum.Packages),
		)
		for index, packageItem := range request.PackagesContent.Platinum.Packages {
			landingPackagesContentPlatinumLandingPackageItem, err := landingPackageItemRepository.Create(
				ctx,
				entity.LandingPackageItem{
					IsEnabled:   packageItem.IsEnabled,
					IsMobile:    packageItem.IsMobile,
					IsDesktop:   packageItem.IsDesktop,
					PackageId:   packageItem.Package,
					ButtonLabel: packageItem.ButtonLabel,
				},
			)
			if err != nil {
				return err
			}

			landingPackagesContentPlatinumLandingPackageDetailItems[index] = entity.LandingPackageDetailItem{
				IsEnabled:              packageItem.IsEnabled,
				IsMobile:               packageItem.IsMobile,
				IsDesktop:              packageItem.IsDesktop,
				LandingPackageDetailId: landingPackagesContentPlatinumLandingPackageDetail.Id,
				LandingPackageItemId:   landingPackagesContentPlatinumLandingPackageItem.Id,
			}
		}
		_, err = landingPackageDetailItemRepository.CreateMany(
			ctx, landingPackagesContentPlatinumLandingPackageDetailItems,
		)
		if err != nil {
			return err
		}

		// Update landing features content
		landingFeaturesContent, err := landingFeaturesContentRepository.Find(ctx)
		if err != nil {
			return err
		}
		landingFeaturesContent.IsEnabled = request.FeaturesContent.IsEnabled
		landingFeaturesContent.IsMobile = request.FeaturesContent.IsMobile
		landingFeaturesContent.IsDesktop = request.FeaturesContent.IsDesktop
		landingFeaturesContent.FooterTitle = request.FeaturesContent.FooterTitle
		landingFeaturesContent.ButtonAbout = request.FeaturesContent.ButtonAbout
		landingFeaturesContent.ButtonPackage = request.FeaturesContent.ButtonPackage

		landingFeaturesContent, err = landingFeaturesContentRepository.Update(
			ctx,
			landingFeaturesContent,
		)
		if err != nil {
			return err
		}

		if _, err = landingSectionHeaderRepository.Update(
			ctx,
			landingFeaturesContent.LandingSectionHeaderId,
			entity.LandingSectionHeader{
				IsEnabled: request.FeaturesContent.Header.IsEnabled,
				IsMobile:  request.FeaturesContent.Header.IsMobile,
				IsDesktop: request.FeaturesContent.Header.IsDesktop,
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
				IsMobile:  benefit.IsMobile,
				IsDesktop: benefit.IsDesktop,
				Title:     benefit.Title,
				Subtitle:  benefit.Subtitle,
				LogoId:    benefit.Logo,
			}
		}
		if _, err = landingFeaturesContentBenefitRepository.CreateMany(ctx, benefits); err != nil {
			return err
		}

		// Update landing moments content with repository
		landingMomentsContent, err := landingMomentsContentRepository.Find(ctx)
		if err != nil {
			return err
		}
		landingMomentsContent.IsEnabled = request.MomentsContent.IsEnabled
		landingMomentsContent.IsMobile = request.MomentsContent.IsMobile
		landingMomentsContent.IsDesktop = request.MomentsContent.IsDesktop

		landingMomentsContent, err = landingMomentsContentRepository.Update(
			ctx,
			landingMomentsContent,
		)
		if err != nil {
			return err
		}

		if _, err = landingSectionHeaderRepository.Update(
			ctx,
			landingMomentsContent.LandingSectionHeaderId,
			entity.LandingSectionHeader{
				IsEnabled: request.MomentsContent.Header.IsEnabled,
				IsMobile:  request.MomentsContent.Header.IsMobile,
				IsDesktop: request.MomentsContent.Header.IsDesktop,
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
				IsMobile:  image.IsMobile,
				IsDesktop: image.IsDesktop,
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
		landingAffiliatesContent, err := landingAffiliatesContentRepository.Find(ctx)
		if err != nil {
			return err
		}
		landingAffiliatesContent.IsEnabled = request.AffiliatesContent.IsEnabled
		landingAffiliatesContent.IsMobile = request.AffiliatesContent.IsMobile
		landingAffiliatesContent.IsDesktop = request.AffiliatesContent.IsDesktop

		landingAffiliatesContent, err = landingAffiliatesContentRepository.Update(
			ctx,
			landingAffiliatesContent,
		)
		if err != nil {
			return err
		}

		if _, err = landingSectionHeaderRepository.Update(
			ctx,
			landingAffiliatesContent.LandingSectionHeaderId,
			entity.LandingSectionHeader{
				IsEnabled: request.AffiliatesContent.Header.IsEnabled,
				IsMobile:  request.AffiliatesContent.Header.IsMobile,
				IsDesktop: request.AffiliatesContent.Header.IsDesktop,
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
				IsMobile:  affiliate.IsMobile,
				IsDesktop: affiliate.IsDesktop,
				Name:      affiliate.Name,
				LogoId:    affiliate.Logo,
				Width:     affiliate.Width,
				Height:    affiliate.Height,
			}
		}
		if _, err = landingAffiliatesContentAffiliateRepository.CreateMany(ctx, affiliates); err != nil {
			return err
		}

		// Update landing testimonial content with repository
		landingTestimonialContent, err := landingTestimonialContentRepository.Find(ctx)
		if err != nil {
			return err
		}
		landingTestimonialContent.IsEnabled = request.TestimonialContent.IsEnabled
		landingTestimonialContent.IsMobile = request.TestimonialContent.IsMobile
		landingTestimonialContent.IsDesktop = request.TestimonialContent.IsDesktop

		landingTestimonialContent, err = landingTestimonialContentRepository.Update(
			ctx,
			landingTestimonialContent,
		)
		if err != nil {
			return err
		}

		if _, err = landingSectionHeaderRepository.Update(
			ctx,
			landingTestimonialContent.LandingSectionHeaderId,
			entity.LandingSectionHeader{
				IsEnabled: request.TestimonialContent.Header.IsEnabled,
				IsMobile:  request.TestimonialContent.Header.IsMobile,
				IsDesktop: request.TestimonialContent.Header.IsDesktop,
				Title:     request.TestimonialContent.Header.Title,
				Subtitle:  request.TestimonialContent.Header.Subtitle,
				TagsLine:  request.TestimonialContent.Header.TagsLine,
			},
		); err != nil {
			return err
		}

		if _, err = landingTestimonialContentReviewRepository.DeleteMany(ctx); err != nil {
			return err
		}

		reviews := make([]entity.LandingTestimonialContentReview, len(request.TestimonialContent.Reviews))
		for index, review := range request.TestimonialContent.Reviews {
			reviews[index] = entity.LandingTestimonialContentReview{
				IsEnabled: review.IsEnabled,
				IsMobile:  review.IsMobile,
				IsDesktop: review.IsDesktop,
				Reviewer:  review.Reviewer,
				Age:       review.Age,
				Address:   review.Address,
				Rating:    review.Rating,
				Review:    review.Review,
			}
		}
		if _, err = landingTestimonialContentReviewRepository.CreateMany(ctx, reviews); err != nil {
			return err
		}

		// Update landing faq content with repository
		landingFaqContent, err := landingFaqContentRepository.Find(ctx)
		if err != nil {
			return err
		}
		landingFaqContent.IsEnabled = request.FaqContent.IsEnabled
		landingFaqContent.IsMobile = request.FaqContent.IsMobile
		landingFaqContent.IsDesktop = request.FaqContent.IsDesktop

		landingFaqContent, err = landingFaqContentRepository.Update(
			ctx,
			landingFaqContent,
		)
		if err != nil {
			return err
		}

		if _, err = landingSectionHeaderRepository.Update(
			ctx,
			landingFaqContent.LandingSectionHeaderId,
			entity.LandingSectionHeader{
				IsEnabled: request.FaqContent.Header.IsEnabled,
				IsMobile:  request.FaqContent.Header.IsMobile,
				IsDesktop: request.FaqContent.Header.IsDesktop,
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
				IsMobile:  faq.IsMobile,
				IsDesktop: faq.IsDesktop,
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
				IsMobile:  menu.IsMobile,
				IsDesktop: menu.IsDesktop,
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
			landingTestimonialContentReviewRepository,
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
			landingTestimonialContent,
			landingFaqContent,
			landingMenus,
		)
		return err
	})

	return response, err
}
