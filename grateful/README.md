## Graceful Shutdown

These examples illustrate how to gracefully shutdown an http server making use of the high level http package.

## Usage

```console
go run main.go
2021-11-13 23:01:35     [INFO]: starting server on: localhost:7777 
```

there are 3 endpoints tkat would consume your time and energy if you let them!

```
/5sec
/10sec
/15sec
```

Depending on how fast and far your `ctrl` and `c` keys are, chose the one that's better fit for the job and see what happens. Uppon pressing the tedious key combination the server will close any iddle connections and wait for those that are still open for a period of 5 seconds, which is the grace timeout. So from here on there are really 2 ways of testing this little examples:

1. use the `10sec` and `15sec` endpoints, as the graceful shutdown tolerance is set to 5 seconds and force the server to free resources, resulting in unfinished requests, or
2. use the `5sec` endpoint and send a SIGINT to the server process while the request is still in progress, the server will then wait for the request to finish and free all resources.
