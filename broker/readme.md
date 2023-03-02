Broker
---------------------------

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