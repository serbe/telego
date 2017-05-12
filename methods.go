package telego

import (
	"encoding/json"
	"net/url"
	"strconv"
)

// getMe - A simple method for testing your bot's auth token. Requires no parameters. Returns basic information
// about the bot in form of a User object.
func (bot *Telebot) getMe() (User, error) {
	r, err := bot.createResponse("getMe", nil)
	if err != nil {
		errLog("getMe createResponse", err)
		return User{}, err
	}
	var user User
	err = json.Unmarshal(r.Result, &user)
	if err != nil {
		errLog("getMe Unmarshal", err)
	}
	return user, err
}

// SendMessage - "sendMessage" Use this method to send text messages. On success, the sent Message is returned.
//
// Parameters	Type	    Required	Description
// chat_id	    Integer	    Yes	        Unique identifier for the target chat or username of the target channel
//              or String				(in the format @channelusername)
// text	        String	    Yes	        Text of the message to be sent
// parse_mode	String	    Optional	Send Markdown or HTML, if you want Telegram apps to show bold, italic,
//										fixed-width text or inline URLs in your bot's message.
// disable_web_page_preview
//              Boolean	    Optional	Disables link previews for links in this message
// disable_notification
//              Boolean	    Optional	Sends the message silently. iOS users will not receive a notification,
//										Android users will receive a notification with no sound.
// reply_to_message_id
//              Integer	    Optional	If the message is a reply, ID of the original message
// reply_markup	            Optional	Additional interface options. A JSON-serialized object for an inline
//										keyboard, custom reply keyboard, instructions
//              InlineKeyboardMarkup or            to remove reply keyboard or to force a reply from the user.
//              ReplyKeyboardMarkup or
//              ReplyKeyboardRemove or
//              ForceReply
//
// Formatting options
// The Bot API supports basic formatting for messages. You can use bold and italic text, as well as inline links and pre-formatted code in your bots' messages.
// Telegram clients will render them accordingly. You can use either markdown-style or HTML-style formatting.
//
// Note that Telegram clients will display an alert to the user before opening an inline link (‘Open this link?’ together with the full URL).
//
// Markdown style
// To use this mode, pass Markdown in the parse_mode field when using sendMessage. Use the following syntax in your message:
//
// *bold text*
// _italic text_
// [text](http://www.example.com/)
// `inline fixed-width code`
// ```text
// pre-formatted fixed-width code block
// ```
// HTML style
//
// To use this mode, pass HTML in the parse_mode field when using sendMessage. The following tags are currently supported:
//
// <b>bold</b>, <strong>bold</strong>
// <i>italic</i>, <em>italic</em>
// <a href="http://www.example.com/">inline URL</a>
// <code>inline fixed-width code</code>
// <pre>pre-formatted fixed-width code block</pre>
// Please note:
//
// Only the tags mentioned above are currently supported.
// Tags must not be nested.
// All <, > and & symbols that are not a part of a tag or an HTML entity must be replaced with the corresponding HTML
// entities (< with &lt;, > with &gt; and & with &amp;).
// All numerical HTML entities are supported.
// The API currently supports only the following named HTML entities: &lt;, &gt;, &amp; and &quot;.
func (bot *Telebot) SendMessage(opt *SendMessageOpts) (Message, error) {
	values := url.Values{}
	if opt.ChatID == "" || opt.Text == "" {
		return Message{}, ErrMissingParam
	}
	values.Set("chat_id", opt.ChatID)
	values.Set("text", opt.Text)
	if len(opt.ParseMode) > 0 {
		values.Set("parse_mode", opt.ParseMode)
	}
	if opt.DisableWebPagePreview {
		values.Set("disable_web_page_preview", "true")
	}
	if opt.DisableNotification {
		values.Set("disable_notification", "true")
	}
	if opt.ReplyToMessageID > 0 {
		values.Set("reply_to_message_id", strconv.Itoa(opt.ReplyToMessageID))
	}
	r, err := bot.createResponse("sendMessage", values)
	if err != nil {
		errLog("SendMessage createResponse", err)
	}
	var message Message
	err = json.Unmarshal(r.Result, &message)
	if err != nil {
		errLog("SendMessage Unmarshal", err)
	}
	return message, err
}

