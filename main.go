package main

import (
	"fmt"

	"github.com/salmanrf/linked_list/impl"
)

func main() {
	ll := impl.New[string]()

	ll.Append("Wombat")
	ll.Append("Capybara")
	ll.Append("Baboon")

	ll.Insert("Manul", 0)

	ll.Insert("Manatee", 200)
	ll.Insert("Manatee", 3)
	
	ll.Traverse(func (val string, index int) {
		fmt.Printf("Index: %d, Value: %s \n", index, val)
	})

	ll.Delete((0))
	
	fmt.Println("========================")
	
	ll.Traverse(func (val string, index int) {
		fmt.Printf("Index: %d, Value: %s \n", index, val)
	})
}