package telego

import "encoding/json"

// Response - response from the Telegram API
type Response struct {
	Ok          bool                `json:"ok"`
	Result      json.RawMessage     `json:"result"`
	ErrorCode   int                 `json:"error_code"`
	Description string              `json:"description"`
	Parameters  *ResponseParameters `json:"parameters"`
}

// User - This object represents a Telegram user or bot.
//
// id	        Integer	Unique identifier for this user or bot
// first_name	String	User‘s or bot’s first name
// last_name	String	Optional. User‘s or bot’s last name
// username	    String	Optional. User‘s or bot’s username
type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name,omitempty"`
	UserName  string `json:"username,omitempty"`
}

// Chat - This object represents a chat.
//
// id	            Integer	Unique identifier for this chat. This number may be greater than 32 bits
//                          and some programming languages may have difficulty/silent defects in
//							interpreting it. But it smaller than 52 bits, so a signed 64 bit integer
//							or double-precision float type are safe for storing this identifier.
// type	            String	Type of chat, can be either “private”, “group”, “supergroup” or “channel”
// title	        String	Optional. Title, for supergroups, channels and group chats
// username	        String	Optional. Username, for private chats, supergroups and channels if available
// first_name	    String	Optional. First name of the other party in a private chat
// last_name	    String	Optional. Last name of the other party in a private chat
// all_members_are_administrators
//					Boolean	Optional. True if a group has ‘All Members Are Admins’ enabled.
type Chat struct {
	ID                          int    `json:"id"`
	Type                        string `json:"type"`
	Title                       string `json:"title,omitempty"`
	UserName                    string `json:"username,omitempty"`
	FirstName                   string `json:"first_name,omitempty"`
	LastName                    string `json:"last_name,omitempty"`
	AllMembersAreAdministrators bool   `json:"all_members_are_administrators,omitempty"`
}

