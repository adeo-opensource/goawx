# Host API

Please refer to `client.md` before reviewing these examples.

## Usage

> List Hosts

```go
result, _, err := client.HostService.ListHost(map[string]string{
    "name": "localhost",
})
if err != nil {
    log.Fatalf("List Hosts err: %s", err)
}

log.Println("List Hosts: ", result)
```

> Create Host

```go
result, err := client.HostService.CreateHost(map[string]interface{}{
    "name":        "test",
    "inventory":   1,
    "description": "test create host",
    "enabled":     true,
    "instance_id": "",
    "variables":   "ansible_host: localhost",
}, map[string]string{})

if err != nil {
    log.Fatalf("Create Host err: %s", err)
}

log.Printf("Host created. Host ID: %d", result.Host.ID")
```

> Update Host

```go
result, err := client.HostService.UpdateHost(3, map[string]interface{}{
    "description": "Add description here",
}, map[string]string{})

if err != nil {
    log.Fatalf("Update Host err: %s", err)
}

log.Printf("Updated Host ID: %d", 3)
```

> Associate Host to Group

```go
result, err := client.HostService.AssociateHostGroup(3, map[string]interface{}{
    "id": 10,
}, map[string]string{})

if err != nil {
    log.Fatalf("Associate Host err: %s", err)
}

log.Printf("Associate Host ID: %d", 3)
```

> Disassociate Host from Group

```go
result, err := client.HostService.DisAssociateHostGroup(3, map[string]interface{}{
    "id": 10,
}, map[string]string{})

if err != nil {
    log.Fatalf("Disassociate Host err: %s", err)
}

log.Printf("Disassociate Host ID: %d", 3)
```

> Delete Host

```go
result, err := client.HostService.DeleteHost(3)

if err != nil {
    log.Fatalf("Host err: %s", err)
}

log.Printf("Host deleted. Host ID: %d", 3)
```