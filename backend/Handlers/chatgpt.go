package Handlers

import (
	"backend/Services"
	"backend/utils"
	"encoding/json"
	"log"
	"net/http"
	"strings"
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
	content := r.URL.Query().Get("text")

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

// GuideByText CalcEmbeddings 入力テキストに基づいて道案内を行う
func GuideByText(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query().Get("text")

	textEmbedded, err := utils.Embeddings(text)
	if err != nil {
		log.Println(err)
		utils.HandleError(w, http.StatusInternalServerError, err.Error())
		return
	}

	m, err := Services.CreateInstructCossimMap(textEmbedded)
	if err != nil {
		log.Println(err)
		utils.HandleError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// コサイン類似度の高い順にソート
	mSorted := utils.InstructSortByCosin(m)

	// TopNを抜き出す処理
	selectedInstructions := Services.SelectInstruction(mSorted, 3)

	// prompt作成処理
	prompt := utils.GenerateTemplate(strings.Join(selectedInstructions, "\n"), text)
	log.Println(prompt)
	chatGPTResponse, err := Services.Chat(prompt)
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
