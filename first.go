package main
import "fmt"

func add(a, b int) int {
	return a + b
}

func diff (a, b int) int {
	return a - b
}

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Second, World!")
	fmt.Println("third, World!")
	fmt.Println("fourth, World!")

	fmt.Println(add(1, 2))
	fmt.Println(diff(1, 2))
	fmt.Println(diff(2, 4))
	fmt.Println(add(2, 4))
	fmt.Println(diff(3, 5))
	fmt.Println(diff(4, 5))
	fmt.Println(diff(5, 5))
}