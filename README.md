# taal-client

This is an small service that runs on a Taal customer's machine and interacts with Taal's API.

All private keys used for signing customer's transactions are only held on the customer's machine.

Binaries for Linux, Mac and Windows can be found at https://taal-gmbh.github.io.

## Direct Usage

```text
Usage
-----
taal-client register <api-key>
  Creates a private key which is held locally, and sends the public key to Taal to be linked to an existing API key.

taal-client start
  Starts listening for requests on :9500.  This value can be changed with the LISTEN environment variable.

  All requests will be sent to https://tapi.taal.com by default unless overridden with the TAAL_URL environment variable.

  DEBUG = 1 will log all transactions to console

Environment variables
---------------------
  LISTEN
  TAAL_URL
  TAAL_TIMEOUT
  DEBUG

Example
-------
  DEBUG=1 LISTEN=localhost:8080 TAAL_URL=http://localhost:4000 TAAL_TIMEOUT=1m ./taal_client start

```

## Registration

Before this client can be used, a valid Taal APIKey needs to be registered in order to bind it with a public key.

1. Register at https://console.taal.com/register
2. Create an APIKey
3. Register the APIKey in TaalClient: ```taal-client register <APIKey>```


![Register sequence](./assets/register.png)


## Writing data

After starting the TaalClient with ```taal-client start``` you can then write data to the blockchain by POSTing to the TaalClient.

```c
curl --location --request POST 'http://localhost:9500/api/v1/write' \
--header 'X-FeesRequired: true' \
--header 'X-Tag: AN_OPTIONAL_TAG' \
--header 'Authorization: Bearer <APIKey>' \
--header 'Content-Type: application/json' \
--data-raw '{
    "key1": "value1",
    "key2": "value2"
}'
```

![Writing sequence](./assets/write.png)

## Reading data

After starting the TaalClient with ```taal-client start``` you can then read data from the blockchain by GETing from the TaalClient.

```c
curl --location --request GET 'http://localhost:9500/api/v1/read/<txid>' \
--header 'Authorization: Bearer <APIKey>'
```
