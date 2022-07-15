# taal-client

This is an small service that runs on a Taal customer's machine and interacts with Taal's API.

All private keys used for signing customer's transactions are only held on the customer's machine.

Binaries for Linux, Mac and Windows can be found at https://taal-gmbh.github.io.


## Direct Usage

```text
Usage
-----
taal-client
  Starts listening for requests on :9500.  This value can be changed via the console (https://localhost:9500) or in the settings.conf file.
  
  All requests will be sent to https://api.taal.com by default unless overridden  via the console (https://localhost:9500) or in the settings.conf file.
```

## Registration

Before this client can be used, a valid Taal APIKey needs to be registered in order to bind it with a public key.

1. Register at https://console.taal.com/register
2. Obtain an APIKey
3. Make sure the TaalClient is running
4. Register the APIKey in TaalClient via the Settings page which can be found at http://localhost:9500


![Register sequence](https://github.com/TAAL-GmbH/taal-client/blob/master/assets/register.png)


## Writing data

After starting the TaalClient with ```taal-client``` you can then write data to the blockchain using the drag and drop interface (http://localhost:9500) or by POSTing directly to the TaalClient API.

```c
curl --location --request POST 'http://localhost:9500/api/v1/write' \
--header 'X-Tag: AN_OPTIONAL_TAG' \
--header 'X-Mode: <raw|hash|encrypt>' \
--header 'X-Key: A_SHARED_SECRET_KEY' \  
--header 'Authorization: Bearer <APIKey>' \
--header 'Content-Type: application/json' \
--data '{
    "key1": "value1",
    "key2": "value2"
}'
```

or to POST a file, use the --data-binary flag:

```
curl --location --request POST 'http://localhost:9500/api/v1/write' \
--header 'X-Mode: raw' \
--header 'Authorization: Bearer <APIKey>' \
--header 'Content-Type: image/png' \
--data-binary @myimage.png
```

### Modes
Currently data can be submitted in 3 different modes
1. Raw: The full data is submitted to the blockchain as raw data.
2. Hash: A SHA256 hash is created of the input data. Only this hash is submitted.
3. Encrypt: Data is encrypted by the given secret using AES encryption. The secret is stored in the local database. When downloading the data, the stored secret used for decryption. Transactions with encrypted data is denoted by a key symbol.

![Writing sequence](https://github.com/TAAL-GmbH/taal-client/blob/master/assets/write.png)

## Reading data

After starting the TaalClient with ```taal-client``` you can then read data from the blockchain by GETing from the TaalClient.

```c
curl --location --request GET 'http://localhost:9500/api/v1/read/<txid>' \
--header 'Authorization: Bearer <APIKey>'
```

## Before usage
Before running the taal-client binary, make sure it is executable by running

```
chmod 755 taal-client
```

### MacOS

For the Mac version the code is not signed with a certificate issued by Apple currently. Therefore when running for the first time the following message will be shown.

![Mac1](https://github.com/TAAL-GmbH/taal-client/blob/master/assets/mac1.png)

In order to still run the application on a please open `Security & Privacy` settings and click on `Open Anyway` as shown in the following picture.

![Mac2](https://github.com/TAAL-GmbH/taal-client/blob/master/assets/mac2.png)

After that when running the application the following message will be shown. This time it has an `Open` button. Press this button in order to run the application.

![Mac3](https://github.com/TAAL-GmbH/taal-client/blob/master/assets/mac3.png)
§