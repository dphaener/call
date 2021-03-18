package commands

import (
	"regexp"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/dphaener/call/logger"
	"github.com/dphaener/call/shell"
)

type Command struct {
	Use         string
	Short       string
	Description string
	Expanded    string
	Args        []string
	Name        string
}

// Parse the dynamically loaded commands from a configuration file and add them
// to the list of Cobra commands along with the passed in default commands.
func Parse() []*cobra.Command {
	var cobraCommands []*cobra.Command

	allCommands := compileCommands()
	confirms := compileConfirms()

	for _, kommand := range allCommands {
		dynamicCommand := kommand.Expanded
		requiredArguments := kommand.Args

		if kommand.Description == "" {
			kommand.Description = kommand.Short
		}

		newCommand := &cobra.Command{
			Use:   kommand.Use,
			Short: kommand.Short,
			Long:  kommand.Description,
			Run: func(cmd *cobra.Command, args []string) {
				var confirm bool

				ec := Expand(dynamicCommand, requiredArguments, args)
				noConfirm := viper.GetBool("noConfirm")

				for i := range confirms {
					if confirms[i].MatchString(dynamicCommand) {
						confirm = true
						break
					}
				}

				if confirm && !noConfirm {
					prompt := promptui.Prompt{
						Label:     "About to run: " + ec.RawCommand,
						IsConfirm: true,
						Default:   "n",
					}

					_, err := prompt.Run()

					if err != nil {
						logger.ExitWithMessage(err)
					}
				}

				logger.NewLine()
				logger.Info("running", ec.RawCommand)
				logger.NewLine()

				err := shell.Interactive(ec.RawCommand)
				if err != nil {
					logger.ExitWithMessage(err)
				}
			},
		}

		cobraCommands = append(cobraCommands, newCommand)
	}

	return cobraCommands
}

func compileCommands() map[string]Command {
	in := viper.Get("commands")
	rawDynamicCommands, _ := in.(map[string]interface{})
	allCommands := make(map[string]Command)

	for commandName, kommand := range rawDynamicCommands {
		var args []string

		value := kommand.(map[string]interface{})
		interfaces := value["args"].([]interface{})
		description, descOk := value["long"].(string)

		if !descOk {
			description = value["short"].(string)
		}

		for i := range interfaces {
			args = append(args, interfaces[i].(string))
		}

		allCommands[commandName] = Command{
			Name:        commandName,
			Use:         value["use"].(string),
			Short:       value["short"].(string),
			Description: description,
			Expanded:    value["expanded"].(string),
			Args:        args,
		}
	}

	return allCommands
}

func compileConfirms() (confirms []*regexp.Regexp) {
	rawConfirms := viper.GetStringSlice("confirms")

	for i := range rawConfirms {
		confirms = append(confirms, regexp.MustCompile(rawConfirms[i]))
	}

	return
}
