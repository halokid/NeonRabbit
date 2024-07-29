Eapi service desing
----------------------------
Extend api for micro service management

# Process flow

*single process*
1. gateway -->>> eapi

# Run

```shell

dapr.exe run --app-id neon_eapi  --app-protocol http  --app-port 19577  --dapr-http-port 3900   -- java -jar ./target/eapi-0.0.1-SNAPSHOT.jar 

```

