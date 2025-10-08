// Go Concurrent Data Pipeline
// Author: Gabriel Demetrios Lafis
// Year: 2025

package pipeline

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"time"
)

// Loader carrega os registros processados para um destino (simulado).
func Loader(in <-chan ProcessedRecord) {
	log.Println("Loader: Iniciando carregamento de dados...")
	file, err := os.Create("processed_data.jsonl")
	if err != nil {
		log.Fatalf("Loader: Falha ao criar arquivo de saída: %v", err)
	}
	defer file.Close()

	for record := range in {
		jsonBytes, err := json.Marshal(record)
		if err != nil {
			log.Printf("Loader: Erro ao serializar registro %s: %v", record.ID, err)
			continue
		}
		_, err = file.WriteString(string(jsonBytes) + "\n")
		if err != nil {
			log.Printf("Loader: Erro ao escrever registro %s no arquivo: %v", record.ID, err)
			continue
		}
		log.Printf("Loader: Carregado %s (Anomaly: %t)", record.ID, record.IsAnomaly)
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
	}
	log.Println("Loader: Carregamento de dados finalizado.")
}