// ForwardMessage - "forwardMessage" Use this method to forward messages of any kind. On success, the sent
// Message is returned.
//
// Parameters	Type	    Required	Description
// chat_id	    Integer or  Yes     	Unique identifier for the target chat or username of the target channel
//              String					(in the format @channelusername)
// from_chat_id	Integer or  Yes	        Unique identifier for the chat where the original message was sent (or
//              String					channel username in the format @channelusername)
// disable_notification
//          	Boolean	    Optional	Sends the message silently. iOS users will not receive a notification,
//										Android users will receive a notification with no sound.
// message_id	Integer	    Yes	        Message identifier in the chat specified in from_chat_id
func (bot *Telebot) ForwardMessage(opt *ForwardMessageOpts) (Message, error) {
	values := url.Values{}
	if opt.ChatID == "" || opt.FromChatID == "" || opt.MessageID == 0 {
		return Message{}, ErrMissingParam
	}
	values.Set("chat_id", opt.ChatID)
	values.Set("from_chat_id", opt.FromChatID)
	if opt.DisableNotification {
		values.Set("disable_notification", "true")
	}
	values.Set("message_id", strconv.Itoa(opt.MessageID))
	r, err := bot.createResponse("forwardMessage", values)
	if err != nil {
		errLog("forwardMessage createResponse", err)
	}
	var message Message
	err = json.Unmarshal(r.Result, &message)
	if err != nil {
		errLog("forwardMessage Unmarshal", err)
	}
	return message, err
}

// SendPhoto - "sendPhoto" Use this method to send photos. On success, the sent Message is returned.
//
// Parameters	Type	    Required	Description
// chat_id	    Integer or  Yes	        Unique identifier for the target chat or username of the target channel
//              String					(in the format @channelusername)
// photo	InputFile or 	Yes			Photo to send. Pass a file_id as String to send a photo that exists on
//				String					the Telegram servers (recommended), pass an HTTP URL as a String for
//										Telegram to get a photo from the Internet, or upload a new photo using
//										multipart/form-data. More info on Sending Files »
// caption		String		Optional	Photo caption (may also be used when resending photos by file_id), 0-200 characters
// disable_notification
//				Boolean		Optional	Sends the message silently. iOS users will not receive a notification,
//										Android users will receive a notification with no sound.
// reply_to_message_id
//				Integer		Optional	If the message is a reply, ID of the original message
// reply_markup				Optional	Additional interface options. A JSON-serialized object for an inline
//			InlineKeyboardMarkup or 	keyboard, custom reply keyboard, instructions to remove reply keyboard
//				ReplyKeyboardMarkup or  or to force a reply from the user.
//				ReplyKeyboardRemove or
//				ForceReply
func (bot *Telebot) SendPhoto(opt *SendPhotoOpts) (Message, error) {
	values := url.Values{}
	if opt.ChatID == "" || opt.Photo == "" {
		return Message{}, ErrMissingParam
	}
	values.Set("chat_id", opt.ChatID)
	values.Set("photo", opt.Photo)
	if len(opt.Caption) > 0 {
		values.Set("caption", opt.Caption)
	}
	if opt.DisableNotification {
		values.Set("disable_notification", "true")
	}
	if opt.ReplyToMessageID > 0 {
		values.Set("reply_to_message_id", strconv.Itoa(opt.ReplyToMessageID))
	}
	r, err := bot.createResponse("sendPhoto", values)
	if err != nil {
		errLog("SendPhoto createResponse", err)
	}
	var message Message
	err = json.Unmarshal(r.Result, &message)
	if err != nil {
		errLog("SendPhoto Unmarshal", err)
	}
	return message, err
}

