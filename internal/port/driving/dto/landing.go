package dto

import (
	"time"

	"github.com/guregu/null/v6"
	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
)

type LandingHeroContentRequest struct {
	IsEnabled   bool
	Title       string
	Description string
	TagsLine    string
	ButtonLabel string
	ButtonUrl   string
	Image       null.Int64
}

type LandingSectionHeaderRequest struct {
	IsEnabled bool
	Title     string
	Subtitle  null.String
	TagsLine  null.String
}

type LandingPackageItemRequest struct {
	IsEnabled   bool
	Package     int64
	ButtonLabel string
}

type LandingSinglePackageContentRequest struct {
	IsEnabled bool
	Header    LandingSectionHeaderRequest
	Silver    LandingPackageItemRequest
	Gold      LandingPackageItemRequest
	Platinum  LandingPackageItemRequest
}

type LandingPackageDetailRequest struct {
	IsEnabled bool
	Header    LandingSectionHeaderRequest
	Packages  []LandingPackageItemRequest
}

type LandingPackagesContentRequest struct {
	IsEnabled bool
	Silver    LandingPackageDetailRequest
	Gold      LandingPackageDetailRequest
	Platinum  LandingPackageDetailRequest
}

type LandingFeaturesContentBenefitRequest struct {
	IsEnabled bool
	Title     string
	Subtitle  string
	Logo      null.Int64
}

type LandingFeaturesContentRequest struct {
	IsEnabled     bool
	Header        LandingSectionHeaderRequest
	Benefits      []LandingFeaturesContentBenefitRequest
	FooterTitle   string
	ButtonAbout   string
	ButtonPackage string
}

type LandingMomentsContentImageRequest struct {
	IsEnabled bool
	Image     int64
}

type LandingMomentsContentRequest struct {
	IsEnabled bool
	Header    LandingSectionHeaderRequest
	Images    []LandingMomentsContentImageRequest
}

type LandingAffiliatesContentAffiliateRequest struct {
	IsEnabled bool
	Name      string
	Logo      null.Int64
	Width     int
	Height    int
}

type LandingAffiliatesContentRequest struct {
	IsEnabled  bool
	Header     LandingSectionHeaderRequest
	Affiliates []LandingAffiliatesContentAffiliateRequest
}

type LandingFaqContentFaqRequest struct {
	IsEnabled bool
	Question  string
	Answer    string
}

type LandingFaqContentRequest struct {
	IsEnabled bool
	Header    LandingSectionHeaderRequest
	Faqs      []LandingFaqContentFaqRequest
}

type LandingMenuRequest struct {
	IsEnabled bool
	Icon      string
	Label     string
	Path      string
}

type LandingRequest struct {
	HeroContent          LandingHeroContentRequest
	SinglePackageContent LandingSinglePackageContentRequest
	PackagesContent      LandingPackagesContentRequest
	FeaturesContent      LandingFeaturesContentRequest
	MomentsContent       LandingMomentsContentRequest
	AffiliatesContent    LandingAffiliatesContentRequest
	FaqContent           LandingFaqContentRequest
	Menus                []LandingMenuRequest
}

type LandingHeroContentResponse struct {
	IsEnabled   bool
	Title       string
	Description string
	TagsLine    string
	ButtonLabel string
	ButtonUrl   string
	Image       null.Value[ImageResponse]
}

type LandingSectionHeaderResponse struct {
	IsEnabled bool
	Title     string
	Subtitle  null.String
	TagsLine  null.String
}

type LandingPackageItemTagResponse struct {
	Icon  string
	Label string
}

type LandingPackageItemDepartureDateResponse struct {
	Date   time.Time
	Status string
}

type LandingPackageItemDetailResponse struct {
	Icon    string
	Label   string
	Value   string
	AltText string
	Rating  int
}

type LandingPackageItemPriceResponse struct {
	QuadPrice        float64
	TriplePrice      float64
	DoublePrice      float64
	InfantPrice      float64
	QuadFinalPrice   float64
	TripleFinalPrice float64
	DoubleFinalPrice float64
	InfantFinalPrice float64
}

type LandingPackageItemResponse struct {
	IsEnabled     bool
	Id            int64
	Thumbnail     null.Value[ImageResponse]
	Tags          []LandingPackageItemTagResponse
	Title         string
	DepartureDate []LandingPackageItemDepartureDateResponse
	Details       []LandingPackageItemDetailResponse
	Price         LandingPackageItemPriceResponse
	ButtonLabel   string
	Category      entity.PackageCategory
}

type LandingSinglePackageContentResponse struct {
	IsEnabled bool
	Header    LandingSectionHeaderResponse
	Silver    null.Value[LandingPackageItemResponse]
	Gold      null.Value[LandingPackageItemResponse]
	Platinum  null.Value[LandingPackageItemResponse]
}

type LandingPackageDetailResponse struct {
	IsEnabled bool
	Header    LandingSectionHeaderResponse
	Packages  []LandingPackageItemResponse
}

type LandingPackagesContentResponse struct {
	IsEnabled bool
	Silver    null.Value[LandingPackageDetailResponse]
	Gold      null.Value[LandingPackageDetailResponse]
	Platinum  null.Value[LandingPackageDetailResponse]
}

type LandingFeaturesContentBenefitResponse struct {
	IsEnabled bool
	Title     string
	Subtitle  string
	Logo      null.Value[ImageResponse]
}

type LandingFeaturesContentResponse struct {
	IsEnabled     bool
	Header        LandingSectionHeaderResponse
	Benefits      []LandingFeaturesContentBenefitResponse
	FooterTitle   string
	ButtonAbout   string
	ButtonPackage string
}

type LandingMomentsContentImageResponse struct {
	IsEnabled bool
	Image     ImageResponse
}

type LandingMomentsContentResponse struct {
	IsEnabled bool
	Header    LandingSectionHeaderResponse
	Images    []LandingMomentsContentImageResponse
}

type LandingAffiliatesContentAffiliateResponse struct {
	IsEnabled bool
	Name      string
	Logo      null.Value[ImageResponse]
	Width     int
	Height    int
}

type LandingAffiliatesContentResponse struct {
	IsEnabled  bool
	Header     LandingSectionHeaderResponse
	Affiliates []LandingAffiliatesContentAffiliateResponse
}

type LandingFaqContentFaqResponse struct {
	IsEnabled bool
	Question  string
	Answer    string
}

type LandingFaqContentResponse struct {
	IsEnabled bool
	Header    LandingSectionHeaderResponse
	Faqs      []LandingFaqContentFaqResponse
}

type LandingMenuResponse struct {
	IsEnabled bool
	Icon      string
	Label     string
	Path      string
}

type LandingResponse struct {
	HeroContent          LandingHeroContentResponse
	SinglePackageContent LandingSinglePackageContentResponse
	PackagesContent      LandingPackagesContentResponse
	FeaturesContent      LandingFeaturesContentResponse
	MomentsContent       LandingMomentsContentResponse
	AffiliatesContent    LandingAffiliatesContentResponse
	FaqContent           LandingFaqContentResponse
	Menus                []LandingMenuResponse
}
