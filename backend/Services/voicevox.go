package Services

import (
	"backend/Types"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func TextToQuery(text string, speaker string) (*Types.VoiceVox, error) {
	url := "http://localhost:50021/audio_query"
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	q := req.URL.Query()
	q.Add("text", text)
	q.Add("speaker", speaker)
	req.URL.RawQuery = q.Encode()

	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Println("Status: " + resp.Status)
		return nil, err
	}

	var response *Types.VoiceVox
	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Println(err)
		return nil, err
	}

	return response, nil
}

func TextToVoice(speaker, enableInterrogativeUpspeak, text string) ([]byte, error) {
	url := "http://localhost:50021/synthesis"

	reqBody, err := TextToQuery(text, speaker)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	jsonReqBody, err := json.MarshalIndent(reqBody, "", "  ")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonReqBody))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	q := req.URL.Query()
	q.Add("speaker", speaker)
	q.Add("enable_interrogative_upspeak", enableInterrogativeUpspeak)
	req.URL.RawQuery = q.Encode()

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "audio/wav")

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		log.Println("Status:", resp.Status)
		return nil, errors.New("Status: " + resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	audioResponse := Types.AudioResponse{
		AudioBinary: body,
	}
	jsonBytes, err := json.MarshalIndent(audioResponse, "", "  ")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	currentDir, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	filePath := filepath.Join(currentDir, "output", "audio.wav")
	if err = os.WriteFile(filePath, body, os.ModePerm); err != nil {
		log.Println(err)
	}

	return jsonBytes, nil
}
