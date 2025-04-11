package repository

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
)

type LandingHeroContentRepository interface {
	Create(ctx context.Context, landingHeroContent entity.LandingHeroContent) (entity.LandingHeroContent, error)
	Find(ctx context.Context) (entity.LandingHeroContent, error)
	Update(ctx context.Context, landingHeroContent entity.LandingHeroContent) (entity.LandingHeroContent, error)
	Delete(ctx context.Context) (entity.LandingHeroContent, error)
}

type LandingSectionHeaderRepository interface {
	Create(ctx context.Context, landingSectionHeader entity.LandingSectionHeader) (entity.LandingSectionHeader, error)
	FindById(ctx context.Context, id int64) (entity.LandingSectionHeader, error)
	FindAll(ctx context.Context, opt FindAllOptions) ([]entity.LandingSectionHeader, error)
	Update(ctx context.Context, id int64, landingSectionHeader entity.LandingSectionHeader) (entity.LandingSectionHeader, error)
	Delete(ctx context.Context, id int64) (entity.LandingSectionHeader, error)
}

type LandingPackageItemRepository interface {
	Create(ctx context.Context, landingPackageItem entity.LandingPackageItem) (entity.LandingPackageItem, error)
	FindById(ctx context.Context, id int64) (entity.LandingPackageItem, error)
	FindAll(ctx context.Context, opt FindAllOptions) ([]entity.LandingPackageItem, error)
	Update(ctx context.Context, id int64, landingPackageItem entity.LandingPackageItem) (entity.LandingPackageItem, error)
	Delete(ctx context.Context, id int64) (entity.LandingPackageItem, error)
}

type LandingSinglePackageContentRepository interface {
	Create(ctx context.Context, landingSinglePackageContent entity.LandingSinglePackageContent) (entity.LandingSinglePackageContent, error)
	Find(ctx context.Context) (entity.LandingSinglePackageContent, error)
	Update(ctx context.Context, landingSinglePackageContent entity.LandingSinglePackageContent) (entity.LandingSinglePackageContent, error)
	Delete(ctx context.Context) (entity.LandingSinglePackageContent, error)
}

type LandingPackageDetailRepository interface {
	Create(ctx context.Context, landingPackageDetail entity.LandingPackageDetail) (entity.LandingPackageDetail, error)
	FindById(ctx context.Context, id int64) (entity.LandingPackageDetail, error)
	FindAll(ctx context.Context, opt FindAllOptions) ([]entity.LandingPackageDetail, error)
	Update(ctx context.Context, id int64, landingPackageDetail entity.LandingPackageDetail) (entity.LandingPackageDetail, error)
	Delete(ctx context.Context, id int64) (entity.LandingPackageDetail, error)
}

type LandingPackageDetailItemRepository interface {
	CreateMany(ctx context.Context, landingPackageDetailItems []entity.LandingPackageDetailItem) ([]entity.LandingPackageDetailItem, error)
	FindAll(ctx context.Context, opt FindAllOptions) ([]entity.LandingPackageDetailItem, error)
	DeleteMany(ctx context.Context) ([]entity.LandingPackageDetailItem, error)
}

type LandingPackagesContentRepository interface {
	Create(ctx context.Context, landingPackagesContent entity.LandingPackagesContent) (entity.LandingPackagesContent, error)
	Find(ctx context.Context) (entity.LandingPackagesContent, error)
	Update(ctx context.Context, landingPackagesContent entity.LandingPackagesContent) (entity.LandingPackagesContent, error)
	Delete(ctx context.Context) (entity.LandingPackagesContent, error)
}

type LandingFeaturesContentRepository interface {
	Create(ctx context.Context, landingFeaturesContent entity.LandingFeaturesContent) (entity.LandingFeaturesContent, error)
	Find(ctx context.Context) (entity.LandingFeaturesContent, error)
	Update(ctx context.Context, landingFeaturesContent entity.LandingFeaturesContent) (entity.LandingFeaturesContent, error)
	Delete(ctx context.Context) (entity.LandingFeaturesContent, error)
}

