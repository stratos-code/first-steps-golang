package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var myString = "résumé 🥳"
	var indexed = myString[0]
	fmt.Printf("%v, %T\n", indexed, indexed)

	var runeSlice = []rune(myString)
	fmt.Printf("%v, %T\n", runeSlice, runeSlice)
	fmt.Printf("%d, %d\n", len(myString), len(runeSlice))

	fmt.Println("Iterating string")
	for i, v := range myString {
		fmt.Println(i, v)
	}
	fmt.Println("Iterating rune slice")
	for i, v := range runeSlice {
		fmt.Println(i, v)
	}

	var myRune = 'a'
	fmt.Printf("\nmyRune = %v\n", myRune)

	var strSlice = []string{"h", "e", "l", "l", "o", " ", "😊"}
	var catStr = ""

	var t0 = time.Now()
	for i := range strSlice {
		catStr += strSlice[i]
	}
	var elapsed = time.Since(t0)
	fmt.Printf("Time concatenating with += : %v\n", elapsed)
	fmt.Println(catStr)

	catStr = ""
	var strBuilder strings.Builder
	t0 = time.Now()
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	catStr = strBuilder.String()
	elapsed = time.Since(t0)
	fmt.Printf("Time concatenating with strings.Builder: %v\n", elapsed)
	fmt.Println(catStr)

}
