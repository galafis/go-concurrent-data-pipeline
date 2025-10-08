package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"
)

// Helper function to capture stdout
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

// TestAdvancedPipelineEndToEnd simulates a full run of the advanced pipeline
// and verifies the output files and metrics.
func TestAdvancedPipelineEndToEnd(t *testing.T) {
	testDir := "./test_output"
	_ = os.RemoveAll(testDir) // Clean up previous test runs
	_ = os.Mkdir(testDir, 0755)

	processedFilePath := filepath.Join(testDir, "processed_data.jsonl")
	failedFilePath := filepath.Join(testDir, "failed_data.jsonl")

	// Override global output paths for testing
	originalProcessedOutput := processedOutput
	originalFailedOutput := failedOutput
	processedOutput = processedFilePath
	failedOutput = failedFilePath
	defer func() {
		processedOutput = originalProcessedOutput
		failedOutput = originalFailedOutput
		_ = os.RemoveAll(testDir)
	}()

	numRecords := 1000
	numWorkers := 5

	// Run the pipeline and capture its console output
	output := captureOutput(func() {
		RunAdvancedPipeline(numRecords, numWorkers)
	})

	// Verify processed_data.jsonl
	processedContent, err := ioutil.ReadFile(processedFilePath)
	if err != nil {
		t.Fatalf("Failed to read processed_data.jsonl: %v", err)
	}
	processedLines := strings.Split(strings.TrimSpace(string(processedContent)), "\n")
	processedCount := 0
	if len(processedLines) > 0 && processedLines[0] != "" {
		processedCount = len(processedLines)
	}

	// Verify failed_data.jsonl
	failedContent, err := ioutil.ReadFile(failedFilePath)
	if err != nil {
		t.Fatalf("Failed to read failed_data.jsonl: %v", err)
	}
	failedLines := strings.Split(strings.TrimSpace(string(failedContent)), "\n")
	failedCount := 0
	if len(failedLines) > 0 && failedLines[0] != "" {
		failedCount = len(failedLines)
	}

	// Check total records processed vs. failed
	if processedCount+failedCount != numRecords {
		t.Errorf("Expected total records (processed + failed) to be %d, got %d. Processed: %d, Failed: %d",
			numRecords, processedCount+failedCount, processedCount, failedCount)
	}

	// Check for specific metrics in the console output
	if !strings.Contains(output, "--- Sumário da Pipeline ---") {
		t.Errorf("Pipeline summary not found in output.\nOutput:\n%s", output)
	}
	if !strings.Contains(output, fmt.Sprintf("Registros Processados: %d", processedCount)) {
		t.Errorf("Processed count mismatch in summary. Expected %d, Output:\n%s", processedCount, output)
	}
	if !strings.Contains(output, fmt.Sprintf("Registros Falhados: %d", failedCount)) {
		t.Errorf("Failed count mismatch in summary. Expected %d, Output:\n%s", failedCount, output)
	}
	if !strings.Contains(output, "Pipeline completed!") {
		t.Errorf("Pipeline completion message not found.\nOutput:\n%s", output)
	}

	// Check content of processed records for transformations and anomaly detection
	anomalyCount := 0
	for _, line := range processedLines {
		if line == "" { continue }
		var record ProcessedRecord
		err := json.Unmarshal([]byte(line), &record)
		if err != nil {
			t.Errorf("Error unmarshalling processed record: %v", err)
			continue
		}
		// Basic check for transformation: value should be squared
		if record.Value != record.OriginalValue*record.OriginalValue {
			t.Errorf("Processed record value not squared: %+v", record)
		}
		// Check for anomaly flag
		if record.IsAnomaly {
			anomalyCount++
		}
	}
	// Since anomaly detection is random, we expect some anomalies but not a fixed number.
	// Just ensure it's not zero if numRecords is large enough.
	if numRecords > 100 && anomalyCount == 0 {
		t.Errorf("Expected some anomalies for %d records, but found 0.", numRecords)
	}

	// Check content of failed records for error messages
	for _, line := range failedLines {
		if line == "" { continue }
		var record DataRecord
		err := json.Unmarshal([]byte(line), &record)
		if err != nil {
			t.Errorf("Error unmarshalling failed record: %v", err)
			continue
		}
		if record.Error == "" {
			t.Errorf("Failed record has no error message: %+v", record)
		}
	}
}

// TestMetricsCollector ensures metrics are correctly aggregated and reported.
func TestMetricsCollector(t *testing.T) {
	mc := newMetricsCollector()

	mc.recordProcessed()
	mc.recordProcessed()
	mc.recordFailed("validation error")
	mc.recordProcessed()
	mc.recordFailed("processing error")
	mc.recordFailed("another error")

	mc.wg.Wait() // Ensure all increments are done

	if mc.processedCount.Load() != 3 {
		t.Errorf("Expected 3 processed records, got %d", mc.processedCount.Load())
	}
	if mc.failedCount.Load() != 3 {
		t.Errorf("Expected 3 failed records, got %d", mc.failedCount.Load())
	}

	// Check error counts
	mc.mu.Lock()
	if mc.errorCounts["validation error"] != 1 {
		t.Errorf("Expected 1 'validation error', got %d", mc.errorCounts["validation error"])
	}
	if mc.errorCounts["processing error"] != 1 {
		t.Errorf("Expected 1 'processing error', got %d", mc.errorCounts["processing error"])
	}
	if mc.errorCounts["another error"] != 1 {
		t.Errorf("Expected 1 'another error', got %d", mc.errorCounts["another error"])
	}
	mc.mu.Unlock()

	// Test reporting
	output := captureOutput(func() {
		mc.report()
	})

	if !strings.Contains(output, "Registros Processados: 3") {
		t.Errorf("Report missing processed count: %s", output)
	}
	if !strings.Contains(output, "Registros Falhados: 3") {
		t.Errorf("Report missing failed count: %s", output)
	}
	if !strings.Contains(output, "validation error: 1") {
		t.Errorf("Report missing 'validation error' count: %s", output)
	}
}

