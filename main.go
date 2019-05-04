package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/lutzky/sonarr-webhook/pkg/sonarr"
)

var (
	port = flag.Int("port", 9999, "Port to listen on")
)

var tmpl = template.Must(template.ParseFiles("template.txt"))

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
