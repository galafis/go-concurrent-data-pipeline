// Go Concurrent Data Pipeline
// Author: Gabriel Demetrios Lafis
// Year: 2025

package pipeline

import (
	"log"
	"math/rand"
	"time"
)

// Validator valida os registros de dados.
func Validator(in <-chan DataRecord, validCh chan<- DataRecord, errorCh chan<- DataRecord) {
	for record := range in {
		log.Printf("Validator: Validando registro %s", record.ID)
		if record.Value < 0 || record.Value > 1000 { // Exemplo de regra de validação
			record.Status = "invalid"
			record.Error = "Value out of expected range (0-1000)"
			errorCh <- record
			log.Printf("Validator: Registro %s inválido (Value: %.2f)", record.ID, record.Value)
		} else {
			validCh <- record
			log.Printf("Validator: Registro %s válido", record.ID)
		}
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
	}
	log.Println("Validator: Validação de dados finalizada.")
}

