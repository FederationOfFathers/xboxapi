package cmd

import (
	"fmt"
	"os"

	"github.com/FederationOfFathers/xboxapi"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

func client() *xboxapi.Client {
	xboxapi.DebugHTTP = viper.GetBool("debug")
	return xboxapi.New(&xboxapi.Config{
		APIKey:   viper.GetString("key"),
		Language: viper.GetString("lang"),
	})
}

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "xboxapi",
	Short: "Command Line Interface for https://xboxapi.com/",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.xboxapi.yaml)")
	RootCmd.PersistentFlags().String("key", "", "your https://xboxapi.com/ API Key")
	RootCmd.PersistentFlags().String("lang", xboxapi.DefaultConfig.Language, "desired response language")
	RootCmd.PersistentFlags().Bool("debug", xboxapi.DebugHTTP, "enable debugging of http request/response")
	viper.BindPFlag("key", RootCmd.PersistentFlags().Lookup("key"))
	viper.BindPFlag("lang", RootCmd.PersistentFlags().Lookup("lang"))
	viper.BindPFlag("debug", RootCmd.PersistentFlags().Lookup("debug"))
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".xboxapi" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".xboxapi")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		// fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
