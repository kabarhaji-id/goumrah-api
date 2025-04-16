package schema

import (
	"time"

	"github.com/guregu/null/v6"
	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type LandingHeroContentRequest struct {
	IsEnabled   bool       `json:"is_enabled"`
	IsMobile    bool       `json:"is_mobile"`
	IsDesktop   bool       `json:"is_desktop"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	TagsLine    string     `json:"tags_line"`
	ButtonLabel string     `json:"button_label"`
	ButtonUrl   string     `json:"button_url"`
	Image       null.Int64 `json:"image"`
}

func (r LandingHeroContentRequest) ToDtoRequest() dto.LandingHeroContentRequest {
	return dto.LandingHeroContentRequest{
		IsEnabled:   r.IsEnabled,
		IsMobile:    r.IsMobile,
		IsDesktop:   r.IsDesktop,
		Title:       r.Title,
		Description: r.Description,
		TagsLine:    r.TagsLine,
		ButtonLabel: r.ButtonLabel,
		ButtonUrl:   r.ButtonUrl,
		Image:       r.Image,
	}
}

type LandingSectionHeaderRequest struct {
	IsEnabled bool        `json:"is_enabled"`
	IsMobile  bool        `json:"is_mobile"`
	IsDesktop bool        `json:"is_desktop"`
	Title     string      `json:"title"`
	Subtitle  null.String `json:"subtitle"`
	TagsLine  null.String `json:"tags_line"`
}

func (r LandingSectionHeaderRequest) ToDtoRequest() dto.LandingSectionHeaderRequest {
	return dto.LandingSectionHeaderRequest{
		IsEnabled: r.IsEnabled,
		IsMobile:  r.IsMobile,
		IsDesktop: r.IsDesktop,
		Title:     r.Title,
		Subtitle:  r.Subtitle,
		TagsLine:  r.TagsLine,
	}
}

type LandingPackageItemRequest struct {
	IsEnabled   bool   `json:"is_enabled"`
	IsMobile    bool   `json:"is_mobile"`
	IsDesktop   bool   `json:"is_desktop"`
	Package     int64  `json:"package"`
	ButtonLabel string `json:"button_label"`
}

func (r LandingPackageItemRequest) ToDtoRequest() dto.LandingPackageItemRequest {
	return dto.LandingPackageItemRequest{
		IsEnabled:   r.IsEnabled,
		IsMobile:    r.IsMobile,
		IsDesktop:   r.IsDesktop,
		Package:     r.Package,
		ButtonLabel: r.ButtonLabel,
	}
}

type LandingSinglePackageContentRequest struct {
	IsEnabled bool                                  `json:"is_enabled"`
	IsMobile  bool                                  `json:"is_mobile"`
	IsDesktop bool                                  `json:"is_desktop"`
	Header    LandingSectionHeaderRequest           `json:"header"`
	Silver    null.Value[LandingPackageItemRequest] `json:"silver"`
	Gold      null.Value[LandingPackageItemRequest] `json:"gold"`
	Platinum  null.Value[LandingPackageItemRequest] `json:"platinum"`
}

func (r LandingSinglePackageContentRequest) ToDtoRequest() dto.LandingSinglePackageContentRequest {
	silver := null.NewValue(dto.LandingPackageItemRequest{}, false)
	if r.Silver.Valid {
		silver = null.ValueFrom(r.Silver.V.ToDtoRequest())
	}

	gold := null.NewValue(dto.LandingPackageItemRequest{}, false)
	if r.Gold.Valid {
		gold = null.ValueFrom(r.Gold.V.ToDtoRequest())
	}

	platinum := null.NewValue(dto.LandingPackageItemRequest{}, false)
	if r.Platinum.Valid {
		platinum = null.ValueFrom(r.Platinum.V.ToDtoRequest())
	}

	return dto.LandingSinglePackageContentRequest{
		IsEnabled: r.IsEnabled,
		IsMobile:  r.IsMobile,
		IsDesktop: r.IsDesktop,
		Header:    r.Header.ToDtoRequest(),
		Silver:    silver,
		Gold:      gold,
		Platinum:  platinum,
	}
}

type LandingPackageDetailRequest struct {
	IsEnabled bool                        `json:"is_enabled"`
	IsMobile  bool                        `json:"is_mobile"`
	IsDesktop bool                        `json:"is_desktop"`
	Header    LandingSectionHeaderRequest `json:"header"`
	Packages  []LandingPackageItemRequest `json:"packages"`
}

func (r LandingPackageDetailRequest) ToDtoRequest() dto.LandingPackageDetailRequest {
	packages := make([]dto.LandingPackageItemRequest, len(r.Packages))
	for i, p := range r.Packages {
		packages[i] = p.ToDtoRequest()
	}
	return dto.LandingPackageDetailRequest{
		IsEnabled: r.IsEnabled,
		IsMobile:  r.IsMobile,
		IsDesktop: r.IsDesktop,
		Header:    r.Header.ToDtoRequest(),
		Packages:  packages,
	}
}

type LandingPackagesContentRequest struct {
	IsEnabled bool                        `json:"is_enabled"`
	IsMobile  bool                        `json:"is_mobile"`
	IsDesktop bool                        `json:"is_desktop"`
	Silver    LandingPackageDetailRequest `json:"silver"`
	Gold      LandingPackageDetailRequest `json:"gold"`
	Platinum  LandingPackageDetailRequest `json:"platinum"`
}

func (r LandingPackagesContentRequest) ToDtoRequest() dto.LandingPackagesContentRequest {
	return dto.LandingPackagesContentRequest{
		IsEnabled: r.IsEnabled,
		IsMobile:  r.IsMobile,
		IsDesktop: r.IsDesktop,
		Silver:    r.Silver.ToDtoRequest(),
		Gold:      r.Gold.ToDtoRequest(),
		Platinum:  r.Platinum.ToDtoRequest(),
	}
}

type LandingTravelDestinationContentDestinationRequest struct {
	IsEnabled bool       `json:"is_enabled"`
	IsMobile  bool       `json:"is_mobile"`
	IsDesktop bool       `json:"is_desktop"`
	Image     null.Int64 `json:"image"`
	Name      string     `json:"name"`
}

func (r LandingTravelDestinationContentDestinationRequest) ToDtoRequest() dto.LandingTravelDestinationContentDestinationRequest {
	return dto.LandingTravelDestinationContentDestinationRequest{
		IsEnabled: r.IsEnabled,
		IsMobile:  r.IsMobile,
		IsDesktop: r.IsDesktop,
		Image:     r.Image,
		Name:      r.Name,
	}
}

type LandingTravelDestinationContentRequest struct {
	IsEnabled    bool                                                `json:"is_enabled"`
	IsMobile     bool                                                `json:"is_mobile"`
	IsDesktop    bool                                                `json:"is_desktop"`
	Header       LandingSectionHeaderRequest                         `json:"header"`
	Destinations []LandingTravelDestinationContentDestinationRequest `json:"destinations"`
}

func (r LandingTravelDestinationContentRequest) ToDtoRequest() dto.LandingTravelDestinationContentRequest {
	destinations := make([]dto.LandingTravelDestinationContentDestinationRequest, len(r.Destinations))
	for i, d := range r.Destinations {
		destinations[i] = d.ToDtoRequest()
	}

	return dto.LandingTravelDestinationContentRequest{
		IsEnabled:    r.IsEnabled,
		IsMobile:     r.IsMobile,
		IsDesktop:    r.IsDesktop,
		Header:       r.Header.ToDtoRequest(),
		Destinations: destinations,
	}
}

type LandingFeaturesContentBenefitRequest struct {
	IsEnabled bool       `json:"is_enabled"`
	IsMobile  bool       `json:"is_mobile"`
	IsDesktop bool       `json:"is_desktop"`
	Title     string     `json:"title"`
	Subtitle  string     `json:"subtitle"`
	Logo      null.Int64 `json:"logo"`
}

func (r LandingFeaturesContentBenefitRequest) ToDtoRequest() dto.LandingFeaturesContentBenefitRequest {
	return dto.LandingFeaturesContentBenefitRequest{
		IsEnabled: r.IsEnabled,
		IsMobile:  r.IsMobile,
		IsDesktop: r.IsDesktop,
		Title:     r.Title,
		Subtitle:  r.Subtitle,
		Logo:      r.Logo,
	}
}

type LandingFeaturesContentRequest struct {
	IsEnabled     bool                                   `json:"is_enabled"`
	IsMobile      bool                                   `json:"is_mobile"`
	IsDesktop     bool                                   `json:"is_desktop"`
	Header        LandingSectionHeaderRequest            `json:"header"`
	Benefits      []LandingFeaturesContentBenefitRequest `json:"benefits"`
	FooterTitle   string                                 `json:"footer_title"`
	ButtonAbout   string                                 `json:"button_about"`
	ButtonPackage string                                 `json:"button_package"`
}

func (r LandingFeaturesContentRequest) ToDtoRequest() dto.LandingFeaturesContentRequest {
	benefits := make([]dto.LandingFeaturesContentBenefitRequest, len(r.Benefits))
	for i, b := range r.Benefits {
		benefits[i] = b.ToDtoRequest()
	}
	return dto.LandingFeaturesContentRequest{
		IsEnabled:     r.IsEnabled,
		IsMobile:      r.IsMobile,
		IsDesktop:     r.IsDesktop,
		Header:        r.Header.ToDtoRequest(),
		Benefits:      benefits,
		FooterTitle:   r.FooterTitle,
		ButtonAbout:   r.ButtonAbout,
		ButtonPackage: r.ButtonPackage,
	}
}

type LandingMomentsContentImageRequest struct {
	IsEnabled bool  `json:"is_enabled"`
	IsMobile  bool  `json:"is_mobile"`
	IsDesktop bool  `json:"is_desktop"`
	Image     int64 `json:"image"`
}

func (r LandingMomentsContentImageRequest) ToDtoRequest() dto.LandingMomentsContentImageRequest {
	return dto.LandingMomentsContentImageRequest{
		IsEnabled: r.IsEnabled,
		IsMobile:  r.IsMobile,
		IsDesktop: r.IsDesktop,
		Image:     r.Image,
	}
}

type LandingMomentsContentRequest struct {
	IsEnabled bool                                `json:"is_enabled"`
	IsMobile  bool                                `json:"is_mobile"`
	IsDesktop bool                                `json:"is_desktop"`
	Header    LandingSectionHeaderRequest         `json:"header"`
	Images    []LandingMomentsContentImageRequest `json:"images"`
}

func (r LandingMomentsContentRequest) ToDtoRequest() dto.LandingMomentsContentRequest {
	images := make([]dto.LandingMomentsContentImageRequest, len(r.Images))
	for i, img := range r.Images {
		images[i] = img.ToDtoRequest()
	}
	return dto.LandingMomentsContentRequest{
		IsEnabled: r.IsEnabled,
		IsMobile:  r.IsMobile,
		IsDesktop: r.IsDesktop,
		Header:    r.Header.ToDtoRequest(),
		Images:    images,
	}
}

type LandingAffiliatesContentAffiliateRequest struct {
	IsEnabled bool       `json:"is_enabled"`
	IsMobile  bool       `json:"is_mobile"`
	IsDesktop bool       `json:"is_desktop"`
	Name      string     `json:"name"`
	Logo      null.Int64 `json:"logo"`
	Width     int        `json:"width"`
	Height    int        `json:"height"`
}

func (r LandingAffiliatesContentAffiliateRequest) ToDtoRequest() dto.LandingAffiliatesContentAffiliateRequest {
	return dto.LandingAffiliatesContentAffiliateRequest{
		IsEnabled: r.IsEnabled,
		IsMobile:  r.IsMobile,
		IsDesktop: r.IsDesktop,
		Name:      r.Name,
		Logo:      r.Logo,
		Width:     r.Width,
		Height:    r.Height,
	}
}

type LandingAffiliatesContentRequest struct {
	IsEnabled  bool                                       `json:"is_enabled"`
	IsMobile   bool                                       `json:"is_mobile"`
	IsDesktop  bool                                       `json:"is_desktop"`
	Header     LandingSectionHeaderRequest                `json:"header"`
	Affiliates []LandingAffiliatesContentAffiliateRequest `json:"affiliates"`
}

func (r LandingAffiliatesContentRequest) ToDtoRequest() dto.LandingAffiliatesContentRequest {
	affiliates := make([]dto.LandingAffiliatesContentAffiliateRequest, len(r.Affiliates))
	for i, a := range r.Affiliates {
		affiliates[i] = a.ToDtoRequest()
	}
	return dto.LandingAffiliatesContentRequest{
		IsEnabled:  r.IsEnabled,
		IsMobile:   r.IsMobile,
		IsDesktop:  r.IsDesktop,
		Header:     r.Header.ToDtoRequest(),
		Affiliates: affiliates,
	}
}

type LandingFaqContentFaqRequest struct {
	IsEnabled bool   `json:"is_enabled"`
	IsMobile  bool   `json:"is_mobile"`
	IsDesktop bool   `json:"is_desktop"`
	Question  string `json:"question"`
	Answer    string `json:"answer"`
}

func (r LandingFaqContentFaqRequest) ToDtoRequest() dto.LandingFaqContentFaqRequest {
	return dto.LandingFaqContentFaqRequest{
		IsEnabled: r.IsEnabled,
		IsMobile:  r.IsMobile,
		IsDesktop: r.IsDesktop,
		Question:  r.Question,
		Answer:    r.Answer,
	}
}

type LandingFaqContentRequest struct {
	IsEnabled bool                          `json:"is_enabled"`
	IsMobile  bool                          `json:"is_mobile"`
	IsDesktop bool                          `json:"is_desktop"`
	Header    LandingSectionHeaderRequest   `json:"header"`
	Faqs      []LandingFaqContentFaqRequest `json:"faqs"`
}

func (r LandingFaqContentRequest) ToDtoRequest() dto.LandingFaqContentRequest {
	faqs := make([]dto.LandingFaqContentFaqRequest, len(r.Faqs))
	for i, faq := range r.Faqs {
		faqs[i] = faq.ToDtoRequest()
	}
	return dto.LandingFaqContentRequest{
		IsEnabled: r.IsEnabled,
		IsMobile:  r.IsMobile,
		IsDesktop: r.IsDesktop,
		Header:    r.Header.ToDtoRequest(),
		Faqs:      faqs,
	}
}

type LandingTestimonialContentReviewRequest struct {
	IsEnabled bool    `json:"is_enabled"`
	IsMobile  bool    `json:"is_mobile"`
	IsDesktop bool    `json:"is_desktop"`
	Reviewer  string  `json:"reviewer"`
	Age       int     `json:"age"`
	Address   string  `json:"address"`
	Rating    float32 `json:"rating"`
	Review    string  `json:"review"`
}

func (r LandingTestimonialContentReviewRequest) ToDtoRequest() dto.LandingTestimonialContentReviewRequest {
	return dto.LandingTestimonialContentReviewRequest{
		IsEnabled: r.IsEnabled,
		IsMobile:  r.IsMobile,
		IsDesktop: r.IsDesktop,
		Reviewer:  r.Reviewer,
		Age:       r.Age,
		Address:   r.Address,
		Rating:    r.Rating,
		Review:    r.Review,
	}
}

type LandingTestimonialContentRequest struct {
	IsEnabled bool                                     `json:"is_enabled"`
	IsMobile  bool                                     `json:"is_mobile"`
	IsDesktop bool                                     `json:"is_desktop"`
	Header    LandingSectionHeaderRequest              `json:"header"`
	Reviews   []LandingTestimonialContentReviewRequest `json:"reviews"`
}

func (r LandingTestimonialContentRequest) ToDtoRequest() dto.LandingTestimonialContentRequest {
	reviews := make([]dto.LandingTestimonialContentReviewRequest, len(r.Reviews))
	for i, review := range r.Reviews {
		reviews[i] = review.ToDtoRequest()
	}

	return dto.LandingTestimonialContentRequest{
		IsEnabled: r.IsEnabled,
		IsMobile:  r.IsMobile,
		IsDesktop: r.IsDesktop,
		Header:    r.Header.ToDtoRequest(),
		Reviews:   reviews,
	}
}

type LandingMenuRequest struct {
	IsEnabled bool   `json:"is_enabled"`
	IsMobile  bool   `json:"is_mobile"`
	IsDesktop bool   `json:"is_desktop"`
	Icon      string `json:"icon"`
	Label     string `json:"label"`
	Path      string `json:"path"`
}

func (r LandingMenuRequest) ToDtoRequest() dto.LandingMenuRequest {
	return dto.LandingMenuRequest{
		IsEnabled: r.IsEnabled,
		IsMobile:  r.IsMobile,
		IsDesktop: r.IsDesktop,
		Icon:      r.Icon,
		Label:     r.Label,
		Path:      r.Path,
	}
}

type LandingRequest struct {
	HeroContent              LandingHeroContentRequest              `json:"hero_content"`
	SinglePackageContent     LandingSinglePackageContentRequest     `json:"single_package_content"`
	PackagesContent          LandingPackagesContentRequest          `json:"packages_content"`
	TravelDestinationContent LandingTravelDestinationContentRequest `json:"travel_destination_content"`
	FeaturesContent          LandingFeaturesContentRequest          `json:"features_content"`
	MomentsContent           LandingMomentsContentRequest           `json:"moments_content"`
	AffiliatesContent        LandingAffiliatesContentRequest        `json:"affiliates_content"`
	TestimonialContent       LandingTestimonialContentRequest       `json:"testimonial_content"`
	FaqContent               LandingFaqContentRequest               `json:"faq_content"`
	Menus                    []LandingMenuRequest                   `json:"menus"`
}

func (r LandingRequest) ToDtoRequest() dto.LandingRequest {
	menus := make([]dto.LandingMenuRequest, len(r.Menus))
	for i, m := range r.Menus {
		menus[i] = m.ToDtoRequest()
	}

	return dto.LandingRequest{
		HeroContent:              r.HeroContent.ToDtoRequest(),
		SinglePackageContent:     r.SinglePackageContent.ToDtoRequest(),
		PackagesContent:          r.PackagesContent.ToDtoRequest(),
		TravelDestinationContent: r.TravelDestinationContent.ToDtoRequest(),
		FeaturesContent:          r.FeaturesContent.ToDtoRequest(),
		MomentsContent:           r.MomentsContent.ToDtoRequest(),
		AffiliatesContent:        r.AffiliatesContent.ToDtoRequest(),
		TestimonialContent:       r.TestimonialContent.ToDtoRequest(),
		FaqContent:               r.FaqContent.ToDtoRequest(),
		Menus:                    menus,
	}
}

type LandingHeroContentResponse struct {
	IsEnabled   bool                      `json:"is_enabled"`
	IsMobile    bool                      `json:"is_mobile"`
	IsDesktop   bool                      `json:"is_desktop"`
	Title       string                    `json:"title"`
	Description string                    `json:"description"`
	TagsLine    string                    `json:"tags_line"`
	ButtonLabel string                    `json:"button_label"`
	ButtonUrl   string                    `json:"button_url"`
	Image       null.Value[ImageResponse] `json:"image"`
}

func NewLandingHeroContentResponse(dto dto.LandingHeroContentResponse) LandingHeroContentResponse {
	image := null.NewValue(ImageResponse{}, false)
	if dto.Image.Valid {
		imageResponse := NewImageResponse(dto.Image.V)

		image = null.ValueFrom(imageResponse)
	}

	return LandingHeroContentResponse{
		IsEnabled:   dto.IsEnabled,
		IsMobile:    dto.IsMobile,
		IsDesktop:   dto.IsDesktop,
		Title:       dto.Title,
		Description: dto.Description,
		TagsLine:    dto.TagsLine,
		ButtonLabel: dto.ButtonLabel,
		ButtonUrl:   dto.ButtonUrl,
		Image:       image,
	}
}

func NewLandingHeroContentResponses(dtos []dto.LandingHeroContentResponse) []LandingHeroContentResponse {
	responses := make([]LandingHeroContentResponse, len(dtos))
	for i, dto := range dtos {
		responses[i] = NewLandingHeroContentResponse(dto)
	}
	return responses
}

type LandingSectionHeaderResponse struct {
	IsEnabled bool        `json:"is_enabled"`
	IsMobile  bool        `json:"is_mobile"`
	IsDesktop bool        `json:"is_desktop"`
	Title     string      `json:"title"`
	Subtitle  null.String `json:"subtitle"`
	TagsLine  null.String `json:"tags_line"`
}

func NewLandingSectionHeaderResponse(dto dto.LandingSectionHeaderResponse) LandingSectionHeaderResponse {
	return LandingSectionHeaderResponse{
		IsEnabled: dto.IsEnabled,
		IsMobile:  dto.IsMobile,
		IsDesktop: dto.IsDesktop,
		Title:     dto.Title,
		Subtitle:  dto.Subtitle,
		TagsLine:  dto.TagsLine,
	}
}

func NewLandingSectionHeaderResponses(dtos []dto.LandingSectionHeaderResponse) []LandingSectionHeaderResponse {
	responses := make([]LandingSectionHeaderResponse, len(dtos))
	for i, dto := range dtos {
		responses[i] = NewLandingSectionHeaderResponse(dto)
	}
	return responses
}

type LandingPackageItemTagResponse struct {
	Icon  string `json:"icon"`
	Label string `json:"label"`
}

func NewLandingPackageItemTagResponse(dto dto.LandingPackageItemTagResponse) LandingPackageItemTagResponse {
	return LandingPackageItemTagResponse{
		Icon:  dto.Icon,
		Label: dto.Label,
	}
}

func NewLandingPackageItemTagResponses(dtos []dto.LandingPackageItemTagResponse) []LandingPackageItemTagResponse {
	responses := make([]LandingPackageItemTagResponse, len(dtos))
	for i, dto := range dtos {
		responses[i] = NewLandingPackageItemTagResponse(dto)
	}
	return responses
}

type LandingPackageItemDepartureDateResponse struct {
	Date   time.Time `json:"date"`
	Status string    `json:"status"`
}

func NewLandingPackageItemDepartureDateResponse(dto dto.LandingPackageItemDepartureDateResponse) LandingPackageItemDepartureDateResponse {
	return LandingPackageItemDepartureDateResponse{
		Date:   dto.Date,
		Status: dto.Status,
	}
}

func NewLandingPackageItemDepartureDateResponses(dtos []dto.LandingPackageItemDepartureDateResponse) []LandingPackageItemDepartureDateResponse {
	responses := make([]LandingPackageItemDepartureDateResponse, len(dtos))
	for i, dto := range dtos {
		responses[i] = NewLandingPackageItemDepartureDateResponse(dto)
	}
	return responses
}

type LandingPackageItemDetailResponse struct {
	Icon    string `json:"icon"`
	Label   string `json:"label"`
	Value   string `json:"value"`
	AltText string `json:"alt_text"`
	Rating  int    `json:"rating"`
}

func NewLandingPackageItemDetailResponse(dto dto.LandingPackageItemDetailResponse) LandingPackageItemDetailResponse {
	return LandingPackageItemDetailResponse{
		Icon:    dto.Icon,
		Label:   dto.Label,
		Value:   dto.Value,
		AltText: dto.AltText,
		Rating:  dto.Rating,
	}
}

func NewLandingPackageItemDetailResponses(dtos []dto.LandingPackageItemDetailResponse) []LandingPackageItemDetailResponse {
	responses := make([]LandingPackageItemDetailResponse, len(dtos))
	for i, dto := range dtos {
		responses[i] = NewLandingPackageItemDetailResponse(dto)
	}
	return responses
}

type LandingPackageItemPriceResponse struct {
	DoublePrice      float64    `json:"double_price"`
	DoubleFinalPrice null.Float `json:"double_final_price"`
	TriplePrice      float64    `json:"triple_price"`
	TripleFinalPrice null.Float `json:"triple_final_price"`
	QuadPrice        float64    `json:"quad_price"`
	QuadFinalPrice   null.Float `json:"quad_final_price"`
	InfantPrice      null.Float `json:"infant_price"`
	InfantFinalPrice null.Float `json:"infant_final_price"`
}

func NewLandingPackageItemPriceResponse(dto dto.LandingPackageItemPriceResponse) LandingPackageItemPriceResponse {
	return LandingPackageItemPriceResponse{
		DoublePrice:      dto.DoublePrice,
		DoubleFinalPrice: dto.DoubleFinalPrice,
		TriplePrice:      dto.TriplePrice,
		TripleFinalPrice: dto.TripleFinalPrice,
		QuadPrice:        dto.QuadPrice,
		QuadFinalPrice:   dto.QuadFinalPrice,
		InfantPrice:      dto.InfantPrice,
		InfantFinalPrice: dto.InfantFinalPrice,
	}
}

func NewLandingPackageItemPriceResponses(dtos []dto.LandingPackageItemPriceResponse) []LandingPackageItemPriceResponse {
	responses := make([]LandingPackageItemPriceResponse, len(dtos))
	for i, dto := range dtos {
		responses[i] = NewLandingPackageItemPriceResponse(dto)
	}
	return responses
}

type LandingPackageItemResponse struct {
	IsEnabled     bool                                      `json:"is_enabled"`
	IsMobile      bool                                      `json:"is_mobile"`
	IsDesktop     bool                                      `json:"is_desktop"`
	Id            int64                                     `json:"id"`
	Thumbnail     null.Value[ImageResponse]                 `json:"thumbnail"`
	Tags          []LandingPackageItemTagResponse           `json:"tags"`
	Title         string                                    `json:"title"`
	DepartureDate []LandingPackageItemDepartureDateResponse `json:"departure_date"`
	Details       []LandingPackageItemDetailResponse        `json:"details"`
	Price         LandingPackageItemPriceResponse           `json:"price"`
	ButtonLabel   string                                    `json:"button_label"`
	Category      entity.PackageCategory                    `json:"category"`
}

func NewLandingPackageItemResponse(dto dto.LandingPackageItemResponse) LandingPackageItemResponse {
	thumbnail := null.NewValue(ImageResponse{}, false)
	if dto.Thumbnail.Valid {
		imageResponse := NewImageResponse(dto.Thumbnail.V)

		thumbnail = null.ValueFrom(imageResponse)
	}

	tags := make([]LandingPackageItemTagResponse, len(dto.Tags))
	for i, tag := range dto.Tags {
		tags[i] = NewLandingPackageItemTagResponse(tag)
	}

	departureDates := make([]LandingPackageItemDepartureDateResponse, len(dto.DepartureDate))
	for i, date := range dto.DepartureDate {
		departureDates[i] = NewLandingPackageItemDepartureDateResponse(date)
	}

	details := make([]LandingPackageItemDetailResponse, len(dto.Details))
	for i, detail := range dto.Details {
		details[i] = NewLandingPackageItemDetailResponse(detail)
	}

	return LandingPackageItemResponse{
		IsEnabled:     dto.IsEnabled,
		IsMobile:      dto.IsMobile,
		IsDesktop:     dto.IsDesktop,
		Id:            dto.Id,
		Thumbnail:     thumbnail,
		Tags:          tags,
		Title:         dto.Title,
		DepartureDate: departureDates,
		Details:       details,
		Price:         NewLandingPackageItemPriceResponse(dto.Price),
		ButtonLabel:   dto.ButtonLabel,
		Category:      dto.Category,
	}
}

func NewLandingPackageItemResponses(dtos []dto.LandingPackageItemResponse) []LandingPackageItemResponse {
	responses := make([]LandingPackageItemResponse, len(dtos))
	for i, dto := range dtos {
		responses[i] = NewLandingPackageItemResponse(dto)
	}
	return responses
}

type LandingSinglePackageContentResponse struct {
	IsEnabled bool                                   `json:"is_enabled"`
	IsMobile  bool                                   `json:"is_mobile"`
	IsDesktop bool                                   `json:"is_desktop"`
	Header    LandingSectionHeaderResponse           `json:"header"`
	Silver    null.Value[LandingPackageItemResponse] `json:"silver"`
	Gold      null.Value[LandingPackageItemResponse] `json:"gold"`
	Platinum  null.Value[LandingPackageItemResponse] `json:"platinum"`
}

func NewLandingSinglePackageContentResponse(dto dto.LandingSinglePackageContentResponse) LandingSinglePackageContentResponse {
	silver := null.NewValue(LandingPackageItemResponse{}, false)
	if dto.Silver.Valid {
		silver = null.ValueFrom(NewLandingPackageItemResponse(dto.Silver.V))
	}

	gold := null.NewValue(LandingPackageItemResponse{}, false)
	if dto.Gold.Valid {
		gold = null.ValueFrom(NewLandingPackageItemResponse(dto.Gold.V))
	}

	platinum := null.NewValue(LandingPackageItemResponse{}, false)
	if dto.Platinum.Valid {
		platinum = null.ValueFrom(NewLandingPackageItemResponse(dto.Platinum.V))
	}

	return LandingSinglePackageContentResponse{
		IsEnabled: dto.IsEnabled,
		IsMobile:  dto.IsMobile,
		IsDesktop: dto.IsDesktop,
		Header:    NewLandingSectionHeaderResponse(dto.Header),
		Silver:    silver,
		Gold:      gold,
		Platinum:  platinum,
	}
}

func NewLandingSinglePackageContentResponses(dtos []dto.LandingSinglePackageContentResponse) []LandingSinglePackageContentResponse {
	responses := make([]LandingSinglePackageContentResponse, len(dtos))
	for i, dto := range dtos {
		responses[i] = NewLandingSinglePackageContentResponse(dto)
	}
	return responses
}

type LandingPackageDetailResponse struct {
	IsEnabled bool                         `json:"is_enabled"`
	IsMobile  bool                         `json:"is_mobile"`
	IsDesktop bool                         `json:"is_desktop"`
	Header    LandingSectionHeaderResponse `json:"header"`
	Packages  []LandingPackageItemResponse `json:"packages"`
}

func NewLandingPackageDetailResponse(dto dto.LandingPackageDetailResponse) LandingPackageDetailResponse {
	packages := make([]LandingPackageItemResponse, len(dto.Packages))
	for i, p := range dto.Packages {
		packages[i] = NewLandingPackageItemResponse(p)
	}
	return LandingPackageDetailResponse{
		IsEnabled: dto.IsEnabled,
		IsMobile:  dto.IsMobile,
		IsDesktop: dto.IsDesktop,
		Header:    NewLandingSectionHeaderResponse(dto.Header),
		Packages:  packages,
	}
}

func NewLandingPackageDetailResponses(dtos []dto.LandingPackageDetailResponse) []LandingPackageDetailResponse {
	responses := make([]LandingPackageDetailResponse, len(dtos))
	for i, dto := range dtos {
		responses[i] = NewLandingPackageDetailResponse(dto)
	}
	return responses
}

type LandingPackagesContentResponse struct {
	IsEnabled bool                         `json:"is_enabled"`
	IsMobile  bool                         `json:"is_mobile"`
	IsDesktop bool                         `json:"is_desktop"`
	Silver    LandingPackageDetailResponse `json:"silver"`
	Gold      LandingPackageDetailResponse `json:"gold"`
	Platinum  LandingPackageDetailResponse `json:"platinum"`
}

func NewLandingPackagesContentResponse(dto dto.LandingPackagesContentResponse) LandingPackagesContentResponse {
	return LandingPackagesContentResponse{
		IsEnabled: dto.IsEnabled,
		IsMobile:  dto.IsMobile,
		IsDesktop: dto.IsDesktop,
		Silver:    NewLandingPackageDetailResponse(dto.Silver),
		Gold:      NewLandingPackageDetailResponse(dto.Gold),
		Platinum:  NewLandingPackageDetailResponse(dto.Platinum),
	}
}

func NewLandingPackagesContentResponses(dtos []dto.LandingPackagesContentResponse) []LandingPackagesContentResponse {
	responses := make([]LandingPackagesContentResponse, len(dtos))
	for i, dto := range dtos {
		responses[i] = NewLandingPackagesContentResponse(dto)
	}
	return responses
}

type LandingTravelDestinationContentDestinationResponse struct {
	IsEnabled bool                      `json:"is_enabled"`
	IsMobile  bool                      `json:"is_mobile"`
	IsDesktop bool                      `json:"is_desktop"`
	Image     null.Value[ImageResponse] `json:"image"`
	Name      string                    `json:"name"`
}

func NewLandingTravelDestinationContentDestinationResponse(dto dto.LandingTravelDestinationContentDestinationResponse) LandingTravelDestinationContentDestinationResponse {
	image := null.NewValue(ImageResponse{}, false)
	if dto.Image.Valid {
		imageResponse := NewImageResponse(dto.Image.V)

		image = null.ValueFrom(imageResponse)
	}

	return LandingTravelDestinationContentDestinationResponse{
		IsEnabled: dto.IsEnabled,
		IsMobile:  dto.IsMobile,
		IsDesktop: dto.IsDesktop,
		Image:     image,
		Name:      dto.Name,
	}
}

func NewLandingTravelDestinationContentDestinationResponses(dtos []dto.LandingTravelDestinationContentDestinationResponse) []LandingTravelDestinationContentDestinationResponse {
	responses := make([]LandingTravelDestinationContentDestinationResponse, len(dtos))
	for i, dto := range dtos {
		responses[i] = NewLandingTravelDestinationContentDestinationResponse(dto)
	}
	return responses
}

type LandingTravelDestinationContentResponse struct {
	IsEnabled    bool                                                 `json:"is_enabled"`
	IsMobile     bool                                                 `json:"is_mobile"`
	IsDesktop    bool                                                 `json:"is_desktop"`
	Header       LandingSectionHeaderResponse                         `json:"header"`
	Destinations []LandingTravelDestinationContentDestinationResponse `json:"destinations"`
}

func NewLandingTravelDestinationContentResponse(dto dto.LandingTravelDestinationContentResponse) LandingTravelDestinationContentResponse {
	destinations := make([]LandingTravelDestinationContentDestinationResponse, len(dto.Destinations))
	for i, d := range dto.Destinations {
		destinations[i] = NewLandingTravelDestinationContentDestinationResponse(d)
	}
	return LandingTravelDestinationContentResponse{
		IsEnabled:    dto.IsEnabled,
		IsMobile:     dto.IsMobile,
		IsDesktop:    dto.IsDesktop,
		Header:       NewLandingSectionHeaderResponse(dto.Header),
		Destinations: destinations,
	}
}

func NewLandingTravelDestinationContentResponses(dtos []dto.LandingTravelDestinationContentResponse) []LandingTravelDestinationContentResponse {
	responses := make([]LandingTravelDestinationContentResponse, len(dtos))
	for i, dto := range dtos {
		responses[i] = NewLandingTravelDestinationContentResponse(dto)
	}
	return responses
}

type LandingFeaturesContentBenefitResponse struct {
	IsEnabled bool                      `json:"is_enabled"`
	IsMobile  bool                      `json:"is_mobile"`
	IsDesktop bool                      `json:"is_desktop"`
	Title     string                    `json:"title"`
	Subtitle  string                    `json:"subtitle"`
	Logo      null.Value[ImageResponse] `json:"logo"`
}

func NewLandingFeaturesContentBenefitResponse(dto dto.LandingFeaturesContentBenefitResponse) LandingFeaturesContentBenefitResponse {
	logo := null.NewValue(ImageResponse{}, false)
	if dto.Logo.Valid {
		imageResponse := NewImageResponse(dto.Logo.V)

		logo = null.ValueFrom(imageResponse)
	}

	return LandingFeaturesContentBenefitResponse{
		IsEnabled: dto.IsEnabled,
		IsMobile:  dto.IsMobile,
		IsDesktop: dto.IsDesktop,
		Title:     dto.Title,
		Subtitle:  dto.Subtitle,
		Logo:      logo,
	}
}

func NewLandingFeaturesContentBenefitResponses(dtos []dto.LandingFeaturesContentBenefitResponse) []LandingFeaturesContentBenefitResponse {
	responses := make([]LandingFeaturesContentBenefitResponse, len(dtos))
	for i, dto := range dtos {
		responses[i] = NewLandingFeaturesContentBenefitResponse(dto)
	}
	return responses
}

type LandingFeaturesContentResponse struct {
	IsEnabled     bool                                    `json:"is_enabled"`
	IsMobile      bool                                    `json:"is_mobile"`
	IsDesktop     bool                                    `json:"is_desktop"`
	Header        LandingSectionHeaderResponse            `json:"header"`
	Benefits      []LandingFeaturesContentBenefitResponse `json:"benefits"`
	FooterTitle   string                                  `json:"footer_title"`
	ButtonAbout   string                                  `json:"button_about"`
	ButtonPackage string                                  `json:"button_package"`
}

func NewLandingFeaturesContentResponse(dto dto.LandingFeaturesContentResponse) LandingFeaturesContentResponse {
	benefits := make([]LandingFeaturesContentBenefitResponse, len(dto.Benefits))
	for i, b := range dto.Benefits {
		benefits[i] = NewLandingFeaturesContentBenefitResponse(b)
	}
	return LandingFeaturesContentResponse{
		IsEnabled:     dto.IsEnabled,
		IsMobile:      dto.IsMobile,
		IsDesktop:     dto.IsDesktop,
		Header:        NewLandingSectionHeaderResponse(dto.Header),
		Benefits:      benefits,
		FooterTitle:   dto.FooterTitle,
		ButtonAbout:   dto.ButtonAbout,
		ButtonPackage: dto.ButtonPackage,
	}
}

func NewLandingFeaturesContentResponses(dtos []dto.LandingFeaturesContentResponse) []LandingFeaturesContentResponse {
	responses := make([]LandingFeaturesContentResponse, len(dtos))
	for i, dto := range dtos {
		responses[i] = NewLandingFeaturesContentResponse(dto)
	}
	return responses
}

type LandingMomentsContentImageResponse struct {
	IsEnabled bool          `json:"is_enabled"`
	IsMobile  bool          `json:"is_mobile"`
	IsDesktop bool          `json:"is_desktop"`
	Image     ImageResponse `json:"image"`
}

func NewLandingMomentsContentImageResponse(dto dto.LandingMomentsContentImageResponse) LandingMomentsContentImageResponse {
	return LandingMomentsContentImageResponse{
		IsEnabled: dto.IsEnabled,
		IsMobile:  dto.IsMobile,
		IsDesktop: dto.IsDesktop,
		Image:     NewImageResponse(dto.Image),
	}
}

func NewLandingMomentsContentImageResponses(dtos []dto.LandingMomentsContentImageResponse) []LandingMomentsContentImageResponse {
	responses := make([]LandingMomentsContentImageResponse, len(dtos))
	for i, dto := range dtos {
		responses[i] = NewLandingMomentsContentImageResponse(dto)
	}
	return responses
}

type LandingMomentsContentResponse struct {
	IsEnabled bool                                 `json:"is_enabled"`
	IsMobile  bool                                 `json:"is_mobile"`
	IsDesktop bool                                 `json:"is_desktop"`
	Header    LandingSectionHeaderResponse         `json:"header"`
	Images    []LandingMomentsContentImageResponse `json:"images"`
}

func NewLandingMomentsContentResponse(dto dto.LandingMomentsContentResponse) LandingMomentsContentResponse {
	images := make([]LandingMomentsContentImageResponse, len(dto.Images))
	for i, img := range dto.Images {
		images[i] = NewLandingMomentsContentImageResponse(img)
	}
	return LandingMomentsContentResponse{
		IsEnabled: dto.IsEnabled,
		IsMobile:  dto.IsMobile,
		IsDesktop: dto.IsDesktop,
		Header:    NewLandingSectionHeaderResponse(dto.Header),
		Images:    images,
	}
}

func NewLandingMomentsContentResponses(dtos []dto.LandingMomentsContentResponse) []LandingMomentsContentResponse {
	responses := make([]LandingMomentsContentResponse, len(dtos))
	for i, dto := range dtos {
		responses[i] = NewLandingMomentsContentResponse(dto)
	}
	return responses
}

type LandingAffiliatesContentAffiliateResponse struct {
	IsEnabled bool                      `json:"is_enabled"`
	IsMobile  bool                      `json:"is_mobile"`
	IsDesktop bool                      `json:"is_desktop"`
	Name      string                    `json:"name"`
	Logo      null.Value[ImageResponse] `json:"logo"`
	Width     int                       `json:"width"`
	Height    int                       `json:"height"`
}

func NewLandingAffiliatesContentAffiliateResponse(dto dto.LandingAffiliatesContentAffiliateResponse) LandingAffiliatesContentAffiliateResponse {
	logo := null.NewValue(ImageResponse{}, false)
	if dto.Logo.Valid {
		imageResponse := NewImageResponse(dto.Logo.V)

		logo = null.ValueFrom(imageResponse)
	}

	return LandingAffiliatesContentAffiliateResponse{
		IsEnabled: dto.IsEnabled,
		IsMobile:  dto.IsMobile,
		IsDesktop: dto.IsDesktop,
		Name:      dto.Name,
		Logo:      logo,
		Width:     dto.Width,
		Height:    dto.Height,
	}
}

func NewLandingAffiliatesContentAffiliateResponses(dtos []dto.LandingAffiliatesContentAffiliateResponse) []LandingAffiliatesContentAffiliateResponse {
	responses := make([]LandingAffiliatesContentAffiliateResponse, len(dtos))
	for i, dto := range dtos {
		responses[i] = NewLandingAffiliatesContentAffiliateResponse(dto)
	}
	return responses
}

type LandingAffiliatesContentResponse struct {
	IsEnabled  bool                                        `json:"is_enabled"`
	IsMobile   bool                                        `json:"is_mobile"`
	IsDesktop  bool                                        `json:"is_desktop"`
	Header     LandingSectionHeaderResponse                `json:"header"`
	Affiliates []LandingAffiliatesContentAffiliateResponse `json:"affiliates"`
}

func NewLandingAffiliatesContentResponse(dto dto.LandingAffiliatesContentResponse) LandingAffiliatesContentResponse {
	affiliates := make([]LandingAffiliatesContentAffiliateResponse, len(dto.Affiliates))
	for i, a := range dto.Affiliates {
		affiliates[i] = NewLandingAffiliatesContentAffiliateResponse(a)
	}
	return LandingAffiliatesContentResponse{
		IsEnabled:  dto.IsEnabled,
		IsMobile:   dto.IsMobile,
		IsDesktop:  dto.IsDesktop,
		Header:     NewLandingSectionHeaderResponse(dto.Header),
		Affiliates: affiliates,
	}
}

func NewLandingAffiliatesContentResponses(dtos []dto.LandingAffiliatesContentResponse) []LandingAffiliatesContentResponse {
	responses := make([]LandingAffiliatesContentResponse, len(dtos))
	for i, dto := range dtos {
		responses[i] = NewLandingAffiliatesContentResponse(dto)
	}
	return responses
}

type LandingTestimonialContentReviewResponse struct {
	IsEnabled bool    `json:"is_enabled"`
	IsMobile  bool    `json:"is_mobile"`
	IsDesktop bool    `json:"is_desktop"`
	Reviewer  string  `json:"reviewer"`
	Age       int     `json:"age"`
	Address   string  `json:"address"`
	Rating    float32 `json:"rating"`
	Review    string  `json:"review"`
}

func NewLandingTestimonialContentReviewResponse(dto dto.LandingTestimonialContentReviewResponse) LandingTestimonialContentReviewResponse {
	return LandingTestimonialContentReviewResponse{
		IsEnabled: dto.IsEnabled,
		IsMobile:  dto.IsMobile,
		IsDesktop: dto.IsDesktop,
		Reviewer:  dto.Reviewer,
		Age:       dto.Age,
		Address:   dto.Address,
		Rating:    dto.Rating,
		Review:    dto.Review,
	}
}

func NewLandingTestimonialContentReviewResponses(dtos []dto.LandingTestimonialContentReviewResponse) []LandingTestimonialContentReviewResponse {
	responses := make([]LandingTestimonialContentReviewResponse, len(dtos))
	for i, dto := range dtos {
		responses[i] = NewLandingTestimonialContentReviewResponse(dto)
	}
	return responses
}

type LandingTestimonialContentResponse struct {
	IsEnabled bool                                      `json:"is_enabled"`
	IsMobile  bool                                      `json:"is_mobile"`
	IsDesktop bool                                      `json:"is_desktop"`
	Header    LandingSectionHeaderResponse              `json:"header"`
	Reviews   []LandingTestimonialContentReviewResponse `json:"reviews"`
}

func NewLandingTestimonialContentResponse(dto dto.LandingTestimonialContentResponse) LandingTestimonialContentResponse {
	reviews := make([]LandingTestimonialContentReviewResponse, len(dto.Reviews))
	for i, review := range dto.Reviews {
		reviews[i] = NewLandingTestimonialContentReviewResponse(review)
	}
	return LandingTestimonialContentResponse{
		IsEnabled: dto.IsEnabled,
		IsMobile:  dto.IsMobile,
		IsDesktop: dto.IsDesktop,
		Header:    NewLandingSectionHeaderResponse(dto.Header),
		Reviews:   reviews,
	}
}

func NewLandingTestimonialContentResponses(dtos []dto.LandingTestimonialContentResponse) []LandingTestimonialContentResponse {
	responses := make([]LandingTestimonialContentResponse, len(dtos))
	for i, dto := range dtos {
		responses[i] = NewLandingTestimonialContentResponse(dto)
	}
	return responses
}

type LandingFaqContentFaqResponse struct {
	IsEnabled bool   `json:"is_enabled"`
	IsMobile  bool   `json:"is_mobile"`
	IsDesktop bool   `json:"is_desktop"`
	Question  string `json:"question"`
	Answer    string `json:"answer"`
}

func NewLandingFaqContentFaqResponse(dto dto.LandingFaqContentFaqResponse) LandingFaqContentFaqResponse {
	return LandingFaqContentFaqResponse{
		IsEnabled: dto.IsEnabled,
		IsMobile:  dto.IsMobile,
		IsDesktop: dto.IsDesktop,
		Question:  dto.Question,
		Answer:    dto.Answer,
	}
}

func NewLandingFaqContentFaqResponses(dtos []dto.LandingFaqContentFaqResponse) []LandingFaqContentFaqResponse {
	responses := make([]LandingFaqContentFaqResponse, len(dtos))
	for i, dto := range dtos {
		responses[i] = NewLandingFaqContentFaqResponse(dto)
	}
	return responses
}

type LandingFaqContentResponse struct {
	IsEnabled bool                           `json:"is_enabled"`
	IsMobile  bool                           `json:"is_mobile"`
	IsDesktop bool                           `json:"is_desktop"`
	Header    LandingSectionHeaderResponse   `json:"header"`
	Faqs      []LandingFaqContentFaqResponse `json:"faqs"`
}

func NewLandingFaqContentResponse(dto dto.LandingFaqContentResponse) LandingFaqContentResponse {
	faqs := make([]LandingFaqContentFaqResponse, len(dto.Faqs))
	for i, faq := range dto.Faqs {
		faqs[i] = NewLandingFaqContentFaqResponse(faq)
	}
	return LandingFaqContentResponse{
		IsEnabled: dto.IsEnabled,
		IsMobile:  dto.IsMobile,
		IsDesktop: dto.IsDesktop,
		Header:    NewLandingSectionHeaderResponse(dto.Header),
		Faqs:      faqs,
	}
}

func NewLandingFaqContentResponses(dtos []dto.LandingFaqContentResponse) []LandingFaqContentResponse {
	responses := make([]LandingFaqContentResponse, len(dtos))
	for i, dto := range dtos {
		responses[i] = NewLandingFaqContentResponse(dto)
	}
	return responses
}

type LandingMenuResponse struct {
	IsEnabled bool   `json:"is_enabled"`
	IsMobile  bool   `json:"is_mobile"`
	IsDesktop bool   `json:"is_desktop"`
	Icon      string `json:"icon"`
	Label     string `json:"label"`
	Path      string `json:"path"`
}

func NewLandingMenuResponse(dto dto.LandingMenuResponse) LandingMenuResponse {
	return LandingMenuResponse{
		IsEnabled: dto.IsEnabled,
		IsMobile:  dto.IsMobile,
		IsDesktop: dto.IsDesktop,
		Icon:      dto.Icon,
		Label:     dto.Label,
		Path:      dto.Path,
	}
}

func NewLandingMenuResponses(dtos []dto.LandingMenuResponse) []LandingMenuResponse {
	responses := make([]LandingMenuResponse, len(dtos))
	for i, dto := range dtos {
		responses[i] = NewLandingMenuResponse(dto)
	}
	return responses
}

type LandingResponse struct {
	HeroContent              LandingHeroContentResponse              `json:"hero_content"`
	SinglePackageContent     LandingSinglePackageContentResponse     `json:"single_package_content"`
	PackagesContent          LandingPackagesContentResponse          `json:"packages_content"`
	TravelDestinationContent LandingTravelDestinationContentResponse `json:"travel_destination_content"`
	FeaturesContent          LandingFeaturesContentResponse          `json:"features_content"`
	MomentsContent           LandingMomentsContentResponse           `json:"moments_content"`
	AffiliatesContent        LandingAffiliatesContentResponse        `json:"affiliates_content"`
	TestimonialContent       LandingTestimonialContentResponse       `json:"testimonial_content"`
	FaqContent               LandingFaqContentResponse               `json:"faq_content"`
	Menus                    []LandingMenuResponse                   `json:"menus"`
}

func NewLandingResponse(dto dto.LandingResponse) LandingResponse {
	return LandingResponse{
		HeroContent:              NewLandingHeroContentResponse(dto.HeroContent),
		SinglePackageContent:     NewLandingSinglePackageContentResponse(dto.SinglePackageContent),
		PackagesContent:          NewLandingPackagesContentResponse(dto.PackagesContent),
		TravelDestinationContent: NewLandingTravelDestinationContentResponse(dto.TravelDestinationContent),
		FeaturesContent:          NewLandingFeaturesContentResponse(dto.FeaturesContent),
		MomentsContent:           NewLandingMomentsContentResponse(dto.MomentsContent),
		AffiliatesContent:        NewLandingAffiliatesContentResponse(dto.AffiliatesContent),
		TestimonialContent:       NewLandingTestimonialContentResponse(dto.TestimonialContent),
		FaqContent:               NewLandingFaqContentResponse(dto.FaqContent),
		Menus:                    NewLandingMenuResponses(dto.Menus),
	}
}

func NewLandingResponses(dtos []dto.LandingResponse) []LandingResponse {
	responses := make([]LandingResponse, len(dtos))
	for i, dto := range dtos {
		responses[i] = NewLandingResponse(dto)
	}
	return responses
}

// type LandingResponse struct {
// 	Id        int64   `json:"id"`
// 	Name      string  `json:"name"`
// 	Latitude  float64 `json:"latitude"`
// 	Longitude float64 `json:"longitude"`
// 	Slug      string  `json:"slug"`

// 	CreatedAt time.Time `json:"created_at"`
// 	UpdatedAt time.Time `json:"updated_at"`
// 	DeletedAt null.Time `json:"deleted_at"`
// }

// func NewLandingResponse(dtoResponse dto.LandingResponse) LandingResponse {
// 	return LandingResponse{
// 		Id:        dtoResponse.Id,
// 		Name:      dtoResponse.Name,
// 		Latitude:  dtoResponse.Latitude,
// 		Longitude: dtoResponse.Longitude,
// 		Slug:      dtoResponse.Slug,
// 		CreatedAt: dtoResponse.CreatedAt,
// 		UpdatedAt: dtoResponse.UpdatedAt,
// 		DeletedAt: dtoResponse.DeletedAt,
// 	}
// }

// func NewLandingResponses(dtoResponses []dto.LandingResponse) []LandingResponse {
// 	embarkationResponses := make([]LandingResponse, len(dtoResponses))

// 	for i, dtoResponse := range dtoResponses {
// 		embarkationResponses[i] = NewLandingResponse(dtoResponse)
// 	}

// 	return embarkationResponses
// }
