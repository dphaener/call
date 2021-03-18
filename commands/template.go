package commands

import (
	"bytes"
	"os"
	"text/template"

	"github.com/dphaener/call/git"
	"github.com/dphaener/call/logger"
)

var templFunctions = template.FuncMap{
	"gitDesc": func() (desc string) {
		desc, err := git.Desc()
		if err != nil {
			logger.ExitWithMessage(err)
		}

		return
	},
	"gitSha": func() (sha string) {
		sha, err := git.Sha()
		if err != nil {
			logger.ExitWithMessage(err)
		}

		return
	},
	"gitBranch": func() (branch string) {
		branch, err := git.Branch()
		if err != nil {
			logger.ExitWithMessage(err)
		}

		return
	},
	"gitShortDesc": func() (shortDesc string) {
		shortDesc, err := git.Shortdesc()
		if err != nil {
			logger.ExitWithMessage(err)
		}

		return
	},
	"gitAuthor": func() (author string) {
		author, err := git.Author()
		if err != nil {
			logger.ExitWithMessage(err)
		}

		return
	},
	"getEnv": func(varName string) string {
		return os.Getenv(varName)
	},
	"gitCommitsBetween": func(item1 string, item2 string) string {
		changes, err := git.CommitsBetween(item2, item1)
		if err != nil {
			logger.ExitWithMessage(err)
		}

		return changes
	},
}

// Process an expanded command template using the arguments provided.
func ProcessTemplate(dynamicCommand string, argumentMap map[string]string) (string, error) {
	t, err := template.New("command").Funcs(templFunctions).Parse(dynamicCommand)

	if err != nil {
		return "", err
	}

	var expanded bytes.Buffer

	if err = t.Execute(&expanded, argumentMap); err != nil {
		return "", err
	}

	return expanded.String(), err
}
