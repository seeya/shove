package email

import (
	"fmt"
	//"net/smtp"
	"time"
	"crypto/tls"
	"gitlab.com/pennersr/shove/internal/services"
	jwemail "github.com/jordan-wright/email"
)

/*
func (ec EmailConfig) send(from string, to []string, body []byte, fc services.FeedbackCollector) error {
	t := time.Now()
	err := smtp.SendMail(fmt.Sprintf("%s:%d", ec.EmailHost, ec.EmailPort), nil, from, to, body)
	duration := time.Since(t)
	fc.CountPush(serviceID, err == nil, duration)

	if err != nil {
		ec.Log.Printf("[ERROR] Send failed: %s", err)
		return err
	}
	return nil
}
*/

func (ec EmailConfig) send(from string, to []string, body []byte, fc services.FeedbackCollector) error {
        t := time.Now()

        email, _ := jwemail.NewEmailFromReader(bytes.NewReader(body))
        err := email.SendWithStartTLS(fmt.Sprintf("%s:%d", ec.EmailHost, ec.EmailPort), nil, &tls.Config{InsecureSkipVerify: true})
        duration := time.Since(t)
        fc.CountPush(serviceID, err == nil, duration)

        if err != nil {
                ec.Log.Printf("[ERROR] Send failed: %s", err)
                return err
        }
        return nil
}
