package commands

import (
	"regexp"

	"github.com/dphaener/call/logger"
	"github.com/dphaener/call/slice/string"
)

var appName string
var rackName string

type ExpandedCommand struct {
	Command    string
	Arguments  []string
	RawCommand string
}

// Expands a given dynamic command. Creates a map of the list of required
// arguments to the passed in arguments and evaluates the arguments to obtain
// their final value.
func Expand(dynamicCommand string, requiredArguments []string, args []string) ExpandedCommand {
	argumentMap := make(map[string]string)

	for i := range requiredArguments {
		argName := requiredArguments[i]
		arg := slice.At(args, i)
		argumentMap[argName] = arg
	}

	expandedCommand, err := ProcessTemplate(dynamicCommand, argumentMap)
	if err != nil {
		logger.ExitWithMessage(err)
	}
	splitRegex := regexp.MustCompile(`\s`)
	splitCommand := splitRegex.Split(expandedCommand, -1)

	return ExpandedCommand{
		Command:    splitCommand[0],
		Arguments:  splitCommand[1:],
		RawCommand: expandedCommand,
	}
}
