build:
	go-bindata -o exporter/templates.go -pkg exporter exporter/templates/*/*.xml

debug:
	go-bindata -o exporter/templates.go -pkg exporter --debug exporter/templates/*/*.xml

# fake out clean and install
clean:
install:

.PHONY: build
