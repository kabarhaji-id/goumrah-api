package postgresql

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driven/repository"
	"github.com/kabarhaji-id/goumrah-api/pkg/sqlbuilder"
)

type itineraryWidgetRepositoryPostgresql struct {
	db DB
}

func NewItineraryWidgetRepository(db DB) repository.ItineraryWidgetRepository {
	return itineraryWidgetRepositoryPostgresql{db}
}

func (r itineraryWidgetRepositoryPostgresql) Create(ctx context.Context, itineraryWidget entity.ItineraryWidget) (entity.ItineraryWidget, error) {
	builder := sqlbuilder.New().
		S(`INSERT INTO "itinerary_widgets" ("activity_id", "hotel_id", "information_id", "transport_id", "recommendation_id", "next_id", "created_at", "updated_at", "deleted_at")`).
		S(`VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW(), NULL)`, itineraryWidget.ActivityId, itineraryWidget.HotelId, itineraryWidget.InformationId, itineraryWidget.TransportId, itineraryWidget.RecommendationId, itineraryWidget.NextId).
		S(`RETURNING "id", "activity_id", "hotel_id", "information_id", "transport_id", "recommendation_id", "next_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&itineraryWidget.Id, &itineraryWidget.ActivityId, &itineraryWidget.HotelId, &itineraryWidget.InformationId, &itineraryWidget.TransportId, &itineraryWidget.RecommendationId, &itineraryWidget.NextId,
		&itineraryWidget.CreatedAt, &itineraryWidget.UpdatedAt, &itineraryWidget.DeletedAt,
	)

	return itineraryWidget, err
}

func (r itineraryWidgetRepositoryPostgresql) FindById(ctx context.Context, id int64) (entity.ItineraryWidget, error) {
	itineraryWidget := entity.ItineraryWidget{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "activity_id", "hotel_id", "information_id", "transport_id", "recommendation_id", "next_id", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "itinerary_widgets" WHERE "id" = $1 AND "deleted_at" IS NULL`, id)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&itineraryWidget.Id, &itineraryWidget.ActivityId, &itineraryWidget.HotelId, &itineraryWidget.InformationId, &itineraryWidget.TransportId, &itineraryWidget.RecommendationId, &itineraryWidget.NextId,
		&itineraryWidget.CreatedAt, &itineraryWidget.UpdatedAt, &itineraryWidget.DeletedAt,
	)

	return itineraryWidget, err
}

func (r itineraryWidgetRepositoryPostgresql) FindAll(ctx context.Context, opt repository.FindAllOptions) ([]entity.ItineraryWidget, error) {
	itineraries := []entity.ItineraryWidget{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "activity_id", "hotel_id", "information_id", "transport_id", "recommendation_id", "next_id", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "itinerary_widgets" WHERE "deleted_at" IS NULL`).
		S(`ORDER BY "id" ASC`)
	if opt.Limit.Valid {
		builder.SA(`LIMIT ?`, opt.Limit)
	}
	if opt.Offset.Valid {
		builder.SA(`OFFSET ?`, opt.Offset)
	}

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return itineraries, err
	}

	for rows.Next() {
		itineraryWidget := entity.ItineraryWidget{}
		err = rows.Scan(
			&itineraryWidget.Id, &itineraryWidget.ActivityId, &itineraryWidget.HotelId, &itineraryWidget.InformationId, &itineraryWidget.TransportId, &itineraryWidget.RecommendationId, &itineraryWidget.NextId,
			&itineraryWidget.CreatedAt, &itineraryWidget.UpdatedAt, &itineraryWidget.DeletedAt,
		)
		if err != nil {
			return itineraries, err
		}

		itineraries = append(itineraries, itineraryWidget)
	}

	return itineraries, nil
}

func (r itineraryWidgetRepositoryPostgresql) Update(ctx context.Context, id int64, itineraryWidget entity.ItineraryWidget) (entity.ItineraryWidget, error) {
	builder := sqlbuilder.New().
		S(
			`UPDATE "itinerary_widgets" SET "activity_id" = $1, "hotel_id" = $2, "information_id" = $3, "transport_id" = $4, "recommendation_id" = $5, "next_id" = $6, "updated_at" = NOW()`,
			itineraryWidget.ActivityId, itineraryWidget.HotelId, itineraryWidget.InformationId, itineraryWidget.TransportId, itineraryWidget.RecommendationId, itineraryWidget.NextId,
		).
		S(`WHERE "id" = $7 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "activity_id", "hotel_id", "information_id", "transport_id", "recommendation_id", "next_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&itineraryWidget.Id, &itineraryWidget.ActivityId, &itineraryWidget.HotelId, &itineraryWidget.InformationId, &itineraryWidget.TransportId, &itineraryWidget.RecommendationId, &itineraryWidget.NextId,
		&itineraryWidget.CreatedAt, &itineraryWidget.UpdatedAt, &itineraryWidget.DeletedAt,
	)

	return itineraryWidget, err
}

