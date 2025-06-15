package responses

type HTTPResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
