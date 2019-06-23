# Go gRPC calculator server and client

## Quickstart

```
$ make help
 Choose a make command:

  build          Build all Go binaries.
  install        Install Go binaries into GOPATH.
  compile        Run the Go compiler.
  test           Run tests.
  upgrade        Upgrade and vendor dependencies.
  proto          Generate protobuf files.
  docker-build   Build calc-server Docker image.
  docker-run     Run calc-server Docker image.
  deploy         Deploy calc-server on Kubernetes IN Docker.
  clean          Clean build cache, remove binaries.
```

## Calculator Server

```
$ ./calc-server --help
Usage of calc-server:
  -config string
    	config file (optional)
  -host string
    	listen address (default "localhost")
  -port uint
    	port number (default 7777)
```

Environment variables `CALC_SERVER_HOST` and `CALC_SERVER_PORT` may be used to set `host` and `port` respectively.

## Calculator Client

```
$ ./calc-client --help
NAME:
   Simple calc-server CLI - An example CLI to send basic request to calc-server

USAGE:
   calc-client [global options] command [command options] [arguments...]

COMMANDS:
     add      Add two 32-bit floating point numbers
     div      Divide a 32-bit floating point number by another one
     mul      Multiply two 32-bit floating point numbers
     sub      Subtract two 32-bit floating point numbers
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --host value  Calc server host to connect to (default: "localhost") [$CALC_SERVER_HOST]
   --port value  Calc server port to connect to (default: 7777) [$CALC_SERVER_PORT]
   --help, -h    show help
```

Environment variables `CALC_SERVER_HOST` and `CALC_SERVER_PORT` may be used to set `host` and `port` respectively.

### Negative values as arguments

Invoking `calc-client` with negative numbers is done by prepending the command line arguments with `--`.
This signalizes to the command line parser that all subsequent arguments should not be parsed as flags.

```
$ calc-client add -- -1 2
1
```

# [12 Factor](https://12factor.net/) Best Practices

1. Codebase under version control (`git`)
2. All dependencies bundled via `go mod vendor`. Third-party tools like `protoc-gen-go` and `kind` are vendored as well.
3. Configuration is injectable via environment variables.
4. There a no `Backing services` with app, however the `calc-client` attaches to any `calc-server` specified via its environment.
5. Separation of `build stage`, `release stage` and `run stage`. The multi-stage `Dockerfile` first builds the `calc-server` binary and then creates a smaller release image with configuration provided through the `envfile`. The `run stage` then is comprised of deploying the release image on `Kubernetes` (via `kind`).
6. `calc-server` is designed to be a stateless application.
7. `calc-client` and `calc-server` are completely self-contained.
8. `calc-server` is horizontally scalable. Adding more concurrency is very simple.
9. `calc-server` can be started/stopped in an instant and it shuts down gracefully.
10. This project is designed for continous deployment, it's just not set up yet.

# Cloud Native Understanding

This project ships `calc-server` in a `Docker` image which can be readily deployed on any Kubernetes cluster via the `deployment.yaml` file also shipped with the source code. This means anyone can get it up and running on most cloud providers and onprem in a matter of minutes. The project aims to automate deployments and configuration as much as possible.

# Expanding this service

* Expand the API to allow for `square root` operation (second payload + response type)
* Set up CI job (Travis or CircleCI on GitHub)
* Set up test coverage reporting.
* Get external connection to Kubernetes to work (on Kubernetes in Docker).
* Add TLS. The `Dockerfile` is prepared for that, CA certificates must be set up.

# Access from the outside to the cluster

The project is certainly lacking a proper `Kubernetes` deployment setup, however due to time limitations this was not further explored so far. One could imagine an `IngressController` being set up to control traffic the the `calc-service` in the current deployment scheme.
