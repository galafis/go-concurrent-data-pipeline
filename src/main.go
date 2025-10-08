// Go Concurrent Data Pipeline
// Author: Gabriel Demetrios Lafis
// Year: 2025

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
	pipeline.RunAdvancedPipeline(50, 3)

	fmt.Println("===========================================")
	fmt.Println("Pipeline completed!")
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

