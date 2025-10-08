// Go Concurrent Data Pipeline
// Author: Gabriel Demetrios Lafis
// Year: 2025

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

// DataRecord representa um registro de dados com mais campos e complexidade.
type DataRecord struct {
	ID        string    `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	SensorID  string    `json:"sensor_id"`
	Value     float64   `json:"value"`
	Unit      string    `json:"unit"`
	Location  string    `json:"location"`
	Status    string    `json:"status"` // Adicionado para indicar status após processamento
	Error     string    `json:"error,omitempty"` // Para registrar erros específicos
}

// ProcessedRecord representa um registro após a transformação.
type ProcessedRecord struct {
	DataRecord
	ProcessedAt time.Time `json:"processed_at"`
	AnomalyScore float64   `json:"anomaly_score"`
	IsAnomaly    bool      `json:"is_anomaly"`
}

// producer gera registros de dados brutos.
func producer(dataCh chan<- DataRecord, wg *sync.WaitGroup, numRecords int) {
	defer wg.Done()
	log.Println("Producer: Iniciando produção de dados...")
	for i := 0; i < numRecords; i++ {
		record := DataRecord{
			ID:        fmt.Sprintf("rec-%04d", i),
			Timestamp: time.Now().Add(time.Duration(i) * time.Second),
			SensorID:  fmt.Sprintf("sensor-%d", rand.Intn(5)+1),
			Value:     rand.Float64()*100 + 10, // Valor entre 10 e 110
			Unit:      "Celsius",
			Location:  fmt.Sprintf("Room-%d", rand.Intn(3)+1),
			Status:    "raw",
		}
		// Simular alguns dados inválidos para teste de validação
		if i%7 == 0 {
			record.Value = -1.0 // Valor inválido
		}
		// Simular alguns dados que podem causar erro de processamento
		if i%11 == 0 {
			record.Unit = "INVALID_UNIT"
		}
		dataCh <- record
		log.Printf("Producer: Produzido %s (Value: %.2f)", record.ID, record.Value)
		time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)
	}
	close(dataCh)
	log.Println("Producer: Produção de dados finalizada.")
}

// validator valida os registros de dados.
func validator(in <-chan DataRecord, validCh chan<- DataRecord, errorCh chan<- DataRecord, wg *sync.WaitGroup) {
	defer wg.Done()
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
	close(validCh)
	log.Println("Validator: Validação de dados finalizada.")
}

// transformer transforma os registros de dados válidos.
func transformer(in <-chan DataRecord, out chan<- ProcessedRecord, errorCh chan<- DataRecord, wg *sync.WaitGroup) {
	defer wg.Done()
	for record := range in {
		log.Printf("Transformer: Transformando registro %s", record.ID)
		// Simular uma transformação mais complexa: cálculo de score de anomalia
		anomalyScore := record.Value * 0.1 // Exemplo simples
		isAnomaly := anomalyScore > 8.0   // Se o valor original for > 80

		// Simular erro de transformação
		if record.Unit == "INVALID_UNIT" {
			record.Status = "transformation_error"
			record.Error = "Invalid unit for transformation"
			errorCh <- record
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
	close(out)
	log.Println("Transformer: Transformação de dados finalizada.")
}

// loader carrega os registros processados para um destino (simulado).
func loader(in <-chan ProcessedRecord, wg *sync.WaitGroup) {
	defer wg.Done()
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

// errorHandler lida com registros que falharam em alguma etapa.
func errorHandler(errorCh <-chan DataRecord, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Println("ErrorHandler: Iniciando tratamento de erros...")
	file, err := os.Create("failed_data.jsonl")
	if err != nil {
		log.Fatalf("ErrorHandler: Falha ao criar arquivo de erros: %v", err)
	}
	defer file.Close()

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
	}
	log.Println("ErrorHandler: Tratamento de erros finalizado.")
}

// metricsCollector coleta e agrega métricas da pipeline.
func metricsCollector(processedCh <-chan ProcessedRecord, errorCh <-chan DataRecord, wg *sync.WaitGroup) {
	defer wg.Done()
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

func RunAdvancedPipeline(numRecords int, numWorkers int) {
	// Canais para comunicação entre as etapas
	dataCh := make(chan DataRecord, 100)       // Producer -> Validator
	validCh := make(chan DataRecord, 100)      // Validator -> Transformer
	processedCh := make(chan ProcessedRecord, 100) // Transformer -> Loader
	errorCh := make(chan DataRecord, 100)      // Erros de Validator ou Transformer -> ErrorHandler

	var wg sync.WaitGroup
	var workerWg sync.WaitGroup // Para esperar todos os workers (validator, transformer)

	// 1. Producer
	wg.Add(1)
	go producer(dataCh, &wg, numRecords)

	// 2. Validators
	for i := 0; i < numWorkers; i++ {
		workerWg.Add(1)
		go validator(dataCh, validCh, errorCh, &workerWg)
	}

	// 3. Transformers
	for i := 0; i < numWorkers; i++ {
		workerWg.Add(1)
		go transformer(validCh, processedCh, errorCh, &workerWg)
	}

	// Goroutine para fechar canais intermediários quando os workers terminarem
	go func() {
		workerWg.Wait()
		close(validCh)
		close(processedCh)
	}()

	// 4. Loader
	wg.Add(1)
	go loader(processedCh, &wg)

	// 5. Error Handler
	wg.Add(1)
	go errorHandler(errorCh, &wg)

	// 6. Metrics Collector
	// Para o metricsCollector, vamos criar canais temporários para evitar que ele consuma dos canais principais
	// antes que loader e errorHandler tenham a chance. Isso garante que todos os dados cheguem aos destinos corretos.
	metricsProcessedCh := make(chan ProcessedRecord, 100)
	metricsErrorCh := make(chan DataRecord, 100)

	// Reencaminhar dados para os canais de métricas
	go func() {
		for p := range processedCh {
			metricsProcessedCh <- p
		}
		close(metricsProcessedCh)
	}()

	go func() {
		for e := range errorCh {
			metricsErrorCh <- e
		}
		close(metricsErrorCh)
	}()

	wg.Add(1)
	go metricsCollector(metricsProcessedCh, metricsErrorCh, &wg)

	wg.Wait()
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	fmt.Println("===========================================")
	fmt.Println("Go Concurrent Data Pipeline")
	fmt.Println("===========================================")

	// Limpar arquivos de saída anteriores
	_ = os.Remove("processed_data.jsonl")
	_ = os.Remove("failed_data.jsonl")

	// Executar a pipeline com 50 registros e 3 workers para validação/transformação
	RunAdvancedPipeline(50, 3)

	fmt.Println("===========================================")
	fmt.Println("Pipeline completed!")
	fmt.Println("===========================================")
}

