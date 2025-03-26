package entity

import (
	"time"

	"github.com/guregu/null/v6"
)

type LandingHeroContent struct {
	Id          int64
	IsEnabled   bool
	IsMobile    bool
	IsDesktop   bool
	Title       string
	Description string
	TagsLine    string
	ButtonLabel string
	ButtonUrl   string
	ImageId     null.Int64

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}

type LandingSectionHeader struct {
	Id        int64
	IsEnabled bool
	IsMobile  bool
	IsDesktop bool
	Title     string
	Subtitle  null.String
	TagsLine  null.String

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}

type LandingPackageItem struct {
	Id          int64
	IsEnabled   bool
	IsMobile    bool
	IsDesktop   bool
	PackageId   int64
	ButtonLabel string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}

type LandingSinglePackageContent struct {
	Id                           int64
	IsEnabled                    bool
	IsMobile                     bool
	IsDesktop                    bool
	LandingSectionHeaderId       int64
	SilverLandingPackageItemId   null.Int64
	GoldLandingPackageItemId     null.Int64
	PlatinumLandingPackageItemId null.Int64

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}

type LandingPackageDetail struct {
	Id                     int64
	IsEnabled              bool
	IsMobile               bool
	IsDesktop              bool
	LandingSectionHeaderId int64

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}

type LandingPackageDetailItem struct {
	IsEnabled              bool
	IsMobile               bool
	IsDesktop              bool
	LandingPackageDetailId int64
	LandingPackageItemId   int64

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}

type LandingPackagesContent struct {
	Id                             int64
	IsEnabled                      bool
	IsMobile                       bool
	IsDesktop                      bool
	SilverLandingPackageDetailId   int64
	GoldLandingPackageDetailId     int64
	PlatinumLandingPackageDetailId int64

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}

type LandingFeaturesContent struct {
	Id                     int64
	IsEnabled              bool
	IsMobile               bool
	IsDesktop              bool
	LandingSectionHeaderId int64
	FooterTitle            string
	ButtonAbout            string
	ButtonPackage          string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}

type LandingFeaturesContentBenefit struct {
	Id        int64
	IsEnabled bool
	IsMobile  bool
	IsDesktop bool
	Title     string
	Subtitle  string
	LogoId    null.Int64

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}

type LandingMomentsContent struct {
	Id                     int64
	IsEnabled              bool
	IsMobile               bool
	IsDesktop              bool
	LandingSectionHeaderId int64

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}

type LandingMomentsContentImage struct {
	IsEnabled bool
	IsMobile  bool
	IsDesktop bool
	ImageId   int64

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}

type LandingAffiliatesContent struct {
	Id                     int64
	IsEnabled              bool
	IsMobile               bool
	IsDesktop              bool
	LandingSectionHeaderId int64

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}

type LandingAffiliatesContentAffiliate struct {
	Id        int64
	IsEnabled bool
	IsMobile  bool
	IsDesktop bool
	Name      string
	LogoId    null.Int64
	Width     int
	Height    int

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}

type LandingFaqContent struct {
	Id                     int64
	IsEnabled              bool
	IsMobile               bool
	IsDesktop              bool
	LandingSectionHeaderId int64

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}

type LandingFaqContentFaq struct {
	Id        int64
	IsEnabled bool
	IsMobile  bool
	IsDesktop bool
	Question  string
	Answer    string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}

type LandingMenu struct {
	Id        int64
	IsEnabled bool
	IsMobile  bool
	IsDesktop bool
	Icon      string
	Label     string
	Path      string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}
