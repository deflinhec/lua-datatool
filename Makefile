# Define
VERSION=0.1.0
BUILD=$(shell git rev-parse HEAD)

# Setup linker flags option for build that interoperate with variable names in src code
ifeq ($(OS),Windows_NT)
	LDFLAGS='-s -w -X "main.Version=$(VERSION)" -X "main.Build=$(BUILD)" -H=windowsgui -extldflags=-static'
	PLATFORM := windows
else
	LDFLAGS='-s -w -X "main.Version=$(VERSION)" -X "main.Build=$(BUILD)"'
	PLATFORM := $(shell uname -s | tr A-Z a-z)
endif

.PHONY: default build translate osx-app assets

default: fmt build tidy

fmt:
	go fmt ./...

tidy:
	go mod tidy

go-text:
	go install golang.org/x/text/cmd/gotext@latest

translate: go-text
	go generate ./internal/translations/translations.go

go-bindata:
	go get -u github.com/go-bindata/go-bindata/...
	go install -a -v github.com/go-bindata/go-bindata/...

.PHONY: assets
assets: go-bindata
	git clean -xdff assets
	go-bindata -nomemcopy -pkg=assets -o=assets/assets.go \
		-debug=$(if $(findstring debug,$(BUILDTAGS)),true,false) \
		-ignore=assets.go -ignore=init.go assets/...

osx-tool:
	go get github.com/machinebox/appify
	go install -a -v github.com/machinebox/appify

osx-app: build
	$(foreach file, $(wildcard $(CURDIR)/build/**/*), \
		$(if $(shell grep ".app" "$(file)"), \
			appify -version $(VERSION) -name $(notdir $(file)) \
				-author deflinhec -icon ./icon.png $(file); \
			rm -rf $(file).app; \
			mv $(notdir $(file)).app $(dir $(file)); \
		,) \
	)

# Sperate "linux-amd64" as GOOS and GOARCH
OSARCH_SPERATOR = $(word $2,$(subst -, ,$1))

# Arch build options
arch-%: export GOARCH=$(call OSARCH_SPERATOR,$*,1)
arch-%: export CGO_ENABLED=1
arch-%: fmt assets tidy
	go build -ldflags $(LDFLAGS) -o ./build/$(GOARCH)/ ./cmd/...

# Local build options
build: export CGO_ENABLED=1
build: fmt assets tidy
	go build -ldflags $(LDFLAGS) -o ./build/$(PLATFORM)/ ./cmd/...


