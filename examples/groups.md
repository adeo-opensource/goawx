# Group API

Please refer to `client.md` before reviewing these examples.

## Usage

> List Groups

```go
result, _, err := client.GroupService.ListGroups(map[string]string{
    "name": "Demo Group",
})
if err != nil {
    log.Fatalf("List Groups err: %s", err)
}

log.Println("List Groups: ", result)
```

> Create Group

```go
result, err := client.GroupService.CreateGroup(map[string]interface{}{
   "name":     "test",
   "description":  "for testing CreateGroup api",
}, map[string]string{})

if err != nil {
    log.Fatalf("Create Group err: %s", err)
}

log.Printf("Group created. Group ID: %d", result.Group.ID")
```

> Update Group

```go
result, err := client.GroupService.UpdateGroup(21, map[string]interface{}{
    "description": "Add description here",
}, map[string]string{})

if err != nil {
    log.Fatalf("Update Group err: %s", err)
}

log.Printf("Updated Group ID: %d", 21)
```

> Delete Group

```go
result, err := client.GroupService.DeleteGroup(12)

if err != nil {
    log.Fatalf("Delete Group err: %s", err)
}

log.Printf("Group deleted. Group ID: %d", 12)
```