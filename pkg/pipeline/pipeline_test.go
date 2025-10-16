package pipeline

import (
	"testing"
	"time"
)

func TestProducer(t *testing.T) {
	numRecords := 10
	dataCh := make(chan DataRecord, numRecords)
	
	go Producer(dataCh, numRecords)
	
	count := 0
	for record := range dataCh {
		count++
		if record.ID == "" {
			t.Errorf("Record ID should not be empty")
		}
		if record.SensorID == "" {
			t.Errorf("SensorID should not be empty, got: %s", record.SensorID)
		}
		if record.Location == "" {
			t.Errorf("Location should not be empty, got: %s", record.Location)
		}
		if record.Timestamp.IsZero() {
			t.Errorf("Timestamp should not be zero")
		}
	}
	
	if count != numRecords {
		t.Errorf("Expected %d records, got %d", numRecords, count)
	}
}

func TestValidator(t *testing.T) {
	dataCh := make(chan DataRecord, 10)
	validCh := make(chan DataRecord, 10)
	errorCh := make(chan DataRecord, 10)
	
	// Test valid record
	go func() {
		dataCh <- DataRecord{ID: "test-1", Value: 50.0, Unit: "unit_A", Timestamp: time.Now()}
		dataCh <- DataRecord{ID: "test-2", Value: -5.0, Unit: "unit_A", Timestamp: time.Now()} // Invalid
		dataCh <- DataRecord{ID: "test-3", Value: 1001.0, Unit: "unit_A", Timestamp: time.Now()} // Invalid
		close(dataCh)
	}()
	
	go func() {
		Validator(dataCh, validCh, errorCh)
		close(validCh)
		close(errorCh)
	}()
	
	validCount := 0
	errorCount := 0
	
	done := make(chan bool, 2)
	go func() {
		for range validCh {
			validCount++
		}
		done <- true
	}()
	
	go func() {
		for record := range errorCh {
			errorCount++
			if record.Error == "" {
				t.Errorf("Error field should not be empty for invalid record")
			}
		}
		done <- true
	}()
	
	<-done
	<-done
	
	if validCount != 1 {
		t.Errorf("Expected 1 valid record, got %d", validCount)
	}
	if errorCount != 2 {
		t.Errorf("Expected 2 error records, got %d", errorCount)
	}
}

func TestTransformer(t *testing.T) {
	validCh := make(chan DataRecord, 10)
	processedCh := make(chan ProcessedRecord, 10)
	errorCh := make(chan DataRecord, 10)
	
	// Test valid transformation
	go func() {
		validCh <- DataRecord{ID: "test-1", Value: 50.0, Unit: "unit_A", Timestamp: time.Now()}
		validCh <- DataRecord{ID: "test-2", Value: 90.0, Unit: "unit_A", Timestamp: time.Now()} // Anomaly
		validCh <- DataRecord{ID: "test-3", Value: 75.0, Unit: "INVALID_UNIT", Timestamp: time.Now()} // Transformation error
		close(validCh)
	}()
	
	go func() {
		Transformer(validCh, processedCh, errorCh)
		close(processedCh)
		close(errorCh)
	}()
	
	processedCount := 0
	anomalyCount := 0
	errorCount := 0
	
	done := make(chan bool, 2)
	go func() {
		for record := range processedCh {
			processedCount++
			if record.AnomalyScore <= 0 {
				t.Errorf("AnomalyScore should be positive, got %f", record.AnomalyScore)
			}
			if record.IsAnomaly {
				anomalyCount++
			}
		}
		done <- true
	}()
	
	go func() {
		for range errorCh {
			errorCount++
		}
		done <- true
	}()
	
	<-done
	<-done
	
	if processedCount != 2 {
		t.Errorf("Expected 2 processed records, got %d", processedCount)
	}
	if anomalyCount != 1 {
		t.Errorf("Expected 1 anomaly, got %d", anomalyCount)
	}
	if errorCount != 1 {
		t.Errorf("Expected 1 transformation error, got %d", errorCount)
	}
}

func TestMetricsCollector(t *testing.T) {
	processedCh := make(chan ProcessedRecord, 10)
	errorCh := make(chan DataRecord, 10)
	
	go func() {
		// Send 5 processed records (2 anomalies)
		for i := 0; i < 5; i++ {
			processedCh <- ProcessedRecord{
				DataRecord: DataRecord{
					ID:    "test-" + string(rune('0'+i)),
					Value: float64(i * 20),
				},
				IsAnomaly: i >= 3, // Last 2 are anomalies
			}
		}
		close(processedCh)
		
		// Send 3 error records
		for i := 0; i < 3; i++ {
			errorCh <- DataRecord{ID: "error-" + string(rune('0'+i))}
		}
		close(errorCh)
	}()
	
	metrics := MetricsCollector(processedCh, errorCh)
	
	if metrics.ProcessedCount != 5 {
		t.Errorf("Expected 5 processed records, got %d", metrics.ProcessedCount)
	}
	if metrics.ErrorCount != 3 {
		t.Errorf("Expected 3 error records, got %d", metrics.ErrorCount)
	}
	if metrics.AnomalyCount != 2 {
		t.Errorf("Expected 2 anomalies, got %d", metrics.AnomalyCount)
	}
	expectedTotal := 0.0 + 20.0 + 40.0 + 60.0 + 80.0
	if metrics.TotalValue != expectedTotal {
		t.Errorf("Expected total value %f, got %f", expectedTotal, metrics.TotalValue)
	}
}

func BenchmarkProducer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dataCh := make(chan DataRecord, 100)
		go Producer(dataCh, 100)
		for range dataCh {
			// Consume all records
		}
	}
}

func BenchmarkValidator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dataCh := make(chan DataRecord, 100)
		validCh := make(chan DataRecord, 100)
		errorCh := make(chan DataRecord, 100)
		
		go func() {
			for j := 0; j < 100; j++ {
				dataCh <- DataRecord{ID: "test", Value: float64(j), Unit: "unit_A", Timestamp: time.Now()}
			}
			close(dataCh)
		}()
		
		go Validator(dataCh, validCh, errorCh)
		
		for range validCh {
		}
		for range errorCh {
		}
	}
}
