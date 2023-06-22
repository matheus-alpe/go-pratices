package concurrency

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	if t != nil {
		Walk(t.Left, ch)
		ch <- t.Value
		Walk(t.Right, ch)
	}
}

func Walking(t *tree.Tree, ch chan int) {
	Walk(t, ch)
	defer close(ch)
}

func Same(t1, t2 *tree.Tree) bool {
	x := make(chan int)
	y := make(chan int)

	go Walking(t1, x)
	go Walking(t2, y)

	for {
		v1, ok1 := <-x
		v2, ok2 := <-y

		if ok1 != ok2 || v1 != v2 {
			return false
		}
		if !ok1 {
			break
		}
	}
	return true
}

func ExerciseEquivalentBinaryTreesExample() {
	fmt.Println("\nExercise Equivalent Binary Tree:")

	c := make(chan int)
	go Walking(tree.New(1), c)

	for i := range c {
		fmt.Println(i)
	}

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
