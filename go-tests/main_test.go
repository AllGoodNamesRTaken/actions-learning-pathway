package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestIntMin covers the IntMin function
func TestIntMin(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{1, 2, 1},
		{2, -1, -1},
		{-3, -7, -7},
		{0, 0, 0},
	}

	for _, tt := range tests {
		result := IntMin(tt.a, tt.b)
		assert.Equal(t, tt.expected, result)
	}
}

// TestGetPortWithEnvVar tests the getPort function when the PORT environment variable is set
func TestGetPortWithEnvVar(t *testing.T) {
	// Set a custom PORT environment variable
	os.Setenv("PORT", "9090")
	defer os.Unsetenv("PORT")

	port := getPort()
	assert.Equal(t, "9090", port)
}

// TestGetPortWithDefault tests the getPort function when the PORT environment variable is not set
func TestGetPortWithDefault(t *testing.T) {
	// Unset PORT environment variable
	os.Unsetenv("PORT")

	port := getPort()
	assert.Equal(t, "8080", port)
}

// TestRootEndpoint covers the "/" route
func TestRootEndpoint(t *testing.T) {
	e := setupServer()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "Hello, Docker! <3", rec.Body.String())
}

// TestHealthEndpoint covers the "/health" route
func TestHealthEndpoint(t *testing.T) {
	e := setupServer()

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `{"Status":"OK"}`, rec.Body.String())
}
