Broker
---------------------------

# Code architecture & data flow
              

# Test Producer & Consumer
```powershell
# ------------------ run two consumers -----------------
cd broker/cmd
go run main.go

# another terminal
cd broker/cmd
go run main.go

# ------------------ run one producer -----------------
cd broker/adaptee
go test -v -run TestKafka_Pub

```

# Build
```powershell

# service
cd broker/cmd
go build -o neon_broker .

# test client
cd tests/dapr 
go build -o neon_client .

```

# description
Dapr's sidecar can listen for both HTTP and gRPC requests even if you specify only one of the ports. By default, if you don't specify the --dapr-http-port or --dapr-grpc-port, Dapr will use the default ports (HTTP on 3500 and gRPC on 50001).

# Dapr argument sample
```powershell

# Your application will now communicate with the Dapr sidecar using HTTP on port 50001.
dapr run --app-id myapp --app-port 50001 --dapr-http-port 3500 --app-protocol http


# Your application will now communicate with the Dapr sidecar using gRPC on port 50001.
dapr run --app-id myapp --app-port 50001 --dapr-http-port 3500 --app-protocol grpc

# `--dapr-http-port` means expose the http port outside to call
# `--dapr-grpc-port` means expose the grpc port outside to call 
# `--app-protocol` means the protocol dapr sidebar call service
dapr run --app-id neon_broker  --app-protocol grpc  --app-port 19527  --dapr-grpc-port 50001 --dapr-http-port 3500  -- go run main.go

```

# Service run & invoke
```powershell

# run broker
cd broker/cmd
dapr run --app-id neon_broker --app-port 19527  --dapr-http-port 3600 -- ./neon_broker
# or
dapr run --app-id neon_broker --app-port 19527  --dapr-http-port 3600  -- go run main.go


# run test client
cd tests/dapr
dapr run --app-id neon_client --app-protocol http --dapr-http-port 3601 -- ./neon_client
# or
dapr run --app-id neon_client --app-protocol http --dapr-http-port 3601 -- go run cmd.go invoke_service.go

```


# Run in different proto
HTTP
```shell 
# TODO: the port `19527` need same as `config.yaml` broker block `app_port` field, now is `19528`
dapr run --app-id neon_broker  --app-protocol http  --app-port 19527  --dapr-http-port 3500   -- go run main.go 

dapr run --app-id neon_broker  --app-protocol http  --app-port 19527  --dapr-http-port 3500  --log-level debug  -- go run main.go 
```

gRPC
```shell 
dapr run --app-id neon_broker  --app-protocol grpc  --app-port 19527  --dapr-grpc-port 50001  -- go run main.go

dapr run --app-id neon_broker  --app-protocol grpc  --app-port 19527  --dapr-grpc-port 50001 --log-level debug -- go run main.go
```





