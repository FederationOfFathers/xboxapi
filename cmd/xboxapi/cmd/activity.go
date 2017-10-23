package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/FederationOfFathers/xboxapi"
	"github.com/spf13/cobra"
)

// activityCmd represents the activity command
var activityCmd = &cobra.Command{
	Use:   "activity",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("please specify an xuid")
		}
		xuid, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}
		var out []xboxapi.ActivityItems
		var token *json.Number
		var client = client()
		for {
			data, err := client.Activity(xuid, token)
			if data != nil {
				out = append(out, data.ActivityItems...)
			}
			if err != nil {
				log.Fatal(err)
			}
			if data == nil {
				break
			}
			if data.ContinuationToken == nil {
				break
			}
			token = data.ContinuationToken
		}
		buf, err := json.MarshalIndent(out, "", "\t")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(buf))
	},
}

func init() {
	RootCmd.AddCommand(activityCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// activityCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// activityCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
