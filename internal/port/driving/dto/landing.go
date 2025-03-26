package dto

import (
	"time"

	"github.com/guregu/null/v6"
	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
)

type LandingHeroContentRequest struct {
	IsEnabled   bool
	IsMobile    bool
	IsDesktop   bool
	Title       string
	Description string
	TagsLine    string
	ButtonLabel string
	ButtonUrl   string
	Image       null.Int64
}

type LandingSectionHeaderRequest struct {
	IsEnabled bool
	IsMobile  bool
	IsDesktop bool
	Title     string
	Subtitle  null.String
	TagsLine  null.String
}

type LandingPackageItemRequest struct {
	IsEnabled   bool
	IsMobile    bool
	IsDesktop   bool
	Package     int64
	ButtonLabel string
}

type LandingSinglePackageContentRequest struct {
	IsEnabled bool
	IsMobile  bool
	IsDesktop bool
	Header    LandingSectionHeaderRequest
	Silver    null.Value[LandingPackageItemRequest]
	Gold      null.Value[LandingPackageItemRequest]
	Platinum  null.Value[LandingPackageItemRequest]
}

type LandingPackageDetailRequest struct {
	IsEnabled bool
	IsMobile  bool
	IsDesktop bool
	Header    LandingSectionHeaderRequest
	Packages  []LandingPackageItemRequest
}

type LandingPackagesContentRequest struct {
	IsEnabled bool
	IsMobile  bool
	IsDesktop bool
	Silver    LandingPackageDetailRequest
	Gold      LandingPackageDetailRequest
	Platinum  LandingPackageDetailRequest
}

type LandingFeaturesContentBenefitRequest struct {
	IsEnabled bool
	IsMobile  bool
	IsDesktop bool
	Title     string
	Subtitle  string
	Logo      null.Int64
}

type LandingFeaturesContentRequest struct {
	IsEnabled     bool
	IsMobile      bool
	IsDesktop     bool
	Header        LandingSectionHeaderRequest
	Benefits      []LandingFeaturesContentBenefitRequest
	FooterTitle   string
	ButtonAbout   string
	ButtonPackage string
}

type LandingMomentsContentImageRequest struct {
	IsEnabled bool
	IsMobile  bool
	IsDesktop bool
	Image     int64
}

type LandingMomentsContentRequest struct {
	IsEnabled bool
	IsMobile  bool
	IsDesktop bool
	Header    LandingSectionHeaderRequest
	Images    []LandingMomentsContentImageRequest
}

type LandingAffiliatesContentAffiliateRequest struct {
	IsEnabled bool
	IsMobile  bool
	IsDesktop bool
	Name      string
	Logo      null.Int64
	Width     int
	Height    int
}

type LandingAffiliatesContentRequest struct {
	IsEnabled  bool
	IsMobile   bool
	IsDesktop  bool
	Header     LandingSectionHeaderRequest
	Affiliates []LandingAffiliatesContentAffiliateRequest
}

type LandingFaqContentFaqRequest struct {
	IsEnabled bool
	IsMobile  bool
	IsDesktop bool
	Question  string
	Answer    string
}

type LandingFaqContentRequest struct {
	IsEnabled bool
	IsMobile  bool
	IsDesktop bool
	Header    LandingSectionHeaderRequest
	Faqs      []LandingFaqContentFaqRequest
}

type LandingMenuRequest struct {
	IsEnabled bool
	IsMobile  bool
	IsDesktop bool
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
	IsMobile    bool
	IsDesktop   bool
	Title       string
	Description string
	TagsLine    string
	ButtonLabel string
	ButtonUrl   string
	Image       null.Value[ImageResponse]
}

type LandingSectionHeaderResponse struct {
	IsEnabled bool
	IsMobile  bool
	IsDesktop bool
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
	DoublePrice      float64
	DoubleFinalPrice null.Float
	TriplePrice      float64
	TripleFinalPrice null.Float
	QuadPrice        float64
	QuadFinalPrice   null.Float
	InfantPrice      null.Float
	InfantFinalPrice null.Float
}

type LandingPackageItemResponse struct {
	IsEnabled     bool
	IsMobile      bool
	IsDesktop     bool
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
	IsMobile  bool
	IsDesktop bool
	Header    LandingSectionHeaderResponse
	Silver    null.Value[LandingPackageItemResponse]
	Gold      null.Value[LandingPackageItemResponse]
	Platinum  null.Value[LandingPackageItemResponse]
}

type LandingPackageDetailResponse struct {
	IsEnabled bool
	IsMobile  bool
	IsDesktop bool
	Header    LandingSectionHeaderResponse
	Packages  []LandingPackageItemResponse
}

type LandingPackagesContentResponse struct {
	IsEnabled bool
	IsMobile  bool
	IsDesktop bool
	Silver    LandingPackageDetailResponse
	Gold      LandingPackageDetailResponse
	Platinum  LandingPackageDetailResponse
}

type LandingFeaturesContentBenefitResponse struct {
	IsEnabled bool
	IsMobile  bool
	IsDesktop bool
	Title     string
	Subtitle  string
	Logo      null.Value[ImageResponse]
}

type LandingFeaturesContentResponse struct {
	IsEnabled     bool
	IsMobile      bool
	IsDesktop     bool
	Header        LandingSectionHeaderResponse
	Benefits      []LandingFeaturesContentBenefitResponse
	FooterTitle   string
	ButtonAbout   string
	ButtonPackage string
}

type LandingMomentsContentImageResponse struct {
	IsEnabled bool
	IsMobile  bool
	IsDesktop bool
	Image     ImageResponse
}

type LandingMomentsContentResponse struct {
	IsEnabled bool
	IsMobile  bool
	IsDesktop bool
	Header    LandingSectionHeaderResponse
	Images    []LandingMomentsContentImageResponse
}

type LandingAffiliatesContentAffiliateResponse struct {
	IsEnabled bool
	IsMobile  bool
	IsDesktop bool
	Name      string
	Logo      null.Value[ImageResponse]
	Width     int
	Height    int
}

type LandingAffiliatesContentResponse struct {
	IsEnabled  bool
	IsMobile   bool
	IsDesktop  bool
	Header     LandingSectionHeaderResponse
	Affiliates []LandingAffiliatesContentAffiliateResponse
}

type LandingFaqContentFaqResponse struct {
	IsEnabled bool
	IsMobile  bool
	IsDesktop bool
	Question  string
	Answer    string
}

type LandingFaqContentResponse struct {
	IsEnabled bool
	IsMobile  bool
	IsDesktop bool
	Header    LandingSectionHeaderResponse
	Faqs      []LandingFaqContentFaqResponse
}

type LandingMenuResponse struct {
	IsEnabled bool
	IsMobile  bool
	IsDesktop bool
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
