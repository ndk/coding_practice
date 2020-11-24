package exercise

func ShiftedArrSearch(shiftArr []int, num int) int {
	li, ri := 0, len(shiftArr)-1
	for li <= ri {
		mi := li + (ri-li)/2
		l, m, r := shiftArr[li], shiftArr[mi], shiftArr[ri]
		switch {
		case num == l:
			return li
		case num == m:
			return mi
		case num == r:
			return ri
		case (num < r && r < l && l < m) ||
			(m < num && num < r && r < l) ||
			(l < m && m < num && num < r) ||
			(r < l && l < m && m < num):
			li = mi
		case (num < m && m < r && r < l) ||
			(l < num && num < m && m < r) ||
			(r < l && l < num && num < m) ||
			(m < r && r < l && l < num):
			ri = mi
		default:
			li++
			ri--
		}
	}
	return -1
}
