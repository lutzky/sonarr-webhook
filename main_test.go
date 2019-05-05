package main

import "testing"

func TestSubjectAndMessage(t *testing.T) {
	testCases := []struct {
		name                     string
		s                        string
		wantSubject, wantMessage string
	}{
		{
			"oneline",
			"just one line",
			"", "just one line",
		},
		{
			"twoline",
			"line1\nline2",
			"", "line1\nline2",
		},
		{
			"one-and-one",
			"subject\n\nmessage",
			"subject", "message",
		},
		{
			"one-and-two",
			"subject\n\nmessage1\nmessage2",
			"subject", "message1\nmessage2",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotSubject, gotMessage := subjectAndMessage(tc.s)
			if gotSubject != tc.wantSubject {
				t.Errorf("Subject mismatch\nGot:  %q; Want: %q", gotSubject, tc.wantSubject)
			}
			if gotMessage != tc.wantMessage {
				t.Errorf("Message mismatch\nGot:  %q; Want: %q", gotMessage, tc.wantMessage)
			}
		})
	}
}
