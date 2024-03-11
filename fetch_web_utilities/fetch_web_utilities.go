package fetch_web_utilities

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"slices"
	"strings"
	"time"

	"golang.org/x/net/html"
)

// it's supposed to be constant, but since go not support slice as constant
// please don't mutate this
var AnticipatedAssetTag = []string{"img", "link", "script", "source"}
var AnticipatedAssetProp = []string{"src", "href"}

type WebUrl struct {
	Url          string
	Hostname     string
	PageFilename string
	AssetDir     string
}

func (w WebUrl) New(u string) WebUrl {
	webUrl := sanitizeUrl(u)
	link, _ := url.Parse(webUrl)
	hostname := strings.TrimPrefix(link.Hostname(), "www.")

	return WebUrl{
		Url:          webUrl,
		Hostname:     hostname,
		PageFilename: hostname + ".html",
		AssetDir:     hostname + "_assets",
	}
}

func (w WebUrl) FetchWebPage() ([]byte, error) {
	// fetch the web page
	resp, err := http.Get(w.Url)
	if err != nil {
		errMsg := fmt.Sprintf("Error fetching the web page: %v", err)
		return nil, errors.New(errMsg)
	}
	defer resp.Body.Close()

	// read the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		errMsg := fmt.Sprintf("Error reading response body: %v", err)
		return nil, errors.New(errMsg)
	}

	// save webpage to disk
	err = os.WriteFile(w.PageFilename, body, 0755)
	if err != nil {
		errMsg := fmt.Sprintf("Error saving webpage as html to disk: %v", err)
		return nil, errors.New(errMsg)
	}

	return body, nil
}

func (w WebUrl) FetchWebAssets() (Metadata, error) {
	// create a directory to store assets
	var metadata Metadata
	site := w.Hostname
	numLinks := 0
	images := 0
	lastFetch := time.Now()
	err := os.Mkdir(w.AssetDir, 0755)
	if err != nil && !os.IsExist(err) {
		errMsg := fmt.Sprintf("Error creating directory for assets "+w.Hostname+" : %v", err)
		return Metadata{}, errors.New(errMsg)
	}

	f, err := os.Open(w.PageFilename)
	if err != nil {
		return Metadata{}, err
	}

	defer f.Close()

	tokenizer := html.NewTokenizer(f)

	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			if tokenizer.Err() == io.EOF {
				break
			}
			return Metadata{}, tokenizer.Err()
		}

		if tokenType == html.StartTagToken {
			token := tokenizer.Token()
			if slices.Contains(AnticipatedAssetTag, token.Data) {
				for _, attr := range token.Attr {
					if slices.Contains(AnticipatedAssetProp, attr.Key) {
						assetURL := attr.Val
						if !strings.HasPrefix(assetURL, "http://") && !strings.HasPrefix(assetURL, "https://") {
							assetURL, _ = url.JoinPath(w.Url, assetURL)
						}
						downloadAsset(assetURL, w.AssetDir)
					}
				}
			}

			if token.Data == "a" {
				numLinks += 1
			} else if token.Data == "img" {
				images += 1
			}
		}
	}

	metadata = metadata.New(site, numLinks, images, lastFetch)

	return metadata, nil
}

func (w WebUrl) ReplaceAssetURLsInHTML() error {
	data, err := os.ReadFile(w.PageFilename)
	if err != nil {
		return err
	}

	doc, err := html.Parse(strings.NewReader(string(data)))
	if err != nil {
		return err
	}

	var replaceURL func(*html.Node)
	replaceURL = func(n *html.Node) {
		if n.Type == html.ElementNode && slices.Contains(AnticipatedAssetTag, n.Data) {
			for i, attr := range n.Attr {
				if slices.Contains(AnticipatedAssetProp, attr.Key) {
					localURL := path.Join(w.AssetDir, path.Base(attr.Val))
					n.Attr[i].Val = localURL
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			replaceURL(c)
		}
	}

	replaceURL(doc)

	var buf strings.Builder
	html.Render(&buf, doc)

	err = os.WriteFile(w.PageFilename, []byte(buf.String()), 0644)
	if err != nil {
		return err
	}

	return nil
}

func sanitizeUrl(url string) string {
	// Check if the URL starts with "http://" or "https://", and add if not present
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}

	return url
}

func downloadAsset(url, dir string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	assetFileName := path.Join(dir, path.Base(url))
	err = os.WriteFile(assetFileName, body, 0644)
	if err != nil {
		return err
	}

	return nil
}
