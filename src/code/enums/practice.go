package main

import "fmt"

type OrderStatus int

const (
	RECIEVED OrderStatus = iota
	CONFIRMED
	PREPARED
	DELIVERED
)

type Vehicle string

const (
	CAR      Vehicle = "Car"
	AIRPLANE Vehicle = "AirPlane"
	RAIL     Vehicle = "Rail"
	SHIP     Vehicle = "Ship"
)

func ChangeStatus(status OrderStatus) {
	fmt.Println("The current status is", status)
}

func GetVehicleType(vehicleType Vehicle) {
	fmt.Println("Vehicle Type is", vehicleType)
}

func main() {
	ChangeStatus(DELIVERED)
	GetVehicleType(SHIP)
}
