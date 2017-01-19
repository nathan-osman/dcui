PKG = github.com/nathan-osman/dcui
CMD = dcui

SOURCES = $(shell find -type f -name '*.go')
BINDATA = $(shell find server/static server/templates)

all: dist/${CMD}

dist/${CMD}: ${SOURCES} server/bindata.go
	CGO_ENABLED=0 go build -o dist/${CMD} ${PKG}/cmd/${CMD}

server/bindata.go: ${BINDATA}
	go generate ${PKG}/...

clean:
	@rm -f dist/${CMD}
	@rm -f server/bindata.go

.PHONY: clean
