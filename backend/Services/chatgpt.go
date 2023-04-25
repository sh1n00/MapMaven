package Services

import (
	"backend/Settings"
	"backend/Types"
	"backend/utils"
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func HealthCheck() (*Types.ChatGPTResponse, error) {
	url := Settings.CHATGPTAPIBASEURL + "chat/completions"
	reqBody := Types.ChatGPTRequest{
		Model:    "gpt-3.5-turbo",
		Messages: []Types.Message{{Role: "user", Content: "Hello"}},
	}

	jsonReqBody, err := json.MarshalIndent(reqBody, "", "  ")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonReqBody))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+Settings.OPENAIAPIKEY)

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		log.Println("Status: ", resp.Status)
		return nil, errors.New("Status:" + resp.Status)
	}

	var chatGPTResponse Types.ChatGPTResponse
	if err = json.NewDecoder(resp.Body).Decode(&chatGPTResponse); err != nil {
		log.Println(err)
		return nil, err
	}

	return &chatGPTResponse, nil
}

func Chat(content string) (*Types.ChatGPTResponse, error) {
	url := Settings.CHATGPTAPIBASEURL + "chat/completions"
	reqBody := Types.ChatGPTRequest{
		Model:    "gpt-3.5-turbo",
		Messages: []Types.Message{{Role: "user", Content: content}},
	}

	jsonReqBody, err := json.MarshalIndent(reqBody, "", "  ")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonReqBody))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+Settings.OPENAIAPIKEY)

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		log.Println("Status: ", resp.Status)
		return nil, errors.New("Status: " + resp.Status)
	}

	var chatGPTResponse Types.ChatGPTResponse
	if err = json.NewDecoder(resp.Body).Decode(&chatGPTResponse); err != nil {
		log.Println(err)
		return nil, err
	}

	return &chatGPTResponse, nil
}

func Embeddings(input string) (*Types.Embedding, error) {
	url := Settings.CHATGPTAPIBASEURL + "embeddings"

	reqBody := Types.EmbeddingRequest{
		Input: input,
		Model: "text-embedding-ada-002",
	}

	jsonReqBody, err := json.MarshalIndent(reqBody, "", "  ")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonReqBody))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+Settings.OPENAIAPIKEY)

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		log.Println("Status: ", resp.Status)
		return nil, err
	}

	var embeddingResponse Types.Embedding
	if err = json.NewDecoder(resp.Body).Decode(&embeddingResponse); err != nil {
		log.Println(err)
		return nil, err
	}

	return &embeddingResponse, nil
}

func CalcCosSimilarity(text1, text2 string) (float64, error) {
	text1Embedded, err := Embeddings(text1)
	if err != nil {
		return 0.0, err
	}

	text2Embedded, err := Embeddings(text2)
	if err != nil {
		return 0.0, err
	}

	text1Vector := text1Embedded.Data[0].Embedding
	text2Vector := text2Embedded.Data[0].Embedding

	cosinSim, err := utils.Cosine(text1Vector, text2Vector)
	if err != nil {
		return 0.0, err
	}

	return cosinSim, nil
}
