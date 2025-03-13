package service

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type PackageSessionService interface {
	CreatePackageSession(ctx context.Context, request dto.PackageSessionRequest) (dto.PackageSessionResponse, error)
	GetPackageSessionById(ctx context.Context, id int64) (dto.PackageSessionResponse, error)
	GetAllPackageSession(ctx context.Context, request dto.GetAllPackageSessionRequest) ([]dto.PackageSessionListResponse, error)
	UpdatePackageSession(ctx context.Context, id int64, request dto.PackageSessionRequest) (dto.PackageSessionResponse, error)
	DeletePackageSession(ctx context.Context, id int64) (dto.PackageSessionResponse, error)
}
