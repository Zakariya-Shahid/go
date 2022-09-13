package main

import (
	"fmt"
)

type car struct {
	brand string
	model string
	reg   int
}

func display(vehicle car) {
	fmt.Print("Brand of the car = ", vehicle.brand, "\n")
	fmt.Print("Model of the car = ", vehicle.model, "\n")
	fmt.Print("Registration number of the car = ", vehicle.reg, "\n")
}

func main() {
	var num int
	//var check string
	fmt.Print("Enter the number of cars : ")
	fmt.Scanf("%d", &num)

	vehicles := make([]car, num)
	//taking inputs for the details of the struc car
	for i := 0; i < num; i++ {

		//fmt.Scan(&check)
		fmt.Print("Enter the brand of the car ", i+1, " : ")
		fmt.Scan(&vehicles[i].brand)
		fmt.Print("Enter the model of the car ", i+1, " : ")
		fmt.Scan(&vehicles[i].model)
		fmt.Print("Enter the registration number of the car ", i+1, " : ")
		fmt.Scan(&vehicles[i].reg)
		fmt.Println("\n")
	}
	//passing the struct in fucntion
	for j := 0; j < num; j++ {

		fmt.Println("\n\tCar ", j+1)
		display(vehicles[j])
		fmt.Println("\n")
	}

}
