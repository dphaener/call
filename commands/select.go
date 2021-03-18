package commands

import (
	"errors"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"

	"github.com/dphaener/call/logger"
	"github.com/dphaener/call/prompt"
)

// Select a command from the list of available commands.
func Select(allCommands []*cobra.Command) (string, error) {
	matchingCommands := Match(allCommands)

	if len(matchingCommands) == 1 {
		return matchingCommands[0].Name(), nil
	} else if len(matchingCommands) > 1 {
		prompt := promptui.Select{
			Label:     "Which Command",
			Items:     matchingCommands,
			Size:      10,
			Templates: commandTemplates,
			Searcher:  commandSearcher(matchingCommands),
		}

		i, _, err := prompt.Run()

		if err != nil {
			logger.ExitWithMessage(err)
		}

		return matchingCommands[i].Name(), nil
	}

	return "", errors.New("Command not found")
}

var commandLineTemplate string = " {{ .Name | blue }}"
var commandTemplates = &promptui.SelectTemplates{
	Label:    "{{ . }}?",
	Active:   prompt.SelectEmoji + commandLineTemplate,
	Inactive: "  " + commandLineTemplate,
	Selected: "Command: " + prompt.CheckEmoji + commandLineTemplate,
	Details: `
--------- Description ----------

{{ .Long | faint }}`,
}

var commandSearcher = func(commands []*cobra.Command) func(string, int) bool {
	searchFunc := func(input string, index int) bool {
		command := commands[index]
		name := strings.Replace(strings.ToLower(command.Name()), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	return searchFunc
}
