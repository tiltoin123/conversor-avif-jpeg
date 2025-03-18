package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Verifica se o usuário passou um argumento
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run main.go <caminho_da_pasta>")
		return
	}

	// Pega o caminho passado como argumento
	rootPath := os.Args[1]

	// Verifica se o caminho existe
	if _, err := os.Stat(rootPath); os.IsNotExist(err) {
		fmt.Println("Erro: O caminho especificado não existe.")
		return
	}

	// Percorre todas as pastas e subpastas
	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Erro ao acessar:", path, "-", err)
			return nil
		}

		// Verifica se é um arquivo .avif
		if !info.IsDir() && strings.HasSuffix(strings.ToLower(info.Name()), ".avif") {
			fmt.Println(path) // Exibe o caminho do arquivo
		}
		return nil
	})
	// Trata erro na busca
	if err != nil {
		fmt.Println("Erro ao vasculhar diretórios:", err)
	}
}
