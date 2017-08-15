package telego

// Invoice - This object contains basic information about an invoice.
//
// title			String	Product name
// description		String	Product description
// start_parameter	String	Unique bot deep-linking parameter that can be used to generate this invoice
// currency			String	Three-letter ISO 4217 currency code
// total_amount		Integer	Total price in the smallest units of the currency (integer, not float/double).
//							For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter
//							in currencies.json, it shows the number of digits past the decimal point for
//							each currency (2 for the majority of currencies).
type Invoice struct {
	Title          string `json:"user"`
	Description    string `json:"description"`
	StartParameter string `json:"start_parameter"`
	Currency       string `json:"currency"`
	TotalAmount    int    `json:"total_amount"`
}

// ShippingAddress - This object represents a shipping address.
//
// country_code	String	ISO 3166-1 alpha-2 country code
// state	String	State, if applicable
// city	String	City
// street_line1	String	First line for the address
// street_line2	String	Second line for the address
// post_code	String	Address post code
type ShippingAddress struct {
	CountryCode string `json:"country_code"`
	State       string `json:"state"`
	City        string `json:"city"`
	StreetLine1 string `json:"street_line1"`
	StreetLine2 string `json:"street_line2"`
	PostCode    string `json:"post_code"`
}

// OrderInfo - This object represents information about an order.
//
// name				String			Optional. User name
// phone_number		String			Optional. User's phone number
// email			String			Optional. User email
// shipping_address	ShippingAddress	Optional. User shipping address
type OrderInfo struct {
	Name            string           `json:"name,omitempty"`
	PhoneNumber     string           `json:"phone_number,omitempty"`
	Email           string           `json:"email,omitempty"`
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
}

// SuccessfulPayment - This object contains basic information about a successful payment.
//
// currency						String	Three-letter ISO 4217 currency code
// total_amount					Integer	Total price in the smallest units of the currency (integer, not float/double).
//										For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter
//										in currencies.json, it shows the number of digits past the decimal point for
//										each currency (2 for the majority of currencies).
// invoice_payload				String	Bot specified invoice payload
// shipping_option_id			String	Optional. Identifier of the shipping option chosen by the user
// order_info				 OrderInfo	Optional. Order info provided by the user
// telegram_payment_charge_id	String	Telegram payment identifier
// provider_payment_charge_id	String	Provider payment identifier
type SuccessfulPayment struct {
	Currency                string     `json:"currency"`
	TotalAmount             int        `json:"total_amount"`
	InvoicePayload          string     `json:"invoice_payload"`
	ShippingOptionID        string     `json:"shipping_option_id,omitempty"`
	OrderInfo               *OrderInfo `json:"order_info,omitempty"`
	TelegramPaymentChargeID string     `json:"telegram_payment_charge_id,omitempty"`
	ProviderPaymentChargeID string     `json:"provider_payment_charge_id,omitempty"`
}

// ShippingQuery - This object contains information about an incoming shipping query.
//
// id				String			Unique query identifier
// from				User			User who sent the query
// invoice_payload	String			Bot specified invoice payload
// shipping_address	ShippingAddress	User specified shipping address
type ShippingQuery struct {
	ID              string           `json:"id"`
	From            *User            `json:"from"`
	InvoicePayload  string           `json:"invoice_payload"`
	ShippingAddress *ShippingAddress `json:"shipping_address"`
}

// PreCheckoutQuery - This object contains information about an incoming pre-checkout query.
//
// id					String		Unique query identifier
// from					User		User who sent the query
// currency				String		Three-letter ISO 4217 currency code
// total_amount			Integer		Total price in the smallest units of the currency (integer, not float/double).
//									For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter
//									in currencies.json, it shows the number of digits past the decimal point for
//									each currency (2 for the majority of currencies).
// invoice_payload		String		Bot specified invoice payload
// shipping_option_id	String		Optional. Identifier of the shipping option chosen by the user
// order_info			OrderInfo	Optional. Order info provided by the user
type PreCheckoutQuery struct {
	ID               string     `json:"id"`
	From             *User      `json:"from"`
	Currency         string     `json:"currency"`
	TotalAmount      int        `json:"total_amount"`
	InvoicePayload   string     `json:"invoice_payload"`
	ShippingOptionID string     `json:"shipping_option_id,omitempty"`
	OrderInfo        *OrderInfo `json:"order_info,omitempty"`
}
