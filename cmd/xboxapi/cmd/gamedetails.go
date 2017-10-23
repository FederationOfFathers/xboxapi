package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// gamedetailsCmd represents the gamedetails command
var gamedetailsCmd = &cobra.Command{
	Use:   "gamedetails",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("please specify an xuid")
		}
		data, err := client().GameDetails(args[0])
		if err != nil {
			log.Fatal(err)
		}
		buf, err := json.MarshalIndent(data, "", "\t")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(buf))
	},
}

func init() {
	RootCmd.AddCommand(gamedetailsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gamedetailsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gamedetailsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
