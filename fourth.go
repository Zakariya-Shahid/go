package main

import (
	"crypto/sha256"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	roll_no  int
	name     string
	address  string
	subjects []string
}

func newStudent(roll_no int, name string, address string, sub []string) *Student {
	student := new(Student)
	student.roll_no = roll_no
	student.name = name
	student.address = address

	student.subjects = sub

	return student
}

func display(list []Student) {
	var length int
	length = len(list)

	for i := 0; i < length; i++ {
		fmt.Println("\n\n\t", strings.Repeat("=", 25), " Student ", i+1, strings.Repeat("=", 25), "\n")
		fmt.Println("\n\tRoll No : ", list[i].roll_no)
		fmt.Println("\n\tName : ", list[i].name)
		fmt.Println("\n\tAddress : ", list[i].address)
		fmt.Println("\n\tCourses : ")

		var hash string

		hash += strconv.Itoa(list[i].roll_no) + list[i].name + list[i].address

		for j := 0; j < len(list[i].subjects); j++ {
			fmt.Println("\n\t\t", j+1, list[i].subjects[j])
			hash += list[i].subjects[j]
		}

		sum := sha256.Sum256([]byte(hash))
		fmt.Printf("Hash = %x\n", sum)
	}
}

func menu() {
	fmt.Println("\n\n\t\t", strings.Repeat("=", 25), "Menu", strings.Repeat("=", 25))
	fmt.Println("Enter 1 to add student.")
	fmt.Println("Enter 2 to print the student list.")
	fmt.Println("Enter 0 to Quit.")
}

func main() {

	studentList := make([]Student, 0)

	subjects := []string{"Blockchain", "PDC", "CNET"}

	var input int
	check := 1

	var st Student

	for check >= 0 {
		menu()
		fmt.Print("Enter a number : ")
		fmt.Scan(&input)

		switch input {
		case 0:
			fmt.Println("\nExiting...")
			fmt.Println("Done.")
			os.Exit(0)
			break
		case 1:
			fmt.Print("\n\nEnter the Roll no : ")
			fmt.Scan(&st.roll_no)
			fmt.Print("Enter the name : ")
			fmt.Scan(&st.name)
			fmt.Print("Enter the address : ")
			fmt.Scan(&st.address)

			st.subjects = subjects

			studentList = append(studentList, st)

			fmt.Println("Student added successfully.\n")
			break
		case 2:
			display(studentList)
			fmt.Println("\n\tEnd of list")
		default:
			fmt.Println("Invalid Input.\n")

		}

	}
}
