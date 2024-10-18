package main

import "fmt"

func main() {
	arr := NewArray()

	fmt.Println("size:", arr.Size())
	fmt.Println("Appening....")
	arr = arr.Append("2")
	arr = arr.Append("3")
	arr = arr.Append("4")
	arr = arr.Append("5")
	arr = arr.Append("6")

	fmt.Println(arr.ToString())
	fmt.Println("size:", arr.Size())
	fmt.Println("Accessing index 1")
	item, err := arr.At(1)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Array at %v is %v\n", 1, item)
	}

	arr = arr.Clear()
	arr.ToString()
}
