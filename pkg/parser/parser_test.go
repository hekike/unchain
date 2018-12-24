package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToSemVerChange(t *testing.T) {
	res := ToSemVerChange("patch")
	assert.Equal(t, Patch, res)

	res = ToSemVerChange("minor")
	assert.Equal(t, Minor, res)

	res = ToSemVerChange("major")
	assert.Equal(t, Major, res)

	res = ToSemVerChange("invalid")
	assert.Equal(t, SemVerChange(""), res)
}
