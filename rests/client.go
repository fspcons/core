package rests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

const (
	//AuthHeader authorization header key
	AuthHeader = "Authorization"
	//ContentTypeHeader content type header key
	ContentTypeHeader = "Content-Type"
	//DefaultTo default t/o setting for http client
	DefaultTo = 15 * time.Second
)

// Config simple http client timeouts config
type Config struct {
	ClientTimeout    *time.Duration
	DialerTimeout    *time.Duration
	KeepAlive        *time.Duration
	HandshakeTimeout *time.Duration
}

// NewClient builder for a http Client with some optional timeout config
func NewClient(cfg *Config) *http.Client {
	c := &http.Client{Timeout: DefaultTo}
	if cfg != nil {
		dialerCfg := &net.Dialer{}
		if cfg.DialerTimeout != nil {
			dialerCfg.Timeout = *cfg.DialerTimeout
		}
		if cfg.DialerTimeout != nil {
			dialerCfg.KeepAlive = *cfg.KeepAlive
		}
		t := &http.Transport{DialContext: (dialerCfg).DialContext}
		if cfg.HandshakeTimeout != nil {
			t.TLSHandshakeTimeout = *cfg.HandshakeTimeout
		}
		c.Transport = t

		if cfg.ClientTimeout != nil {
			c.Timeout = *cfg.ClientTimeout
		}

		return c
	}

	return c
}

// Dispatch a simple 1-liner request method (on request success, always returns http response so status can be asserted)
// inform a non-nil payload param to add a body to the request
// inform a non-nil respObj to get an unmarshall on the response body
func Dispatch(ctx context.Context, client *http.Client, url, method string, headers map[string]string, payload any, respObj any) (*http.Response, error) {
	var body io.Reader
	if payload != nil {
		j, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}

		body = bytes.NewBuffer(j)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return resp, err
	}

	if respObj != nil {
		err = json.NewDecoder(resp.Body).Decode(respObj)
	}

	return resp, err
}

// WithAuthHeader header builder fn for jwt auth
func WithAuthHeader(h map[string]string, token string) map[string]string {
	h = ensureHMap(h)
	h[AuthHeader] = fmt.Sprintf("Bearer %s", token)
	return h
}

// WithContentType header builder fn for content type
func WithContentType(h map[string]string, contentType string) map[string]string {
	h = ensureHMap(h)
	h[ContentTypeHeader] = contentType
	return h
}

func ensureHMap(h map[string]string) map[string]string {
	if h == nil {
		h = make(map[string]string)
	}
	return h
}
