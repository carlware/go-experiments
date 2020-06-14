package egs

import "fmt"

type Car struct {
	Make  string
	Model string
}

func (c Car) Tires() int { return 4 }
func (c Car) PrintInfo() {
	fmt.Printf("%v has %d tires\n", c, c.Tires())
}

type CarWithSpare struct {
	Car
}

func (o CarWithSpare) Tires() int { return 5 }

func (c CarWithSpare) PrintInfo() {
	fmt.Printf("%v has %d tires\n", c, c.Tires())
}

func TestCar() {
	accord := Car{"Honda", "Accord"}
	accord.PrintInfo()
	highlander := CarWithSpare{Car{"Toyota", "Highlander"}}
	highlander.PrintInfo()
	fmt.Printf("%v has %d tires\n", highlander.Car, highlander.Tires())
	accord.PrintInfo()
}
