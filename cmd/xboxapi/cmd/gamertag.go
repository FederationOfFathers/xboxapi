package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

// gamertagCmd represents the gamertag command
var gamertagCmd = &cobra.Command{
	Use:   "gamertag",
	Short: "Return the GamerTag for an XUID integer",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("please specify an xuid")
		}
		xuid, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}
		gt, err := client().GamerTag(xuid)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(gt)
	},
}

func init() {
	RootCmd.AddCommand(gamertagCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gamertagCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gamertagCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
