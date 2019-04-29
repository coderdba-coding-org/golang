package main

type Emp struct {
        empname   string
        empid     int
        employed  bool       
        salary    Salary
        skills    []string
}

// NOTE: Pointer receiver ensure that the value of the Emp structure elements get set directly

func (emp *Emp) initialize() error {

  emp.setEmployed()
  return nil

}

func (emp *Emp) setEmployed() error {

  emp.employed = true
  return nil

}

func (emp *Emp) setSalary() error {

  emp.salary.initialize()  
  return nil

}
