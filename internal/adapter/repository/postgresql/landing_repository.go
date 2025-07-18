package postgresql

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driven/repository"
	"github.com/kabarhaji-id/goumrah-api/pkg/sqlbuilder"
)

type landingHeroContentRepositoryPostgresql struct {
	db DB
}

func NewLandingHeroContentRepository(db DB) repository.LandingHeroContentRepository {
	return landingHeroContentRepositoryPostgresql{db}
}

func (r landingHeroContentRepositoryPostgresql) Create(ctx context.Context, landingHeroContent entity.LandingHeroContent) (entity.LandingHeroContent, error) {
	builder := sqlbuilder.New().
		S(`INSERT INTO "landing_hero_content" ("is_enabled", "is_mobile", "is_desktop", "title", "description", "tags_line", "button_label", "button_url", "image_id", "created_at", "updated_at", "deleted_at")`).
		S(`VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, NOW(), NOW(), NULL)`, landingHeroContent.IsEnabled, landingHeroContent.IsMobile, landingHeroContent.IsDesktop, landingHeroContent.Title, landingHeroContent.Description, landingHeroContent.TagsLine, landingHeroContent.ButtonLabel, landingHeroContent.ButtonUrl, landingHeroContent.ImageId).
		S(`RETURNING "id", "is_enabled", "is_mobile", "is_desktop", "title", "description", "tags_line", "button_label", "button_url", "image_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingHeroContent.Id, &landingHeroContent.IsEnabled, &landingHeroContent.IsMobile, &landingHeroContent.IsDesktop, &landingHeroContent.Title, &landingHeroContent.Description, &landingHeroContent.TagsLine, &landingHeroContent.ButtonLabel, &landingHeroContent.ButtonUrl, &landingHeroContent.ImageId,
		&landingHeroContent.CreatedAt, &landingHeroContent.UpdatedAt, &landingHeroContent.DeletedAt,
	)

	return landingHeroContent, err
}

func (r landingHeroContentRepositoryPostgresql) Find(ctx context.Context) (entity.LandingHeroContent, error) {
	landingHeroContent := entity.LandingHeroContent{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "is_enabled", "is_mobile", "is_desktop", "title", "description", "tags_line", "button_label", "button_url", "image_id", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "landing_hero_content" WHERE "deleted_at" IS NULL`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingHeroContent.Id, &landingHeroContent.IsEnabled, &landingHeroContent.IsMobile, &landingHeroContent.IsDesktop, &landingHeroContent.Title, &landingHeroContent.Description, &landingHeroContent.TagsLine, &landingHeroContent.ButtonLabel, &landingHeroContent.ButtonUrl, &landingHeroContent.ImageId,
		&landingHeroContent.CreatedAt, &landingHeroContent.UpdatedAt, &landingHeroContent.DeletedAt,
	)

	return landingHeroContent, err
}

func (r landingHeroContentRepositoryPostgresql) Update(ctx context.Context, landingHeroContent entity.LandingHeroContent) (entity.LandingHeroContent, error) {
	builder := sqlbuilder.New().
		S(`UPDATE "landing_hero_content" SET "is_enabled" = $1, "is_mobile" = $2, "is_desktop" = $3, "title" = $4, "description" = $5, "tags_line" = $6, "button_label" = $7, "button_url" = $8, "image_id" = $9, "updated_at" = NOW()`, landingHeroContent.IsEnabled, landingHeroContent.IsMobile, landingHeroContent.IsDesktop, landingHeroContent.Title, landingHeroContent.Description, landingHeroContent.TagsLine, landingHeroContent.ButtonLabel, landingHeroContent.ButtonUrl, landingHeroContent.ImageId).
		S(`WHERE "deleted_at" IS NULL`).
		S(`RETURNING "id", "is_enabled", "is_mobile", "is_desktop", "title", "description", "tags_line", "button_label", "button_url", "image_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingHeroContent.Id, &landingHeroContent.IsEnabled, &landingHeroContent.IsMobile, &landingHeroContent.IsDesktop, &landingHeroContent.Title, &landingHeroContent.Description, &landingHeroContent.TagsLine, &landingHeroContent.ButtonLabel, &landingHeroContent.ButtonUrl, &landingHeroContent.ImageId,
		&landingHeroContent.CreatedAt, &landingHeroContent.UpdatedAt, &landingHeroContent.DeletedAt,
	)

	return landingHeroContent, err
}

func (r landingHeroContentRepositoryPostgresql) Delete(ctx context.Context) (entity.LandingHeroContent, error) {
	landingHeroContent := entity.LandingHeroContent{}

	builder := sqlbuilder.New().
		S(`DELETE FROM "landing_hero_content"`).
		S(`RETURNING "id", "is_enabled", "is_mobile", "is_desktop", "title", "description", "tags_line", "button_label", "button_url", "image_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingHeroContent.Id, &landingHeroContent.IsEnabled, &landingHeroContent.IsMobile, &landingHeroContent.IsDesktop, &landingHeroContent.Title, &landingHeroContent.Description, &landingHeroContent.TagsLine, &landingHeroContent.ButtonLabel, &landingHeroContent.ButtonUrl, &landingHeroContent.ImageId,
		&landingHeroContent.CreatedAt, &landingHeroContent.UpdatedAt, &landingHeroContent.DeletedAt,
	)

	return landingHeroContent, err
}

type landingSectionHeaderRepositoryPostgresql struct {
	db DB
}

func NewLandingSectionHeaderRepository(db DB) repository.LandingSectionHeaderRepository {
	return landingSectionHeaderRepositoryPostgresql{db}
}

func (r landingSectionHeaderRepositoryPostgresql) Create(ctx context.Context, landingSectionHeader entity.LandingSectionHeader) (entity.LandingSectionHeader, error) {
	builder := sqlbuilder.New().
		S(`INSERT INTO "landing_section_headers" ("is_enabled", "is_mobile", "is_desktop", "title", "subtitle", "tags_line", "created_at", "updated_at", "deleted_at")`).
		S(`VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW(), NULL)`, landingSectionHeader.IsEnabled, landingSectionHeader.IsMobile, landingSectionHeader.IsDesktop, landingSectionHeader.Title, landingSectionHeader.Subtitle, landingSectionHeader.TagsLine).
		S(`RETURNING "id", "is_enabled", "is_mobile", "is_desktop", "title", "subtitle", "tags_line", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingSectionHeader.Id, &landingSectionHeader.IsEnabled, &landingSectionHeader.IsMobile, &landingSectionHeader.IsDesktop, &landingSectionHeader.Title, &landingSectionHeader.Subtitle, &landingSectionHeader.TagsLine,
		&landingSectionHeader.CreatedAt, &landingSectionHeader.UpdatedAt, &landingSectionHeader.DeletedAt,
	)

	return landingSectionHeader, err
}

func (r landingSectionHeaderRepositoryPostgresql) FindById(ctx context.Context, id int64) (entity.LandingSectionHeader, error) {
	builder := sqlbuilder.New().
		S(`SELECT "id", "is_enabled", "is_mobile", "is_desktop", "title", "subtitle", "tags_line", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "landing_section_headers" WHERE "id" = $1 AND "deleted_at" IS NULL`, id)

	landingSectionHeader := entity.LandingSectionHeader{}
	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingSectionHeader.Id, &landingSectionHeader.IsEnabled, &landingSectionHeader.IsMobile, &landingSectionHeader.IsDesktop, &landingSectionHeader.Title, &landingSectionHeader.Subtitle, &landingSectionHeader.TagsLine,
		&landingSectionHeader.CreatedAt, &landingSectionHeader.UpdatedAt, &landingSectionHeader.DeletedAt,
	)

	return landingSectionHeader, err
}

func (r landingSectionHeaderRepositoryPostgresql) FindAll(ctx context.Context, opt repository.FindAllOptions) ([]entity.LandingSectionHeader, error) {
	builder := sqlbuilder.New().
		S(`SELECT "id", "is_enabled", "is_mobile", "is_desktop", "title", "subtitle", "tags_line", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "landing_section_headers" WHERE "deleted_at" IS NULL`).
		S(`ORDER BY "id" ASC`)
	if opt.Limit.Valid {
		builder.SA(`LIMIT ?`, opt.Limit)
	}
	if opt.Offset.Valid {
		builder.SA(`OFFSET ?`, opt.Offset)
	}

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return nil, err
	}

	landingSectionHeaders := []entity.LandingSectionHeader{}
	for rows.Next() {
		landingSectionHeader := entity.LandingSectionHeader{}
		err := rows.Scan(
			&landingSectionHeader.Id, &landingSectionHeader.IsEnabled, &landingSectionHeader.IsMobile, &landingSectionHeader.IsDesktop, &landingSectionHeader.Title, &landingSectionHeader.Subtitle, &landingSectionHeader.TagsLine,
			&landingSectionHeader.CreatedAt, &landingSectionHeader.UpdatedAt, &landingSectionHeader.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		landingSectionHeaders = append(landingSectionHeaders, landingSectionHeader)
	}

	return landingSectionHeaders, nil
}

func (r landingSectionHeaderRepositoryPostgresql) Update(ctx context.Context, id int64, landingSectionHeader entity.LandingSectionHeader) (entity.LandingSectionHeader, error) {
	builder := sqlbuilder.New().
		S(`UPDATE "landing_section_headers" SET "is_enabled" = $1, "is_mobile" = $2, "is_desktop" = $3, "title" = $4, "subtitle" = $5, "tags_line" = $6, "updated_at" = NOW()`).
		S(`WHERE "id" = $7 AND "deleted_at" IS NULL`, landingSectionHeader.IsEnabled, landingSectionHeader.IsMobile, landingSectionHeader.IsDesktop, landingSectionHeader.Title, landingSectionHeader.Subtitle, landingSectionHeader.TagsLine, id).
		S(`RETURNING "id", "is_enabled", "is_mobile", "is_desktop", "title", "subtitle", "tags_line", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingSectionHeader.Id, &landingSectionHeader.IsEnabled, &landingSectionHeader.IsMobile, &landingSectionHeader.IsDesktop, &landingSectionHeader.Title, &landingSectionHeader.Subtitle, &landingSectionHeader.TagsLine,
		&landingSectionHeader.CreatedAt, &landingSectionHeader.UpdatedAt, &landingSectionHeader.DeletedAt,
	)

	return landingSectionHeader, err
}

func (r landingSectionHeaderRepositoryPostgresql) Delete(ctx context.Context, id int64) (entity.LandingSectionHeader, error) {
	builder := sqlbuilder.New().
		S(`UPDATE "landing_section_headers" SET "deleted_at" = NOW()`).
		S(`WHERE "id" = $1 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "is_enabled", "is_mobile", "is_desktop", "title", "subtitle", "tags_line", "created_at", "updated_at", "deleted_at"`)

	landingSectionHeader := entity.LandingSectionHeader{}
	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingSectionHeader.Id, &landingSectionHeader.IsEnabled, &landingSectionHeader.IsMobile, &landingSectionHeader.IsDesktop, &landingSectionHeader.Title, &landingSectionHeader.Subtitle, &landingSectionHeader.TagsLine,
		&landingSectionHeader.CreatedAt, &landingSectionHeader.UpdatedAt, &landingSectionHeader.DeletedAt,
	)

	return landingSectionHeader, err
}

type landingPackageItemRepositoryPostgresql struct {
	db DB
}

func NewLandingPackageItemRepository(db DB) repository.LandingPackageItemRepository {
	return landingPackageItemRepositoryPostgresql{db}
}

