package egs

import "fmt"

type Collection []string
type MapFunc func(string) string
type MapFunc2 func(string, int) string

func (cars Collection) Map(fn MapFunc) Collection {
	mappedCars := make(Collection, 0, len(cars))
	for _, car := range cars {
		mappedCars = append(mappedCars, fn(car))
	}
	return mappedCars
}

func (cars Collection) Map2(fn MapFunc2) Collection {
	mappedCars := make(Collection, 0, len(cars))
	for _, car := range cars {
		mappedCars = append(mappedCars, fn(car, 2))
	}
	return mappedCars
}

func Upgrade() MapFunc {
	return func(car string) string {
		return fmt.Sprintf("%s %s", car, "LX")
	}
}

func TestLambda() {
	cars := &Collection{"Honda Accord", "Lexus IS 250"}

	fmt.Println("Upgrade() is not a Lambda Expression:")
	fmt.Printf("> cars.Map(Upgrade()): %+v\n\n", cars.Map(Upgrade()))

	fmt.Println("Anonymous function is not a Lambda Expression:")
	fmt.Printf("> cars.Map(func(...{...}): %+v\n\n", cars.Map2(func(car string, num int) string {
		return fmt.Sprintf("%s %s%d", car, "LX", num)
	}))
}
