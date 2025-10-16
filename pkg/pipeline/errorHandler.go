// Go Concurrent Data Pipeline
// Author: Gabriel Demetrios Lafis
// Year: 2025

package pipeline

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

// ErrorHandler lida com registros que falharam em alguma etapa.
func ErrorHandler(errorCh <-chan DataRecord) {
	log.Println("ErrorHandler: Iniciando tratamento de erros...")
	file, err := os.Create("failed_data.jsonl")
	if err != nil {
		log.Fatalf("ErrorHandler: Falha ao criar arquivo de erros: %v", err)
	}
	defer func() { _ = file.Close() }()

	for record := range errorCh {
		jsonBytes, err := json.Marshal(record)
		if err != nil {
			log.Printf("ErrorHandler: Erro ao serializar registro de erro %s: %v", record.ID, err)
			continue
		}
		_, err = file.WriteString(string(jsonBytes) + "\n")
		if err != nil {
			log.Printf("ErrorHandler: Erro ao escrever registro de erro %s no arquivo: %v", record.ID, err)
			continue
		}
		log.Printf("ErrorHandler: Registro %s falhou permanentemente. Motivo: %s", record.ID, record.Error)
		time.Sleep(time.Duration(5) * time.Millisecond)
	}
	log.Println("ErrorHandler: Tratamento de erros finalizado.")
}

