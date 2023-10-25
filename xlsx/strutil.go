package xlsx

import "strconv"

const (
	base10 = 10
	bit32  = 32
)

func parseInt32(val string) (int32, error) {
	n, err := strconv.ParseInt(val, base10, bit32)
	return int32(n), err
}
