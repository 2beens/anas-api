package internal

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

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

	log.Printf("returning therapy day %d for for user %d", dayIndex, userId)

	WriteResponseBytes(w, "application/json", dayJson)
}

func (h *TherapyHandler) handleGetTherapyAudio(w http.ResponseWriter, r *http.Request) {
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

	f, err := os.Open(day.AudioFilePath)
	if err != nil {
		log.Errorf("get therapy audio [%s]: %s", day.AudioFilePath, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pathParts := strings.Split(day.AudioFilePath, string(filepath.Separator))
	fileName := pathParts[len(pathParts)-1]

	reader := bufio.NewReader(f)
	buf := new(bytes.Buffer)
	bytesRead, err := buf.ReadFrom(reader)
	if err != nil {
		log.Errorf("get therapy audio [%s]: %s", day.AudioFilePath, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "audio/mpeg")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", bytesRead))
	// to encourage the browser to download the mp3 rather then streaming
	w.Header().Set("Content-Disposition", fmt.Sprintf("filename=%s", fileName))

	_, err = io.Copy(w, buf)
	if err != nil {
		log.Errorf("get therapy audio, copy [%s]: %s", day.AudioFilePath, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
