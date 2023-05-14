# Logging service for bookCRUD 

## Technologies
- Audit Log service gRPC Server with Message Queue

## Before Run
To run, you must specify environment variables, for example, in the .env file
```
export DB_DATABASE=TESGO
export DB_URI=mongodb+srv:...

export SERVER_PORT=9000
```

start RabbitMQ with Docker 
```
docker run -it --rm --name rabbitmq -p 5672:5672 rabbitmq
```

## Build & Run (gRPC )
```
go build -o app cmd/grpc/main.go && ./app
```

## Build & Run (gRPC with Message Queue)
```
go build -o app cmd/rabbitMQ/main.go && ./app
```