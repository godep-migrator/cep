go-processing
=============

Intake -> [Processing] -> Dispatch

## Dependencies

1. go 1.2
2. godep
3. cassandra 2.0+
4. kafka 0.8

## Getting Started

**0**. Start cassandra and kafka via `cassandra` and `bin/zookeeper-server-start.sh` and `bin/kafka-server-start.sh`

**1**. Install godep

```shell
$ go get github.com/kr/godep
```

**2**. Get dependencies

```shell
$ godep restore
```

**3**. Compile and run

```shell
$ go run processor.go
```

### Before checking in

**1**. Make sure it compiles

```shell
$ go install
```

**2**. Run code formatter

```shell
go fmt ./...
```

**3**. Run tests

```shell
$ go test ./...
```

### Creating a keyspace

`create keyspace processor_development with replication = {'class': 'SimpleStrategy', 'replication_factor': 3};`
