TAG := d7203eb719304866a7eb7033ef03d421459335b8

schema-update:
	curl https://raw.githubusercontent.com/tdlib/td/${TAG}/td/generate/scheme/td_api.tl 2>/dev/null > ./data/td_api.tl
	curl https://raw.githubusercontent.com/tdlib/td/${TAG}/td/telegram/Td.cpp 2>/dev/null > ./data/Td.cpp

generate-json:
	go run ./cmd/generate-json.go \
		-output "./data/td_api.json"

generate-code:
	go run ./cmd/generate-code.go \
		-outputDir "./client" \
		-package client \
		-functionFile function.go \
		-typeFile type.go \
		-unmarshalerFile unmarshaler.go
#	go fmt ./...
