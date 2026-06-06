<div align="center">

# dev

**Um task runner minimalista com TUI para descoberta de tarefas.**

[![Go](https://img.shields.io/badge/Go-1.24+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

[🇺🇸 English](README.md)

</div>

---

`dev` é uma alternativa simples ao Makefile e Taskfile. Defina tarefas em um arquivo `dev.yaml` e execute-as por uma interface TUI interativa ou diretamente pela linha de comando.

```
┌────────────────────────────────────────┐
│ Dev                                    │
├────────────────────────────────────────┤
│ ▶ build                                │
│   lint                                 │
│   run                                  │
│   test                                 │
├────────────────────────────────────────┤
│ Descrição:                             │
│ Compilar a aplicação                   │
└────────────────────────────────────────┘

  ↑/↓ navegar  •  enter executar  •  q sair
```

---

## Índice

- [Instalação](#instalação)
  - [Linux e macOS — curl](#linux-e-macos--curl)
  - [Windows — PowerShell](#windows--powershell)
  - [go install](#go-install)
  - [Compilar do código-fonte](#compilar-do-código-fonte)
- [Configuração](#configuração)
- [Uso](#uso)
- [Atalhos de teclado](#atalhos-de-teclado)
- [Exemplos](#exemplos)

---

## Instalação

### Linux e macOS — curl

A forma mais rápida. Não precisa ter Go instalado.

```bash
curl -fsSL https://raw.githubusercontent.com/JoaoPedr0Maciel/dev/main/install.sh | sh
```

O script vai:
1. Detectar seu sistema operacional e arquitetura
2. Baixar o binário correto da [última release](https://github.com/JoaoPedr0Maciel/dev/releases/latest)
3. Instalar em `/usr/local/bin` (ou `~/.local/bin` se não tiver permissão de escrita)

Se o `dev` não for encontrado após a instalação, adicione o diretório ao `PATH`:

```bash
# Adicione ao ~/.bashrc ou ~/.zshrc
export PATH="$PATH:$HOME/.local/bin"
source ~/.bashrc   # ou source ~/.zshrc
```

---

### Windows — PowerShell

```powershell
$url = "https://github.com/JoaoPedr0Maciel/dev/releases/latest/download/dev_windows_amd64.exe"
$dest = "$env:USERPROFILE\.local\bin\dev.exe"
New-Item -ItemType Directory -Force -Path (Split-Path $dest) | Out-Null
Invoke-WebRequest -Uri $url -OutFile $dest
```

Em seguida, adicione `%USERPROFILE%\.local\bin` ao `PATH`:

1. Abra o **Menu Iniciar** → pesquise **"Variáveis de Ambiente"**
2. Em **Variáveis do usuário**, selecione `Path` → **Editar**
3. Clique em **Novo** e adicione: `%USERPROFILE%\.local\bin`
4. Clique em **OK** e reinicie o terminal

---

### go install

Se você tiver Go 1.24+ instalado:

```bash
go install github.com/JoaoPedr0Maciel/dev/cmd/dev@latest
```

Certifique-se de que `$GOPATH/bin` (normalmente `~/go/bin`) está no seu `PATH`:

```bash
export PATH="$PATH:$HOME/go/bin"
```

---

### Compilar do código-fonte

Funciona em Linux, macOS e Windows:

```bash
git clone https://github.com/JoaoPedr0Maciel/dev.git
cd dev
go build -o dev ./cmd/dev/

# Mova para um diretório no PATH (Linux/macOS)
sudo mv dev /usr/local/bin/

# No Windows, mova dev.exe para uma pasta que já está no PATH
```

---

## Configuração

Crie um arquivo `dev.yaml` na **raiz do seu projeto**:

```yaml
tasks:
  build:
    description: Compilar a aplicação
    cmd: go build .

  test:
    description: Executar todos os testes
    cmd: go test ./...

  run:
    description: Iniciar o servidor de desenvolvimento
    cmd: go run .

  lint:
    description: Executar o linter
    cmd: go vet ./...
```

### Campos

| Campo | Obrigatório | Descrição |
|---|---|---|
| `description` | Não | Texto exibido no painel de descrição da TUI |
| `cmd` | Sim | Comando shell a ser executado (`sh -c` no Linux/macOS) |

> **Atenção:** o `dev` procura pelo `dev.yaml` no **diretório atual**. Execute sempre a partir da raiz do projeto.

---

## Uso

### Modo TUI

Execute `dev` sem argumentos para abrir a interface interativa:

```bash
dev
```

Use o teclado para navegar pela lista de tarefas, ler a descrição e executar. O resultado é exibido com status de saída, duração e output.

### Modo direto

Execute uma tarefa pelo nome sem abrir a TUI:

```bash
dev <nome-da-task>
```

**Exemplo de sucesso:**

```
$ dev test

Running test...

✓ Success

Task:     test
Duration: 421ms
```

**Exemplo de falha:**

```
$ dev build

Running build...

✗ Failed

./main.go:12:2: undefined: someFunc
```

---

## Atalhos de teclado

| Tecla | Ação |
|---|---|
| `↑` ou `k` | Mover seleção para cima |
| `↓` ou `j` | Mover seleção para baixo |
| `Enter` | Executar a tarefa selecionada |
| `q` ou `Ctrl+C` | Sair |

---

## Exemplos

### Projeto Go típico

```yaml
tasks:
  build:
    description: Compilar o binário
    cmd: go build -o bin/app ./cmd/app/

  test:
    description: Rodar testes com race detector
    cmd: go test -race ./...

  cover:
    description: Abrir relatório de cobertura
    cmd: go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out

  lint:
    description: Executar golangci-lint
    cmd: golangci-lint run ./...

  tidy:
    description: Organizar módulos Go
    cmd: go mod tidy
```

### Projeto Node.js

```yaml
tasks:
  install:
    description: Instalar dependências
    cmd: npm install

  dev:
    description: Iniciar servidor de desenvolvimento
    cmd: npm run dev

  build:
    description: Build para produção
    cmd: npm run build

  test:
    description: Executar testes
    cmd: npm test

  lint:
    description: Verificar código-fonte
    cmd: npm run lint
```

### Fluxo com Docker

```yaml
tasks:
  up:
    description: Subir todos os containers
    cmd: docker compose up -d

  down:
    description: Parar todos os containers
    cmd: docker compose down

  logs:
    description: Acompanhar logs dos containers
    cmd: docker compose logs -f

  build:
    description: Rebuild das imagens
    cmd: docker compose build --no-cache
```

---

## Licença

MIT © [JoaoPedr0Maciel](https://github.com/JoaoPedr0Maciel)
