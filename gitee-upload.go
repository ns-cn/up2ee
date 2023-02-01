package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"time"
)

const PRE = "cloud"

func uploadImages(images []string) []string {
	var downloadUrls []string = make([]string, len(images))
	createRepository()
	for _, image := range images {
		if image != "" {
			downloadUrl := uploadImage(image)
			if downloadUrl != "" {
				downloadUrls = append(downloadUrls, downloadUrl)
			}
		}
	}
	return downloadUrls
}

func createRepository() {
	url := "https://gitee.com/api/v5/user/repos"
	method := "POST"
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("access_token", accessData.AccessToken)
	_ = writer.WriteField("name", Repository)
	_ = writer.WriteField("auto_init", "true")
	_ = writer.Close()
	client := &http.Client{}
	req, _ := http.NewRequest(method, url, payload)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, _ := client.Do(req)
	defer res.Body.Close()
}

func uploadImage(image string) (downloadUrl string) {
	FileName := uuid.New().String()
	now := time.Now()
	url := fmt.Sprintf("https://gitee.com/api/v5/repos/%s/%s/contents/%d/%s/%s.png", UserName, Repository, now.Year(), now.Month().String(), FileName)
	method := "POST"
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("access_token", accessData.AccessToken)
	_ = writer.WriteField("content", image)
	_ = writer.WriteField("message", fmt.Sprintf("upload file:%s", FileName))
	_ = writer.WriteField("branch", "master")
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	var ResContent CreateFileData
	err = json.Unmarshal(body, &ResContent)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	downloadUrl = ResContent.Content.DownloadUrl
	return
}
