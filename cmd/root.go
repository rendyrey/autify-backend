/*
Copyright © 2024 Rendy Reynaldy <rendyreynaldy@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	fwu "github.com/rendyrey/rendy_web_fetcher/fetch_web_utilities"
	"github.com/spf13/cobra"
)

var (
	metadataFlag            bool
	metadataDisplayOnlyFlag bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rendy_web_fetcher [flags] [URLs...]",
	Short: "Web Fetcher by Rendy Reynaldy",
	Long: `Web Fetcher by Rendy
A command-line tool for retrieving web content. 
Fetch HTML, CSS, Images, and JavaScript resources effortlessly for offline browsing or analysis.`,
	// Run: func(cmd *cobra.Command, args []string) {},
	Run: fetchTheWeb,
}

func fetchTheWeb(cmd *cobra.Command, args []string) {
	params := removeFlagsInParams(os.Args[1:])
	urls := make([]string, len(params))
	copy(urls, params)

	if len(urls) == 0 {
		fmt.Println(cmd.Usage())
	}

	for _, url_ := range urls {
		var website fwu.WebUrl
		siteUrl := website.New(url_)

		fmt.Println("Fetching", siteUrl.Url, "...")

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

		if metadataDisplayOnlyFlag {
			metadata.Display()
		}

		if metadataFlag {
			metadata.Display()
			metadata.Save()
		}
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	rootCmd.Flags().BoolVarP(&metadataFlag, "metadata", "m", false, "Display metadata in console and save it in json format")
	rootCmd.Flags().BoolVarP(&metadataDisplayOnlyFlag, "metadata-display", "d", false, "Only display metadata in console")
	rootCmd.MarkFlagsMutuallyExclusive("metadata", "metadata-display")
}
