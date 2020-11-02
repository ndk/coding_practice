package exercise

import "sort"

// https://www.coursera.org/learn/algorithms-part1/lecture/aJdPT/ordered-operations-in-bsts

type value struct {
	index   int
	number  int
	canOdd  bool
	canEven bool
}

type node struct {
	value
	left  *node
	right *node
}

func put(root *node, v value) *node {
	if root == nil {
		return &node{value: v}
	}
	switch {
	case v.number < root.number:
		root.left = put(root.left, v)
	case root.number < v.number:
		root.right = put(root.right, v)
	default:
		root.value = v
	}
	return root
}

func floor(root *node, n int) *node {
	if root == nil {
		return nil
	}
	switch {
	case n < root.number:
		return floor(root.left, n)
	case root.number < n:
		if f := floor(root.right, n); f != nil {
			return f
		}
	}
	return root
}

func ceil(root *node, n int) *node {
	if root == nil {
		return nil
	}
	switch {
	case n < root.number:
		if f := ceil(root.left, n); f != nil {
			return f
		}
		return root
	case root.number < n:
		return ceil(root.right, n)
	}
	return root
}

func oddEvenJumps(a []int) int {
	i := len(a) - 1
	visited := put(nil, value{i, a[i], true, true})

	result := 1
	for i--; i >= 0; i-- {
		v := value{index: i, number: a[i]}
		if ceil := ceil(visited, v.number); ceil != nil && ceil.canEven {
			v.canOdd = true
			result++
		}
		if floor := floor(visited, v.number); floor != nil && floor.canOdd {
			v.canEven = true
		}
		put(visited, v)
	}

	return result
}

///////////////////////////////////////////////////////////////////////////////
//	The fastest solution on Leetcode at the moment
func oddEvenJumps2(A []int) int {
	size := len(A)

	indexs := make([]int, size)
	for i := range indexs {
		indexs[i] = i
	}

	sort.Slice(indexs, func(i int, j int) bool {
		if A[indexs[i]] == A[indexs[j]] {
			return indexs[i] < indexs[j]
		}
		return A[indexs[i]] < A[indexs[j]]
	})

	nextHigher := nextIndex(indexs)

	ascToDes(A, indexs)

	nextLower := nextIndex(indexs)

	higher, lower := make([]int, size), make([]int, size)
	higher[size-1], lower[size-1] = 1, 1
	for i := size - 2; i >= 0; i-- {
		higher[i], lower[i] = lower[nextHigher[i]], higher[nextLower[i]]
	}
	return sum(higher)
}

func nextIndex(indexs []int) []int {
	size := len(indexs)
	res := make([]int, size)
	stack := make([]int, 0, size)
	for _, j := range indexs {
		for len(stack) > 0 && stack[len(stack)-1] < j {
			pop := stack[len(stack)-1]
			res[pop] = j
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, j)
	}
	return res
}

func ascToDes(A, indexs []int) {
	i, size := 0, len(A)
	for i+1 < size {
		if A[indexs[i]] != A[indexs[i+1]] {
			i++
			continue
		}
		a, j := A[indexs[i]], i+1
		for j+1 < size && A[indexs[j+1]] == a {
			j++
		}
		reverse(indexs, i, j)
		i = j + 1
	}
	reverse(indexs, 0, size-1)
}

func reverse(A []int, i, j int) {
	for i < j {
		A[i], A[j] = A[j], A[i]
		i++
		j--
	}
}

func sum(A []int) int {
	res := 0
	for _, a := range A {
		res += a
	}
	return res
}