// Message - This object represents a message.
//
// message_id	    Integer	Unique message identifier inside this chat
// from				User	Optional. Sender, can be empty for messages sent to channels
// date	            Integer	Date the message was sent in Unix time
// chat	            Chat	Conversation the message belongs to
// forward_from	    User	Optional. For forwarded messages, sender of the original message
// forward_from_chat
//					Chat	Optional. For messages forwarded from a channel, information about the original
//							channel
// forward_from_message_id
//					Integer	Optional. For forwarded channel posts, identifier of the original message in
//							the channel
// forward_date	    Integer	Optional. For forwarded messages, date the original message was sent in Unix time
// reply_to_message	Message	Optional. For replies, the original message. Note that the Message object in
//							this field will not contain further reply_to_message fields even if it itself is a reply.
// edit_date	    Integer	Optional. Date the message was last edited in Unix time
// text				String	Optional. For text messages, the actual UTF-8 text of the message, 0-4096 characters.
// entities	      Array of 	Optional. For text messages, special entities like usernames, URLs, bot commands, etc.
//           MessageEntity	that appear in the text
// audio	        Audio	Optional. Message is an audio file, information about the file
// document	       Document	Optional. Message is a general file, information about the file
// game	            Game	Optional. Message is a game, information about the game. More about games »
// photo	      Array of 	Optional. Message is a photo, available sizes of the photo
//				  PhotoSize
// sticker	        Sticker	Optional. Message is a sticker, information about the sticker
// video	        Video	Optional. Message is a video, information about the video
// voice	        Voice	Optional. Message is a voice message, information about the file
// caption	        String	Optional. Caption for the document, photo or video, 0-200 characters
// contact	        Contact	Optional. Message is a shared contact, information about the contact
// location	       Location	Optional. Message is a shared location, information about the location
// venue	        Venue	Optional. Message is a venue, information about the venue
// new_chat_member	User	Optional. A new member was added to the group, information about them (this member may be the bot itself)
// left_chat_member	User	Optional. A member was removed from the group, information about them (this member may be the bot itself)
// new_chat_title	String	Optional. A chat title was changed to this value
// new_chat_photo Array of	Optional. A chat photo was change to this value
//				 PhotoSize
// delete_chat_photo True	Optional. Service message: the chat photo was deleted
// group_chat_created True	Optional. Service message: the group has been created
// supergroup_chat_created
//					True	Optional. Service message: the supergroup has been created. This field can‘t be received in a message coming through updates,
//                          because bot can’t be a member of a supergroup when it is created. It can only be found in reply_to_message if someone replies to a very
//                          first message in a directly created supergroup.
// channel_chat_created	True	Optional. Service message: the channel has been created. This field can‘t be received in a message coming through updates, because
//                          bot can’t be a member of a channel when it is created. It can only be found in reply_to_message if someone replies to a very first message in a channel.
// migrate_to_chat_id	Integer	Optional. The group has been migrated to a supergroup with the specified identifier. This number may be greater than 32 bits and
//                          some programming languages may have difficulty/silent defects in interpreting it. But it smaller than 52 bits, so a signed 64 bit integer
//                          or double-precision float type are safe for storing this identifier.
// migrate_from_chat_id	Integer	Optional. The supergroup has been migrated from a group with the specified identifier. This number may be greater than 32 bits and some
//                          programming languages may have difficulty/silent defects in interpreting it. But it smaller than 52 bits, so a signed 64 bit integer or
//                          double-precision float type are safe for storing this identifier.
// pinned_message	Message	Optional. Specified message was pinned. Note that the Message object in this field will not contain further reply_to_message fields even
//                          if it is itself a reply.
type Message struct {
	MessageID             int              `json:"message_id"`
	Date                  int              `json:"date"`
	Chat                  *Chat            `json:"chat"`
	From                  *User            `json:"from,omitempty"`
	ForwardFrom           *User            `json:"forward_from,omitempty"`
	ForwardFromChat       *Chat            `json:"forward_from_chat,omitempty"`
	ForwardFromMessageID  int              `json:"forward_from_message_id,omitempty"`
	ForwardDate           int              `json:"forward_date,omitempty"`
	ReplyToMessage        *Message         `json:"reply_to_message,omitempty"`
	EditDate              int              `json:"edit_date,omitempty"`
	Text                  string           `json:"text,omitempty"`
	Entities              []*MessageEntity `json:"entities,omitempty"`
	Audio                 *Audio           `json:"audio,omitempty"`
	Document              *Document        `json:"document,omitempty"`
	Game                  *Game            `json:"game,omitempty"`
	Photo                 []*PhotoSize     `json:"photo,omitempty"`
	Sticker               *Sticker         `json:"sticker,omitempty"`
	Video                 *Video           `json:"video,omitempty"`
	Voice                 *Voice           `json:"voice,omitempty"`
	Caption               string           `json:"caption,omitempty"`
	Contact               *Contact         `json:"contact,omitempty"`
	Location              *Location        `json:"location,omitempty"`
	Venue                 *Venue           `json:"venue,omitempty"`
	NewChatMember         *User            `json:"new_chat_member,omitempty"`
	LeftChatMember        *User            `json:"left_chat_member,omitempty"`
	NewChatTitle          string           `json:"new_chat_title,omitempty"`
	NewChatPhoto          []*PhotoSize     `json:"new_chat_photo,omitempty"`
	DeleteChatPhoto       bool             `json:"delete_chat_photo,omitempty"`
	GroupChatCreated      bool             `json:"group_chat_created,omitempty"`
	SupergroupChatCreated bool             `json:"supergroup_chat_created,omitempty"`
	ChannelChatCreated    bool             `json:"channel_chat_created,omitempty"`
	MigrateToChatID       int              `json:"migrate_to_chat_id,omitempty"`
	MigrateFromChatID     int              `json:"migrate_from_chat_id,omitempty"`
	PinnedMessage         *Message         `json:"pinned_message,omitempty"`
}

