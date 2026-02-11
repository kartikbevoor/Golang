package main

import (
	"fmt"
	"os"
)

func recursion() {

}

// factorial using recursion
func factorial(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// tail recursion: when recursive call is the last operation in the function
func factorial2(n, acc int) int {
	if n == 0 {
		return acc
	}
	return factorial2(n-1, n*acc)
}

// Go does NOT optimize tail recursion
// (No Tail Call Optimization - TCO)

func sumSlice(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return nums[0] + sumSlice(nums[1:])
}

// Binary tree
type Node struct {
	value int
	left  *Node
	right *Node
}

// inorder traversal
func inOrder(root *Node) {
	if root == nil {
		return
	}
	inOrder(root.left)
	fmt.Println(root.value)
	inOrder(root.right)
}

// post order traversal
func postOrder(root *Node) {
	if root == nil {
		return
	}

	postOrder(root.left)
	postOrder(root.right)
	fmt.Println(root.value)
}

// pre order traversal
func preOrder(root *Node) {
	if root == nil {
		return
	}

	fmt.Println(root.value)
	preOrder(root.left)
	preOrder(root.right)
}

// recursion for file transfer
func walk(path string) {
	files, _ := os.ReadDir(path)

	for _, file := range files {
		fullPath := path + "/" + file.Name()
		if file.IsDir() {
			walk(fullPath)
		} else {
			fmt.Println(fullPath)
		}
	}
}

// Memoization (Optimizing Recursion): fixing fibonachi using maps
var memo = make(map[int]int)

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	if val, ok := memo[n]; ok {
		return val
	}

	memo[n] = fibonacci(n-1) + fibonacci(n-2)
	return memo[n]
}
