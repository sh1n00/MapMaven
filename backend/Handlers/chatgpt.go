package Handlers

import (
	"backend/Settings"
	"backend/Types"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// HealthCheck ChatGPTAPI通信確認用
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	url := "https://api.openai.com/v1/chat/completions"
	reqBody := Types.ChatGPTRequest{
		Model:    "gpt-3.5-turbo",
		Messages: []Types.Message{{Role: "user", Content: "Hello"}},
	}
	jsonReqBody, err := json.Marshal(reqBody)
	if err != nil {
		log.Println(err)
		return
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonReqBody))
	if err != nil {
		log.Println(err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+Settings.OPENAIAPIKEY)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("Status: ", resp.Status)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	var response Types.ChatGPTResponse
	if err = json.Unmarshal(body, &response); err != nil {
		log.Println(err)
		return
	}

	if err = json.NewEncoder(w).Encode(response); err != nil {
		log.Println(err)
		return
	}
}

// Chat ChatGPTにMessageを送る
func Chat(w http.ResponseWriter, r *http.Request) {
	content := r.URL.Query().Get("content")
	url := "https://api.openai.com/v1/chat/completions"
	reqBody := Types.ChatGPTRequest{
		Model:    "gpt-3.5-turbo",
		Messages: []Types.Message{{Role: "user", Content: content}},
	}

	jsonReqBody, err := json.Marshal(reqBody)
	if err != nil {
		log.Println(err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonReqBody))
	if err != nil {
		log.Println(err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+Settings.OPENAIAPIKEY)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("Status: ", resp.Status)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	var response Types.ChatGPTResponse
	if err = json.Unmarshal(body, &response); err != nil {
		log.Println(err)
		return
	}

	if err = json.NewEncoder(w).Encode(response); err != nil {
		log.Println(err)
		return
	}
}
