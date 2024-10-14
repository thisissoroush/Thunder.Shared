package model

import (
	"github.com/thisissoroush/thunder.shared/primitive/enum"

	"github.com/google/uuid"
)

type BaseRequest struct {
	RequestId        uuid.UUID
	RequestIPAddress string
	RequestUserId    uuid.UUID
	RequestCulture   enum.Language
}
