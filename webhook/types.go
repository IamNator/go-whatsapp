package webhook

// WebhookRequest represents the main webhook payload from WhatsApp Cloud API
type WebhookRequest struct {
	Object string  `json:"object"` // whatsapp_business_account
	Entry  []Entry `json:"entry"`
}

// Entry represents a webhook event from WhatsApp Business API
type Entry struct {
	ID      string   `json:"id"`      // WhatsApp Business Account ID
	Changes []Change `json:"changes"` // Array of changes in this webhook event
}

// Change contains the actual webhook payload data
type Change struct {
	Value Value  `json:"value"` // Contains messages, contacts, and status data
	Field string `json:"field"` // Usually "messages"
}

// Value holds the main webhook data including messages and metadata
type Value struct {
	MessagingProduct string    `json:"messaging_product"` // Always "whatsapp"
	Metadata         Metadata  `json:"metadata"`          // Business phone info
	Contacts         []Contact `json:"contacts,omitempty"`
	Messages         []Message `json:"messages,omitempty"`
	Statuses         []Status  `json:"statuses,omitempty"`
}

// Metadata contains WhatsApp Business phone number information
type Metadata struct {
	DisplayPhoneNumber string `json:"display_phone_number"`
	PhoneNumberID      string `json:"phone_number_id"`
}

// Contact contains customer/user profile information
type Contact struct {
	Profile Profile `json:"profile"`
	WaID    string  `json:"wa_id"` // WhatsApp ID (usually phone number)
}

// Profile contains user display information
type Profile struct {
	Name string `json:"name"`
}

// Message represents any type of WhatsApp message
type Message struct {
	From        string           `json:"from"`
	ID          string           `json:"id"`        // wamid.ID format
	Timestamp   string           `json:"timestamp"` // Unix timestamp
	Text        *Text            `json:"text,omitempty"`
	Type        string           `json:"type"` // text, image, video, etc.
	Image       *Media           `json:"image,omitempty"`
	Video       *Media           `json:"video,omitempty"`
	Document    *Media           `json:"document,omitempty"`
	Location    *Location        `json:"location,omitempty"`
	Contacts    []ContactMessage `json:"contacts,omitempty"`
	Reaction    *Reaction        `json:"reaction,omitempty"`
	Context     *Context         `json:"context,omitempty"`
	Button      *Button          `json:"button,omitempty"`
	Interactive *Interactive     `json:"interactive,omitempty"`
	Referral    *Referral        `json:"referral,omitempty"`
	Order       *Order           `json:"order,omitempty"`
	System      *System          `json:"system,omitempty"`
	Errors      []Error          `json:"errors,omitempty"`
}

// Text represents a text message content
type Text struct {
	Body string `json:"body"` // Message text content
}

// Media represents image, video, or document message attachments
type Media struct {
	Caption  string `json:"caption,omitempty"` // Optional media caption
	MimeType string `json:"mime_type"`         // MIME type of media
	SHA256   string `json:"sha256"`            // Hash of media content
	ID       string `json:"id"`                // Media ID for downloading
}

// Location represents a shared location message
type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Name      string  `json:"name,omitempty"`    // Location name if available
	Address   string  `json:"address,omitempty"` // Address if available
}

// ContactMessage represents a shared contact message
type ContactMessage struct {
	Addresses []Address `json:"addresses,omitempty"`
	Birthday  string    `json:"birthday,omitempty"`
	Emails    []Email   `json:"emails,omitempty"`
	Name      Name      `json:"name"`
	Org       Org       `json:"org,omitempty"`
	Phones    []Phone   `json:"phones,omitempty"`
	URLs      []URL     `json:"urls,omitempty"`
}

// Address represents a contact's address information
type Address struct {
	City        string `json:"city"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
	State       string `json:"state"`
	Street      string `json:"street"`
	Type        string `json:"type"` // HOME or WORK
	Zip         string `json:"zip"`
}

// Email represents a contact's email information
type Email struct {
	Email string `json:"email"`
	Type  string `json:"type"` // WORK or HOME
}

// Name represents contact name details
type Name struct {
	FormattedName string `json:"formatted_name"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	MiddleName    string `json:"middle_name,omitempty"`
	Suffix        string `json:"suffix,omitempty"`
	Prefix        string `json:"prefix,omitempty"`
}

// Org represents business/organization information
type Org struct {
	Company    string `json:"company"`
	Department string `json:"department"`
	Title      string `json:"title"`
}