func (r landingPackageItemRepositoryPostgresql) Create(ctx context.Context, landingPackageItem entity.LandingPackageItem) (entity.LandingPackageItem, error) {
	builder := sqlbuilder.New().
		S(`INSERT INTO "landing_package_items" ("is_enabled", "is_mobile", "is_desktop", "package_id", "button_label", "created_at", "updated_at", "deleted_at")`).
		S(`VALUES ($1, $2, $3, $4, $5, NOW(), NOW(), NULL)`, landingPackageItem.IsEnabled, landingPackageItem.IsMobile, landingPackageItem.IsDesktop, landingPackageItem.PackageId, landingPackageItem.ButtonLabel).
		S(`RETURNING "id", "is_enabled", "is_mobile", "is_desktop", "package_id", "button_label", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingPackageItem.Id, &landingPackageItem.IsEnabled, &landingPackageItem.IsMobile, &landingPackageItem.IsDesktop, &landingPackageItem.PackageId, &landingPackageItem.ButtonLabel,
		&landingPackageItem.CreatedAt, &landingPackageItem.UpdatedAt, &landingPackageItem.DeletedAt,
	)

	return landingPackageItem, err
}

func (r landingPackageItemRepositoryPostgresql) FindById(ctx context.Context, id int64) (entity.LandingPackageItem, error) {
	builder := sqlbuilder.New().
		S(`SELECT "id", "is_enabled", "is_mobile", "is_desktop", "package_id", "button_label", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "landing_package_items" WHERE "id" = $1 AND "deleted_at" IS NULL`, id)

	landingPackageItem := entity.LandingPackageItem{}
	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingPackageItem.Id, &landingPackageItem.IsEnabled, &landingPackageItem.IsMobile, &landingPackageItem.IsDesktop, &landingPackageItem.PackageId, &landingPackageItem.ButtonLabel,
		&landingPackageItem.CreatedAt, &landingPackageItem.UpdatedAt, &landingPackageItem.DeletedAt,
	)

	return landingPackageItem, err
}

func (r landingPackageItemRepositoryPostgresql) FindAll(ctx context.Context, opt repository.FindAllOptions) ([]entity.LandingPackageItem, error) {
	builder := sqlbuilder.New().
		S(`SELECT "id", "is_enabled", "is_mobile", "is_desktop", "package_id", "button_label", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "landing_package_items" WHERE "deleted_at" IS NULL`).
		S(`ORDER BY "id" ASC`)
	if opt.Limit.Valid {
		builder.SA(`LIMIT ?`, opt.Limit)
	}
	if opt.Offset.Valid {
		builder.SA(`OFFSET ?`, opt.Offset)
	}

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return nil, err
	}

	landingPackageItems := []entity.LandingPackageItem{}
	for rows.Next() {
		landingPackageItem := entity.LandingPackageItem{}
		err := rows.Scan(
			&landingPackageItem.Id, &landingPackageItem.IsEnabled, &landingPackageItem.IsMobile, &landingPackageItem.IsDesktop, &landingPackageItem.PackageId, &landingPackageItem.ButtonLabel,
			&landingPackageItem.CreatedAt, &landingPackageItem.UpdatedAt, &landingPackageItem.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		landingPackageItems = append(landingPackageItems, landingPackageItem)
	}

	return landingPackageItems, nil
}

func (r landingPackageItemRepositoryPostgresql) Update(ctx context.Context, id int64, landingPackageItem entity.LandingPackageItem) (entity.LandingPackageItem, error) {
	builder := sqlbuilder.New().
		S(`UPDATE "landing_package_items" SET "is_enabled" = $1, "is_mobile" = $2, "is_desktop" = $3, "package_id" = $4, "button_label" = $5, "updated_at" = NOW()`).
		S(`WHERE "id" = $6 AND "deleted_at" IS NULL`, landingPackageItem.IsEnabled, landingPackageItem.IsMobile, landingPackageItem.IsDesktop, landingPackageItem.PackageId, landingPackageItem.ButtonLabel, id).
		S(`RETURNING "id", "is_enabled", "is_mobile", "is_desktop", "package_id", "button_label", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingPackageItem.Id, &landingPackageItem.IsEnabled, &landingPackageItem.IsMobile, &landingPackageItem.IsDesktop, &landingPackageItem.PackageId, &landingPackageItem.ButtonLabel,
		&landingPackageItem.CreatedAt, &landingPackageItem.UpdatedAt, &landingPackageItem.DeletedAt,
	)

	return landingPackageItem, err
}

func (r landingPackageItemRepositoryPostgresql) Delete(ctx context.Context, id int64) (entity.LandingPackageItem, error) {
	builder := sqlbuilder.New().
		S(`UPDATE "landing_package_items" SET "deleted_at" = NOW()`).
		S(`WHERE "id" = $1 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "is_enabled", "is_mobile", "is_desktop", "package_id", "button_label", "created_at", "updated_at", "deleted_at"`)

	landingPackageItem := entity.LandingPackageItem{}
	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingPackageItem.Id, &landingPackageItem.IsEnabled, &landingPackageItem.IsMobile, &landingPackageItem.IsDesktop, &landingPackageItem.PackageId, &landingPackageItem.ButtonLabel,
		&landingPackageItem.CreatedAt, &landingPackageItem.UpdatedAt, &landingPackageItem.DeletedAt,
	)

	return landingPackageItem, err
}

type landingSinglePackageContentRepositoryPostgresql struct {
	db DB
}

func NewLandingSinglePackageContentRepository(db DB) repository.LandingSinglePackageContentRepository {
	return landingSinglePackageContentRepositoryPostgresql{db}
}

func (r landingSinglePackageContentRepositoryPostgresql) Create(ctx context.Context, landingSinglePackageContent entity.LandingSinglePackageContent) (entity.LandingSinglePackageContent, error) {
	builder := sqlbuilder.New().
		S(`INSERT INTO "landing_single_package_content" ("is_enabled", "is_mobile", "is_desktop", "landing_section_header_id", "silver_landing_package_item_id", "gold_landing_package_item_id", "platinum_landing_package_item_id", "created_at", "updated_at", "deleted_at")`).
		S(`VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(), NOW(), NULL)`, landingSinglePackageContent.IsEnabled, landingSinglePackageContent.IsMobile, landingSinglePackageContent.IsDesktop, landingSinglePackageContent.LandingSectionHeaderId, landingSinglePackageContent.SilverLandingPackageItemId, landingSinglePackageContent.GoldLandingPackageItemId, landingSinglePackageContent.PlatinumLandingPackageItemId).
		S(`RETURNING "id", "is_enabled", "is_mobile", "is_desktop", "landing_section_header_id", "silver_landing_package_item_id", "gold_landing_package_item_id", "platinum_landing_package_item_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingSinglePackageContent.Id, &landingSinglePackageContent.IsEnabled, &landingSinglePackageContent.IsMobile, &landingSinglePackageContent.IsDesktop, &landingSinglePackageContent.LandingSectionHeaderId,
		&landingSinglePackageContent.SilverLandingPackageItemId, &landingSinglePackageContent.GoldLandingPackageItemId,
		&landingSinglePackageContent.PlatinumLandingPackageItemId, &landingSinglePackageContent.CreatedAt,
		&landingSinglePackageContent.UpdatedAt, &landingSinglePackageContent.DeletedAt,
	)

	return landingSinglePackageContent, err
}

func (r landingSinglePackageContentRepositoryPostgresql) Find(ctx context.Context) (entity.LandingSinglePackageContent, error) {
	landingSinglePackageContent := entity.LandingSinglePackageContent{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "is_enabled", "is_mobile", "is_desktop", "landing_section_header_id", "silver_landing_package_item_id", "gold_landing_package_item_id", "platinum_landing_package_item_id", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "landing_single_package_content" WHERE "deleted_at" IS NULL`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingSinglePackageContent.Id, &landingSinglePackageContent.IsEnabled, &landingSinglePackageContent.IsMobile, &landingSinglePackageContent.IsDesktop, &landingSinglePackageContent.LandingSectionHeaderId,
		&landingSinglePackageContent.SilverLandingPackageItemId, &landingSinglePackageContent.GoldLandingPackageItemId,
		&landingSinglePackageContent.PlatinumLandingPackageItemId, &landingSinglePackageContent.CreatedAt,
		&landingSinglePackageContent.UpdatedAt, &landingSinglePackageContent.DeletedAt,
	)

	return landingSinglePackageContent, err
}

func (r landingSinglePackageContentRepositoryPostgresql) Update(ctx context.Context, landingSinglePackageContent entity.LandingSinglePackageContent) (entity.LandingSinglePackageContent, error) {
	builder := sqlbuilder.New().
		S(`UPDATE "landing_single_package_content" SET "is_enabled" = $1, "is_mobile" = $2, "is_desktop" = $3, "landing_section_header_id" = $4, "silver_landing_package_item_id" = $5, "gold_landing_package_item_id" = $6, "platinum_landing_package_item_id" = $7, "updated_at" = NOW()`, landingSinglePackageContent.IsEnabled, landingSinglePackageContent.IsMobile, landingSinglePackageContent.IsDesktop, landingSinglePackageContent.LandingSectionHeaderId, landingSinglePackageContent.SilverLandingPackageItemId, landingSinglePackageContent.GoldLandingPackageItemId, landingSinglePackageContent.PlatinumLandingPackageItemId).
		S(`WHERE "deleted_at" IS NULL`).
		S(`RETURNING "id", "is_enabled", "is_mobile", "is_desktop", "landing_section_header_id", "silver_landing_package_item_id", "gold_landing_package_item_id", "platinum_landing_package_item_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingSinglePackageContent.Id, &landingSinglePackageContent.IsEnabled, &landingSinglePackageContent.IsMobile, &landingSinglePackageContent.IsDesktop, &landingSinglePackageContent.LandingSectionHeaderId,
		&landingSinglePackageContent.SilverLandingPackageItemId, &landingSinglePackageContent.GoldLandingPackageItemId,
		&landingSinglePackageContent.PlatinumLandingPackageItemId, &landingSinglePackageContent.CreatedAt,
		&landingSinglePackageContent.UpdatedAt, &landingSinglePackageContent.DeletedAt,
	)

	return landingSinglePackageContent, err
}

func (r landingSinglePackageContentRepositoryPostgresql) Delete(ctx context.Context) (entity.LandingSinglePackageContent, error) {
	landingSinglePackageContent := entity.LandingSinglePackageContent{}

	builder := sqlbuilder.New().
		S(`DELETE FROM "landing_single_package_content"`).
		S(`RETURNING "id", "is_enabled", "is_mobile", "is_desktop", "landing_section_header_id", "silver_landing_package_item_id", "gold_landing_package_item_id", "platinum_landing_package_item_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingSinglePackageContent.Id, &landingSinglePackageContent.IsEnabled, &landingSinglePackageContent.IsMobile, &landingSinglePackageContent.IsDesktop, &landingSinglePackageContent.LandingSectionHeaderId,
		&landingSinglePackageContent.SilverLandingPackageItemId, &landingSinglePackageContent.GoldLandingPackageItemId,
		&landingSinglePackageContent.PlatinumLandingPackageItemId, &landingSinglePackageContent.CreatedAt,
		&landingSinglePackageContent.UpdatedAt, &landingSinglePackageContent.DeletedAt,
	)

	return landingSinglePackageContent, err
}

type landingPackageDetailRepositoryPostgresql struct {
	db DB
}

func NewLandingPackageDetailRepository(db DB) repository.LandingPackageDetailRepository {
	return landingPackageDetailRepositoryPostgresql{db}
}

