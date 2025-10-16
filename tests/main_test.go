package tests

import (
	"encoding/json"
	"io"
	"math/rand"
	"os"
	"strings"
	"testing"
	"time"

	"go-concurrent-data-pipeline/pkg/pipeline"
)

func TestRunAdvancedPipeline(t *testing.T) {
	// Limpar arquivos de saída antes do teste
	_ = os.Remove("processed_data.jsonl")
	_ = os.Remove("failed_data.jsonl")

	numRecords := 50
	numWorkers := 3

	metrics := pipeline.RunAdvancedPipeline(numRecords, numWorkers)

	// Verificar se os arquivos de saída foram criados
	processedFile, err := os.Open("processed_data.jsonl")
	if err != nil {
		t.Fatalf("Falha ao abrir processed_data.jsonl: %v", err)
	}
	defer func() { _ = processedFile.Close() }()

	failedFile, err := os.Open("failed_data.jsonl")
	if err != nil {
		t.Fatalf("Falha ao abrir failed_data.jsonl: %v", err)
	}
	defer func() { _ = failedFile.Close() }()

	// Contar linhas nos arquivos de saída
	processedContent, _ := io.ReadAll(processedFile)
	processedLines := strings.Split(strings.TrimSpace(string(processedContent)), "\n")
	processedCount := 0
	if len(processedLines) > 0 && processedLines[0] != "" {
		processedCount = len(processedLines)
	}

	failedContent, _ := io.ReadAll(failedFile)
	failedLines := strings.Split(strings.TrimSpace(string(failedContent)), "\n")
	failedCount := 0
	if len(failedLines) > 0 && failedLines[0] != "" {
		failedCount = len(failedLines)
	}

	// Verificar as métricas retornadas
	if metrics.ProcessedCount != processedCount {
		t.Errorf("Contagem de métricas processadas (%d) não corresponde ao arquivo (%d)", 
			metrics.ProcessedCount, processedCount)
	}

	if metrics.ErrorCount != failedCount {
		t.Errorf("Contagem de métricas de erro (%d) não corresponde ao arquivo (%d)", 
			metrics.ErrorCount, failedCount)
	}

	// A lógica de contagem de registros: todos os registros devem ser processados ou ter falhado
	if processedCount+failedCount != numRecords {
		t.Errorf("Total de registros processados (%d) + falhados (%d) não é igual ao total produzido (%d)", 
			processedCount, failedCount, numRecords)
	}

	// Verificar se há registros anômalos no arquivo processado
	anomalyFound := false
	for _, line := range processedLines {
		if line == "" { continue }
		var record pipeline.ProcessedRecord
		err := json.Unmarshal([]byte(line), &record)
		if err != nil {
			t.Errorf("Erro ao deserializar registro processado: %v", err)
			continue
		}
		if record.IsAnomaly {
			anomalyFound = true
			break
		}
	}
	if !anomalyFound {
		t.Errorf("Nenhum registro anômalo encontrado no arquivo processed_data.jsonl. Isso pode indicar um problema na lógica de anomalia ou nos dados de teste.")
	}

	// Verificar se há registros com erro no arquivo failed
	errorRecordFound := false
	for _, line := range failedLines {
		if line == "" { continue }
		var record pipeline.DataRecord
		err := json.Unmarshal([]byte(line), &record)
		if err != nil {
			t.Errorf("Erro ao deserializar registro falhado: %v", err)
			continue
		}
		if record.Error != "" {
			errorRecordFound = true
			break
		}
	}
	if !errorRecordFound {
		t.Errorf("Nenhum registro com erro encontrado no arquivo failed_data.jsonl. Isso pode indicar um problema na lógica de erro ou nos dados de teste.")
	}

	// Verificar se as métricas de anomalia são corretas
	if metrics.AnomalyCount == 0 {
		t.Errorf("Nenhuma anomalia detectada nas métricas, esperado pelo menos uma")
	}

	// Limpar arquivos de saída após o teste
	_ = os.Remove("processed_data.jsonl")
	_ = os.Remove("failed_data.jsonl")
}

func TestMain(m *testing.M) {
	// Configurar o gerador de números aleatórios para resultados reproduzíveis nos testes
	rand.Seed(time.Now().UnixNano())

	// Executar os testes
	exitCode := m.Run()

	// Sair com o código de status dos testes
	os.Exit(exitCode)
}

