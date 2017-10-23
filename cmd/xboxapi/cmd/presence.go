package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

// presenceCmd represents the presence command
var presenceCmd = &cobra.Command{
	Use:   "presence",
	Short: "Get presence information from the specified xuid",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("please specify an xuid")
		}
		xuid, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}
		gc, err := client().Presence(xuid)
		if err != nil {
			log.Fatal(err)
		}
		buf, err := json.MarshalIndent(gc, "", "\t")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(buf))
	},
}

func init() {
	RootCmd.AddCommand(presenceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// presenceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// presenceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
