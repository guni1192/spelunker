package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "spelunker",
		Short: "",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("hello world")
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
