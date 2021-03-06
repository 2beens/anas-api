package internal

import (
	"net/http"
	log "github.com/sirupsen/logrus"
)

func WriteResponse(w http.ResponseWriter, contentType, message string) {
	WriteResponseBytes(w, contentType, []byte(message))
}

func WriteResponseBytes(w http.ResponseWriter, contentType string, message []byte) {
	if contentType != "" {
		w.Header().Add("Content-Type", contentType)
	}

	if _, err := w.Write(message); err != nil {
		log.Errorf("failed to write response [%s]: %s", message, err)
	}
}

