package telego

// InlineQuery - This object represents an incoming inline query. When the user sends an empty query, your bot
// could return some default or trending results.
//
// Field	Type		Description
// id		String		Unique identifier for this query
// from		User		Sender
// location	Location	Optional. Sender location, only for bots that request user location
// query	String		Text of the query (up to 512 characters)
// offset	String		Offset of the results to be returned, can be controlled by the bot
type InlineQuery struct {
	ID       string    `json:"id"`
	From     *User     `json:"from"`
	Location *Location `json:"location,omitempty"`
	Query    string    `json:"query"`
	Offset   string    `json:"offset"`
}

// answerInlineQuery
// Use this method to send answers to an inline query. On success, True is returned.
// No more than 50 results per query are allowed.
//
// Parameters	Type	Required	Description
// inline_query_id	String	Yes	Unique identifier for the answered query
// results	Array of InlineQueryResult	Yes	A JSON-serialized array of results for the inline query
// cache_time	Integer	Optional	The maximum amount of time in seconds that the result of the inline query may be cached on the server. Defaults to 300.
// is_personal	Boolean	Optional	Pass True, if results may be cached on the server side only for the user that sent the query. By default, results may be returned to any user who sends the same query
// next_offset	String	Optional	Pass the offset that a client should send in the next query with the same text to receive more results. Pass an empty string if there are no more results or if you don‘t support pagination. Offset length can’t exceed 64 bytes.
// switch_pm_text	String	Optional	If passed, clients will display a button with specified text that switches the user to a private chat with the bot and sends the bot a start message with the parameter switch_pm_parameter
// switch_pm_parameter	String	Optional	Deep-linking parameter for the /start message sent to the bot when user presses the switch button. 1-64 characters, only A-Z, a-z, 0-9, _ and - are allowed.
//
// Example: An inline bot that sends YouTube videos can ask the user to connect the bot to their YouTube account to adapt search results accordingly. To do this, it displays a ‘Connect your YouTube account’ button above the results, or even before showing any. The user presses the button, switches to a private chat with the bot and, in doing so, passes a start parameter that instructs the bot to return an oauth link. Once done, the bot can offer a switch_inline button so that the user can easily return to the chat where they wanted to use the bot's inline capabilities.

// InlineQueryResult
// This object represents one result of an inline query. Telegram clients currently support results of the following 20 types:
//
// InlineQueryResultCachedAudio
// InlineQueryResultCachedDocument
// InlineQueryResultCachedGif
// InlineQueryResultCachedMpeg4Gif
// InlineQueryResultCachedPhoto
// InlineQueryResultCachedSticker
// InlineQueryResultCachedVideo
// InlineQueryResultCachedVoice
// InlineQueryResultArticle
// InlineQueryResultAudio
// InlineQueryResultContact
// InlineQueryResultGame
// InlineQueryResultDocument
// InlineQueryResultGif
// InlineQueryResultLocation
// InlineQueryResultMpeg4Gif
// InlineQueryResultPhoto
// InlineQueryResultVenue
// InlineQueryResultVideo
// InlineQueryResultVoice
// InlineQueryResultArticle
// Represents a link to an article or web page.
//
// Field	Type	Description
// type	String	Type of the result, must be article
// id	String	Unique identifier for this result, 1-64 Bytes
// title	String	Title of the result
// input_message_content	InputMessageContent	Content of the message to be sent
// reply_markup	InlineKeyboardMarkup	Optional. Inline keyboard attached to the message
// url	String	Optional. URL of the result
// hide_url	Boolean	Optional. Pass True, if you don't want the URL to be shown in the message
// description	String	Optional. Short description of the result
// thumb_url	String	Optional. Url of the thumbnail for the result
// thumb_width	Integer	Optional. Thumbnail width
// thumb_height	Integer	Optional. Thumbnail height

