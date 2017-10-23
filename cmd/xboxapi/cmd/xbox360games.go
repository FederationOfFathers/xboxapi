package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

// xbox360gamesCmd represents the xbox360games command
var xbox360gamesCmd = &cobra.Command{
	Use:   "xbox360games",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("please specify an xuid")
		}
		xuid, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}
		data, err := client().Xbox360Games(xuid)
		if err != nil {
			log.Fatal(err)
		}
		for {
			next, err := data.More()
			if next != nil {
				data.Titles = append(data.Titles, next.Titles...)
			} else {
				break
			}
			if err != nil {
				break
			}
		}
		buf, err := json.MarshalIndent(data.Titles, "", "\t")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(buf))
	},
}

func init() {
	RootCmd.AddCommand(xbox360gamesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// xbox360gamesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// xbox360gamesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
