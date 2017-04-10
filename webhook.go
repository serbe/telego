package telego

// setWebhook
// Use this method to specify a url and receive incoming updates via an outgoing webhook. Whenever there is an
// update for the bot, we will send an HTTPS POST request to the specified url, containing a JSON-serialized
// Update. In case of an unsuccessful request, we will give up after a reasonable amount of attempts. Returns true.
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
// allowed_updates	Array of String
//								Optional	List the types of updates you want your bot to receive. For example,
//											specify [“message”, “edited_channel_post”, “callback_query”] to only
//											receive updates of these types. See Update for a complete list of
//											available update types. Specify an empty list to receive all updates
//											regardless of type (default). If not specified, the previous setting
//											will be used.
//
// Please note that this parameter doesn't affect updates created before the call to the setWebhook, so unwanted
// updates may be received for a short period of time.

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
