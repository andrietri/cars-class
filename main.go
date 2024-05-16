package main

import (
	"fmt"
)

// define type of tire
type Tire interface {
	TireType() string
}

type RubberTire struct{}
type WoodenTire struct{}
type IronTire struct{}

func (b RubberTire) TireType() string {
	return "Ban Karet"
}

func (b WoodenTire) TireType() string {
	return "Ban Kayu"
}

func (b IronTire) TireType() string {
	return "Ban Besi"
}

// define wheel
type Wheel struct {
	tire Tire
}

// define door
type Door struct {
	name string
}

func (p Door) Knock() {
	if p.name == "right" {
		fmt.Println("Knock")
	} else {
		fmt.Println("Beep")
	}
}

func (p Door) Open() {
	if p.name == "Right" {
		fmt.Println("Beep")
	} else {
		fmt.Println("Knock")
	}
}

// define car
type Car struct {
	wheel     [4]Wheel
	rightDoor Door
	leftDoor  Door
}

// function for changes wheel of car
func (m *Car) ChangeWheel(position int, tire Tire) {
	if position < 0 || position >= 4 {
		fmt.Println("Posisi roda tidak valid")
		return
	}
	m.wheel[position] = Wheel{tire: tire}
	position++
	fmt.Printf("Roda di posisi %d diganti dengan %s\n", position, tire.TireType())
}

func main() {
	car := Car{
		rightDoor: Door{name: "right"},
		leftDoor:  Door{name: "left"},
	}

	car.ChangeWheel(0, RubberTire{})
	car.ChangeWheel(1, WoodenTire{})
	car.ChangeWheel(2, IronTire{})
	car.ChangeWheel(3, RubberTire{})

	fmt.Println("Pintu Kanan:")
	car.rightDoor.Knock()
	car.rightDoor.Open()

	fmt.Println("Pintu Kiri:")
	car.leftDoor.Knock()
	car.leftDoor.Open()
}
