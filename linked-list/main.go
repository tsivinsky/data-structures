package main

import "fmt"

func main() {
	list := new(LinkedList)

	list.Insert(12)
	list.Insert(69)
	list.Insert(420)

	fmt.Println(list)

	node := list.Find(12)

	fmt.Println(node.Prev)

	list.Delete(node)
	list.Delete(list.Find(420))
	list.Delete(list.Find(13))

	fmt.Println(list)

}
