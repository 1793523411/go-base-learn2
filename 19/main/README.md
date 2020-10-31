## docker  redis

`docker run --name gobase-redis -p 6380:6379 -d daocloud.io/redis`

`docker exec -it gobase-redis bash` or `docker run -it --network host --rm redis:5.0.7 redis-cli` 

go-redis :

```
32.21:443: connect: connection refused
$ go get -u github.com/go-redis/redis
unrecognized import path "go.opentelemetry.io/otel/api/global": https fetch: Get "https://go.opentelemetry.io/otel/api/global?go-get=1": dial tcp 216.239.32.21:443: connect: connection refused
unrecognized import path "go.opentelemetry.io/otel/api/metric": https fetch: Get "https://go.opentelemetry.io/otel/api/metric?go-get=1": dial tcp 216.239.32.21:443: connect: connection refused
unrecognized import path "go.opentelemetry.io/otel/api/trace": https fetch: Get "https://go.opentelemetry.io/otel/api/trace?go-get=1": dial tcp 216.239.32.21:443: connect: connection refused
unrecognized import path "go.opentelemetry.io/otel/label": https fetch: Get "https://go.opentelemetry.io/otel/label?go-get=1": dial tcp 216.239.32.21:443: connect: connection refused


```

run err:

```
../../../github.com/go-redis/redis/internal/instruments.go:6:2: cannot find package "go.opentelemetry.io/otel/api/global" in any of:
        /opt/go/src/go.opentelemetry.io/otel/api/global (from $GOROOT)
        /home/ygjzs/goproject/src/go.opentelemetry.io/otel/api/global (from $GOPATH)
../../../github.com/go-redis/redis/internal/instruments.go:7:2: cannot find package "go.opentelemetry.io/otel/api/metric" in any of:
        /opt/go/src/go.opentelemetry.io/otel/api/metric (from $GOROOT)
        /home/ygjzs/goproject/src/go.opentelemetry.io/otel/api/metric (from $GOPATH)
../../../github.com/go-redis/redis/internal/util.go:10:2: cannot find package "go.opentelemetry.io/otel/api/trace" in any of:
        /opt/go/src/go.opentelemetry.io/otel/api/trace (from $GOROOT)
        /home/ygjzs/goproject/src/go.opentelemetry.io/otel/api/trace (from $GOPATH)
../../../github.com/go-redis/redis/options.go:18:2: cannot find package "go.opentelemetry.io/otel/label" in any of:
        /opt/go/src/go.opentelemetry.io/otel/label (from $GOROOT)
        /home/ygjzs/goproject/src/go.opentelemetry.io/otel/label (from $GOPATH)
```