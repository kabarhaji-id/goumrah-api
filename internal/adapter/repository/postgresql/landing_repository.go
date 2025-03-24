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
		S(`INSERT INTO "landing_hero_content" ("is_enabled", "title", "description", "tags_line", "button_label", "button_url", "image_id", "created_at", "updated_at", "deleted_at")`).
		S(`VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(), NOW(), NULL)`, landingHeroContent.IsEnabled, landingHeroContent.Title, landingHeroContent.Description, landingHeroContent.TagsLine, landingHeroContent.ButtonLabel, landingHeroContent.ButtonUrl, landingHeroContent.ImageId).
		S(`RETURNING "id", "is_enabled", "title", "description", "tags_line", "button_label", "button_url", "image_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingHeroContent.Id, &landingHeroContent.IsEnabled, &landingHeroContent.Title, &landingHeroContent.Description, &landingHeroContent.TagsLine, &landingHeroContent.ButtonLabel, &landingHeroContent.ButtonUrl, &landingHeroContent.ImageId,
		&landingHeroContent.CreatedAt, &landingHeroContent.UpdatedAt, &landingHeroContent.DeletedAt,
	)

	return landingHeroContent, err
}

func (r landingHeroContentRepositoryPostgresql) Find(ctx context.Context) (entity.LandingHeroContent, error) {
	landingHeroContent := entity.LandingHeroContent{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "is_enabled", "title", "description", "tags_line", "button_label", "button_url", "image_id", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "landing_hero_content" WHERE "deleted_at" IS NULL`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingHeroContent.Id, &landingHeroContent.IsEnabled, &landingHeroContent.Title, &landingHeroContent.Description, &landingHeroContent.TagsLine, &landingHeroContent.ButtonLabel, &landingHeroContent.ButtonUrl, &landingHeroContent.ImageId,
		&landingHeroContent.CreatedAt, &landingHeroContent.UpdatedAt, &landingHeroContent.DeletedAt,
	)

	return landingHeroContent, err
}

func (r landingHeroContentRepositoryPostgresql) Update(ctx context.Context, landingHeroContent entity.LandingHeroContent) (entity.LandingHeroContent, error) {
	builder := sqlbuilder.New().
		S(`UPDATE "landing_hero_content" SET "is_enabled" = $1, "title" = $2, "description" = $3, "tags_line" = $4, "button_label" = $5, "button_url" = $6, "image_id" = $7, "updated_at" = NOW()`, landingHeroContent.IsEnabled, landingHeroContent.Title, landingHeroContent.Description, landingHeroContent.TagsLine, landingHeroContent.ButtonLabel, landingHeroContent.ButtonUrl, landingHeroContent.ImageId).
		S(`WHERE "id" = $8 AND "deleted_at" IS NULL`, landingHeroContent.Id).
		S(`RETURNING "id", "is_enabled", "title", "description", "tags_line", "button_label", "button_url", "image_id", "created_at", "updated_at", "deleted_at`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingHeroContent.Id, &landingHeroContent.IsEnabled, &landingHeroContent.Title, &landingHeroContent.Description, &landingHeroContent.TagsLine, &landingHeroContent.ButtonLabel, &landingHeroContent.ButtonUrl, &landingHeroContent.ImageId,
		&landingHeroContent.CreatedAt, &landingHeroContent.UpdatedAt, &landingHeroContent.DeletedAt,
	)

	return landingHeroContent, err
}

