/* base62 encoding

base62 is usually used to encode short URLs.
*/
package base62

const (
	alphabet = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Version = "0.1.0"
)

// Encode encoders num to base62 string.
func Encode(num uint64) string {
	if num == 0 {
		return "0"
	}

	arr := []uint8{}
	base := uint64(len(alphabet))

	for num > 0 {
		rem := num % base
		num = num / base
		arr = append(arr, alphabet[rem])
	}

	// Reverse the result array
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	return string(arr)
}
