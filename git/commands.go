package git

import (
	"github.com/bitrise-io/go-utils/command"
)

// Log shows the commit logs. The format parameter controls what is shown and how.
func (g *Git) Diff(commitStart string, commitEnd string, path string, otherOptions... string) *command.Model {

	var options []string

	if len(otherOptions) > 0 {
		options = append(options, otherOptions...)
	}

	// Handle commit range
	if len(commitStart) > 0 && len(commitEnd) > 0 {
		options = append(options, commitStart, commitEnd)
	}

	if len(path) > 0 {
		options = append(options, "--", path)
	}


	log := append([]string{"diff"}, options...)
	return g.command(log...)
}

func (g *Git) RevList(commit string, otherOptions... string) *command.Model {
	var options []string

	if len(otherOptions) >0 {
		options = append(options, otherOptions...)
	}
	options = append(options, commit)
	log := append([]string{"rev-list"}, options...)
	return g.command(log...)
}
