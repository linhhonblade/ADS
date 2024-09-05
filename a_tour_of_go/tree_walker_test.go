package main

import (
	"fmt"
	"golang.org/x/tour/tree"
	"testing"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	ch <- t.Value
	Walk(t.Left, ch)
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	check := make(map[int]bool)
	go func() {
		Walk(t1, ch1)
		close(ch1)
	}()
	for i := range ch1 {
		if _, ok := check[i]; !ok {
			check[i] = true
			fmt.Println("Tree 1: ", i)
		}
	}
	go func() {
		Walk(t2, ch2)
		close(ch2)
	}()
	for k := range ch2 {
		if _, ok := check[k]; !ok {
			return false
		} else {
			fmt.Println("Tree 2: ", k)
		}
	}
	return true
}

func TestSameTree(t *testing.T) {
	tree1 := tree.New(1)
	tree2 := tree.New(1)
	t.Run("Test Same Tree", func(t *testing.T) {
		res := Same(tree1, tree2)
		if !res {
			t.Errorf("Res should be True, got False\n")
		}
	})
}

func TestDiffTree(t *testing.T) {
	tree1 := tree.New(1)
	tree3 := tree.New(2)
	t.Run("Test Different Tree", func(t *testing.T) {
		res := Same(tree1, tree3)
		if res {
			t.Errorf("Res should be False, got True\n")
		}
	})
}
