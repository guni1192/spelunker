package main

import (
	"fmt"
	"os"

	"github.com/guni1192/spelunker/pkg/gh"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "spelunker",
		Short: "",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			c := gh.NewGitHubClient()
			fmt.Println("Starting GitHub Repository archive check")
			url := args[0]
			archived, err := c.IsArchivedFromURL(url)
			if err != nil {
				return fmt.Errorf("failed to get repository status: %w", err)
			}
			if *archived {
				fmt.Println("ARCHIVED", url)
			} else {
				fmt.Println("ACTIVE", url)
			}

			return nil
		},
	}
)

func init() {
}

func main() {

	err := rootCmd.Execute()

	if err != nil {
		fmt.Fprintln(os.Stderr, "%w", err)
		os.Exit(1)
	}

}