func (r itineraryWidgetRepositoryPostgresql) Delete(ctx context.Context, id int64) (entity.ItineraryWidget, error) {
	itineraryWidget := entity.ItineraryWidget{}

	builder := sqlbuilder.New().
		S(`UPDATE "itinerary_widgets" SET "deleted_at" = NOW()`).
		S(`WHERE "id" = $1 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "activity_id", "hotel_id", "information_id", "transport_id", "recommendation_id", "next_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&itineraryWidget.Id, &itineraryWidget.ActivityId, &itineraryWidget.HotelId, &itineraryWidget.InformationId, &itineraryWidget.TransportId, &itineraryWidget.RecommendationId, &itineraryWidget.NextId,
		&itineraryWidget.CreatedAt, &itineraryWidget.UpdatedAt, &itineraryWidget.DeletedAt,
	)

	return itineraryWidget, err
}

type itineraryWidgetActivityRepositoryPostgresql struct {
	db DB
}

func NewItineraryWidgetActivityRepository(db DB) repository.ItineraryWidgetActivityRepository {
	return itineraryWidgetActivityRepositoryPostgresql{db}
}

func (r itineraryWidgetActivityRepositoryPostgresql) Create(ctx context.Context, itineraryWidgetActivity entity.ItineraryWidgetActivity) (entity.ItineraryWidgetActivity, error) {
	builder := sqlbuilder.New().
		S(`INSERT INTO "itinerary_widget_activities" ("title", "description", "created_at", "updated_at", "deleted_at")`).
		S(`VALUES ($1, $2, NOW(), NOW(), NULL)`, itineraryWidgetActivity.Title, itineraryWidgetActivity.Description).
		S(`RETURNING "id", "title", "description", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&itineraryWidgetActivity.Id, &itineraryWidgetActivity.Title, &itineraryWidgetActivity.Description,
		&itineraryWidgetActivity.CreatedAt, &itineraryWidgetActivity.UpdatedAt, &itineraryWidgetActivity.DeletedAt,
	)

	return itineraryWidgetActivity, err
}

func (r itineraryWidgetActivityRepositoryPostgresql) FindById(ctx context.Context, id int64) (entity.ItineraryWidgetActivity, error) {
	itineraryWidgetActivity := entity.ItineraryWidgetActivity{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "title", "description", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "itinerary_widget_activities" WHERE "id" = $1 AND "deleted_at" IS NULL`, id)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&itineraryWidgetActivity.Id, &itineraryWidgetActivity.Title, &itineraryWidgetActivity.Description,
		&itineraryWidgetActivity.CreatedAt, &itineraryWidgetActivity.UpdatedAt, &itineraryWidgetActivity.DeletedAt,
	)

	return itineraryWidgetActivity, err
}

func (r itineraryWidgetActivityRepositoryPostgresql) FindAll(ctx context.Context, opt repository.FindAllOptions) ([]entity.ItineraryWidgetActivity, error) {
	itineraries := []entity.ItineraryWidgetActivity{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "title", "description", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "itinerary_widget_activities" WHERE "deleted_at" IS NULL`).
		S(`ORDER BY "id" ASC`)
	if opt.Limit.Valid {
		builder.SA(`LIMIT ?`, opt.Limit)
	}
	if opt.Offset.Valid {
		builder.SA(`OFFSET ?`, opt.Offset)
	}

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return itineraries, err
	}

	for rows.Next() {
		itineraryWidgetActivity := entity.ItineraryWidgetActivity{}
		err = rows.Scan(
			&itineraryWidgetActivity.Id, &itineraryWidgetActivity.Title, &itineraryWidgetActivity.Description,
			&itineraryWidgetActivity.CreatedAt, &itineraryWidgetActivity.UpdatedAt, &itineraryWidgetActivity.DeletedAt,
		)
		if err != nil {
			return itineraries, err
		}

		itineraries = append(itineraries, itineraryWidgetActivity)
	}

	return itineraries, nil
}

