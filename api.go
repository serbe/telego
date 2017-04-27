package telego

import "encoding/json"
import "errors"

var errMissingParam = errors.New("Missing param")

// Response - response from the Telegram API
type Response struct {
	Ok          bool                `json:"ok"`
	Result      json.RawMessage     `json:"result"`
	ErrorCode   int                 `json:"error_code"`
	Description string              `json:"description"`
	Parameters  *ResponseParameters `json:"parameters"`
}
