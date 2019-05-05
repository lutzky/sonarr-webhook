package mail

import (
	"testing"
)

func TestBuildRawMessage(t *testing.T) {
	testCases := []struct {
		name                   string
		from, to, subject, msg string
		want                   string
	}{
		{
			"simple",
			"src@src.net", "dest@dest.net", "Hello sir", "How are you?",
			"From: src@src.net\r\nTo: dest@dest.net\r\nSubject: Hello sir\r\n\r\nHow are you?",
		},
		{
			"LF-newline",
			"src@src.net", "dest@dest.net", "Hello sir", "How are you?\nGood day",
			"From: src@src.net\r\nTo: dest@dest.net\r\nSubject: Hello sir\r\n\r\nHow are you?\r\nGood day",
		},
		{
			"CRLF-newline",
			"src@src.net", "dest@dest.net", "Hello sir", "How are you?\r\nGood day",
			"From: src@src.net\r\nTo: dest@dest.net\r\nSubject: Hello sir\r\n\r\nHow are you?\r\nGood day",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := buildRawMessage(tc.from, tc.to, tc.subject, tc.msg)
			if got != tc.want {
				t.Errorf("buildRawMessage mismatch.\nGot:  %q\nWant: %q", got, tc.want)
			}
		})
	}
}
