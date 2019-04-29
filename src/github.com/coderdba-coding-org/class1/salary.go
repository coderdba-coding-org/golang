package main

type Salary float64


// Both with and without * seem to work (pointer and value retrievers)
//func (salary *Salary) initialize() Salary {
func (salary Salary) initialize() Salary {

  return 1000.0

}

// Only without * seem to work (value retriever)
func (salary Salary) calctax() Salary {

  return salary * 0.2

}
