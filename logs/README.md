# Logs Directory

This directory is used for storing pipeline execution logs.

Log files are gitignored to prevent committing potentially sensitive operational data.

## Log Files

The pipeline may generate the following log files:
- `pipeline.log` - Main execution log
- `error.log` - Error-specific logs
- `metrics.log` - Performance metrics logs

## Configuration

Log behavior can be configured in `config/config.yaml`:
- Log level (DEBUG, INFO, WARN, ERROR)
- Log file locations
- Log rotation settings

## Viewing Logs

```bash
# View the most recent logs
tail -f logs/pipeline.log

# Search for errors
grep ERROR logs/pipeline.log

# View last 100 lines
tail -n 100 logs/pipeline.log
```
