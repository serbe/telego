package telego

import (
	"encoding/json"
	"net/url"
	"strconv"
)

// SendGame - "sendGame" Use this method to send a game. On success, the sent Message is returned.
//
// Parameters			Type	Required	Description
// chat_id				Integer	Yes			Unique identifier for the target chat
// game_short_name		String	Yes			Short name of the game, serves as the unique identifier for the game.
//											Set up your games via Botfather.
// disable_notification	Boolean	Optional	Sends the message silently. iOS users will not receive a notification,
//											Android users will receive a notification with no sound.
// reply_to_message_id	Integer	Optional	If the message is a reply, ID of the original message
// reply_markup					Optional	A JSON-serialized object for an inline keyboard. If empty, one ‘Play
//				InlineKeyboardMarkup		game_title’ button will be shown. If not empty, the first button must
//											launch the game.
// func (bot *Telebot) SendGame(chatID int, gameShortName string, disableNotification bool, replyToMessageID int, replyMarkup InlineKeyboardMarkup) (Message, error) {
func (bot *Telebot) SendGame(chatID int, gameShortName string, disableNotification bool, replyToMessageID int) (Message, error) {
	values := url.Values{}
	values.Set("chat_id", strconv.Itoa(chatID))
	values.Set("game_short_name", gameShortName)
	if disableNotification {
		values.Set("disable_notification", "true")
	}
	if replyToMessageID > 0 {
		values.Set("reply_to_message_id", strconv.Itoa(replyToMessageID))
	}
	// if replyMarkup != nil {

	// }
	r, err := bot.createResponse("sendGame", values)
	var message Message
	err = json.Unmarshal(r.Result, &message)
	if err != nil {
		errLog("SendMessage Unmarshal", err)
	}
	return message, err
}

// Game - This object represents a game. Use BotFather to create and edit games, their short names will act as
// unique identifiers.
//
// Parameters	Type				Description
// title	    String	    		Title of the game
// description	String	    		Description of the game
// photo	    Array of PhotoSize	Photo that will be displayed in the game message in chats.
// text	        String	    		Optional. Brief description of the game or high scores included in the game message.
//									Can be automatically edited to include current high scores for the game when
//									the bot calls setGameScore, or manually edited using editMessageText. 0-4096 characters.
// text_entities	Array of 		Optional. Special entities that appear in text, such as usernames, URLs, bot
//				MessageEntity		commands, etc.
// animation	Animation			Optional. Animation that will be displayed in the game message in chats. Upload
//									via BotFather
type Game struct {
	Title        string           `json:"title"`
	Description  string           `json:"description"`
	Photo        []*PhotoSize     `json:"photo"`
	Text         string           `json:"text, omitempty"`
	TextEntities []*MessageEntity `json:"text_entities, omitempty"`
	Animation    *Animation       `json:"animation, omitempty"`
}

// Animation - You can provide an animation for your game so that it looks stylish in chats (check out Lumberjack
// for an example). This object represents an animation file to be displayed in the message containing a game.
//
// Parameters	Type		Description
// file_id		String		Unique file identifier
// thumb		PhotoSize	Optional. Animation thumbnail as defined by sender
// file_name	String		Optional. Original animation filename as defined by sender
// mime_type	String		Optional. MIME type of the file as defined by sender
// file_size	Integer		Optional. File size
type Animation struct {
	FileID   string     `json:"file_id"`
	Thumb    *PhotoSize `json:"thumb, omitempty"`
	FileName string     `json:"file_name, omitempty"`
	MimeType string     `json:"mime_type, omitempty"`
	FileSize int        `json:"file_size, omitempty"`
}

// CallbackGame - A placeholder, currently holds no information. Use BotFather to set up your game.
type CallbackGame struct{}

// setGameScore
// Use this method to set the score of the specified user in a game. On success, if the message was sent by the bot, returns the edited Message, otherwise returns True. Returns an error, if the new score is not greater than the user's current score in the chat and force is False.
//
// Parameters	Type	Required	Description
// user_id	Integer	Yes	User identifier
// score	Integer	Yes	New score, must be non-negative
// force	Boolean	Optional	Pass True, if the high score is allowed to decrease. This can be useful when fixing mistakes or banning cheaters
// disable_edit_message	Boolean	Optional	Pass True, if the game message should not be automatically edited to include the current scoreboard
// chat_id	Integer	Optional	Required if inline_message_id is not specified. Unique identifier for the target chat
// message_id	Integer	Optional	Required if inline_message_id is not specified. Identifier of the sent message
// inline_message_id	String	Optional	Required if chat_id and message_id are not specified. Identifier of the inline message

// getGameHighScores
// Use this method to get data for high score tables. Will return the score of the specified user and several of his neighbors in a game. On success, returns an Array of GameHighScore objects.

// This method will currently return scores for the target user, plus two of his closest neighbors on each side. Will also return the top three users if the user and his neighbors are not among them. Please note that this behavior is subject to change.
// Parameters	Type	Required	Description
// user_id	Integer	Yes	Target user id
// chat_id	Integer	Optional	Required if inline_message_id is not specified. Unique identifier for the target chat
// message_id	Integer	Optional	Required if inline_message_id is not specified. Identifier of the sent message
// inline_message_id	String	Optional	Required if chat_id and message_id are not specified. Identifier of the inline message

// GameHighScore - This object represents one row of the high scores table for a game.
//
// Parameters	Type		Description
// position		Integer		Position in high score table for the game
// user	    	User		User
// score		Integer		Score
type GameHighScore struct {
	Position int   `json:"position"`
	User     *User `json:"user"`
	Score    int   `json:"score"`
}
