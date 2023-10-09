package emails

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/NdoleStudio/httpsms/pkg/entities"
	"github.com/matcornic/hermes/v2"
	"github.com/palantir/stacktrace"
)

type hermesNotificationEmailFactory struct {
	factory
	config    *HermesGeneratorConfig
	generator hermes.Hermes
}

// NewHermesNotificationEmailFactory creates a new instance of the UserEmailFactory
func NewHermesNotificationEmailFactory(config *HermesGeneratorConfig) NotificationEmailFactory {
	return &hermesNotificationEmailFactory{
		config:    config,
		generator: config.Generator(),
	}
}

func (factory *hermesNotificationEmailFactory) DiscordMessageFailed(user *entities.User, eventName, owner, errorMessage, channelID string, httpResponseStatusCode *int) (*Email, error) {
	email := hermes.Email{
		Body: hermes.Body{
			Title: "Hello",
			Intros: []string{
				fmt.Sprintf("We ran into an error while fowarding an incoming SMS to your discord server at %s", user.UserTimeString(time.Now())),
			},
			Dictionary: []hermes.Entry{
				{"Discord Channel ID", channelID},
				{"Event Name", eventName},
				{"Phone Number", factory.formatPhoneNumber(owner)},
				{"HTTP Response Code", factory.formatHTTPResponseCode(httpResponseStatusCode)},
				{"Error Message / HTTP Response", errorMessage},
			},
			Actions: []hermes.Action{
				{
					Instructions: "Usually this error happens because you have revoked permissions for the httpSMS discord app on your discord channel. You can always grant httpSMS permission to post to your discord channel under the settings page.",
					Button: hermes.Button{
						Color:     "#329ef4",
						TextColor: "#FFFFFF",
						Text:      "Settings",
						Link:      "https://httpsms.com/settings",
					},
				},
			},
			Signature: "Cheers",
			Outros: []string{
				fmt.Sprintf("Don't hesitate to contact us by replying to this email."),
			},
		},
	}

	html, err := factory.generator.GenerateHTML(email)
	if err != nil {
		return nil, stacktrace.Propagate(err, "cannot generate html email")
	}

	text, err := factory.generator.GeneratePlainText(email)
	if err != nil {
		return nil, stacktrace.Propagate(err, "cannot generate text email")
	}

	return &Email{
		ToEmail: user.Email,
		Subject: "📢 We could not forward an incoming message to your server",
		HTML:    html,
		Text:    text,
	}, nil
}

func (factory *hermesNotificationEmailFactory) WebhookSendFailed(user *entities.User, eventName, eventID, owner, errorMessage, url string, httpResponseStatusCode *int) (*Email, error) {
	email := hermes.Email{
		Body: hermes.Body{
			Title: "Hello",
			Intros: []string{
				fmt.Sprintf("We ran into an error while fowarding a webhook event from httpSMS to your webserver at %s", user.UserTimeString(time.Now())),
			},
			Dictionary: []hermes.Entry{
				{"Server URL", url},
				{"Event Name", eventName},
				{"Event ID", eventID},
				{"Phone Number", factory.formatPhoneNumber(owner)},
				{"HTTP Response Code", factory.formatHTTPResponseCode(httpResponseStatusCode)},
				{"Error Message / HTTP Response", errorMessage},
			},
			Actions: []hermes.Action{
				{
					Instructions: "Usually this error happens because your webserver is either offline or inaccessible, you can always configure the webhook endpoint on the httpSMS website under the settings page.",
					Button: hermes.Button{
						Color:     "#329ef4",
						TextColor: "#FFFFFF",
						Text:      "Settings",
						Link:      "https://httpsms.com/settings",
					},
				},
			},
			Signature: "Cheers",
			Outros: []string{
				fmt.Sprintf("Don't hesitate to contact us by replying to this email."),
			},
		},
	}

	html, err := factory.generator.GenerateHTML(email)
	if err != nil {
		return nil, stacktrace.Propagate(err, "cannot generate html email")
	}

	text, err := factory.generator.GeneratePlainText(email)
	if err != nil {
		return nil, stacktrace.Propagate(err, "cannot generate text email")
	}

	return &Email{
		ToEmail: user.Email,
		Subject: "📢 We could not forward a webhook event to your server",
		HTML:    html,
		Text:    text,
	}, nil
}