// MessageEntity - This object represents one special entity in a text message. For example, hashtags, usernames, URLs, etc.
//
// type	    String	Type of the entity. Can be mention (@username), hashtag, bot_command, url, email, bold (bold
//					text), italic (italic text), code (monowidth string), pre (monowidth block), text_link (for
//					clickable text URLs), text_mention (for users without usernames)
// offset	Integer	Offset in UTF-16 code units to the start of the entity
// length	Integer	Length of the entity in UTF-16 code units
// url	    String	Optional. For “text_link” only, url that will be opened after user taps on the text
// user	    User	Optional. For “text_mention” only, the mentioned user
type MessageEntity struct {
	Type   string `json:"type"`
	Offset int    `json:"offset"`
	Length int    `json:"length"`
	URL    string `json:"url,omitempty"`
	User   *User  `json:"user,omitempty"`
}

// PhotoSize - This object represents one size of a photo or a file / sticker thumbnail.
// file_id	    String	Unique identifier for this file
// width	    Integer	Photo width
// height	    Integer	Photo height
// file_size	Integer	Optional. File size
type PhotoSize struct {
	FileID   string `json:"file_id"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	FileSize int    `json:"file_size,omitempty"`
}

// Audio - This object represents an audio file to be treated as music by the Telegram clients.
// file_id	    String	Unique identifier for this file
// duration	    Integer	Duration of the audio in seconds as defined by sender
// performer	String	Optional. Performer of the audio as defined by sender or by audio tags
// title	    String	Optional. Title of the audio as defined by sender or by audio tags
// mime_type	String	Optional. MIME type of the file as defined by sender
// file_size	Integer	Optional. File size
type Audio struct {
	FileID    string `json:"file_id"`
	Duration  int    `json:"duration"`
	Performer string `json:"performer,omitempty"`
	Title     string `json:"title,omitempty"`
	MimeType  string `json:"mime_type,omitempty"`
	FileSize  int    `json:"file_size,omitempty"`
}

// Document - This object represents a general file (as opposed to photos, voice messages and audio files).
// file_id	    String		Unique file identifier
// thumb	    PhotoSize	Optional. Document thumbnail as defined by sender
// file_name	String		Optional. Original filename as defined by sender
// mime_type	String		Optional. MIME type of the file as defined by sender
// file_size	Integer		Optional. File size
type Document struct {
	FileID   string     `json:"file_id"`
	Thumb    *PhotoSize `json:"thumb,omitempty"`
	FileName string     `json:"file_name,omitempty"`
	MimeType string     `json:"mime_type,omitempty"`
	FileSize int        `json:"file_size,omitempty"`
}

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

// Video - This object represents a video file.
// file_id	    String		Unique identifier for this file
// width	    Integer		Video width as defined by sender
// height	    Integer		Video height as defined by sender
// duration	    Integer		Duration of the video in seconds as defined by sender
// thumb	    PhotoSize	Optional. Video thumbnail
// mime_type	String		Optional. Mime type of a file as defined by sender
// file_size	Integer		Optional. File size
type Video struct {
	FileID   string     `json:"file_id"`
	Width    int        `json:"width"`
	Height   int        `json:"height"`
	Duration int        `json:"duration"`
	Thumb    *PhotoSize `json:"thumb,omitempty"`
	MimeType string     `json:"mime_type,omitempty"`
	FileSize int        `json:"file_size,omitempty"`
}

// Voice - This object represents a voice note.
// file_id	    String	Unique identifier for this file
// duration	    Integer	Duration of the audio in seconds as defined by sender
// mime_type	String	Optional. MIME type of the file as defined by sender
// file_size	Integer	Optional. File size
type Voice struct {
	FileID   string `json:"file_id"`
	Duration int    `json:"duration"`
	MimeType string `json:"mime_type,omitempty"`
	FileSize int    `json:"file_size,omitempty"`
}

// Contact - This object represents a phone contact.
// phone_number	String	Contact's phone number
// first_name	String	Contact's first name
// last_name	String	Optional. Contact's last name
// user_id	    Integer	Optional. Contact's user identifier in Telegram
type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name,omitempty"`
	UserID      int    `json:"user_id,omitempty"`
}

