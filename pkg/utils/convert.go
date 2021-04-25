package utils

import "strconv"

func Str2Int(str string) (int, error) {
	i64, err := strconv.ParseInt(str, 10, 32)
	return int(i64), err
}

func StrMust2Int(str string) int {
	v, err := Str2Int(str)
	if err != nil {
		panic(err)
	}
	return v
}

func Str2UInt(str string) (uint32, error) {
	u64, err := strconv.ParseUint(str, 10, 32)
	return uint32(u64), err
}

func StrMust2UInt(str string) uint32 {
	v, err := Str2UInt(str)
	if err != nil {
		panic(err)
	}
	return v
}
