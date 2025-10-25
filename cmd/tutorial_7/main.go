package main

import "fmt"

func main() {
	var p *int32 = new(int32)
	var i int32
	fmt.Printf("The value of p points to is: %v\nthe address is: %v", *p, p)
	fmt.Printf("\nThe value of i is: %v\nthe address is: %v", i, &i)

	*p = 10
	fmt.Printf("\nThe value of p points to is: %v\nthe address is: %v", *p, p)

	p = &i
	fmt.Printf("\nThe value of p points to is: %v\nthe address is: %v", *p, p)

	*p = 20
	fmt.Printf("\nThe value of p points to is: %v\nthe address is: %v", *p, p)
	fmt.Printf("\nThe value of i is: %v\n", i)

	var k int32 = 2
	i = k
	fmt.Printf("\nAddres of i: %v\nAddres of k: %v\n", &i, &k)

	var slice = []int{1, 2, 3}
	var sliceCopy = slice
	sliceCopy[2] = 4
	fmt.Println(slice, &slice[0])         // prints [1 2 4] 0xc0000140b0
	fmt.Println(sliceCopy, &sliceCopy[0]) // prints [1 2 4] 0xc0000140b0

	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("\nThe memory location of the thing1 array is %p", &thing1)
	var thing2 = square(thing1)
	fmt.Printf("\nThe memory location of the thing2 array is %p\n", &thing2)

	fmt.Printf("\nThe original thing1 array is: %v", thing1)
	fmt.Printf("\nThe result is: %v\n", thing2)

	var thing3 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("\nThe memory location of the thing3 array is %p", &thing3)
	var result = squarePtr(&thing3)
	fmt.Printf("\nThe result is: %v", result)
	fmt.Printf("\nThe thing3 is: %v", thing3)
	fmt.Printf("\nThe memory location of the result array is %p\n", &result)

}

// Function that receives an array and returns a new array with squared values
func square(thing2 [5]float64) [5]float64 {
	fmt.Printf("\nThe memory location of the thing2 array is %p", &thing2)

	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return thing2
}

// Function that receives a pointer to an array and modifies the original array
func squarePtr(arr *[5]float64) [5]float64 {
	fmt.Printf("\nThe memory location of the arr array is %p", arr)
	for i := range arr {
		arr[i] = arr[i] * arr[i]
	}

	return *arr
}
