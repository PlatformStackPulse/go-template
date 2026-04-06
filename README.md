# Go Template

![Go Version](https://img.shields.io/badge/Go-1.22+-blue?style=flat-square&logo=go)
![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg?style=flat-square)
![CI Status](https://github.com/PlatformStackPulse/go-template/actions/workflows/ci.yml/badge.svg)
[![codecov](https://codecov.io/gh/PlatformStackPulse/go-template/branch/main/graph/badge.svg)](https://codecov.io/gh/PlatformStackPulse/go-template)
![DevContainer](https://img.shields.io/static/v1?label=DevContainer&message=Ready&color=blue&style=flat-square&logo=visual-studio-code)

<p align="center">
  <strong>Production-Ready Go Template Repository</strong><br>
  Enterprise-grade CI/CD, DevSecOps, GitOps workflow, clean architecture, and developer-first tooling.
</p>

---

## 🎯 Overview

This repository is a **standardized, reusable GitHub template** for creating production-ready Go applications. It provides everything needed to build secure, scalable, single-binary command-line tools and platform utilities.

### ✨ Highlights

- ✅ **Clean Architecture** — Domain, Usecase, and Adapter layers
- ✅ **CLI Framework** — Cobra-based command structure
- ✅ **Feature Flags** — Built-in feature management system
- ✅ **Structured Logging** — Using Go's standard `slog` package
- ✅ **Comprehensive Testing** — Unit, integration, and example tests
- ✅ **DevSecOps** — Integrated security scanning (gosec, govulncheck, CodeQL)
- ✅ **CI/CD Automation** — GitHub Actions with linting, testing, coverage, and releases
- ✅ **Semantic Versioning** — Automatic version management and release automation
- ✅ **Conventional Commits** — Enforced commit message format
- ✅ **DevContainer** — Pre-configured development environment
- ✅ **Multi-Platform Builds** — macOS, Linux, Windows support
- ✅ **Docker Support** — Multi-stage Dockerfile and docker-compose
- ✅ **Infrastructure as Code** — Terraform and Kubernetes examples
- ✅ **Code Quality** — golangci-lint, go vet, coverage enforcement

---

## 🚀 Quick Start

### Using as a GitHub Template

```bash
# Create a new repository using this template
gh repo create my-go-app --template go-template

# Clone and navigate
git clone https://github.com/PlatformStackPulse/my-go-app
cd my-go-app

# Install development tools
make dev-setup

# Run tests
make test

# Build
make build

# Run
./bin/go-template hello --name "World"
```

### Local Development

```bash
# Setup development environment
make dev-setup

# Run in watch mode (requires air)
make watch

# Tests with coverage
make coverage

# Security scan
make security

# Format and lint
make fmt lint
```

### Using DevContainer

```bash
# Open in DevContainer (VS Code)
# Press Ctrl+Shift+P and select "Dev Containers: Reopen in Container"

# Inside container:
make install
make test
make build
```

---

## 📁 Project Structure

```
go-template/
├── cmd/
│   └── app/
│       └── main.go             # Application entry point
├── internal/                   # Private packages
│   ├── adapter/               # External integrations (HTTP, DB, etc.)
│   ├── cli/                   # Cobra commands
│   ├── config/                # Configuration management
│   ├── domain/                # Domain entities
│   ├── feature/               # Feature flags system
│   ├── logger/                # Structured logging
│   └── usecase/               # Business logic
├── pkg/                       # Public packages
│   ├── health/                # Health check handlers
│   └── version/               # Version information
├── test/
│   ├── integration/           # Integration tests
│   └── unit/                  # Unit tests
├── deploy/
│   ├── kubernetes/            # Kubernetes manifests
│   └── terraform/             # Terraform IaC
├── scripts/                   # Build and utility scripts
├── .devcontainer/             # DevContainer config
├── .github/
│   ├── workflows/             # GitHub Actions
│   ├── ISSUE_TEMPLATE/        # Issue templates
│   └── pull_request_template.md
├── .golangci.yml              # Linter configuration
├── Dockerfile                 # Multi-stage build
├── Makefile                   # Build automation
├── go.mod / go.sum            # Dependency management
└── README.md
```

---

## 🔄 Architecture

### Clean Architecture Pattern

```
┌─────────────────────────────────────┐
│           CLI Layer (Cobra)         │
├─────────────────────────────────────┤
│     Adapters (Ports & Adapters)     │
├─────────────────────────────────────┤
│      Usecases (Business Logic)      │
├─────────────────────────────────────┤
│       Domain (Entities & Rules)     │
└─────────────────────────────────────┘
```

### Design Patterns Included

- **Command Pattern** — Cobra commands
- **Mediator Pattern** — Feature manager
- **Orchestrator Pattern** — Usecase layer
- **Dependency Injection** — Idiomatic Go

---

## 🛠 Available Commands

### Makefile Targets

```bash
make help              # Show all available targets
make build             # Build the application
make run               # Build and run
make test              # Run all tests with coverage
make test-unit         # Run unit tests only
make test-integration  # Run integration tests only
make coverage          # Generate coverage report
make clean             # Remove build artifacts
make install           # Install dependencies
make lint              # Run linters
make fmt               # Format code
make vet               # Run go vet
make security          # Run security checks
make sec-update        # Check for security updates
make dev-setup         # Setup development environment
make changelog         # Regenerate CHANGELOG.md from commits
make changelog-check   # Verify CHANGELOG.md is current
make watch             # Watch for changes (air)
make version           # Show version info
make all               # Run all tasks
```

### CLI Commands

```bash
./bin/go-template --help
./bin/go-template hello                    # Print "Hello, World!"
./bin/go-template hello --name Alice       # Print "Hello, Alice!"
./bin/go-template version                  # Show version information
```

---

## 🧪 Testing

### Running Tests

```bash
# All tests with coverage
make test

# Specific test suite
make test-unit
make test-integration

# Generate HTML coverage report
make coverage

# Run with race detector
go test -race ./...
```

### Test Coverage

- **Minimum threshold:** 70%
- **Enforced in CI:** Yes
- **Table-driven tests:** Domain and usecase layers
- **Mock patterns:** Via interfaces

---

## 🔐 Security & Quality

### Integrated Security Tools

| Tool | Purpose | Status |
|------|---------|--------|
| `golangci-lint` | Code quality & style | ✅ CI/CD |
| `gosec` | Security issues | ✅ CI/CD |
| `govulncheck` | Dependency vulnerabilities | ✅ CI/CD |
| `CodeQL` | Static analysis | ✅ CI/CD |
| `Dependabot` | Dependency updates | ✅ Automated |

### Security Checks

```bash
# Run all security scans
make security

# Check for vulnerabilities
make sec-update

# Manual gosec scan
gosec ./...

# Manual govulncheck
govulncheck ./...
```

---

## 📦 Versioning & Releases

### Semantic Versioning

This project follows [Semantic Versioning](https://semver.org/):

```
MAJOR.MINOR.PATCH
  ↑      ↑      ↑
  │      │      └─ Bug fixes
  │      └────────── New features (backward compatible)
  └───────────────── Breaking changes
```

### Conventional Commits

Commit messages must follow the [Conventional Commits](https://www.conventionalcommits.org/) specification:

```
<type>(<scope>): <description>

<body>

<footer>
```

**Types:**
- `feat` — New feature
- `fix` — Bug fix
- `docs` — Documentation
- `style` — Code style
- `refactor` — Code refactoring
- `perf` — Performance improvement
- `test` — Test additions/updates
- `chore` — Build, dependencies, etc.

**Examples:**
```
feat: add hello command with name parameter
fix: resolve panic in logger initialization
docs: update README with quick start guide
chore: upgrade Go from 1.21 to 1.22
```

### Automated Release Process

1. **Create and push tag:**
   ```bash
   git tag v1.2.3
   git push origin v1.2.3
   ```

2. **Release workflow triggers:**
   - Multi-platform builds
   - GitHub Release creation
   - Docker image push
   - SBOM generation

3. **Artifacts available at:** `https://github.com/PlatformStackPulse/my-app/releases/tag/v1.2.3`

### Changelog Strategy

- `CHANGELOG.md` is generated from Conventional Commits using `git-chglog`.
- On merges/pushes to `main`, `.github/workflows/changelog.yml` updates and commits `CHANGELOG.md` automatically.
- Maintainers can regenerate locally with:

```bash
make changelog
```

---

## 🌟 Feature Flags

### Quick Example

```go
fm := feature.NewManager()

if fm.IsEnabled(feature.FeatureHello) {
    fmt.Println("Feature is enabled!")
}
```

### Define Flags

Edit `internal/feature/flags.go`:

```go
const (
    FeatureNewAPI    Flag = "feature_new_api"
    FeatureMetrics   Flag = "feature_metrics"
)
```

### Enable Flags

```bash
# Environment variable format
export FEATURE_FLAGS="feature_new_api=true,feature_metrics=false"
./bin/go-template hello
```

### Future Extensions

- Config file support (YAML/JSON)
- Remote flag service integration
- Gradual rollout (percentage)
- User-based targeting
- A/B testing support

---

## 🐳 Docker

### Build Image

```bash
# Build locally
docker build -t go-template:latest .

# Run container
docker run --rm go-template:latest hello

# Run with feature flags
docker run --rm \
  -e FEATURE_FLAGS="feature_hello=true" \
  go-template:latest hello
```

### Docker Compose

```bash
# Development
docker-compose up dev

# Run tests
docker-compose run test

# Run linter
docker-compose run lint
```

---

## 🚀 CI/CD Pipelines

### GitHub Actions Workflows

#### 1. **CI Pipeline** (`.github/workflows/ci.yml`)
Runs on every push and PR:
- Code linting (golangci-lint)
- Format verification
- Unit & integration tests
- Coverage enforcement (≥70%)
- Security scans (gosec, CodeQL)
- PR commit message validation
- Multi-version Go testing (1.21, 1.22)

#### 2. **Release Pipeline** (`.github/workflows/release.yml`)
Triggers on version tags (`v*.*.*`):
- Multi-platform builds (Linux, macOS, Windows)
- Checksum generation
- GitHub Release creation with artifacts
- Docker image build and push
- SBOM generation

#### 3. **Dependency Management** (`.github/workflows/dependencies.yml`)
Weekly automated tasks:
- Dependency updates
- Security vulnerability scans
- Automated PR creation

---

## 📚 Infrastructure Templates

### Kubernetes

Deploy to Kubernetes:

```bash
kubectl apply -f deploy/kubernetes/

# Check deployment
kubectl get pods -l app=go-template
```

### Terraform

Deploy to AWS Lambda:

```bash
cd deploy/terraform

terraform init
terraform plan -var-file=terraform.dev.tfvars
terraform apply -var-file=terraform.dev.tfvars
```

---

## 📖 Configuration

### Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `DEBUG` | Enable debug logging | `false` |
| `APP_NAME` | Application name | `go-template` |
| `APP_VERSION` | Application version | `dev` |
| `FEATURE_FLAGS` | Enabled feature flags | "" |

### Example `.env`

```env
DEBUG=true
APP_NAME=my-app
APP_VERSION=v1.0.0
FEATURE_FLAGS=feature_hello=true,feature_metrics=true
```

---

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/awesome-feature`
3. Commit with conventional format: `git commit -m "feat: add awesome feature"`
4. Push to branch: `git push origin feature/awesome-feature`
5. Open a Pull Request

### PR Requirements

- ✅ Must follow Conventional Commits
- ✅ Tests passing (coverage ≥ 70%)
- ✅ Code linting passing
- ✅ No security warnings
- ✅ Documentation updated

---

## 📋 Checklist for New Projects

When using this template:

- [ ] Update `go.mod` module path
- [ ] Update `module path in code imports`
- [ ] Create `.github/CODEOWNERS` file
- [ ] Update README with project-specific info
- [ ] Configure branch protection rules
- [ ] Setup Dependabot alerts
- [ ] Update repository description
- [ ] Add repository topics
- [ ] Customize GitHub Actions (if needed)
- [ ] Update Terraform variables for your environment

---

## 📜 License

This project is licensed under the MIT License — see [LICENSE](LICENSE) file for details.

---

## 🙏 Acknowledgments

Built with best practices from:
- [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [Conventional Commits](https://www.conventionalcommits.org/)
- [Semantic Versioning](https://semver.org/)
- [Go Best Practices](https://golang.org/doc/effective_go)
- [GoReleaser](https://goreleaser.com/)
- [Cobra](https://cobra.dev/)

---

## 📞 Support

- 📖 [Documentation](https://github.com/PlatformStackPulse/go-template/wiki)
- 🐛 [Report Issues](https://github.com/PlatformStackPulse/go-template/issues)
- 💬 [Discussions](https://github.com/PlatformStackPulse/go-template/discussions)

---

<p align="center">
  <sub>Built with ❤️ for platform engineers and Go developers</sub>
</p>
