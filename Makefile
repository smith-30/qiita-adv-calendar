COMMIT = $$(git rev-parse --short HEAD)

.PHONY: build

build:
	go build -v -ldflags "-X main.revision=\"$(COMMIT)\"" -o build/bin/qiita-adv