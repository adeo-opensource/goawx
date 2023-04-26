# Job API

Please refer to `client.md` before reviewing these examples.

## Usage

> Get Job

```go
result, err := client.JobService.GetJob(yourJobId, map[string]string{})
if err != nil {
    log.Fatalf("Get Job err: %s", err)
}

log.Println("Get job: ", result)
```

> Cancel Job

```go
result, err := client.JobService.CancelJob(yourJobId, map[string]interface{}{}, map[string]string{})
if err != nil {
    log.Fatalf("Cancel Job err: %s", err)
}

log.Println("Cancel job: ", result)
```

> Relaunch Job

```go
result, err := client.JobService.RelaunchJob(yourJobId, map[string]interface{}{"hosts": "all"}, map[string]string{})
if err != nil {
    log.Fatalf("Relaunch Job err: %s", err)
}

log.Println("Relaunch job: ", result)
```

> Get Host Summaries

```go
result, _, err := client.JobService.GetHostSummaries(yourJobId, map[string]string{})
if err != nil {
    log.Fatalf("Get Host Summaries err: %s", err)
}

log.Println("Get Host Summaries: ", result)
```

> Get Job Events

```go
result, _, err := client.JobService.GetJobEvents(yourJobId, map[string]string{
    "order_by":  "start_line",
    "page_size": "1000000",
})
if err != nil {
    log.Fatalf("Get Job Events err: %s", err)
}

log.Println("Get Job Events: ", result)
```