// InlineQueryResultPhoto
// Represents a link to a photo. By default, this photo will be sent by the user with optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the photo.
//
// Field	Type	Description
// type	String	Type of the result, must be photo
// id	String	Unique identifier for this result, 1-64 bytes
// photo_url	String	A valid URL of the photo. Photo must be in jpeg format. Photo size must not exceed 5MB
// thumb_url	String	URL of the thumbnail for the photo
// photo_width	Integer	Optional. Width of the photo
// photo_height	Integer	Optional. Height of the photo
// title	String	Optional. Title for the result
// description	String	Optional. Short description of the result
// caption	String	Optional. Caption of the photo to be sent, 0-200 characters
// reply_markup	InlineKeyboardMarkup	Optional. Inline keyboard attached to the message
// input_message_content	InputMessageContent	Optional. Content of the message to be sent instead of the photo

// InlineQueryResultGif
// Represents a link to an animated GIF file. By default, this animated GIF file will be sent by the user with optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
//
// Field	Type	Description
// type	String	Type of the result, must be gif
// id	String	Unique identifier for this result, 1-64 bytes
// gif_url	String	A valid URL for the GIF file. File size must not exceed 1MB
// gif_width	Integer	Optional. Width of the GIF
// gif_height	Integer	Optional. Height of the GIF
// thumb_url	String	URL of the static thumbnail for the result (jpeg or gif)
// title	String	Optional. Title for the result
// caption	String	Optional. Caption of the GIF file to be sent, 0-200 characters
// reply_markup	InlineKeyboardMarkup	Optional. Inline keyboard attached to the message
// input_message_content	InputMessageContent	Optional. Content of the message to be sent instead of the GIF animation

// InlineQueryResultMpeg4Gif
// Represents a link to a video animation (H.264/MPEG-4 AVC video without sound). By default, this animated MPEG-4 file will be sent by the user with optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
//
// Field	Type	Description
// type	String	Type of the result, must be mpeg4_gif
// id	String	Unique identifier for this result, 1-64 bytes
// mpeg4_url	String	A valid URL for the MP4 file. File size must not exceed 1MB
// mpeg4_width	Integer	Optional. Video width
// mpeg4_height	Integer	Optional. Video height
// thumb_url	String	URL of the static thumbnail (jpeg or gif) for the result
// title	String	Optional. Title for the result
// caption	String	Optional. Caption of the MPEG-4 file to be sent, 0-200 characters
// reply_markup	InlineKeyboardMarkup	Optional. Inline keyboard attached to the message
// input_message_content	InputMessageContent	Optional. Content of the message to be sent instead of the video animation

// InlineQueryResultVideo
// Represents a link to a page containing an embedded video player or a video file. By default, this video file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the video.
//
// Field	Type	Description
// type	String	Type of the result, must be video
// id	String	Unique identifier for this result, 1-64 bytes
// video_url	String	A valid URL for the embedded video player or video file
// mime_type	String	Mime type of the content of video url, “text/html” or “video/mp4”
// thumb_url	String	URL of the thumbnail (jpeg only) for the video
// title	String	Title for the result
// caption	String	Optional. Caption of the video to be sent, 0-200 characters
// video_width	Integer	Optional. Video width
// video_height	Integer	Optional. Video height
// video_duration	Integer	Optional. Video duration in seconds
// description	String	Optional. Short description of the result
// reply_markup	InlineKeyboardMarkup	Optional. Inline keyboard attached to the message
// input_message_content	InputMessageContent	Optional. Content of the message to be sent instead of the video

// InlineQueryResultAudio
// Represents a link to an mp3 audio file. By default, this audio file will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the audio.
//
// Field	Type	Description
// type	String	Type of the result, must be audio
// id	String	Unique identifier for this result, 1-64 bytes
// audio_url	String	A valid URL for the audio file
// title	String	Title
// caption	String	Optional. Caption, 0-200 characters
// performer	String	Optional. Performer
// audio_duration	Integer	Optional. Audio duration in seconds
// reply_markup	InlineKeyboardMarkup	Optional. Inline keyboard attached to the message
// input_message_content	InputMessageContent	Optional. Content of the message to be sent instead of the audio
// Note: This will only work in Telegram versions released after 9 April, 2016. Older clients will ignore them.

