package prompt

import (
	"io"
	"os"

	"github.com/manifoldco/promptui"
)

type Select struct {
	Label     string
	Items     interface{}
	Size      int
	Templates *promptui.SelectTemplates
	Searcher  func(string, int) bool
	Stdout    io.WriteCloser
}

// Runs the select prompt. Delegates out to the promptui Run method. This is a
// convenience wrapper that always uses the BellSkipper as the stdout in order
// to prevent the bell noise when using the arrow keys to navigate the search
// menu.
func (s Select) Run() (index int, err error) {
	var size int = s.Size
	if size == 0 {
		size = 10
	}

	prompt := promptui.Select{
		Label:     s.Label,
		Items:     s.Items,
		Size:      size,
		Templates: s.Templates,
		Searcher:  s.Searcher,
		Stdout:    &BellSkipper{},
	}

	index, _, err = prompt.Run()

	return
}

// bellSkipper implements an io.WriteCloser that skips the terminal bell
// character (ASCII code 7), and writes the rest to os.Stderr. It is used to
// replace readline.Stdout, that is the package used by promptui to display the
// prompts.
//
// This is a workaround for the bell issue documented in
// https://github.com/manifoldco/promptui/issues/49.
type BellSkipper struct{}

// Write implements an io.WriterCloser over os.Stderr, but it skips the terminal
// bell character.
func (bs *BellSkipper) Write(b []byte) (int, error) {
	const charBell = 7 // c.f. readline.CharBell
	if len(b) == 1 && b[0] == charBell {
		return 0, nil
	}
	return os.Stderr.Write(b)
}

// Close implements an io.WriterCloser over os.Stderr.
func (bs *BellSkipper) Close() error {
	return os.Stderr.Close()
}
