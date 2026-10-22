package main

import (
	"bufio"
	"fmt"
	"image"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/gen2brain/avif"
	"golang.org/x/image/webp"
)

const (
	targetFolder   = "Fotos Produtos"
	originalFolder = "originais"
)

func main() {
	var exitCode int
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "\nErro: %v\n", err)
		exitCode = 1
	}

	fmt.Println("\nPressione Enter para sair...")
	_, _ = bufio.NewReader(os.Stdin).ReadBytes('\n')

	if exitCode != 0 {
		os.Exit(exitCode)
	}
}

func isSupportedFormat(ext string) bool {
	switch strings.ToLower(ext) {
	case ".webp", ".avif":
		return true
	default:
		return false
	}
}

func run() error {
	baseDir := targetFolder
	info, err := os.Stat(baseDir)
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(baseDir, 0755); err != nil {
				return fmt.Errorf("falha ao criar pasta '%s': %w", baseDir, err)
			}
			fmt.Printf("Pasta '%s' criada com sucesso.\nPor favor, coloque os arquivos (.webp, .avif) ali dentro e execute o programa novamente.\n", targetFolder)
			return nil
		}
		return fmt.Errorf("falha ao acessar '%s': %w", baseDir, err)
	}
	if !info.IsDir() {
		return fmt.Errorf("'%s' existe mas não é um diretório", baseDir)
	}

	originaisDir := filepath.Join(baseDir, originalFolder)
	if err := os.MkdirAll(originaisDir, 0755); err != nil {
		return fmt.Errorf("falha ao criar pasta '%s': %w", originaisDir, err)
	}

	entries, err := os.ReadDir(baseDir)
	if err != nil {
		return fmt.Errorf("falha ao ler arquivos de '%s': %w", baseDir, err)
	}

	var filesToConvert []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if isSupportedFormat(filepath.Ext(entry.Name())) {
			filesToConvert = append(filesToConvert, entry.Name())
		}
	}

	if len(filesToConvert) == 0 {
		fmt.Println("Nenhum arquivo (.webp, .avif) encontrado para converter.")
		fmt.Println("Processo terminado.")
		return nil
	}

	total := len(filesToConvert)
	successCount := 0

	for i, fileName := range filesToConvert {
		srcPath := filepath.Join(baseDir, fileName)
		baseName := strings.TrimSuffix(fileName, filepath.Ext(fileName))
		dstPngPath := filepath.Join(baseDir, baseName+".png")
		dstOriginalPath := filepath.Join(originaisDir, fileName)

		fmt.Printf("[%d/%d] Convertendo: %s -> %s\n", i+1, total, fileName, filepath.Base(dstPngPath))

		if err := convertImageToPng(srcPath, dstPngPath); err != nil {
			fmt.Fprintf(os.Stderr, "  Falha ao converter %s: %v\n", fileName, err)
			continue
		}

		if err := moveFile(srcPath, dstOriginalPath); err != nil {
			fmt.Fprintf(os.Stderr, "  Falha ao mover %s para originais: %v\n", fileName, err)
			continue
		}

		successCount++
	}

	fmt.Printf("\nConversão concluída: %d de %d arquivos processados com sucesso.\n", successCount, total)
	fmt.Println("Processo terminado com sucesso.")
	return nil
}

func convertImageToPng(srcPath, dstPath string) error {
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return fmt.Errorf("não foi possível abrir o arquivo de origem: %w", err)
	}
	defer srcFile.Close()

	ext := strings.ToLower(filepath.Ext(srcPath))
	var img image.Image
	switch ext {
	case ".webp":
		img, err = webp.Decode(srcFile)
	case ".avif":
		img, err = avif.Decode(srcFile)
	default:
		return fmt.Errorf("formato não suportado: %s", ext)
	}
	if err != nil {
		return fmt.Errorf("falha na decodificação %s: %w", ext, err)
	}

	dstFile, err := os.Create(dstPath)
	if err != nil {
		return fmt.Errorf("não foi possível criar o arquivo de destino png: %w", err)
	}
	defer dstFile.Close()

	if err := png.Encode(dstFile, img); err != nil {
		return fmt.Errorf("falha ao codificar png: %w", err)
	}

	return nil
}

// moveFile move o arquivo de origem para o destino com fallback de cópia e remoção
func moveFile(src, dst string) error {
	if err := os.Rename(src, dst); err == nil {
		return nil
	}

	// Fallback caso rename falhe entre volumes/sistemas de arquivos distintos
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	if _, err = io.Copy(out, in); err != nil {
		return err
	}

	if err = out.Sync(); err != nil {
		return err
	}

	in.Close()
	return os.Remove(src)
}
