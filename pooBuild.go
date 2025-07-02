package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gooBuild",
	Short: "Build a JS bundle file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		sourceFile, _ := cmd.Flags().GetString("source")
		targetFile, _ := cmd.Flags().GetString("target")
		fmt.Println("source")
		fmt.Println(sourceFile)
		fmt.Println("target")
		fmt.Println(targetFile)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}