func (r itineraryWidgetActivityRepositoryPostgresql) Update(ctx context.Context, id int64, itineraryWidgetActivity entity.ItineraryWidgetActivity) (entity.ItineraryWidgetActivity, error) {
	builder := sqlbuilder.New().
		S(
			`UPDATE "itinerary_widget_activities" SET "title" = $1, "description" = $2, "updated_at" = NOW()`,
			itineraryWidgetActivity.Title, itineraryWidgetActivity.Description,
		).
		S(`WHERE "id" = $3 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "title", "description", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&itineraryWidgetActivity.Id, &itineraryWidgetActivity.Title, &itineraryWidgetActivity.Description,
		&itineraryWidgetActivity.CreatedAt, &itineraryWidgetActivity.UpdatedAt, &itineraryWidgetActivity.DeletedAt,
	)

	return itineraryWidgetActivity, err
}

func (r itineraryWidgetActivityRepositoryPostgresql) Delete(ctx context.Context, id int64) (entity.ItineraryWidgetActivity, error) {
	itineraryWidgetActivity := entity.ItineraryWidgetActivity{}

	builder := sqlbuilder.New().
		S(`UPDATE "itinerary_widget_activities" SET "deleted_at" = NOW()`).
		S(`WHERE "id" = $1 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "title", "description", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&itineraryWidgetActivity.Id, &itineraryWidgetActivity.Title, &itineraryWidgetActivity.Description,
		&itineraryWidgetActivity.CreatedAt, &itineraryWidgetActivity.UpdatedAt, &itineraryWidgetActivity.DeletedAt,
	)

	return itineraryWidgetActivity, err
}

func (r itineraryWidgetActivityRepositoryPostgresql) AttachImages(ctx context.Context, id int64, imageIds []int64) ([]int64, error) {
	if len(imageIds) == 0 {
		return []int64{}, nil
	}

	builder := sqlbuilder.New().
		S(`INSERT INTO "itinerary_widget_activity_images" ("itinerary_widget_activity_id", "image_id", "created_at", "updated_at", "deleted_at") VALUES`)

	imageIdsLen := len(imageIds)
	for index, imageId := range imageIds {
		builder.SA(`(?, ?, NOW(), NOW(), NULL)`, id, imageId)
		if index+1 < imageIdsLen {
			builder.S(",")
		}
	}

	_, err := r.db.Exec(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return nil, err
	}

	return imageIds, nil
}

func (r itineraryWidgetActivityRepositoryPostgresql) FindImages(ctx context.Context, id int64) ([]entity.Image, error) {
	images := []entity.Image{}

	builder := sqlbuilder.New().
		S(`SELECT "images"."id", "images"."src", "images"."alt", "images"."category", "images"."title", "images"."created_at", "images"."updated_at", "images"."deleted_at"`).
		S(`FROM "itinerary_widget_activity_images"`).
		S(`INNER JOIN "itinerary_widget_activities" ON "itinerary_widget_activities"."id" = "itinerary_widget_activity_images"."itinerary_widget_activity_id"`).
		S(`INNER JOIN "images" ON "images"."id" = "itinerary_widget_activity_images"."image_id"`).
		S(`WHERE "itinerary_widget_activity_images"."itinerary_widget_activity_id" = $1 AND "itinerary_widget_activity_images"."deleted_at" IS NULL AND "itinerary_widget_activities"."deleted_at" IS NULL AND "images"."deleted_at" IS NULL`, id)

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return images, err
	}

	for rows.Next() {
		image := entity.Image{}
		err = rows.Scan(
			&image.Id, &image.Src, &image.Alt, &image.Category, &image.Title,
			&image.CreatedAt, &image.UpdatedAt, &image.DeletedAt,
		)
		if err != nil {
			return images, err
		}

		images = append(images, image)
	}

	return images, nil
}

func (r itineraryWidgetActivityRepositoryPostgresql) FindImageIds(ctx context.Context, id int64) ([]int64, error) {
	imageIds := []int64{}

	builder := sqlbuilder.New().
		S(`SELECT "images"."id"`).
		S(`FROM "itinerary_widget_activity_images"`).
		S(`INNER JOIN "itinerary_widget_activities" ON "itinerary_widget_activities"."id" = "itinerary_widget_activity_images"."itinerary_widget_activity_id"`).
		S(`INNER JOIN "images" ON "images"."id" = "itinerary_widget_activity_images"."image_id"`).
		S(`WHERE "itinerary_widget_activity_images"."itinerary_widget_activity_id" = $1 AND "itinerary_widget_activity_images"."deleted_at" IS NULL AND "itinerary_widget_activities"."deleted_at" IS NULL AND "images"."deleted_at" IS NULL`, id)

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return imageIds, err
	}

	for rows.Next() {
		var imageId int64
		err = rows.Scan(&imageId)
		if err != nil {
			return imageIds, err
		}

		imageIds = append(imageIds, imageId)
	}

	return imageIds, nil
}