func (r landingPackageDetailRepositoryPostgresql) Create(ctx context.Context, landingPackageDetail entity.LandingPackageDetail) (entity.LandingPackageDetail, error) {
	builder := sqlbuilder.New().
		S(`INSERT INTO "landing_package_details" ("is_enabled", "is_mobile", "is_desktop", "landing_section_header_id", "created_at", "updated_at", "deleted_at")`).
		S(`VALUES ($1, $2, $3, $4, NOW(), NOW(), NULL)`, landingPackageDetail.IsEnabled, landingPackageDetail.IsMobile, landingPackageDetail.IsDesktop, landingPackageDetail.LandingSectionHeaderId).
		S(`RETURNING "id", "is_enabled", "is_mobile", "is_desktop", "landing_section_header_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingPackageDetail.Id, &landingPackageDetail.IsEnabled, &landingPackageDetail.IsMobile, &landingPackageDetail.IsDesktop, &landingPackageDetail.LandingSectionHeaderId,
		&landingPackageDetail.CreatedAt, &landingPackageDetail.UpdatedAt, &landingPackageDetail.DeletedAt,
	)

	return landingPackageDetail, err
}

func (r landingPackageDetailRepositoryPostgresql) FindById(ctx context.Context, id int64) (entity.LandingPackageDetail, error) {
	builder := sqlbuilder.New().
		S(`SELECT "id", "is_enabled", "is_mobile", "is_desktop", "landing_section_header_id", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "landing_package_details" WHERE "id" = $1 AND "deleted_at" IS NULL`, id)

	landingPackageDetail := entity.LandingPackageDetail{}
	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingPackageDetail.Id, &landingPackageDetail.IsEnabled, &landingPackageDetail.IsMobile, &landingPackageDetail.IsDesktop, &landingPackageDetail.LandingSectionHeaderId,
		&landingPackageDetail.CreatedAt, &landingPackageDetail.UpdatedAt, &landingPackageDetail.DeletedAt,
	)

	return landingPackageDetail, err
}

func (r landingPackageDetailRepositoryPostgresql) FindAll(ctx context.Context, opt repository.FindAllOptions) ([]entity.LandingPackageDetail, error) {
	builder := sqlbuilder.New().
		S(`SELECT "id", "is_enabled", "is_mobile", "is_desktop", "landing_section_header_id", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "landing_package_details" WHERE "deleted_at" IS NULL`).
		S(`ORDER BY "id" ASC`)
	if opt.Limit.Valid {
		builder.SA(`LIMIT ?`, opt.Limit)
	}
	if opt.Offset.Valid {
		builder.SA(`OFFSET ?`, opt.Offset)
	}

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return nil, err
	}

	landingPackageDetails := []entity.LandingPackageDetail{}
	for rows.Next() {
		landingPackageDetail := entity.LandingPackageDetail{}
		err := rows.Scan(
			&landingPackageDetail.Id, &landingPackageDetail.IsEnabled, &landingPackageDetail.IsMobile, &landingPackageDetail.IsDesktop, &landingPackageDetail.LandingSectionHeaderId,
			&landingPackageDetail.CreatedAt, &landingPackageDetail.UpdatedAt, &landingPackageDetail.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		landingPackageDetails = append(landingPackageDetails, landingPackageDetail)
	}

	return landingPackageDetails, nil
}

func (r landingPackageDetailRepositoryPostgresql) Update(ctx context.Context, id int64, landingPackageDetail entity.LandingPackageDetail) (entity.LandingPackageDetail, error) {
	builder := sqlbuilder.New().
		S(`UPDATE "landing_package_details" SET "is_enabled" = $1, "is_mobile" = $2, "is_desktop" = $3, "landing_section_header_id" = $4, "updated_at" = NOW()`).
		S(`WHERE "id" = $5 AND "deleted_at" IS NULL`, landingPackageDetail.IsEnabled, landingPackageDetail.IsMobile, landingPackageDetail.IsDesktop, landingPackageDetail.LandingSectionHeaderId, id).
		S(`RETURNING "id", "is_enabled", "is_mobile", "is_desktop", "landing_section_header_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingPackageDetail.Id, &landingPackageDetail.IsEnabled, &landingPackageDetail.IsMobile, &landingPackageDetail.IsDesktop, &landingPackageDetail.LandingSectionHeaderId,
		&landingPackageDetail.CreatedAt, &landingPackageDetail.UpdatedAt, &landingPackageDetail.DeletedAt,
	)

	return landingPackageDetail, err
}

func (r landingPackageDetailRepositoryPostgresql) Delete(ctx context.Context, id int64) (entity.LandingPackageDetail, error) {
	builder := sqlbuilder.New().
		S(`UPDATE "landing_package_details" SET "deleted_at" = NULL`).
		S(`WHERE "id" = $1 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "is_enabled", "is_mobile", "is_desktop", "landing_section_header_id", "created_at", "updated_at", "deleted_at"`)

	landingPackageDetail := entity.LandingPackageDetail{}
	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingPackageDetail.Id, &landingPackageDetail.IsEnabled, &landingPackageDetail.IsMobile, &landingPackageDetail.IsDesktop, &landingPackageDetail.LandingSectionHeaderId,
		&landingPackageDetail.CreatedAt, &landingPackageDetail.UpdatedAt, &landingPackageDetail.DeletedAt,
	)

	return landingPackageDetail, err
}

type landingPackageDetailItemRepositoryPostgresql struct {
	db DB
}

func NewLandingPackageDetailItemRepository(db DB) repository.LandingPackageDetailItemRepository {
	return landingPackageDetailItemRepositoryPostgresql{db}
}

func (r landingPackageDetailItemRepositoryPostgresql) CreateMany(ctx context.Context, landingPackageDetailItems []entity.LandingPackageDetailItem) ([]entity.LandingPackageDetailItem, error) {
	if len(landingPackageDetailItems) == 0 {
		return []entity.LandingPackageDetailItem{}, nil
	}

	builder := sqlbuilder.New().
		S(`INSERT INTO "landing_package_detail_items" ("is_enabled", "is_mobile", "is_desktop", "landing_package_detail_id", "landing_package_item_id", "created_at", "updated_at") VALUES`)

	for i, item := range landingPackageDetailItems {
		builder.SA(`(?, ?, ?, ?, ?, NOW(), NOW())`, item.IsEnabled, item.IsMobile, item.IsDesktop, item.LandingPackageDetailId, item.LandingPackageItemId)
		if i+1 < len(landingPackageDetailItems) {
			builder.SA(",")
		}
	}

	builder.S(`RETURNING "is_enabled", "is_mobile", "is_desktop", "landing_package_detail_id", "landing_package_item_id", "created_at", "updated_at", "deleted_at"`)

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	createdItems := []entity.LandingPackageDetailItem{}
	for rows.Next() {
		item := entity.LandingPackageDetailItem{}
		err := rows.Scan(&item.IsEnabled, &item.IsMobile, &item.IsDesktop, &item.LandingPackageDetailId, &item.LandingPackageItemId, &item.CreatedAt, &item.UpdatedAt, &item.DeletedAt)
		if err != nil {
			return nil, err
		}
		createdItems = append(createdItems, item)
	}

	return createdItems, nil
}

