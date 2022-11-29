package utils

import "strconv"

func Uint64ToString(i uint64) string {
	return strconv.FormatUint(i, 10)
}

func StringToUint64(s string) (uint64, error) {
	return strconv.ParseUint(s, 10, 64)
}

func StringToFloat32(s string) (float32, error) {
	temp, err := strconv.ParseFloat(s, 32)
	return float32(temp), err
}

func StringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}