func (r itineraryWidgetActivityRepositoryPostgresql) DetachImages(ctx context.Context, id int64) ([]int64, error) {
	builder := sqlbuilder.New().
		S(`DELETE FROM "itinerary_widget_activity_images"`).
		S(`WHERE "itinerary_widget_activity_id" = $1`, id).
		S(`RETURNING "image_id"`)

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	imageIds := []int64{}
	for rows.Next() {
		var imageId int64
		if err = rows.Scan(&imageId); err != nil {
			return nil, err
		}

		imageIds = append(imageIds, imageId)
	}

	return imageIds, nil
}

type itineraryWidgetHotelRepositoryPostgresql struct {
	db DB
}

func NewItineraryWidgetHotelRepository(db DB) repository.ItineraryWidgetHotelRepository {
	return itineraryWidgetHotelRepositoryPostgresql{db}
}

func (r itineraryWidgetHotelRepositoryPostgresql) Create(ctx context.Context, itineraryWidgetHotel entity.ItineraryWidgetHotel) (entity.ItineraryWidgetHotel, error) {
	builder := sqlbuilder.New().
		S(`INSERT INTO "itinerary_widget_hotels" ("hotel_id", "created_at", "updated_at", "deleted_at")`).
		S(`VALUES ($1, NOW(), NOW(), NULL)`, itineraryWidgetHotel.HotelId).
		S(`RETURNING "id", "hotel_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&itineraryWidgetHotel.Id, &itineraryWidgetHotel.HotelId,
		&itineraryWidgetHotel.CreatedAt, &itineraryWidgetHotel.UpdatedAt, &itineraryWidgetHotel.DeletedAt,
	)

	return itineraryWidgetHotel, err
}

func (r itineraryWidgetHotelRepositoryPostgresql) FindById(ctx context.Context, id int64) (entity.ItineraryWidgetHotel, error) {
	itineraryWidgetHotel := entity.ItineraryWidgetHotel{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "hotel_id", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "itinerary_widget_hotels" WHERE "id" = $1 AND "deleted_at" IS NULL`, id)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&itineraryWidgetHotel.Id, &itineraryWidgetHotel.HotelId,
		&itineraryWidgetHotel.CreatedAt, &itineraryWidgetHotel.UpdatedAt, &itineraryWidgetHotel.DeletedAt,
	)

	return itineraryWidgetHotel, err
}

func (r itineraryWidgetHotelRepositoryPostgresql) FindAll(ctx context.Context, opt repository.FindAllOptions) ([]entity.ItineraryWidgetHotel, error) {
	itineraryWidgetHotels := []entity.ItineraryWidgetHotel{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "hotel_id", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "itinerary_widget_hotels" WHERE "deleted_at" IS NULL`).
		S(`ORDER BY "id" ASC`)
	if opt.Limit.Valid {
		builder.SA(`LIMIT ?`, opt.Limit)
	}
	if opt.Offset.Valid {
		builder.SA(`OFFSET ?`, opt.Offset)
	}

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return itineraryWidgetHotels, err
	}

	for rows.Next() {
		itineraryWidgetHotel := entity.ItineraryWidgetHotel{}
		err = rows.Scan(
			&itineraryWidgetHotel.Id, &itineraryWidgetHotel.HotelId,
			&itineraryWidgetHotel.CreatedAt, &itineraryWidgetHotel.UpdatedAt, &itineraryWidgetHotel.DeletedAt,
		)
		if err != nil {
			return itineraryWidgetHotels, err
		}

		itineraryWidgetHotels = append(itineraryWidgetHotels, itineraryWidgetHotel)
	}

	return itineraryWidgetHotels, nil
}

func (r itineraryWidgetHotelRepositoryPostgresql) Update(ctx context.Context, id int64, itineraryWidgetHotel entity.ItineraryWidgetHotel) (entity.ItineraryWidgetHotel, error) {
	builder := sqlbuilder.New().
		S(
			`UPDATE "itinerary_widget_hotels" SET "hotel_id" = $1, "updated_at" = NOW()`,
			itineraryWidgetHotel.HotelId,
		).
		S(`WHERE "id" = $2 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "title", "description", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&itineraryWidgetHotel.Id, &itineraryWidgetHotel.HotelId,
		&itineraryWidgetHotel.CreatedAt, &itineraryWidgetHotel.UpdatedAt, &itineraryWidgetHotel.DeletedAt,
	)

	return itineraryWidgetHotel, err
}

