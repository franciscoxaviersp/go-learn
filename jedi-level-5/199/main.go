package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle   vehicle
	fourWheel bool
}

type sedan struct {
	vehicle vehicle
	luxury  bool
}

func main() {

	truck_1 := truck{
		vehicle: vehicle{
			doors: 4,
			color: "green",
		},
		fourWheel: true,
	}

	fmt.Println(truck_1)
	fmt.Println("Truck color: ", truck_1.vehicle.color)

	sedan_1 := sedan{
		vehicle: vehicle{
			doors: 2,
			color: "red",
		},
		luxury: false,
	}

	fmt.Println(sedan_1)
	fmt.Println("Sedan color: ", sedan_1.vehicle.color)
}
