# tech:lounge Masterclass

Run EventSourcingDB:

```shell
docker run \
  -it \
  --init \
  -p 3000:3000 \
  thenativeweb/eventsourcingdb:1.0.1 \
  run \
  --api-token secret \
  --data-directory-temporary \
  --http-enabled \
  --https-enabled=false \
  --with-ui
```

Environment variables to be used:

- `PORT` (default `4000`)
- `ESDB_URL` (default `http://localhost:3000`)
- `ESDB_API_TOKEN` (default `secret`)

Build the Docker image:

```shell
make build-docker
```