func (r itineraryWidgetHotelRepositoryPostgresql) Delete(ctx context.Context, id int64) (entity.ItineraryWidgetHotel, error) {
	itineraryWidgetHotel := entity.ItineraryWidgetHotel{}

	builder := sqlbuilder.New().
		S(`UPDATE "itinerary_widget_hotels" SET "deleted_at" = NOW()`).
		S(`WHERE "id" = $1 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "hotel_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&itineraryWidgetHotel.Id, &itineraryWidgetHotel.HotelId,
		&itineraryWidgetHotel.CreatedAt, &itineraryWidgetHotel.UpdatedAt, &itineraryWidgetHotel.DeletedAt,
	)

	return itineraryWidgetHotel, err
}

type itineraryWidgetInformationRepositoryPostgresql struct {
	db DB
}

func NewItineraryWidgetInformationRepository(db DB) repository.ItineraryWidgetInformationRepository {
	return itineraryWidgetInformationRepositoryPostgresql{db}
}

func (r itineraryWidgetInformationRepositoryPostgresql) Create(ctx context.Context, itineraryWidgetInformation entity.ItineraryWidgetInformation) (entity.ItineraryWidgetInformation, error) {
	builder := sqlbuilder.New().
		S(`INSERT INTO "itinerary_widget_informations" ("description", "created_at", "updated_at", "deleted_at")`).
		S(`VALUES ($1, NOW(), NOW(), NULL)`, itineraryWidgetInformation.Description).
		S(`RETURNING "id", "description", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&itineraryWidgetInformation.Id, &itineraryWidgetInformation.Description,
		&itineraryWidgetInformation.CreatedAt, &itineraryWidgetInformation.UpdatedAt, &itineraryWidgetInformation.DeletedAt,
	)

	return itineraryWidgetInformation, err
}

func (r itineraryWidgetInformationRepositoryPostgresql) FindById(ctx context.Context, id int64) (entity.ItineraryWidgetInformation, error) {
	itineraryWidgetInformation := entity.ItineraryWidgetInformation{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "description", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "itinerary_widget_informations" WHERE "id" = $1 AND "deleted_at" IS NULL`, id)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&itineraryWidgetInformation.Id, &itineraryWidgetInformation.Description,
		&itineraryWidgetInformation.CreatedAt, &itineraryWidgetInformation.UpdatedAt, &itineraryWidgetInformation.DeletedAt,
	)

	return itineraryWidgetInformation, err
}

func (r itineraryWidgetInformationRepositoryPostgresql) FindAll(ctx context.Context, opt repository.FindAllOptions) ([]entity.ItineraryWidgetInformation, error) {
	itineraryWidgetInformations := []entity.ItineraryWidgetInformation{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "description", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "itinerary_widget_informations" WHERE "deleted_at" IS NULL`).
		S(`ORDER BY "id" ASC`)
	if opt.Limit.Valid {
		builder.SA(`LIMIT ?`, opt.Limit)
	}
	if opt.Offset.Valid {
		builder.SA(`OFFSET ?`, opt.Offset)
	}

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return itineraryWidgetInformations, err
	}

	for rows.Next() {
		itineraryWidgetInformation := entity.ItineraryWidgetInformation{}
		err = rows.Scan(
			&itineraryWidgetInformation.Id, &itineraryWidgetInformation.Description,
			&itineraryWidgetInformation.CreatedAt, &itineraryWidgetInformation.UpdatedAt, &itineraryWidgetInformation.DeletedAt,
		)
		if err != nil {
			return itineraryWidgetInformations, err
		}

		itineraryWidgetInformations = append(itineraryWidgetInformations, itineraryWidgetInformation)
	}

	return itineraryWidgetInformations, nil
}

func (r itineraryWidgetInformationRepositoryPostgresql) Update(ctx context.Context, id int64, itineraryWidgetInformation entity.ItineraryWidgetInformation) (entity.ItineraryWidgetInformation, error) {
	builder := sqlbuilder.New().
		S(
			`UPDATE "itinerary_widget_informations" SET "description" = $1, "updated_at" = NOW()`,
			itineraryWidgetInformation.Description,
		).
		S(`WHERE "id" = $2 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "title", "description", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&itineraryWidgetInformation.Id, &itineraryWidgetInformation.Description,
		&itineraryWidgetInformation.CreatedAt, &itineraryWidgetInformation.UpdatedAt, &itineraryWidgetInformation.DeletedAt,
	)

	return itineraryWidgetInformation, err
}

func (r itineraryWidgetInformationRepositoryPostgresql) Delete(ctx context.Context, id int64) (entity.ItineraryWidgetInformation, error) {
	itineraryWidgetInformation := entity.ItineraryWidgetInformation{}

	builder := sqlbuilder.New().
		S(`UPDATE "itinerary_widget_informations" SET "deleted_at" = NOW()`).
		S(`WHERE "id" = $1 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "description", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&itineraryWidgetInformation.Id, &itineraryWidgetInformation.Description,
		&itineraryWidgetInformation.CreatedAt, &itineraryWidgetInformation.UpdatedAt, &itineraryWidgetInformation.DeletedAt,
	)

	return itineraryWidgetInformation, err
}

