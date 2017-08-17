package telego

// SendMessageOpts - options for SendMessage
type SendMessageOpts struct {
	ChatID                string `json:"chat_id"`
	Text                  string `json:"text"`
	ParseMode             string `json:"parse_mode,omitempty"`
	DisableWebPagePreview bool   `json:"disable_web_page_preview,omitempty"`
	DisableNotification   bool   `json:"disable_notification,omitempty"`
	ReplyToMessageID      int    `json:"reply_to_message_id,omitempty"`
	// reply_markup	            Optional		Additional interface options. A JSON-serialized object for an inline
	//											keyboard, custom reply keyboard, instructions
	//              InlineKeyboardMarkup or     to remove reply keyboard or to force a reply from the user.
	//              ReplyKeyboardMarkup or
	//              ReplyKeyboardRemove or
	//              ForceReply
}

// ForwardMessageOpts - options forForwardMessage
type ForwardMessageOpts struct {
	ChatID              string `json:"chat_id"`
	FromChatID          string `json:"from_chat_id"`
	DisableNotification bool   `json:"disable_motification"`
	MessageID           int    `json:"message_id"`
}

// SendPhotoOpts - options for SendPhoto
type SendPhotoOpts struct {
	ChatID              string `json:"chat_id"`
	Photo               string `json:"photo"`
	Caption             string `json:"caption,omitempty"`
	DisableNotification bool   `json:"disable_notification,omitempty"`
	ReplyToMessageID    int    `json:"reply_to_message_id,omitempty"`
	// reply_markup				Optional	Additional interface options. A JSON-serialized object for an inline
	//			InlineKeyboardMarkup or 	keyboard, custom reply keyboard, instructions to remove reply keyboard
	//				ReplyKeyboardMarkup or  or to force a reply from the user.
	//				ReplyKeyboardRemove or
	//				ForceReply
}

// GetUpdatesOpt - options for GetUpdates
type GetUpdatesOpt struct {
	Offset         int      `json:"offset,omitempty"`
	Limit          int      `json:"limit,omitempty"`
	Timeout        int      `json:"timeout,omitempty"`
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}

// SetWebhookOpt - options for SetWebhook
type SetWebhookOpt struct {
	URL            string   `json:"url"`
	Certificate    string   `json:"certificate,omitempty"`
	MaxConnections int      `json:"max_connections,omitempty"`
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}

// WebhookInfoOpt  - options for WebhookInfo
type WebhookInfoOpt struct {
	URL                  string   `json:"url"`
	HasCustomCertificate bool     `json:"has_custom_certificate"`
	PendingUpdateCount   int      `json:"pending_update_count"`
	LastErrorDate        int      `json:"last_error_date,omitempty"`
	LastErrorMessage     string   `json:"last_error_message,omitempty"`
	MaxConnections       int      `json:"max_connections,omitempty"`
	AllowedUpdates       []string `json:"allowed_updates,omitempty"`
}

// SendAudioOpt - options for SendAudio
type SendAudioOpt struct {
	ChatID              string `json:"chat_id"`
	Audio               string `json:"audio"`
	Caption             string `json:"caption,omitempty"`
	Duration            int    `json:"duration,omitempty"`
	Performer           string `json:"performer,omitempty"`
	Title               string `json:"title,omitempty"`
	DisableNotification bool   `json:"disable_notification,omitempty"`
	ReplyToMessageID    int    `json:"reply_to_message_id,omitempty"`
	// reply_markup				Optional	Additional interface options. A JSON-serialized object for an inline
	//				InlineKeyboardMarkup	keyboard, custom reply keyboard, instructions to remove reply keyboard
	//				or ReplyKeyboardMarkup	or to force a reply from the user.
	//				or ReplyKeyboardRemove
	//				or ForceReply
}

// SendDocumentOpt - options for sendDocument
type SendDocumentOpt struct {
	ChatID              string `json:"chat_id"`
	Document            string `json:"document"`
	Caption             string `json:"caption,omitempty"`
	DisableNotification bool   `json:"disable_notification,omitempty"`
	ReplyToMessageID    int    `json:"reply_to_message_id,omitempty"`
	// reply_markup				Optional	Additional interface options. A JSON-serialized object for an inline
	//				InlineKeyboardMarkup	keyboard, custom reply keyboard, instructions to remove reply keyboard
	//				or ReplyKeyboardMarkup	or to force a reply from the user.
	//				or ReplyKeyboardRemove
	//				or ForceReply
}

