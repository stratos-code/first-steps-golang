package main

import "fmt"

func main() {
	var intArr [3]int32 = [3]int32{1, 2, 3}
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])
	fmt.Println(intArr[1:])
	fmt.Println(intArr[:2])

	intArr[1] = 20
	fmt.Println(intArr)

	fmt.Println(len(intArr)) // length is fixed, always 3

	// address of the array
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	var intArr2 = [...]int32{4, 5, 6, 7, 8} // ellipsis infers the length
	fmt.Println(intArr2)
	fmt.Println(len(intArr2))

	// slices
	/*
		Go slices are dynamically-sized, flexible views into the elements of an array.
		Unlike arrays, slices do not have a fixed length and can grow or shrink as needed.
		A slice is defined by a pointer to an array, a length, and a capacity.
		Slices are created by slicing arrays or other slices, or by using the built-in make function.
		Appending to a slice may allocate a new underlying array if the capacity is exceeded.
		Slices are the idiomatic way to work with sequences in Go.
	*/
	var intSlice []int32 = []int32{10, 20, 30}
	fmt.Println(intSlice)
	fmt.Printf("Length: %d, Capacity: %d\n", len(intSlice), cap(intSlice))
	fmt.Println(&intSlice[0])
	intSlice = append(intSlice, 40)
	fmt.Println(intSlice)
	fmt.Printf("Length: %d, Capacity: %d\n", len(intSlice), cap(intSlice)) // len 4, cap 6
	fmt.Println(&intSlice[0])
	//fmt.Println(intSlice[4]) // despite capacity is 6, we can't access index 5

	// spread operator
	intSlice2 := []int32{50, 60, 70}
	intSlice = append(intSlice, intSlice2...) // append all elements of intSlice2
	fmt.Println(intSlice)
	fmt.Printf("Length: %d, Capacity: %d\n", len(intSlice), cap(intSlice)) // len 7, cap 12
	fmt.Println(&intSlice[0])

	// make function
	intSlice3 := make([]int32, 5, 10) // length 5, capacity 10
	fmt.Println(intSlice3)            // [0 0 0 0 0]
	fmt.Printf("Length: %d, Capacity: %d\n", len(intSlice3), cap(intSlice3))
	intSlice3[0] = 100
	intSlice3[1] = 200
	fmt.Println(intSlice3)

	// -------------------- Maps ------------------------------

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 23, "Eve": 45}
	fmt.Println(myMap2["Adam"])
	fmt.Println(myMap2["Jason"]) // zero value 0

	var age, ok = myMap2["Jason"]
	if ok {
		fmt.Println("Jason's age is", age)
	} else {
		fmt.Println("Jason's age not found")
	}

	fmt.Println(myMap2)
	delete(myMap2, "Eve")
	fmt.Println(myMap2)

	myMap2["Frank"] = 33

	// iterating a map
	for name := range myMap2 {
		fmt.Printf("Name: %v\n", name)
	}

	// iterating a map
	for name, age := range myMap2 {
		fmt.Printf("Name: %v, Age: %v\n", name, age)
	}

	// iterating an array
	for i, v := range intArr {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}

	// while loop
	var i int = 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	i = 0
	for {
		if i >= 10 {
			break
		}
		fmt.Println(i)
		i++
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
