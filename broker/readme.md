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
dapr run --app-id neon_broker  --app-protocol grpc  --app-port 19527  --dapr-grpc-port 3500  -- go run main.go

dapr run --app-id neon_broker  --app-protocol grpc  --app-port 19527  --dapr-grpc-port 3500 --log-level debug -- go run main.go
```





