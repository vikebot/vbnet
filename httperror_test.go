package vbnet

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHTTPError(t *testing.T) {
	tests := []struct {
		name     string
		message  string
		httpCode int
		code     int
		inner    error
	}{
		{"standard", "unauthorized access to resource", 403, 10000, nil},
		{"standard with inner error", "teapot is too hot", 418, 0, errors.New("teapot-error")},
		{"empty message and negative numbers", "", -1000, -1000, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := NewHTTPError(tt.message, tt.httpCode, tt.code, tt.inner)

			assert.NotNil(t, err, "NewHTTPError() should not be nil")
			assert.Equal(t, tt.message, err.Message(), "invalid message")
			assert.Equal(t, tt.httpCode, err.HTTPCode(), "invalid httpStatusCode")
			assert.Equal(t, tt.code, err.Code(), "invalid code")
			assert.Equal(t, tt.inner, err.Inner(), "invalid inner error")
		})
	}
}

func TestHttpErr_Error(t *testing.T) {
	tests := []struct {
		name     string
		message  string
		httpCode int
		code     int
		inner    error
		expected string
	}{
		{"standard", "unauthorized access to resource", 403, 10000, nil, "vbnet.10000: unauthorized access to resource (HTTP 403)"},
		{"standard with inner error", "teapot is too hot", 418, 0, errors.New("teapot\"-error"), "vbnet.0: teapot is too hot (HTTP 418), due-to: teapot\"-error"},
		{"empty message and negative numbers", "", -1000, -1000, nil, "vbnet.-1000:  (HTTP -1000)"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := NewHTTPError(tt.message, tt.httpCode, tt.code, tt.inner)

			assert.NotNil(t, err, "NewHTTPError() should not be nil")
			assert.Equal(t, tt.expected, err.Error(), "invalid error message construction")
		})
	}
}
