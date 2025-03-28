package main

import (
	"image/jpeg"
	"os"
	"path/filepath"
	"strings"

	"github.com/gen2brain/avif"
)

func main() {
	// Pega o diretório inicial como argumento ou usa o atual
	dir := "."
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}

	// Contador pra rastrear quantos arquivos foram convertidos
	converted := 0

	// Percorre todas as pastas e subpastas
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Só processa arquivos com extensão .avif
		if !info.IsDir() && strings.ToLower(filepath.Ext(path)) == ".avif" {
			err := convertAndDelete(path)
			if err != nil {
				println("Erro ao converter", path, ":", err.Error())
			} else {
				converted++
			}
		}
		return nil
	})

	if err != nil {
		println("Erro ao vasculhar diretórios:", err.Error())
		return
	}

	println("Conversão concluída. Arquivos convertidos:", converted)
}

// convertAndDelete converte um arquivo .avif pra .jpeg e deleta o original se der certo
func convertAndDelete(avifPath string) error {
	// Abre o arquivo AVIF
	avifFile, err := os.Open(avifPath)
	if err != nil {
		return err
	}
	defer avifFile.Close()

	// Decodifica o AVIF pra image.Image
	img, err := avif.Decode(avifFile)
	if err != nil {
		return err
	}

	// Gera o nome do arquivo de saída (.jpeg)
	jpegPath := strings.TrimSuffix(avifPath, filepath.Ext(avifPath)) + ".jpg"

	// Cria o arquivo de saída JPEG
	outFile, err := os.Create(jpegPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// Codifica como JPEG
	err = jpeg.Encode(outFile, img, &jpeg.Options{Quality: 85})
	if err != nil {
		return err
	}

	// Se chegou aqui, a conversão foi bem-sucedida. Deleta o arquivo .avif
	err = os.Remove(avifPath)
	if err != nil {
		println("Aviso: Conversão OK, mas falha ao deletar", avifPath, ":", err.Error())
		return nil // Não retorna erro pra não interromper o processo
	}

	println("Convertido e deletado:", avifPath, "->", jpegPath)
	return nil
}
