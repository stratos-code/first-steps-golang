package main

import "fmt"

// generics
func main() {
	// functionExample()
	exampleStructs()
}

func functionExample() {
	var intSlice = []int{1, 2, 3, 4, 5}
	fmt.Println(sumSlice[int](intSlice))

	var float32Slice = []float32{1.1, 2.2, 3.3, 4.4, 5.5}
	fmt.Println(sumSlice[float32](float32Slice))

	var intSlice2 = []int{}
	fmt.Println(isEmpty[int](intSlice2))

	var float32Slice2 = []float32{1.1, 2.2, 3.3, 4.4, 5.5}
	fmt.Println(isEmpty[float32](float32Slice2))

}

func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

func isEmpty[T any](slice []T) bool {
	return len(slice) == 0
}

// with structs
type gasEngine struct {
	gallons float32
	mpg     float32
}

type electricEngine struct {
	kwh   float32
	mpkwh float32
}

type car[T gasEngine | electricEngine] struct {
	carMake  string
	carModel string
	engine   T
}

func exampleStructs() {
	var gasCar = car[gasEngine]{
		carMake:  "Toyota",
		carModel: "Corolla",
		engine: gasEngine{
			gallons: 10,
			mpg:     30,
		},
	}

	var electricCar = car[electricEngine]{
		carMake:  "Tesla",
		carModel: "Model 3",
		engine: electricEngine{
			kwh:   50,
			mpkwh: 4,
		},
	}

	fmt.Printf("Gas car: %v, Electric car: %v\n", gasCar, electricCar)
}
