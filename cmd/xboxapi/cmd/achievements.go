package cmd

import (
	"encoding/json"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// achievementsCmd represents the achievements command
var achievementsCmd = &cobra.Command{
	Use:   "achievements",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			log.Fatal("please specify an xuid and titleID")
		}
		titleID, err := strconv.Atoi(args[1])
		if err != nil {
			log.Fatal(err)
		}
		rsp, err := client().Achievements(args[0], titleID)
		if err != nil {
			log.Fatal(err)
		}
		buf, _ := json.MarshalIndent(rsp, "", "  ")
		os.Stdout.Write(buf)
		os.Stdout.Write([]byte("\n"))
	},
}

func init() {
	RootCmd.AddCommand(achievementsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// achievementsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// achievementsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
