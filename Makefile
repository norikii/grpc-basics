gen:
	protoc --proto_path=proto proto/*.proto --go_opt=paths=source_relative --go_out=plugins=grpc:pb

clean:
	rm pb/*.go

server:
	go run cmd/server/main.go --port 8080

client:
	go run cmd/client/main.go --address 0.0.0.0:8080