package gpt

import (
	"embed"
)

//go:embed default.html
var embeddedFiles embed.FS

// RequestData represents data received in request
type RequestData struct {
	Message string `json:"message"`
	//Name    string `json:"name"`
	//Checked bool   `json:"checked"`
}

// Response Data represents data returned in response
type ResponseData struct {
	Message string `json:"message"`
}
