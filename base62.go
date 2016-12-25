/*
Package base62 supports base62 encoding and decoding

base62 is usually used to encode short URLs.
*/
package base62

import (
	"math"
	"strings"
)

const (
	// Alphabet is the set of digits in base62
	Alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// Version is the current version
	Version = "0.2.0"
)

var base = uint64(len(Alphabet))

// Encode encodes num to base62 string.
func Encode(num uint64) string {
	if num == 0 {
		return "0"
	}

	arr := []uint8{}
	base := uint64(len(Alphabet))

	for num > 0 {
		rem := num % base
		num = num / base
		arr = append(arr, Alphabet[rem])
	}

	// Reverse the result array
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	return string(arr)
}

// Decode decodes base62 string to a number.
func Decode(b62 string) uint64 {
	size := len(b62)
	num := uint64(0)
	base := float64(len(Alphabet))

	for i, ch := range b62 {
		idx := i + 1
		loc := uint64(strings.IndexRune(Alphabet, ch))
		pow := uint64(math.Pow(base, float64(size-idx)))
		num += loc * pow
	}

	return num
}
