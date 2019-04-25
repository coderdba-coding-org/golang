package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(prmsgCmd)
}

var prmsgCmd = &cobra.Command{
  Use:   "prmsg",
  Short: "Prints a message",
  Long:  `Prints a message you send it`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("This is your message: ", args[0])
  },
}