// sendAudio
// Use this method to send audio files, if you want Telegram clients to display them in the music player. Your audio
// must be in the .mp3 format. On success, the sent Message is returned. Bots can currently send audio files of up
// to 50 MB in size, this limit may be changed in the future.
// For sending voice messages, use the sendVoice method instead.
//
// Parameters	Type		Required	Description
// chat_id		Integer 	Yes			Unique identifier for the target chat or username of the target channel (in
//				or String				the format @channelusername)
// audio		InputFile 	Yes			Audio file to send. Pass a file_id as String to send an audio file that
//				or String				exists on the Telegram servers (recommended), pass an HTTP URL as a String
//										for Telegram to get an audio file from the Internet, or upload a new one
//										using multipart/form-data. More info on Sending Files »
// caption		String		Optional	Audio caption, 0-200 characters
// duration		Integer		Optional	Duration of the audio in seconds
// performer	String		Optional	Performer
// title		String		Optional	Track name
// disable_notification		Optional	Sends the message silently. iOS users will not receive a notification,
//				Boolean					Android users will receive a notification with no sound.
// reply_to_message_id		Optional	If the message is a reply, ID of the original message
//				Integer
// reply_markup				Optional	Additional interface options. A JSON-serialized object for an inline
//				InlineKeyboardMarkup	keyboard, custom reply keyboard, instructions to remove reply keyboard
//				or ReplyKeyboardMarkup	or to force a reply from the user.
//				or ReplyKeyboardRemove
//				or ForceReply

// sendDocument
// Use this method to send general files. On success, the sent Message is returned. Bots can currently send files
// of any type of up to 50 MB in size, this limit may be changed in the future.
//
// Parameters	Type		Required	Description
// chat_id		Integer 	Yes			Unique identifier for the target chat or username of the target channel
//				or String				(in the format @channelusername)
// document		InputFile 	Yes			File to send. Pass a file_id as String to send a file that exists on the
//				or String				Telegram servers (recommended), pass an HTTP URL as a String for Telegram
//										to get a file from the Internet, or upload a new one using multipart/form-data.
//										More info on Sending Files »
// caption		String		Optional	Document caption (may also be used when resending documents by file_id),
//										0-200 characters
// disable_notification		Optional	Sends the message silently. iOS users will not receive a notification,
//				Boolean					Android users will receive a notification with no sound.
// reply_to_message_id		Optional	If the message is a reply, ID of the original message
//				Integer
// reply_markup				Optional	Additional interface options. A JSON-serialized object for an inline
//				InlineKeyboardMarkup	keyboard, custom reply keyboard, instructions to remove reply keyboard
//				or ReplyKeyboardMarkup	or to force a reply from the user.
//				or ReplyKeyboardRemove
//				or ForceReply

// sendSticker
// Use this method to send .webp stickers. On success, the sent Message is returned.
//
// Parameters	Type	Required	Description
// chat_id	Integer or String	Yes	Unique identifier for the target chat or username of the target channel (in the format @channelusername)
// sticker	InputFile or String	Yes	Sticker to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a .webp file from the Internet, or upload a new one using multipart/form-data. More info on Sending Files »
// disable_notification	Boolean	Optional	Sends the message silently. iOS users will not receive a notification, Android users will receive a notification with no sound.
// reply_to_message_id	Integer	Optional	If the message is a reply, ID of the original message
// reply_markup	InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply	Optional	Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.

// sendVideo
// Use this method to send video files, Telegram clients support mp4 videos (other formats may be sent as Document). On success, the sent Message is returned. Bots can currently send video files of up to 50 MB in size, this limit may be changed in the future.
//
// Parameters	Type	Required	Description
// chat_id	Integer or String	Yes	Unique identifier for the target chat or username of the target channel (in the format @channelusername)
// video	InputFile or String	Yes	Video to send. Pass a file_id as String to send a video that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a video from the Internet, or upload a new video using multipart/form-data. More info on Sending Files »
// duration	Integer	Optional	Duration of sent video in seconds
// width	Integer	Optional	Video width
// height	Integer	Optional	Video height
// caption	String	Optional	Video caption (may also be used when resending videos by file_id), 0-200 characters
// disable_notification	Boolean	Optional	Sends the message silently. iOS users will not receive a notification, Android users will receive a notification with no sound.
// reply_to_message_id	Integer	Optional	If the message is a reply, ID of the original message
// reply_markup	InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply	Optional	Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.

