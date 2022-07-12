package gomod

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func ReadGoSum(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open go.sum: %w", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	urls := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		url := parseGoSumLine(line)
		if url != "" {
			urls = append(urls, url)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read go.sum: %w", err)
	}
	return urls, nil

}

func parseGoSumLine(line string) string {
	regex := regexp.MustCompile(`^github.com/\S+/\S+`)
	url := regex.FindString(line)
	url = strings.Trim(url, " ")
	return url
}
