package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	"testing"
	"text/template"

	"github.com/lutzky/sonarr-webhook/pkg/sonarr"
)

var update = flag.Bool("update", false, "Update golden test files")

func TestTemplate(t *testing.T) {
	tmpl = template.Must(template.ParseFiles("template.txt"))
	payload, err := os.Open("testdata/payload.json")
	if err != nil {
		t.Fatal(err)
	}
	defer payload.Close()

	var s sonarr.SonarrEvent

	json.NewDecoder(payload).Decode(&s)

	var b bytes.Buffer

	if err := tmpl.Execute(&b, struct {
		SonarrEvent sonarr.SonarrEvent
	}{
		SonarrEvent: s,
	}); err != nil {
		t.Fatal(err)
	}

	got := b.Bytes()

	if *update {
		ioutil.WriteFile("testdata/want.txt", got, 0644)
	}

	want, err := ioutil.ReadFile("testdata/want.txt")
	if err != nil {
		t.Fatal(err)
	}

	if string(want) != string(got) {
		t.Error("testdata/want.txt doesn't match; run with -update and use git diff")
	}
}
