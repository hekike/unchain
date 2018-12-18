package npm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

// Package package.json
type Package struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// HasPackage checks if folder has a package.json
func HasPackage(path string) bool {
	if _, err := os.Stat(path + "/package.json"); !os.IsNotExist(err) {
		return true
	}
	return false
}

// ParsePackage returns the version in the package.json file
func ParsePackage(path string) (*Package, error) {
	jsonFile, err := os.Open(path + "/package.json")
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
func Bump(path string, version string, change string) (string, error) {
	// TODO: set workdir
	message := fmt.Sprintf("chore(package): bump version to %s", version)
	cmd := exec.Command("npm", "version", change, "--message", message)
	cmd.Dir = path
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("[Bump] exec command: %v %s", err, string(out))
	}
	return string(out), nil
}
