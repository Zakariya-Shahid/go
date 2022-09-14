package main

import (
	"fmt"
)

type employee struct {
	name     string
	position string
	salary   int
}

type company struct {
	companyName string
	employees   []employee
}

func displayEmployee(emp employee) {
	fmt.Print("\nName = ", emp.name, "\n")
	fmt.Print("Position = ", emp.position, "\n")
	fmt.Print("Salary = ", emp.salary, "\n")
}

func displayCompany(comp company) {
	fmt.Println("\nCompany name = ", comp.companyName)

	for j := 0; j < len(comp.employees); j++ {

		fmt.Println("\n\tEmployee ", j+1)
		displayEmployee(comp.employees[j])
		fmt.Println("\n")
	}

}

func main() {
	var num int
	//var check string
	fmt.Print("Enter the number of employees you want to register : ")
	fmt.Scanf("%d", &num)

	employeeList := make([]employee, num)
	emp1 := employee{"Zakariya", "Developer", 340000}
	employeeList[0] = emp1
	//taking inputs for the details of the struc employee
	for i := 1; i < num; i++ {

		//fmt.Scan(&check)
		fmt.Print("Enter the name of employee ", i+1, " : ")
		fmt.Scan(&employeeList[i].name)
		fmt.Print("Enter the position of the employee ", i+1, " : ")
		fmt.Scan(&employeeList[i].position)
		fmt.Print("Enter the salary of the employee ", i+1, " : ")
		fmt.Scan(&employeeList[i].salary)
		fmt.Println("\n")
	}

	//creating company struct
	comp := company{"Tetra", employeeList}

	displayCompany(comp)
}
