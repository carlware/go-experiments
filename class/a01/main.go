package main

import (
	"fmt"
)

func Suma(c, d int) float32 {
	res := c + d
	return float32(res)
}

func main() {
	type gato struct {
		Nombre string
		Edad   uint8
	}

	gatos := []struct {
		Nombre string
		Fn     func(a, b int) float32
		Edad   uint8
	}{
		{
			Nombre: "blanco",
			Edad:   10,
			Fn:     Suma,
		},
		{
			Nombre: "negro",
			Edad:   15,
			Fn:     Suma,
		},
		{
			Nombre: "rojo",
			Edad:   1,
			Fn:     Suma,
		},
	}

	fmt.Println(gatos[1].Fn(2, 3))
	fmt.Println(gatos)
}
