package model

// SemVerChange describes the semver change type
type SemVerChange string

const (
	// Patch semver change
	Patch SemVerChange = "patch"
	// Minor semver change
	Minor SemVerChange = "minor"
	// Major semver change
	Major SemVerChange = "major"
)

// ConventionalCommit parsed commit
type ConventionalCommit struct {
	Hash         string
	Type         string
	Component    string
	Description  string
	Body         string
	Footer       string
	Breaking     string
	SemVerChange SemVerChange
	SemVer       string
}
