package main

import (
	"fmt"
	"os/exec"
        "os"
)

func usage() {

    fmt.Printf("Program arguments - count: %d, arguments: %s \n", len(os.Args),  os.Args) 
    
    if len(os.Args) < 2 {
    
       fmt.Printf("Usage: %s command_to_run \n", os.Args[0])
       os.Exit(1)

    }
}

func main() {  

    var (
          cmdOut []byte  //cmdOut is output of command comes as a byte array - do a string(cmdOut) to get string
          err    error
    )
    
    usage()

    cmd := os.Args[1]

    fmt.Println ("Running a command: ", cmd)

    if cmdOut, err = exec.Command(cmd).Output(); err != nil {

		fmt.Fprintln(os.Stderr, "There was an error running the command: ", err)
    
		os.Exit(1)
	}

    //fmt.Fprintln(os.Stdout, "Output as bytes is: ", cmdOut)
    fmt.Fprintln(os.Stdout, "Output as strings is: \n",  string(cmdOut))

    fmt.Println ("Ran it well")
}

