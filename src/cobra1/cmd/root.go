package cmd

import (
  "fmt"
  "os"

  "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
  Use:   "cobra1",
  Short: "cobra1 is a runner",
  Long: `A cobra command program
                just to get a hang of it`,
                
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("it is running!")
  },
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
