package telego

import "net/http"

const urlAPI = "https://api.telegram.org/bot"

var (
	botDebugLog bool
	botErrorLog bool
)

// Bot - telego bot main struct
type Bot struct {
	Token  string
	Client *http.Client
	Self   *User
}

// NewBot - create new bot
func NewBot(token string, debugLog bool, errorLog bool) (*Bot, error) {
	botDebugLog = debugLog
	botErrorLog = errorLog

	bot := new(Bot)
	bot.Token = token
	bot.Client = new(http.Client)

	user, err := bot.getMe()
	if err != nil {
		errLog("bot.getMe", err)
		return &Bot{}, err
	}

	bot.Self = &user

	return bot, nil
}
