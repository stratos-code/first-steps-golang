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
	fmt.Println(result)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1 / intNum2)                   // División entera
	fmt.Println(float64(intNum1) / float64(intNum2)) // División float
	fmt.Println(intNum1 % intNum2)                   // Módulo

	var myString string = "Hello\nWorld"
	fmt.Println(myString)
	myString = `Hello
World`
	fmt.Println(myString)

	var booleano bool = true
	fmt.Println(booleano)

	fmt.Println(len(myString))       // tamaño en bytes no caracteres (11 caracteres ASCII -> 11 bytes)
	fmt.Println(len(myString + "🫥")) // 12 caracteres, pero 15 bytes ya que el emoji ocupa 4 bytes

	fmt.Println(utf8.RuneCountInString(myString + "🫥")) // número de caracteres

	var myRune rune = 'a' // un rune es un int32 que representa un caracter Unicode
	fmt.Println(myRune)

	var myBoolean bool = false
	fmt.Println(myBoolean)

	// Default values
	fmt.Println("\nDefault values")
	var defaultInt int
	var defaultFloat float64
	var defaultString string
	var defaultBool bool
	fmt.Println(defaultInt)    // 0
	fmt.Println(defaultFloat)  // 0
	fmt.Println(defaultString) // ""
	fmt.Println(defaultBool)   // false

	var1, var2, var3 := 1, true, "Hello"
	fmt.Println(var1, var2, var3)

	const myConst string = "const value"
	// myConst = "new value" // Error: cannot assign to myConst
}