// Phone represents contact phone number information
type Phone struct {
	Phone string `json:"phone"`
	WaID  string `json:"wa_id,omitempty"` // WhatsApp ID if available
	Type  string `json:"type"`            // HOME or WORK
}

// URL represents contact URL information
type URL struct {
	URL  string `json:"url"`
	Type string `json:"type"` // HOME or WORK
}

// Reaction represents a message reaction
type Reaction struct {
	MessageID string `json:"message_id"` // ID of message being reacted to
	Emoji     string `json:"emoji"`      // Reaction emoji
}

// Context represents message context like replies
type Context struct {
	From            string           `json:"from"`
	ID              string           `json:"id"`
	ReferredProduct *ReferredProduct `json:"referred_product,omitempty"`
}

// ReferredProduct represents product catalog reference
type ReferredProduct struct {
	CatalogID         string `json:"catalog_id"`
	ProductRetailerID string `json:"product_retailer_id"`
}

// Button represents button interaction data
type Button struct {
	Text    string `json:"text"`    // Button text
	Payload string `json:"payload"` // Button payload
}

// Interactive represents interactive message responses
type Interactive struct {
	Type        string       `json:"type"` // button_reply or list_reply
	ButtonReply *ButtonReply `json:"button_reply,omitempty"`
	ListReply   *ListReply   `json:"list_reply,omitempty"`
}

// ButtonReply represents a button response
type ButtonReply struct {
	ID    string `json:"id"`    // Button identifier
	Title string `json:"title"` // Button text
}

// ListReply represents a list selection response
type ListReply struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// Referral represents Click to WhatsApp ad referral data
type Referral struct {
	SourceURL    string `json:"source_url"`
	SourceID     string `json:"source_id"`
	SourceType   string `json:"source_type"` // ad or post
	Headline     string `json:"headline"`    // Ad title
	Body         string `json:"body"`        // Ad description
	MediaType    string `json:"media_type"`  // image or video
	ImageURL     string `json:"image_url,omitempty"`
	VideoURL     string `json:"video_url,omitempty"`
	ThumbnailURL string `json:"thumbnail_url,omitempty"`
	CtwaClid     string `json:"ctwa_clid"`
}

// Order represents a product order message
type Order struct {
	CatalogID    string        `json:"catalog_id"`
	ProductItems []ProductItem `json:"product_items"`
	Text         string        `json:"text,omitempty"`
}

// ProductItem represents an ordered product
type ProductItem struct {
	ProductRetailerID string `json:"product_retailer_id"`
	Quantity          string `json:"quantity"`
	ItemPrice         string `json:"item_price"`
	Currency          string `json:"currency"`
}

// System represents system notifications
type System struct {
	Body    string `json:"body"`                // System message content
	NewWaID string `json:"new_wa_id,omitempty"` // New phone number for user
	Type    string `json:"type"`                // e.g., user_changed_number
}

// Error represents error information
type Error struct {
	Code      int       `json:"code"`
	Title     string    `json:"title"`
	Message   string    `json:"message,omitempty"`
	ErrorData ErrorData `json:"error_data,omitempty"`
	Details   string    `json:"details,omitempty"`
	Href      string    `json:"href,omitempty"` // e.g https://developers.facebook.com/docs/whatsapp/cloud-api/support/error-codes/
}

// ErrorData represents additional error data
type ErrorData struct {
	Details string `json:"details"`
}

type StatusValue string

const (
	StatusSent      StatusValue = "sent"
	StatusDelivered StatusValue = "delivered"
	StatusRead      StatusValue = "read"
	StatusFailed    StatusValue = "failed"
)

// Status represents message delivery status updates
type Status struct {
	ID           string        `json:"id"`     // Message ID
	Status       StatusValue   `json:"status"` // sent, delivered, read, failed
	Timestamp    string        `json:"timestamp"`
	RecipientID  string        `json:"recipient_id"`
	Conversation *Conversation `json:"conversation,omitempty"`
	Pricing      *Pricing      `json:"pricing,omitempty"`
	Errors       []Error       `json:"errors,omitempty"`
}

// Conversation represents conversation information
type Conversation struct {
	ID                  string `json:"id"`
	ExpirationTimestamp string `json:"expiration_timestamp,omitempty"`
	Origin              Origin `json:"origin"`
}

// Origin represents conversation origin information
type Origin struct {
	Type string `json:"type"` // e.g., user_initiated, business_initiated
}

// Pricing represents message pricing information
type Pricing struct {
	Billable     bool   `json:"billable"`      // Whether message is billable
	PricingModel string `json:"pricing_model"` // Usually "CBP"
	Category     string `json:"category"`      // Message category
}
