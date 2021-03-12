package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type Response struct {
	ID         string `json:"id"`
	Output_url string `json:"output_url"`
}

// Deprecated
func getWaifu2xImgURLFromLocal(imgPath string) (error, string) {

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, errFile1 := os.Open(imgPath)
	defer file.Close()
	part1, errFile1 := writer.CreateFormFile("image", filepath.Base(imgPath))
	_, errFile1 = io.Copy(part1, file)
	if errFile1 != nil {
		return errFile1, ""
	}
	err := writer.Close()
	if err != nil {
		return err, ""
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, apiURL, payload)

	if err != nil {
		return err, ""
	}
	req.Header.Add("api-key", apiKey)

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		return err, ""
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err, ""
	}

	var resp Response
	json.Unmarshal([]byte(body), &resp)

	return nil, resp.Output_url
}

func getWaifu2xImgFromURL(imgURL string) (error, string) {

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("image", imgURL)
	err := writer.Close()
	if err != nil {
		return err, ""
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, apiURL, payload)

	if err != nil {
		return err, ""
	}
	req.Header.Add("api-key", apiKey)

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		return err, ""
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err, ""
	}

	var resp Response
	json.Unmarshal([]byte(body), &resp)

	return nil, resp.Output_url
}
