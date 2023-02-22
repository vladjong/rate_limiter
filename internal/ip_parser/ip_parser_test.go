package ipparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseIp(t *testing.T) {
	assert := assert.New(t)
	in := "10.120.100.255"
	expected := "10.120.64.0"
	mask := 18

	ipParser := New(int8(mask))

	result, err := ipParser.GetParentIp(in)

	assert.Nil(err)
	assert.Equal(expected, result)
}