// Location - This object represents a point on the map.
// longitude	Float	Longitude as defined by sender
// latitude	    Float	Latitude as defined by sender
type Location struct {
	Longitude float32 `json:"longitude"`
	Latitude  float32 `json:"latitude"`
}

// Venue - This object represents a venue.
//
// location	        Location	Venue location
// title	        String	    Name of the venue
// address	        String	    Address of the venue
// foursquare_id	String	    Optional. Foursquare identifier of the venue
type Venue struct {
	Location     *Location `json:"location"`
	Title        string    `json:"title"`
	Address      string    `json:"address"`
	FoursquareID string    `json:"foursquare_id,omitempty"`
}

// UserProfilePhotos - This object represent a user's profile pictures.
//
// total_count	Integer						Total number of profile pictures the target user has
// photos		Array of Array of PhotoSize	Requested profile pictures (in up to 4 sizes each)
type UserProfilePhotos struct {
	TotalCount int            `json:"total_count"`
	Photos     [][]*PhotoSize `json:"photos"`
}

// File - This object represents a file ready to be downloaded. The file can be downloaded via the
// link https://api.telegram.org/file/bot<token>/<file_path>. It is guaranteed that the link
// will be valid for at least 1 hour. When the link expires, a new one can be requested by
// calling getFile.
// Maximum file size to download is 20 MB
//
// file_id		String	Unique identifier for this file
// file_size	Integer	Optional. File size, if known
// file_path	String	Optional. File path. Use https://api.telegram.org/file/bot<token>/<file_path>
//						to get the file.
type File struct {
	FileID   string `json:"file_id"`
	FileSize int    `json:"file_size,omitempty"`
	FilePath string `json:"file_path,omitempty"`
}

// ReplyKeyboardMarkup - This object represents a custom keyboard with reply options
// (see Introduction to bots for details and examples).
//
// keyboard				Array of Array 		Array of button rows, each represented by an Array
//                 		of KeyboardButton	of KeyboardButton objects
// resize_keyboard		Boolean				Optional. Requests clients to resize the keyboard vertically
//                                			for optimal fit (e.g., make the keyboard smaller if there
//											are just two rows of buttons). Defaults to false, in which
//											case the custom keyboard is always of the same height as the
//											app's standard keyboard.
// one_time_keyboard	Boolean				Optional. Requests clients to hide the keyboard as soon as
//											it's been used. The keyboard will still be available, but
//											clients will automatically display the usual letter-keyboard
//											in the chat – the user can press a special button in the input
//											field to see the custom keyboard again. Defaults to false.
// selective			Boolean				Optional. Use this parameter if you want to show the keyboard
//											to specific users only. Targets: 1) users that are @mentioned
//											in the text of the Message object; 2) if the bot's message is
//											a reply (has reply_to_message_id), sender of the original message.
// Example: A user requests to change the bot‘s language, bot replies to the request with a keyboard to
// select the new language. Other users in the group don’t see the keyboard.
type ReplyKeyboardMarkup struct {
	Keyboard        [][]*KeyboardButton `json:"keyboard"`
	ResizeKeyboard  bool                `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard bool                `json:"one_time_keyboard,omitempty"`
	Selective       bool                `json:"selective,omitempty"`
}

// KeyboardButton - This object represents one button of the reply keyboard.
// For simple text buttons String can be used instead of this object to specify
// text of the button. Optional fields are mutually exclusive.
//
// text				String	Text of the button. If none of the optional fields are used, it will be sent
//							to the bot as a message when the button is pressed
// request_contact	Boolean	Optional. If True, the user's phone number will be sent as a contact when the
//							button is pressed. Available in private chats only
// request_location	Boolean	Optional. If True, the user's current location will be sent when the button is
//							pressed. Available in private chats only
// Note: request_contact and request_location options will only work in Telegram versions released after
// 9 April, 2016. Older clients will ignore them.
type KeyboardButton struct {
	Text            string `json:"text"`
	RequestContact  bool   `json:"request_contact,omitempty"`
	RequestLocation bool   `json:"request_location,omitempty"`
}

// ReplyKeyboardRemove - Upon receiving a message with this object, Telegram clients will remove the
// current custom keyboard and display the default letter-keyboard. By default, custom keyboards are
// displayed until a new keyboard is sent by a bot. An exception is made for one-time keyboards that
// are hidden immediately after the user presses a button (see ReplyKeyboardMarkup).
//
// remove_keyboard	True	Requests clients to remove the custom keyboard (user will not be able to
//							summon this keyboard; if you want to hide the keyboard from sight but keep
//							it accessible, use one_time_keyboard in ReplyKeyboardMarkup)
// selective		Boolean	Optional. Use this parameter if you want to remove the keyboard for specific
//							users only. Targets: 1) users that are @mentioned in the text of the Message
//							object; 2) if the bot's message is a reply (has reply_to_message_id), sender
//							of the original message.
// Example: A user votes in a poll, bot returns confirmation message in reply to the vote and removes the
// keyboard for that user, while still showing the keyboard with poll options to users who haven't voted yet.
type ReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"`
	Selective      bool `json:"selective,omitempty"`
}

