package gh

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseGitHubURL(t *testing.T) {
	url := "github.com/guni1192/dotfiles"

	repo, err := parseGitHubURL(url)
	assert.Nil(t, err)
	assert.Equal(t, "guni1192", repo.owner, "unexpected repo.owner")
	assert.Equal(t, "dotfiles", repo.name, "unexpected repo.name")
}

func TestParseGitHubURLPrefixHTTP(t *testing.T) {
	url := "http://github.com/guni1192/dotfiles"

	repo, err := parseGitHubURL(url)
	assert.Nil(t, err)
	assert.Equal(t, "guni1192", repo.owner, "unexpected repo.owner")
	assert.Equal(t, "dotfiles", repo.name, "unexpected repo.name")
}

func TestParseGitHubURLPrefixHTTPS(t *testing.T) {
	url := "https://github.com/guni1192/dotfiles"

	repo, err := parseGitHubURL(url)
	assert.Nil(t, err)
	assert.Equal(t, "guni1192", repo.owner, "unexpected repo.owner")
	assert.Equal(t, "dotfiles", repo.name, "unexpected repo.name")
}

func TestParseGitHubURLSuffixGit(t *testing.T) {
	url := "https://github.com/guni1192/dotfiles.git"

	repo, err := parseGitHubURL(url)
	assert.Nil(t, err)
	assert.Equal(t, "guni1192", repo.owner, "unexpected repo.owner")
	assert.Equal(t, "dotfiles", repo.name, "unexpected repo.name")
}

func TestParseGitHubURLInvalidURL(t *testing.T) {
	url := "https://github.com/"

	_, err := parseGitHubURL(url)
	assert.NotNil(t, err)
}
