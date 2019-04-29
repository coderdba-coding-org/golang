package main

type Emp struct {
        empname   string
        empid     int
        employed  bool       
        skills    []string
}

func (emp *Emp) initialize() error {

  emp.setEmployed()
  return nil

}

func (emp *Emp) setEmployed() error {

  emp.employed = true
  return nil

}
