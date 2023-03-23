# Job Template API

Please refer to `client.md` before reviewing these examples.

## Usage

> List Job Templates

```go
result, _, err := client.JobTemplateService.ListJobTemplates(map[string]string{})
if err != nil {
    log.Fatalf("List Job Templates err: %s", err)
}

log.Println("List Job Templates: ", result)
```

> Launch Job Template

```go
result, err := client.JobTemplateService.Launch(yourJobTemplateId, map[string]interface{}{
    "inventory": yourInventoryId,
}, map[string]string{})
if err != nil {
    log.Fatalf("Lauch err: %s", err)
}

log.Println("Launch Job Template: ", result)
```

> Create Job Template

```go
result, err := client.JobTemplateService.CreateJobTemplate(map[string]interface{}{
    "name":        "Example Create Job Template",
    "description": "Created from awx-go Example",
    "job_type":    "run",
    "inventory":   1,
    "project":     1,
    "playbook":    "playbook.yml",
    "verbosity":   0,
}, map[string]string{})

if err != nil {
    log.Fatalf("Create job template err: %s", err)
}
log.Printf("Job template created. JobTemplate ID: %d", result.ID)
```

> Update Job Template

```go
result, err := client.JobTemplateService.UpdateJobTemplate(5, map[string]interface{}{
    "description": "Update Job Template",
}, map[string]string{})

if err != nil {
    log.Fatalf("Update job template err: %s", err)
}
log.Printf("Job template Updated. JobTemplate ID: %d", result.ID)
```

> Delete Job Template

```go
result, err := client.JobTemplateService.DeleteJobTemplate(5)

if err != nil {
    log.Fatalf("Delete job template err: %s", err)
}
log.Printf("Job template Deleted. JobTemplate ID: %d", result.ID)
```