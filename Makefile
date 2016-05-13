build:
	go-bindata -o exporter/templates.go -pkg exporter --prefix "exporter/" exporter/templates/*/*.xml exporter/hqmfr2_template_oid_map.json exporter/hqmf_template_oid_map.json

debug:
	go-bindata -o exporter/templates.go -pkg exporter --debug --prefix "exporter/" exporter/templates/*/*.xml exporter/hqmfr2_template_oid_map.json exporter/hqmf_template_oid_map.json

# fake out clean and install
clean:
install:

.PHONY: build
