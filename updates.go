package telego

import (
	"encoding/json"
	"net/url"
	"strconv"
)

// Update - This object represents an incoming update.
// At most one of the optional parameters can be present in any given update.
//
// Field				Type				Description
// update_id			Integer				The update‘s unique identifier. Update identifiers start from a certain
//											positive number and increase sequentially. This ID becomes especially
//											handy if you’re using Webhooks, since it allows you to ignore repeated
//											updates or to restore the correct update sequence, should they get out
//											of order.
// message				Message				Optional. New incoming message of any kind — text, photo, sticker, etc.
// edited_message		Message				Optional. New version of a message that is known to the bot and was edited
// channel_post			Message				Optional. New incoming channel post of any kind — text, photo, sticker, etc.
// edited_channel_post	Message				Optional. New version of a channel post that is known to the bot and was edited
// inline_query			InlineQuery			Optional. New incoming inline query
// chosen_inline_result	ChosenInlineResult	Optional. The result of an inline query that was chosen by a user and sent
//											to their chat partner.
// callback_query		CallbackQuery		Optional. New incoming callback query
// shipping_query		ShippingQuery		Optional. New incoming shipping query. Only for invoices with flexible price
// pre_checkout_query	PreCheckoutQuery	Optional. New incoming pre-checkout query. Contains full information about checkout
type Update struct {
	UpdateID           int64               `json:"update_id"`
	Message            *Message            `json:"message,omitempty"`
	EditedMessage      *Message            `json:"edited_message,omitempty"`
	ChannelPost        *Message            `json:"channel_post,omitempty"`
	EditedChannelPost  *Message            `json:"edited_channel_post,omitempty"`
	InlineQuery        *InlineQuery        `json:"inline_query,omitempty"`
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"`
	CallbackQuery      *CallbackQuery      `json:"callback_query,omitempty"`
	ShippingQuery      *ShippingQuery      `json:"shipping_query,omitempty"`
	PreCheckoutQuery   *PreCheckoutQuery   `json:"pre_checkout_query,omitempty"`
}

// GetUpdates - "getUpdates" Use this method to receive incoming updates using long polling (wiki). An Array of
// Update objects is returned.
//
// Parameters		Type			Required	Description
// offset			Integer			Optional	Identifier of the first update to be returned. Must be greater by
//												one than the highest among the identifiers of previously received
//												updates. By default, updates starting with the earliest unconfirmed
//												update are returned. An update is considered confirmed as soon as
//												getUpdates is called with an offset higher than its update_id. The
//												negative offset can be specified to retrieve updates starting from
//												-offset update from the end of the updates queue. All previous
//												updates will forgotten.
// limit			Integer			Optional	Limits the number of updates to be retrieved. Values between 1—100
//												are accepted. Defaults to 100.
// timeout			Integer			Optional	Timeout in seconds for long polling. Defaults to 0, i.e. usual
//												short polling. Should be positive, short polling should be used for
//												testing purposes only.
// allowed_updates	Array of String	Optional	List the types of updates you want your bot to receive. For example,
//												specify [“message”, “edited_channel_post”, “callback_query”] to only
//												receive updates of these types. See Update for a complete list of
//												available update types. Specify an empty list to receive all updates
//												regardless of type (default). If not specified, the previous setting
//												will be used.
// Please note that this parameter doesn't affect updates created before the call to the getUpdates, so unwanted
// updates may be received for a short period of time.
// Notes
// 1. This method will not work if an outgoing webhook is set up.
// 2. In order to avoid getting duplicate updates, recalculate offset after each server response.
func (t *Telebot) GetUpdates(opt *GetUpdatesOpt) ([]Update, error) {
	values := url.Values{}
	if opt.Offset != 0 {
		values.Set("offset", strconv.Itoa(opt.Offset))
	}
	if opt.Limit > 0 && opt.Limit < 101 {
		values.Set("limit", strconv.Itoa(opt.Limit))
	} else {
		values.Set("limit", "100")
	}
	if opt.Timeout > 0 {
		values.Set("timeout", strconv.Itoa(opt.Timeout))
	} else {
		values.Set("timeout", "0")
	}
	if len(opt.AllowedUpdates) > 0 {
		values["allowed_updates"] = opt.AllowedUpdates
	}
	r, err := t.createResponse("getUpdates", values)
	if err != nil {
		errLog("GetUpdates createResponse", err)
		return nil, err
	}
	var updates []Update
	err = json.Unmarshal(r.Result, &updates)
	if err != nil {
		errLog("GetUpdates Unmarshal", err)
	}
	return updates, nil
}

