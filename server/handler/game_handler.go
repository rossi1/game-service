package handler

import (
	"encoding/json"
	"game-service/utils"
	"net/http"
)

var totalScore = map[string]int{"Computer": 0, "User": 0, "Tie": 0, "Rounds": 0}

func Response(w http.ResponseWriter, res interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")

	content, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)

	w.Write(content)
}

func (h *Handler) Play(w http.ResponseWriter, r *http.Request) {
	var req userInput

	err := json.NewDecoder(r.Body).Decode(&req)

	defer r.Body.Close()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	result, winner, user, computer := utils.DetermineWinner(req.Player, utils.GenerateComputerChoice())
	totalScore[winner]++
	totalScore["Rounds"]++

	res := gameResult{
		Computer: computer,
		Player:   user,
		Results:  result,
	}
	Response(w, res, http.StatusOK)

}

func (h *Handler) Choices(w http.ResponseWriter, r *http.Request) {
	choices := utils.GetChoices()
	Response(w, choices, http.StatusOK)
}

func (h *Handler) ScoreBoard(w http.ResponseWriter, r *http.Request) {
	scores := totalScore
	Response(w, scores, http.StatusOK)
}

func (h *Handler) ResetScoreBoard(w http.ResponseWriter, r *http.Request) {
	for entry := range totalScore {
		delete(totalScore, entry)
	}
	Response(w, totalScore, http.StatusOK)
}

func (h *Handler) Choice(w http.ResponseWriter, r *http.Request) {
	choice := utils.GetRandomChoice()
	Response(w, choice, http.StatusOK)

}
