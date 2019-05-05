package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/lutzky/sonarr-webhook/pkg/sonarr"
)

var (
	port = flag.Int("port", 9999, "Port to listen on")
)

var tmpl = template.Must(template.ParseFiles("template.txt"))

func subjectAndMessage(s string) (string, string) {
	scanner := bufio.NewScanner(strings.NewReader(s))

	var b1, b2 bytes.Buffer
	seenBlankLine := false

	for scanner.Scan() {
		line := scanner.Text()
		if !seenBlankLine && line == "" {
			seenBlankLine = true
			continue
		}
		if !seenBlankLine {
			fmt.Fprintln(&b1, line)
		} else {
			fmt.Fprintln(&b2, line)
		}
	}

	if b2.Len() == 0 {
		// No subject, just a message
		return "", strings.TrimSpace(b1.String())
	}
	return strings.TrimSpace(b1.String()), strings.TrimSpace(b2.String())
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var s sonarr.SonarrEvent
		defer r.Body.Close()
		json.NewDecoder(r.Body).Decode(&s)

		if err := tmpl.Execute(os.Stdout, struct {
			SonarrEvent sonarr.SonarrEvent
		}{
			SonarrEvent: s,
		}); err != nil {
			fmt.Printf("Failed to execute template: %v\n", err)
		}
	})
	fmt.Println("Listening on port", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
