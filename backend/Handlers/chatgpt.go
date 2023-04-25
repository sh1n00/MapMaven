package Handlers

import (
	"backend/Services"
	"backend/utils"
	"encoding/binary"
	"encoding/json"
	"log"
	"math"
	"net/http"
)

// HealthCheck ChatGPTAPI通信確認用
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	chatGPTResponse, err := Services.HealthCheck()
	if err != nil {
		utils.HandleError(w, http.StatusInternalServerError, err.Error())
	}

	if err = json.NewEncoder(w).Encode(chatGPTResponse); err != nil {
		log.Println(err)
		utils.HandleError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

// Chat ChatGPTにMessageを送る
func Chat(w http.ResponseWriter, r *http.Request) {
	content := r.URL.Query().Get("content")

	chatGPTResponse, err := Services.Chat(content)
	if err != nil {
		log.Println(err)
		utils.HandleError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err = json.NewEncoder(w).Encode(chatGPTResponse); err != nil {
		log.Println(err)
		utils.HandleError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func Embeddings(w http.ResponseWriter, r *http.Request) {
	input := r.URL.Query().Get("input")

	embeddings, err := Services.Embeddings(input)
	if err != nil {
		log.Println(err)
		utils.HandleError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err = json.NewEncoder(w).Encode(embeddings); err != nil {
		log.Println(err)
		utils.HandleError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func CalcCosSimilarity(w http.ResponseWriter, r *http.Request) {
	text1 := r.URL.Query().Get("text1")
	text2 := r.URL.Query().Get("text2")

	cosin, err := Services.CalcCosSimilarity(text1, text2)
	if err != nil {
		log.Println(err)
		utils.HandleError(w, http.StatusInternalServerError, err.Error())
		return
	}

	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, math.Float64bits(cosin))
	w.Write(b)
}
