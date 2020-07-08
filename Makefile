GO       := GOSUMDB=off go
GOBUILD  := CGO_ENABLED=0 $(GO) build $(BUILD_FLAG)

EXECUTABLES := $(addprefix bin/,$(notdir $(wildcard cmd/*)))

.PHONY: default
default: clean build

.PHONY: build
build: $(EXECUTABLES)

$(EXECUTABLES):
	$(GOBUILD) -o $@ ./cmd/$(notdir $@)

.PHONY: clean
clean:
	rm -f $(EXECUTABLES)

.PHONY: get
get:
	$(GO) get ./...
	$(GO) mod verify
	$(GO) mod tidy

.PHONY: update
update:
	$(GO) get -u -v all
	$(GO) mod verify
	$(GO) mod tidy

.PHONY: fmt
fmt:
	gofmt -s -l -w $(FILES) $(TESTS)
