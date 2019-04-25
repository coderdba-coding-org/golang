package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
  Use:   "version",
  Short: "Prints version number",
  Long:  `Prints version number for this application`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Cobra3 version v0.1")
  },
}
