package requests

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/sync/errgroup"
)

type Request struct {
	url     string
	body    io.Reader
	method  HttpMethods
	headers http.Header
	timeout time.Duration
	client  http.Client
}

func New(url string) *Request {
	return &Request{
		url:     url,
		timeout: time.Second * 10,
		client: http.Client{
			Timeout: time.Second * 10,
			Transport: &http.Transport{
				MaxIdleConns:        10,
				IdleConnTimeout:     90 * time.Second,
				DisableKeepAlives:   false,
				DisableCompression:  false, // Automatically add gzip compression to headers
				ForceAttemptHTTP2:   true,
				TLSHandshakeTimeout: 1 * time.Second,
			},
		},
	}
}

func (r *Request) JSONBody(data Dict) *Request {
	res, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	r.body = bytes.NewReader(res)
	r.headers.Set("Content-Type", "application/json")
	return r
}

func (r *Request) Headers(h *HTTPHeaders) *Request {
	header := h.Build()
	r.headers = header
	return r
}

func (r *Request) Send(ctx context.Context) ([]byte, error) {
	if r.url == "" || r.method == "" {
		return nil, errors.New("ulr or method cannot be empty")
	}

	u, err := url.Parse(r.url)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, string(r.method), u.String(), r.body)
	if err != nil {
		return nil, err
	}

	response, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	bodyData, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(bodyData))
	return bodyData, err
}

func SendBatch(ctx context.Context, requests []*Request) ([][]byte, error) {
	g, ctx := errgroup.WithContext(ctx)
	g.SetLimit(10)

	results := make([][]byte, len(requests))
	for i, req := range requests {
		i, req := i, req
		g.Go(func() error {
			res, err := req.Send(ctx)
			if err != nil {
				return err
			}
			results[i] = res
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return results, nil
}
