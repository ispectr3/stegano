package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kaique/stegano-go/stego"
)

const (
	colorReset = "\033[0m"
	colorGreen = "\033[32m"
	colorCyan  = "\033[36m"
	colorRed   = "\033[31m"
)

func printBanner() {
	banner := `
   _____ __                                  ______     
  / ___// /____  ____ _____ _____  ____     / ____/___  
  \__ \/ __/ _ \/ __ '/ __ '/ __ \/ __ \   / / __/ __ \ 
 ___/ / /_/  __/ /_/ / /_/ / / / / /_/ /  / /_/ / /_/ / 
/____/\__/\___/\__, /\__,_/_/ /_/\____/   \____/\____/  
              /____/                                    
`
	fmt.Printf("%s%s%s\n", colorCyan, banner, colorReset)
	fmt.Printf("%sStegano-Go - LSB Steganography Tool for PNG Images%s\n\n", colorGreen, colorReset)
}

func main() {
	mode := flag.String("mode", "", "Modo de operação: encode, decode, capacity")
	imagePath := flag.String("image", "", "Caminho da imagem PNG de entrada")
	outputPath := flag.String("output", "", "Caminho da imagem PNG de saída (para encode)")
	message := flag.String("message", "", "Mensagem de texto para ocultar")
	filePath := flag.String("file", "", "Caminho do arquivo contendo a mensagem (alternativa ao -message)")
	password := flag.String("password", "", "Senha opcional para criptografar a mensagem (AES-256)")
	
	flag.Usage = func() {
		printBanner()
		fmt.Fprintf(os.Stderr, "Uso:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if *mode == "" || *imagePath == "" {
		flag.Usage()
		os.Exit(1)
	}

	printBanner()

	switch *mode {
	case "encode":
		if *outputPath == "" {
			fmt.Printf("%sErro: Caminho de saída (-output) não especificado para o modo encode.%s\n", colorRed, colorReset)
			os.Exit(1)
		}
		
		var msgData []byte
		if *filePath != "" {
			data, err := os.ReadFile(*filePath)
			if err != nil {
				fmt.Printf("%sErro ao ler arquivo: %v%s\n", colorRed, err, colorReset)
				os.Exit(1)
			}
			msgData = data
		} else if *message != "" {
			msgData = []byte(*message)
		} else {
			fmt.Printf("%sErro: Nenhuma mensagem fornecida (-message ou -file).%s\n", colorRed, colorReset)
			os.Exit(1)
		}

		err := stego.Encode(*imagePath, *outputPath, msgData, *password)
		if err != nil {
			fmt.Printf("%sErro ao ocultar mensagem: %v%s\n", colorRed, err, colorReset)
			os.Exit(1)
		}
		fmt.Printf("%sMensagem ocultada com sucesso em %s!%s\n", colorGreen, *outputPath, colorReset)

	case "decode":
		msgData, err := stego.Decode(*imagePath, *password)
		if err != nil {
			fmt.Printf("%sErro ao extrair mensagem: %v%s\n", colorRed, err, colorReset)
			os.Exit(1)
		}
		fmt.Printf("%sMensagem extraída:%s\n%s\n", colorGreen, colorReset, string(msgData))

	case "capacity":
		capacity, err := stego.CalculateCapacity(*imagePath)
		if err != nil {
			fmt.Printf("%sErro ao calcular capacidade: %v%s\n", colorRed, err, colorReset)
			os.Exit(1)
		}
		fmt.Printf("%sCapacidade da imagem %s:%s\n", colorCyan, *imagePath, colorReset)
		fmt.Printf("Total de bytes disponíveis para ocultar: %d bytes (aprox. %.2f KB)\n", capacity, float64(capacity)/1024)

	default:
		fmt.Printf("%sModo inválido. Use: encode, decode, capacity.%s\n", colorRed, colorReset)
		flag.Usage()
		os.Exit(1)
	}
}