func (r landingPackageDetailItemRepositoryPostgresql) FindAll(ctx context.Context, opt repository.FindAllOptions) ([]entity.LandingPackageDetailItem, error) {
	builder := sqlbuilder.New().
		S(`SELECT "is_enabled", "is_mobile", "is_desktop", "landing_package_detail_id", "landing_package_item_id", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "landing_package_detail_items" WHERE "deleted_at" IS NULL`)
	if landingPackageDetailId, ok := opt.Where["landing_package_detail_id"]; ok {
		builder.SA(`AND "landing_package_detail_id" = ?`, landingPackageDetailId)
	}
	if opt.Limit.Valid {
		builder.SA(`LIMIT ?`, opt.Limit)
	}
	if opt.Offset.Valid {
		builder.SA(`OFFSET ?`, opt.Offset)
	}

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []entity.LandingPackageDetailItem{}
	for rows.Next() {
		item := entity.LandingPackageDetailItem{}
		err := rows.Scan(&item.IsEnabled, &item.IsMobile, &item.IsDesktop, &item.LandingPackageDetailId, &item.LandingPackageItemId, &item.CreatedAt, &item.UpdatedAt, &item.DeletedAt)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func (r landingPackageDetailItemRepositoryPostgresql) DeleteMany(ctx context.Context) ([]entity.LandingPackageDetailItem, error) {
	builder := sqlbuilder.New().
		S(`DELETE FROM "landing_package_detail_items" WHERE "deleted_at" IS NULL`).
		S(`RETURNING "is_enabled", "is_mobile", "is_desktop", "landing_package_detail_id", "landing_package_item_id", "created_at", "updated_at", "deleted_at"`)

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	deletedItems := []entity.LandingPackageDetailItem{}
	for rows.Next() {
		item := entity.LandingPackageDetailItem{}
		err := rows.Scan(&item.IsEnabled, &item.IsMobile, &item.IsDesktop, &item.LandingPackageDetailId, &item.LandingPackageItemId, &item.CreatedAt, &item.UpdatedAt, &item.DeletedAt)
		if err != nil {
			return nil, err
		}
		deletedItems = append(deletedItems, item)
	}

	return deletedItems, nil
}

type landingPackagesContentRepositoryPostgresql struct {
	db DB
}

func NewLandingPackagesContentRepository(db DB) repository.LandingPackagesContentRepository {
	return landingPackagesContentRepositoryPostgresql{db}
}

func (r landingPackagesContentRepositoryPostgresql) Create(ctx context.Context, landingPackagesContent entity.LandingPackagesContent) (entity.LandingPackagesContent, error) {
	builder := sqlbuilder.New().
		S(`INSERT INTO "landing_packages_content" ("is_enabled", "is_mobile", "is_desktop", "silver_landing_package_detail_id", "gold_landing_package_detail_id", "platinum_landing_package_detail_id", "created_at", "updated_at", "deleted_at")`).
		S(`VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW(), NULL)`, landingPackagesContent.IsEnabled, landingPackagesContent.IsMobile, landingPackagesContent.IsDesktop, landingPackagesContent.SilverLandingPackageDetailId, landingPackagesContent.GoldLandingPackageDetailId, landingPackagesContent.PlatinumLandingPackageDetailId).
		S(`RETURNING "id", "is_enabled", "is_mobile", "is_desktop", "silver_landing_package_detail_id", "gold_landing_package_detail_id", "platinum_landing_package_detail_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingPackagesContent.Id, &landingPackagesContent.IsEnabled, &landingPackagesContent.IsMobile, &landingPackagesContent.IsDesktop, &landingPackagesContent.SilverLandingPackageDetailId,
		&landingPackagesContent.GoldLandingPackageDetailId, &landingPackagesContent.PlatinumLandingPackageDetailId,
		&landingPackagesContent.CreatedAt, &landingPackagesContent.UpdatedAt, &landingPackagesContent.DeletedAt,
	)

	return landingPackagesContent, err
}

func (r landingPackagesContentRepositoryPostgresql) Find(ctx context.Context) (entity.LandingPackagesContent, error) {
	landingPackagesContent := entity.LandingPackagesContent{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "is_enabled", "is_mobile", "is_desktop", "silver_landing_package_detail_id", "gold_landing_package_detail_id", "platinum_landing_package_detail_id", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "landing_packages_content" WHERE "deleted_at" IS NULL`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingPackagesContent.Id, &landingPackagesContent.IsEnabled, &landingPackagesContent.IsMobile, &landingPackagesContent.IsDesktop, &landingPackagesContent.SilverLandingPackageDetailId,
		&landingPackagesContent.GoldLandingPackageDetailId, &landingPackagesContent.PlatinumLandingPackageDetailId,
		&landingPackagesContent.CreatedAt, &landingPackagesContent.UpdatedAt, &landingPackagesContent.DeletedAt,
	)

	return landingPackagesContent, err
}

func (r landingPackagesContentRepositoryPostgresql) Update(ctx context.Context, landingPackagesContent entity.LandingPackagesContent) (entity.LandingPackagesContent, error) {
	builder := sqlbuilder.New().
		S(`UPDATE "landing_packages_content" SET "is_enabled" = $1, "is_mobile" = $2, "is_desktop" = $3, "silver_landing_package_detail_id" = $4, "gold_landing_package_detail_id" = $5, "platinum_landing_package_detail_id" = $6, "updated_at" = NOW()`, landingPackagesContent.IsEnabled, landingPackagesContent.IsMobile, landingPackagesContent.IsDesktop, landingPackagesContent.SilverLandingPackageDetailId, landingPackagesContent.GoldLandingPackageDetailId, landingPackagesContent.PlatinumLandingPackageDetailId).
		S(`WHERE "deleted_at" IS NULL`).
		S(`RETURNING "id", "is_enabled", "is_mobile", "is_desktop", "silver_landing_package_detail_id", "gold_landing_package_detail_id", "platinum_landing_package_detail_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingPackagesContent.Id, &landingPackagesContent.IsEnabled, &landingPackagesContent.IsMobile, &landingPackagesContent.IsDesktop, &landingPackagesContent.SilverLandingPackageDetailId,
		&landingPackagesContent.GoldLandingPackageDetailId, &landingPackagesContent.PlatinumLandingPackageDetailId,
		&landingPackagesContent.CreatedAt, &landingPackagesContent.UpdatedAt, &landingPackagesContent.DeletedAt,
	)

	return landingPackagesContent, err
}

func (r landingPackagesContentRepositoryPostgresql) Delete(ctx context.Context) (entity.LandingPackagesContent, error) {
	landingPackagesContent := entity.LandingPackagesContent{}

	builder := sqlbuilder.New().
		S(`DELETE FROM "landing_packages_content"`).
		S(`RETURNING "id", "is_enabled", "is_mobile", "is_desktop", "silver_landing_package_detail_id", "gold_landing_package_detail_id", "platinum_landing_package_detail_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingPackagesContent.Id, &landingPackagesContent.IsEnabled, &landingPackagesContent.IsMobile, &landingPackagesContent.IsDesktop, &landingPackagesContent.SilverLandingPackageDetailId,
		&landingPackagesContent.GoldLandingPackageDetailId, &landingPackagesContent.PlatinumLandingPackageDetailId,
		&landingPackagesContent.CreatedAt, &landingPackagesContent.UpdatedAt, &landingPackagesContent.DeletedAt,
	)

	return landingPackagesContent, err
}

type landingTravelDestinationContentRepositoryPostgresql struct {
	db DB
}

func NewLandingTravelDestinationContentRepository(db DB) repository.LandingTravelDestinationContentRepository {
	return landingTravelDestinationContentRepositoryPostgresql{db}
}

func (r landingTravelDestinationContentRepositoryPostgresql) Create(ctx context.Context, landingTravelDestinationContent entity.LandingTravelDestinationContent) (entity.LandingTravelDestinationContent, error) {
	builder := sqlbuilder.New().
		S(`INSERT INTO "landing_travel_destination_content" ("is_enabled", "is_mobile", "is_desktop", "landing_section_header_id", "created_at", "updated_at", "deleted_at")`).
		S(`VALUES ($1, $2, $3, $4, NOW(), NOW(), NULL)`, landingTravelDestinationContent.IsEnabled, landingTravelDestinationContent.IsMobile, landingTravelDestinationContent.IsDesktop, landingTravelDestinationContent.LandingSectionHeaderId).
		S(`RETURNING "id", "is_enabled", "is_mobile", "is_desktop", "landing_section_header_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingTravelDestinationContent.Id, &landingTravelDestinationContent.IsEnabled, &landingTravelDestinationContent.IsMobile, &landingTravelDestinationContent.IsDesktop, &landingTravelDestinationContent.LandingSectionHeaderId,
		&landingTravelDestinationContent.CreatedAt, &landingTravelDestinationContent.UpdatedAt, &landingTravelDestinationContent.DeletedAt,
	)

	return landingTravelDestinationContent, err
}

func (r landingTravelDestinationContentRepositoryPostgresql) Find(ctx context.Context) (entity.LandingTravelDestinationContent, error) {
	landingTravelDestinationContent := entity.LandingTravelDestinationContent{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "is_enabled", "is_mobile", "is_desktop", "landing_section_header_id", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "landing_travel_destination_content" WHERE "deleted_at" IS NULL`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingTravelDestinationContent.Id, &landingTravelDestinationContent.IsEnabled, &landingTravelDestinationContent.IsMobile, &landingTravelDestinationContent.IsDesktop, &landingTravelDestinationContent.LandingSectionHeaderId,
		&landingTravelDestinationContent.CreatedAt, &landingTravelDestinationContent.UpdatedAt, &landingTravelDestinationContent.DeletedAt,
	)

	return landingTravelDestinationContent, err
}

func (r landingTravelDestinationContentRepositoryPostgresql) Update(ctx context.Context, landingTravelDestinationContent entity.LandingTravelDestinationContent) (entity.LandingTravelDestinationContent, error) {
	builder := sqlbuilder.New().
		S(`UPDATE "landing_travel_destination_content" SET "is_enabled" = $1, "is_mobile" = $2, "is_desktop" = $3, "landing_section_header_id" = $4, "updated_at" = NOW()`, landingTravelDestinationContent.IsEnabled, landingTravelDestinationContent.IsMobile, landingTravelDestinationContent.IsDesktop, landingTravelDestinationContent.LandingSectionHeaderId).
		S(`WHERE "deleted_at" IS NULL`).
		S(`RETURNING "id", "is_enabled", "is_mobile", "is_desktop", "landing_section_header_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingTravelDestinationContent.Id, &landingTravelDestinationContent.IsEnabled, &landingTravelDestinationContent.IsMobile, &landingTravelDestinationContent.IsDesktop, &landingTravelDestinationContent.LandingSectionHeaderId,
		&landingTravelDestinationContent.CreatedAt, &landingTravelDestinationContent.UpdatedAt, &landingTravelDestinationContent.DeletedAt,
	)

	return landingTravelDestinationContent, err
}

func (r landingTravelDestinationContentRepositoryPostgresql) Delete(ctx context.Context) (entity.LandingTravelDestinationContent, error) {
	landingTravelDestinationContent := entity.LandingTravelDestinationContent{}

	builder := sqlbuilder.New().
		S(`DELETE FROM "landing_travel_destination_content"`).
		S(`RETURNING "id", "is_enabled", "is_mobile", "is_desktop", "landing_section_header_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingTravelDestinationContent.Id, &landingTravelDestinationContent.IsEnabled, &landingTravelDestinationContent.IsMobile, &landingTravelDestinationContent.IsDesktop, &landingTravelDestinationContent.LandingSectionHeaderId,
		&landingTravelDestinationContent.CreatedAt, &landingTravelDestinationContent.UpdatedAt, &landingTravelDestinationContent.DeletedAt,
	)

	return landingTravelDestinationContent, err
}

type landingTravelDestinationContentDestinationPostgresql struct {
	db DB
}

func NewLandingTravelDestinationContentDestinationRepository(db DB) repository.LandingTravelDestinationContentDestinationRepository {
	return landingTravelDestinationContentDestinationPostgresql{db}
}

func (r landingTravelDestinationContentDestinationPostgresql) CreateMany(ctx context.Context, landingTravelDestinationContentDestinations []entity.LandingTravelDestinationContentDestination) ([]entity.LandingTravelDestinationContentDestination, error) {
	if len(landingTravelDestinationContentDestinations) == 0 {
		return []entity.LandingTravelDestinationContentDestination{}, nil
	}

	builder := sqlbuilder.New().
		S(`INSERT INTO "landing_travel_destination_content_destinations" ("is_enabled", "is_mobile", "is_desktop", "image_id", "name", "created_at", "updated_at") VALUES`)
	for i, item := range landingTravelDestinationContentDestinations {
		builder.SA(`(?, ?, ?, ?, ?, NOW(), NOW())`, item.IsEnabled, item.IsMobile, item.IsDesktop, item.ImageId, item.Name)

		if i+1 < len(landingTravelDestinationContentDestinations) {
			builder.SA(",")
		}
	}
	builder.S(`RETURNING "is_enabled", "is_mobile", "is_desktop", "image_id", "name", "created_at", "updated_at", "deleted_at"`)

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	landingTravelDestinationContentDestinations = []entity.LandingTravelDestinationContentDestination{}
	for rows.Next() {
		landingTravelDestinationContentDestination := entity.LandingTravelDestinationContentDestination{}

		err := rows.Scan(
			&landingTravelDestinationContentDestination.IsEnabled, &landingTravelDestinationContentDestination.IsMobile, &landingTravelDestinationContentDestination.IsDesktop,
			&landingTravelDestinationContentDestination.ImageId, &landingTravelDestinationContentDestination.Name,
			&landingTravelDestinationContentDestination.CreatedAt, &landingTravelDestinationContentDestination.UpdatedAt, &landingTravelDestinationContentDestination.DeletedAt,
		)
		if err != nil {
			return nil, err
		}

		landingTravelDestinationContentDestinations = append(landingTravelDestinationContentDestinations, landingTravelDestinationContentDestination)
	}

	return landingTravelDestinationContentDestinations, nil
}

func (r landingTravelDestinationContentDestinationPostgresql) FindAll(ctx context.Context, opt repository.FindAllOptions) ([]entity.LandingTravelDestinationContentDestination, error) {
	builder := sqlbuilder.New().
		S(`SELECT "is_enabled", "is_mobile", "is_desktop", "image_id", "name", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "landing_travel_destination_content_destinations" WHERE "deleted_at" IS NULL`)
	if opt.Limit.Valid {
		builder.SA(`LIMIT ?`, opt.Limit)
	}
	if opt.Offset.Valid {
		builder.SA(`OFFSET ?`, opt.Offset)
	}

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	landingTravelDestinationContentDestinations := []entity.LandingTravelDestinationContentDestination{}
	for rows.Next() {
		landingTravelDestinationContentDestination := entity.LandingTravelDestinationContentDestination{}

		err := rows.Scan(
			&landingTravelDestinationContentDestination.IsEnabled, &landingTravelDestinationContentDestination.IsMobile, &landingTravelDestinationContentDestination.IsDesktop,
			&landingTravelDestinationContentDestination.ImageId, &landingTravelDestinationContentDestination.Name,
			&landingTravelDestinationContentDestination.CreatedAt, &landingTravelDestinationContentDestination.UpdatedAt, &landingTravelDestinationContentDestination.DeletedAt,
		)
		if err != nil {
			return nil, err
		}

		landingTravelDestinationContentDestinations = append(landingTravelDestinationContentDestinations, landingTravelDestinationContentDestination)
	}

	return landingTravelDestinationContentDestinations, nil
}

func (r landingTravelDestinationContentDestinationPostgresql) DeleteMany(ctx context.Context) ([]entity.LandingTravelDestinationContentDestination, error) {
	builder := sqlbuilder.New().
		S(`DELETE FROM "landing_travel_destination_content_destinations" WHERE "deleted_at" IS NULL`).
		S(`RETURNING "is_enabled", "is_mobile", "is_desktop", "image_id", "name", "created_at", "updated_at", "deleted_at"`)

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	landingTravelDestinationContentDestinations := []entity.LandingTravelDestinationContentDestination{}
	for rows.Next() {
		landingTravelDestinationContentDestination := entity.LandingTravelDestinationContentDestination{}

		err := rows.Scan(
			&landingTravelDestinationContentDestination.IsEnabled, &landingTravelDestinationContentDestination.IsMobile, &landingTravelDestinationContentDestination.IsDesktop,
			&landingTravelDestinationContentDestination.ImageId, &landingTravelDestinationContentDestination.Name,
			&landingTravelDestinationContentDestination.CreatedAt, &landingTravelDestinationContentDestination.UpdatedAt, &landingTravelDestinationContentDestination.DeletedAt,
		)
		if err != nil {
			return nil, err
		}

		landingTravelDestinationContentDestinations = append(landingTravelDestinationContentDestinations, landingTravelDestinationContentDestination)
	}

	return landingTravelDestinationContentDestinations, nil
}

type landingFeaturesContentRepositoryPostgresql struct {
	db DB
}

func NewLandingFeaturesContentRepository(db DB) repository.LandingFeaturesContentRepository {
	return landingFeaturesContentRepositoryPostgresql{db}
}

func (r landingFeaturesContentRepositoryPostgresql) Create(ctx context.Context, landingFeaturesContent entity.LandingFeaturesContent) (entity.LandingFeaturesContent, error) {
	builder := sqlbuilder.New().
		S(`INSERT INTO "landing_features_content" ("is_enabled", "is_mobile", "is_desktop", "landing_section_header_id", "footer_title", "button_about", "button_package", "created_at", "updated_at", "deleted_at")`).
		S(`VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(), NOW(), NULL)`, landingFeaturesContent.IsEnabled, landingFeaturesContent.IsMobile, landingFeaturesContent.IsDesktop, landingFeaturesContent.LandingSectionHeaderId, landingFeaturesContent.FooterTitle, landingFeaturesContent.ButtonAbout, landingFeaturesContent.ButtonPackage).
		S(`RETURNING "id", "is_enabled", "is_mobile", "is_desktop", "landing_section_header_id", "footer_title", "button_about", "button_package", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingFeaturesContent.Id, &landingFeaturesContent.IsEnabled, &landingFeaturesContent.IsMobile, &landingFeaturesContent.IsDesktop, &landingFeaturesContent.LandingSectionHeaderId,
		&landingFeaturesContent.FooterTitle, &landingFeaturesContent.ButtonAbout, &landingFeaturesContent.ButtonPackage,
		&landingFeaturesContent.CreatedAt, &landingFeaturesContent.UpdatedAt, &landingFeaturesContent.DeletedAt,
	)

	return landingFeaturesContent, err
}

func (r landingFeaturesContentRepositoryPostgresql) Find(ctx context.Context) (entity.LandingFeaturesContent, error) {
	landingFeaturesContent := entity.LandingFeaturesContent{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "is_enabled", "is_mobile", "is_desktop", "landing_section_header_id", "footer_title", "button_about", "button_package", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "landing_features_content" WHERE "deleted_at" IS NULL`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingFeaturesContent.Id, &landingFeaturesContent.IsEnabled, &landingFeaturesContent.IsMobile, &landingFeaturesContent.IsDesktop, &landingFeaturesContent.LandingSectionHeaderId,
		&landingFeaturesContent.FooterTitle, &landingFeaturesContent.ButtonAbout, &landingFeaturesContent.ButtonPackage,
		&landingFeaturesContent.CreatedAt, &landingFeaturesContent.UpdatedAt, &landingFeaturesContent.DeletedAt,
	)

	return landingFeaturesContent, err
}

func (r landingFeaturesContentRepositoryPostgresql) Update(ctx context.Context, landingFeaturesContent entity.LandingFeaturesContent) (entity.LandingFeaturesContent, error) {
	builder := sqlbuilder.New().
		S(`UPDATE "landing_features_content" SET "is_enabled" = $1, "is_mobile" = $2, "is_desktop" = $3, "landing_section_header_id" = $4, "footer_title" = $5, "button_about" = $6, "button_package" = $7, "updated_at" = NOW()`, landingFeaturesContent.IsEnabled, landingFeaturesContent.IsMobile, landingFeaturesContent.IsDesktop, landingFeaturesContent.LandingSectionHeaderId, landingFeaturesContent.FooterTitle, landingFeaturesContent.ButtonAbout, landingFeaturesContent.ButtonPackage).
		S(`WHERE "deleted_at" IS NULL`).
		S(`RETURNING "id", "is_enabled", "is_mobile", "is_desktop", "landing_section_header_id", "footer_title", "button_about", "button_package", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingFeaturesContent.Id, &landingFeaturesContent.IsEnabled, &landingFeaturesContent.IsMobile, &landingFeaturesContent.IsDesktop, &landingFeaturesContent.LandingSectionHeaderId,
		&landingFeaturesContent.FooterTitle, &landingFeaturesContent.ButtonAbout, &landingFeaturesContent.ButtonPackage,
		&landingFeaturesContent.CreatedAt, &landingFeaturesContent.UpdatedAt, &landingFeaturesContent.DeletedAt,
	)

	return landingFeaturesContent, err
}

func (r landingFeaturesContentRepositoryPostgresql) Delete(ctx context.Context) (entity.LandingFeaturesContent, error) {
	landingFeaturesContent := entity.LandingFeaturesContent{}

	builder := sqlbuilder.New().
		S(`DELETE FROM "landing_features_content"`).
		S(`RETURNING "id", "is_enabled", "is_mobile", "is_desktop", "landing_section_header_id", "footer_title", "button_about", "button_package", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingFeaturesContent.Id, &landingFeaturesContent.IsEnabled, &landingFeaturesContent.IsMobile, &landingFeaturesContent.IsDesktop, &landingFeaturesContent.LandingSectionHeaderId,
		&landingFeaturesContent.FooterTitle, &landingFeaturesContent.ButtonAbout, &landingFeaturesContent.ButtonPackage,
		&landingFeaturesContent.CreatedAt, &landingFeaturesContent.UpdatedAt, &landingFeaturesContent.DeletedAt,
	)

	return landingFeaturesContent, err
}

type landingFeaturesContentBenefitRepositoryPostgresql struct {
	db DB
}

func NewLandingFeaturesContentBenefitRepository(db DB) repository.LandingFeaturesContentBenefitRepository {
	return landingFeaturesContentBenefitRepositoryPostgresql{db}
}

func (r landingFeaturesContentBenefitRepositoryPostgresql) CreateMany(ctx context.Context, landingFeaturesContentBenefits []entity.LandingFeaturesContentBenefit) ([]entity.LandingFeaturesContentBenefit, error) {
	if len(landingFeaturesContentBenefits) == 0 {
		return []entity.LandingFeaturesContentBenefit{}, nil
	}

	builder := sqlbuilder.New().
		S(`INSERT INTO "landing_features_content_benefits" ("is_enabled", "is_mobile", "is_desktop", "title", "subtitle", "logo_id", "created_at", "updated_at") VALUES`)

	for i, benefit := range landingFeaturesContentBenefits {
		builder.SA(`(?, ?, ?, ?, ?, ?, NOW(), NOW())`, benefit.IsEnabled, benefit.IsMobile, benefit.IsDesktop, benefit.Title, benefit.Subtitle, benefit.LogoId)
		if i+1 < len(landingFeaturesContentBenefits) {
			builder.SA(",")
		}
	}

	builder.S(`RETURNING "id", "is_enabled", "is_mobile", "is_desktop", "title", "subtitle", "logo_id", "created_at", "updated_at", "deleted_at"`)

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	createdBenefits := []entity.LandingFeaturesContentBenefit{}
	for rows.Next() {
		benefit := entity.LandingFeaturesContentBenefit{}
		err := rows.Scan(
			&benefit.Id, &benefit.IsEnabled, &benefit.IsMobile, &benefit.IsDesktop, &benefit.Title,
			&benefit.Subtitle, &benefit.LogoId,
			&benefit.CreatedAt, &benefit.UpdatedAt, &benefit.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		createdBenefits = append(createdBenefits, benefit)
	}

	return createdBenefits, nil
}

func (r landingFeaturesContentBenefitRepositoryPostgresql) FindAll(ctx context.Context, opt repository.FindAllOptions) ([]entity.LandingFeaturesContentBenefit, error) {
	builder := sqlbuilder.New().
		S(`SELECT "id", "is_enabled", "is_mobile", "is_desktop", "title", "subtitle", "logo_id", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "landing_features_content_benefits" WHERE "deleted_at" IS NULL`).
		S(`ORDER BY "id" ASC`)
	if opt.Limit.Valid {
		builder.SA(`LIMIT ?`, opt.Limit)
	}
	if opt.Offset.Valid {
		builder.SA(`OFFSET ?`, opt.Offset)
	}

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	benefits := []entity.LandingFeaturesContentBenefit{}
	for rows.Next() {
		benefit := entity.LandingFeaturesContentBenefit{}
		err := rows.Scan(
			&benefit.Id, &benefit.IsEnabled, &benefit.IsMobile, &benefit.IsDesktop, &benefit.Title, &benefit.Subtitle, &benefit.LogoId,
			&benefit.CreatedAt, &benefit.UpdatedAt, &benefit.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		benefits = append(benefits, benefit)
	}

	return benefits, nil
}

func (r landingFeaturesContentBenefitRepositoryPostgresql) DeleteMany(ctx context.Context) ([]entity.LandingFeaturesContentBenefit, error) {
	builder := sqlbuilder.New().
		S(`DELETE FROM "landing_features_content_benefits"`).
		S(`RETURNING "id", "is_enabled", "is_mobile", "is_desktop", "title", "subtitle", "logo_id", "created_at", "updated_at", "deleted_at"`)

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	deletedBenefits := []entity.LandingFeaturesContentBenefit{}
	for rows.Next() {
		benefit := entity.LandingFeaturesContentBenefit{}
		err := rows.Scan(
			&benefit.Id, &benefit.IsEnabled, &benefit.IsMobile, &benefit.IsDesktop, &benefit.Title, &benefit.Subtitle, &benefit.LogoId,
			&benefit.CreatedAt, &benefit.UpdatedAt, &benefit.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		deletedBenefits = append(deletedBenefits, benefit)
	}

	return deletedBenefits, nil
}

type landingMomentsContentRepositoryPostgresql struct {
	db DB
}

func NewLandingMomentsContentRepository(db DB) repository.LandingMomentsContentRepository {
	return landingMomentsContentRepositoryPostgresql{db}
}

func (r landingMomentsContentRepositoryPostgresql) Create(ctx context.Context, landingMomentsContent entity.LandingMomentsContent) (entity.LandingMomentsContent, error) {
	builder := sqlbuilder.New().
		S(`INSERT INTO "landing_moments_content" ("is_enabled", "is_mobile", "is_desktop", "landing_section_header_id", "created_at", "updated_at", "deleted_at")`).
		S(`VALUES ($1, $2, $3, $4, NOW(), NOW(), NULL)`, landingMomentsContent.IsEnabled, landingMomentsContent.IsMobile, landingMomentsContent.IsDesktop, landingMomentsContent.LandingSectionHeaderId).
		S(`RETURNING "id", "is_enabled", "is_mobile", "is_desktop", "landing_section_header_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingMomentsContent.Id, &landingMomentsContent.IsEnabled, &landingMomentsContent.IsMobile, &landingMomentsContent.IsDesktop, &landingMomentsContent.LandingSectionHeaderId,
		&landingMomentsContent.CreatedAt, &landingMomentsContent.UpdatedAt, &landingMomentsContent.DeletedAt,
	)

	return landingMomentsContent, err
}

func (r landingMomentsContentRepositoryPostgresql) Find(ctx context.Context) (entity.LandingMomentsContent, error) {
	landingMomentsContent := entity.LandingMomentsContent{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "is_enabled", "is_mobile", "is_desktop", "landing_section_header_id", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "landing_moments_content" WHERE "deleted_at" IS NULL`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingMomentsContent.Id, &landingMomentsContent.IsEnabled, &landingMomentsContent.IsMobile, &landingMomentsContent.IsDesktop, &landingMomentsContent.LandingSectionHeaderId,
		&landingMomentsContent.CreatedAt, &landingMomentsContent.UpdatedAt, &landingMomentsContent.DeletedAt,
	)

	return landingMomentsContent, err
}

func (r landingMomentsContentRepositoryPostgresql) Update(ctx context.Context, landingMomentsContent entity.LandingMomentsContent) (entity.LandingMomentsContent, error) {
	builder := sqlbuilder.New().
		S(`UPDATE "landing_moments_content" SET "is_enabled" = $1, "is_mobile" = $2, "is_desktop" = $3, "landing_section_header_id" = $4, "updated_at" = NOW()`, landingMomentsContent.IsEnabled, landingMomentsContent.IsMobile, landingMomentsContent.IsDesktop, landingMomentsContent.LandingSectionHeaderId).
		S(`WHERE "deleted_at" IS NULL`).
		S(`RETURNING "id", "is_enabled", "is_mobile", "is_desktop", "landing_section_header_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingMomentsContent.Id, &landingMomentsContent.IsEnabled, &landingMomentsContent.IsMobile, &landingMomentsContent.IsDesktop, &landingMomentsContent.LandingSectionHeaderId,
		&landingMomentsContent.CreatedAt, &landingMomentsContent.UpdatedAt, &landingMomentsContent.DeletedAt,
	)

	return landingMomentsContent, err
}

func (r landingMomentsContentRepositoryPostgresql) Delete(ctx context.Context) (entity.LandingMomentsContent, error) {
	landingMomentsContent := entity.LandingMomentsContent{}

	builder := sqlbuilder.New().
		S(`DELETE FROM "landing_moments_content"`).
		S(`RETURNING "id", "is_enabled", "is_mobile", "is_desktop", "landing_section_header_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingMomentsContent.Id, &landingMomentsContent.IsEnabled, &landingMomentsContent.IsMobile, &landingMomentsContent.IsDesktop, &landingMomentsContent.LandingSectionHeaderId,
		&landingMomentsContent.CreatedAt, &landingMomentsContent.UpdatedAt, &landingMomentsContent.DeletedAt,
	)

	return landingMomentsContent, err
}

type landingMomentsContentImageRepositoryPostgresql struct {
	db DB
}

func NewLandingMomentsContentImageRepository(db DB) repository.LandingMomentsContentImageRepository {
	return landingMomentsContentImageRepositoryPostgresql{db}
}

func (r landingMomentsContentImageRepositoryPostgresql) CreateMany(ctx context.Context, landingMomentsContentImages []entity.LandingMomentsContentImage) ([]entity.LandingMomentsContentImage, error) {
	if len(landingMomentsContentImages) == 0 {
		return []entity.LandingMomentsContentImage{}, nil
	}

	builder := sqlbuilder.New().
		S(`INSERT INTO "landing_moments_content_images" ("is_enabled", "is_mobile", "is_desktop", "image_id", "created_at", "updated_at") VALUES`)

	for i, image := range landingMomentsContentImages {
		builder.SA(`(?, ?, ?, ?, NOW(), NOW())`, image.IsEnabled, image.IsMobile, image.IsDesktop, image.ImageId)
		if i+1 < len(landingMomentsContentImages) {
			builder.SA(",")
		}
	}

	builder.S(`RETURNING "is_enabled", "is_mobile", "is_desktop", "image_id", "created_at", "updated_at", "deleted_at"`)

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	createdImages := []entity.LandingMomentsContentImage{}
	for rows.Next() {
		image := entity.LandingMomentsContentImage{}
		err := rows.Scan(&image.IsEnabled, &image.IsMobile, &image.IsDesktop, &image.ImageId, &image.CreatedAt, &image.UpdatedAt, &image.DeletedAt)
		if err != nil {
			return nil, err
		}
		createdImages = append(createdImages, image)
	}

	return createdImages, nil
}

func (r landingMomentsContentImageRepositoryPostgresql) FindAll(ctx context.Context, opt repository.FindAllOptions) ([]entity.LandingMomentsContentImage, error) {
	builder := sqlbuilder.New().
		S(`SELECT "is_enabled", "is_mobile", "is_desktop", "image_id", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "landing_moments_content_images" WHERE "deleted_at" IS NULL`).
		S(`ORDER BY "image_id" ASC`)
	if opt.Limit.Valid {
		builder.SA(`LIMIT ?`, opt.Limit)
	}
	if opt.Offset.Valid {
		builder.SA(`OFFSET ?`, opt.Offset)
	}

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	images := []entity.LandingMomentsContentImage{}
	for rows.Next() {
		image := entity.LandingMomentsContentImage{}
		err := rows.Scan(&image.IsEnabled, &image.IsMobile, &image.IsDesktop, &image.ImageId, &image.CreatedAt, &image.UpdatedAt, &image.DeletedAt)
		if err != nil {
			return nil, err
		}
		images = append(images, image)
	}

	return images, nil
}

func (r landingMomentsContentImageRepositoryPostgresql) DeleteMany(ctx context.Context) ([]entity.LandingMomentsContentImage, error) {
	builder := sqlbuilder.New().
		S(`DELETE FROM "landing_moments_content_images"`).
		S(`RETURNING "is_enabled", "is_mobile", "is_desktop", "image_id", "created_at", "updated_at", "deleted_at"`)

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	deletedImages := []entity.LandingMomentsContentImage{}
	for rows.Next() {
		image := entity.LandingMomentsContentImage{}
		err := rows.Scan(&image.IsEnabled, &image.IsMobile, &image.IsDesktop, &image.ImageId, &image.CreatedAt, &image.UpdatedAt, &image.DeletedAt)
		if err != nil {
			return nil, err
		}
		deletedImages = append(deletedImages, image)
	}

	return deletedImages, nil
}

type landingAffiliatesContentRepositoryPostgresql struct {
	db DB
}

func NewLandingAffiliatesContentRepository(db DB) repository.LandingAffiliatesContentRepository {
	return landingAffiliatesContentRepositoryPostgresql{db}
}

func (r landingAffiliatesContentRepositoryPostgresql) Create(ctx context.Context, landingAffiliatesContent entity.LandingAffiliatesContent) (entity.LandingAffiliatesContent, error) {
	builder := sqlbuilder.New().
		S(`INSERT INTO "landing_affiliates_content" ("is_enabled", "is_mobile", "is_desktop", "landing_section_header_id", "created_at", "updated_at", "deleted_at")`).
		S(`VALUES ($1, $2, $3, $4, NOW(), NOW(), NULL)`, landingAffiliatesContent.IsEnabled, landingAffiliatesContent.IsMobile, landingAffiliatesContent.IsDesktop, landingAffiliatesContent.LandingSectionHeaderId).
		S(`RETURNING "id", "is_enabled", "is_mobile", "is_desktop", "landing_section_header_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingAffiliatesContent.Id, &landingAffiliatesContent.IsEnabled, &landingAffiliatesContent.IsMobile, &landingAffiliatesContent.IsDesktop, &landingAffiliatesContent.LandingSectionHeaderId,
		&landingAffiliatesContent.CreatedAt, &landingAffiliatesContent.UpdatedAt, &landingAffiliatesContent.DeletedAt,
	)

	return landingAffiliatesContent, err
}

func (r landingAffiliatesContentRepositoryPostgresql) Find(ctx context.Context) (entity.LandingAffiliatesContent, error) {
	landingAffiliatesContent := entity.LandingAffiliatesContent{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "is_enabled", "is_mobile", "is_desktop", "landing_section_header_id", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "landing_affiliates_content" WHERE "deleted_at" IS NULL`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingAffiliatesContent.Id, &landingAffiliatesContent.IsEnabled, &landingAffiliatesContent.IsMobile, &landingAffiliatesContent.IsDesktop, &landingAffiliatesContent.LandingSectionHeaderId,
		&landingAffiliatesContent.CreatedAt, &landingAffiliatesContent.UpdatedAt, &landingAffiliatesContent.DeletedAt,
	)

	return landingAffiliatesContent, err
}

func (r landingAffiliatesContentRepositoryPostgresql) Update(ctx context.Context, landingAffiliatesContent entity.LandingAffiliatesContent) (entity.LandingAffiliatesContent, error) {
	builder := sqlbuilder.New().
		S(`UPDATE "landing_affiliates_content" SET "is_enabled" = $1, "is_mobile" = $2, "is_desktop" = $3, "landing_section_header_id" = $4, "updated_at" = NOW()`, landingAffiliatesContent.IsEnabled, landingAffiliatesContent.IsMobile, landingAffiliatesContent.IsDesktop, landingAffiliatesContent.LandingSectionHeaderId).
		S(`WHERE "deleted_at" IS NULL`).
		S(`RETURNING "id", "is_enabled", "is_mobile", "is_desktop", "landing_section_header_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingAffiliatesContent.Id, &landingAffiliatesContent.IsEnabled, &landingAffiliatesContent.IsMobile, &landingAffiliatesContent.IsDesktop, &landingAffiliatesContent.LandingSectionHeaderId,
		&landingAffiliatesContent.CreatedAt, &landingAffiliatesContent.UpdatedAt, &landingAffiliatesContent.DeletedAt,
	)

	return landingAffiliatesContent, err
}

func (r landingAffiliatesContentRepositoryPostgresql) Delete(ctx context.Context) (entity.LandingAffiliatesContent, error) {
	landingAffiliatesContent := entity.LandingAffiliatesContent{}

	builder := sqlbuilder.New().
		S(`DELETE FROM "landing_affiliates_content"`).
		S(`RETURNING "id", "is_enabled", "is_mobile", "is_desktop", "landing_section_header_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingAffiliatesContent.Id, &landingAffiliatesContent.IsEnabled, &landingAffiliatesContent.IsMobile, &landingAffiliatesContent.IsDesktop, &landingAffiliatesContent.LandingSectionHeaderId,
		&landingAffiliatesContent.CreatedAt, &landingAffiliatesContent.UpdatedAt, &landingAffiliatesContent.DeletedAt,
	)

	return landingAffiliatesContent, err
}

type landingAffiliatesContentAffiliateRepositoryPostgresql struct {
	db DB
}

func NewLandingAffiliatesContentAffiliateRepository(db DB) repository.LandingAffiliatesContentAffiliateRepository {
	return landingAffiliatesContentAffiliateRepositoryPostgresql{db}
}

func (r landingAffiliatesContentAffiliateRepositoryPostgresql) CreateMany(ctx context.Context, landingAffiliatesContentAffiliates []entity.LandingAffiliatesContentAffiliate) ([]entity.LandingAffiliatesContentAffiliate, error) {
	if len(landingAffiliatesContentAffiliates) == 0 {
		return []entity.LandingAffiliatesContentAffiliate{}, nil
	}

	builder := sqlbuilder.New().
		S(`INSERT INTO "landing_affiliates_content_affiliates" ("is_enabled", "is_mobile", "is_desktop", "name", "logo_id", "width", "height", "created_at", "updated_at") VALUES`)

	for i, affiliate := range landingAffiliatesContentAffiliates {
		builder.SA(`(?, ?, ?, ?, ?, ?, ?, NOW(), NOW())`, affiliate.IsEnabled, affiliate.IsMobile, affiliate.IsDesktop, affiliate.Name, affiliate.LogoId, affiliate.Width, affiliate.Height)
		if i+1 < len(landingAffiliatesContentAffiliates) {
			builder.SA(",")
		}
	}

	builder.S(`RETURNING "id", "is_enabled", "is_mobile", "is_desktop", "name", "logo_id", "width", "height", "created_at", "updated_at", "deleted_at"`)

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	createdAffiliates := []entity.LandingAffiliatesContentAffiliate{}
	for rows.Next() {
		affiliate := entity.LandingAffiliatesContentAffiliate{}
		err := rows.Scan(
			&affiliate.Id, &affiliate.IsEnabled, &affiliate.IsMobile, &affiliate.IsDesktop, &affiliate.Name, &affiliate.LogoId, &affiliate.Width, &affiliate.Height,
			&affiliate.CreatedAt, &affiliate.UpdatedAt, &affiliate.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		createdAffiliates = append(createdAffiliates, affiliate)
	}

	return createdAffiliates, nil
}

func (r landingAffiliatesContentAffiliateRepositoryPostgresql) FindAll(ctx context.Context, opt repository.FindAllOptions) ([]entity.LandingAffiliatesContentAffiliate, error) {
	builder := sqlbuilder.New().
		S(`SELECT "id", "is_enabled", "is_mobile", "is_desktop", "name", "logo_id", "width", "height", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "landing_affiliates_content_affiliates" WHERE "deleted_at" IS NULL`).
		S(`ORDER BY "id" ASC`)
	if opt.Limit.Valid {
		builder.SA(`LIMIT ?`, opt.Limit)
	}
	if opt.Offset.Valid {
		builder.SA(`OFFSET ?`, opt.Offset)
	}

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	affiliates := []entity.LandingAffiliatesContentAffiliate{}
	for rows.Next() {
		affiliate := entity.LandingAffiliatesContentAffiliate{}
		err := rows.Scan(
			&affiliate.Id, &affiliate.IsEnabled, &affiliate.IsMobile, &affiliate.IsDesktop, &affiliate.Name, &affiliate.LogoId, &affiliate.Width, &affiliate.Height,
			&affiliate.CreatedAt, &affiliate.UpdatedAt, &affiliate.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		affiliates = append(affiliates, affiliate)
	}

	return affiliates, nil
}

func (r landingAffiliatesContentAffiliateRepositoryPostgresql) DeleteMany(ctx context.Context) ([]entity.LandingAffiliatesContentAffiliate, error) {
	builder := sqlbuilder.New().
		S(`DELETE FROM "landing_affiliates_content_affiliates"`).
		S(`RETURNING "id", "is_enabled", "is_mobile", "is_desktop", "name", "logo_id", "width", "height", "created_at", "updated_at", "deleted_at"`)

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	deletedAffiliates := []entity.LandingAffiliatesContentAffiliate{}
	for rows.Next() {
		affiliate := entity.LandingAffiliatesContentAffiliate{}
		err := rows.Scan(
			&affiliate.Id, &affiliate.IsEnabled, &affiliate.IsMobile, &affiliate.IsDesktop, &affiliate.Name, &affiliate.LogoId, &affiliate.Width, &affiliate.Height,
			&affiliate.CreatedAt, &affiliate.UpdatedAt, &affiliate.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		deletedAffiliates = append(deletedAffiliates, affiliate)
	}

	return deletedAffiliates, nil
}

type landingTestimonialContentRepositoryPostgresql struct {
	db DB
}

func NewLandingTestimonialContentRepository(db DB) repository.LandingTestimonialContentRepository {
	return landingTestimonialContentRepositoryPostgresql{db}
}

func (r landingTestimonialContentRepositoryPostgresql) Create(ctx context.Context, landingTestimonialContent entity.LandingTestimonialContent) (entity.LandingTestimonialContent, error) {
	builder := sqlbuilder.New().
		S(`INSERT INTO "landing_testimonial_content" ("is_enabled", "is_mobile", "is_desktop", "landing_section_header_id", "created_at", "updated_at", "deleted_at")`).
		S(`VALUES ($1, $2, $3, $4, NOW(), NOW(), NULL)`, landingTestimonialContent.IsEnabled, landingTestimonialContent.IsMobile, landingTestimonialContent.IsDesktop, landingTestimonialContent.LandingSectionHeaderId).
		S(`RETURNING "id", "is_enabled", "is_mobile", "is_desktop", "landing_section_header_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingTestimonialContent.Id, &landingTestimonialContent.IsEnabled, &landingTestimonialContent.IsMobile, &landingTestimonialContent.IsDesktop, &landingTestimonialContent.LandingSectionHeaderId,
		&landingTestimonialContent.CreatedAt, &landingTestimonialContent.UpdatedAt, &landingTestimonialContent.DeletedAt,
	)

	return landingTestimonialContent, err
}

func (r landingTestimonialContentRepositoryPostgresql) Find(ctx context.Context) (entity.LandingTestimonialContent, error) {
	landingTestimonialContent := entity.LandingTestimonialContent{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "is_enabled", "is_mobile", "is_desktop", "landing_section_header_id", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "landing_testimonial_content" WHERE "deleted_at" IS NULL`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingTestimonialContent.Id, &landingTestimonialContent.IsEnabled, &landingTestimonialContent.IsMobile, &landingTestimonialContent.IsDesktop, &landingTestimonialContent.LandingSectionHeaderId,
		&landingTestimonialContent.CreatedAt, &landingTestimonialContent.UpdatedAt, &landingTestimonialContent.DeletedAt,
	)

	return landingTestimonialContent, err
}

func (r landingTestimonialContentRepositoryPostgresql) Update(ctx context.Context, landingTestimonialContent entity.LandingTestimonialContent) (entity.LandingTestimonialContent, error) {
	builder := sqlbuilder.New().
		S(`UPDATE "landing_testimonial_content" SET "is_enabled" = $1, "is_mobile" = $2, "is_desktop" = $3, "landing_section_header_id" = $4, "updated_at" = NOW()`, landingTestimonialContent.IsEnabled, landingTestimonialContent.IsMobile, landingTestimonialContent.IsDesktop, landingTestimonialContent.LandingSectionHeaderId).
		S(`WHERE "deleted_at" IS NULL`).
		S(`RETURNING "id", "is_enabled", "is_mobile", "is_desktop", "landing_section_header_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingTestimonialContent.Id, &landingTestimonialContent.IsEnabled, &landingTestimonialContent.IsMobile, &landingTestimonialContent.IsDesktop, &landingTestimonialContent.LandingSectionHeaderId,
		&landingTestimonialContent.CreatedAt, &landingTestimonialContent.UpdatedAt, &landingTestimonialContent.DeletedAt,
	)

	return landingTestimonialContent, err
}

func (r landingTestimonialContentRepositoryPostgresql) Delete(ctx context.Context) (entity.LandingTestimonialContent, error) {
	landingTestimonialContent := entity.LandingTestimonialContent{}

	builder := sqlbuilder.New().
		S(`DELETE FROM "landing_testimonial_content"`).
		S(`RETURNING "id", "is_enabled", "is_mobile", "is_desktop", "landing_section_header_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingTestimonialContent.Id, &landingTestimonialContent.IsEnabled, &landingTestimonialContent.IsMobile, &landingTestimonialContent.IsDesktop, &landingTestimonialContent.LandingSectionHeaderId,
		&landingTestimonialContent.CreatedAt, &landingTestimonialContent.UpdatedAt, &landingTestimonialContent.DeletedAt,
	)

	return landingTestimonialContent, err
}

type landingTestimonialContentReviewRepositoryPostgresql struct {
	db DB
}

func NewLandingTestimonialContentReviewRepository(db DB) repository.LandingTestimonialContentReviewRepository {
	return landingTestimonialContentReviewRepositoryPostgresql{db}
}

func (r landingTestimonialContentReviewRepositoryPostgresql) CreateMany(ctx context.Context, landingTestimonialContentReviews []entity.LandingTestimonialContentReview) ([]entity.LandingTestimonialContentReview, error) {
	if len(landingTestimonialContentReviews) == 0 {
		return []entity.LandingTestimonialContentReview{}, nil
	}

	builder := sqlbuilder.New().
		S(`INSERT INTO "landing_testimonial_content_reviews" ("is_enabled", "is_mobile", "is_desktop", "reviewer", "age", "address", "rating", "review", "created_at", "updated_at") VALUES`)

	for i, review := range landingTestimonialContentReviews {
		builder.SA(`(?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())`, review.IsEnabled, review.IsMobile, review.IsDesktop, review.Reviewer, review.Age, review.Address, review.Rating, review.Review)
		if i+1 < len(landingTestimonialContentReviews) {
			builder.SA(",")
		}
	}

	builder.S(`RETURNING "id", "is_enabled", "is_mobile", "is_desktop", "reviewer", "age", "address", "rating", "review", "created_at", "updated_at", "deleted_at"`)

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	createdTestimonials := []entity.LandingTestimonialContentReview{}
	for rows.Next() {
		faq := entity.LandingTestimonialContentReview{}
		err := rows.Scan(
			&faq.Id, &faq.IsEnabled, &faq.IsMobile, &faq.IsDesktop, &faq.Reviewer, &faq.Age, &faq.Address, &faq.Rating, &faq.Review,
			&faq.CreatedAt, &faq.UpdatedAt, &faq.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		createdTestimonials = append(createdTestimonials, faq)
	}

	return createdTestimonials, nil
}

func (r landingTestimonialContentReviewRepositoryPostgresql) FindAll(ctx context.Context, opt repository.FindAllOptions) ([]entity.LandingTestimonialContentReview, error) {
	builder := sqlbuilder.New().
		S(`SELECT "id", "is_enabled", "is_mobile", "is_desktop", "reviewer", "age", "address", "rating", "review", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "landing_testimonial_content_reviews" WHERE "deleted_at" IS NULL`).
		S(`ORDER BY "id" ASC`)
	if opt.Limit.Valid {
		builder.SA(`LIMIT ?`, opt.Limit)
	}
	if opt.Offset.Valid {
		builder.SA(`OFFSET ?`, opt.Offset)
	}

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	reviews := []entity.LandingTestimonialContentReview{}
	for rows.Next() {
		review := entity.LandingTestimonialContentReview{}
		err := rows.Scan(
			&review.Id, &review.IsEnabled, &review.IsMobile, &review.IsDesktop, &review.Reviewer, &review.Age, &review.Address, &review.Rating, &review.Review,
			&review.CreatedAt, &review.UpdatedAt, &review.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}

	return reviews, nil
}

func (r landingTestimonialContentReviewRepositoryPostgresql) DeleteMany(ctx context.Context) ([]entity.LandingTestimonialContentReview, error) {
	builder := sqlbuilder.New().
		S(`DELETE FROM "landing_testimonial_content_reviews"`).
		S(`RETURNING "id", "is_enabled", "is_mobile", "is_desktop", "reviewer", "age", "address", "rating", "review", "created_at", "updated_at", "deleted_at"`)

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	deletedTestimonials := []entity.LandingTestimonialContentReview{}
	for rows.Next() {
		faq := entity.LandingTestimonialContentReview{}
		err := rows.Scan(
			&faq.Id, &faq.IsEnabled, &faq.IsMobile, &faq.IsDesktop, &faq.Reviewer, &faq.Age, &faq.Address, &faq.Rating, &faq.Review,
			&faq.CreatedAt, &faq.UpdatedAt, &faq.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		deletedTestimonials = append(deletedTestimonials, faq)
	}

	return deletedTestimonials, nil
}

type landingFaqContentRepositoryPostgresql struct {
	db DB
}

func NewLandingFaqContentRepository(db DB) repository.LandingFaqContentRepository {
	return landingFaqContentRepositoryPostgresql{db}
}

func (r landingFaqContentRepositoryPostgresql) Create(ctx context.Context, landingFaqContent entity.LandingFaqContent) (entity.LandingFaqContent, error) {
	builder := sqlbuilder.New().
		S(`INSERT INTO "landing_faq_content" ("is_enabled", "is_mobile", "is_desktop", "landing_section_header_id", "created_at", "updated_at", "deleted_at")`).
		S(`VALUES ($1, $2, $3, $4, NOW(), NOW(), NULL)`, landingFaqContent.IsEnabled, landingFaqContent.IsMobile, landingFaqContent.IsDesktop, landingFaqContent.LandingSectionHeaderId).
		S(`RETURNING "id", "is_enabled", "is_mobile", "is_desktop", "landing_section_header_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingFaqContent.Id, &landingFaqContent.IsEnabled, &landingFaqContent.IsMobile, &landingFaqContent.IsDesktop, &landingFaqContent.LandingSectionHeaderId,
		&landingFaqContent.CreatedAt, &landingFaqContent.UpdatedAt, &landingFaqContent.DeletedAt,
	)

	return landingFaqContent, err
}

func (r landingFaqContentRepositoryPostgresql) Find(ctx context.Context) (entity.LandingFaqContent, error) {
	landingFaqContent := entity.LandingFaqContent{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "is_enabled", "is_mobile", "is_desktop", "landing_section_header_id", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "landing_faq_content" WHERE "deleted_at" IS NULL`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingFaqContent.Id, &landingFaqContent.IsEnabled, &landingFaqContent.IsMobile, &landingFaqContent.IsDesktop, &landingFaqContent.LandingSectionHeaderId,
		&landingFaqContent.CreatedAt, &landingFaqContent.UpdatedAt, &landingFaqContent.DeletedAt,
	)

	return landingFaqContent, err
}

func (r landingFaqContentRepositoryPostgresql) Update(ctx context.Context, landingFaqContent entity.LandingFaqContent) (entity.LandingFaqContent, error) {
	builder := sqlbuilder.New().
		S(`UPDATE "landing_faq_content" SET "is_enabled" = $1, "is_mobile" = $2, "is_desktop" = $3, "landing_section_header_id" = $4, "updated_at" = NOW()`, landingFaqContent.IsEnabled, landingFaqContent.IsMobile, landingFaqContent.IsDesktop, landingFaqContent.LandingSectionHeaderId).
		S(`WHERE "deleted_at" IS NULL`).
		S(`RETURNING "id", "is_enabled", "is_mobile", "is_desktop", "landing_section_header_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingFaqContent.Id, &landingFaqContent.IsEnabled, &landingFaqContent.IsMobile, &landingFaqContent.IsDesktop, &landingFaqContent.LandingSectionHeaderId,
		&landingFaqContent.CreatedAt, &landingFaqContent.UpdatedAt, &landingFaqContent.DeletedAt,
	)

	return landingFaqContent, err
}

func (r landingFaqContentRepositoryPostgresql) Delete(ctx context.Context) (entity.LandingFaqContent, error) {
	landingFaqContent := entity.LandingFaqContent{}

	builder := sqlbuilder.New().
		S(`DELETE FROM "landing_faq_content"`).
		S(`RETURNING "id", "is_enabled", "is_mobile", "is_desktop", "landing_section_header_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingFaqContent.Id, &landingFaqContent.IsEnabled, &landingFaqContent.IsMobile, &landingFaqContent.IsDesktop, &landingFaqContent.LandingSectionHeaderId,
		&landingFaqContent.CreatedAt, &landingFaqContent.UpdatedAt, &landingFaqContent.DeletedAt,
	)

	return landingFaqContent, err
}

type landingFaqContentFaqRepositoryPostgresql struct {
	db DB
}

func NewLandingFaqContentFaqRepository(db DB) repository.LandingFaqContentFaqRepository {
	return landingFaqContentFaqRepositoryPostgresql{db}
}

func (r landingFaqContentFaqRepositoryPostgresql) CreateMany(ctx context.Context, landingFaqContentFaqs []entity.LandingFaqContentFaq) ([]entity.LandingFaqContentFaq, error) {
	if len(landingFaqContentFaqs) == 0 {
		return []entity.LandingFaqContentFaq{}, nil
	}

	builder := sqlbuilder.New().
		S(`INSERT INTO "landing_faq_content_faqs" ("is_enabled", "is_mobile", "is_desktop", "question", "answer", "created_at", "updated_at") VALUES`)

	for i, faq := range landingFaqContentFaqs {
		builder.SA(`(?, ?, ?, ?, ?, NOW(), NOW())`, faq.IsEnabled, faq.IsMobile, faq.IsDesktop, faq.Question, faq.Answer)
		if i+1 < len(landingFaqContentFaqs) {
			builder.SA(",")
		}
	}

	builder.S(`RETURNING "id", "is_enabled", "is_mobile", "is_desktop", "question", "answer", "created_at", "updated_at", "deleted_at"`)

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	createdFaqs := []entity.LandingFaqContentFaq{}
	for rows.Next() {
		faq := entity.LandingFaqContentFaq{}
		err := rows.Scan(
			&faq.Id, &faq.IsEnabled, &faq.IsMobile, &faq.IsDesktop, &faq.Question, &faq.Answer,
			&faq.CreatedAt, &faq.UpdatedAt, &faq.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		createdFaqs = append(createdFaqs, faq)
	}

	return createdFaqs, nil
}

func (r landingFaqContentFaqRepositoryPostgresql) FindAll(ctx context.Context, opt repository.FindAllOptions) ([]entity.LandingFaqContentFaq, error) {
	builder := sqlbuilder.New().
		S(`SELECT "id", "is_enabled", "is_mobile", "is_desktop", "question", "answer", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "landing_faq_content_faqs" WHERE "deleted_at" IS NULL`).
		S(`ORDER BY "id" ASC`)
	if opt.Limit.Valid {
		builder.SA(`LIMIT ?`, opt.Limit)
	}
	if opt.Offset.Valid {
		builder.SA(`OFFSET ?`, opt.Offset)
	}

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	faqs := []entity.LandingFaqContentFaq{}
	for rows.Next() {
		faq := entity.LandingFaqContentFaq{}
		err := rows.Scan(
			&faq.Id, &faq.IsEnabled, &faq.IsMobile, &faq.IsDesktop, &faq.Question, &faq.Answer,
			&faq.CreatedAt, &faq.UpdatedAt, &faq.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		faqs = append(faqs, faq)
	}

	return faqs, nil
}

func (r landingFaqContentFaqRepositoryPostgresql) DeleteMany(ctx context.Context) ([]entity.LandingFaqContentFaq, error) {
	builder := sqlbuilder.New().
		S(`DELETE FROM "landing_faq_content_faqs"`).
		S(`RETURNING "id", "is_enabled", "is_mobile", "is_desktop", "question", "answer", "created_at", "updated_at", "deleted_at"`)

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	deletedFaqs := []entity.LandingFaqContentFaq{}
	for rows.Next() {
		faq := entity.LandingFaqContentFaq{}
		err := rows.Scan(
			&faq.Id, &faq.IsEnabled, &faq.IsMobile, &faq.IsDesktop, &faq.Question, &faq.Answer,
			&faq.CreatedAt, &faq.UpdatedAt, &faq.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		deletedFaqs = append(deletedFaqs, faq)
	}

	return deletedFaqs, nil
}

type landingMenuRepositoryPostgresql struct {
	db DB
}

func NewLandingMenuRepository(db DB) repository.LandingMenuRepository {
	return landingMenuRepositoryPostgresql{db}
}

func (r landingMenuRepositoryPostgresql) CreateMany(ctx context.Context, landingMenus []entity.LandingMenu) ([]entity.LandingMenu, error) {
	if len(landingMenus) == 0 {
		return []entity.LandingMenu{}, nil
	}

	builder := sqlbuilder.New().
		S(`INSERT INTO "landing_menus" ("is_enabled", "is_mobile", "is_desktop", "icon", "label", "path", "created_at", "updated_at") VALUES`)

	for i, menu := range landingMenus {
		builder.SA(`(?, ?, ?, ?, ?, ?, NOW(), NOW())`, menu.IsEnabled, menu.IsMobile, menu.IsDesktop, menu.Icon, menu.Label, menu.Path)
		if i+1 < len(landingMenus) {
			builder.SA(",")
		}
	}

	builder.S(`RETURNING "id", "is_enabled", "is_mobile", "is_desktop", "icon", "label", "path", "created_at", "updated_at", "deleted_at"`)

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	createdMenus := []entity.LandingMenu{}
	for rows.Next() {
		menu := entity.LandingMenu{}
		err := rows.Scan(
			&menu.Id, &menu.IsEnabled, &menu.IsMobile, &menu.IsDesktop, &menu.Icon, &menu.Label, &menu.Path,
			&menu.CreatedAt, &menu.UpdatedAt, &menu.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		createdMenus = append(createdMenus, menu)
	}

	return createdMenus, nil
}

func (r landingMenuRepositoryPostgresql) FindAll(ctx context.Context, opt repository.FindAllOptions) ([]entity.LandingMenu, error) {
	builder := sqlbuilder.New().
		S(`SELECT "id", "is_enabled", "is_mobile", "is_desktop", "icon", "label", "path", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "landing_menus" WHERE "deleted_at" IS NULL`).
		S(`ORDER BY "id" ASC`)
	if opt.Limit.Valid {
		builder.SA(`LIMIT ?`, opt.Limit)
	}
	if opt.Offset.Valid {
		builder.SA(`OFFSET ?`, opt.Offset)
	}

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	menus := []entity.LandingMenu{}
	for rows.Next() {
		menu := entity.LandingMenu{}
		err := rows.Scan(
			&menu.Id, &menu.IsEnabled, &menu.IsMobile, &menu.IsDesktop, &menu.Icon, &menu.Label, &menu.Path,
			&menu.CreatedAt, &menu.UpdatedAt, &menu.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		menus = append(menus, menu)
	}

	return menus, nil
}

func (r landingMenuRepositoryPostgresql) DeleteMany(ctx context.Context) ([]entity.LandingMenu, error) {
	builder := sqlbuilder.New().
		S(`DELETE FROM "landing_menus"`).
		S(`RETURNING "id", "is_enabled", "is_mobile", "is_desktop", "icon", "label", "path", "created_at", "updated_at", "deleted_at"`)

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	deletedMenus := []entity.LandingMenu{}
	for rows.Next() {
		menu := entity.LandingMenu{}
		err := rows.Scan(
			&menu.Id, &menu.IsEnabled, &menu.IsMobile, &menu.IsDesktop, &menu.Icon, &menu.Label, &menu.Path,
			&menu.CreatedAt, &menu.UpdatedAt, &menu.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		deletedMenus = append(deletedMenus, menu)
	}

	return deletedMenus, nil
}
