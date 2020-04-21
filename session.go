package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/securecookie"
)

var cookieHandler = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))

type SessionData struct {
	Email     string
	Nounce    string
	Name      string
	LastLogin string
}

func GetToken() string {
	crutime := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(crutime, 10))
	token := fmt.Sprintf("%x", h.Sum(nil))
	return token
}

func GetEmail(request *http.Request) string {
	cookie, err := request.Cookie("session")
	if err != nil {
		//log.Println("missing session cookie")
		return ""
	}
	cookieValue := make(map[string]string)
	err = cookieHandler.Decode("session", cookie.Value, &cookieValue)
	if err != nil {
		//log.Println("invalid session cookie")
		return ""
	}
	if cookieValue["email"] == "" {
		log.Println("no email in session cookie")
		return ""
	}
	if !UserNonceMatches(cookieValue["nonce"], cookieValue["email"]) {
		log.Println("nonce does not match in session cookie")
		return ""
	}
	return cookieValue["email"]
}

func GetSession(request *http.Request) (s SessionData) {
	cookie, err := request.Cookie("session")
	if err != nil {
		return
	}
	cookieValue := SessionData{}
	err = cookieHandler.Decode("session", cookie.Value, &cookieValue)
	if err != nil {
		log.Fatalln("Error decoding session data: ", cookieValue)
	}
	return
}

func SetSession(response http.ResponseWriter, email string, nonce string) {
	if email == "" || nonce == "" {
		return
	}
	value := map[string]string{
		"email": email,
		"nonce": nonce,
	}
	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
}

func ClearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}
