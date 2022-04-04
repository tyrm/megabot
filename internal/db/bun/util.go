package bun

import (
	"encoding/binary"
)

func offset(i, c int) int { return i * c }

func int2bytes(i int64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(i))
	return b
}

func bytes2int(b []byte) int64 {
	return int64(binary.LittleEndian.Uint64(b))
}
