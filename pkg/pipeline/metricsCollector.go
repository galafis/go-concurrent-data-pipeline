// Go Concurrent Data Pipeline
// Author: Gabriel Demetrios Lafis
// Year: 2025

package pipeline

import (
	"log"
)

// MetricsCollector coleta e agrega métricas da pipeline.
// Retorna as métricas coletadas para permitir validação em testes.
func MetricsCollector(processedCh <-chan ProcessedRecord, errorCh <-chan DataRecord) Metrics {
	log.Println("MetricsCollector: Iniciando coleta de métricas...")
	metrics := Metrics{}

	// Usar um select para ler de múltiplos canais
	for processedCh != nil || errorCh != nil {
		select {
		case record, ok := <-processedCh:
			if !ok {
				processedCh = nil // Canal fechado
			} else {
				metrics.ProcessedCount++
				metrics.TotalValue += record.Value
				if record.IsAnomaly {
					metrics.AnomalyCount++
				}
				log.Printf("MetricsCollector: Coletando métrica para registro processado %s", record.ID)
			}
		case record, ok := <-errorCh:
			if !ok {
				errorCh = nil // Canal fechado
			} else {
				metrics.ErrorCount++
				log.Printf("MetricsCollector: Coletando métrica para registro com erro %s", record.ID)
			}
		}
	}

	log.Println("MetricsCollector: Coleta de métricas finalizada.")
	log.Printf("--- Sumário da Pipeline ---")
	log.Printf("Registros Processados com Sucesso: %d", metrics.ProcessedCount)
	log.Printf("Registros com Erro: %d", metrics.ErrorCount)
	log.Printf("Registros Anômalos: %d", metrics.AnomalyCount)
	log.Printf("Valor Total Processado: %.2f", metrics.TotalValue)
	log.Printf("---------------------------")
	
	return metrics
}

