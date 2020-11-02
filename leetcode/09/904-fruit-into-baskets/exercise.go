package exercise

func totalFruit(tree []int) int {
	if len(tree) <= 2 {
		return len(tree)
	}

	longest := 0
	//
	// p                             *
	//      [3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4]
	// a              |  |     |  |
	//   fa           *  |     |  |
	//   la              |     *  |
	// b                 |        |
	//   fb              *        |
	//   lb                       *
	fa, la := 0, 0
	for ; la < len(tree) && tree[la] == tree[fa]; la++ {
	}
	if la == len(tree) {
		return len(tree)
	}
	fb, lb := la, la
	la--
	p := lb + 1
	for ; p < len(tree); p++ {
		if tree[p] == tree[fa] {
			la = p
			continue
		}
		if tree[p] == tree[fb] {
			lb = p
			continue
		}
		last := lb
		if last < la {
			last = la
		}
		if last-fa+1 > longest {
			longest = last - fa + 1
		}
		if lb < la {
			fa, fb, lb = lb+1, p, p
		} else {
			fa, la, fb, lb = la+1, lb, p, p
		}
	}
	if p-fa > longest {
		longest = p - fa
	}

	return longest
}

///////////////////////////////////////////////////////////////////////////////
// The fastest solution on Leetcode at the moment
func totalFruit2(tree []int) int {
	var res, cur, count_b, a, b int
	for _, c := range tree {
		if a == c || b == c {
			cur++
		} else {
			if cur > res {
				res = cur
			}
			cur = count_b + 1
		}
		if b == c {
			count_b++
		} else {
			count_b = 1
		}
		if b != c {
			a, b = b, c
		}
	}
	if cur > res {
		return cur
	}
	return res
}
