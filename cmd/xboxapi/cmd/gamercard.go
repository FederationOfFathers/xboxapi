package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

// gamercardCmd represents the gamercard command
var gamercardCmd = &cobra.Command{
	Use:   "gamercard",
	Short: "Return the gamer card for a specific xuid",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("please specify an xuid")
		}
		xuid, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}
		gc, err := client().GamerCard(xuid)
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
	RootCmd.AddCommand(gamercardCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gamercardCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gamercardCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
