# The base go-image
FROM golang:latest AS builder

COPY . /go/src/github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data
WORKDIR /go/src/github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data

COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy only `.go` files, if you want all files to be copied then replace `with `COPY . .` for the code below.
COPY . .

# Install our third-party application for hot-reloading capability.
RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]
RUN ["go", "install", "github.com/githubnemo/CompileDaemon"]

ENTRYPOINT CompileDaemon -polling -log-prefix=false -build="go build ." -command="./temposcale rpc_serve" -directory="./"
