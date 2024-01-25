package types

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Coordinates struct {
	Latitude  float64 `json:"lat" db:"lat"`
	Longitude float64 `json:"lon" db:"lon"`
}

type UUID = uuid.UUID

// Network types

type HandlerFunc func(ctx *fiber.Ctx) error

type EmptyResponse struct{}

type DataResponse struct {
	Message  string `json:"message,omitempty"`
	Data     any    `json:"data"`
	Metadata any    `json:"metadata,omitempty"`
}

type ErrorResponse struct {
	Message  string `json:"message"`
	Error    error  `json:"error"`
	Metadata any    `json:"metadata,omitempty"`
}