// InlineQueryResultVoice
// Represents a link to a voice recording in an .ogg container encoded with OPUS. By default, this voice recording will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the the voice message.
//
// Field	Type	Description
// type	String	Type of the result, must be voice
// id	String	Unique identifier for this result, 1-64 bytes
// voice_url	String	A valid URL for the voice recording
// title	String	Recording title
// caption	String	Optional. Caption, 0-200 characters
// voice_duration	Integer	Optional. Recording duration in seconds
// reply_markup	InlineKeyboardMarkup	Optional. Inline keyboard attached to the message
// input_message_content	InputMessageContent	Optional. Content of the message to be sent instead of the voice recording
// Note: This will only work in Telegram versions released after 9 April, 2016. Older clients will ignore them.

// InlineQueryResultDocument
// Represents a link to a file. By default, this file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the file. Currently, only .PDF and .ZIP files can be sent using this method.
//
// Field	Type	Description
// type	String	Type of the result, must be document
// id	String	Unique identifier for this result, 1-64 bytes
// title	String	Title for the result
// caption	String	Optional. Caption of the document to be sent, 0-200 characters
// document_url	String	A valid URL for the file
// mime_type	String	Mime type of the content of the file, either “application/pdf” or “application/zip”
// description	String	Optional. Short description of the result
// reply_markup	InlineKeyboardMarkup	Optional. Inline keyboard attached to the message
// input_message_content	InputMessageContent	Optional. Content of the message to be sent instead of the file
// thumb_url	String	Optional. URL of the thumbnail (jpeg only) for the file
// thumb_width	Integer	Optional. Thumbnail width
// thumb_height	Integer	Optional. Thumbnail height
// Note: This will only work in Telegram versions released after 9 April, 2016. Older clients will ignore them.

// InlineQueryResultLocation
// Represents a location on a map. By default, the location will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the location.
//
// Field	Type	Description
// type	String	Type of the result, must be location
// id	String	Unique identifier for this result, 1-64 Bytes
// latitude	Float number	Location latitude in degrees
// longitude	Float number	Location longitude in degrees
// title	String	Location title
// reply_markup	InlineKeyboardMarkup	Optional. Inline keyboard attached to the message
// input_message_content	InputMessageContent	Optional. Content of the message to be sent instead of the location
// thumb_url	String	Optional. Url of the thumbnail for the result
// thumb_width	Integer	Optional. Thumbnail width
// thumb_height	Integer	Optional. Thumbnail height
// Note: This will only work in Telegram versions released after 9 April, 2016. Older clients will ignore them.

// InlineQueryResultVenue
// Represents a venue. By default, the venue will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the venue.
//
// Field	Type	Description
// type	String	Type of the result, must be venue
// id	String	Unique identifier for this result, 1-64 Bytes
// latitude	Float	Latitude of the venue location in degrees
// longitude	Float	Longitude of the venue location in degrees
// title	String	Title of the venue
// address	String	Address of the venue
// foursquare_id	String	Optional. Foursquare identifier of the venue if known
// reply_markup	InlineKeyboardMarkup	Optional. Inline keyboard attached to the message
// input_message_content	InputMessageContent	Optional. Content of the message to be sent instead of the venue
// thumb_url	String	Optional. Url of the thumbnail for the result
// thumb_width	Integer	Optional. Thumbnail width
// thumb_height	Integer	Optional. Thumbnail height
// Note: This will only work in Telegram versions released after 9 April, 2016. Older clients will ignore them.

// InlineQueryResultContact
// Represents a contact with a phone number. By default, this contact will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the contact.
//
// Field	Type	Description
// type	String	Type of the result, must be contact
// id	String	Unique identifier for this result, 1-64 Bytes
// phone_number	String	Contact's phone number
// first_name	String	Contact's first name
// last_name	String	Optional. Contact's last name
// reply_markup	InlineKeyboardMarkup	Optional. Inline keyboard attached to the message
// input_message_content	InputMessageContent	Optional. Content of the message to be sent instead of the contact
// thumb_url	String	Optional. Url of the thumbnail for the result
// thumb_width	Integer	Optional. Thumbnail width
// thumb_height	Integer	Optional. Thumbnail height
// Note: This will only work in Telegram versions released after 9 April, 2016. Older clients will ignore them.

