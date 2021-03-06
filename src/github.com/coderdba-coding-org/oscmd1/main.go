package main

import (
	"fmt"
	"os/exec"
        "os"
)

func main() {  

    var (
          //cmdOut is output of command comes as a byte array - do a string(cmdOut) to get string
          cmdOut []byte
          err    error
    )
    
    

    cmd := "hostname"

    fmt.Println ("Running a command: ", cmd)

    if cmdOut, err = exec.Command(cmd).Output(); err != nil {

		fmt.Fprintln(os.Stderr, "There was an error running the command: ", err)
    
		os.Exit(1)
	}

    fmt.Fprintln(os.Stdout, "Hostname as bytes is: ", cmdOut)
    fmt.Fprintln(os.Stdout, "Hostname as strings is: ",  string(cmdOut))

    fmt.Println ("Ran it well")
}
