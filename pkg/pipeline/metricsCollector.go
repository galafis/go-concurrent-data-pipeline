// Go Concurrent Data Pipeline
// Author: Gabriel Demetrios Lafis
// Year: 2025

package pipeline

import (
	"log"
)

// MetricsCollector coleta e agrega métricas da pipeline.
func MetricsCollector(processedCh <-chan ProcessedRecord, errorCh <-chan DataRecord) {
	log.Println("MetricsCollector: Iniciando coleta de métricas...")
	processedCount := 0
	errorCount := 0
	anomalyCount := 0
	totalValue := 0.0

	// Usar um select para ler de múltiplos canais
	for {
		select {
		case record, ok := <-processedCh:
			if !ok {
				processedCh = nil // Canal fechado
				break
			}
			processedCount++
			totalValue += record.Value
			if record.IsAnomaly {
				anomalyCount++
			}
			log.Printf("MetricsCollector: Coletando métrica para registro processado %s", record.ID)
		case record, ok := <-errorCh:
			if !ok {
				errorCh = nil // Canal fechado
				break
			}
			errorCount++
			log.Printf("MetricsCollector: Coletando métrica para registro com erro %s", record.ID)
		}

		if processedCh == nil && errorCh == nil {
			break // Ambos os canais estão fechados
		}
	}

	log.Println("MetricsCollector: Coleta de métricas finalizada.")
	log.Printf("--- Sumário da Pipeline ---")
	log.Printf("Registros Processados com Sucesso: %d", processedCount)
	log.Printf("Registros com Erro: %d", errorCount)
	log.Printf("Registros Anômalos: %d", anomalyCount)
	log.Printf("Valor Total Processado: %.2f", totalValue)
	log.Printf("---------------------------")
}

