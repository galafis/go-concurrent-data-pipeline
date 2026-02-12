# Sample Data

This directory contains sample data files for testing and demonstration purposes.

## Files

### `sample_input.jsonl`

Contains 10 sample data records in JSONL format. Each record has:
- `id`: Unique identifier
- `timestamp`: ISO 8601 timestamp
- `sensor_id`: Sensor identifier (sensor-1 through sensor-5)
- `value`: Numeric value (0-100)
- `unit`: Measurement unit (unit_A)
- `location`: Geographic location (North, South, East, West, Center)
- `status`: Record status (raw, processed, etc.)

## Usage

These sample files serve as reference for the data format the pipeline produces internally:

```bash
# Run the pipeline (it generates data programmatically)
go run src/main.go
```

## Data Format

All data files use JSONL (JSON Lines) format:
- One JSON object per line
- No commas between objects
- Easy to stream and process incrementally
- Compatible with many data processing tools

## Creating Your Own Data

To create your own test data:

1. Follow the same JSON schema
2. Use one JSON object per line
3. Ensure all required fields are present
4. Save with `.jsonl` extension

Example:
```json
{"id":"custom-001","timestamp":"2025-01-15T12:00:00Z","sensor_id":"sensor-1","value":42.0,"unit":"unit_A","location":"North","status":"raw"}
```