func (r landingHeroContentRepositoryPostgresql) Delete(ctx context.Context) (entity.LandingHeroContent, error) {
	landingHeroContent := entity.LandingHeroContent{}

	builder := sqlbuilder.New().
		S(`DELETE FROM "landing_hero_content"`).
		S(`RETURNING "id", "is_enabled", "title", "description", "tags_line", "button_label", "button_url", "image_id", "created_at", "updated_at", "deleted_at`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingHeroContent.Id, &landingHeroContent.IsEnabled, &landingHeroContent.Title, &landingHeroContent.Description, &landingHeroContent.TagsLine, &landingHeroContent.ButtonLabel, &landingHeroContent.ButtonUrl, &landingHeroContent.ImageId,
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
		S(`INSERT INTO "landing_section_headers" ("is_enabled", "title", "subtitle", "tags_line", "created_at", "updated_at", "deleted_at")`).
		S(`VALUES ($1, $2, $3, $4, NOW(), NOW(), NULL)`, landingSectionHeader.IsEnabled, landingSectionHeader.Title, landingSectionHeader.Subtitle, landingSectionHeader.TagsLine).
		S(`RETURNING "id", "is_enabled", "title", "subtitle", "tags_line", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingSectionHeader.Id, &landingSectionHeader.IsEnabled, &landingSectionHeader.Title, &landingSectionHeader.Subtitle, &landingSectionHeader.TagsLine,
		&landingSectionHeader.CreatedAt, &landingSectionHeader.UpdatedAt, &landingSectionHeader.DeletedAt,
	)

	return landingSectionHeader, err
}

func (r landingSectionHeaderRepositoryPostgresql) FindById(ctx context.Context, id int64) (entity.LandingSectionHeader, error) {
	builder := sqlbuilder.New().
		S(`SELECT "id", "is_enabled", "title", "subtitle", "tags_line", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "landing_section_headers" WHERE "id" = $1 AND "deleted_at" IS NULL`, id)

	landingSectionHeader := entity.LandingSectionHeader{}
	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingSectionHeader.Id, &landingSectionHeader.IsEnabled, &landingSectionHeader.Title, &landingSectionHeader.Subtitle, &landingSectionHeader.TagsLine,
		&landingSectionHeader.CreatedAt, &landingSectionHeader.UpdatedAt, &landingSectionHeader.DeletedAt,
	)

	return landingSectionHeader, err
}

func (r landingSectionHeaderRepositoryPostgresql) FindAll(ctx context.Context, opt repository.FindAllOptions) ([]entity.LandingSectionHeader, error) {
	builder := sqlbuilder.New().
		S(`SELECT "id", "is_enabled", "title", "subtitle", "tags_line", "created_at", "updated_at", "deleted_at"`).
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
			&landingSectionHeader.Id, &landingSectionHeader.IsEnabled, &landingSectionHeader.Title, &landingSectionHeader.Subtitle, &landingSectionHeader.TagsLine,
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
		S(`UPDATE "landing_section_headers" SET "is_enabled" = $1, "title" = $2, "subtitle" = $3, "tags_line" = $4, "updated_at" = NOW()`).
		S(`WHERE "id" = $5 AND "deleted_at" IS NULL`, landingSectionHeader.IsEnabled, landingSectionHeader.Title, landingSectionHeader.Subtitle, landingSectionHeader.TagsLine, id).
		S(`RETURNING "id", "is_enabled", "title", "subtitle", "tags_line", "created_at", "updated_at", "deleted_at`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingSectionHeader.Id, &landingSectionHeader.IsEnabled, &landingSectionHeader.Title, &landingSectionHeader.Subtitle, &landingSectionHeader.TagsLine,
		&landingSectionHeader.CreatedAt, &landingSectionHeader.UpdatedAt, &landingSectionHeader.DeletedAt,
	)

	return landingSectionHeader, err
}

func (r landingSectionHeaderRepositoryPostgresql) Delete(ctx context.Context, id int64) (entity.LandingSectionHeader, error) {
	builder := sqlbuilder.New().
		S(`UPDATE "landing_section_headers" SET "deleted_at" = NOW()`).
		S(`WHERE "id" = $1 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "is_enabled", "title", "subtitle", "tags_line", "created_at", "updated_at", "deleted_at`)

	landingSectionHeader := entity.LandingSectionHeader{}
	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingSectionHeader.Id, &landingSectionHeader.IsEnabled, &landingSectionHeader.Title, &landingSectionHeader.Subtitle, &landingSectionHeader.TagsLine,
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
		S(`INSERT INTO "landing_package_items" ("is_enabled", "package_id", "button_label", "created_at", "updated_at", "deleted_at")`).
		S(`VALUES ($1, $2, $3, NOW(), NOW(), NULL)`, landingPackageItem.IsEnabled, landingPackageItem.PackageId, landingPackageItem.ButtonLabel).
		S(`RETURNING "id", "is_enabled", "package_id", "button_label", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingPackageItem.Id, &landingPackageItem.IsEnabled, &landingPackageItem.PackageId, &landingPackageItem.ButtonLabel,
		&landingPackageItem.CreatedAt, &landingPackageItem.UpdatedAt, &landingPackageItem.DeletedAt,
	)

	return landingPackageItem, err
}

func (r landingPackageItemRepositoryPostgresql) FindById(ctx context.Context, id int64) (entity.LandingPackageItem, error) {
	builder := sqlbuilder.New().
		S(`SELECT "id", "is_enabled", "package_id", "button_label", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "landing_package_items" WHERE "id" = $1 AND "deleted_at" IS NULL`, id)

	landingPackageItem := entity.LandingPackageItem{}
	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingPackageItem.Id, &landingPackageItem.IsEnabled, &landingPackageItem.PackageId, &landingPackageItem.ButtonLabel,
		&landingPackageItem.CreatedAt, &landingPackageItem.UpdatedAt, &landingPackageItem.DeletedAt,
	)

	return landingPackageItem, err
}

func (r landingPackageItemRepositoryPostgresql) FindAll(ctx context.Context, opt repository.FindAllOptions) ([]entity.LandingPackageItem, error) {
	builder := sqlbuilder.New().
		S(`SELECT "id", "is_enabled", "package_id", "button_label", "created_at", "updated_at", "deleted_at"`).
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
			&landingPackageItem.Id, &landingPackageItem.IsEnabled, &landingPackageItem.PackageId, &landingPackageItem.ButtonLabel,
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
		S(`UPDATE "landing_package_items" SET "is_enabled" = $1, "package_id" = $2, "button_label" = $3, "updated_at" = NOW()`).
		S(`WHERE "id" = $4 AND "deleted_at" IS NULL`, landingPackageItem.IsEnabled, landingPackageItem.PackageId, landingPackageItem.ButtonLabel, id).
		S(`RETURNING "id", "is_enabled", "package_id", "button_label", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingPackageItem.Id, &landingPackageItem.IsEnabled, &landingPackageItem.PackageId, &landingPackageItem.ButtonLabel,
		&landingPackageItem.CreatedAt, &landingPackageItem.UpdatedAt, &landingPackageItem.DeletedAt,
	)

	return landingPackageItem, err
}

func (r landingPackageItemRepositoryPostgresql) Delete(ctx context.Context, id int64) (entity.LandingPackageItem, error) {
	builder := sqlbuilder.New().
		S(`UPDATE "landing_package_items" SET "deleted_at" = NOW()`).
		S(`WHERE "id" = $1 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "is_enabled", "package_id", "button_label", "created_at", "updated_at", "deleted_at"`)

	landingPackageItem := entity.LandingPackageItem{}
	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingPackageItem.Id, &landingPackageItem.IsEnabled, &landingPackageItem.PackageId, &landingPackageItem.ButtonLabel,
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
		S(`INSERT INTO "landing_single_package_content" ("is_enabled", "landing_section_header_id", "silver_landing_package_item_id", "gold_landing_package_item_id", "platinum_landing_package_item_id", "created_at", "updated_at", "deleted_at")`).
		S(`VALUES ($1, $2, $3, $4, $5, NOW(), NOW(), NULL)`, landingSinglePackageContent.IsEnabled, landingSinglePackageContent.LandingSectionHeaderId, landingSinglePackageContent.SilverLandingPackageItemId, landingSinglePackageContent.GoldLandingPackageItemId, landingSinglePackageContent.PlatinumLandingPackageItemId).
		S(`RETURNING "id", "is_enabled", "landing_section_header_id", "silver_landing_package_item_id", "gold_landing_package_item_id", "platinum_landing_package_item_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingSinglePackageContent.Id, &landingSinglePackageContent.IsEnabled, &landingSinglePackageContent.LandingSectionHeaderId,
		&landingSinglePackageContent.SilverLandingPackageItemId, &landingSinglePackageContent.GoldLandingPackageItemId,
		&landingSinglePackageContent.PlatinumLandingPackageItemId, &landingSinglePackageContent.CreatedAt,
		&landingSinglePackageContent.UpdatedAt, &landingSinglePackageContent.DeletedAt,
	)

	return landingSinglePackageContent, err
}

func (r landingSinglePackageContentRepositoryPostgresql) Find(ctx context.Context) (entity.LandingSinglePackageContent, error) {
	landingSinglePackageContent := entity.LandingSinglePackageContent{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "is_enabled", "landing_section_header_id", "silver_landing_package_item_id", "gold_landing_package_item_id", "platinum_landing_package_item_id", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "landing_single_package_content" WHERE "deleted_at" IS NULL`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingSinglePackageContent.Id, &landingSinglePackageContent.IsEnabled, &landingSinglePackageContent.LandingSectionHeaderId,
		&landingSinglePackageContent.SilverLandingPackageItemId, &landingSinglePackageContent.GoldLandingPackageItemId,
		&landingSinglePackageContent.PlatinumLandingPackageItemId, &landingSinglePackageContent.CreatedAt,
		&landingSinglePackageContent.UpdatedAt, &landingSinglePackageContent.DeletedAt,
	)

	return landingSinglePackageContent, err
}

func (r landingSinglePackageContentRepositoryPostgresql) Update(ctx context.Context, landingSinglePackageContent entity.LandingSinglePackageContent) (entity.LandingSinglePackageContent, error) {
	builder := sqlbuilder.New().
		S(`UPDATE "landing_single_package_content" SET "is_enabled" = $1, "landing_section_header_id" = $2, "silver_landing_package_item_id" = $3, "gold_landing_package_item_id" = $4, "platinum_landing_package_item_id" = $5, "updated_at" = NOW()`).
		S(`WHERE "id" = $6 AND "deleted_at" IS NULL`, landingSinglePackageContent.IsEnabled, landingSinglePackageContent.LandingSectionHeaderId, landingSinglePackageContent.SilverLandingPackageItemId, landingSinglePackageContent.GoldLandingPackageItemId, landingSinglePackageContent.PlatinumLandingPackageItemId, landingSinglePackageContent.Id).
		S(`RETURNING "id", "is_enabled", "landing_section_header_id", "silver_landing_package_item_id", "gold_landing_package_item_id", "platinum_landing_package_item_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingSinglePackageContent.Id, &landingSinglePackageContent.IsEnabled, &landingSinglePackageContent.LandingSectionHeaderId,
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
		S(`RETURNING "id", "is_enabled", "landing_section_header_id", "silver_landing_package_item_id", "gold_landing_package_item_id", "platinum_landing_package_item_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingSinglePackageContent.Id, &landingSinglePackageContent.IsEnabled, &landingSinglePackageContent.LandingSectionHeaderId,
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
		S(`INSERT INTO "landing_package_details" ("is_enabled", "landing_section_header_id", "created_at", "updated_at", "deleted_at")`).
		S(`VALUES ($1, $2, NOW(), NOW(), NULL)`, landingPackageDetail.IsEnabled, landingPackageDetail.LandingSectionHeaderId).
		S(`RETURNING "id", "is_enabled", "landing_section_header_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingPackageDetail.Id, &landingPackageDetail.IsEnabled, &landingPackageDetail.LandingSectionHeaderId,
		&landingPackageDetail.CreatedAt, &landingPackageDetail.UpdatedAt, &landingPackageDetail.DeletedAt,
	)

	return landingPackageDetail, err
}

func (r landingPackageDetailRepositoryPostgresql) FindById(ctx context.Context, id int64) (entity.LandingPackageDetail, error) {
	builder := sqlbuilder.New().
		S(`SELECT "id", "is_enabled", "landing_section_header_id", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "landing_package_details" WHERE "id" = $1 AND "deleted_at" IS NULL`, id)

	landingPackageDetail := entity.LandingPackageDetail{}
	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingPackageDetail.Id, &landingPackageDetail.IsEnabled, &landingPackageDetail.LandingSectionHeaderId,
		&landingPackageDetail.CreatedAt, &landingPackageDetail.UpdatedAt, &landingPackageDetail.DeletedAt,
	)

	return landingPackageDetail, err
}

func (r landingPackageDetailRepositoryPostgresql) FindAll(ctx context.Context, opt repository.FindAllOptions) ([]entity.LandingPackageDetail, error) {
	builder := sqlbuilder.New().
		S(`SELECT "id", "is_enabled", "landing_section_header_id", "created_at", "updated_at", "deleted_at"`).
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
			&landingPackageDetail.Id, &landingPackageDetail.IsEnabled, &landingPackageDetail.LandingSectionHeaderId,
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
		S(`UPDATE "landing_package_details" SET "is_enabled" = $1, "landing_section_header_id" = $2, "updated_at" = NOW()`).
		S(`WHERE "id" = $3 AND "deleted_at" IS NULL`, landingPackageDetail.IsEnabled, landingPackageDetail.LandingSectionHeaderId, id).
		S(`RETURNING "id", "is_enabled", "landing_section_header_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingPackageDetail.Id, &landingPackageDetail.IsEnabled, &landingPackageDetail.LandingSectionHeaderId,
		&landingPackageDetail.CreatedAt, &landingPackageDetail.UpdatedAt, &landingPackageDetail.DeletedAt,
	)

	return landingPackageDetail, err
}

func (r landingPackageDetailRepositoryPostgresql) Delete(ctx context.Context, id int64) (entity.LandingPackageDetail, error) {
	builder := sqlbuilder.New().
		S(`UPDATE "landing_package_details" SET "deleted_at" = NULL`).
		S(`WHERE "id" = $1 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "is_enabled", "landing_section_header_id", "created_at", "updated_at", "deleted_at"`)

	landingPackageDetail := entity.LandingPackageDetail{}
	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingPackageDetail.Id, &landingPackageDetail.IsEnabled, &landingPackageDetail.LandingSectionHeaderId,
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
	builder := sqlbuilder.New().
		S(`INSERT INTO "landing_package_detail_items" ("is_enabled", "landing_package_detail_id", "landing_package_item_id", "created_at", "updated_at")`).
		S(`VALUES`)

	for i, item := range landingPackageDetailItems {
		if i > 0 {
			builder.SA(",")
		}
		builder.SA(`(?, ?, ?, NOW(), NOW())`, item.IsEnabled, item.LandingPackageDetailId, item.LandingPackageItemId)
	}

	builder.S(`RETURNING "is_enabled", "landing_package_detail_id", "landing_package_item_id", "created_at", "updated_at", "deleted_at"`)

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	createdItems := []entity.LandingPackageDetailItem{}
	for rows.Next() {
		item := entity.LandingPackageDetailItem{}
		err := rows.Scan(&item.IsEnabled, &item.LandingPackageDetailId, &item.LandingPackageItemId, &item.CreatedAt, &item.UpdatedAt, &item.DeletedAt)
		if err != nil {
			return nil, err
		}
		createdItems = append(createdItems, item)
	}

	return createdItems, nil
}

func (r landingPackageDetailItemRepositoryPostgresql) FindAll(ctx context.Context, opt repository.FindAllOptions) ([]entity.LandingPackageDetailItem, error) {
	builder := sqlbuilder.New().
		S(`SELECT "is_enabled", "landing_package_detail_id", "landing_package_item_id", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "landing_package_detail_items" WHERE "deleted_at" IS NULL`).
		S(`ORDER BY "landing_package_detail_id" ASC`)
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
		err := rows.Scan(&item.IsEnabled, &item.LandingPackageDetailId, &item.LandingPackageItemId, &item.CreatedAt, &item.UpdatedAt, &item.DeletedAt)
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
		S(`RETURNING "is_enabled", "landing_package_detail_id", "landing_package_item_id", "created_at", "updated_at", "deleted_at"`)

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	deletedItems := []entity.LandingPackageDetailItem{}
	for rows.Next() {
		item := entity.LandingPackageDetailItem{}
		err := rows.Scan(&item.IsEnabled, &item.LandingPackageDetailId, &item.LandingPackageItemId, &item.CreatedAt, &item.UpdatedAt, &item.DeletedAt)
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
		S(`INSERT INTO "landing_packages_content" ("is_enabled", "silver_landing_package_detail_id", "gold_landing_package_detail_id", "platinum_landing_package_detail_id", "created_at", "updated_at", "deleted_at")`).
		S(`VALUES ($1, $2, $3, $4, NOW(), NOW(), NULL)`, landingPackagesContent.IsEnabled, landingPackagesContent.SilverLandingPackageDetailId, landingPackagesContent.GoldLandingPackageDetailId, landingPackagesContent.PlatinumLandingPackageDetailId).
		S(`RETURNING "id", "is_enabled", "silver_landing_package_detail_id", "gold_landing_package_detail_id", "platinum_landing_package_detail_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingPackagesContent.Id, &landingPackagesContent.IsEnabled, &landingPackagesContent.SilverLandingPackageDetailId,
		&landingPackagesContent.GoldLandingPackageDetailId, &landingPackagesContent.PlatinumLandingPackageDetailId,
		&landingPackagesContent.CreatedAt, &landingPackagesContent.UpdatedAt, &landingPackagesContent.DeletedAt,
	)

	return landingPackagesContent, err
}

func (r landingPackagesContentRepositoryPostgresql) Find(ctx context.Context) (entity.LandingPackagesContent, error) {
	landingPackagesContent := entity.LandingPackagesContent{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "is_enabled", "silver_landing_package_detail_id", "gold_landing_package_detail_id", "platinum_landing_package_detail_id", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "landing_packages_content" WHERE "deleted_at" IS NULL`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingPackagesContent.Id, &landingPackagesContent.IsEnabled, &landingPackagesContent.SilverLandingPackageDetailId,
		&landingPackagesContent.GoldLandingPackageDetailId, &landingPackagesContent.PlatinumLandingPackageDetailId,
		&landingPackagesContent.CreatedAt, &landingPackagesContent.UpdatedAt, &landingPackagesContent.DeletedAt,
	)

	return landingPackagesContent, err
}

func (r landingPackagesContentRepositoryPostgresql) Update(ctx context.Context, landingPackagesContent entity.LandingPackagesContent) (entity.LandingPackagesContent, error) {
	builder := sqlbuilder.New().
		S(`UPDATE "landing_packages_content" SET "is_enabled" = $1, "silver_landing_package_detail_id" = $2, "gold_landing_package_detail_id" = $3, "platinum_landing_package_detail_id" = $4, "updated_at" = NOW()`).
		S(`WHERE "id" = $5 AND "deleted_at" IS NULL`, landingPackagesContent.IsEnabled, landingPackagesContent.SilverLandingPackageDetailId, landingPackagesContent.GoldLandingPackageDetailId, landingPackagesContent.PlatinumLandingPackageDetailId, landingPackagesContent.Id).
		S(`RETURNING "id", "is_enabled", "silver_landing_package_detail_id", "gold_landing_package_detail_id", "platinum_landing_package_detail_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingPackagesContent.Id, &landingPackagesContent.IsEnabled, &landingPackagesContent.SilverLandingPackageDetailId,
		&landingPackagesContent.GoldLandingPackageDetailId, &landingPackagesContent.PlatinumLandingPackageDetailId,
		&landingPackagesContent.CreatedAt, &landingPackagesContent.UpdatedAt, &landingPackagesContent.DeletedAt,
	)

	return landingPackagesContent, err
}

func (r landingPackagesContentRepositoryPostgresql) Delete(ctx context.Context) (entity.LandingPackagesContent, error) {
	landingPackagesContent := entity.LandingPackagesContent{}

	builder := sqlbuilder.New().
		S(`DELETE FROM "landing_packages_content"`).
		S(`RETURNING "id", "is_enabled", "silver_landing_package_detail_id", "gold_landing_package_detail_id", "platinum_landing_package_detail_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingPackagesContent.Id, &landingPackagesContent.IsEnabled, &landingPackagesContent.SilverLandingPackageDetailId,
		&landingPackagesContent.GoldLandingPackageDetailId, &landingPackagesContent.PlatinumLandingPackageDetailId,
		&landingPackagesContent.CreatedAt, &landingPackagesContent.UpdatedAt, &landingPackagesContent.DeletedAt,
	)

	return landingPackagesContent, err
}

type landingFeaturesContentRepositoryPostgresql struct {
	db DB
}

func NewLandingFeaturesContentRepository(db DB) repository.LandingFeaturesContentRepository {
	return landingFeaturesContentRepositoryPostgresql{db}
}

func (r landingFeaturesContentRepositoryPostgresql) Create(ctx context.Context, landingFeaturesContent entity.LandingFeaturesContent) (entity.LandingFeaturesContent, error) {
	builder := sqlbuilder.New().
		S(`INSERT INTO "landing_features_content" ("is_enabled", "landing_section_header_id", "footer_title", "button_about", "button_package", "created_at", "updated_at", "deleted_at")`).
		S(`VALUES ($1, $2, $3, $4, $5, NOW(), NOW(), NULL)`, landingFeaturesContent.IsEnabled, landingFeaturesContent.LandingSectionHeaderId, landingFeaturesContent.FooterTitle, landingFeaturesContent.ButtonAbout, landingFeaturesContent.ButtonPackage).
		S(`RETURNING "id", "is_enabled", "landing_section_header_id", "footer_title", "button_about", "button_package", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingFeaturesContent.Id, &landingFeaturesContent.IsEnabled, &landingFeaturesContent.LandingSectionHeaderId,
		&landingFeaturesContent.FooterTitle, &landingFeaturesContent.ButtonAbout, &landingFeaturesContent.ButtonPackage,
		&landingFeaturesContent.CreatedAt, &landingFeaturesContent.UpdatedAt, &landingFeaturesContent.DeletedAt,
	)

	return landingFeaturesContent, err
}

func (r landingFeaturesContentRepositoryPostgresql) Find(ctx context.Context) (entity.LandingFeaturesContent, error) {
	landingFeaturesContent := entity.LandingFeaturesContent{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "is_enabled", "landing_section_header_id", "footer_title", "button_about", "button_package", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "landing_features_content" WHERE "deleted_at" IS NULL`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingFeaturesContent.Id, &landingFeaturesContent.IsEnabled, &landingFeaturesContent.LandingSectionHeaderId,
		&landingFeaturesContent.FooterTitle, &landingFeaturesContent.ButtonAbout, &landingFeaturesContent.ButtonPackage,
		&landingFeaturesContent.CreatedAt, &landingFeaturesContent.UpdatedAt, &landingFeaturesContent.DeletedAt,
	)

	return landingFeaturesContent, err
}

func (r landingFeaturesContentRepositoryPostgresql) Update(ctx context.Context, landingFeaturesContent entity.LandingFeaturesContent) (entity.LandingFeaturesContent, error) {
	builder := sqlbuilder.New().
		S(`UPDATE "landing_features_content" SET "is_enabled" = $1, "landing_section_header_id" = $2, "footer_title" = $3, "button_about" = $4, "button_package" = $5, "updated_at" = NOW()`).
		S(`WHERE "id" = $6 AND "deleted_at" IS NULL`, landingFeaturesContent.IsEnabled, landingFeaturesContent.LandingSectionHeaderId, landingFeaturesContent.FooterTitle, landingFeaturesContent.ButtonAbout, landingFeaturesContent.ButtonPackage, landingFeaturesContent.Id).
		S(`RETURNING "id", "is_enabled", "landing_section_header_id", "footer_title", "button_about", "button_package", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingFeaturesContent.Id, &landingFeaturesContent.IsEnabled, &landingFeaturesContent.LandingSectionHeaderId,
		&landingFeaturesContent.FooterTitle, &landingFeaturesContent.ButtonAbout, &landingFeaturesContent.ButtonPackage,
		&landingFeaturesContent.CreatedAt, &landingFeaturesContent.UpdatedAt, &landingFeaturesContent.DeletedAt,
	)

	return landingFeaturesContent, err
}

func (r landingFeaturesContentRepositoryPostgresql) Delete(ctx context.Context) (entity.LandingFeaturesContent, error) {
	landingFeaturesContent := entity.LandingFeaturesContent{}

	builder := sqlbuilder.New().
		S(`DELETE FROM "landing_features_content"`).
		S(`RETURNING "id", "is_enabled", "landing_section_header_id", "footer_title", "button_about", "button_package", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingFeaturesContent.Id, &landingFeaturesContent.IsEnabled, &landingFeaturesContent.LandingSectionHeaderId,
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
	builder := sqlbuilder.New().
		S(`INSERT INTO "landing_features_content_benefits" ("is_enabled", "title", "subtitle", "logo_id", "created_at", "updated_at") VALUES`)

	for i, benefit := range landingFeaturesContentBenefits {
		if i > 0 {
			builder.SA(",")
		}
		builder.SA(`(?, ?, ?, ?, NOW(), NOW())`, benefit.IsEnabled, benefit.Title, benefit.Subtitle, benefit.LogoId)
	}

	builder.S(`RETURNING "id", "is_enabled", "title", "subtitle", "logo_id", "created_at", "updated_at", "deleted_at"`)

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	createdBenefits := []entity.LandingFeaturesContentBenefit{}
	for rows.Next() {
		benefit := entity.LandingFeaturesContentBenefit{}
		err := rows.Scan(
			&benefit.Id, &benefit.IsEnabled, &benefit.Title,
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
		S(`SELECT "id", "is_enabled", "title", "subtitle", "logo_id", "created_at", "updated_at", "deleted_at"`).
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
			&benefit.Id, &benefit.IsEnabled, &benefit.Title, &benefit.Subtitle, &benefit.LogoId,
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
		S(`RETURNING "id", "is_enabled", "title", "subtitle", "logo_id", "created_at", "updated_at", "deleted_at"`)

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	deletedBenefits := []entity.LandingFeaturesContentBenefit{}
	for rows.Next() {
		benefit := entity.LandingFeaturesContentBenefit{}
		err := rows.Scan(
			&benefit.Id, &benefit.IsEnabled, &benefit.Title, &benefit.Subtitle, &benefit.LogoId,
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
		S(`INSERT INTO "landing_moments_content" ("is_enabled", "landing_section_header_id", "created_at", "updated_at", "deleted_at")`).
		S(`VALUES ($1, $2, NOW(), NOW(), NULL)`, landingMomentsContent.IsEnabled, landingMomentsContent.LandingSectionHeaderId).
		S(`RETURNING "id", "is_enabled", "landing_section_header_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingMomentsContent.Id, &landingMomentsContent.IsEnabled, &landingMomentsContent.LandingSectionHeaderId,
		&landingMomentsContent.CreatedAt, &landingMomentsContent.UpdatedAt, &landingMomentsContent.DeletedAt,
	)

	return landingMomentsContent, err
}

func (r landingMomentsContentRepositoryPostgresql) Find(ctx context.Context) (entity.LandingMomentsContent, error) {
	landingMomentsContent := entity.LandingMomentsContent{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "is_enabled", "landing_section_header_id", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "landing_moments_content" WHERE "deleted_at" IS NULL`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingMomentsContent.Id, &landingMomentsContent.IsEnabled, &landingMomentsContent.LandingSectionHeaderId,
		&landingMomentsContent.CreatedAt, &landingMomentsContent.UpdatedAt, &landingMomentsContent.DeletedAt,
	)

	return landingMomentsContent, err
}

func (r landingMomentsContentRepositoryPostgresql) Update(ctx context.Context, landingMomentsContent entity.LandingMomentsContent) (entity.LandingMomentsContent, error) {
	builder := sqlbuilder.New().
		S(`UPDATE "landing_moments_content" SET "is_enabled" = $1, "landing_section_header_id" = $2, "updated_at" = NOW()`).
		S(`WHERE "id" = $3 AND "deleted_at" IS NULL`, landingMomentsContent.IsEnabled, landingMomentsContent.LandingSectionHeaderId, landingMomentsContent.Id).
		S(`RETURNING "id", "is_enabled", "landing_section_header_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingMomentsContent.Id, &landingMomentsContent.IsEnabled, &landingMomentsContent.LandingSectionHeaderId,
		&landingMomentsContent.CreatedAt, &landingMomentsContent.UpdatedAt, &landingMomentsContent.DeletedAt,
	)

	return landingMomentsContent, err
}

func (r landingMomentsContentRepositoryPostgresql) Delete(ctx context.Context) (entity.LandingMomentsContent, error) {
	landingMomentsContent := entity.LandingMomentsContent{}

	builder := sqlbuilder.New().
		S(`DELETE FROM "landing_moments_content"`).
		S(`RETURNING "id", "is_enabled", "landing_section_header_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingMomentsContent.Id, &landingMomentsContent.IsEnabled, &landingMomentsContent.LandingSectionHeaderId,
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
	builder := sqlbuilder.New().
		S(`INSERT INTO "landing_moments_content_images" ("is_enabled", "image_id", "created_at", "updated_at")`).
		S(`VALUES`)

	for i, image := range landingMomentsContentImages {
		if i > 0 {
			builder.SA(",")
		}
		builder.SA(`(?, ?, NOW(), NOW())`, image.IsEnabled, image.ImageId)
	}

	builder.S(`RETURNING "is_enabled", "image_id", "created_at", "updated_at", "deleted_at"`)

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	createdImages := []entity.LandingMomentsContentImage{}
	for rows.Next() {
		image := entity.LandingMomentsContentImage{}
		err := rows.Scan(&image.IsEnabled, &image.ImageId, &image.CreatedAt, &image.UpdatedAt, &image.DeletedAt)
		if err != nil {
			return nil, err
		}
		createdImages = append(createdImages, image)
	}

	return createdImages, nil
}

func (r landingMomentsContentImageRepositoryPostgresql) FindAll(ctx context.Context, opt repository.FindAllOptions) ([]entity.LandingMomentsContentImage, error) {
	builder := sqlbuilder.New().
		S(`SELECT "is_enabled", "image_id", "created_at", "updated_at", "deleted_at"`).
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
		err := rows.Scan(&image.IsEnabled, &image.ImageId, &image.CreatedAt, &image.UpdatedAt, &image.DeletedAt)
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
		S(`RETURNING "is_enabled", "image_id", "created_at", "updated_at", "deleted_at"`)

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	deletedImages := []entity.LandingMomentsContentImage{}
	for rows.Next() {
		image := entity.LandingMomentsContentImage{}
		err := rows.Scan(&image.IsEnabled, &image.ImageId, &image.CreatedAt, &image.UpdatedAt, &image.DeletedAt)
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
		S(`INSERT INTO "landing_affiliates_content" ("is_enabled", "landing_section_header_id", "created_at", "updated_at", "deleted_at")`).
		S(`VALUES ($1, $2, NOW(), NOW(), NULL)`, landingAffiliatesContent.IsEnabled, landingAffiliatesContent.LandingSectionHeaderId).
		S(`RETURNING "id", "is_enabled", "landing_section_header_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingAffiliatesContent.Id, &landingAffiliatesContent.IsEnabled, &landingAffiliatesContent.LandingSectionHeaderId,
		&landingAffiliatesContent.CreatedAt, &landingAffiliatesContent.UpdatedAt, &landingAffiliatesContent.DeletedAt,
	)

	return landingAffiliatesContent, err
}

func (r landingAffiliatesContentRepositoryPostgresql) Find(ctx context.Context) (entity.LandingAffiliatesContent, error) {
	landingAffiliatesContent := entity.LandingAffiliatesContent{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "is_enabled", "landing_section_header_id", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "landing_affiliates_content" WHERE "deleted_at" IS NULL`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingAffiliatesContent.Id, &landingAffiliatesContent.IsEnabled, &landingAffiliatesContent.LandingSectionHeaderId,
		&landingAffiliatesContent.CreatedAt, &landingAffiliatesContent.UpdatedAt, &landingAffiliatesContent.DeletedAt,
	)

	return landingAffiliatesContent, err
}

func (r landingAffiliatesContentRepositoryPostgresql) Update(ctx context.Context, landingAffiliatesContent entity.LandingAffiliatesContent) (entity.LandingAffiliatesContent, error) {
	builder := sqlbuilder.New().
		S(`UPDATE "landing_affiliates_content" SET "is_enabled" = $1, "landing_section_header_id" = $2, "updated_at" = NOW()`).
		S(`WHERE "id" = $3 AND "deleted_at" IS NULL`, landingAffiliatesContent.IsEnabled, landingAffiliatesContent.LandingSectionHeaderId, landingAffiliatesContent.Id).
		S(`RETURNING "id", "is_enabled", "landing_section_header_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingAffiliatesContent.Id, &landingAffiliatesContent.IsEnabled, &landingAffiliatesContent.LandingSectionHeaderId,
		&landingAffiliatesContent.CreatedAt, &landingAffiliatesContent.UpdatedAt, &landingAffiliatesContent.DeletedAt,
	)

	return landingAffiliatesContent, err
}

func (r landingAffiliatesContentRepositoryPostgresql) Delete(ctx context.Context) (entity.LandingAffiliatesContent, error) {
	landingAffiliatesContent := entity.LandingAffiliatesContent{}

	builder := sqlbuilder.New().
		S(`DELETE FROM "landing_affiliates_content"`).
		S(`RETURNING "id", "is_enabled", "landing_section_header_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingAffiliatesContent.Id, &landingAffiliatesContent.IsEnabled, &landingAffiliatesContent.LandingSectionHeaderId,
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
	builder := sqlbuilder.New().
		S(`INSERT INTO "landing_affiliates_content_affiliates" ("is_enabled", "name", "logo_id", "width", "height", "created_at", "updated_at") VALUES`)

	for i, affiliate := range landingAffiliatesContentAffiliates {
		if i > 0 {
			builder.SA(",")
		}
		builder.SA(`(?, ?, ?, ?, ?, NOW(), NOW())`, affiliate.IsEnabled, affiliate.Name, affiliate.LogoId, affiliate.Width, affiliate.Height)
	}

	builder.S(`RETURNING "id", "is_enabled", "name", "logo_id", "width", "height", "created_at", "updated_at", "deleted_at"`)

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	createdAffiliates := []entity.LandingAffiliatesContentAffiliate{}
	for rows.Next() {
		affiliate := entity.LandingAffiliatesContentAffiliate{}
		err := rows.Scan(
			&affiliate.Id, &affiliate.IsEnabled, &affiliate.Name, &affiliate.LogoId, &affiliate.Width, &affiliate.Height,
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
		S(`SELECT "id", "is_enabled", "name", "logo_id", "width", "height", "created_at", "updated_at", "deleted_at"`).
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
			&affiliate.Id, &affiliate.IsEnabled, &affiliate.Name, &affiliate.LogoId, &affiliate.Width, &affiliate.Height,
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
		S(`RETURNING "id", "is_enabled", "name", "logo_id", "width", "height", "created_at", "updated_at", "deleted_at"`)

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	deletedAffiliates := []entity.LandingAffiliatesContentAffiliate{}
	for rows.Next() {
		affiliate := entity.LandingAffiliatesContentAffiliate{}
		err := rows.Scan(
			&affiliate.Id, &affiliate.IsEnabled, &affiliate.Name, &affiliate.LogoId, &affiliate.Width, &affiliate.Height,
			&affiliate.CreatedAt, &affiliate.UpdatedAt, &affiliate.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		deletedAffiliates = append(deletedAffiliates, affiliate)
	}

	return deletedAffiliates, nil
}

type landingFaqContentRepositoryPostgresql struct {
	db DB
}

func NewLandingFaqContentRepository(db DB) repository.LandingFaqContentRepository {
	return landingFaqContentRepositoryPostgresql{db}
}

func (r landingFaqContentRepositoryPostgresql) Create(ctx context.Context, landingFaqContent entity.LandingFaqContent) (entity.LandingFaqContent, error) {
	builder := sqlbuilder.New().
		S(`INSERT INTO "landing_faq_content" ("is_enabled", "landing_section_header_id", "created_at", "updated_at", "deleted_at")`).
		S(`VALUES ($1, $2, NOW(), NOW(), NULL)`, landingFaqContent.IsEnabled, landingFaqContent.LandingSectionHeaderId).
		S(`RETURNING "id", "is_enabled", "landing_section_header_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingFaqContent.Id, &landingFaqContent.IsEnabled, &landingFaqContent.LandingSectionHeaderId,
		&landingFaqContent.CreatedAt, &landingFaqContent.UpdatedAt, &landingFaqContent.DeletedAt,
	)

	return landingFaqContent, err
}

func (r landingFaqContentRepositoryPostgresql) Find(ctx context.Context) (entity.LandingFaqContent, error) {
	landingFaqContent := entity.LandingFaqContent{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "is_enabled", "landing_section_header_id", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "landing_faq_content" WHERE "deleted_at" IS NULL`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingFaqContent.Id, &landingFaqContent.IsEnabled, &landingFaqContent.LandingSectionHeaderId,
		&landingFaqContent.CreatedAt, &landingFaqContent.UpdatedAt, &landingFaqContent.DeletedAt,
	)

	return landingFaqContent, err
}

func (r landingFaqContentRepositoryPostgresql) Update(ctx context.Context, landingFaqContent entity.LandingFaqContent) (entity.LandingFaqContent, error) {
	builder := sqlbuilder.New().
		S(`UPDATE "landing_faq_content" SET "is_enabled" = $1, "landing_section_header_id" = $2, "updated_at" = NOW()`).
		S(`WHERE "id" = $3 AND "deleted_at" IS NULL`, landingFaqContent.IsEnabled, landingFaqContent.LandingSectionHeaderId, landingFaqContent.Id).
		S(`RETURNING "id", "is_enabled", "landing_section_header_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingFaqContent.Id, &landingFaqContent.IsEnabled, &landingFaqContent.LandingSectionHeaderId,
		&landingFaqContent.CreatedAt, &landingFaqContent.UpdatedAt, &landingFaqContent.DeletedAt,
	)

	return landingFaqContent, err
}

func (r landingFaqContentRepositoryPostgresql) Delete(ctx context.Context) (entity.LandingFaqContent, error) {
	landingFaqContent := entity.LandingFaqContent{}

	builder := sqlbuilder.New().
		S(`DELETE FROM "landing_faq_content"`).
		S(`RETURNING "id", "is_enabled", "landing_section_header_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&landingFaqContent.Id, &landingFaqContent.IsEnabled, &landingFaqContent.LandingSectionHeaderId,
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
	builder := sqlbuilder.New().
		S(`INSERT INTO "landing_faq_content_faqs" ("is_enabled", "question", "answer", "created_at", "updated_at") VALUES`)

	for i, faq := range landingFaqContentFaqs {
		if i > 0 {
			builder.SA(",")
		}
		builder.SA(`(?, ?, ?, NOW(), NOW())`, faq.IsEnabled, faq.Question, faq.Answer)
	}

	builder.S(`RETURNING "id", "is_enabled", "question", "answer", "created_at", "updated_at", "deleted_at"`)

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	createdFaqs := []entity.LandingFaqContentFaq{}
	for rows.Next() {
		faq := entity.LandingFaqContentFaq{}
		err := rows.Scan(
			&faq.Id, &faq.IsEnabled, &faq.Question, &faq.Answer,
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
		S(`SELECT "id", "is_enabled", "question", "answer", "created_at", "updated_at", "deleted_at"`).
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
			&faq.Id, &faq.IsEnabled, &faq.Question, &faq.Answer,
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
		S(`RETURNING "id", "is_enabled", "name", "logo_id", "width", "height", "created_at", "updated_at", "deleted_at"`)

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	deletedFaqs := []entity.LandingFaqContentFaq{}
	for rows.Next() {
		faq := entity.LandingFaqContentFaq{}
		err := rows.Scan(
			&faq.Id, &faq.IsEnabled, &faq.Question, &faq.Answer,
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
	builder := sqlbuilder.New().
		S(`INSERT INTO "landing_menus" ("is_enabled", "icon", "label", "path", "created_at", "updated_at") VALUES`)

	for i, menu := range landingMenus {
		if i > 0 {
			builder.SA(",")
		}
		builder.SA(`(?, ?, ?, ?, NOW(), NOW())`, menu.IsEnabled, menu.Icon, menu.Label, menu.Path)
	}

	builder.S(`RETURNING "id", "is_enabled", "icon", "label", "path", "created_at", "updated_at", "deleted_at"`)

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	createdMenus := []entity.LandingMenu{}
	for rows.Next() {
		menu := entity.LandingMenu{}
		err := rows.Scan(
			&menu.Id, &menu.IsEnabled, &menu.Icon, &menu.Label, &menu.Path,
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
		S(`SELECT "id", "is_enabled", "icon", "label", "path", "created_at", "updated_at", "deleted_at"`).
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
			&menu.Id, &menu.IsEnabled, &menu.Icon, &menu.Label, &menu.Path,
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
		S(`RETURNING "id", "is_enabled", "name", "logo_id", "width", "height", "created_at", "updated_at", "deleted_at"`)

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	deletedMenus := []entity.LandingMenu{}
	for rows.Next() {
		menu := entity.LandingMenu{}
		err := rows.Scan(
			&menu.Id, &menu.IsEnabled, &menu.Icon, &menu.Label, &menu.Path,
			&menu.CreatedAt, &menu.UpdatedAt, &menu.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		deletedMenus = append(deletedMenus, menu)
	}

	return deletedMenus, nil
}
