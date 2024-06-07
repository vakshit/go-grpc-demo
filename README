# go-grpc-demo

### How to run

Make sure you have golang installed on your system. If not, it can be downloaded from [here](https://go.dev/doc/install). v1.20.14 preferred.

This pplication also required installation of proto-compiler. It can be installed as:

```bash
# mac OS
brew install protobuf
# Ubuntu
sudo apt install protobuf
```

Once the setup is complete:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go mod download
make generate_protos
go run main.go
```

you can now make requests from [Postman](https://www.postman.com/downloads/).
