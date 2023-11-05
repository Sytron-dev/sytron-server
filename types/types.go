package types

import (
	"github.com/gofiber/fiber/v2"
)

type Coordinates struct {
	Latitude  float64 `bson:"lat" json:"lat"`
	Longitude float64 `bson:"lon"    json:"lon"`
}

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
