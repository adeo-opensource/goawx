# Project API

Please refer to `client.md` before reviewing these examples.

## Usage

> List Projects

```go
result, _, err := client.ProjectService.ListProjects(map[string]string{
    "name": "Demo Project",
})
if err != nil {
    log.Fatalf("List Projects err: %s", err)
}

log.Println("List Projects: ", result)
```

> Create Project

```go
result, err := client.ProjectService.CreateProject(map[string]interface{}{
   "name":         "TestProject",
   "description":  "for testing CreateProject api",
   "organization": 1,
   "kind":         "",
   "host_filter":  "",
   "variables":    "",
}, map[string]string{})

if err != nil {
    log.Fatalf("Create Projects err: %s", err)
}

log.Printf("Project created. Project ID: %d", result.Project.ID)
```

> Update Project

```go
result, err := client.ProjectService.UpdateProject(4, map[string]interface{}{
   "description":  "Update Example",
}, map[string]string{})

if err != nil {
    log.Fatalf("Update Projects err: %s", err)
}

log.Printf("Project Updated. Project ID: %d", result.ID)
```

> Delete Project

```go
result, err := client.ProjectService.DeleteProject(4)

if err != nil {
    log.Fatalf("Delete Projects err: %s", err)
}

log.Printf("Project Deleted. Project ID: %d", result.ID)
```
