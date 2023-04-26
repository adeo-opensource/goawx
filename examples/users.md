# User API

Please refer to `client.md` before reviewing these examples.

## Usage

> List Users

```go
result, _, err := client.UserService.ListUsers(map[string]string{
    "name": "Demo User",
})
if err != nil {
    log.Fatalf("List Users err: %s", err)
}

log.Println("List User: ", result)
```

> Create User

```go
result, err := client.UserService.CreateUser(map[string]interface{}{
   "username":     "test",
   "description":  "for testing CreateUser api",
}, map[string]string{})

if err != nil {
    log.Fatalf("Create User err: %s", err)
}

log.Printf("User created. Username: %s", result.User.Username)
```

> Update User

```go
result, err := client.UserService.UpdateUser(1, map[string]interface{}{
   "description":  "for testing Update api",
}, map[string]string{})

if err != nil {
    log.Fatalf("Update User err: %s", err)
}

log.Printf("Update finised. Description: %s", result.User.Description)
```

> Delete User

```go
result, err := client.UserService.DeleteUser(1)

if err != nil {
    log.Fatalf("Delete user err: %s", err)
}

log.Printf("User Deleted")
```