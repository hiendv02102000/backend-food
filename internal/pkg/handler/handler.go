package handler

import (
	"backend-food/pkg/infrastucture/db"
)

type HTTPHandler struct {
}

func NewHTTPHandler(db db.Database) *HTTPHandler {

	return &HTTPHandler{}
}
