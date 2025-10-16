# Contributing to Go Concurrent Data Pipeline

Thank you for your interest in contributing! This document provides guidelines for contributing to this project.

## Code of Conduct

Be respectful and inclusive. We welcome contributions from everyone.

## How to Contribute

### Reporting Issues

- Check if the issue already exists
- Provide detailed reproduction steps
- Include Go version and OS information
- Share relevant log output

### Suggesting Features

- Explain the use case clearly
- Describe the proposed solution
- Consider backward compatibility
- Be open to alternative approaches

### Submitting Pull Requests

1. **Fork the repository**
2. **Create a feature branch**
   ```bash
   git checkout -b feature/my-new-feature
   ```
3. **Make your changes**
   - Write clear, idiomatic Go code
   - Follow existing code style
   - Add tests for new functionality
   - Update documentation as needed
4. **Run tests**
   ```bash
   go test ./...
   ```
5. **Run linter**
   ```bash
   golangci-lint run
   ```
6. **Commit your changes**
   - Write clear, descriptive commit messages
   - Reference relevant issues
7. **Push to your fork**
8. **Open a Pull Request**

## Development Guidelines

### Code Style

- Follow [Effective Go](https://go.dev/doc/effective_go)
- Use `gofmt` to format code
- Run `golangci-lint` before committing
- Keep functions small and focused
- Write self-documenting code with clear variable names

### Testing

- Write unit tests for all new functions
- Maintain or improve code coverage
- Use table-driven tests where appropriate
- Test both success and failure cases
- Use meaningful test names

### Documentation

- Add godoc comments for exported functions and types
- Update README.md for user-facing changes
- Update ARCHITECTURE.md for design changes
- Include code examples where helpful

### Commit Messages

Follow conventional commits format:

```
type(scope): subject

body

footer
```

Types: `feat`, `fix`, `docs`, `style`, `refactor`, `test`, `chore`

Example:
```
feat(transformer): add support for custom anomaly thresholds

Added configuration option to customize anomaly detection threshold.
Previously hardcoded at 8.0, now configurable per deployment.

Closes #123
```

## Project Structure

```
.
├── src/              # Main application entry point
├── pkg/pipeline/     # Core pipeline implementation
├── tests/            # Integration tests
├── docs/             # Documentation
├── data/             # Sample data
├── config/           # Configuration files
└── .github/          # GitHub Actions workflows
```

## Testing Your Changes

### Unit Tests
```bash
go test ./pkg/pipeline -v
```

### Integration Tests
```bash
go test ./tests -v
```

### Benchmarks
```bash
go test ./pkg/pipeline -bench=. -benchmem
```

### Coverage
```bash
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## Getting Help

- Open an issue for questions
- Check existing documentation
- Review closed issues for similar problems

## License

By contributing, you agree that your contributions will be licensed under the MIT License.
