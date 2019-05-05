package main

import (
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

	got, err := runTemplate(s)
	if err != nil {
		t.Fatal(err)
	}

	if *update {
		ioutil.WriteFile("testdata/want.txt", []byte(got), 0644)
	}

	want, err := ioutil.ReadFile("testdata/want.txt")
	if err != nil {
		t.Fatal(err)
	}

	if string(want) != got {
		t.Error("testdata/want.txt doesn't match; run with -update and use git diff")
	}
}
