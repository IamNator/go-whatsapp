package webhook

import (
	"encoding/json"
	"io"
)

// HandleWebhook parses WhatsApp webhook requests
//
// Example usage with standard http:
//
//	func WebhookHandler(w http.ResponseWriter, r *http.Request) {
//	    if r.Method != http.MethodPost {
//	        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
//	        return
//	    }
//
//	    webhook, err := Parse(r.Body)
//	    if err != nil {
//	        http.Error(w, err.Error(), http.StatusBadRequest)
//	        return
//	    }
//
//	    // Handle messages
//	    for _, entry := range webhook.Entry {
//	        for _, change := range entry.Changes {
//	            ...
//	        }
//	    }
//
//	    w.WriteHeader(http.StatusOK)
//	}
//
// Example usage with Fiber:
//
//	app.Post("/webhook", func(c *fiber.Ctx) error {
//	    webhook, err := Parse(bytes.NewReader(c.Body()))
//	    if err != nil {
//	        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
//	            "error": err.Error(),
//	        })
//	    }
//
//	    // Handle messages
//	    for _, entry := range webhook.Entry {
//	        ...
//	    }
//
//	    return c.SendStatus(fiber.StatusOK)
//	})
//
// Example usage with Gin:
//
//	r.POST("/webhook", func(c *gin.Context) {
//	    webhook, err := Parse(c.Request.Body)
//	    if err != nil {
//	        c.JSON(http.StatusBadRequest, gin.H{
//	            "error": err.Error(),
//	        })
//	        return
//	    }
//
//	    // Handle messages
//	    for _, entry := range webhook.Entry {
//	        ...
//	    }
//
//	    c.Status(http.StatusOK)
//	})
func Parse(reader io.Reader) (*WebhookRequest, error) {
	var webhook WebhookRequest
	if err := json.NewDecoder(reader).Decode(&webhook); err != nil {
		return nil, err
	}

	return &webhook, nil
}

func ParseBytes(data []byte) (*WebhookRequest, error) {
	var webhook WebhookRequest
	if err := json.Unmarshal(data, &webhook); err != nil {
		return nil, err
	}

	return &webhook, nil
}

func ParseAny(input interface{}) (*WebhookRequest, error) {
	// Handle different input types
	switch v := input.(type) {
	case io.Reader:
		return Parse(v)
	case []byte:
		return ParseBytes(v)
	default:
		return nil, &json.UnmarshalTypeError{} // Return an error for unsupported types
	}
}
