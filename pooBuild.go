package main

import (
	"fmt"

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
