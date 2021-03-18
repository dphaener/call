package template

import (
	"error"
	"fmt"

	"github.com/manifoldco/promptui"

	"github.com/dphaener/call/logger"
	"github.com/dphaener/call/prompt"
)

type Select struct {
	Line     string
	Details  Details
	Selected string
}

// Compile the template for use as a promptui template.
func (s Select) Compile() (templates *promptui.SelectTemplates) {
	templates = &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   fmt.Sprintf("%s %s", prompt.SelectEmoji, s.Line),
		Inactive: fmt.Sprintf("   %s", s.Line),
		Selected: fmt.Sprintf("%s %s", prompt.CheckEmoji, s.Selected),
		Details:  s.Details.Compile(),
	}

	return
}

type Details struct {
	Label  string
	Fields []Field
}

type Field struct {
	Name     string
	Format   string
	Template string
}

// Compile the tempalte details into a string suitable for use in a promptui
// template.
func (d Details) Compile() (details string) {
	details = fmt.Sprintf("\n--------- %s ----------", d.Label)

	for _, field := range d.Fields {
		switch field.Format {
		case "datetime":
			details = fmt.Sprintf(datetimeTemplate, details, field.Name, field.Name)
		case "custom":
			if field.Template == "" {
				err := error.New("You must provide a template if using the custom format")
				logger.ExitWithMessage(err)
			}
			details = fmt.Sprintf("%s\n%s", details, field.Template)
		default:
			if field.Template == "" {
				details = fmt.Sprintf(defaultTemplate, details, field.Name, field.Name)
			} else {
				details = fmt.Sprintf("%s\n%s", details, field.Template)
			}
		}
	}

	return
}

var datetimeTemplate = "%s\n{{ \"%s:\" | faint }} {{ .%s.Format \"Jan 02, 2006 15:04:05 UTC\" }}"
var defaultTemplate = "%s\n{{ \"%s:\" | faint }} {{ .%s }}"
