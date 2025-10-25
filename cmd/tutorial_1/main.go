package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int = 32767
	intNum = intNum + 1
	fmt.Println(intNum)

	var floatNum float64 = 12345678.9
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println("resultado", result)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println("divison entera", intNum1/intNum2)                  // División entera
	fmt.Println("divison float", float64(intNum1)/float64(intNum2)) // División float
	fmt.Printf("Suma floats %.20f\n", 0.1+0.2)                      // Suma floats
	fmt.Println("modulo", intNum1%intNum2)                          // Módulo

	var myString string = "Hello\nWorld"
	fmt.Println(myString)
	myString = `Hello
World`
	fmt.Println(myString)

	var booleano bool = true
	fmt.Println(booleano)

	fmt.Println("tamaño en bytes", len(myString))     // tamaño en bytes no caracteres (11 caracteres ASCII -> 11 bytes)
	fmt.Println("tamaño en bytes", len(myString+"🫥")) // 12 caracteres, pero 15 bytes ya que el emoji ocupa 4 bytes

	fmt.Println("número de caracteres", utf8.RuneCountInString(myString+"🫥")) // número de caracteres

	var myRune rune = 'a' // un rune es un int32 que representa un caracter Unicode
	fmt.Println("myRune", myRune)

	var myBoolean bool = false
	fmt.Println("myBoolean", myBoolean)

	// Default values
	fmt.Println("\nDefault values")
	var defaultInt int
	var defaultFloat float64
	var defaultString string
	var defaultBool bool
	fmt.Println("defaultInt", defaultInt)       // 0
	fmt.Println("defaultFloat", defaultFloat)   // 0
	fmt.Println("defaultString", defaultString) // ""
	fmt.Println("defaultBool", defaultBool)     // false

	// Short variable declaration
	var1, var2, var3 := 1, true, "Hello"
	fmt.Println("Asignación múltiple", var1, var2, var3)

	const myConst string = "const value"
	// myConst = "new value" // Error: cannot assign to myConst
}
