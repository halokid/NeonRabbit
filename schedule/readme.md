Schedule service desing
----------------------------
remote call the automation jobs by AS platform API

# Process flow

*single process*
1. Schedule -->>> sp (this is sample, add GW, all trace is:  gw -->>> Schedule -->>> sp)
2. Schedule -->>> sn
3. Schedule -->>> cb


*mutiple process*
1. Schedule -->>> sp -->>> sn
2. Schedule -->>> broker(grpc) -->>> sn 


# Run

```shell

dapr.exe run --app-id neon_schedule  --app-protocol http  --app-port 19567  --dapr-http-port 3700   -- go run main.go

```

