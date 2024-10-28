package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandleValidRequest(t *testing.T) {

	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code, "Expected status code 200")
	assert.NotEmpty(t, responseRecorder.Body.String(), "Expected non-empty response body")
}

func TestMainHandleInvalidCity(t *testing.T) {

	req := httptest.NewRequest("GET", "/cafe?count=2&city=unknown", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code, "Expected status code 400")
	assert.Equal(t, "wrong city value", responseRecorder.Body.String(), "Expected error message 'wrong city value")
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {

	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code, "Expected 200 OK")
	assert.Len(t, strings.Split(responseRecorder.Body.String(), ","), totalCount, "Expected all cafes")
}
