package validator

import (
	"context"
	"net/mail"

	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
	"github.com/kabarhaji-id/goumrah-api/pkg/regexpattern"
)

type UserValidator struct {
}

func NewUserValidator() UserValidator {
	return UserValidator{}
}

func (v UserValidator) ValidateRequest(ctx context.Context, request dto.UserRequest) error {
	fullNameLength := len(request.FullName)
	if fullNameLength < 1 {
		return newError("FullName", mustBeNotEmpty)
	}
	if fullNameLength > 100 {
		return newError("FullName", maxChars(100))
	}

	phoneNumberLength := len(request.PhoneNumber)
	if phoneNumberLength < 1 {
		return newError("PhoneNumber", mustBeNotEmpty)
	}
	if phoneNumberLength > 20 {
		return newError("PhoneNumber", maxChars(20))
	}
	if !regexpattern.PhoneNumber().MatchString(request.PhoneNumber) {
		return newError("PhoneNumber", invalidPhoneNumber)
	}

	emailLength := len(request.Email)
	if emailLength < 1 {
		return newError("Email", mustBeNotEmpty)
	}
	if emailLength > 256 {
		return newError("Email", maxChars(256))
	}
	if _, err := mail.ParseAddress(request.Email); err != nil {
		return newError("Email", invalidEmail)
	}

	addressLength := len(request.Address)
	if addressLength < 1 {
		return newError("Address", mustBeNotEmpty)
	}
	if addressLength > 500 {
		return newError("Address", maxChars(500))
	}

	return nil
}

func (v UserValidator) ValidateId(ctx context.Context, id int64) error {
	if id < 1 {
		return newError("id", mustBeGte(1))
	}

	return nil
}

func (v UserValidator) ValidateGetAllRequest(ctx context.Context, request dto.GetAllUserRequest) error {
	if request.Page < 1 {
		return newError("Page", mustBeGte(1))
	}

	if request.PerPage < 1 {
		return newError("PerPage", mustBeGte(1))
	}

	return nil
}