// sendVoice
// Use this method to send audio files, if you want Telegram clients to display the file as a playable voice message. For this to work, your audio must be in an .ogg file encoded with OPUS (other formats may be sent as Audio or Document). On success, the sent Message is returned. Bots can currently send voice messages of up to 50 MB in size, this limit may be changed in the future.
//
// Parameters	Type	Required	Description
// chat_id	Integer or String	Yes	Unique identifier for the target chat or username of the target channel (in the format @channelusername)
// voice	InputFile or String	Yes	Audio file to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More info on Sending Files »
// caption	String	Optional	Voice message caption, 0-200 characters
// duration	Integer	Optional	Duration of the voice message in seconds
// disable_notification	Boolean	Optional	Sends the message silently. iOS users will not receive a notification, Android users will receive a notification with no sound.
// reply_to_message_id	Integer	Optional	If the message is a reply, ID of the original message
// reply_markup	InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply	Optional	Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.

// sendLocation
// Use this method to send point on the map. On success, the sent Message is returned.
//
// Parameters	Type	Required	Description
// chat_id	Integer or String	Yes	Unique identifier for the target chat or username of the target channel (in the format @channelusername)
// latitude	Float number	Yes	Latitude of location
// longitude	Float number	Yes	Longitude of location
// disable_notification	Boolean	Optional	Sends the message silently. iOS users will not receive a notification, Android users will receive a notification with no sound.
// reply_to_message_id	Integer	Optional	If the message is a reply, ID of the original message
// reply_markup	InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply	Optional	Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.

// sendVenue
// Use this method to send information about a venue. On success, the sent Message is returned.
//
// Parameters	Type	Required	Description
// chat_id	Integer or String	Yes	Unique identifier for the target chat or username of the target channel (in the format @channelusername)
// latitude	Float number	Yes	Latitude of the venue
// longitude	Float number	Yes	Longitude of the venue
// title	String	Yes	Name of the venue
// address	String	Yes	Address of the venue
// foursquare_id	String	Optional	Foursquare identifier of the venue
// disable_notification	Boolean	Optional	Sends the message silently. iOS users will not receive a notification, Android users will receive a notification with no sound.
// reply_to_message_id	Integer	Optional	If the message is a reply, ID of the original message
// reply_markup	InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply	Optional	Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.

// sendContact
// Use this method to send phone contacts. On success, the sent Message is returned.
//
// Parameters	Type	Required	Description
// chat_id	Integer or String	Yes	Unique identifier for the target chat or username of the target channel (in the format @channelusername)
// phone_number	String	Yes	Contact's phone number
// first_name	String	Yes	Contact's first name
// last_name	String	Optional	Contact's last name
// disable_notification	Boolean	Optional	Sends the message silently. iOS users will not receive a notification, Android users will receive a notification with no sound.
// reply_to_message_id	Integer	Optional	If the message is a reply, ID of the original message
// reply_markup	InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply	Optional	Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove keyboard or to force a reply from the user.

// sendChatAction
// Use this method when you need to tell the user that something is happening on the bot's side. The status is set for 5 seconds or less (when a message arrives from your bot, Telegram clients clear its typing status). Returns True on success.
//
// Example: The ImageBot needs some time to process a request and upload the image. Instead of sending a text message along the lines of “Retrieving image, please wait…”, the bot may use sendChatAction with action = upload_photo. The user will see a “sending photo” status for the bot.
// We only recommend using this method when a response from the bot will take a noticeable amount of time to arrive.
//
// Parameters	Type	Required	Description
// chat_id	Integer or String	Yes	Unique identifier for the target chat or username of the target channel (in the format @channelusername)
// action	String	Yes	Type of action to broadcast. Choose one, depending on what the user is about to receive: typing for text messages, upload_photo for photos, record_video or upload_video for videos, record_audio or upload_audio for audio files, upload_document for general files, find_location for location data.

// getUserProfilePhotos
// Use this method to get a list of profile pictures for a user. Returns a UserProfilePhotos object.
//
// Parameters	Type	Required	Description
// user_id	Integer	Yes	Unique identifier of the target user
// offset	Integer	Optional	Sequential number of the first photo to be returned. By default, all photos are returned.
// limit	Integer	Optional	Limits the number of photos to be retrieved. Values between 1—100 are accepted. Defaults to 100.

// getFile
// Use this method to get basic info about a file and prepare it for downloading. For the moment, bots can download files of up to 20MB in size. On success, a File object is returned. The file can then be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>, where <file_path> is taken from the response. It is guaranteed that the link will be valid for at least 1 hour. When the link expires, a new one can be requested by calling getFile again.
//
// Parameters	Type	Required	Description
// file_id	String	Yes	File identifier to get info about
// Note: This function may not preserve the original file name and MIME type. You should save the file's MIME type and name (if available) when the File object is received.

