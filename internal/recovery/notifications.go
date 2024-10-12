package recovery

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
	"os"
)

func sendEmail(email *string, code *string) (*string, error) {

	client := sendgrid.NewSendClient(os.Getenv("SG_KEY"))
	from := mail.NewEmail("Password Recovery", "caiodtn@gmail.com")
	to := mail.NewEmail("User", *email)

	message := mail.NewV3Mail()
	message.SetFrom(from)
	message.SetTemplateID(os.Getenv("SG_TEMPLATE_ID"))
	message.Subject = "Password Recovery Code"

	personalization := mail.NewPersonalization()
	personalization.AddTos(to)
	personalization.SetDynamicTemplateData("code", code)
	message.AddPersonalizations(personalization)

	response, err := client.Send(message)
	if err != nil {
		return nil, err
	}

	messageIDs := response.Headers["X-Message-Id"]
	if len(messageIDs) == 0 {
		return nil, fmt.Errorf("no message ID returned")
	}
	messageID := messageIDs[0]

	return &messageID, nil
}

func sendSMS(phone *string, code *string) (*string, error) {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: os.Getenv("TWILIO_SID"),
		Password: os.Getenv("TWILIO_TOKEN"),
	})

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(fmt.Sprintf("+55%s", *phone))
	params.SetFrom(os.Getenv("TWILIO_PHONE"))
	params.SetBody(fmt.Sprintf("Your password recovery code is: %s", *code))

	message, err := client.Api.CreateMessage(params)
	if err != nil {
		return nil, err
	}

	if message.Sid == nil {
		return nil, fmt.Errorf("no message SID returned")
	}

	return message.Sid, nil
}
