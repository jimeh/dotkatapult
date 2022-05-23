NAME = dotkatapult
BINARY = bin/${NAME}
VERSION ?= $(shell cat VERSION)
SOURCES = $(shell find . -name '*.go' -o -name 'Makefile')

$(BINARY): $(SOURCES)
	CGO_ENABLED=0 go build -a -o ${BINARY} -ldflags "-s -w"

.PHONY: build
build: $(BINARY)

.PHONY: run
run: $(BINARY)
	$(BINARY)

.PHONY: clean
clean:
	$(eval BIN_DIR := $(shell dirname ${BINARY}))
	if [ -f ${BINARY} ]; then rm ${BINARY}; fi
	if [ -d ${BIN_DIR} ]; then rmdir ${BIN_DIR}; fi

.PHONY: docker
docker:
	docker build -t "$(shell whoami)/$(NAME)" .