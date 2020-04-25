package pattern

import (
	"strings"
)

const Pattern = "abcdefghijklmnopqrstuvwxyz"

// https://en.wikipedia.org/wiki/De_Bruijn_sequence
func DeBruijn(k string, n int) string {
	var sequence []int
	var a []int = make([]int, len(k)*n)

	var db func(int, int)
	db = func(t, p int) {
		if t > n {
			if n%p == 0 {
				sequence = append(sequence, a[1:p+1]...)
			}
		} else {
			a[t] = a[t-p]
			db(t+1, p)
			for j := a[t-p] + 1; j < len(k); j++ {
				a[t] = j
				db(t+1, t)
			}
		}
	}
	db(1, 1)

	var res []byte
	for _, s := range sequence {
		res = append(res, k[s])
	}
	return string(res)
}

func Create(length int) string {
	haystack := DeBruijn(Pattern, 4)
	return haystack[:length]
}

func Offset(needle string) int {
	haystack := DeBruijn(Pattern, 4)
	return strings.Index(haystack, needle)
}

func Hex2str(x uint64, bigendian bool) string {
	var b []byte
	for x > 0 {
		b = append(b, byte(x&uint64(0xff)))
		x >>= 8
	}
	res := string(b)
	if bigendian {
		res = reverse(res)
	}
	return res
}

func reverse(s string) string {
	res := []rune(s)
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		res[l], res[r] = res[r], res[l]
	}
	return string(res)
}
