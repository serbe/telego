package telego

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// Errors
var (
	ErrMissingParam  = errors.New("Missing param")
	ErrForbiddenHTTP = errors.New("Forbidden http")
)

func (bot *Telebot) createResponse(method string, values url.Values) (Response, error) {
	fullURL := urlAPI + bot.Token + "/" + method
	resp, err := bot.Client.PostForm(fullURL, values)
	if err != nil {
		errLog("bot.Client.PostForm", err)
		return Response{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusForbidden {
		err = ErrForbiddenHTTP
		errLog("http.StatusForbidden", err)
		return Response{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errLog("ioutil.ReadAll", err)
		return Response{}, err
	}

	debugLog("debugLog " + method + " " + string(body))

	var response Response

	json.Unmarshal(body, &response)

	if !response.Ok {
		err = errors.New(response.Description)
		errLog("!response.Ok", err)
		return response, err
	}

	return response, nil
}

func errLog(s string, err error) {
	if botErrorLog {
		log.Println("Error in ", s, err)
	}
}

func debugLog(s string) {
	if botDebugLog {
		log.Println(s)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errLog("rootHandler ReadAll", err)
		return
	}
	log.Println(string(body))
}