// SetWebhook - "setWebhook" Use this method to specify a url and receive incoming updates via an outgoing webhook.
// Whenever there is an update for the bot, we will send an HTTPS POST request to the specified url, containing a
// JSON-serialized Update. In case of an unsuccessful request, we will give up after a reasonable amount of attempts.
// Returns true.
//
// If you'd like to make sure that the Webhook request comes from Telegram, we recommend using a secret path in
// the URL, e.g. https://www.example.com/<token>. Since nobody else knows your bot‘s token, you can be pretty
// sure it’s us.
//
// Parameters		Type		Required	Description
// url				String		Yes			HTTPS url to send updates to. Use an empty string to remove webhook
// 											integration
// certificate		InputFile	Optional	Upload your public key certificate so that the root certificate in
//											use can be checked. See our self-signed guide for details.
// max_connections	Integer		Optional	Maximum allowed number of simultaneous HTTPS connections to the
//											webhook for update delivery, 1-100. Defaults to 40. Use lower values
//											to limit the load on your bot‘s server, and higher values to increase
//											your bot’s throughput.
// allowed_updates	Array		Optional	List the types of updates you want your bot to receive. For example,
//					of String				specify [“message”, “edited_channel_post”, “callback_query”] to only
//											receive updates of these types. See Update for a complete list of
//											available update types. Specify an empty list to receive all updates
//											regardless of type (default). If not specified, the previous setting
//											will be used.
//
// Please note that this parameter doesn't affect updates created before the call to the setWebhook, so unwanted
// updates may be received for a short period of time.
func (t *Telebot) SetWebhook(URL string, opt *SetWebhookOpt) error {
	values := url.Values{}
	values.Set("url", URL)
	if opt.Certificate != "" {
		values.Set("certificate", opt.Certificate)
	}
	if opt.MaxConnections > 0 {
		values.Set("max_connections", strconv.Itoa(opt.MaxConnections))
	}
	if len(opt.AllowedUpdates) > 0 {
		values["allowed_updates"] = opt.AllowedUpdates
	}
	r, err := t.createResponse("setWebhook", values)
	if err != nil {
		errLog("SetWebhook createResponse", err)
		return err
	}
	var result bool
	err = json.Unmarshal(r.Result, &result)
	if err != nil {
		errLog("GetUpdates Unmarshal", err)
	}
	return err
}

// deleteWebhook
// Use this method to remove webhook integration if you decide to switch back to getUpdates. Returns True on success. Requires no parameters.

// getWebhookInfo
// Use this method to get current webhook status. Requires no parameters. On success, returns a WebhookInfo object. If the bot is using getUpdates, will return an object with the url field empty.

// WebhookInfo
// Contains information about the current status of a webhook.
//
// Field					Type		Description
// url						String		Webhook URL, may be empty if webhook is not set up
// has_custom_certificate	Boolean		True, if a custom certificate was provided for webhook certificate checks
// pending_update_count		Integer		Number of updates awaiting delivery
// last_error_date			Integer		Optional. Unix time for the most recent error that happened when trying to
//										deliver an update via webhook
// last_error_message		String		Optional. Error message in human-readable format for the most recent error
//										that happened when trying to deliver an update via webhook
// max_connections			Integer		Optional. Maximum allowed number of simultaneous HTTPS connections to the
//										webhook for update delivery
// allowed_updates	Array of String		Optional. A list of update types the bot is subscribed to. Defaults to all
//										update types
