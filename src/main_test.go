package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

// Mock para simular a saída do console
func captureOutput(f func()) string {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	f()
	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = oldStdout
	return string(out)
}

func TestProducerConsumerPipeline(t *testing.T) {
	// Testar com 3 workers e 10 itens de dados
	numWorkers := 3
	numData := 10

	output := captureOutput(func() {
		runPipeline(numWorkers, numData)
	})

	// Verificar se a saída contém as mensagens esperadas
	if !strings.Contains(output, "Pipeline completed!") {
		t.Errorf("Expected %q in output, got: %s", "Pipeline completed!", output)
	}

	// Verificar se todos os dados foram produzidos e consumidos
	for i := 1; i <= numData; i++ {
		expectedProduced := fmt.Sprintf("Produced: {%d Data-%d}", i, i)
		if !strings.Contains(output, expectedProduced) {
			t.Errorf("Expected %q in output, got: %s", expectedProduced, output)
		}

		expectedConsumed := fmt.Sprintf("Consumed: {%d Data-%d}", i, i)
		if !strings.Contains(output, expectedConsumed) {
			t.Errorf("Expected %q in output, got: %s", expectedConsumed, output)
		}
	}

	// Verificar se os workers processaram dados
	for i := 1; i <= numWorkers; i++ {
		expectedWorkerProcessing := fmt.Sprintf("Worker %d processing:", i)
		if !strings.Contains(output, expectedWorkerProcessing) {
			t.Errorf("Expected a worker to process data, but no output for: %s", expectedWorkerProcessing)
		}
	}

	// Testar com 0 workers para garantir que o produtor ainda funcione e feche o canal
	numWorkers = 0
	numData = 3

	output0 := captureOutput(func() {
		runPipeline(numWorkers, numData)
	})

	for i := 1; i <= numData; i++ {
		expectedProduced := fmt.Sprintf("Produced: {%d Data-%d}", i, i)
		if !strings.Contains(output0, expectedProduced) {
			t.Errorf("Producer failed with 0 workers: Expected %q in output, got: %s", expectedProduced, output0)
		}
	}

	if strings.Contains(output0, "Consumed:") {
		t.Errorf("Consumed data found with 0 workers: %s", output0)
	}
}

