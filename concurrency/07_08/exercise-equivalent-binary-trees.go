package main

import (
	"fmt"
	"golang.org/x/tour/tree"
	"sync"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	fmt.Println(t.Value)
	Walk(t.Right, ch)
}

// Same determines whether the trees
// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1, c2 := make(chan int), make(chan int)

	go Walk(t1, c1)
	go Walk(t2, c2)

	for {
		v1, ok1 := <-c1
		v2, ok2 := <-c2
		if ok1 != ok2 || v1 != v2 {
			return false
		}
		if !ok1 && !ok2 {
			break
		}
	}
	return true
}

func main() {
	var wg sync.WaitGroup
	c := make(chan int)

	wg.Add(1)
	go func(c chan int) {
		defer wg.Done()
		Walk(tree.New(1), c)
	}(c)
	wg.Wait()

	// fmt.Println(Same(tree.New(1), tree.New(1)))
	// fmt.Println(Same(tree.New(2), tree.New(2)))
	// fmt.Println(Same(tree.New(3), tree.New(3)))

	// fmt.Println(Same(tree.New(1), tree.New(2)))
	// fmt.Println(Same(tree.New(1), tree.New(3)))
	// fmt.Println(Same(tree.New(1), tree.New(4)))

	// fmt.Println(Same(tree.New(2), tree.New(1)))
	// fmt.Println(Same(tree.New(3), tree.New(1)))
	// fmt.Println(Same(tree.New(4), tree.New(1)))
}
