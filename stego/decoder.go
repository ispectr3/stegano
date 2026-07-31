package stego

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"image/png"
	"os"
)

// Decode extracts a hidden message from a PNG image
func Decode(inputPath, password string) ([]byte, error) {
	// 1. Load image
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open image: %w", err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("failed to decode PNG: %w", err)
	}

	bounds := img.Bounds()

	var extractedBits []uint8
	
	// Max length to read initially (4 bytes for length = 32 bits)
	headerBits := 32
	
	// Read all bits to find the message
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			
			extractedBits = append(extractedBits, uint8(r>>8)&1)
			extractedBits = append(extractedBits, uint8(g>>8)&1)
			extractedBits = append(extractedBits, uint8(b>>8)&1)
		}
	}

	if len(extractedBits) < headerBits {
		return nil, fmt.Errorf("image too small to contain a message")
	}

	// Extract the 4-byte length
	lengthBytes := bytesFromBits(extractedBits[:headerBits])
	var payloadLength uint32
	err = binary.Read(bytes.NewReader(lengthBytes), binary.BigEndian, &payloadLength)
	if err != nil {
		return nil, err
	}

	totalBitsNeeded := headerBits + int(payloadLength)*8
	if totalBitsNeeded > len(extractedBits) {
		return nil, fmt.Errorf("corrupted data or no message found (length read: %d bytes)", payloadLength)
	}
	
	// Sanity check, if payload length is absurdly large
	if payloadLength > uint32(len(extractedBits)/8) {
		return nil, fmt.Errorf("invalid message length extracted")
	}

	// Extract message bytes
	messageBits := extractedBits[headerBits:totalBitsNeeded]
	message := bytesFromBits(messageBits)

	// Decrypt if password is provided
	if password != "" {
		message, err = Decrypt(message, password)
		if err != nil {
			return nil, fmt.Errorf("decryption failed (wrong password?): %w", err)
		}
	}

	return message, nil
}
