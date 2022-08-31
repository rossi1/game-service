package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/stephenafamo/janus"
)

type userInput struct {
	Player int `json:"player"`
}

type gameResult struct {
	Computer int    `json:"computer"`
	Player   int    `json:"player"`
	Results  string `json:"results"`
}

type Handler struct {
	janus.Broker
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) GetRouter() http.Handler {
	router := chi.NewRouter()
	router.Use(h.RecommendedMiddlewares()...)
	router.Get("/choice", h.Choice)
	router.Get("/choices", h.Choices)
	router.Get("/score-board", h.ScoreBoard)
	router.Get("/reset", h.ResetScoreBoard)
	router.Post("/play", h.Play)

	return router
}
