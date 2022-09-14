package main

import "fmt"

type car struct {
	brand string
	model string
	reg   int
}

func display(vehicle car) {
	fmt.Print("Brand = ", vehicle.brand, "\n")
	fmt.Print("Model = ", vehicle.model, "\n")
	fmt.Print("Registration number = ", vehicle.reg, "\n")
}

func main() {
	var vehicle car

	//taking inputs for the details of the struc car
	fmt.Print("\nEnter the brand of the car : ")
	fmt.Scanln(&vehicle.brand)
	fmt.Print("\nEnter the model of the car : ")
	fmt.Scanln(&vehicle.model)
	fmt.Print("\nEnter the registration number of the car : ")
	fmt.Scan(&vehicle.reg)

	//passing the struct in fucntion
	fmt.Println("\n")
	fmt.Println("\n\t\tCar\n\n")
	display(vehicle)
	fmt.Println("\n")

}
