package responses

import (
	"fmt"
	"net/http"
)

func NewInternalServerError(route string) HTTPResponse {
	return HTTPResponse{
		Message: fmt.Sprintf("%s: internal server error", route),
		Code:    http.StatusInternalServerError,
	}
}
