package main

import (
	"fmt"
)

func main() {

   emp1 := Emp{
        empname:   "abc",
        empid:     101,
   } 

   fmt.Println ("Employee details: ", emp1)

   emp1.initialize()

   fmt.Println ("Employee details after initializing: ", emp1)

}
