# lyr-service

lyr-service is a go rest api that wraps an elasticsearch database

## Usage

Clone the repository to a new folder then run the following commands:
``` bash
cd api
go build api

cd elasticpersist
go build elasticpersist

go install lyr-service
lyr-service
```

## Dependencies

gorilla Mux
gorilla Handlers
go elasticsearch
