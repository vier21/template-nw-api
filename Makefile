.PHONY: test genkey genpub

test:
	@go test ./...
genkey:
	@openssl genrsa -out secret/keypair/pkcs8_privateKey.pem
genpub:
	@openssl rsa -in secret/keypair/pkcs8_privateKey.pem -pubout >> secret/keypair/publicKey.pub
	

