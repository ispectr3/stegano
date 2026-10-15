# Stegano

**PNG LSB steganography CLI in Go with optional AES-256-GCM payload encryption**

[![CI](https://github.com/ispectr3/stegano/actions/workflows/ci.yml/badge.svg)](https://github.com/ispectr3/stegano/actions/workflows/ci.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Go: 1.21+](https://img.shields.io/badge/Go-1.21%2B-00ADD8.svg?logo=go)](https://golang.org)
[![Encryption: AES-256-GCM](https://img.shields.io/badge/Encryption-AES--256--GCM-green.svg)](stego/crypto.go)

Stegano embeds and extracts arbitrary bytes in PNG images by writing one bit into the least significant bit of each RGB channel. The command-line interface accepts either a text message or a file. When a password is supplied, the payload is encrypted before it is embedded.

> **Authorized use:** use this tool only with images and data you own or are explicitly authorized to handle. Steganography and encryption are dual-use capabilities; follow applicable law and organizational policy, and do not use this project to conceal or transfer data without permission.

## Confirmed features

- **Encode:** writes a length-prefixed payload to an output PNG using RGB LSBs.
- **Decode:** reads the payload from a PNG and prints the extracted bytes to standard output.
- **Text or file input:** `-message` and `-file` are alternatives for the `encode` mode.
- **Optional authenticated encryption:** `-password` uses AES-256-GCM. The key is the SHA-256 digest of the password and a random nonce is stored with the ciphertext.
- **Capacity inspection:** `-mode capacity` calculates the available payload size from the image dimensions.
- **Standard library only:** the module has no external Go dependencies and uses the standard image/PNG and cryptography packages.

## Requirements and installation

- Go 1.21 or newer.
- A PNG image for encoding, decoding, or capacity inspection.

```bash
git clone https://github.com/ispectr3/stegano.git
cd stegano
go build -o stegano main.go
```

## Usage

### Check image capacity

The command reports the capacity available after the 4-byte payload-length header:

```bash
./stegano -mode capacity -image cover.png
```

### Encode a text message

```bash
./stegano -mode encode \
  -image cover.png \
  -output stego.png \
  -message "Confidential Operations Plan" \
  -password "MySecretPassphrase"
```

### Encode a file

```bash
./stegano -mode encode \
  -image cover.png \
  -output stego.png \
  -file payload.bin
```

Add `-password` to the file example when authenticated encryption is required.

### Decode a message

Use the same password that was supplied during encoding, if any:

```bash
./stegano -mode decode \
  -image stego.png \
  -password "MySecretPassphrase"
```

The CLI writes decoded data to standard output; redirect it when a file is appropriate, for example `> extracted.txt`.

## Tests

The repository's CI runs the test suite with Go 1.21 and 1.22 and builds the CLI binary:

```bash
go test -v ./...
go build -v -o stegano main.go
```

The current unit tests cover AES-256-GCM round trips and rejection of an incorrect password. Image encode/decode integration tests are not currently included.

## Repository layout

```text
.
├── main.go                 # CLI flags and encode/decode/capacity modes
├── stego/
│   ├── encoder.go          # PNG loading and RGB LSB embedding
│   ├── decoder.go          # RGB LSB extraction and payload framing
│   ├── crypto.go           # AES-256-GCM encryption and decryption
│   ├── utils.go            # Capacity and bit/byte conversion helpers
│   └── crypto_test.go      # Encryption unit tests
├── samples/README.md       # Notes for adding PNG test images
├── .github/workflows/ci.yml
├── go.mod
└── LICENSE
```

## Limitations and implementation notes

- Only PNG input is supported. The output is also encoded as PNG; JPEG and other image formats are not handled.
- The payload starts with a 32-bit big-endian length and has no magic marker or version field. Decoding an ordinary PNG, or a PNG made with incompatible framing, can therefore fail with a corrupted or invalid length error.
- Capacity is based on three RGB channels and one bit per channel. The `capacity` mode reports `(width × height × 3 / 8) - 4` bytes. Encoding also stores the 4-byte header, and enabling encryption adds the AES-GCM nonce and authentication tag, so the maximum plaintext is smaller when a password is used.
- Passwords are converted directly to a 32-byte key with SHA-256; the implementation does not provide a salt or a slow password-based key-derivation function.
- The decode mode prints the extracted bytes as a string and has no separate output-file option.
- The project does not claim that hidden data is undetectable or that encryption provides anonymity.

## License

Stegano is open-source software licensed under the [MIT License](LICENSE).
