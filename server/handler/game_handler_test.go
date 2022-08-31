package handler

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func setupRouter() (*chi.Mux, *Handler) {
	router := chi.NewRouter()
	handler := NewHandler()
	return router, handler
}

func TestHander_ChoiceRoute(t *testing.T) {
	router, handler := setupRouter()
	router.Get("/choice", handler.Choice)
	req := httptest.NewRequest("GET", "/choice", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.NotEmpty(t, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestHander_ChoicesRoute(t *testing.T) {
	router, handler := setupRouter()
	router.Get("/choices", handler.Choices)
	req := httptest.NewRequest("GET", "/choices", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.NotEmpty(t, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestHander_PlayRoute_WithError(t *testing.T) {
	router, handler := setupRouter()
	router.Post("/play", handler.Play)
	req := httptest.NewRequest("POST", "/play", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
