package git

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/dphaener/call/shell"
	"github.com/dphaener/call/slice/string"
)

var descRegex *regexp.Regexp = regexp.MustCompile(`[^\/\w\s\[\]\-\.\,\(\)]`)

// Returns an abbreviated description of the current git branch in the form
// of <commit-sha>-<branch-name>-<commit-description>
func Desc() (description string, err error) {
	sha, err := Sha()
	if err != nil {
		return
	}

	branch, err := Branch()
	if err != nil {
		return
	}

	shortdesc, err := Shortdesc()
	if err != nil {
		return
	}

	rawDesc := strings.Join([]string{sha, branch, shortdesc}, "-")
	description = string(descRegex.ReplaceAll([]byte(rawDesc), []byte("")))

	return
}

// Returns an abbreviated commit sha for the current branch.
func Sha() (sha string, err error) {
	out, err := shell.Execute("git", "rev-parse", "HEAD")
	if err != nil {
		return
	}

	sha = string(out)[:9]
	return
}

// Returns the current branch name.
func Branch() (branch string, err error) {
	out, err := shell.Execute("git", "rev-parse", "--abbrev-ref", "HEAD")
	if err != nil {
		return
	}

	branch = strings.TrimSpace(string(out))
	return
}

// Returns an abbreviated description of the most recent commit.
func Shortdesc() (shortdesc string, err error) {
	out, err := shell.Execute("git", "log", "-1", "--pretty=%B")
	if err != nil {
		return
	}

	desc := strings.TrimSpace(string(out))
	shortdesc = strings.Split(desc, "\n")[0]

	if len(shortdesc) > 70 {
		shortdesc = shortdesc[:70]
	}

	return
}

// Returns the name of the author of the most recent commit.
func Author() (author string, err error) {
	out, err := shell.Execute("git", "log", "-1", "--pretty=format:%an")
	if err != nil {
		return
	}

	author = strings.TrimSpace(string(out))
	return
}

// Returns a list of all commits since the last tag. Also parses the list and
// removes the commit sha, replacing it with a "-", whici is suitable for use
// in a changelog.
func RecentChanges() (changes string, err error) {
	tag, err := LatestTag()
	if err != nil {
		return
	}

	changes, err = CommitsBetween(tag, "HEAD")

	return
}

// Syncs the given branch with the upstream branch.
func Sync(branch string) (err error) {
	branchArg := fmt.Sprintf("%s:%s", branch, branch)
	_, err = shell.Execute("git", "fetch", "origin", branchArg)
	return
}

// Fetch the most recent tag for a repo.
func LatestTag() (tag string, err error) {
	tag, err = shell.Execute("git", "describe", "--tags", "--abbrev=0")
	if err != nil {
		return
	}

	tag = strings.Replace(tag, "\n", "", -1)
	return
}

// Fetch the last n tags from the current repository.
func LastNTags(n int) (tags []string, err error) {
	tagString, err := shell.Execute("git", "tag", "-l", "--sort=version:refname")
	if err != nil {
		return
	}

	tags = strings.Split(tagString, "\n")
	tags = slice.Compact(tags)
	tags = slice.Reverse(tags)
	tags = slice.To(tags, n)

	return
}

// Fetch the commits between two given tags, branches, or commits.
func CommitsBetween(item1 string, item2 string) (changes string, err error) {
	args := []string{
		"--no-pager",
		"log",
		fmt.Sprintf("%s..%s", item1, item2),
		"--pretty=format:%s",
		"--reverse",
		"--no-merges",
	}
	commits, err := shell.Execute("git", args...)

	for _, line := range strings.Split(commits, "\n") {
		changes = changes + "- " + line + "\n"
	}
	changes = strings.TrimRight(changes, "\n")

	return
}
