package telego

// Sticker - This object represents a sticker.
// file_id	    String		Unique identifier for this file
// width	    Integer		Sticker width
// height	    Integer		Sticker height
// thumb	    PhotoSize	Optional. Sticker thumbnail in .webp or .jpg format
// emoji	    String		Optional. Emoji associated with the sticker
// file_size	Integer		Optional. File size
type Sticker struct {
	FileID   string     `json:"file_id"`
	Width    int        `json:"width"`
	Height   int        `json:"height"`
	Thumb    *PhotoSize `json:"thumb,omitempty"`
	Emoji    string     `json:"emoji,omitempty"`
	FileSize int        `json:"file_size,omitempty"`
}

// sendSticker
// Use this method to send .webp stickers. On success, the sent Message is returned.
//
// Parameters	Type	Required	Description
// chat_id	Integer or String	Yes	Unique identifier for the target chat or username of the target channel (in the format @channelusername)
// sticker	InputFile or String	Yes	Sticker to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a .webp file from the Internet, or upload a new one using multipart/form-data. More info on Sending Files »
// disable_notification	Boolean	Optional	Sends the message silently. iOS users will not receive a notification, Android users will receive a notification with no sound.
// reply_to_message_id	Integer	Optional	If the message is a reply, ID of the original message
// reply_markup	InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply	Optional	Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
