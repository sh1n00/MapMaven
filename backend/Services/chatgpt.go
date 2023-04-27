package Services

import (
	"backend/Settings"
	"backend/Types"
	"backend/db"
	"backend/utils"
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

// HealthCheck ヘルスチェック用
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

// Chat ChatGPTにプロンプトを渡す処理
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

// CreateInstructCossimMap Instructionとコサイン類似度のmapを返す
func CreateInstructCossimMap(textEmbedded *Types.Embedding) (map[string]float64, error) {
	keys, err := db.RedisClient.Conn.Keys("*").Result()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	m := make(map[string]float64)
	for _, key := range keys {
		val, err := db.RedisClient.Conn.Get(key).Result()
		if err != nil {
			log.Println(err)
			return nil, err
		}
		var embedding *Types.Embedding
		if err = json.Unmarshal([]byte(val), &embedding); err != nil {
			log.Println(err)
			return nil, err
		}
		cosSim, err := CalcCosSimilarity(textEmbedded, embedding)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		m[key] = cosSim
	}
	return m, nil
}

// SelectInstruction 指定した数だけコサイン類似度が高いinstructionsを返す
func SelectInstruction(mSorted utils.PairList, top int) []string {
	var selectedInstructions []string
	for i, instruction := range mSorted {
		if i > top {
			break
		}
		selectedInstructions = append(selectedInstructions, instruction.Key)
	}
	return selectedInstructions
}

// CalcCosSimilarity コサイン類似度の計算を行う
func CalcCosSimilarity(text1, text2 *Types.Embedding) (float64, error) {
	text1Vector := text1.Data[0].Embedding
	text2Vector := text2.Data[0].Embedding

	cosinSim, err := utils.Cosine(text1Vector, text2Vector)
	if err != nil {
		return 0.0, err
	}

	return cosinSim, nil
}
