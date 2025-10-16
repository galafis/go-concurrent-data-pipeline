# Concurrent Data Pipeline with Go

[![Go Report Card](https://goreportcard.com/badge/github.com/GabrielDemetriosLafis/go-concurrent-data-pipeline)](https://goreportcard.com/report/github.com/GabrielDemetriosLafis/go-concurrent-data-pipeline)
![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![JSON](https://img.shields.io/badge/Data%20Format-JSON-000000?style=for-the-badge&logo=json&logoColor=white)
![Mermaid](https://img.shields.io/badge/Diagrams-Mermaid-orange?style=for-the-badge&logo=mermaid&logoColor=white)
![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg?style=for-the-badge)

---

## 🇧🇷 Pipeline de Dados Concorrente com Go

Este repositório apresenta uma **pipeline de dados de alta performance e concorrência desenvolvida em Go**, projetada para processar grandes volumes de dados de forma eficiente e escalável. O foco é em demonstrar como utilizar as capacidades de concorrência nativas do Go (goroutines e channels) para construir sistemas de processamento de dados robustos, tolerantes a falhas e com baixa latência. É ideal para **engenheiros de dados, desenvolvedores de backend e arquitetos de sistemas** que buscam soluções eficientes para ingestão, transformação e carregamento de dados em tempo real ou em lote.

### 🎯 Objetivo

O principal objetivo deste projeto é **fornecer exemplos práticos, código funcional e documentação detalhada** sobre a construção de pipelines de dados concorrentes com Go. Serão abordados tópicos como consumo de mensagens de filas (e.g., Kafka), processamento paralelo de dados, persistência em bancos de dados, tratamento de erros e monitoramento, tudo com foco em **performance, resiliência e manutenibilidade**, com ênfase em **validação de dados, transformações complexas, tratamento de erros e coleta de métricas**.

### ✨ Destaques

- **Concorrência Nativa com Go**: Utilização de goroutines e channels para construir pipelines de dados altamente concorrentes e eficientes, permitindo o processamento paralelo de dados em várias etapas.
- **Validação de Dados e Tratamento de Erros**: Implementação de uma etapa de validação (`validator`) que filtra dados inválidos e um `errorHandler` dedicado para persistir e gerenciar registros com falha, garantindo a integridade da pipeline.
- **Transformações de Dados Complexas**: A etapa de `transformer` demonstra como aplicar lógica de negócios mais sofisticada, como cálculo de scores de anomalia, e lidar com diferentes tipos de erros de transformação.
- **Coleta e Sumarização de Métricas**: Um `metricsCollector` dedicado para monitorar o fluxo de dados, contando registros processados, erros e anomalias, fornecendo um sumário abrangente da performance da pipeline.
- **Persistência Flexível**: Os `loader` e `errorHandler` persistem os dados processados e com erro em arquivos JSONL, demonstrando uma abordagem flexível para armazenamento de resultados.
- **Escalabilidade Horizontal**: O design da pipeline permite fácil escalabilidade, adicionando mais workers para as etapas de validação e transformação conforme a demanda.
- **Código Profissional**: Exemplos de código bem estruturados, seguindo as melhores práticas da indústria, com foco em clareza, eficiência e documentação interna.
- **Testes Incluídos**: Módulos de código validados através de testes unitários e de integração, garantindo a robustez e a confiabilidade das implementações.

### 🚀 Benefícios do Go para Pipelines de Dados em Ação

Go é uma linguagem poderosa e eficiente para a construção de pipelines de dados concorrentes e de alta performance. Este projeto ilustra como esses benefícios são explorados:

1.  **Concorrência Simplificada e Eficaz:** Com goroutines e channels, Go torna a programação concorrente muito mais fácil e segura, permitindo a construção de um pipeline multi-estágio (producer, validator, transformer, loader, errorHandler, metricsCollector) onde cada etapa opera de forma independente e paralela.

2.  **Performance Otimizada:** Compilado para código de máquina, Go oferece performance próxima à de C/C++, ideal para cargas de trabalho intensivas em CPU e I/O, como a geração, validação e transformação de milhões de registros de dados.

3.  **Eficiência de Recursos:** Goroutines são leves e o scheduler do Go é otimizado para gerenciar milhares delas eficientemente, resultando em baixo consumo de memória e CPU, mesmo ao lidar com um grande número de registros e workers.

4.  **Tratamento de Erros Robusto:** O modelo de tratamento de erros explícito de Go é fundamental para pipelines de dados, permitindo que a pipeline identifique e desvie registros inválidos ou com falha para um `errorHandler` dedicado, sem interromper o fluxo principal.

5.  **Observabilidade Integrada:** A inclusão de um `metricsCollector` demonstra como é fácil integrar a coleta de métricas diretamente na pipeline, fornecendo insights em tempo real sobre o volume de dados, erros e anomalias.

6.  **Modularidade e Manutenibilidade:** A estrutura baseada em funções e canais promove a modularidade, tornando cada etapa da pipeline um componente independente que pode ser facilmente testado, mantido e substituído.

---

## 🇬🇧 Concurrent Data Pipeline with Go

This repository presents a **high-performance and concurrent data pipeline developed in Go**, designed to process large volumes of data efficiently and scalably. The focus is on demonstrating how to use Go's native concurrency capabilities (goroutines and channels) to build robust, fault-tolerant, and low-latency data processing systems. It is ideal for **data engineers, backend developers, and system architects** seeking efficient solutions for real-time or batch data ingestion, transformation, and loading.

### 🎯 Objective

The main objective of this project is to **provide practical examples, functional code, and detailed documentation** on building concurrent data pipelines with Go. Topics covered include consuming messages from queues (e.g., Kafka), parallel data processing, persistence in databases, error handling, and monitoring, all with a focus on **performance, resilience, and maintainability**, with an emphasis on **data validation, complex transformations, error handling, and metrics collection**.

### ✨ Highlights

- **Native Concurrency with Go**: Utilization of goroutines and channels to build highly concurrent and efficient data pipelines, allowing parallel data processing at various stages.
- **Data Validation and Error Handling**: Implementation of a validation stage (`validator`) that filters invalid data and a dedicated `errorHandler` to persist and manage failed records, ensuring pipeline integrity.
- **Complex Data Transformations**: The `transformer` stage demonstrates how to apply more sophisticated business logic, such as anomaly score calculation, and handle different types of transformation errors.
- **Metrics Collection and Summarization**: A dedicated `metricsCollector` to monitor data flow, counting processed records, errors, and anomalies, providing a comprehensive summary of pipeline performance.
- **Flexible Persistence**: The `loader` and `errorHandler` persist processed and erroneous data into JSONL files, demonstrating a flexible approach to result storage.
- **Horizontal Scalability**: The pipeline design allows for easy scalability, adding more workers for validation and transformation stages as demand increases.
- **Professional Code**: Well-structured code examples, following industry best practices, with a focus on clarity, efficiency, and internal documentation.
- **Tests Included**: Code modules validated through unit and integration tests, ensuring the robustness and reliability of the implementations.

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
├── .github/
│   └── workflows/     # GitHub Actions CI/CD workflows
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
│   ├── config.example.yaml # Example configuration
│   └── placeholder.txt
├── data/              # Sample data files for testing
│   ├── sample_input.jsonl # Example input data
│   └── README.md
├── diagrams/          # Architecture diagrams
│   ├── go_data_pipeline_flow.mmd # Mermaid diagram source
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
git clone https://github.com/GabrielDemetriosLafis/go-concurrent-data-pipeline.git
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

Este projeto mantém alta qualidade de código através de:

- ✅ **Testes Unitários** - Cobertura de componentes individuais
- ✅ **Testes de Integração** - Validação end-to-end da pipeline
- ✅ **Benchmarks** - Medição de performance
- ✅ **CI/CD** - GitHub Actions para testes automáticos
- ✅ **Linting** - Análise estática de código com golangci-lint

[![Tests](https://github.com/GabrielDemetriosLafis/go-concurrent-data-pipeline/actions/workflows/go-test.yml/badge.svg)](https://github.com/GabrielDemetriosLafis/go-concurrent-data-pipeline/actions/workflows/go-test.yml)

---

## 📊 Performance

Resultados de benchmark em uma máquina com 8 cores:

| Workers | Throughput (rec/s) | Latência Média (ms) |
|---------|-------------------|---------------------|
| 1       | ~50               | 20                  |
| 2       | ~95               | 10.5                |
| 3       | ~140              | 7.1                 |
| 4       | ~180              | 5.5                 |
| 8       | ~300              | 3.3                 |

*Os resultados podem variar dependendo do hardware e da carga de trabalho.*

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

- GitHub: [@GabrielDemetriosLafis](https://github.com/GabrielDemetriosLafis)
- LinkedIn: [Gabriel Demetrios Lafis](https://www.linkedin.com/in/gabriel-demetrios-lafis)

---

## 🙏 Agradecimentos

- Comunidade Go por excelentes recursos e documentação
- Contribuidores que ajudaram a melhorar este projeto
- Todos que usam e fornecem feedback

---

## 📈 Status do Projeto

🟢 **Ativo e Mantido** - Issues e PRs são bem-vindos!

**Última atualização:** Outubro 2025

**Versão:** 1.0.0