type itineraryWidgetTransportRepositoryPostgresql struct {
	db DB
}

func NewItineraryWidgetTransportRepository(db DB) repository.ItineraryWidgetTransportRepository {
	return itineraryWidgetTransportRepositoryPostgresql{db}
}

func (r itineraryWidgetTransportRepositoryPostgresql) Create(ctx context.Context, itineraryWidgetTransport entity.ItineraryWidgetTransport) (entity.ItineraryWidgetTransport, error) {
	builder := sqlbuilder.New().
		S(`INSERT INTO "itinerary_widget_transports" ("transportation", "from", "to", "created_at", "updated_at", "deleted_at")`).
		S(`VALUES ($1, $2, $3, NOW(), NOW(), NULL)`, itineraryWidgetTransport.Transportation, itineraryWidgetTransport.From, itineraryWidgetTransport.To).
		S(`RETURNING "id", "description", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&itineraryWidgetTransport.Id, &itineraryWidgetTransport.Transportation, &itineraryWidgetTransport.From, &itineraryWidgetTransport.To,
		&itineraryWidgetTransport.CreatedAt, &itineraryWidgetTransport.UpdatedAt, &itineraryWidgetTransport.DeletedAt,
	)

	return itineraryWidgetTransport, err
}

func (r itineraryWidgetTransportRepositoryPostgresql) FindById(ctx context.Context, id int64) (entity.ItineraryWidgetTransport, error) {
	itineraryWidgetTransport := entity.ItineraryWidgetTransport{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "transportation", "from", "to", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "itinerary_widget_transports" WHERE "id" = $1 AND "deleted_at" IS NULL`, id)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&itineraryWidgetTransport.Id, &itineraryWidgetTransport.Transportation, &itineraryWidgetTransport.From, &itineraryWidgetTransport.To,
		&itineraryWidgetTransport.CreatedAt, &itineraryWidgetTransport.UpdatedAt, &itineraryWidgetTransport.DeletedAt,
	)

	return itineraryWidgetTransport, err
}

func (r itineraryWidgetTransportRepositoryPostgresql) FindAll(ctx context.Context, opt repository.FindAllOptions) ([]entity.ItineraryWidgetTransport, error) {
	itineraryWidgetTransports := []entity.ItineraryWidgetTransport{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "transportation", "from", "to", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "itinerary_widget_transports" WHERE "deleted_at" IS NULL`).
		S(`ORDER BY "id" ASC`)
	if opt.Limit.Valid {
		builder.SA(`LIMIT ?`, opt.Limit)
	}
	if opt.Offset.Valid {
		builder.SA(`OFFSET ?`, opt.Offset)
	}

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return itineraryWidgetTransports, err
	}

	for rows.Next() {
		itineraryWidgetTransport := entity.ItineraryWidgetTransport{}
		err = rows.Scan(
			&itineraryWidgetTransport.Id, &itineraryWidgetTransport.Transportation, &itineraryWidgetTransport.From, &itineraryWidgetTransport.To,
			&itineraryWidgetTransport.CreatedAt, &itineraryWidgetTransport.UpdatedAt, &itineraryWidgetTransport.DeletedAt,
		)
		if err != nil {
			return itineraryWidgetTransports, err
		}

		itineraryWidgetTransports = append(itineraryWidgetTransports, itineraryWidgetTransport)
	}

	return itineraryWidgetTransports, nil
}

func (r itineraryWidgetTransportRepositoryPostgresql) Update(ctx context.Context, id int64, itineraryWidgetTransport entity.ItineraryWidgetTransport) (entity.ItineraryWidgetTransport, error) {
	builder := sqlbuilder.New().
		S(
			`UPDATE "itinerary_widget_transports" SET "transportation" = $1, "from" = $2, "to" = $3, "updated_at" = NOW()`,
			itineraryWidgetTransport.Transportation, itineraryWidgetTransport.From, itineraryWidgetTransport.To,
		).
		S(`WHERE "id" = $4 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "title", "description", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&itineraryWidgetTransport.Id, &itineraryWidgetTransport.Transportation, &itineraryWidgetTransport.From, &itineraryWidgetTransport.To,
		&itineraryWidgetTransport.CreatedAt, &itineraryWidgetTransport.UpdatedAt, &itineraryWidgetTransport.DeletedAt,
	)

	return itineraryWidgetTransport, err
}

