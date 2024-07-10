package main

func secondAdd(a, b int)int {
	return a + b
}

func secondAdd2(a, b int)int {
	return a + b + 2
}

func main() {
	secondAdd(1, 2)
	secondAdd(1, 2)
	secondAdd(1, 2)
	secondAdd(1, 2)
	secondAdd(1, 2)
	secondAdd(1, 2)

	secondAdd2(1, 2)
}
