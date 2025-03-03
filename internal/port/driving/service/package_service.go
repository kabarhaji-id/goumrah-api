package service

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type PackageService interface {
	CreatePackage(ctx context.Context, request dto.PackageRequest) (dto.PackageResponse, error)
	GetPackageById(ctx context.Context, id int64) (dto.PackageResponse, error)
	GetAllPackage(ctx context.Context, request dto.GetAllPackageRequest) ([]dto.PackageListResponse, error)
	UpdatePackage(ctx context.Context, id int64, request dto.PackageRequest) (dto.PackageResponse, error)
	DeletePackage(ctx context.Context, id int64) (dto.PackageResponse, error)
}
