build:
	go-bindata -o exporter/templates.go -pkg exporter --prefix "exporter/" exporter/templates/*/*.xml

debug:
	go-bindata -o exporter/templates.go -pkg exporter --debug --prefix "expporter/" exporter/templates/*/*.xml

# fake out clean and install
clean:
install:

.PHONY: build
