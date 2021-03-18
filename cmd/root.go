package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/user"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/dphaener/call/commands"
	"github.com/dphaener/call/logger"
	"github.com/dphaener/call/slice/string"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:     "cmd",
	Short:   "A configurable CLI built to help developers with productivity",
	Long:    `cmd is a CLI tool built to allow heavy customization via configuration.`,
	Version: "1.0.0",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	initConfig()

	dynamicCommands := commands.Parse()
	rootCmd.AddCommand(dynamicCommands...)

	osArgs := os.Args[1:]
	cmd, _, err := rootCmd.Find(osArgs)

	if err != nil || cmd == nil || len(osArgs) == 0 {
		command, matchErr := commands.Select(rootCmd.Commands())

		if matchErr == nil {
			positionalArgs := slice.From(os.Args, 2)
			args := append([]string{command}, positionalArgs...)
			rootCmd.SetArgs(args)
		}
	}
	if err := rootCmd.Execute(); err != nil {
		logger.ExitWithMessage(err)
	}
}

func init() {}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	user, err := user.Current()
	if err != nil {
		logger.ExitWithMessage(err)
	}

	configPath := fmt.Sprintf("%s/.config/cmd", user.HomeDir)

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(configPath)
	err = viper.ReadInConfig()

	if err != nil && !os.IsNotExist(err) {
		logger.ExitWithMessage(err)
	}

	if err != nil && os.IsNotExist(err) {
		err = errors.New("This is basically useless without a config file, so go make one!")
		logger.ExitWithMessage(err)
	}
}
