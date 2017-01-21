PKG = github.com/nathan-osman/dcui
CMD = dcui

SOURCES = $(shell find -type f -name '*.go')
BINDATA = $(shell find server/static server/templates)

all: dist/${CMD}

dist/${CMD}: ${SOURCES} server/ab0x.go
	CGO_ENABLED=0 go build -o dist/${CMD} ${PKG}/cmd/${CMD}

server/ab0x.go: ${BINDATA} b0x.yaml
	fileb0x b0x.yaml

clean:
	@rm -f dist/${CMD}
	@rm -f server/ab0x.go

.PHONY: clean
