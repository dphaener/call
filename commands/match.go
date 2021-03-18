package commands

import (
	"os"
	"regexp"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/dphaener/call/slice/string"
)

// Match a given command from the CLI to the list of all commands.
func Match(allCommands []*cobra.Command) []*cobra.Command {
	var matchingCommands []*cobra.Command

	command := slice.At(os.Args, 1)

	commandAliases := viper.GetStringMapString("aliases")
	if commandAliases[command] != "" {
		command = commandAliases[command]
	}

	regex := regexp.MustCompile(`\A` + command)

	for i := range allCommands {
		if regex.MatchString(allCommands[i].Name()) {
			matchingCommands = append(matchingCommands, allCommands[i])
		}
	}

	return matchingCommands
}
