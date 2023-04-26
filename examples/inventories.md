# Inventories API

Please refer to `client.md` before reviewing these examples.

## Usage

> List Inventories

```go
result, _, err := client.InventoriesService.ListInventories(map[string]string{
    "name": "Demo Inventory",
})
if err != nil {
    log.Fatalf("List Inventories err: %s", err)
}

log.Println("List Inventories: ", result)
```

> Create Inventory

```go
result, err := client.InventoriesService.CreateInventory(map[string]interface{}{
   "name":         "TestInventory",
   "description":  "for testing CreateInventory api",
   "organization": 1,
   "kind":         "",
   "host_filter":  "",
   "variables":    "",
}, map[string]string{})

if err != nil {
    log.Fatalf("Create Inventories err: %s", err)
}

log.Printf("Inventory created. Inventory ID: %d", result.Inventory.ID")
```

> Update Inventory

```go
result, err := client.InventoriesService.UpdateInventory(1, map[string]interface{}{
   "name":         "TestInventory - 1",
   "description":  "for testing UpdateInventory api",
   "organization": 1,
   "kind":         "",
   "host_filter":  "",
   "variables":    "",
}, nil)

if err != nil {
    log.Fatalf("Update Inventories err: %s", err)
}

log.Printf("Update result %v", result.Name)
```

> Delete Inventory

```go
result, err := client.InventoriesService.DeleteInventory(5)
if err != nil {
    log.Fatalf("Delete Inventories err: %s", err)
}

log.Println("Inventroy 5 Deleted")
```

> GetInventory

```go
result, err := client.InventoriesService.GetInventory(1, map[string]string{})
if err != nil {
    log.Fatalf("Get Inventory by Name err: %s", err)
}

log.Printf("Demo Inventory id %d\n", result.ID)
```