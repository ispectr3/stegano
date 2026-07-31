package stego

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

// Encode hides a message inside a PNG image using LSB and saves it to outputPath
func Encode(inputPath, outputPath string, message []byte, password string) error {
	// 1. Load image
	file, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("failed to open image: %w", err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		return fmt.Errorf("failed to decode PNG: %w", err)
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	// 2. Encrypt if password is provided
	if password != "" {
		message, err = Encrypt(message, password)
		if err != nil {
			return fmt.Errorf("encryption failed: %w", err)
		}
	}

	// 3. Prepare payload: [Length (4 bytes)] + [Message (N bytes)]
	payloadLength := uint32(len(message))
	buf := new(bytes.Buffer)
	err = binary.Write(buf, binary.BigEndian, payloadLength)
	if err != nil {
		return err
	}
	buf.Write(message)
	payload := buf.Bytes()

	// Check capacity
	capacity := (width * height * 3) / 8 // 3 channels (RGB)
	if len(payload) > capacity {
		return fmt.Errorf("message too large: payload is %d bytes, max capacity is %d bytes", len(payload), capacity)
	}

	// 4. Create new RGBA image for editing
	outImg := image.NewRGBA(bounds)
	
	payloadBits := bitsFromBytes(payload)
	bitIndex := 0
	totalBits := len(payloadBits)

	// 5. Embed payload into LSB
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			originalColor := img.At(x, y)
			r, g, b, a := originalColor.RGBA()
			
			// Convert to 8-bit color
			r8 := uint8(r >> 8)
			g8 := uint8(g >> 8)
			b8 := uint8(b >> 8)
			a8 := uint8(a >> 8)

			if bitIndex < totalBits {
				r8 = (r8 & 0xFE) | payloadBits[bitIndex]
				bitIndex++
			}
			if bitIndex < totalBits {
				g8 = (g8 & 0xFE) | payloadBits[bitIndex]
				bitIndex++
			}
			if bitIndex < totalBits {
				b8 = (b8 & 0xFE) | payloadBits[bitIndex]
				bitIndex++
			}

			outImg.Set(x, y, color.RGBA{R: r8, G: g8, B: b8, A: a8})
		}
	}

	// 6. Save image
	out, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer out.Close()

	err = png.Encode(out, outImg)
	if err != nil {
		return fmt.Errorf("failed to encode output PNG: %w", err)
	}

	return nil
}
