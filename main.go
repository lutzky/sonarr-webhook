package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/smtp"
	"os"
	"strings"
	"text/template"

	"github.com/lutzky/sonarr-webhook/mail"
	"github.com/lutzky/sonarr-webhook/pkg/sonarr"
)

var (
	port         = flag.Int("port", 9999, "Port to listen on")
	templateFile = flag.String("template", "template.txt", "Template file")
	configFile   = flag.String("config", "config.json", "Config file")
)

var tmpl *template.Template

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

type Config struct {
	Mail struct {
		From     string
		Server   string
		Username string
		Password string
	}
}

func loadConfig() Config {
	var c Config
	f, err := os.Open(*configFile)
	if err != nil {
		log.Fatal(err)
	}
	if err := json.NewDecoder(f).Decode(&c); err != nil {
		log.Fatal(err)
	}
	return c
}

func runTemplate(event sonarr.SonarrEvent) (string, error) {
	var b bytes.Buffer

	if err := tmpl.Execute(&b, struct {
		SonarrEvent sonarr.SonarrEvent
	}{
		SonarrEvent: event,
	}); err != nil {
		return "", err
	}

	return b.String(), nil
}

func main() {
	flag.Parse()
	tmpl = template.Must(template.ParseFiles(*templateFile))

	config := loadConfig()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var ev sonarr.SonarrEvent
		defer r.Body.Close()
		if err := json.NewDecoder(r.Body).Decode(&ev); err != nil {
			http.Error(w, fmt.Sprintf("Invalid SonarrEvent: %v", err), http.StatusBadRequest)
			return
		}

		to := r.FormValue("to")

		if to == "" {
			http.Error(w, "Missing ?to=target@email.com", http.StatusBadRequest)
			return
		}

		response, err := runTemplate(ev)

		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Printf("Error executing template: %v", err)
			return
		}

		subject, message := subjectAndMessage(response)
		host, _, _ := net.SplitHostPort(config.Mail.Server)
		auth := smtp.PlainAuth("", config.Mail.Username, config.Mail.Password, host)
		if err := mail.Send(config.Mail.Server, auth, config.Mail.From, to, subject, message); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Printf("Error sending mail: %v", err)
			return
		}
	})
	fmt.Println("Listening on port", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
