package npm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
)

// Package package.json
type Package struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// HasPackage checks if folder has a package.json
func HasPackage(dir string) bool {
	filePath := path.Join(dir, "/package.json")
	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		return true
	}
	return false
}

// ParsePackage returns the version in the package.json file
func ParsePackage(dir string) (*Package, error) {
	filePath := path.Join(dir, "/package.json")
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("[ParsePackage] open package.json: %v", err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var pkg Package
	json.Unmarshal(byteValue, &pkg)

	return &pkg, nil
}

// Bump version in package.json and git
func Bump(dir string, version string, change string) (string, error) {
	// TODO: set workdir
	message := fmt.Sprintf("chore(package): bump version to %s", version)
	cmd := exec.Command("npm", "version", change, "--message", message)
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("[Bump] exec command: %v %s", err, string(out))
	}
	return string(out), nil
}