// InlineKeyboardMarkup - This object represents an inline keyboard that appears right next to the message
// it belongs to.
//
// inline_keyboard	Array of Array of InlineKeyboardButton	Array of button rows, each represented by an
//															Array of InlineKeyboardButton objects
// Note: This will only work in Telegram versions released after 9 April, 2016. Older clients will display
// unsupported message.
type InlineKeyboardMarkup struct {
	InlineKeyboard [][]*InlineKeyboardButton `json:"inline_keyboard"`
}

// InlineKeyboardButton  - This object represents one button of an inline keyboard. You must use exactly
// one of the optional fields.
//
// text					String	Label text on the button
// url					String	Optional. HTTP url to be opened when button is pressed
// callback_data		String	Optional. Data to be sent in a callback query to the bot when button is
//								pressed, 1-64 bytes
// switch_inline_query	String	Optional. If set, pressing the button will prompt the user to select one
//								of their chats, open that chat and insert the bot‘s username and the
//								specified inline query in the input field. Can be empty, in which case just
//								the bot’s username will be inserted.
// switch_inline_query_current_chat
//						String	Optional. If set, pressing the button will insert the bot‘s username and
//								the specified inline query in the current chat's input field. Can be empty,
//								in which case only the bot’s username will be inserted.
// This offers a quick way for the user to open your bot in inline mode in the same chat – good for
// selecting something from multiple options.
// callback_game		CallbackGame
//								Optional. Description of the game that will be launched when the user
//								presses the button.
// NOTE: This type of button must always be the first button in the first row.
type InlineKeyboardButton struct {
	Text                         string        `json:"text"`
	URL                          string        `json:"url,omitempty"`
	CallbackData                 string        `json:"callback_data,omitempty"`
	SwitchInlineQuery            string        `json:"switch_inline_query,omitempty"`
	WwitchInlineQueryCurrentChat string        `json:"switch_inline_query_current_chat,omitempty"`
	CallbackGame                 *CallbackGame `json:"callback_game,omitempty"`
}

