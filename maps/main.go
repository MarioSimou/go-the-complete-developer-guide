package main

import "fmt"

func main() {
	// declaration
	// 1) with var
	// var colors map[string]string

	// 2) with make
	colors := make(map[string]string)
	colors["white"] = "#ffffff"
	colors["red"] = "#ff0000"
	colors["green"] = "#00ff00"

	// delete a property
	delete(colors, "white")

	print(colors)

}

func print(m map[string]string) {
	for k, v := range m {
		fmt.Println(k, v)
	}
}
