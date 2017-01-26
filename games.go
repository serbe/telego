package telego

// sendGame
// Use this method to send a game. On success, the sent Message is returned.

// Parameters	Type	Required	Description
// chat_id	Integer	Yes	Unique identifier for the target chat
// game_short_name	String	Yes	Short name of the game, serves as the unique identifier for the game. Set up your games via Botfather.
// disable_notification	Boolean	Optional	Sends the message silently. iOS users will not receive a notification, Android users will receive a notification with no sound.
// reply_to_message_id	Integer	Optional	If the message is a reply, ID of the original message
// reply_markup	InlineKeyboardMarkup	Optional	A JSON-serialized object for an inline keyboard. If empty, one ‘Play game_title’ button will be shown. If not empty, the first button must launch the game.

// Game - This object represents a game. Use BotFather to create and edit games, their short names will act as unique identifiers.
// title	    String	    Title of the game
// description	String	    Description of the game
// photo	    Array of PhotoSize	Photo that will be displayed in the game message in chats.
// text	        String	    Optional. Brief description of the game or high scores included in the game message. Can be automatically edited to include current high scores for the game when the bot calls setGameScore, or manually edited using editMessageText. 0-4096 characters.
// text_entities	Array of MessageEntity	Optional. Special entities that appear in text, such as usernames, URLs, bot commands, etc.
// animation	Animation	Optional. Animation that will be displayed in the game message in chats. Upload via BotFather
type Game struct {
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Photo       *[]PhotoSize `json:"photo"`
	// Optional
	Text         string           `json:"text"`
	TextEntities *[]MessageEntity `json:"text_entities"`
	Animation    *Animation       `json:"animation"`
}

// Animation - You can provide an animation for your game so that it looks stylish in chats (check out Lumberjack for an example). This object represents an animation file to be displayed in the message containing a game.
// file_id	String	Unique file identifier
// thumb	PhotoSize	Optional. Animation thumbnail as defined by sender
// file_name	String	Optional. Original animation filename as defined by sender
// mime_type	String	Optional. MIME type of the file as defined by sender
// file_size	Integer	Optional. File size
type Animation struct {
	FileID string `json:"file_id"`
	// Optional
	Thumb    *PhotoSize `json:"thumb"`
	FileName string     `json:"file_name"`
	MimeType string     `json:"mime_type"`
	FileSize int        `json:"file_size"`
}

// CallbackGame - A placeholder, currently holds no information. Use BotFather to set up your game.
type CallbackGame struct{}

// setGameScore
// Use this method to set the score of the specified user in a game. On success, if the message was sent by the bot, returns the edited Message, otherwise returns True. Returns an error, if the new score is not greater than the user's current score in the chat and force is False.

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
// position	Integer	Position in high score table for the game
// user	    User	User
// score	Integer	Score
type GameHighScore struct {
	Position int   `json:"position"`
	User     *User `json:"user"`
	Score    int   `json:"score"`
}
