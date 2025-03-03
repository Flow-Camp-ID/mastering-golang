# Nama aplikasi
APP_NAME=restfulapi

# Variabel untuk perintah Go
GO=go
GO_BUILD=$(GO) build
GO_RUN=$(GO) run

# Target install dependencies
deps:
	$(GO) mod tidy

# Target Run
run:
	$(GO) run main.go

# Target help
help:
	@echo "Usage: make <target>"
	@echo ""
	@echo "Targets:"
	@echo "  run     : Menjalankan aplikasi Go"
	@echo "  deps    : Install dependencies"
	@echo "  help    : Menampilkan bantuan ini"

.PHONY: run  deps  help