package stego

import (
	"image/png"
	"os"
)

// CalculateCapacity calculates the maximum number of bytes that can be hidden in an image
func CalculateCapacity(inputPath string) (int, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	img, err := png.DecodeConfig(file)
	if err != nil {
		return 0, err
	}

	// 3 channels (RGB) * Width * Height / 8 bits per byte
	capacity := (img.Width * img.Height * 3) / 8
	// Subtract 4 bytes for length header
	return capacity - 4, nil
}

// bitsFromBytes converts a byte slice into a slice of bits (0 or 1)
func bitsFromBytes(data []byte) []uint8 {
	bits := make([]uint8, len(data)*8)
	for i, b := range data {
		for j := 0; j < 8; j++ {
			bits[i*8+j] = uint8((b >> (7 - j)) & 1)
		}
	}
	return bits
}

// bytesFromBits converts a slice of bits (0 or 1) back into a byte slice
func bytesFromBits(bits []uint8) []byte {
	bytes := make([]byte, len(bits)/8)
	for i := 0; i < len(bytes); i++ {
		var b uint8
		for j := 0; j < 8; j++ {
			b |= (bits[i*8+j] << (7 - j))
		}
		bytes[i] = b
	}
	return bytes
}
