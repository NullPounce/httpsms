package emails

import (
	"github.com/NdoleStudio/httpsms/pkg/entities"
	"github.com/google/uuid"
)

// NotificationEmailFactory generates emails to users about a message
type NotificationEmailFactory interface {
	// MessageExpired sends an email when the user's message is expired
	MessageExpired(user *entities.User, messageID uuid.UUID, owner, contact, content string) (*Email, error)

	// MessageFailed sends an email when the user's message is failed
	MessageFailed(user *entities.User, messageID uuid.UUID, owner, contact, content, reason string) (*Email, error)

	// DiscordMessageFailed sends an email when the user's discord message is failed
	DiscordMessageFailed(user *entities.User, eventName, owner, errorMessage, channelID string, responseHTTPStatusCode *int) (*Email, error)

	// WebhookSendFailed sends an email when the user's webhook message is failed
	WebhookSendFailed(user *entities.User, eventName, eventID, owner, errorMessage, url string, responseHTTPStatusCode *int) (*Email, error)
}
