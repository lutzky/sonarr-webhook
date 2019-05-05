// package mail is a convenient wrapper around net/smtp, automatically
// populating SMTP headers.

package mail

import (
	"net/smtp"
	"strings"
)

// Send sends an email
func Send(server string, a smtp.Auth, from, to, subject, msg string) error {
	rawMsg := buildRawMessage(from, to, subject, msg)
	return smtp.SendMail(server, a, from, []string{to}, []byte(rawMsg))
}

func buildRawMessage(from, to, subject, msg string) string {
	msgCRLF := strings.NewReplacer("\r\n", "\r\n", "\n", "\r\n").Replace(msg)
	return "From: " + from + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" + msgCRLF
}
