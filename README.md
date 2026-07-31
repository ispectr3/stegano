# 🛡️ Stegano-Go

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Security](https://img.shields.io/badge/security-steganography-red.svg?style=for-the-badge)
![License](https://img.shields.io/badge/license-MIT-blue.svg?style=for-the-badge)

Uma ferramenta de linha de comando desenvolvida em Go para ocultar e extrair mensagens secretas em imagens PNG utilizando a técnica de Esteganografia LSB (Least Significant Bit).

---

## 🚀 Funcionalidades

- **Ocultar Textos e Arquivos:** Esconda mensagens ou conteúdos de arquivos dentro de imagens PNG aparentemente normais.
- **Extração Segura:** Recupere a mensagem escondida da imagem processada.
- **Técnica LSB:** Utiliza o Bit Menos Significativo para alterar levemente os pixels da imagem (RGB), tornando a alteração imperceptível a olho nu.
- **Criptografia AES-256:** (Opcional) Proteja suas mensagens com uma senha. Os dados são criptografados antes de serem ocultados, garantindo segurança adicional.
- **Verificação de Capacidade:** Verifique quantos bytes podem ser escondidos em uma determinada imagem antes de tentar ocultar dados.
- **CLI Amigável:** Interface colorida com ASCII art e parâmetros claros.

## 📦 Instalação

Certifique-se de ter o [Go](https://golang.org/dl/) instalado (versão 1.21+).

1. Clone o repositório:
```bash
git clone https://github.com/kaique/stegano-go.git
cd stegano-go
```

2. Compile o projeto:
```bash
go build -o stegano-go main.go
```

## 🛠️ Como Usar

O `stegano-go` possui três modos principais: `encode`, `decode` e `capacity`.

### 1. Ocultar Mensagem (Encode)
Ocultar um texto diretamente pelo terminal:
```bash
./stegano-go -mode encode -image cover.png -output secret.png -message "Esta é uma mensagem ultra secreta"
```

Ocultar com proteção de senha (Criptografia AES):
```bash
./stegano-go -mode encode -image cover.png -output secret.png -message "Mensagem secreta" -password "MinhaSenhaForte"
```

Ocultar um arquivo de texto:
```bash
./stegano-go -mode encode -image cover.png -output secret.png -file payload.txt
```

### 2. Extrair Mensagem (Decode)
Extrair uma mensagem não criptografada:
```bash
./stegano-go -mode decode -image secret.png
```

Extrair uma mensagem criptografada:
```bash
./stegano-go -mode decode -image secret.png -password "MinhaSenhaForte"
```

### 3. Verificar Capacidade
Verifica o quanto de dados pode ser inserido em uma imagem PNG específica:
```bash
./stegano-go -mode capacity -image cover.png
```

## 🖼️ Exemplos e Demonstração

Para testar a ferramenta, você pode adicionar suas próprias imagens PNG na pasta `samples/`. 
*(Nota: O repositório não inclui imagens de teste, por favor adicione as suas próprias para testar).*

## ⚠️ Disclaimer

Esta ferramenta foi desenvolvida com propósitos puramente **educacionais** e de pesquisa em segurança da informação. O autor não se responsabiliza pelo uso indevido deste software.

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.
