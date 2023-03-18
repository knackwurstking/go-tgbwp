build:
	@go build -v -buildmode=plugin -o ~/.local/share/knackwurstking/tgbwp/plugins/ip.so ./plugins/ip
	@mkdir --parents ~/.local/bin/
	@go build -v -o ~/.local/bin ./cmd/tgbwp
