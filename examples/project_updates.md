# Project API

Please refer to `client.md` before reviewing these examples.

## Usage

> Project Updates Cancel

```go
err := client.ProjectUpdatesService.ProjectUpdateCancel(4)

if err != nil {
    log.Fatalf("Cancel Update Projects err: %s", err)
}

log.Printf("Update Project cancelled.")
```

> Project Updates Get Update

```go
err := client.ProjectUpdatesService.ProjectUpdateGet(4)

if err != nil {
    log.Fatalf("Get Update Projects err: %s", err)
}

log.Printf("Get Project done.")
```