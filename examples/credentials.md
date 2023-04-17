# Credentials API

Please refer to `client.md` before reviewing these examples.

## Usage

> List Credentials

```go
result, _, err := client.CredentialsService.ListCredentials(map[string]string{})
if err != nil {
    log.Fatalf("List Credentials err: %s", err)
}

log.Println("List Credentials: ", result)
```

> Create Credentials

```go
result, err := client.CredentialsService.CreateCredentials(map[string]interface{}{
        "credential_type": 5, // AWS Credential
        "inputs": map[string]interface{}{
            "username": "badusername",
            "password": "badpassword",  // pragma: allowlist secret
        },
        "name":         "my_aws_creds",
        "organization": 1, // Default organisation
    }, map[string]string{})
if err != nil {
    t.Error(err)
}

log.Println("Create Credentials: ", result)
```

> Get Credentials by ID

```go
result, err := client.CredentialsService.GetCredentialsByID(1, map[string]string{})
if err != nil {
    t.Error(err)
}

log.Println("Obtained Credentials: ", result)
```

> Update Credentials by ID

```go
result, err := client.CredentialsService.UpdateCredentialsByID(1, map[string]interface{}{
        "inputs": map[string]interface{}{
            "username": "badusername",
            "password": "MUCH_BETTER_P4SS0WRD_RIGHT?",  // pragma: allowlist secret
        },
    }, map[string]string{})
if err != nil {
    t.Error(err)
}

log.Println("Updated Credentials: ", result)
```

> Delete Credentials by ID

```go
err := client.CredentialsService.DeleteCredentialsByID(1, map[string]string{})
if err != nil {
    t.Error(err)
}

log.Println("Credentials Deleted")
```