// kickChatMember
// Use this method to kick a user from a group or a supergroup. In the case of supergroups, the user will not be able to return to the group on their own using invite links, etc., unless unbanned first. The bot must be an administrator in the group for this to work. Returns True on success.
//
// Note: This will method only work if the ‘All Members Are Admins’ setting is off in the target group. Otherwise members may only be removed by the group's creator or by the member that added them.
// Parameters	Type	Required	Description
// chat_id	Integer or String	Yes	Unique identifier for the target group or username of the target supergroup (in the format @supergroupusername)
// user_id	Integer	Yes	Unique identifier of the target user

// LeaveChat - "leaveChat" Use this method for your bot to leave a group, supergroup or channel. Returns True on success.
//
// Parameters	Type		Required	Description
// chat_id		Integer or 	Yes			Unique identifier for the target chat or username of the target supergroup
// 				String					or channel (in the format @channelusername)
func (bot *Telebot) LeaveChat(chatID string) (bool, error) {
	values := url.Values{}
	values["chat_id"] = []string{chatID}
	r, err := bot.createResponse("leaveChat", values)
	if err != nil {
		errLog("LeaveChat createResponse", err)
	}
	var result bool
	err = json.Unmarshal(r.Result, &result)
	if err != nil {
		errLog("LeaveChat Unmarshal", err)
	}
	return result, err
}

// unbanChatMember
// Use this method to unban a previously kicked user in a supergroup. The user will not return to the group automatically, but will be able to join via link, etc. The bot must be an administrator in the group for this to work. Returns True on success.
//
// Parameters	Type	Required	Description
// chat_id	Integer or String	Yes	Unique identifier for the target group or username of the target supergroup (in the format @supergroupusername)
// user_id	Integer	Yes	Unique identifier of the target user

// getChat
// Use this method to get up to date information about the chat (current name of the user for one-on-one conversations, current username of a user, group or channel, etc.). Returns a Chat object on success.
//
// Parameters	Type	Required	Description
// chat_id	Integer or String	Yes	Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)

// getChatAdministrators
// Use this method to get a list of administrators in a chat. On success, returns an Array of ChatMember objects that contains information about all chat administrators except other bots. If the chat is a group or a supergroup and no administrators were appointed, only the creator will be returned.
//
// Parameters	Type	Required	Description
// chat_id	Integer or String	Yes	Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)

// getChatMembersCount
// Use this method to get the number of members in a chat. Returns Int on success.
//
// Parameters	Type	Required	Description
// chat_id	Integer or String	Yes	Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)

// getChatMember
// Use this method to get information about a member of a chat. Returns a ChatMember object on success.
//
// Parameters	Type	Required	Description
// chat_id	Integer or String	Yes	Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
// user_id	Integer	Yes	Unique identifier of the target user

// answerCallbackQuery
// Use this method to send answers to callback queries sent from inline keyboards. The answer will be displayed to the user as a notification at the top of the chat screen or as an alert. On success, True is returned.
//
// Alternatively, the user can be redirected to the specified Game URL. For this option to work, you must first create a game for your bot via BotFather and accept the terms. Otherwise, you may use links like telegram.me/your_bot?start=XXXX that open your bot with a parameter.
// Parameters	Type	Required	Description
// callback_query_id	String	Yes	Unique identifier for the query to be answered
// text	String	Optional	Text of the notification. If not specified, nothing will be shown to the user, 0-200 characters
// show_alert	Boolean	Optional	If true, an alert will be shown by the client instead of a notification at the top of the chat screen. Defaults to false.
// url	String	Optional	URL that will be opened by the user's client. If you have created a Game and accepted the conditions via @Botfather, specify the URL that opens your game – note that this will only work if the query comes from a callback_game button.

// Otherwise, you may use links like telegram.me/your_bot?start=XXXX that open your bot with a parameter.
// cache_time	Integer	Optional	The maximum amount of time in seconds that the result of the callback query may be cached client-side. Telegram apps will support caching starting in version 3.14. Defaults to 0.
