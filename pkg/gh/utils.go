package gh

import (
	"fmt"
	"strings"
)

func parseGitHubURL(url string) (*repository, error) {

	u := url
	u = strings.TrimPrefix(u, "http://")
	u = strings.TrimPrefix(u, "https://")
	u = strings.TrimSuffix(u, ".git")

	arr := strings.Split(u, "/")

	if len(arr) != 3 {
		return nil, fmt.Errorf("invalid github url")
	}

	owner := arr[1]
	repo := arr[2]

	return &repository{owner, repo}, nil
}
