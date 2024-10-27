package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)

	expectedBody := "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент"
	assert.Equal(t, expectedBody, responseRecorder.Body.String())

	assert.NotEmpty(t, responseRecorder.Body.String())

	cafesReturned := responseRecorder.Body.String()
	cafesList := splitCafesString(cafesReturned)
	assert.Equal(t, totalCount, len(cafesList))
}

func splitCafesString(cafesStr string) []string {
	return strings.Split(cafesStr, ",")
}