// TestDataGenerator ensures data generation produces expected number of records.
func TestDataGenerator(t *testing.T) {
	numRecords := 100
	dataChan := make(chan DataRecord, numRecords)
	var wg sync.WaitGroup

	wg.Add(1)
	go generateData(numRecords, dataChan, &wg)

	wg.Wait()
	close(dataChan)

	count := 0
	for range dataChan {
		count++
	}

	if count != numRecords {
		t.Errorf("Expected %d records, got %d", numRecords, count)
	}
}

// TestProcessRecordWorker ensures workers process records correctly.
func TestProcessRecordWorker(t *testing.T) {
	inputChan := make(chan DataRecord, 1)
	processedChan := make(chan ProcessedRecord, 1)
	failedChan := make(chan DataRecord, 1)
	mc := newMetricsCollector()
	var wg sync.WaitGroup

	wg.Add(1)
	go processRecordWorker(inputChan, processedChan, failedChan, mc, &wg)

	// Test a valid record
	validRecord := DataRecord{ID: "1", Value: 10, Category: "A"}
	inputChan <- validRecord
	close(inputChan)

	wg.Wait()
	close(processedChan)
	close(failedChan)

	processedCount := 0
	for pr := range processedChan {
		processedCount++
		if pr.Value != 100 { // 10 * 10
			t.Errorf("Expected processed value to be 100, got %f", pr.Value)
		}
	}
	if processedCount != 1 {
		t.Errorf("Expected 1 processed record, got %d", processedCount)
	}

	// Reset for failed record test
	inputChan = make(chan DataRecord, 1)
	processedChan = make(chan ProcessedRecord, 1)
	failedChan = make(chan DataRecord, 1)
	mc = newMetricsCollector()
	wg.Add(1)
	go processRecordWorker(inputChan, processedChan, failedChan, mc, &wg)

	// Test a record that should fail validation (Value < 0)
	failedRecord := DataRecord{ID: "2", Value: -5, Category: "B"}
	inputChan <- failedRecord
	close(inputChan)

	wg.Wait()
	close(processedChan)
	close(failedChan)

	failedCount := 0
	for fr := range failedChan {
		failedCount++
		if !strings.Contains(fr.Error, "Value cannot be negative") {
			t.Errorf("Expected 'Value cannot be negative' error, got %s", fr.Error)
		}
	}
	if failedCount != 1 {
		t.Errorf("Expected 1 failed record, got %d", failedCount)
	}
}

// TestWriteProcessedRecords ensures processed records are written correctly.
func TestWriteProcessedRecords(t *testing.T) {
	testDir := "./test_write_processed"
	_ = os.RemoveAll(testDir)
	_ = os.Mkdir(testDir, 0755)
	filePath := filepath.Join(testDir, "processed.jsonl")

	processedChan := make(chan ProcessedRecord, 2)
	var wg sync.WaitGroup

	wg.Add(1)
	go writeProcessedRecords(processedChan, filePath, &wg)

	processedChan <- ProcessedRecord{ID: "1", OriginalValue: 5, Value: 25, IsAnomaly: false, Timestamp: time.Now()}
	processedChan <- ProcessedRecord{ID: "2", OriginalValue: 10, Value: 100, IsAnomaly: true, Timestamp: time.Now()}
	close(processedChan)

	wg.Wait()

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read processed file: %v", err)
	}
	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	if len(lines) != 2 {
		t.Errorf("Expected 2 lines, got %d", len(lines))
	}

	var r1, r2 ProcessedRecord
	_ = json.Unmarshal([]byte(lines[0]), &r1)
	_ = json.Unmarshal([]byte(lines[1]), &r2)

	if r1.ID != "1" || r2.ID != "2" {
		t.Errorf("Records not written correctly: %v, %v", r1, r2)
	}
	_ = os.RemoveAll(testDir)
}

// TestWriteFailedRecords ensures failed records are written correctly.
func TestWriteFailedRecords(t *testing.T) {
	testDir := "./test_write_failed"
	_ = os.RemoveAll(testDir)
	_ = os.Mkdir(testDir, 0755)
	filePath := filepath.Join(testDir, "failed.jsonl")

	failedChan := make(chan DataRecord, 2)
	var wg sync.WaitGroup

	wg.Add(1)
	go writeFailedRecords(failedChan, filePath, &wg)

	failedChan <- DataRecord{ID: "3", Value: -1, Error: "negative value"}
	failedChan <- DataRecord{ID: "4", Value: 0, Error: "zero value"}
	close(failedChan)

	wg.Wait()

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read failed file: %v", err)
	}
	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	if len(lines) != 2 {
		t.Errorf("Expected 2 lines, got %d", len(lines))
	}

	var r1, r2 DataRecord
	_ = json.Unmarshal([]byte(lines[0]), &r1)
	_ = json.Unmarshal([]byte(lines[1]), &r2)

	if r1.ID != "3" || r2.ID != "4" {
		t.Errorf("Records not written correctly: %v, %v", r1, r2)
	}
	_ = os.RemoveAll(testDir)
}