func (r itineraryWidgetTransportRepositoryPostgresql) Delete(ctx context.Context, id int64) (entity.ItineraryWidgetTransport, error) {
	itineraryWidgetTransport := entity.ItineraryWidgetTransport{}

	builder := sqlbuilder.New().
		S(`UPDATE "itinerary_widget_transports" SET "deleted_at" = NOW()`).
		S(`WHERE "id" = $1 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "transportation", "from", "to", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&itineraryWidgetTransport.Id, &itineraryWidgetTransport.Transportation, &itineraryWidgetTransport.From, &itineraryWidgetTransport.To,
		&itineraryWidgetTransport.CreatedAt, &itineraryWidgetTransport.UpdatedAt, &itineraryWidgetTransport.DeletedAt,
	)

	return itineraryWidgetTransport, err
}

type itineraryWidgetRecommendationRepositoryPostgresql struct {
	db DB
}

func NewItineraryWidgetRecommendationRepository(db DB) repository.ItineraryWidgetRecommendationRepository {
	return itineraryWidgetRecommendationRepositoryPostgresql{db}
}

func (r itineraryWidgetRecommendationRepositoryPostgresql) Create(ctx context.Context, itineraryWidgetRecommendation entity.ItineraryWidgetRecommendation) (entity.ItineraryWidgetRecommendation, error) {
	builder := sqlbuilder.New().
		S(`INSERT INTO "itinerary_widget_recommendations" ("description", "created_at", "updated_at", "deleted_at")`).
		S(`VALUES ($1, NOW(), NOW(), NULL)`, itineraryWidgetRecommendation.Description).
		S(`RETURNING "id", "description", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&itineraryWidgetRecommendation.Id, &itineraryWidgetRecommendation.Description,
		&itineraryWidgetRecommendation.CreatedAt, &itineraryWidgetRecommendation.UpdatedAt, &itineraryWidgetRecommendation.DeletedAt,
	)

	return itineraryWidgetRecommendation, err
}

func (r itineraryWidgetRecommendationRepositoryPostgresql) FindById(ctx context.Context, id int64) (entity.ItineraryWidgetRecommendation, error) {
	itineraryWidgetRecommendation := entity.ItineraryWidgetRecommendation{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "description", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "itinerary_widget_recommendations" WHERE "id" = $1 AND "deleted_at" IS NULL`, id)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&itineraryWidgetRecommendation.Id, &itineraryWidgetRecommendation.Description,
		&itineraryWidgetRecommendation.CreatedAt, &itineraryWidgetRecommendation.UpdatedAt, &itineraryWidgetRecommendation.DeletedAt,
	)

	return itineraryWidgetRecommendation, err
}

func (r itineraryWidgetRecommendationRepositoryPostgresql) FindAll(ctx context.Context, opt repository.FindAllOptions) ([]entity.ItineraryWidgetRecommendation, error) {
	itineraryWidgetRecommendations := []entity.ItineraryWidgetRecommendation{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "description", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "itinerary_widget_recommendations" WHERE "deleted_at" IS NULL`).
		S(`ORDER BY "id" ASC`)
	if opt.Limit.Valid {
		builder.SA(`LIMIT ?`, opt.Limit)
	}
	if opt.Offset.Valid {
		builder.SA(`OFFSET ?`, opt.Offset)
	}

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return itineraryWidgetRecommendations, err
	}

	for rows.Next() {
		itineraryWidgetRecommendation := entity.ItineraryWidgetRecommendation{}
		err = rows.Scan(
			&itineraryWidgetRecommendation.Id, &itineraryWidgetRecommendation.Description,
			&itineraryWidgetRecommendation.CreatedAt, &itineraryWidgetRecommendation.UpdatedAt, &itineraryWidgetRecommendation.DeletedAt,
		)
		if err != nil {
			return itineraryWidgetRecommendations, err
		}

		itineraryWidgetRecommendations = append(itineraryWidgetRecommendations, itineraryWidgetRecommendation)
	}

	return itineraryWidgetRecommendations, nil
}

func (r itineraryWidgetRecommendationRepositoryPostgresql) Update(ctx context.Context, id int64, itineraryWidgetRecommendation entity.ItineraryWidgetRecommendation) (entity.ItineraryWidgetRecommendation, error) {
	builder := sqlbuilder.New().
		S(
			`UPDATE "itinerary_widget_recommendations" SET "description" = $1, "updated_at" = NOW()`,
			itineraryWidgetRecommendation.Description,
		).
		S(`WHERE "id" = $2 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "description", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&itineraryWidgetRecommendation.Id, &itineraryWidgetRecommendation.Description,
		&itineraryWidgetRecommendation.CreatedAt, &itineraryWidgetRecommendation.UpdatedAt, &itineraryWidgetRecommendation.DeletedAt,
	)

	return itineraryWidgetRecommendation, err
}

func (r itineraryWidgetRecommendationRepositoryPostgresql) Delete(ctx context.Context, id int64) (entity.ItineraryWidgetRecommendation, error) {
	itineraryWidgetRecommendation := entity.ItineraryWidgetRecommendation{}

	builder := sqlbuilder.New().
		S(`UPDATE "itinerary_widget_recommendations" SET "deleted_at" = NOW()`).
		S(`WHERE "id" = $1 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "description", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&itineraryWidgetRecommendation.Id, &itineraryWidgetRecommendation.Description,
		&itineraryWidgetRecommendation.CreatedAt, &itineraryWidgetRecommendation.UpdatedAt, &itineraryWidgetRecommendation.DeletedAt,
	)

	return itineraryWidgetRecommendation, err
}

