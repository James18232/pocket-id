package onetimeaccess

import (
	"github.com/pocket-id/pocket-id/backend/internal/utils"
)

type OneTimeAccessTokenCreateDto struct {
	TTL utils.JSONDuration `json:"ttl" binding:"ttl"`
	PermittedClientId *string            `json:"permittedClientId"`
}

type emailAsUnauthenticatedUserDto struct {
	Email        string `json:"email" binding:"required,email" unorm:"nfc"`
	RedirectPath string `json:"redirectPath"`
}

type emailAsAdminDto struct {
	TTL utils.JSONDuration `json:"ttl" binding:"ttl"`
}
