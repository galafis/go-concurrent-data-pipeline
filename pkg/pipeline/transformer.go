// Go Concurrent Data Pipeline
// Author: Gabriel Demetrios Lafis
// Year: 2025

package pipeline

import (
	"log"
	"math/rand"
	"time"
)

// Transformer transforma os registros de dados válidos.
func Transformer(in <-chan DataRecord, out chan<- ProcessedRecord, errCh chan<- DataRecord) {
	for record := range in {
		log.Printf("Transformer: Transformando registro %s", record.ID)
		// Simular uma transformação mais complexa: cálculo de score de anomalia
		anomalyScore := record.Value * 0.1 // Exemplo simples
		isAnomaly := anomalyScore > 8.0   // Se o valor original for > 80

		// Simular erro de transformação
		if record.Unit == "INVALID_UNIT" {
			record.Status = "transformation_error"
			record.Error = "Invalid unit for transformation"
			errCh <- record
			log.Printf("Transformer: Erro ao transformar registro %s (Invalid Unit)", record.ID)
			continue
		}

		processedRecord := ProcessedRecord{
			DataRecord:  record,
			ProcessedAt: time.Now(),
			AnomalyScore: anomalyScore,
			IsAnomaly:    isAnomaly,
		}
		processedRecord.Status = "processed"
		out <- processedRecord
		log.Printf("Transformer: Registro %s transformado (AnomalyScore: %.2f)", record.ID, anomalyScore)
		time.Sleep(time.Duration(rand.Intn(30)) * time.Millisecond)
	}
	log.Println("Transformer: Transformação de dados finalizada.")
}

