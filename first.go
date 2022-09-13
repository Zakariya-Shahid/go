package main

import "fmt"

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
	var vehicle car

	//taking inputs for the details of the struc car
	fmt.Println("Enter the brand of the car : ")
	fmt.Scanln(&vehicle.brand)
	fmt.Println("Enter the model of the car : ")
	fmt.Scanln(&vehicle.model)
	fmt.Println("Enter the registration number of the car : ")
	fmt.Scan(&vehicle.reg)

	//passing the struct in fucntion
	display(vehicle)

}
