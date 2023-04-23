package Handlers

import (
	"backend/Types"
	"backend/utils"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
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

func TextToAudio(w http.ResponseWriter, r *http.Request) {
	url := "http://localhost:50021/synthesis"
	speaker := r.URL.Query().Get("speaker")
	enableInterrogativeUpspeak := r.URL.Query().Get("enable_interrogative_upspeak")
	text := r.URL.Query().Get("text")

	reqBody, err := TextToQuery(text, speaker)
	if err != nil {
		log.Println(err)
		utils.HandleError(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonReqBody, err := json.MarshalIndent(reqBody, "", "  ")
	log.Println(string(jsonReqBody))
	if err != nil {
		log.Println(err)
		utils.HandleError(w, http.StatusInternalServerError, err.Error())
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonReqBody))
	if err != nil {
		log.Println(err)
		utils.HandleError(w, http.StatusInternalServerError, err.Error())
		return
	}

	q := req.URL.Query()
	q.Add("speaker", speaker)
	q.Add("enable_interrogative_upspeak", enableInterrogativeUpspeak)
	req.URL.RawQuery = q.Encode()

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "audio/wav")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		utils.HandleError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Println("Status:", resp.Status)
		utils.HandleError(w, resp.StatusCode, resp.Status)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		utils.HandleError(w, http.StatusInternalServerError, err.Error())
		return
	}

	audioReponse := Types.AudioResponse{
		AudioBinary: body,
	}

	jsonBytes, err := json.Marshal(audioReponse)
	if err != nil {
		log.Println(err)
		utils.HandleError(w, http.StatusInternalServerError, err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)

	if err = os.WriteFile("./output/audio.wav", body, os.ModePerm); err != nil {
		log.Println(err)
		utils.HandleError(w, http.StatusInternalServerError, err.Error())
		return
	}
}
