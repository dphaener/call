package prompt

import (
	"reflect"
	"regexp"
)

type Matcher interface {
	SearchField(int) string
	SearchItems() interface{}
	MapItems([]interface{})
}

// Given a matcher that satisfies the Matcher interface, attempts to match the
// given matchString against the list of SearchItems returned by the matcher.
// If a match is found the match interface returned will contain the match.
// If no match is found the original list from the Matcher is mutated to either
// contain the list of potential matches if any were found, or the original full
// list of items from the Matcher.
func Match(m Matcher, matchString string) (match interface{}) {
	var matchList []interface{}
	var matchedItems []interface{}

	allItems := reflect.ValueOf(m.SearchItems())
	regex := regexp.MustCompile(`(?i)` + matchString)

	for i := 0; i < allItems.Len(); i++ {
		if regex.MatchString(m.SearchField(i)) {
			matchedItems = append(matchedItems, allItems.Index(i).Interface())
		}
	}

	switch len(matchedItems) {
	case 0:
		for i := 0; i < allItems.Len(); i++ {
			matchList = append(matchList, allItems.Index(i).Interface())
		}
	case 1:
		match = matchedItems[0]
		return
	default:
		matchList = matchedItems
	}

	m.MapItems(matchList)

	return
}
