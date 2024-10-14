package model

import (
	"github.com/thisissoroush/thunder.shared/primitive/enum"
)

type BaseRequest struct {
	RequestId        UUID
	RequestIPAddress string
	RequestUserId    UUID
	RequestCulture   enum.Language
}
