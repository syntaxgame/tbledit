package utils

import (
	"math"
)

func IntToBytes(dec uint64, bytes byte, reverse bool) []byte {

	arr := make([]byte, 8)
	if dec == 0 {
		return arr[:bytes]
	}

	if bytes < 8 {
		dec = dec % (1 << (bytes * 8))
	}

	for i := 7; i >= 0; i-- {
		power := uint64(math.Pow(256, float64(i)))
		arr[i] = byte(dec / power)
		if arr[i] > 0 {
			dec -= uint64(arr[i]) * power
		}
	}

	if reverse {
		return arr[:bytes]
	}

	return reverseBytes(arr)[8-bytes : 8]
}

func BytesToInt(arr []byte, reverse bool) int64 {

	var total int64
	_arr := make([]byte, len(arr))
	copy(_arr, arr)

	if reverse {
		_arr = reverseBytes(_arr)
	}

	for i := 0; i < len(_arr); i++ {
		power := int64(math.Pow(256, float64(len(_arr)-(i+1))))
		total += int64(_arr[i]) * power
	}

	return total
}

func BytesToFloat(arr []byte, reverse bool) float32 {
	return math.Float32frombits(uint32(BytesToInt(arr, reverse)))
}

func FloatToBytes(dec float32, bytes byte, reverse bool) []byte {
	return IntToBytes(uint64(math.Float32bits(dec)), bytes, reverse)
}

func reverseBytes(input []byte) []byte {
	if len(input) == 0 {
		return input
	}

	return append(reverseBytes(input[1:]), input[0])
}
