package changelog

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"time"

	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

var changelogFile = "CHANGELOG.md"

// Prepend prepend content to file
func Prepend(dir string, content string) error {
	filePath := path.Join(dir, changelogFile)
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return fmt.Errorf("[Prepend] open file: %v", err)
	}
	defer f.Close()

	current, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("[Prepend] read file: %v", err)
	}

	writer := bufio.NewWriter(f)
	writer.WriteString(content)
	writer.Write(current)

	if err := writer.Flush(); err != nil {
		return fmt.Errorf("[Prepend] flush file: %v", err)
	}

	return nil
}

// GitCommit adds CHANGELOG.md to Git
func GitCommit(dir string, version string) error {
	r, err := git.PlainOpen(dir)
	if err != nil {
		return fmt.Errorf("[GitAdd] open repo: %v", err)
	}

	w, err := r.Worktree()
	if err != nil {
		return fmt.Errorf("[GitAdd] worktree: %v", err)
	}

	_, err = w.Add(changelogFile)
	if err != nil {
		return fmt.Errorf("[GitAdd] worktree add (%s): %v", changelogFile, err)
	}

	message := fmt.Sprintf("chore(changelog): update for version %s", version)
	_, err = w.Commit(message, &git.CommitOptions{
		Author: &object.Signature{
			Name:  "Release",
			Email: "release",
			When:  time.Now(),
		},
	})
	if err != nil {
		return fmt.Errorf("[GitAdd] git commit: %v", err)
	}

	return nil
}
