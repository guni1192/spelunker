package gomod

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseGoSumLine(t *testing.T) {
	url := parseGoSumLine("github.com/stretchr/testify v1.8.0/go.mod h1:yNjHg4UonilssWZ8iaSj1OCr/vHnekPRkoO+kdMU+MU=")
	assert.Equal(t, "github.com/stretchr/testify", url)
}
