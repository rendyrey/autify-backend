package main

import (
	"fmt"
	"os"

	fwu "github.com/rendyrey/autify-backend/fetch_web_utilities"
)

func main() {
	urls := make([]string, len(os.Args))
	copy(urls, os.Args)

	if len(urls) == 1 {
		fmt.Println("Usage: go run app.go [URLs...]")
	}

	for _, url_ := range urls[1:] {
		var website fwu.WebUrl
		siteUrl := website.New(url_)

		fmt.Println(siteUrl.Url)

		_, err := siteUrl.FetchWebPage()

		if err != nil {
			fmt.Println(err)
			return
		}

		// parse HTML to extract asset URLs and download them
		metadata, err := siteUrl.FetchWebAssets()
		if err != nil {
			fmt.Println("Error parsing HTML:", err)
			return
		}

		// Replace asset URLs in HTML file with local paths
		err = siteUrl.ReplaceAssetURLsInHTML()
		if err != nil {
			fmt.Println("Error replacing asset URLs:", err)
			return
		}
		fmt.Println("HTML file updated with local asset paths")

		metadata.Display()
		metadata.Save()
	}

}
