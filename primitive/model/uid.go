package model

import "github.com/google/uuid"

type UUID uuid.UUID

func NewUid() UUID {
	return UUID(uuid.New())
}

var UID UUID
