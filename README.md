NeonRabbit
-------------------

A neon style clever robbit.



# Feature

- [ ] support Terraform(for Azure)
- [ ] support Splunk
- [ ] Task flow for target OS like pipeline style



# Devops
```shell
# consul
consul.exe agent -dev -ui -client 0.0.0.0

# zipkin
java -jar zipkin-server-2.24.0-exec.jar

# kafka
# go to path to kafka
.\bin\windows\zookeeper-server-start.bat .\config\zookeeper.properties

.\bin\windows\kafka-server-start.bat  .\config\server.properties

```