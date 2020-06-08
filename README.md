# lyr-service

lyr-service is a go rest api that wraps an elasticsearch database.

## Usage

Clone the repository to a new folder then run the following commands:

```bash
cd api
go build api

cd elasticpersist
go build elasticpersist

cd utils
go build utils

go install lyr-service
lyr-service
```

## Dependencies

- [gorilla Mux](https://github.com/gorilla/mux)
- [gorilla Handlers](https://github.com/gorilla/handlers)
- [go-elasticsearch](https://github.com/elastic/go-elasticsearch)