func (factory *hermesNotificationEmailFactory) MessageExpired(user *entities.User, messageID uuid.UUID, owner string, contact string, content string) (*Email, error) {
	email := hermes.Email{
		Body: hermes.Body{
			Title: "Hello",
			Intros: []string{
				fmt.Sprintf("The SMS message which you sent to %s has expired at %s and you will need to resend this message.", factory.formatPhoneNumber(contact), user.UserTimeString(time.Now())),
			},
			Dictionary: []hermes.Entry{
				{"ID", messageID.String()},
				{"From", factory.formatPhoneNumber(owner)},
				{"To", factory.formatPhoneNumber(contact)},
				{"Message", content},
			},
			Actions: []hermes.Action{
				{
					Instructions: "Messages usually expire because we couldn't connect with your mobile phone to send the outgoing SMS. You can fix this by making sure your phone is connected to the internet and also connect your phone to the charger all the time since Android may kill the httpSMS app if it has been active for a very long time so save phone battery.",
					Button: hermes.Button{
						Color:     "#329ef4",
						TextColor: "#FFFFFF",
						Text:      "View Messages",
						Link:      "https://httpsms.com/threads",
					},
				},
			},
			Signature: "Cheers",
			Outros: []string{
				fmt.Sprintf("Don't hesitate to contact us by replying to this email."),
			},
		},
	}

	html, err := factory.generator.GenerateHTML(email)
	if err != nil {
		return nil, stacktrace.Propagate(err, "cannot generate html email")
	}

	text, err := factory.generator.GeneratePlainText(email)
	if err != nil {
		return nil, stacktrace.Propagate(err, "cannot generate text email")
	}

	return &Email{
		ToEmail: user.Email,
		Subject: "📢 Your SMS message has expired on httpSMS",
		HTML:    html,
		Text:    text,
	}, nil
}

func (factory *hermesNotificationEmailFactory) MessageFailed(user *entities.User, messageID uuid.UUID, owner, contact, content, reason string) (*Email, error) {
	email := hermes.Email{
		Body: hermes.Body{
			Title: "Hello",
			Intros: []string{
				fmt.Sprintf("The SMS message which you sent to %s has failed at %s and you will need to resend this message.", factory.formatPhoneNumber(contact), user.UserTimeString(time.Now())),
			},
			Dictionary: []hermes.Entry{
				{"ID", messageID.String()},
				{"From", factory.formatPhoneNumber(owner)},
				{"To", factory.formatPhoneNumber(contact)},
				{"Message", content},
				{"Failure Reason", reason},
			},
			Actions: []hermes.Action{
				{
					Instructions: "Check the default SMS messaging app on your phone to find out the exact reason why the message failed. Usually messages fail because the httpSMS app phone has been un-installed or it is not active. Logout and login again on the mobile app on your Android phone and retry sending the SMS.",
					Button: hermes.Button{
						Color:     "#329ef4",
						TextColor: "#FFFFFF",
						Text:      "View Messages",
						Link:      "https://httpsms.com/threads",
					},
				},
			},
			Signature: "Cheers",
			Outros: []string{
				fmt.Sprintf("Don't hesitate to contact us by replying to this email."),
			},
		},
	}

	html, err := factory.generator.GenerateHTML(email)
	if err != nil {
		return nil, stacktrace.Propagate(err, "cannot generate html email")
	}

	text, err := factory.generator.GeneratePlainText(email)
	if err != nil {
		return nil, stacktrace.Propagate(err, "cannot generate text email")
	}

	return &Email{
		ToEmail: user.Email,
		Subject: "📢 Your SMS message has failed on httpSMS",
		HTML:    html,
		Text:    text,
	}, nil
}
