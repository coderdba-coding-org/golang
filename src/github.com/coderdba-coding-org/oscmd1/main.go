package main

import (
	"fmt"
	"os/exec"
        "os"
)

func main() {  

    var (
          cmdOut []byte
          //cmdOut []string
          err    error
    )
    
    //fmt.Println("hostname: ", exec.Command("hostname"))

    cmdName := "hostname"
    if cmdOut, err = exec.Command(cmdName).Output(); err != nil {

		fmt.Fprintln(os.Stderr, "There was an error running the command: ", err)
    
		//os.Exit(1)

	}

    fmt.Fprintln(os.Stderr, cmdOut)
    fmt.Fprintln(os.Stderr, string(cmdOut))
    fmt.Println ("Ran it well")
}