// SendVideoOpt - options for sendVideo
type SendVideoOpt struct {
	ChatID              string `json:"chat_id"`
	Video               string `json:"video"`
	Duration            int    `json:"duration,omitempty"`
	Width               int    `json:"width,omitempty"`
	Height              int    `json:"height,omitempty"`
	Caption             string `json:"caption,omitempty"`
	DisableNotification bool   `json:"disable_notification,omitempty"`
	ReplyToMessageID    int    `json:"reply_to_message_id,omitempty"`
	// reply_markup					Optional	Additional interface options. A JSON-serialized object for an inline keyboard,
	//			InlineKeyboardMarkup 			custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
	//			or ReplyKeyboardMarkup
	//			or ReplyKeyboardRemove
	//			or ForceReply
}

// SendVoiceOpt - options for sendVoice
type SendVoiceOpt struct {
	ChatID              string `json:"chat_id"`
	Voice               string `json:"voice"`
	Caption             string `json:"caption,omitempty"`
	Duration            int    `json:"duration,omitempty"`
	DisableNotification bool   `json:"disable_notification,omitempty"`
	ReplyToMessageID    int    `json:"reply_to_message_id,omitempty"`
	// reply_markup					Optional	Additional interface options. A JSON-serialized object for an inline keyboard,
	//			InlineKeyboardMarkup 			custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
	//			or ReplyKeyboardMarkup
	//			or ReplyKeyboardRemove
	//			or ForceReply
}

// SendVideoNoteOpt - options for sendVideo
type SendVideoNoteOpt struct {
	ChatID              string `json:"chat_id"`
	VideoNote           string `json:"video_note"`
	Duration            int    `json:"duration,omitempty"`
	Length              int    `json:"length,omitempty"`
	DisableNotification bool   `json:"disable_notification,omitempty"`
	ReplyToMessageID    int    `json:"reply_to_message_id,omitempty"`
	// reply_markup					Optional	Additional interface options. A JSON-serialized object for an inline keyboard,
	//			InlineKeyboardMarkup 			custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
	//			or ReplyKeyboardMarkup
	//			or ReplyKeyboardRemove
	//			or ForceReply
}

// SendLocationOpt - option for sendLocation
type SendLocationOpt struct {
	ChatID              string  `json:"chat_id"`
	Latitude            float64 `json:"latitude"`
	Longitude           float64 `json:"longitude"`
	DisableNotification bool    `json:"disable_notification,omitempty"`
	ReplyToMessageID    int     `json:"reply_to_message_id,omitempty"`
	// reply_markup		InlineKeyboardMarkup 	Additional interface options. A JSON-serialized object for an inline
	//				or ReplyKeyboardMarkup 		keyboard, custom reply keyboard, instructions to remove reply keyboard
	//				or ReplyKeyboardRemove 		or to force a reply from the user.
	//				or ForceReply	Optional
}

// SendVenueOpt - option for SendVenue
type SendVenueOpt struct {
	ChatID              string  `json:"chat_id"`
	Latitude            float64 `json:"latitude"`
	Longitude           float64 `json:"longitude"`
	Title               string  `json:"title"`
	Address             string  `json:"address"`
	FoursquareID        string  `json:"foursquare_id,omitempty"`
	DisableNotification bool    `json:"disable_notification,omitempty"`
	ReplyToMessageID    int     `json:"reply_to_message_id,omitempty"`
	// reply_markup		InlineKeyboardMarkup 	Additional interface options. A JSON-serialized object for an inline
	//				or ReplyKeyboardMarkup 		keyboard, custom reply keyboard, instructions to remove reply keyboard
	//				or ReplyKeyboardRemove 		or to force a reply from the user.
	//				or ForceReply	Optional
}

// ---------------------

type SendGameOpt struct {
	ChatID              string `json:"chat_id"`
	GameShortName       string `json:"game_short_name"`
	DisableNotification bool   `json:"disable_notification,omitempty"`
	ReplyToMessageID    int    `json:"reply_to_message_id,omitempty"`
}
