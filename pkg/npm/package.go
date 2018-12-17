package npm

import (
	"encoding/json"
	"io/ioutil"
	"os"
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
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var pkg Package
	json.Unmarshal(byteValue, &pkg)

	return &pkg, nil
}
