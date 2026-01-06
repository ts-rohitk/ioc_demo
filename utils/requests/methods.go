package requests

import "net/http"

const (
	POST    HttpMethods = http.MethodPost
	GET     HttpMethods = http.MethodGet
	PUT     HttpMethods = http.MethodPut
	DELETE  HttpMethods = http.MethodDelete
	PATCH   HttpMethods = http.MethodPatch
	OPTIONS HttpMethods = http.MethodOptions
)

func (r *Request) Post() *Request {
	r.method = POST
	return r
}

func (r *Request) Get() *Request {
	r.method = GET
	return r
}

func (r *Request) Put() *Request {
	r.method = PUT
	return r
}

func (r *Request) Delete() *Request {
	r.method = DELETE
	return r
}

func (r *Request) Patch() *Request {
	r.method = PATCH
	return r
}
