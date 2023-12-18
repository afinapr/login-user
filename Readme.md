# How to Run (localhost)

## Windows
- store google credentials as `cred.json` and placed it in root directory
- setup cloudsql proxy
`cloud-sql-proxy.x64.exe poc-ddb-mlops:asia-southeast2:bri-ct-hc-data -p=3307 -c=./cred.json`

- build and run
`go run .`