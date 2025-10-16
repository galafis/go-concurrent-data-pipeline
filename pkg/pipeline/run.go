// Go Concurrent Data Pipeline
// Author: Gabriel Demetrios Lafis
// Year: 2025

package pipeline

import (
	"log"
	"sync"
)

func RunAdvancedPipeline(numRecords int, numWorkers int) Metrics {
	// Canais para comunicação entre as etapas
	dataCh := make(chan DataRecord, 100)       // Producer -> Validator
	validCh := make(chan DataRecord, 100)      // Validator -> Transformer
	processedCh := make(chan ProcessedRecord, 100) // Transformer -> Fan-out
	errorCh := make(chan DataRecord, 100)      // Erros de Validator ou Transformer -> Fan-out
	
	// Canais dedicados para cada consumidor
	loaderCh := make(chan ProcessedRecord, 100)
	errorHandlerCh := make(chan DataRecord, 100)
	metricsProcessedCh := make(chan ProcessedRecord, 100)
	metricsErrorCh := make(chan DataRecord, 100)

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
	
	// Fan-out para processedCh - distribui para Loader e MetricsCollector
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(loaderCh)
		defer close(metricsProcessedCh)
		for record := range processedCh {
			loaderCh <- record
			metricsProcessedCh <- record
		}
	}()
	
	// Fan-out para errorCh - distribui para ErrorHandler e MetricsCollector
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(errorHandlerCh)
		defer close(metricsErrorCh)
		for record := range errorCh {
			errorHandlerCh <- record
			metricsErrorCh <- record
		}
	}()

	// 4. Loader
	wg.Add(1)
	go func() {
		defer wg.Done()
		Loader(loaderCh)
	}()

	// 5. Error Handler
	wg.Add(1)
	go func() {
		defer wg.Done()
		ErrorHandler(errorHandlerCh)
	}()

	// 6. Metrics Collector
	var metrics Metrics
	wg.Add(1)
	go func() {
		defer wg.Done()
		metrics = MetricsCollector(metricsProcessedCh, metricsErrorCh)
	}()

	wg.Wait() // Espera todas as etapas da pipeline serem concluídas
	log.Println("Pipeline completed!")
	return metrics
}


