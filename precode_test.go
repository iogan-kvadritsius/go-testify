package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenOk(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=4&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	responce := responseRecorder.Result()
	// проверка статус 200
	require.Equal(t, 200, responce.StatusCode, "Unexpected status code")
	// проверка тела ответа не пустое
	require.NotEmpty(t, responce.Body, "responce body is empty")
}

func TestMainHandlerWhenCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=4&city=NotCity", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	responce := responseRecorder.Result()
	// проверка ответа код 400
	assert.Equal(t, http.StatusBadRequest, responce.StatusCode, "expected status code")

	body, _ := io.ReadAll(responce.Body)
	require.Contains(t, string(body), "wrong city value")
}

func TestMainHandlerWhenCount(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	responce := responseRecorder.Result()

	body, _ := io.ReadAll(responce.Body)
	list := strings.Split(string(body), ",")
	// проверка длины списка
	assert.Len(t, list, totalCount, "Unexpected number of cafes")
}
