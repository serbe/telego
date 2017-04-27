package telego

type sendMessageOpts struct {
	ChatID                string `json:"chat_id"`
	Text                  string `json:"text"`
	ParseMode             string `json:"parse_mode, omitempty"`
	DisableWebPagePreview bool   `json:"disable_web_page_preview, omitempty"`
	DisableNotification   bool   `json:"disable_notification, omitempty"`
	ReplyToMessageID      int    `json:"reply_to_message_id, omitempty"`
	// reply_markup	            Optional		Additional interface options. A JSON-serialized object for an inline
	//											keyboard, custom reply keyboard, instructions
	//              InlineKeyboardMarkup or     to remove reply keyboard or to force a reply from the user.
	//              ReplyKeyboardMarkup or
	//              ReplyKeyboardRemove or
	//              ForceReply
}

type forwardMessageOpts struct {
	ChatID              string `json:"chat_id"`
	FromChatID          string `json:"from_chat_id"`
	DisableNotification bool   `json:"disable_motification"`
	MessageID           int    `json:"message_id"`
}

type sendPhotoOpts struct {
	ChatID              string `json:"chat_id"`
	Photo               string `json:"photo"`
	Caption             string `json:"caption, omitempty"`
	DisableNotification bool   `json:"disable_notification, omitempty"`
	ReplyToMessageID    int    `json:"reply_to_message_id, omitempty"`
	// reply_markup				Optional	Additional interface options. A JSON-serialized object for an inline
	//			InlineKeyboardMarkup or 	keyboard, custom reply keyboard, instructions to remove reply keyboard
	//				ReplyKeyboardMarkup or  or to force a reply from the user.
	//				ReplyKeyboardRemove or
	//				ForceReply
}
