package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const (
	ReCaptchaSecret = "6Lc2JFEUAAAAADxut8ESIM-xWsXjb9cKTFhgAfCV"
	ReCaptchaAPI    = "https://www.google.com/recaptcha/api/siteverify"
)

type ReCaptchaResponse struct {
	Success bool `json:"success"`
}

func IsHuman(challenge, ip string) bool {
	if challenge == "" || ip == "" {
		return false
	}
	rs, err := http.PostForm(ReCaptchaAPI, url.Values{
		"secret":   {ReCaptchaSecret},
		"response": {challenge},
		"remoteip": {ip},
	})
	if err != nil {
		log.Println("Error with recaptcha: ", err.Error())
		return true
	}
	defer rs.Body.Close()
	body, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		log.Println("Error with recaptcha: ", err.Error())
		return true
	}
	var r ReCaptchaResponse
	// log.Println(string(body))
	json.Unmarshal(body, &r)
	return r.Success
}
