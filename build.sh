#!/usr/bin/env bash
set -e

# Diretório raiz do projeto
ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$ROOT_DIR"

OUTPUT_DIR="dist"
mkdir -p "$OUTPUT_DIR"

echo "=== Compilando Conversor WebP ==="

# Compilação para Linux (64-bit)
echo "-> Compilando para Linux (amd64)..."
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o "$OUTPUT_DIR/conversorWebp" .
cp "$OUTPUT_DIR/conversorWebp" "$ROOT_DIR/conversorWebp"

# Compilação para Windows (64-bit)
echo "-> Compilando para Windows (amd64)..."
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o "$OUTPUT_DIR/conversorWebp.exe" .
cp "$OUTPUT_DIR/conversorWebp.exe" "$ROOT_DIR/conversorWebp.exe"

echo "=== Compilação concluída com sucesso! ==="
echo "Arquivos gerados:"
echo " - $OUTPUT_DIR/conversorWebp (Linux)"
echo " - $OUTPUT_DIR/conversorWebp.exe (Windows)"