// InlineQueryResultGame
// Represents a Game.
//
// Field	Type	Description
// type	String	Type of the result, must be game
// id	String	Unique identifier for this result, 1-64 bytes
// game_short_name	String	Short name of the game
// reply_markup	InlineKeyboardMarkup	Optional. Inline keyboard attached to the message
// Note: This will only work in Telegram versions released after October 1, 2016. Older clients will not display any inline results if a game result is among them.

// InlineQueryResultCachedPhoto
// Represents a link to a photo stored on the Telegram servers. By default, this photo will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the photo.
//
// Field	Type	Description
// type	String	Type of the result, must be photo
// id	String	Unique identifier for this result, 1-64 bytes
// photo_file_id	String	A valid file identifier of the photo
// title	String	Optional. Title for the result
// description	String	Optional. Short description of the result
// caption	String	Optional. Caption of the photo to be sent, 0-200 characters
// reply_markup	InlineKeyboardMarkup	Optional. Inline keyboard attached to the message
// input_message_content	InputMessageContent	Optional. Content of the message to be sent instead of the photo

// InlineQueryResultCachedGif
// Represents a link to an animated GIF file stored on the Telegram servers. By default, this animated GIF file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with specified content instead of the animation.
//
// Field	Type	Description
// type	String	Type of the result, must be gif
// id	String	Unique identifier for this result, 1-64 bytes
// gif_file_id	String	A valid file identifier for the GIF file
// title	String	Optional. Title for the result
// caption	String	Optional. Caption of the GIF file to be sent, 0-200 characters
// reply_markup	InlineKeyboardMarkup	Optional. Inline keyboard attached to the message
// input_message_content	InputMessageContent	Optional. Content of the message to be sent instead of the GIF animation

// InlineQueryResultCachedMpeg4Gif
// Represents a link to a video animation (H.264/MPEG-4 AVC video without sound) stored on the Telegram servers. By default, this animated MPEG-4 file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
//
// Field	Type	Description
// type	String	Type of the result, must be mpeg4_gif
// id	String	Unique identifier for this result, 1-64 bytes
// mpeg4_file_id	String	A valid file identifier for the MP4 file
// title	String	Optional. Title for the result
// caption	String	Optional. Caption of the MPEG-4 file to be sent, 0-200 characters
// reply_markup	InlineKeyboardMarkup	Optional. Inline keyboard attached to the message
// input_message_content	InputMessageContent	Optional. Content of the message to be sent instead of the video animation

// InlineQueryResultCachedSticker
// Represents a link to a sticker stored on the Telegram servers. By default, this sticker will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the sticker.
//
// Field	Type	Description
// type	String	Type of the result, must be sticker
// id	String	Unique identifier for this result, 1-64 bytes
// sticker_file_id	String	A valid file identifier of the sticker
// reply_markup	InlineKeyboardMarkup	Optional. Inline keyboard attached to the message
// input_message_content	InputMessageContent	Optional. Content of the message to be sent instead of the sticker
// Note: This will only work in Telegram versions released after 9 April, 2016. Older clients will ignore them.

// InlineQueryResultCachedDocument
// Represents a link to a file stored on the Telegram servers. By default, this file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the file.
//
// Field	Type	Description
// type	String	Type of the result, must be document
// id	String	Unique identifier for this result, 1-64 bytes
// title	String	Title for the result
// document_file_id	String	A valid file identifier for the file
// description	String	Optional. Short description of the result
// caption	String	Optional. Caption of the document to be sent, 0-200 characters
// reply_markup	InlineKeyboardMarkup	Optional. Inline keyboard attached to the message
// input_message_content	InputMessageContent	Optional. Content of the message to be sent instead of the file
// Note: This will only work in Telegram versions released after 9 April, 2016. Older clients will ignore them.

