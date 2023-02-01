package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func access() (accessData AccessData, err error) {
	url := "https://gitee.com/oauth/token"
	method := "POST"
	data := fmt.Sprintf("grant_type=password&username=%s&password=%s&client_id=%s&client_secret=%s&scope=projects gists", UserName, UserPassword, ClientId, ClientSecret)
	payload := strings.NewReader(data)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err := client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	if res.StatusCode != 200 {
		err = errors.New("fail to login")
		return
	}
	_ = json.Unmarshal(body, &accessData)
	return
}
