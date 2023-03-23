# Ping API

Please refer to `client.md` before reviewing these examples.

## Usage

> Ping

```go
result, err := client.PingService.Ping()
if err != nil {
    log.Fatalf("Ping awx err: %s", err)
}

log.Println("Ping awx: ", result)
```
