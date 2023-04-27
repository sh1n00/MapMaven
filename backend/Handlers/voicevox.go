package Handlers

import (
	"backend/Services"
	"backend/utils"
	"encoding/json"
	"log"
	"net/http"
)

func TextToAudio(w http.ResponseWriter, r *http.Request) {
	speaker := r.URL.Query().Get("speaker")
	enableInterrogativeUpspeak := r.URL.Query().Get("enable_interrogative_upspeak")
	text := r.URL.Query().Get("text")

	audioJson, err := Services.TextToVoice(speaker, enableInterrogativeUpspeak, text)
	if err != nil {
		log.Println(err)
		utils.HandleError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(audioJson); err != nil {
		log.Println(err)
		return
	}

}
