build:
	@go build -buildmode=plugin -o ~/.local/share/knackwurstking/tgbwp/plugins/ip.so
	@go build cmd/tgbwp -o ~/.local/bin/tgbwp
