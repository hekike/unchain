package markdown

import (
	"bytes"
	"fmt"
	"time"

	"github.com/hekike/conventional-commits/pkg/model"
)

// Generate generate markdown output
func Generate(version string, commits []model.ConventionalCommit) string {
	var out bytes.Buffer
	var patch = false
	var minor = false
	var major = false

	// Tag Header
	date := time.Now().Format("2006-01-02")
	out.WriteString(fmt.Sprintf("<a name=\"%s\"></a>\n", version))
	out.WriteString(fmt.Sprintf("## %s (%s)\n\n\n", version, date))

	// Patch
	for _, commit := range commits {
		if commit.SemVerChange == model.Patch {
			if patch == false {
				out.WriteString("#### Bug Fixes\n")
			}
			out.WriteString(getCommitLine(&commit))
			patch = true
		}
	}
	if patch == true {
		out.WriteString("\n")
	}

	// Minor
	for _, commit := range commits {
		if commit.SemVerChange == model.Minor {
			if minor == false {
				out.WriteString("\n#### Features\n")
			}
			out.WriteString(getCommitLine(&commit))
			minor = true
		}
	}
	if minor == true {
		out.WriteString("\n")
	}

	// Major
	for _, commit := range commits {
		if commit.SemVerChange == model.Major {
			if major == false {
				out.WriteString("\n#### Breaking Changes\n")
			}
			out.WriteString(getBreakingLine(&commit))
			major = true
		}
	}
	if major == true {
		out.WriteString("\n")
	}

	return out.String()
}

func getCommitLine(commit *model.ConventionalCommit) string {
	var out bytes.Buffer

	out.WriteString("\n* ")
	if len(commit.Component) > 0 {
		c := fmt.Sprintf("**%s:** ", commit.Component)
		out.WriteString(c)
	}
	out.WriteString(commit.Description)
	out.WriteString(" ")
	out.WriteString(commit.Hash)

	return out.String()
}

func getBreakingLine(commit *model.ConventionalCommit) string {
	var out bytes.Buffer

	out.WriteString("\n* ")
	out.WriteString(commit.Breaking)
	out.WriteString(" ")
	out.WriteString(commit.Hash)

	return out.String()
}
