PKG = github.com/nathan-osman/dcui
CMD = dcui

all: dist/${CMD}

dist/${CMD}: server/bindata.go
	go build -o dist/${CMD} ${PKG}/cmd/${CMD}

server/bindata.go:
	go generate ${PKG}/...

clean:
	@rm -f dist/${CMD}
	@rm -f server/bindata.go

.PHONY: clean
