# Logging service for bookCRUD 

## Technologies
- Audit Log service gRPC Server

## Before Run
To run, you must specify environment variables, for example, in the .env file
```
export DB_DATABASE=TESGO
export DB_URI=mongodb+srv:...

export SERVER_PORT=9000
```

## Build & Run 
```
source .env && go build -o app cmd/main.go && ./app
```