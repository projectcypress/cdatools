build:
	go-bindata -nocompress -o exporter/templates.go -pkg exporter --prefix "exporter/" exporter/templates/*/*/*.xml exporter/hqmf_template_oid_map.json exporter/hqmf_qrda_oids.json exporter/hqmf_qrda_oids_r3_1.json

debug:
	go-bindata -nocompress -o exporter/templates.go -pkg exporter --debug --prefix "exporter/" exporter/templates/*/*/*.xml exporter/hqmf_template_oid_map.json exporter/hqmf_qrda_oids.json exporter/hqmf_qrda_oids_r3_1.json

# fake out clean and install
clean:
install:

coverage:
	go test -coverprofile=exporter.out ./exporter
	go tool cover -html=exporter.out
	go test -coverprofile=importer.out ./importer
	go tool cover -html=importer.out
	go test -coverprofile=models.out ./models
	go tool cover -html=models.out
	rm exporter.out importer.out models.out

.PHONY: build
