// Go Concurrent Data Pipeline
// Author: Gabriel Demetrios Lafis
// Year: 2025

package pipeline

import (
	"log"
	"sync"
)

func RunAdvancedPipeline(numRecords int, numWorkers int) {
	// Canais para comunicação entre as etapas
	dataCh := make(chan DataRecord, 100)       // Producer -> Validator
	validCh := make(chan DataRecord, 100)      // Validator -> Transformer
	processedCh := make(chan ProcessedRecord, 100) // Transformer -> Loader
	errorCh := make(chan DataRecord, 100)      // Erros de Validator ou Transformer -> ErrorHandler

	var wg sync.WaitGroup // Main WaitGroup for all goroutines

	// WaitGroups para coordenar o fechamento dos canais
	var validatorWg sync.WaitGroup
	var transformerWg sync.WaitGroup
	var errorWg sync.WaitGroup // Para goroutines que escrevem em errorCh (Validators e Transformers)

	// 1. Producer
	wg.Add(1)
	go func() {
		defer wg.Done()
		Producer(dataCh, numRecords)
	}()

	// 2. Validators
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		validatorWg.Add(1)
		errorWg.Add(1) // Validators escrevem em errorCh
		go func() {
			defer wg.Done()
			defer validatorWg.Done()
			defer errorWg.Done()
			Validator(dataCh, validCh, errorCh)
		}()
	}

	// Goroutine para fechar validCh após todos os validators terminarem
	go func() {
		validatorWg.Wait()
		close(validCh)
	}()

	// 3. Transformers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		transformerWg.Add(1)
		errorWg.Add(1) // Transformers escrevem em errorCh
		go func() {
			defer wg.Done()
			defer transformerWg.Done()
			defer errorWg.Done()
			Transformer(validCh, processedCh, errorCh)
		}()
	}

	// Goroutine para fechar processedCh após todos os transformers terminarem
	go func() {
		transformerWg.Wait()
		close(processedCh)
	}()

	// Goroutine para fechar errorCh após todos os produtores de erro terminarem
	go func() {
		errorWg.Wait()
		close(errorCh)
	}()

	// 4. Loader
	wg.Add(1)
	go func() {
		defer wg.Done()
		Loader(processedCh)
	}()

	// 5. Error Handler
	wg.Add(1)
	go func() {
		defer wg.Done()
		ErrorHandler(errorCh)
	}()

	// 6. Metrics Collector
	wg.Add(1)
	go func() {
		defer wg.Done()
		MetricsCollector(processedCh, errorCh)
	}()

	wg.Wait() // Espera todas as etapas da pipeline serem concluídas
	log.Println("Pipeline completed!")
}

