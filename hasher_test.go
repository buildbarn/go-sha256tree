package sha256tree_test

import (
	"crypto/sha256"
	"math"
	"strconv"
	"testing"

	"github.com/buildbarn/go-sha256tree"
	"github.com/stretchr/testify/require"
)

// For any input of 1024 or less, the resulting hash should be identical
// to that of plain SHA-256.
func TestSmall(t *testing.T) {
	var input [1024]byte
	for i := 0; i < len(input); i++ {
		input[i] = byte(i % 251)
	}
	for i := 0; i <= len(input); i++ {
		expected := sha256.Sum256(input[:i])
		w := sha256tree.New(int64(i))
		w.Write(input[:i])
		require.Equal(t, expected[:], w.Sum(nil))
	}
}

func TestLarge(t *testing.T) {
	hasherSmall := sha256tree.New(0)
	hasherLarge := sha256tree.New(math.MaxInt64)

	for _, testVector := range []struct {
		inputSizeBytes int
		hash           []byte
	}{
		{1025, []byte{0x36, 0xc0, 0x99, 0x8b, 0x21, 0x83, 0x9e, 0xf7, 0x43, 0x00, 0xb9, 0xde, 0x47, 0xd9, 0x6d, 0x1f, 0x62, 0x32, 0x3d, 0xc8, 0x1f, 0x2b, 0x42, 0x31, 0xe9, 0x8c, 0xe7, 0x0c, 0xd6, 0xff, 0xe7, 0x50}},
		{2048, []byte{0xb5, 0x84, 0x99, 0x63, 0x86, 0xf0, 0x17, 0x93, 0x75, 0x1c, 0x5c, 0xf0, 0xc3, 0x95, 0x61, 0xf5, 0x1b, 0x7e, 0x99, 0x24, 0xb8, 0x18, 0x94, 0x3b, 0x3c, 0xb2, 0xf6, 0x92, 0x8c, 0xea, 0x0f, 0xa9}},
		{2049, []byte{0x73, 0x18, 0xd2, 0x02, 0x9b, 0x03, 0x92, 0xed, 0xf4, 0xcf, 0x10, 0x9e, 0xdb, 0x5a, 0x08, 0x6b, 0x4b, 0xda, 0xdb, 0xb7, 0x95, 0x0f, 0x71, 0x0a, 0x14, 0x83, 0xeb, 0x88, 0x1d, 0x9e, 0x5d, 0x44}},
		{3072, []byte{0xdf, 0xc6, 0x1c, 0x0a, 0x04, 0x1f, 0x79, 0xd5, 0x5d, 0x53, 0xbf, 0xe3, 0x1c, 0x6c, 0xda, 0x7d, 0xf7, 0x7f, 0xdc, 0x8e, 0x6f, 0xba, 0xc1, 0x14, 0x3d, 0x70, 0xb7, 0x14, 0x4f, 0xdf, 0x69, 0x37}},
		{3073, []byte{0x51, 0x7d, 0x20, 0xc0, 0xe5, 0x83, 0x5f, 0x06, 0x0a, 0x1b, 0xd6, 0x38, 0x8e, 0xd6, 0x85, 0x74, 0xf6, 0x34, 0x24, 0xbd, 0xac, 0x2a, 0x2c, 0x3a, 0x35, 0xa5, 0xc2, 0xef, 0x85, 0x9d, 0x8f, 0xe2}},
		{4096, []byte{0x2f, 0x72, 0xbb, 0x93, 0x88, 0x00, 0x12, 0x16, 0x8c, 0x02, 0x7f, 0x67, 0x81, 0x52, 0x7f, 0xf0, 0x81, 0x77, 0xc7, 0xc8, 0xdc, 0xcb, 0x44, 0x3f, 0x4d, 0x2c, 0x63, 0x89, 0xc1, 0x86, 0x63, 0x3d}},
		{4097, []byte{0xc3, 0xec, 0x94, 0x2c, 0x1b, 0x8f, 0x45, 0x80, 0x32, 0x0d, 0x3a, 0x06, 0xbc, 0xf4, 0xf8, 0xfe, 0x1f, 0x5d, 0xb2, 0xbe, 0x79, 0x7a, 0xb6, 0x70, 0x61, 0xea, 0x4c, 0x2a, 0x95, 0xf2, 0x08, 0xf2}},
		{5120, []byte{0xa7, 0x69, 0x24, 0xf6, 0x53, 0x5b, 0x4b, 0x47, 0x33, 0x77, 0xc2, 0x85, 0xec, 0x27, 0xac, 0xc8, 0x4c, 0xc5, 0x8e, 0x95, 0xab, 0x1e, 0x9e, 0x29, 0xb1, 0xbb, 0x6a, 0x4a, 0x3f, 0xb9, 0xd0, 0xb3}},
		{5121, []byte{0x98, 0xf9, 0x87, 0xc3, 0xe9, 0xfc, 0x05, 0x7a, 0x70, 0x87, 0x37, 0x15, 0xb6, 0x79, 0xb8, 0x9a, 0x66, 0x3d, 0x0d, 0xf8, 0x06, 0x85, 0x9b, 0x6c, 0xe7, 0x3f, 0x83, 0x79, 0xb0, 0x6a, 0x10, 0xff}},
		{6144, []byte{0x37, 0x2f, 0x98, 0x8a, 0xf4, 0x12, 0x04, 0x1b, 0x68, 0x0a, 0xb2, 0x36, 0xfe, 0xef, 0x45, 0x62, 0x63, 0x80, 0x06, 0x2b, 0xeb, 0x75, 0x14, 0xbb, 0xf9, 0x36, 0x07, 0xae, 0xdd, 0x28, 0xfc, 0x9a}},
		{6145, []byte{0x6d, 0xc4, 0xb7, 0x8e, 0xfd, 0x77, 0x04, 0x53, 0x41, 0x7b, 0x2f, 0xfd, 0xc7, 0x4b, 0x27, 0x05, 0x47, 0x93, 0xef, 0xe6, 0x12, 0x2e, 0xcd, 0x7e, 0xe0, 0x98, 0x67, 0x0e, 0xd7, 0xc4, 0x65, 0x1c}},
		{7168, []byte{0x43, 0x68, 0x63, 0x12, 0xc0, 0xca, 0xbc, 0xcf, 0x9d, 0x5a, 0xd5, 0x09, 0xef, 0xa0, 0x96, 0xe3, 0xd7, 0x43, 0xc6, 0x3c, 0x7a, 0x51, 0xf1, 0x22, 0x47, 0x3c, 0x57, 0x94, 0x9e, 0x4d, 0xd9, 0xa0}},
		{7169, []byte{0xad, 0x72, 0x92, 0x97, 0xab, 0x36, 0xcd, 0x09, 0x96, 0x65, 0xb2, 0x7c, 0x42, 0x47, 0x47, 0x4a, 0x55, 0x18, 0xe4, 0xcd, 0x0b, 0xe4, 0x43, 0xf5, 0xf3, 0x1d, 0x95, 0xed, 0xda, 0x08, 0x42, 0x9b}},
		{8192, []byte{0xfc, 0xfd, 0xde, 0x6f, 0xe5, 0x91, 0x78, 0xe1, 0x77, 0x08, 0xc5, 0xba, 0x64, 0x79, 0x19, 0xc3, 0xb1, 0x41, 0xa4, 0x4c, 0x9d, 0x19, 0x70, 0x78, 0x2e, 0x59, 0x7e, 0x14, 0x65, 0x26, 0x69, 0x32}},
		{8193, []byte{0x11, 0x3c, 0x6e, 0x3a, 0x24, 0x52, 0xf3, 0x88, 0xb6, 0xfa, 0xd1, 0x3d, 0xfa, 0xb6, 0x6e, 0xe0, 0xbf, 0xf5, 0x97, 0xa0, 0xa9, 0xa5, 0x17, 0xad, 0x8d, 0x01, 0x65, 0xf7, 0x19, 0x0b, 0x60, 0x3e}},
		{16384, []byte{0xa7, 0xa1, 0x01, 0x49, 0xa8, 0xcb, 0x00, 0xbe, 0x53, 0x70, 0x00, 0x56, 0x0e, 0xdb, 0x83, 0xb1, 0x96, 0x30, 0x6b, 0x78, 0x0b, 0x72, 0xfa, 0xd8, 0xaf, 0x21, 0x8f, 0x36, 0x9f, 0x75, 0xfc, 0x19}},
		{31744, []byte{0x2c, 0xdf, 0x76, 0x62, 0x63, 0x6c, 0x17, 0x3d, 0x4b, 0x23, 0x6f, 0x6e, 0xa0, 0x3b, 0xf8, 0x4c, 0x65, 0xe7, 0xf6, 0x48, 0x7b, 0x53, 0xb2, 0xa6, 0x1c, 0x42, 0x0e, 0x26, 0xcf, 0x8a, 0x98, 0xc7}},
		{102400, []byte{0x06, 0x68, 0xd6, 0x9e, 0x53, 0x31, 0x84, 0x0d, 0x2f, 0x18, 0x23, 0xd7, 0x17, 0xb7, 0xb3, 0xf5, 0xd1, 0xfd, 0xc8, 0xa0, 0x95, 0x04, 0xcd, 0xdb, 0x69, 0x2b, 0x87, 0xff, 0x83, 0xd5, 0x0e, 0x5f}},
	} {
		t.Run(strconv.FormatInt(int64(testVector.inputSizeBytes), 10), func(t *testing.T) {
			input := make([]byte, 0, testVector.inputSizeBytes)
			for i := 0; i < testVector.inputSizeBytes; i++ {
				input = append(input, byte(i%251))
			}

			// Test the hasher for small objects.
			hasherSmall.Reset()
			hasherSmall.Write(input)
			require.Equal(t, testVector.hash, hasherSmall.Sum(nil))

			// Test the hasher for large objects.
			hasherLarge.Reset()
			hasherLarge.Write(input)
			require.Equal(t, testVector.hash, hasherLarge.Sum(nil))
		})
	}
}
