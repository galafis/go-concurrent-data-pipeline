package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"
)

// captureOutput é uma função auxiliar para capturar a saída do console.
func captureOutput(f func()) string {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	f()
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = oldStdout
	return string(out)
}

func TestRunAdvancedPipeline(t *testing.T) {
	// Limpar arquivos de saída antes do teste
	_ = os.Remove("processed_data.jsonl")
	_ = os.Remove("failed_data.jsonl")

	numRecords := 50
	numWorkers := 3

	output := captureOutput(func() {
		RunAdvancedPipeline(numRecords, numWorkers)
	})

	// Verificar se os arquivos de saída foram criados
	processedFile, err := os.Open("processed_data.jsonl")
	if err != nil {
		t.Fatalf("Falha ao abrir processed_data.jsonl: %v", err)
	}
	defer processedFile.Close()

	failedFile, err := os.Open("failed_data.jsonl")
	if err != nil {
		t.Fatalf("Falha ao abrir failed_data.jsonl: %v", err)
	}
	defer failedFile.Close()

	// Contar linhas nos arquivos de saída
	processedContent, _ := ioutil.ReadAll(processedFile)
	processedLines := strings.Split(strings.TrimSpace(string(processedContent)), "\n")
	processedCount := 0
	if len(processedLines) > 0 && processedLines[0] != "" {
		processedCount = len(processedLines)
	}

	failedContent, _ := ioutil.ReadAll(failedFile)
	failedLines := strings.Split(strings.TrimSpace(string(failedContent)), "\n")
	failedCount := 0
	if len(failedLines) > 0 && failedLines[0] != "" {
		failedCount = len(failedLines)
	}

	// Verificar o sumário de métricas na saída do console
	if !strings.Contains(output, "--- Sumário da Pipeline ---") {
		t.Errorf("Sumário da pipeline não encontrado na saída.\nOutput:\n%s", output)
	}

	expectedProcessed := numRecords - (numRecords/7) - (numRecords/11) // Aproximado, pois alguns podem falhar em ambas as etapas
	if processedCount+failedCount != numRecords {
		t.Errorf("Total de registros processados (%d) + falhados (%d) não é igual ao total produzido (%d). Output: %s", processedCount, failedCount, numRecords, output)
	}

	// Verificar se há registros anômalos no arquivo processado
	anomalyFound := false
	for _, line := range processedLines {
		if line == "" { continue }
		var record ProcessedRecord
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
		var record DataRecord
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

	// Verificar se a pipeline foi concluída com sucesso
	if !strings.Contains(output, "Pipeline completed!") {
		t.Errorf("Mensagem de conclusão da pipeline não encontrada.\nOutput:\n%s", output)
	}

	// Limpar arquivos de saída após o teste
	_ = os.Remove("processed_data.jsonl")
	_ = os.Remove("failed_data.jsonl")
}

func TestMain(m *testing.M) {
	// Configurar o gerador de números aleatórios para resultados reproduzíveis nos testes
	rand.Seed(1)

	// Executar os testes
	exitCode := m.Run()

	// Sair com o código de status dos testes
	os.Exit(exitCode)
}

