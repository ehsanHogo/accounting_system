package casting

import (
	"strconv"
)

func StringToUint(s string) uint {
	u, _ := strconv.ParseUint(s, 10, 64)

	return uint(u)
}

func UintToString(u uint) string {
	return strconv.FormatUint(uint64(u), 10)
}
