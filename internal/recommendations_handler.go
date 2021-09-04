package internal

import (
	"net/http"

	"strconv"

	"encoding/json"

	"github.com/2beens/anas-api/internal/recommendations"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type RecommendationsHandler struct {
	api recommendations.Api
}

func NewRecommendationsHandler(
	api recommendations.Api,
) *RecommendationsHandler {
	return &RecommendationsHandler{
		api: api,
	}
}

func (h *RecommendationsHandler) handleRecommendationsToday(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIdStr := vars["userId"]
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		log.Errorf("handle recommendations today: %s", err)
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	dbRecs, err := h.api.GetAll(userId)
	if err != nil {
		log.Errorf("handle recommendations today: %s", err)
		// TODO: do not return err details in the future
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var recs []*recommendations.Recommendation
	for _, r := range dbRecs {
		recs = append(recs, r.ToRecommendation())
	}

	recsJson, err := json.Marshal(recs)
	if err != nil {
		log.Errorf("marshal recommendations: %s", err)
		http.Error(w, "marshal recommendations error", http.StatusInternalServerError)
		return
	}

	log.Printf("returning recommendations for today for user %d", userId)

	WriteResponseBytes(w, "application/json", recsJson)
}
