run:
	@go run ./cmd/tgbwp

build:
	@mkdir --parents ~/.local/bin/
	@go build -v -o ~/.local/bin ./cmd/tgbwp

build_test:
	@go build -v -buildmode=plugin -o ~/.local/share/tgbwp/plugins/test.so ./plugins/test

build_ip:
	@go build -v -buildmode=plugin -o ~/.local/share/tgbwp/plugins/ip.so ./plugins/ip
