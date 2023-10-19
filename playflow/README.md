PlayFlow 
-------------------
code process for some task implement in the program, can remote call to run

# Service run & invoke
```powershell

# run task
cd playbook/cmd
dapr run --app-id neon_playbook  --dapr-http-port 3602 --  /opt/anaconda3/bin/python main.py 

# run in auto task node
cd /opt/neonrabbit/NeonRabbit/playflow/cmd;sudo /usr/local/bin/dapr  run --app-id neon_playbook  --dapr-http-port 3602 -- /usr/local/python39/bin/python3 main.py
```


# Flow integration Tasks report
pub data for success
```json
{
    "code": 0,
    "message": "success",
    "data": [
        {
          "task_name": "sample1",
          "state": "success",
          "result": "sample1 run success"
        },
        {
          "task_name": "sample2",
          "state": "success",
          "result": "sample2 run success"
        },
        {
          "task_name": "sample3",
          "state": "fail",
          "result": "sample3 run fail"
        },
        {
          "task_name": "sample4",
          "state": "noStart",
          "result": ""
        }
    ]
}
```