func (r itineraryWidgetRecommendationRepositoryPostgresql) AttachImages(ctx context.Context, id int64, imageIds []int64) ([]int64, error) {
	if len(imageIds) == 0 {
		return []int64{}, nil
	}

	builder := sqlbuilder.New().
		S(`INSERT INTO "itinerary_widget_recommendation_images" ("itinerary_widget_recommendation_id", "image_id", "created_at", "updated_at", "deleted_at") VALUES`)

	imageIdsLen := len(imageIds)
	for index, imageId := range imageIds {
		builder.SA(`(?, ?, NOW(), NOW(), NULL)`, id, imageId)
		if index+1 < imageIdsLen {
			builder.S(",")
		}
	}

	_, err := r.db.Exec(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return nil, err
	}

	return imageIds, nil
}

func (r itineraryWidgetRecommendationRepositoryPostgresql) FindImages(ctx context.Context, id int64) ([]entity.Image, error) {
	images := []entity.Image{}

	builder := sqlbuilder.New().
		S(`SELECT "images"."id", "images"."src", "images"."alt", "images"."category", "images"."title", "images"."created_at", "images"."updated_at", "images"."deleted_at"`).
		S(`FROM "itinerary_widget_recommendation_images"`).
		S(`INNER JOIN "itinerary_widget_recommendations" ON "itinerary_widget_recommendations"."id" = "itinerary_widget_recommendation_images"."itinerary_widget_recommendation_id"`).
		S(`INNER JOIN "images" ON "images"."id" = "itinerary_widget_recommendation_images"."image_id"`).
		S(`WHERE "itinerary_widget_recommendation_images"."itinerary_widget_recommendation_id" = $1 AND "itinerary_widget_recommendation_images"."deleted_at" IS NULL AND "itinerary_widget_recommendations"."deleted_at" IS NULL AND "images"."deleted_at" IS NULL`, id)

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return images, err
	}

	for rows.Next() {
		image := entity.Image{}
		err = rows.Scan(
			&image.Id, &image.Src, &image.Alt, &image.Category, &image.Title,
			&image.CreatedAt, &image.UpdatedAt, &image.DeletedAt,
		)
		if err != nil {
			return images, err
		}

		images = append(images, image)
	}

	return images, nil
}

func (r itineraryWidgetRecommendationRepositoryPostgresql) FindImageIds(ctx context.Context, id int64) ([]int64, error) {
	imageIds := []int64{}

	builder := sqlbuilder.New().
		S(`SELECT "images"."id"`).
		S(`FROM "itinerary_widget_recommendation_images"`).
		S(`INNER JOIN "itinerary_widget_recommendations" ON "itinerary_widget_recommendations"."id" = "itinerary_widget_recommendation_images"."itinerary_widget_recommendation_id"`).
		S(`INNER JOIN "images" ON "images"."id" = "itinerary_widget_recommendation_images"."image_id"`).
		S(`WHERE "itinerary_widget_recommendation_images"."itinerary_widget_recommendation_id" = $1 AND "itinerary_widget_recommendation_images"."deleted_at" IS NULL AND "itinerary_widget_recommendations"."deleted_at" IS NULL AND "images"."deleted_at" IS NULL`, id)

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return imageIds, err
	}

	for rows.Next() {
		var imageId int64
		err = rows.Scan(&imageId)
		if err != nil {
			return imageIds, err
		}

		imageIds = append(imageIds, imageId)
	}

	return imageIds, nil
}

func (r itineraryWidgetRecommendationRepositoryPostgresql) DetachImages(ctx context.Context, id int64) ([]int64, error) {
	builder := sqlbuilder.New().
		S(`DELETE FROM "itinerary_widget_recommendation_images"`).
		S(`WHERE "itinerary_widget_recommendation_id" = $1`, id).
		S(`RETURNING "image_id"`)

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	imageIds := []int64{}
	for rows.Next() {
		var imageId int64
		if err = rows.Scan(&imageId); err != nil {
			return nil, err
		}

		imageIds = append(imageIds, imageId)
	}

	return imageIds, nil
}