// CallbackQuery - This object represents an incoming callback query from a callback button in an inline
// keyboard. If the button that originated the query was attached to a message sent by the bot, the field
// message will be present. If the button was attached to a message sent via the bot (in inline mode), the
// field inline_message_id will be present. Exactly one of the fields data or game_short_name will be
// present.
//
// id					String	Unique identifier for this query
// from					User	Sender
// message				Message	Optional. Message with the callback button that originated the query.
//								Note that message content and message date will not be available if the
//								message is too old
// inline_message_id	String	Optional. Identifier of the message sent via the bot in inline mode, that
//								originated the query.
// chat_instance		String	Global identifier, uniquely corresponding to the chat to which the message
//								with the callback button was sent. Useful for high scores in games.
// data					String	Optional. Data associated with the callback button. Be aware that a bad
//								client can send arbitrary data in this field.
// game_short_name		String	Optional. Short name of a Game to be returned, serves as the unique
//								identifier for the game
// NOTE: After the user presses an inline button, Telegram clients will display a progress bar until
// you call answerCallbackQuery. It is, therefore, necessary to react by calling answerCallbackQuery even
// if no notification to the user is needed (e.g., without specifying any of the optional parameters).
type CallbackQuery struct {
	ID              string   `json:"id"`
	From            *User    `json:"from"`
	Message         *Message `json:"message,omitempty"`
	InlineMessageID string   `json:"inline_message_id,omitempty"`
	ChatInstance    string   `json:"chat_instance,omitempty"`
	Data            string   `json:"data,omitempty"`
	GameShortName   string   `json:"game_short_name,omitempty"`
}

// ForceReply - Upon receiving a message with this object, Telegram clients will display a reply
// interface to the user (act as if the user has selected the bot‘s message and tapped ’Reply'). This
// can be extremely useful if you want to create user-friendly step-by-step interfaces without having to
// sacrifice privacy mode.
//
// force_reply	True	Shows reply interface to the user, as if they manually selected the bot‘s message
//						and tapped ’Reply'
// selective	Boolean	Optional. Use this parameter if you want to force reply from specific users only.
//						Targets: 1) users that are @mentioned in the text of the Message object;
//						2) if the bot's message is a reply (has reply_to_message_id), sender of the
//						original message.
// Example: A poll bot for groups runs in privacy mode (only receives commands, replies to its messages
// and mentions). There could be two ways to create a new poll:
// Explain the user how to send a command with parameters (e.g. /newpoll question answer1 answer2).
// May be appealing for hardcore users but lacks modern day polish.
// Guide the user through a step-by-step process. ‘Please send me your question’, ‘Cool, now let’s add
// the first answer option‘, ’Great. Keep adding answer options, then send /done when you‘re ready’.
// The last option is definitely more attractive. And if you use ForceReply in your bot‘s questions, it
// will receive the user’s answers even if it only receives replies, commands and mentions — without any
// extra work for the user.
type ForceReply struct {
	ForceReply bool `json:"force_reply,omitempty"`
	Selective  bool `json:"selective,omitempty"`
}

// ChatMember - This object contains information about one member of the chat.
//
// user	    User	Information about the user
// status	String	The member's status in the chat. Can be “creator”, “administrator”, “member”, “left”
// 					or “kicked”
type ChatMember struct {
	User   *User  `json:"user"`
	Status string `json:"status"`
}

// ResponseParameters - Contains information about why a request was unsuccessfull.
//
// migrate_to_chat_id	Integer	Optional. The group has been migrated to a supergroup with the specified
//								identifier. This number may be greater than 32 bits and some programming
//								languages may have difficulty/silent defects in interpreting it. But it
//								is smaller than 52 bits, so a signed 64 bit integer or double-precision
//								float type are safe for storing this identifier.
// retry_after	        Integer	Optional. In case of exceeding flood control, the number of seconds left
//								to wait before the request can be repeated
type ResponseParameters struct {
	MigrateToChatID int `json:"migrate_to_chat_id,omitempty"`
	RetryAfter      int `json:"retry_after,omitempty"`
}
