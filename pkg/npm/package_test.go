package npm

import (
	"fmt"
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

func TestBump(t *testing.T) {
	runner = TestRunner{}

	d := "./mock"
	v := "1.0.0"
	s := "major"

	res, err := Bump(d, v, s)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(
		t,
		fmt.Sprintf(
			"%s,version,%s,--message,chore(package): bump version to %s\n",
			d, s, v,
		),
		res,
	)
}

func TestPublish(t *testing.T) {
	runner = TestRunner{}

	d := "./mock"

	res, err := Publish(d)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, fmt.Sprintf("%s,publish\n", d), res)
}
