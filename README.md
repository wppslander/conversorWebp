# Conversor WebP e AVIF para PNG

Utilitário em Go para conversão em lote de imagens `.webp` e `.avif` para `.png`.

---

## Download dos Executáveis (Acesso Rápido)

Você não precisa compilar nem instalar Go para usar. Baixe a versão correspondente ao seu sistema operacional:

- 🪟 **[Download para Windows (conversorWebp.exe)](https://github.com/wppslander/conversorWebp/releases/latest/download/conversorWebp.exe)**
- 🐧 **[Download para Linux (conversorWebp)](https://github.com/wppslander/conversorWebp/releases/latest/download/conversorWebp)**

---

## Requisitos do Sistema

- **Go (Golang)**: Versão 1.24 ou superior instalada na máquina (caso deseje compilar manualmente).
  - Verifique com: `go version`

---

## Dependências do Projeto

O projeto utiliza bibliotecas sem CGo para decodificação multiplataforma:

| Pacote | Função | Repositório |
|---|---|---|
| `golang.org/x/image/webp` | Decodificador WebP | [golang.org/x/image](https://pkg.go.dev/golang.org/x/image/webp) |
| `github.com/gen2brain/avif` | Decodificador AVIF | [github.com/gen2brain/avif](https://pkg.go.dev/github.com/gen2brain/avif) |
| `image/png` | Codificador PNG | Biblioteca padrão do Go |

As dependências já estão declaradas nos arquivos `go.mod` e `go.sum`.

---

## Como Replicar do Zero

### 1. Clonar e restaurar dependências
```bash
# Baixar e verificar as dependências do módulo
go mod download
go mod verify
```

*(Se estiver recriando o projeto do zero sem `go.mod`, execute:)*
```bash
go mod init conversorWebp
go get golang.org/x/image/webp
go get github.com/gen2brain/avif
```

---

## Compilação

### Opção A: Usando o script compilador (Recomendado)
Gera os binários para Linux e Windows em uma única etapa:

```bash
chmod +x build.sh
./build.sh
```

Os executáveis otimizados serão gerados na raiz e na pasta `dist/`:
- `conversorWebp` (Linux amd64)
- `conversorWebp.exe` (Windows amd64)

### Opção B: Compilação manual via `go build`

- **Para Linux:**
  ```bash
  go build -ldflags="-s -w" -o conversorWebp .
  ```

- **Para Windows (Cross-compilation):**
  ```bash
  GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o conversorWebp.exe .
  ```

---

## Como Usar

1. Coloque o executável (`conversorWebp` ou `conversorWebp.exe`) na pasta desejada.
2. Execute o programa:
   - **Linux:** `./conversorWebp`
   - **Windows:** Clique duas vezes em `conversorWebp.exe` ou execute via terminal.
3. **Primeira execução:**
   - Se a pasta `Fotos Produtos` não existir, o programa a cria automaticamente e solicita que você insira as imagens `.webp` ou `.avif` dentro dela.
4. **Execução com imagens:**
   - O programa converte cada arquivo `.webp` e `.avif` para `.png`.
   - As imagens originais são movidas para `Fotos Produtos/originais/`.
   - Ao finalizar, o programa emite aviso de conclusão.
