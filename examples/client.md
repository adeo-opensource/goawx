# Creating a Client

When using this SDK you're going to need to set up a client and establish a connection to the remote AWX/Tower
instance. Here's an example of how you might go about doing this:

```go
package main

import (
    "log"
    "github.com/adeo-opensource/goawx/client"
)

func main() {
    client := awx.NewAWX("http://awx.your_server_host.com", "your_awx_username", "your_awx_passwd", nil)
    // ...
}
```

There are four parameters you can provide:

* Your AWX/Tower DNS hostname or IP address (remembering to include the `http://` too)
* The username you wish to authenticate as
* The password you wish to authenticate with
* And an optional `*http.Client` you can use to custom how the SDK communicates with your AWX/Tower instance(s)

Throughout the rest of these example documents the above `client` variable will be referred to as a correctly
configured client to an operational AWX/Tower instance.
