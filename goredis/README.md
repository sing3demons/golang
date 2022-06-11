# goredis

## start app

```
docker compose up influxdb grafana mariadb redis -d
go run .
docker compose run --rm k6 run /scripts/test.js
```

### run k6 on docker

```k6 on docker
docker run --rm -i grafana/k6 run - <scripts/test.js
docker run --rm -i grafana/k6 run --out influxdb=http://influxdb:8086/k6 - <scripts/test.js
```

## Redis Configuration

```redis
bind 0.0.0.0
appendonly yes
SAVE ""
```
