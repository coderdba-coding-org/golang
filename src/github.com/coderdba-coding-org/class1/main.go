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
  
   emp1.salary = emp1.salary.initialize()
   fmt.Println ("Employee details after initializing salary: ", emp1)

   //not working yet
   fmt.Println ("Tax amount: ", emp1.salary.calctax())

}
