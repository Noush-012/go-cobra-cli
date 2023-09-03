/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package ping

import (
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

var (
	url    string
	client = http.Client{
		Timeout: 5 * time.Second,
	}
)

func ping(domain string) (int, error) {
	urlPath := "http://" + domain
	req, _ := http.NewRequest("HEAD", urlPath, nil)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	// bytResp, _ := io.ReadAll(resp.Body)

	defer resp.Body.Close()
	return resp.StatusCode, nil

}

// PingCmd represents the ping command
var PingCmd = &cobra.Command{
	Use:   "ping",
	Short: "This pings a remote url and returns a response",
	Long:  `To ping a remote url and returns a response example : www.google.com`,
	Run: func(cmd *cobra.Command, args []string) {
		// logic of PingCmd

		if resp, err := ping(url); err != nil {
			fmt.Println(err.Error())

		} else {
			fmt.Println(resp)
		}

	},
}

func init() {
	// rootCmd.AddCommand(pingCmd)
	// Here you will define your flags and configuration settings.
	PingCmd.Flags().StringVarP(&url, "url", "u", "", "The url path to ping")

	if err := PingCmd.MarkFlagRequired("url"); err != nil {
		fmt.Println(err)

	}
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
