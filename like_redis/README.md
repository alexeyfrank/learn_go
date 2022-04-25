# Dummy implementation of client-service application like Redis

## How to use

1. Start server `go run main.go --server --addr=localhost:8810`
2. Connect to server `go run main.go --client --addr=localhost:8810`
3. Play with it:
```
> SET key 1
1
> GET key
1
> INC key 10
11
> DEC key 5
6
```

## TODO

1. Write specs
2. Switch to Protobuf
3. Implement more in-memory data structures
4. Implement persistence
5. ???
