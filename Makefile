build:
	@go build -o bin/gobank

run: build
	@./bin/gobank

test:
	@go test ./...

# etc/resolve.conf
# 8.8.8.8

#
#Get-NetIPConfiguration "vEthernet (WSL)"
