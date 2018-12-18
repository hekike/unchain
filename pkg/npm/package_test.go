package npm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasPackage(t *testing.T) {
	res := HasPackage("./mock")
	assert.Equal(t, true, res)
}

func TestParsePackage(t *testing.T) {
	res, err := ParsePackage("./mock")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, "mock", res.Name)
	assert.Equal(t, "1.0.0", res.Version)
}
