package main

import (
	"fmt"
)

func main() {
	fmt.Println("Main func start!\n")

	var insert_begin int = 0
	var insert_end int = 1000000
	var update_begin int = 500000
	var update_end int = 1000000
	if update_begin < insert_begin {
		update_begin = insert_begin
	}

	if insert_end < update_end {
		update_end = insert_end
	}

	Init()
	Insert(insert_begin, insert_end)
	Query()
	Update(10, 100)
	//Delete()
	Fini()
	fmt.Println("Main func end!\n")
}
