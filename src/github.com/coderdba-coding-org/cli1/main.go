package main

import (
  "fmt"
  "log"
  "os"

  "github.com/urfave/cli"
)

func main() {
  app := cli.NewApp()
  app.Name = "cli1"
  app.Usage = "make an explosive entrance"
  app.Action = func(c *cli.Context) error {
    fmt.Println("boom! I say!")


    // number of arguments
    fmt.Println("Number of arguments: ", c.NArg())

		//print first argument
    fmt.Println("Hello println first argument: ", c.Args().Get(0))
    //fmt.Printf("Hello printf %q ", c.Args().Get(0))

		// Loop through args
		for i := 0; i < c.NArg(); i++ {

		fmt.Println("Argument Println: ", i, c.Args().Get(i))
		fmt.Printf("%s %d %s \n", "Argument Printf: ", i, c.Args().Get(i))

	  }

		//
    return nil
  }

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}
