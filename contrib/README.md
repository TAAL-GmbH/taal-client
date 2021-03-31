# Running Taal Client in Docker


## Build

```bash
docker build -t taal-client:latest -f ./contrib/Dockerfile .
```

## Run

Run the server daemon.

```bash
docker run --rm=true --name taal  -e MAPI_URL=https://mapi.staging.taal.com/mapi -p 8080:8080 -e LISTEN=localhost:8080  -d taal-client:latest
```
* retreive the `$API_KEY` from Taal Console testnet package

To stop the container:

```bash
docker stop taal
```

## Register API keys

The user can register muliple API keys.  To register an API key, do the following:

```bash
docker exec taal /go/bin/taal-client register $API_KEY
```
* retreive the `$API_KEY` from a Taal Console package