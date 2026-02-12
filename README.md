# Concurrent Data Pipeline with Go

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![JSON](https://img.shields.io/badge/Data%20Format-JSON-000000?style=for-the-badge&logo=json&logoColor=white)
![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg?style=for-the-badge)

---

## 🇧🇷 Pipeline de Dados Concorrente com Go

Uma pipeline de dados concorrente escrita em Go, usando goroutines e channels para processar registros em paralelo. O projeto demonstra validação, transformação, detecção de anomalias e persistência de dados em formato JSONL.

### 🎯 Objetivo

Demonstrar a construção de uma pipeline de dados concorrente em Go, cobrindo validação de dados, transformações, tratamento de erros e coleta de métricas.

### ✨ Destaques

- **Concorrência com goroutines e channels**: Pipeline multi-estágio com producer, validator, transformer, loader, error handler e metrics collector rodando em paralelo.
- **Validação de dados**: Filtragem de registros inválidos com roteamento para um error handler dedicado.
- **Transformações com detecção de anomalias**: Cálculo de scores de anomalia e classificação de registros.
- **Coleta de métricas**: Contagem de registros processados, erros e anomalias ao longo da execução.
- **Persistência em JSONL**: Saída de registros processados e com falha em arquivos JSONL.
- **Workers configuráveis**: Número de goroutines de validação e transformação ajustável.
- **Testes unitários e de integração**: Cobertura das etapas individuais e da pipeline completa.

---

## 🇬🇧 Concurrent Data Pipeline with Go

A concurrent data pipeline written in Go, using goroutines and channels to process records in parallel. The project demonstrates validation, transformation, anomaly detection, and JSONL data persistence.

### 🎯 Objective

Show how to build a concurrent data pipeline in Go, covering data validation, transformations, error handling, and metrics collection.

### ✨ Highlights

- **Concurrency with goroutines and channels**: Multi-stage pipeline with producer, validator, transformer, loader, error handler, and metrics collector running in parallel.
- **Data validation**: Invalid record filtering with routing to a dedicated error handler.
- **Transformations with anomaly detection**: Anomaly score calculation and record classification.
- **Metrics collection**: Counts of processed records, errors, and anomalies throughout execution.
- **JSONL persistence**: Output of processed and failed records to JSONL files.
- **Configurable workers**: Adjustable number of validation and transformation goroutines.
- **Unit and integration tests**: Coverage of individual stages and the full pipeline.

### 📊 Visualization

![Go Data Pipeline Flow](diagrams/go_data_pipeline_flow.png)

*Diagrama ilustrativo da arquitetura da pipeline de dados concorrente em Go, destacando as etapas de processamento e a comunicação entre elas.*


---

## 🛠️ Tecnologias Utilizadas / Technologies Used

| Categoria         | Tecnologia      | Descrição                                                                 |
| :---------------- | :-------------- | :------------------------------------------------------------------------ |
| **Linguagem**     | Go              | Linguagem principal para desenvolvimento da pipeline de dados concorrente. |
| **Concorrência**  | Goroutines, Channels | Primitivas nativas do Go para programação concorrente e comunicação segura. |
| **Formato de Dados** | JSONL           | Formato de arquivo para armazenamento de dados processados e com erro.    |
| **Testes**        | `testing`       | Pacote padrão do Go para escrita de testes unitários e de integração.     |
| **Logging**       | `log`           | Pacote padrão do Go para registro de eventos e mensagens da pipeline.     |
| **Diagramação**   | Mermaid         | Para criação de diagramas de arquitetura e fluxo de dados no README.      |

---

## 📁 Repository Structure

```
go-concurrent-data-pipeline/
├── src/               # Main application entry point (main.go)
├── pkg/pipeline/      # Core pipeline implementation modules
│   ├── types.go       # Data structures and type definitions
│   ├── producer.go    # Data generation/ingestion stage
│   ├── validator.go   # Data validation stage
│   ├── transformer.go # Data transformation and enrichment stage
│   ├── loader.go      # Data persistence stage
│   ├── errorHandler.go # Error handling and logging stage
│   ├── metricsCollector.go # Metrics aggregation stage
│   ├── run.go         # Pipeline orchestration logic
│   └── pipeline_test.go # Unit tests for pipeline components
├── tests/             # Integration tests
│   └── main_test.go   # End-to-end pipeline tests
├── docs/              # Documentation
│   ├── ARCHITECTURE.md # Detailed architecture documentation
│   └── CONTRIBUTING.md # Contribution guidelines
├── config/            # Configuration files
│   └── config.example.yaml # Example configuration
├── data/              # Sample data files for testing
│   ├── sample_input.jsonl # Example input data
│   └── README.md
├── diagrams/          # Architecture diagrams
│   ├── go_data_pipeline_flow.mmd # Mermaid diagram source
│   ├── go_pipeline.mmd           # Simplified pipeline diagram
│   └── go_data_pipeline_flow.png # Rendered diagram
├── images/            # Images for documentation
├── logs/              # Log files directory (gitignored)
│   └── README.md
├── .gitignore         # Git ignore rules
├── go.mod             # Go module definition
├── LICENSE            # MIT License
└── README.md          # This file
```

---

## 🚀 Getting Started

### Pré-requisitos

- Go 1.18 ou superior ([Baixar Go](https://go.dev/dl/))
- Git

### Instalação

```bash
# Clone o repositório
git clone https://github.com/galafis/go-concurrent-data-pipeline.git
cd go-concurrent-data-pipeline

# Verifique a instalação do Go
go version

# Execute os testes para validar a instalação
go test ./... -v

# Compile o projeto
go build -o pipeline ./src/main.go

# Execute o pipeline
./pipeline
```

### Execução Rápida

```bash
# Execute diretamente sem compilar
go run src/main.go

# Execute com número customizado de registros e workers
# (modifique os parâmetros em src/main.go)
```

### Visualizando os Resultados

Após executar o pipeline, você encontrará:

```bash
# Dados processados com sucesso
cat processed_data.jsonl

# Registros que falharam na validação/transformação
cat failed_data.jsonl
```

### Exemplo de Uso Avançado (Go)

O exemplo abaixo demonstra a execução da pipeline de dados concorrente, utilizando o módulo `pkg/pipeline` para orquestrar as etapas de geração de registros, validação, transformação (com detecção de anomalias), carregamento para arquivos de saída e tratamento de erros. Um `metricsCollector` sumariza o desempenho da pipeline, fornecendo uma visão completa do fluxo de dados.

```go
package main

import (
	"fmt"
	"log"
	"os"
	"go-concurrent-data-pipeline/pkg/pipeline"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	fmt.Println("===========================================")
	fmt.Println("Go Concurrent Data Pipeline")
	fmt.Println("===========================================")

	// Limpar arquivos de saída anteriores
	_ = os.Remove("processed_data.jsonl")
	_ = os.Remove("failed_data.jsonl")

	// Executar a pipeline com 50 registros e 3 workers para validação/transformação
	// Os logs detalhados serão exibidos no console e as métricas no final.
	metrics := pipeline.RunAdvancedPipeline(50, 3)

	fmt.Println("===========================================")
	fmt.Println("Pipeline completed!")
	fmt.Printf("Final Metrics: Processed=%d, Errors=%d, Anomalies=%d\n", 
		metrics.ProcessedCount, metrics.ErrorCount, metrics.AnomalyCount)
	fmt.Println("===========================================")

	// Opcional: Ler os arquivos de saída para verificar o conteúdo
	fmt.Println("\nConteúdo de processed_data.jsonl:")
	processedContent, err := os.ReadFile("processed_data.jsonl")
	if err == nil {
		fmt.Println(string(processedContent))
	} else {
		fmt.Println("  (Arquivo não encontrado ou vazio)")
	}

	fmt.Println("\nConteúdo de failed_data.jsonl:")
	failedContent, err := os.ReadFile("failed_data.jsonl")
	if err == nil {
		fmt.Println(string(failedContent))
	} else {
		fmt.Println("  (Arquivo não encontrado ou vazio)")
	}
}
```

### Executando os Testes

```bash
# Execute todos os testes
go test ./... -v

# Execute apenas testes unitários da pipeline
go test ./pkg/pipeline -v

# Execute apenas testes de integração
go test ./tests -v

# Execute com cobertura de código
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out

# Execute benchmarks
go test ./pkg/pipeline -bench=. -benchmem
```

### Exemplo de Saída

```
===========================================
Go Concurrent Data Pipeline
===========================================
2025/10/16 01:46:08 Producer: Iniciando produção de dados...
2025/10/16 01:46:08 Validator: Validando registro rec-0000
2025/10/16 01:46:08 Transformer: Transformando registro rec-0001
2025/10/16 01:46:08 Loader: Carregado rec-0001 (Anomaly: false)
...
2025/10/16 01:46:09 --- Sumário da Pipeline ---
2025/10/16 01:46:09 Registros Processados com Sucesso: 41
2025/10/16 01:46:09 Registros com Erro: 9
2025/10/16 01:46:09 Registros Anômalos: 8
2025/10/16 01:46:09 Valor Total Processado: 2137.08
2025/10/16 01:46:09 ---------------------------
===========================================
Pipeline completed!
Final Metrics: Processed=41, Errors=9, Anomalies=8
===========================================
```

---

## 🤝 Contribuição

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues, enviar pull requests ou sugerir melhorias.

Por favor, leia [CONTRIBUTING.md](docs/CONTRIBUTING.md) para detalhes sobre o processo de contribuição e código de conduta.

### Como Contribuir

1. Fork o repositório
2. Crie uma branch para sua feature (`git checkout -b feature/MinhaFeature`)
3. Commit suas mudanças (`git commit -m 'Adiciona MinhaFeature'`)
4. Push para a branch (`git push origin feature/MinhaFeature`)
5. Abra um Pull Request

---

## 📚 Documentação Adicional

- [Arquitetura Detalhada](docs/ARCHITECTURE.md) - Decisões de design e padrões utilizados
- [Guia de Contribuição](docs/CONTRIBUTING.md) - Como contribuir com o projeto
- [Dados de Exemplo](data/README.md) - Informações sobre os dados de teste

---

## 🧪 Testes e Qualidade

- **Testes Unitários** — Cobertura dos componentes individuais da pipeline
- **Testes de Integração** — Validação end-to-end da pipeline completa
- **Benchmarks** — Benchmarks para Producer e Validator disponíveis em `pkg/pipeline/pipeline_test.go`

---

## 🔍 Casos de Uso

Este projeto pode ser adaptado para diversos cenários:

1. **IoT Data Processing** - Processar streams de dados de sensores IoT
2. **Log Aggregation** - Agregar e processar logs de aplicações
3. **ETL Pipelines** - Extract, Transform, Load de dados
4. **Real-time Analytics** - Análise de dados em tempo real
5. **Data Validation** - Validação em lote de grandes volumes de dados
6. **Anomaly Detection** - Detecção de anomalias em séries temporais

---

## 🛠️ Ferramentas e Recursos

### Recomendações de Desenvolvimento

- **IDE**: VSCode com extensão Go, GoLand, ou Vim com vim-go
- **Debugging**: Delve debugger (`dlv`)
- **Profiling**: pprof para análise de performance
- **Monitoring**: Prometheus + Grafana (para produção)

### Recursos para Aprendizado

- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go Concurrency Patterns](https://go.dev/blog/pipelines)
- [Advanced Go Concurrency Patterns](https://go.dev/blog/io2013-talk-concurrency)

---

## ❓ FAQ

**P: Quantos workers devo usar?**  
R: Comece com o número de CPUs disponíveis. Ajuste baseado em profiling e métricas.

**P: Como escalar para milhões de registros?**  
R: Aumente o número de workers e considere processar em batches. Para volumes muito grandes, considere distribuir em múltiplas instâncias.

**P: Posso usar com dados de fontes externas?**  
R: Sim! Modifique o Producer para ler de Kafka, databases, APIs, etc.

**P: Como adicionar persistência em banco de dados?**  
R: Modifique o Loader para escrever em PostgreSQL, MongoDB, etc., em vez de arquivos.

**P: É adequado para produção?**  
R: O código demonstra padrões de produção, mas adicione configuração, monitoramento, e tratamento de erros mais robusto para uso em produção.

---

## 📝 Licença

Este projeto está licenciado sob a Licença MIT - veja o arquivo [LICENSE](LICENSE) para detalhes.

```
MIT License

Copyright (c) 2025 Gabriel Demetrios Lafis

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.
```

---

## 👨‍💻 Autor

**Gabriel Demetrios Lafis**

- GitHub: [@galafis](https://github.com/galafis)
- LinkedIn: [Gabriel Demetrios Lafis](https://www.linkedin.com/in/gabriel-demetrios-lafis)

---

## 🙏 Agradecimentos

- Comunidade Go por excelentes recursos e documentação
- Contribuidores que ajudaram a melhorar este projeto
- Todos que usam e fornecem feedback

---

## 📈 Status do Projeto

**Última atualização:** Outubro 2025

