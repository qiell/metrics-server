## Problem Statement

Implement and Containerize a REST API server that ingests metrics from its clients and generates on-demand stats reports. 
Description 
Imagine you’re writing a metrics ingestion server as part of an Infrastructure monitoring tool. Implement an HTTP server that exposes the following REST API: 
- Ingestion 
```
    Method: POST 
    Path: /metrics 
    Headers: 
        content-type: application/json 
    JSON body structure:
    { 
        "percentage_cpu_used": <integer between 0-100>, 
        "percentage_memory_used": <integer between 0-100> 
    } 
    Responses: 
        1. 200 If api has successfully ingested the data supplied 
        2. 500 if anything goes wrong
``` 

Sample curl request (api is deployed locally and bound to port 8080): 
```bash
   $ curl -XPOST \
    -H "Content-Type: application/json" \ 
    --data '{"percentage_cpu_used": 55, "percentage_memory_used": 90}' \ http://127.0.0.1:8080/metrics
```
   
The API accepts a json object containing info about the sender’s cpu and memory percentage utilization. The api must store this information in-memory along with the IP address of the sender. 
A sender can send their metrics to the API at different points of time. For eg- sender A might send metrics 25 times at 2-second intervals, B might send 500 times at 0.5 sec intervals. For this task, you can ignore concurrency issues like race conditions while managing the data in-memory.

- Report Generation 
```
    Method: GET 
    Path: /report 
    Headers: 
        content-type: application/json 
    Responses: 
        1. 500 if anything goes wrong 
        2. 200 and a json payload with the following structure: 
    [ 
        { 
            "ip": <IP address of machine>, 
            "max_cpu": <maximum cpu %>, 
            "max_memory": <maximum memory %> 
        }, 
    ]
``` 

The JSON payload is an array in which each object corresponds to a unique IP whose metrics were ingested by the API. You need to return the maximum cpu & memory utilization that the client ever reached.
Sample curl request: 
```bash
    $ curl -XGET \ 
    -H "Content-Type: application/json" \ 
    http://127.0.0.1:8080/report 
    [ 
        { 
            "ip": “10.12.10.16”, 
            "max_cpu": 87, 
            "max_memory": 98 
        }, 
        { 
            "ip": “10.12.10.78”, 
            "max_cpu": 44, 
            "max_memory": 50 
        } 
    ] 
```

### Directory structure
```tree
├── bin
│   └── server
├── build
│   └── Dockerfile.metrics-server
├── constants
│   └── constants.go
├── events
│   ├── broker.go
│   ├── metrics.go
│   └── types.go
├── main.go
├── Makefile
├── README.md
├── scripts
│   ├── build-server
│   └── version
└── store
    ├── inmemory
    │   └── inmemory.go
    ├── store.go
    └── types.go

```

### How to build and run? 
- Copy or clone the repository to `$GOPATH/src/github.com/qiell/metrics-server`
- Change directory to `$GOPATH/src/github.com/qiell/metrics-server`
- Run `make` to build the container image
- Use `docker run -p 8080:8080 <container-image-name>` to run the conatiner 