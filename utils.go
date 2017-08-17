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

func (bot *Bot) createResponse(method string, values url.Values) (Response, error) {
	fullURL := urlAPI + bot.Token + "/" + method
	resp, err := bot.Client.PostForm(fullURL, values)
	if err != nil {
		errLog("createResponse bot.Client.PostForm", err)
		return Response{}, err
	}

	defer func() {
		err = resp.Body.Close()
		if err != nil {
			errLog("createResponse resp.Body.Close", err)
		}
	}()

	if resp.StatusCode == http.StatusForbidden {
		err = ErrForbiddenHTTP
		errLog("createResponse http.StatusForbidden", err)
		return Response{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errLog("createResponse ioutil.ReadAll", err)
		return Response{}, err
	}

	debugLog("debugLog " + method + " " + string(body))

	var response Response

	err = json.Unmarshal(body, &response)
	if err != nil {
		errLog("createResponse json.Unmarshal", err)
		return Response{}, err
	}

	if !response.Ok {
		err = errors.New(response.Description)
		errLog("createResponse !response.Ok", err)
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

// func rootHandler(w http.ResponseWriter, r *http.Request) {
// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		errLog("rootHandler ReadAll", err)
// 		return
// 	}
// 	defer func() {
// 		err = r.Body.Close()
// 		if err != nil {
// 			errLog("createResponse resp.Body.Close", err)
// 		}
// 	}()
// 	log.Println(string(body))
// }