type LandingFeaturesContentBenefitRepository interface {
	CreateMany(ctx context.Context, landingFeaturesContentBenefits []entity.LandingFeaturesContentBenefit) ([]entity.LandingFeaturesContentBenefit, error)
	FindAll(ctx context.Context, opt FindAllOptions) ([]entity.LandingFeaturesContentBenefit, error)
	DeleteMany(ctx context.Context) ([]entity.LandingFeaturesContentBenefit, error)
}

type LandingMomentsContentRepository interface {
	Create(ctx context.Context, landingMomentsContent entity.LandingMomentsContent) (entity.LandingMomentsContent, error)
	Find(ctx context.Context) (entity.LandingMomentsContent, error)
	Update(ctx context.Context, landingMomentsContent entity.LandingMomentsContent) (entity.LandingMomentsContent, error)
	Delete(ctx context.Context) (entity.LandingMomentsContent, error)
}

type LandingMomentsContentImageRepository interface {
	CreateMany(ctx context.Context, landingMomentsContentImages []entity.LandingMomentsContentImage) ([]entity.LandingMomentsContentImage, error)
	FindAll(ctx context.Context, opt FindAllOptions) ([]entity.LandingMomentsContentImage, error)
	DeleteMany(ctx context.Context) ([]entity.LandingMomentsContentImage, error)
}

type LandingAffiliatesContentRepository interface {
	Create(ctx context.Context, landingAffiliatesContent entity.LandingAffiliatesContent) (entity.LandingAffiliatesContent, error)
	Find(ctx context.Context) (entity.LandingAffiliatesContent, error)
	Update(ctx context.Context, landingAffiliatesContent entity.LandingAffiliatesContent) (entity.LandingAffiliatesContent, error)
	Delete(ctx context.Context) (entity.LandingAffiliatesContent, error)
}

type LandingAffiliatesContentAffiliateRepository interface {
	CreateMany(ctx context.Context, landingAffiliatesContentAffiliates []entity.LandingAffiliatesContentAffiliate) ([]entity.LandingAffiliatesContentAffiliate, error)
	FindAll(ctx context.Context, opt FindAllOptions) ([]entity.LandingAffiliatesContentAffiliate, error)
	DeleteMany(ctx context.Context) ([]entity.LandingAffiliatesContentAffiliate, error)
}

type LandingFaqContentRepository interface {
	Create(ctx context.Context, landingFaqContent entity.LandingFaqContent) (entity.LandingFaqContent, error)
	Find(ctx context.Context) (entity.LandingFaqContent, error)
	Update(ctx context.Context, landingFaqContent entity.LandingFaqContent) (entity.LandingFaqContent, error)
	Delete(ctx context.Context) (entity.LandingFaqContent, error)
}

type LandingFaqContentFaqRepository interface {
	CreateMany(ctx context.Context, landingFaqContentFaq []entity.LandingFaqContentFaq) ([]entity.LandingFaqContentFaq, error)
	FindAll(ctx context.Context, opt FindAllOptions) ([]entity.LandingFaqContentFaq, error)
	DeleteMany(ctx context.Context) ([]entity.LandingFaqContentFaq, error)
}

type LandingTestimonialContentRepository interface {
	Create(ctx context.Context, landingTestimonialContent entity.LandingTestimonialContent) (entity.LandingTestimonialContent, error)
	Find(ctx context.Context) (entity.LandingTestimonialContent, error)
	Update(ctx context.Context, landingTestimonialContent entity.LandingTestimonialContent) (entity.LandingTestimonialContent, error)
	Delete(ctx context.Context) (entity.LandingTestimonialContent, error)
}

type LandingTestimonialContentReviewRepository interface {
	CreateMany(ctx context.Context, landingTestimonialContentReviews []entity.LandingTestimonialContentReview) ([]entity.LandingTestimonialContentReview, error)
	FindAll(ctx context.Context, opt FindAllOptions) ([]entity.LandingTestimonialContentReview, error)
	DeleteMany(ctx context.Context) ([]entity.LandingTestimonialContentReview, error)
}

type LandingMenuRepository interface {
	CreateMany(ctx context.Context, landingMenus []entity.LandingMenu) ([]entity.LandingMenu, error)
	FindAll(ctx context.Context, opt FindAllOptions) ([]entity.LandingMenu, error)
	DeleteMany(ctx context.Context) ([]entity.LandingMenu, error)
}
