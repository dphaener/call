package prompt

import (
	"strings"
)

type Searcher interface {
	SearchField(int) string
}

// Given a searcher that satisfies the Searcher interface, attempts to find the
// field in the list of items at the given index. Returns a function that
// satisfies the promptui Searcher interface.
func Search(s Searcher) func(string, int) bool {
	return func(input string, index int) bool {
		field := strings.Replace(strings.ToLower(s.SearchField(index)), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(field, input)
	}
}
