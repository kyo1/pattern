package cmd

import (
	"strings"
)

const Pattern = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func Create(length int) string {
	var res []byte
	for i := 0; i < len(Pattern); i++ {
		for j := 0; j < len(Pattern); j++ {
			for k := 0; k < len(Pattern); k++ {
				if i == k {
					continue
				}
				if len(res) < length {
					res = append(res, Pattern[i])
				}
				if len(res) < length {
					res = append(res, Pattern[j])
				}
				if len(res) < length {
					res = append(res, Pattern[k])
				}
			}
		}
	}
	return string(res)
}

func Offset(needle string) int {
	haystack := Create(234484)
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
		res = Reverse(res)
	}
	return res
}

func Reverse(s string) string {
	res := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}
