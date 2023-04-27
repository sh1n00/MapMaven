package utils

import (
	"backend/Settings"
	"backend/Types"
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"math"
	"net/http"
)

// Embeddings textをEmbeddingした結果を返す
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

// Cosine ref: https://github.com/gaspiman/cosine_similarity
func Cosine(a []float64, b []float64) (cosine float64, err error) {
	if len(a) != len(b) {
		return 0.0, errors.New("vector Length Different")
	}
	sum := 0.0
	s1 := 0.0
	s2 := 0.0
	for k := 0; k < len(a); k++ {
		sum += a[k] * b[k]
		s1 += math.Pow(a[k], 2)
		s2 += math.Pow(b[k], 2)
	}
	if s1 == 0 || s2 == 0 {
		return 0.0, errors.New("vectors should not be null (all zeros)")
	}
	return sum / (math.Sqrt(s1) * math.Sqrt(s2)), nil
}
