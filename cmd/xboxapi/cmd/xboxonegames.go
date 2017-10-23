package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/FederationOfFathers/xboxapi"
	"github.com/spf13/cobra"
)

// xboxonegamesCmd represents the xboxonegames command
var xboxonegamesCmd = &cobra.Command{
	Use:   "xboxonegames",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("please specify an xuid")
		}
		xuid, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}
		var out = []xboxapi.XboxOneTitle{}
		var token *json.Number
		for {
			data, err := client().XboxOneGames(xuid, token)
			if err != nil {
				log.Fatal(err)
			}
			out = append(out, data.Titles...)
			if data.PagingInfo.ContinuationToken == nil {
				break
			}
			token = data.PagingInfo.ContinuationToken
		}
		buf, err := json.MarshalIndent(out, "", "\t")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(buf))
	},
}

func init() {
	RootCmd.AddCommand(xboxonegamesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// xboxonegamesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// xboxonegamesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
