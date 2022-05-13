/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log into you mogenius account",
	Long:  `Before you can interact with the cli you need to log into mogenius.`,
	Run: func(cmd *cobra.Command, args []string) {
		client := &http.Client{
			Timeout: time.Second * 10,
		}

		data := url.Values{}
		data.Set("email", "bene@mogenius.com")
		data.Set("password", "mye!ptw!DMT-uvy8eyc")
		data.Add("grant_type", "password")
		data.Add("device", "mogenius-cli")
		encodedData := data.Encode()
		fmt.Println(encodedData)
		req, err := http.NewRequest("POST", "https://api.dev.mogenius.com/auth/login", strings.NewReader(encodedData))
		if err != nil {
			log.Fatalln(err)
		}
		deviceStr := fmt.Sprintf("mo-cli v=%s arch=%s go=%s", runtime.GOARCH, runtime.GOARCH, runtime.Version())
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
		req.Header.Add("x-device", deviceStr)
		req.Header.Add("Authorization", "Basic RFItTkY0TFI1clkyd0V3TlFZV1BfOkVXTE9lY0lmT0U1aFFOQV9hM04xRg==")
		response, err := client.Do(req)

		// defer response.Body.Close()

		//We Read the response body on the line below.
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatalln(err)
		}
		//Convert the body to type string
		sb := string(body)
		manifestJson, _ := json.MarshalIndent(sb, "", "  ")
		fmt.Println(string(manifestJson))
		log.Printf(sb)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
