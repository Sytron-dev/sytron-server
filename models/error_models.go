package models

type ErrorResponse struct {
	Message string `json:"message"`
	Error   error  `json:"error"`
}

type DataResponse struct {
	Message  string `json:"message,omitempty"`
	Data     any    `json:"data"`
	Metadata any    `json:"metadata,omitempty"`
}
