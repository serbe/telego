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
	Certificate    string   `json:"certificate,omitempty"`
	MaxConnections int      `json:"max_connections,omitempty"`
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}
