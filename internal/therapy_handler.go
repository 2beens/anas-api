package internal

import (
	"net/http"
	"strconv"

	"fmt"

	"encoding/json"

	"github.com/2beens/anas-api/internal/therapy"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type TherapyHandler struct {
	api therapy.Api
}

func NewTherapyHandler(api therapy.Api) *TherapyHandler {
	return &TherapyHandler{
		api: api,
	}
}

func (h *TherapyHandler) handleGetTherapy(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIdStr := vars["userId"]
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	dayIndexStr := vars["dayIndex"]
	dayIndex, err := strconv.Atoi(dayIndexStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid day index: %d", dayIndexStr), http.StatusBadRequest)
		return
	}

	day, err := h.api.GetDay(userId, dayIndex)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dayJson, err := json.Marshal(day)
	if err != nil {
		log.Errorf("marshal day: %s", err)
		http.Error(w, "marshal day error", http.StatusInternalServerError)
		return
	}

	WriteResponseBytes(w, "application/json", dayJson)
}
