gen-keys:
	mkdir -p ./keys
	ssh-keygen -N "" -t rsa -b 4096 -m PEM -f ./keys/rs256.rsa
	openssl rsa -in ./keys/rs256.rsa -pubout -outform PEM -out ./keys/rs256.rsa.pub
swag-gen:
	swag init -g internal/delivery/server.go