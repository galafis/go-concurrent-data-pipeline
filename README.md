# 📊 Go Concurrent Data Pipeline

[![Go](https://img.shields.io/badge/Go-1.22-00ADD8.svg)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

[English](#english) | [Português](#português)

---

## English

### 🎯 Overview

**Go Concurrent Data Pipeline** — High-throughput concurrent data pipeline in Go. Leverages goroutines and channels for parallel data processing, transformation, and loading.

Total source lines: **827** across **11** files in **1** language.

### ✨ Key Features

- **Production-Ready Architecture**: Modular, well-documented, and following best practices
- **Comprehensive Implementation**: Complete solution with all core functionality
- **Clean Code**: Type-safe, well-tested, and maintainable codebase
- **Easy Deployment**: Docker support for quick setup and deployment

### 🚀 Quick Start

#### Prerequisites
- Go 1.22+


#### Installation

1. **Clone the repository**
```bash
git clone https://github.com/galafis/go-concurrent-data-pipeline.git
cd go-concurrent-data-pipeline
```

2. **Install dependencies**
```bash
go mod download
```

#### Running

```bash
go run ./...
```


### 🧪 Testing

```bash
go test ./...
```

### 📁 Project Structure

```
go-concurrent-data-pipeline/
├── config/
│   └── config.example.yaml
├── data/
│   └── README.md
├── diagrams/
├── docs/
│   ├── ARCHITECTURE.md
│   └── CONTRIBUTING.md
├── images/
├── logs/
│   └── README.md
├── pkg/
│   └── pipeline/
│       ├── errorHandler.go
│       ├── loader.go
│       ├── metricsCollector.go
│       ├── pipeline_test.go
│       ├── producer.go
│       ├── run.go
│       ├── transformer.go
│       ├── types.go
│       └── validator.go
├── src/
│   └── main.go
├── tests/
│   └── main_test.go
└── README.md
```

### 🛠️ Tech Stack

| Technology | Usage |
|------------|-------|
| Go | 11 files |

### 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

### 👤 Author

**Gabriel Demetrios Lafis**

- GitHub: [@galafis](https://github.com/galafis)
- LinkedIn: [Gabriel Demetrios Lafis](https://linkedin.com/in/gabriel-demetrios-lafis)

---

## Português

### 🎯 Visão Geral

**Go Concurrent Data Pipeline** — High-throughput concurrent data pipeline in Go. Leverages goroutines and channels for parallel data processing, transformation, and loading.

Total de linhas de código: **827** em **11** arquivos em **1** linguagem.

### ✨ Funcionalidades Principais

- **Arquitetura Pronta para Produção**: Modular, bem documentada e seguindo boas práticas
- **Implementação Completa**: Solução completa com todas as funcionalidades principais
- **Código Limpo**: Type-safe, bem testado e manutenível
- **Fácil Implantação**: Suporte Docker para configuração e implantação rápidas

### 🚀 Início Rápido

#### Pré-requisitos
- Go 1.22+


#### Instalação

1. **Clone the repository**
```bash
git clone https://github.com/galafis/go-concurrent-data-pipeline.git
cd go-concurrent-data-pipeline
```

2. **Install dependencies**
```bash
go mod download
```

#### Execução

```bash
go run ./...
```

### 🧪 Testes

```bash
go test ./...
```

### 📁 Estrutura do Projeto

```
go-concurrent-data-pipeline/
├── config/
│   └── config.example.yaml
├── data/
│   └── README.md
├── diagrams/
├── docs/
│   ├── ARCHITECTURE.md
│   └── CONTRIBUTING.md
├── images/
├── logs/
│   └── README.md
├── pkg/
│   └── pipeline/
│       ├── errorHandler.go
│       ├── loader.go
│       ├── metricsCollector.go
│       ├── pipeline_test.go
│       ├── producer.go
│       ├── run.go
│       ├── transformer.go
│       ├── types.go
│       └── validator.go
├── src/
│   └── main.go
├── tests/
│   └── main_test.go
└── README.md
```

### 🛠️ Stack Tecnológica

| Tecnologia | Uso |
|------------|-----|
| Go | 11 files |

### 📄 Licença

Este projeto está licenciado sob a Licença MIT - veja o arquivo [LICENSE](LICENSE) para detalhes.

### 👤 Autor

**Gabriel Demetrios Lafis**

- GitHub: [@galafis](https://github.com/galafis)
- LinkedIn: [Gabriel Demetrios Lafis](https://linkedin.com/in/gabriel-demetrios-lafis)
