package requests

import (
	"fmt"
	"net/http"
)

type HTTPHeaders struct {
	headers http.Header
}

func (h *HTTPHeaders) Set(key, value string) *HTTPHeaders {
	h.headers.Set(key, value)
	return h
}

func (h *HTTPHeaders) Add(key, value string) *HTTPHeaders {
	h.headers.Add(key, value)
	return h
}

func (h *HTTPHeaders) ContentTypeJSON() *HTTPHeaders {
	h.headers.Set("Content-Type", "application/json")
	return h
}

func (h *HTTPHeaders) AUTHHeader(token string) *HTTPHeaders {
	h.headers.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	return h
}

func (h *HTTPHeaders) Build() http.Header {
	return h.headers
}

func NewHeader() *HTTPHeaders {
	return &HTTPHeaders{
		headers: make(http.Header),
	}
}
