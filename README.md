# Golang grpc, http example

Service returns fibonacci numbers from requested range


### Generate grpc:
```bash
brew install protoc-gen-go
brew install protoc-gen-go-grpc

protoc --go_out=./internal/api --go_opt=paths=source_relative \
    --go-grpc_out=./internal/api --go-grpc_opt=paths=source_relative \
    api_grpc.proto
```

### deployment:
```
docker build -t fibo-example .
docker run -d \
    -e HTTP_ADDR=:8077 \
    -e GRPC_ADDR=:8078 \
    -p 8077:8077 \
    -p 8078:8078 \
    fibo-example
```


### Http test:
```bash
export FIBO_HTTP_HOST=localhost:8077
curl -i -X GET "${FIBO_HTTP_HOST}/fibo/50-100"
```

### Grpc test:
```bash
brew install grpcurl

export FIBO_GRPC_HOST=localhost:8078
grpcurl -plaintext -import-path ./ -proto api_grpc.proto -d @ ${FIBO_GRPC_HOST} Fibo/GetFiboNumbers <<EOM
{ "from": 2, "to": 5 }
EOM
```
