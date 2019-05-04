// Package sonarr implements the sonarr webhook schema.
//
// This is grabbed from https://github.com/Sonarr/Sonarr/wiki/Webhook-Schema,
// and converted to Golang by using https://app.quicktype.io.

package sonarr

import "encoding/json"

func UnmarshalSonarrEvent(data []byte) (SonarrEvent, error) {
	var r SonarrEvent
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SonarrEvent) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Sonarr Webhook Event
type SonarrEvent struct {
	EpisodeFile *EpisodeFile `json:"episodeFile,omitempty"`
	Episodes    []Episode    `json:"episodes"`
	EventType   EventType    `json:"eventType"`
	IsUpgrade   *bool        `json:"isUpgrade,omitempty"`
	Release     *Release     `json:"release,omitempty"`
	Series      Series       `json:"series"`
}

type EpisodeFile struct {
	ID             int64   `json:"id"`
	Path           string  `json:"path"`
	Quality        *string `json:"quality,omitempty"`
	QualityVersion *int64  `json:"qualityVersion,omitempty"`
	RelativePath   string  `json:"relativePath"`
	ReleaseGroup   *string `json:"releaseGroup,omitempty"`
	SceneName      *string `json:"sceneName,omitempty"`
}

type Episode struct {
	AirDate        *string `json:"airDate,omitempty"`
	AirDateUTC     *string `json:"airDateUtc,omitempty"`
	EpisodeNumber  int64   `json:"episodeNumber"`
	ID             int64   `json:"id"`
	Quality        *string `json:"quality,omitempty"`        // Deprecated: will be removed in a future version
	QualityVersion *int64  `json:"qualityVersion,omitempty"` // Deprecated: will be removed in a future version
	ReleaseGroup   *string `json:"releaseGroup,omitempty"`   // Deprecated: will be removed in a future version
	SceneName      *string `json:"sceneName,omitempty"`      // Deprecated: will be removed in a future version
	SeasonNumber   int64   `json:"seasonNumber"`
	Title          string  `json:"title"`
}

type Release struct {
	Indexer        *string `json:"indexer,omitempty"`
	Quality        *string `json:"quality,omitempty"`
	QualityVersion *int64  `json:"qualityVersion,omitempty"`
	ReleaseGroup   *string `json:"releaseGroup,omitempty"`
	ReleaseTitle   *string `json:"releaseTitle,omitempty"`
	Size           *int64  `json:"size,omitempty"`
}

type Series struct {
	ID     int64  `json:"id"`
	Path   string `json:"path"`
	Title  string `json:"title"`
	TvdbID *int64 `json:"tvdbId,omitempty"`
}

type EventType string

const (
	Download EventType = "Download"
	Grab     EventType = "Grab"
	Rename   EventType = "Rename"
	Test     EventType = "Test"
)
