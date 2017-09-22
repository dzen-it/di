# DI-scovery service

## Overview

Di is designed as an alternative to complex services for discovering hosts and/or balancing of them. The configuration is done through environment variables. Communication is through  HTTP REST API.

## Configs
The environment variables are used for configuration 

| ENV | DEFAULT | Description |
|---|---|---|
| **DI_MODE** | "cache" | Sets the session storage mode. If multiple instances of `di` are used, then use *"redis"* |
| **DI_HTTPS** | false | Switcher of HTTPS |
| **DI_CERTFILE** | "" | The path to the certificate file |
| **DI_KEYFILE**  | "" | The path to the secret key file |
| **DI_ADDRESS** | ":8080" | Listened HTTP/S address |
| **DI_REDIS** | "localhost:6379" | Address of redis server |
| **DI_REDISDB** | 0 | Number of redis DB |

## REST API   
At the moment there is a fixed minimum request interval of 1 minute.

| Endpoint | Method | Body | Resp. Code | Description | 
|---|---|---|---|---|
| /node | PUT | ```{'name':'node_name', 'address':'10.10.10.1'}  ``` | 201 | Overwrites the previously sent name. Without balancing. |
| /node | POST | ```{'name':'node_name', 'address':'10.10.10.1'}  ``` | 201 | If the name is the same as the one already sent, then the balancing mode is applied to the addresses |
| /node/{name} | GET | | 200 | Returns address of node | 

## Build
```bash 
$ git clone git@github.com:dzen-it/di.git 
$ cd di && go get ./ && go build -o bin/di 

# Run!
$ ./bin/di 
```

---

*I am happy to accept suggestions for improving the service.*
