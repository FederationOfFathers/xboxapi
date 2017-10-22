package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// xuidCmd represents the xuid command
var xuidCmd = &cobra.Command{
	Use:   "xuid",
	Short: "Return the XUID for a given GamerTag",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("Please specify a gamertag")
		}
		xuid, err := client().XUID(args[0])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(xuid)
	},
}

func init() {
	RootCmd.AddCommand(xuidCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// xuidCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// xuidCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
