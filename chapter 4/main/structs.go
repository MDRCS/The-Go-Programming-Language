package main

import (
	"fmt"
	"time"
)

const (
	layoutISO = "2006-01-02"
	layoutUS  = "January 2, 2006"
)


type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

var emp Employee


func main() {
	date := "1996-07-16"
	t, _ := time.Parse(layoutISO, date)
	fmt.Println(t)                  // 1999-12-31 00:00:00 +0000 UTC
	fmt.Println(t.Format(layoutUS)) // December 31, 1999
	emp = Employee{1,"Med","Abou firass el hamdani",t,"AI SWE",100000,3}
	fmt.Println(emp)

	position := &emp.Position
	*position = "Senior " + *position
	fmt.Println(emp)

	var EmployeeOfTheMonth *Employee = &emp
	EmployeeOfTheMonth.Position += " (proactive team player)"

	fmt.Println(EmployeeOfTheMonth)
	EmployeeByID(1).Salary = 0 //fired for no reason
	fmt.Println(emp)
}

func EmployeeByID(id int) *Employee {
	return &emp
}