// InlineQueryResultCachedVideo
// Represents a link to a video file stored on the Telegram servers. By default, this video file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the video.
//
// Field	Type	Description
// type	String	Type of the result, must be video
// id	String	Unique identifier for this result, 1-64 bytes
// video_file_id	String	A valid file identifier for the video file
// title	String	Title for the result
// description	String	Optional. Short description of the result
// caption	String	Optional. Caption of the video to be sent, 0-200 characters
// reply_markup	InlineKeyboardMarkup	Optional. Inline keyboard attached to the message
// input_message_content	InputMessageContent	Optional. Content of the message to be sent instead of the video

// InlineQueryResultCachedVoice
// Represents a link to a voice message stored on the Telegram servers. By default, this voice message will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the voice message.
//
// Field	Type	Description
// type	String	Type of the result, must be voice
// id	String	Unique identifier for this result, 1-64 bytes
// voice_file_id	String	A valid file identifier for the voice message
// title	String	Voice message title
// caption	String	Optional. Caption, 0-200 characters
// reply_markup	InlineKeyboardMarkup	Optional. Inline keyboard attached to the message
// input_message_content	InputMessageContent	Optional. Content of the message to be sent instead of the voice message
// Note: This will only work in Telegram versions released after 9 April, 2016. Older clients will ignore them.

// InlineQueryResultCachedAudio
// Represents a link to an mp3 audio file stored on the Telegram servers. By default, this audio file will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the audio.
//
// Field	Type	Description
// type	String	Type of the result, must be audio
// id	String	Unique identifier for this result, 1-64 bytes
// audio_file_id	String	A valid file identifier for the audio file
// caption	String	Optional. Caption, 0-200 characters
// reply_markup	InlineKeyboardMarkup	Optional. Inline keyboard attached to the message
// input_message_content	InputMessageContent	Optional. Content of the message to be sent instead of the audio
// Note: This will only work in Telegram versions released after 9 April, 2016. Older clients will ignore them.

// InputMessageContent
// This object represents the content of a message to be sent as a result of an inline query. Telegram clients currently support the following 4 types:
//
// InputTextMessageContent
// InputLocationMessageContent
// InputVenueMessageContent
// InputContactMessageContent
// InputTextMessageContent
// Represents the content of a text message to be sent as the result of an inline query.
//
// Field	Type	Description
// message_text	String	Text of the message to be sent, 1-4096 characters
// parse_mode	String	Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in your bot's message.
// disable_web_page_preview	Boolean	Optional. Disables link previews for links in the sent message

// InputLocationMessageContent
// Represents the content of a location message to be sent as the result of an inline query.
//
// Field	Type	Description
// latitude	Float	Latitude of the location in degrees
// longitude	Float	Longitude of the location in degrees
// Note: This will only work in Telegram versions released after 9 April, 2016. Older clients will ignore them.

// InputVenueMessageContent
// Represents the content of a venue message to be sent as the result of an inline query.
//
// Field	Type	Description
// latitude	Float	Latitude of the venue in degrees
// longitude	Float	Longitude of the venue in degrees
// title	String	Name of the venue
// address	String	Address of the venue
// foursquare_id	String	Optional. Foursquare identifier of the venue, if known
// Note: This will only work in Telegram versions released after 9 April, 2016. Older clients will ignore them.

// InputContactMessageContent
// Represents the content of a contact message to be sent as the result of an inline query.
//
// Field	Type	Description
// phone_number	String	Contact's phone number
// first_name	String	Contact's first name
// last_name	String	Optional. Contact's last name
// Note: This will only work in Telegram versions released after 9 April, 2016. Older clients will ignore them.

// ChosenInlineResult - Represents a result of an inline query that was chosen by the user and sent to their
// chat partner.
//
// Field				Type		Description
// result_id			String		The unique identifier for the result that was chosen
// from					User		The user that chose the result
// location				Location	Optional. Sender location, only for bots that require user location
// inline_message_id	String		Optional. Identifier of the sent inline message. Available only if there is an
//									inline keyboard attached to the message. Will be also received in callback
//									queries and can be used to edit the message.
// query				String		The query that was used to obtain the result
type ChosenInlineResult struct {
	ResultID        string    `json:"result_id"`
	From            *User     `json:"from"`
	Location        *Location `json:"location,omitempty"`
	InlineMessageID string    `json:"inline_message_id,omitempty"`
	Query           string    `json:"query"`
}
