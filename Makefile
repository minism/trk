OUT := trk
PKG := github.com/minism/trk
VERSION := $(shell git describe --always --long --dirty)

build:
	go build -v -o ${OUT} -ldflags="-X github.com/minism/trk/internal/version.Version=${VERSION}" ${PKG}

lint:
	staticcheck ./...

clean:
	-@rm -f ${OUT}

license:
	addlicense -c "Josh Bothun" -l mit -v -y 2024 -s .

.PHONY: build lint clean license
