package validator

import (
	"fmt"

	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type LandingValidator struct {
}

func NewLandingValidator() LandingValidator {
	return LandingValidator{}
}

func (v LandingValidator) validateSectionHeaderRequest(
	prefix string,
	sectionHeaderRequest dto.LandingSectionHeaderRequest,
) error {
	titleLength := len(sectionHeaderRequest.Title)
	if titleLength < 1 {
		return newError(fmt.Sprintf("%s.Title", prefix), mustBeNotEmpty)
	}
	if titleLength > 100 {
		return newError(fmt.Sprintf("%s.Title", prefix), maxChars(100))
	}

	if sectionHeaderRequest.Subtitle.Valid {
		subtitleLength := len(sectionHeaderRequest.Subtitle.String)
		if subtitleLength > 100 {
			return newError(fmt.Sprintf("%s.Subtitle", prefix), maxChars(100))
		}
	}

	if sectionHeaderRequest.TagsLine.Valid {
		tagsLineLength := len(sectionHeaderRequest.TagsLine.String)
		if tagsLineLength > 100 {
			return newError(fmt.Sprintf("%s.TagsLine", prefix), maxChars(100))
		}
	}

	return nil
}

func (v LandingValidator) validatePackageItemRequest(
	prefix string,
	packageItemRequest dto.LandingPackageItemRequest,
) error {
	if packageItemRequest.Package < 1 {
		return newError(fmt.Sprintf("%s.Package", prefix), mustBeGte(1))
	}

	buttonLabelLength := len(packageItemRequest.ButtonLabel)
	if buttonLabelLength < 1 {
		return newError(fmt.Sprintf("%s.ButtonLabel", prefix), mustBeNotEmpty)
	}
	if buttonLabelLength > 100 {
		return newError(fmt.Sprintf("%s.ButtonLabel", prefix), maxChars(100))
	}

	return nil
}

func (v LandingValidator) validatePackageDetailRequest(
	prefix string,
	packageDetailRequest dto.LandingPackageDetailRequest,
) error {
	if err := v.validateSectionHeaderRequest(
		fmt.Sprintf("%s.Header", prefix),
		packageDetailRequest.Header,
	); err != nil {
		return err
	}

	for i, packageItemRequest := range packageDetailRequest.Packages {
		if err := v.validatePackageItemRequest(
			fmt.Sprintf("%s.Packages.%d", prefix, i),
			packageItemRequest,
		); err != nil {
			return err
		}
	}

	return nil
}

func (v LandingValidator) validateFeatureContentBenefitRequest(
	prefix string,
	featureContentBenefitRequest dto.LandingFeaturesContentBenefitRequest,
) error {
	titleLength := len(featureContentBenefitRequest.Title)
	if titleLength < 1 {
		return newError(fmt.Sprintf("%s.Title", prefix), mustBeNotEmpty)
	}
	if titleLength > 100 {
		return newError(fmt.Sprintf("%s.Title", prefix), maxChars(100))
	}

	subtitleLength := len(featureContentBenefitRequest.Subtitle)
	if subtitleLength < 1 {
		return newError(fmt.Sprintf("%s.Subtitle", prefix), mustBeNotEmpty)
	}
	if subtitleLength > 500 {
		return newError(fmt.Sprintf("%s.Subtitle", prefix), maxChars(500))
	}

	if featureContentBenefitRequest.Logo.Valid {
		if featureContentBenefitRequest.Logo.Int64 < 1 {
			return newError(fmt.Sprintf("%s.Logo", prefix), mustBeGte(1))
		}
	}

	return nil
}

func (v LandingValidator) validateAffiliatesContentAffiliateRequest(
	prefix string,
	affiliatesContentAffiliateRequest dto.LandingAffiliatesContentAffiliateRequest,
) error {
	nameLength := len(affiliatesContentAffiliateRequest.Name)
	if nameLength < 1 {
		return newError(fmt.Sprintf("%s.Name", prefix), mustBeNotEmpty)
	}
	if nameLength > 100 {
		return newError(fmt.Sprintf("%s.Name", prefix), maxChars(100))
	}

	if affiliatesContentAffiliateRequest.Logo.Valid {
		if affiliatesContentAffiliateRequest.Logo.Int64 < 1 {
			return newError(fmt.Sprintf("%s.Logo", prefix), mustBeGte(1))
		}
	}

	if affiliatesContentAffiliateRequest.Width < 1 {
		return newError(fmt.Sprintf("%s.Width", prefix), mustBeGte(1))
	}

	if affiliatesContentAffiliateRequest.Height < 1 {
		return newError(fmt.Sprintf("%s.Height", prefix), mustBeGte(1))
	}

	return nil
}

func (v LandingValidator) validateTestimonialContentReviewRequest(
	prefix string,
	testimonialContentReviewRequest dto.LandingTestimonialContentReviewRequest,
) error {
	reviewerLength := len(testimonialContentReviewRequest.Reviewer)
	if reviewerLength < 1 {
		return newError(fmt.Sprintf("%s.Reviewer", prefix), mustBeNotEmpty)
	}
	if reviewerLength > 100 {
		return newError(fmt.Sprintf("%s.Reviewer", prefix), maxChars(100))
	}

	if testimonialContentReviewRequest.Age < 1 {
		return newError(fmt.Sprintf("%s.Age", prefix), mustBeGte(1))
	}

	addressLength := len(testimonialContentReviewRequest.Address)
	if addressLength < 1 {
		return newError(fmt.Sprintf("%s.Address", prefix), mustBeNotEmpty)
	}
	if addressLength > 500 {
		return newError(fmt.Sprintf("%s.Address", prefix), maxChars(500))
	}

	if testimonialContentReviewRequest.Rating < 1 || testimonialContentReviewRequest.Rating > 5 {
		return newError(fmt.Sprintf("%s.Rating", prefix), mustBetween(1, 5))
	}

	if len(testimonialContentReviewRequest.Review) < 1 {
		return newError(fmt.Sprintf("%s.Review", prefix), mustBeNotEmpty)
	}

	return nil
}

func (v LandingValidator) validateFaqContentFaqRequest(
	prefix string,
	faqContentFaqRequest dto.LandingFaqContentFaqRequest,
) error {
	questionLength := len(faqContentFaqRequest.Question)
	if questionLength < 1 {
		return newError(fmt.Sprintf("%s.Question", prefix), mustBeNotEmpty)
	}
	if questionLength > 100 {
		return newError(fmt.Sprintf("%s.Question", prefix), maxChars(100))
	}

	answerLength := len(faqContentFaqRequest.Answer)
	if answerLength < 1 {
		return newError(fmt.Sprintf("%s.Answer", prefix), mustBeNotEmpty)
	}
	if answerLength > 500 {
		return newError(fmt.Sprintf("%s.Answer", prefix), maxChars(500))
	}

	return nil
}

func (v LandingValidator) validateTravelDestinationContentDestinationRequest(
	prefix string,
	travelDestinationContentDestinationRequest dto.LandingTravelDestinationContentDestinationRequest,
) error {
	if travelDestinationContentDestinationRequest.Image.Valid {
		if travelDestinationContentDestinationRequest.Image.Int64 < 1 {
			return newError(fmt.Sprintf("%s.Image", prefix), mustBeGte(1))
		}
	}

	nameLength := len(travelDestinationContentDestinationRequest.Name)
	if nameLength < 1 {
		return newError(fmt.Sprintf("%s.Name", prefix), mustBeNotEmpty)
	}
	if nameLength > 100 {
		return newError(fmt.Sprintf("%s.Name", prefix), maxChars(100))
	}

	return nil
}

func (v LandingValidator) validateHeroContentRequest(
	prefix string,
	heroContentRequest dto.LandingHeroContentRequest,
) error {
	titleLength := len(heroContentRequest.Title)
	if titleLength < 1 {
		return newError(fmt.Sprintf("%s.Title", prefix), mustBeNotEmpty)
	}
	if titleLength > 100 {
		return newError(fmt.Sprintf("%s.Title", prefix), maxChars(100))
	}

	descriptionLength := len(heroContentRequest.Description)
	if descriptionLength < 1 {
		return newError(fmt.Sprintf("%s.Description", prefix), mustBeNotEmpty)
	}
	if descriptionLength > 500 {
		return newError(fmt.Sprintf("%s.Description", prefix), maxChars(500))
	}

	tagsLineLength := len(heroContentRequest.TagsLine)
	if tagsLineLength < 1 {
		return newError(fmt.Sprintf("%s.TagsLine", prefix), mustBeNotEmpty)
	}
	if tagsLineLength > 100 {
		return newError(fmt.Sprintf("%s.TagsLine", prefix), maxChars(100))
	}

	buttonLabelLength := len(heroContentRequest.ButtonLabel)
	if buttonLabelLength < 1 {
		return newError(fmt.Sprintf("%s.ButtonLabel", prefix), mustBeNotEmpty)
	}
	if buttonLabelLength > 100 {
		return newError(fmt.Sprintf("%s.ButtonLabel", prefix), maxChars(100))
	}

	buttonUrlLength := len(heroContentRequest.ButtonUrl)
	if buttonUrlLength < 1 {
		return newError(fmt.Sprintf("%s.ButtonUrl", prefix), mustBeNotEmpty)
	}
	if buttonUrlLength > 100 {
		return newError(fmt.Sprintf("%s.ButtonUrl", prefix), maxChars(100))
	}

	if heroContentRequest.Image.Valid {
		if heroContentRequest.Image.Int64 < 1 {
			return newError(fmt.Sprintf("%s.Image", prefix), mustBeGte(1))
		}
	}

	return nil
}

func (v LandingValidator) validateSinglePackageContentRequest(
	prefix string,
	singlePackageContentRequest dto.LandingSinglePackageContentRequest,
) error {
	if err := v.validateSectionHeaderRequest(
		fmt.Sprintf("%s.Header", prefix),
		singlePackageContentRequest.Header,
	); err != nil {
		return err
	}

	if singlePackageContentRequest.Silver.Valid {
		if err := v.validatePackageItemRequest(
			fmt.Sprintf("%s.Silver", prefix),
			singlePackageContentRequest.Silver.V,
		); err != nil {
			return err
		}
	}

	if singlePackageContentRequest.Gold.Valid {
		if err := v.validatePackageItemRequest(
			fmt.Sprintf("%s.Gold", prefix),
			singlePackageContentRequest.Gold.V,
		); err != nil {
			return err
		}
	}

	if singlePackageContentRequest.Platinum.Valid {
		if err := v.validatePackageItemRequest(
			fmt.Sprintf("%s.Platinum", prefix),
			singlePackageContentRequest.Platinum.V,
		); err != nil {
			return err
		}
	}

	return nil
}

func (v LandingValidator) validatePackagesContentRequest(
	prefix string,
	packagesContentRequest dto.LandingPackagesContentRequest,
) error {
	if err := v.validatePackageDetailRequest(
		fmt.Sprintf("%s.Silver", prefix),
		packagesContentRequest.Silver,
	); err != nil {
		return err
	}

	if err := v.validatePackageDetailRequest(
		fmt.Sprintf("%s.Gold", prefix),
		packagesContentRequest.Gold,
	); err != nil {
		return err
	}

	if err := v.validatePackageDetailRequest(
		fmt.Sprintf("%s.Platinum", prefix),
		packagesContentRequest.Platinum,
	); err != nil {
		return err
	}

	return nil
}

func (v LandingValidator) validateTravelDestinationContentRequest(
	prefix string,
	travelDestinationContentRequest dto.LandingTravelDestinationContentRequest,
) error {
	if err := v.validateSectionHeaderRequest(
		fmt.Sprintf("%s.Header", prefix),
		travelDestinationContentRequest.Header,
	); err != nil {
		return err
	}

	for i, travelDestinationRequest := range travelDestinationContentRequest.Destinations {
		if err := v.validateTravelDestinationContentDestinationRequest(
			fmt.Sprintf("%s.Destinations.%d", prefix, i),
			travelDestinationRequest,
		); err != nil {
			return err
		}
	}

	return nil
}

func (v LandingValidator) validateFeaturesContentRequest(
	prefix string,
	featuresContentRequest dto.LandingFeaturesContentRequest,
) error {
	if err := v.validateSectionHeaderRequest(
		fmt.Sprintf("%s.Header", prefix),
		featuresContentRequest.Header,
	); err != nil {
		return err
	}

	for i, featureContentBenefitRequest := range featuresContentRequest.Benefits {
		if err := v.validateFeatureContentBenefitRequest(
			fmt.Sprintf("%s.Benefits.%d", prefix, i),
			featureContentBenefitRequest,
		); err != nil {
			return err
		}
	}

	footerTitleLength := len(featuresContentRequest.FooterTitle)
	if footerTitleLength < 1 {
		return newError(fmt.Sprintf("%s.FooterTitle", prefix), mustBeNotEmpty)
	}
	if footerTitleLength > 100 {
		return newError(fmt.Sprintf("%s.FooterTitle", prefix), maxChars(100))
	}

	buttonAboutLength := len(featuresContentRequest.ButtonAbout)
	if buttonAboutLength < 1 {
		return newError(fmt.Sprintf("%s.ButtonAbout", prefix), mustBeNotEmpty)
	}
	if buttonAboutLength > 100 {
		return newError(fmt.Sprintf("%s.ButtonAbout", prefix), maxChars(100))
	}

	buttonPackageLength := len(featuresContentRequest.ButtonPackage)
	if buttonPackageLength < 1 {
		return newError(fmt.Sprintf("%s.ButtonPackage", prefix), mustBeNotEmpty)
	}
	if buttonPackageLength > 100 {
		return newError(fmt.Sprintf("%s.ButtonPackage", prefix), maxChars(100))
	}

	return nil
}

func (v LandingValidator) validateMomentsContentRequest(
	prefix string,
	momentsContentRequest dto.LandingMomentsContentRequest,
) error {
	if err := v.validateSectionHeaderRequest(
		fmt.Sprintf("%s.Header", prefix),
		momentsContentRequest.Header,
	); err != nil {
		return err
	}

	for i, image := range momentsContentRequest.Images {
		if image.Image < 1 {
			return newError(fmt.Sprintf("%s.Images.%d.Image", prefix, i), mustBeGte(1))
		}
	}

	return nil
}

func (v LandingValidator) validateAffiliatesContentRequest(
	prefix string,
	affiliatesContentRequest dto.LandingAffiliatesContentRequest,
) error {
	if err := v.validateSectionHeaderRequest(
		fmt.Sprintf("%s.Header", prefix),
		affiliatesContentRequest.Header,
	); err != nil {
		return err
	}

	for i, affiliateRequest := range affiliatesContentRequest.Affiliates {
		if err := v.validateAffiliatesContentAffiliateRequest(
			fmt.Sprintf("%s.Affiliates.%d", prefix, i),
			affiliateRequest,
		); err != nil {
			return err
		}
	}

	return nil
}

func (v LandingValidator) validateTestimonialContentRequest(
	prefix string,
	testimonialContentRequest dto.LandingTestimonialContentRequest,
) error {
	if err := v.validateSectionHeaderRequest(
		fmt.Sprintf("%s.Header", prefix),
		testimonialContentRequest.Header,
	); err != nil {
		return err
	}

	for i, testimonialRequest := range testimonialContentRequest.Reviews {
		if err := v.validateTestimonialContentReviewRequest(
			fmt.Sprintf("%s.Reviews.%d", prefix, i),
			testimonialRequest,
		); err != nil {
			return err
		}
	}

	return nil
}

func (v LandingValidator) validateFaqContentRequest(
	prefix string,
	faqContentRequest dto.LandingFaqContentRequest,
) error {
	if err := v.validateSectionHeaderRequest(
		fmt.Sprintf("%s.Header", prefix),
		faqContentRequest.Header,
	); err != nil {
		return err
	}

	for i, faqRequest := range faqContentRequest.Faqs {
		if err := v.validateFaqContentFaqRequest(
			fmt.Sprintf("%s.Faqs.%d", prefix, i),
			faqRequest,
		); err != nil {
			return err
		}
	}

	return nil
}

func (v LandingValidator) validateMenuRequest(
	prefix string,
	menuRequest dto.LandingMenuRequest,
) error {
	iconLength := len(menuRequest.Icon)
	if iconLength < 1 {
		return newError(fmt.Sprintf("%s.Icon", prefix), mustBeNotEmpty)
	}
	if iconLength > 100 {
		return newError(fmt.Sprintf("%s.Icon", prefix), maxChars(100))
	}

	labelLength := len(menuRequest.Label)
	if labelLength < 1 {
		return newError(fmt.Sprintf("%s.Label", prefix), mustBeNotEmpty)
	}
	if labelLength > 100 {
		return newError(fmt.Sprintf("%s.Label", prefix), maxChars(100))
	}

	pathLength := len(menuRequest.Path)
	if pathLength < 1 {
		return newError(fmt.Sprintf("%s.Path", prefix), mustBeNotEmpty)
	}
	if pathLength > 100 {
		return newError(fmt.Sprintf("%s.Path", prefix), maxChars(100))
	}

	return nil
}

func (v LandingValidator) ValidateRequest(request dto.LandingRequest) error {
	if err := v.validateHeroContentRequest("HeroContent", request.HeroContent); err != nil {
		return err
	}

	if err := v.validateSinglePackageContentRequest("SinglePackageContent", request.SinglePackageContent); err != nil {
		return err
	}

	if err := v.validatePackagesContentRequest("PackagesContent", request.PackagesContent); err != nil {
		return err
	}

	if err := v.validateTravelDestinationContentRequest("TravelDestinationContent", request.TravelDestinationContent); err != nil {
		return err
	}

	if err := v.validateFeaturesContentRequest("FeaturesContent", request.FeaturesContent); err != nil {
		return err
	}

	if err := v.validateMomentsContentRequest("MomentsContent", request.MomentsContent); err != nil {
		return err
	}

	if err := v.validateAffiliatesContentRequest("AffiliatesContent", request.AffiliatesContent); err != nil {
		return err
	}

	if err := v.validateTestimonialContentRequest("TestimonialContent", request.TestimonialContent); err != nil {
		return err
	}

	if err := v.validateFaqContentRequest("FaqContent", request.FaqContent); err != nil {
		return err
	}

	for i, menuRequest := range request.Menus {
		if err := v.validateMenuRequest(
			fmt.Sprintf("Menus.%d", i),
			menuRequest,
		); err != nil {
			return err
		}
	}

	return nil
}
