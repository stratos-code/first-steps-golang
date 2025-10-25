package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
	owner
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
	owner
}

type owner struct {
	name string
}

func (e gasEngine) milesLeft() uint8 {
	return e.mpg * e.gallons
}

func (e electricEngine) milesLeft() uint8 {
	return e.mpkwh * e.kwh
}

// func (e gasEngine) milesLeftFloat() float64 {
// 	return float64(e.mpg) * float64(e.gallons)
// }

// func (e electricEngine) milesLeftFloat() float64 {
// 	return float64(e.mpkwh) * float64(e.kwh)
// }

type engine interface {
	milesLeft() uint8
	// milesLeftFloat() float64
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it there!")
	} else {
		fmt.Println("Need to fuel up first!")
	}
}

func main() {
	var myEngine gasEngine = gasEngine{25, 15, owner{"Alex"}}

	myEngine.mpg = 20

	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.name)
	fmt.Println(myEngine)

	var myEngine2 = struct {
		mpg     uint8
		gallons uint8
		owner
	}{30, 10, owner{"Maria"}} // Struct anónima

	fmt.Println(myEngine2)

	fmt.Printf("Miles left: %v\n", myEngine.milesLeft())

	canMakeIt(myEngine, 80)
	canMakeIt(electricEngine{3, 30, owner{"Eve"}}, 80)
}

// Forzar verificación en tiempo de compilación
// de que los tipos cumplen con la interfaz
var _ engine = (*gasEngine)(nil)
var _ engine = (*electricEngine)(nil)
