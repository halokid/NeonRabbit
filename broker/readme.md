Broker
---------------------------

# Code architecture & data flow
code architecture design is base on the program data flow

Go import design must single flow, can not import each other

`every next flow can be the son file or folder`
income -->>> service
              |-- handler
                    |-- brokerx
                    |-- pkgx
                          |-- vo
                          |-- logger
              

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

# Service invoke
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





