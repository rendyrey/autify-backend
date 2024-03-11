package fetch_web_utilities

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type Metadata struct {
	Site      string    `json:"site"`
	NumLinks  int       `json:"num_links"`
	Images    int       `json:"images"`
	LastFetch time.Time `json:"last_fetch"`
}

func (metadata Metadata) New(site string, numLinks, images int, lastFetch time.Time) Metadata {
	return Metadata{
		Site:      site,
		NumLinks:  numLinks,
		Images:    images,
		LastFetch: lastFetch,
	}
}

func (metadata Metadata) Save() error {
	fileName := strings.ToLower(metadata.Site) + "_metadata.json"
	json, err := json.Marshal(metadata)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func (metadata Metadata) Display() {
	fmt.Println("site:", metadata.Site)
	fmt.Println("num_links:", metadata.NumLinks)
	fmt.Println("images", metadata.Images)
	fmt.Println("last_fetch", metadata.LastFetch)
	fmt.Printf("\n")
}
