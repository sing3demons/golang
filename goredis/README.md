

```k6 on docker
docker compose run --rm k6 run /scripts/test.js -u5 -d5s
docker run --rm -i grafana/k6 run - <scripts/test.js  
docker run --rm -i grafana/k6 run --out influxdb=http://influxdb:8086/k6 - <scripts/test.js 
```

```
docker compose up influxdb grafana mariadb redis -d
docker compose run --rm k6 run /scripts/test.js -u5 -d5s
```