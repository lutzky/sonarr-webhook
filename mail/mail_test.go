package mail

import (
	"testing"
)

func TestBuildRawMessage(t *testing.T) {
	testCases := []struct {
		name             string
		to, subject, msg string
		want             string
	}{
		{
			"simple",
			"dest@dest.net", "Hello sir", "How are you?",
			"To: dest@dest.net\r\nSubject: Hello sir\r\n\r\nHow are you?",
		},
		{
			"LF-newline",
			"dest@dest.net", "Hello sir", "How are you?\nGood day",
			"To: dest@dest.net\r\nSubject: Hello sir\r\n\r\nHow are you?\r\nGood day",
		},
		{
			"CRLF-newline",
			"dest@dest.net", "Hello sir", "How are you?\r\nGood day",
			"To: dest@dest.net\r\nSubject: Hello sir\r\n\r\nHow are you?\r\nGood day",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := buildRawMessage(tc.to, tc.subject, tc.msg)
			if got != tc.want {
				t.Errorf("buildRawMessage mismatch.\nGot:  %q\nWant: %q", got, tc.want)
			}
		})
	}
}
