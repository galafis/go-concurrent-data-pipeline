// Go Concurrent Data Pipeline
// Author: Gabriel Demetrios Lafis
// Year: 2025

package pipeline

import (
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

