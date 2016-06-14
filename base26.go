package base26

import "math/big"

var (
	base26 = [26]byte{
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H',
		'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P',
		'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X',
		'Y', 'Z'}

	index = map[byte]int{
		'A': 0, 'B': 1, 'C': 2, 'D': 3, 'E': 4, 'F': 5,
		'G': 6, 'H': 7, 'I': 8, 'J': 9, 'K': 10, 'L': 11,
		'M': 12, 'N': 13, 'O': 14, 'P': 15, 'Q': 16,
		'R': 17, 'S': 18, 'T': 19, 'U': 20, 'V': 21,
		'W': 22, 'X': 23, 'Y': 24, 'Z': 25,
		'a': 0, 'b': 1, 'c': 2, 'd': 3, 'e': 4, 'f': 5,
		'g': 6, 'h': 7, 'i': 8, 'j': 9, 'k': 10, 'l': 11,
		'm': 12, 'n': 13, 'o': 14, 'p': 15, 'q': 16,
		'r': 17, 's': 18, 't': 19, 'u': 20, 'v': 21,
		'w': 22, 'x': 23, 'y': 24, 'z': 25,
	}
)

// Encode encodes a uint64 value to string in base26 format
func Encode(value uint64) string {
	var res [16]byte
	var i int
	for i = len(res) - 1; value != 0; i-- {
		res[i] = base26[value%26]
		value /= 26
	}
	return string(res[i+1:])
}

// Decode decodes a base26-encoded string back to uint64
func Decode(s string) uint64 {
	res := uint64(0)
	l := len(s) - 1
	b26 := big.NewInt(26)
	bidx := big.NewInt(0)
	bpow := big.NewInt(0)
	for idx := range s {
		c := s[l-idx]
		byteOffset := index[c]
		bidx.SetUint64(uint64(idx))
		res += uint64(byteOffset) * bpow.Exp(b26, bidx, nil).Uint64()
	}
	return res
}
