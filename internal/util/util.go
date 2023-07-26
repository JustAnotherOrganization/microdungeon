package util

import "strconv"

const (
	base10 = 10
	b32    = 32
)

func Ptr[V any](v V) *V {
	return &v
}

func FormatID(id int32) string {
	return strconv.FormatInt(int64(id), base10)
}

func ParseID(id string) (int32, error) {
	i64, err := strconv.ParseInt(id, base10, b32)
	if err != nil {
		return 0, err
	}

	// This is safe because we already know it is 32
	return int32(i64), nil
}
