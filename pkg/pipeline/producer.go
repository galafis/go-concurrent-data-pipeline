// Go Concurrent Data Pipeline
// Author: Gabriel Demetrios Lafis
// Year: 2025

package pipeline

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

// Producer gera registros de dados simulados.
func Producer(out chan<- DataRecord, numRecords int) {
	log.Println("Producer: Iniciando produção de dados...")
	locations := []string{"North", "South", "East", "West", "Center"}
	
	for i := 0; i < numRecords; i++ {
		record := DataRecord{
			ID:        fmt.Sprintf("rec-%04d", i),
			Value:     rand.Float64() * 100, // Valor entre 0 e 100
			Unit:      "unit_A",
			Timestamp: time.Now(),
			Status:    "raw",
			SensorID:  fmt.Sprintf("sensor-%d", (i%5)+1), // Simula 5 sensores diferentes
			Location:  locations[i%len(locations)],        // Rotaciona entre as localizações
		}
		// Simular alguns erros para teste
		if i%10 == 0 {
			record.Value = -1.0 // Valor inválido
		}
		if i%11 == 0 {
			record.Unit = "INVALID_UNIT" // Unidade inválida para transformação
		}
		out <- record
		log.Printf("Producer: Produzido %s (Value: %.2f, Sensor: %s, Location: %s)", 
			record.ID, record.Value, record.SensorID, record.Location)
		time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)
	}
	close(out)
	log.Println("Producer: Produção de dados finalizada.")
}

