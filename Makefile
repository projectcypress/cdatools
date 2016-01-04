build:
	ego -o ./exporter/cat1_templates.go -package exporter ./exporter/templates

# fake out clean and install
clean:
install:

.PHONY: build
