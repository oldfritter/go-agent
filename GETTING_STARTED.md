# Getting Started

Follow these steps to instrument your application.  More information is
available in the [GUIDE.md](GUIDE.md).

## Step 0: Installation

The New Relic Go agent is a Go library. It has two dependencies on gRPC
libraries - see [go.mod](v3/go.mod). Install the Go agent the same way you
would install any other Go library. The simplest way is to run:

```
go get github.com/oldfritter/go-agent
```

Then import the package in your application:
```go
import "github.com/oldfritter/go-agent/v3/oldfritter"
```

## Step 1: Create an Application

In your `main` function, or an `init` block, create an
[Application](https://godoc.org/github.com/oldfritter/go-agent/v3/oldfritter/#Application) using
 [ConfigOptions](https://godoc.org/github.com/oldfritter/go-agent/v3/oldfritter/#ConfigOption).
 Available configurations are listed [here](https://godoc.org/github.com/oldfritter/go-agent/v3/oldfritter/#Config).
[Application](https://godoc.org/github.com/oldfritter/go-agent/v3/oldfritter/#Application) is the
starting point for all instrumentation.

```go
func main() {
    // Create an Application:
    app, err := oldfritter.NewApplication(
        // Name your application
        oldfritter.ConfigAppName("Your Application Name"),
        // Fill in your New Relic license key
        oldfritter.ConfigLicense("__YOUR_NEW_RELIC_LICENSE_KEY__"),
        // Add logging:
        oldfritter.ConfigDebugLogger(os.Stdout),
        // Optional: add additional changes to your configuration via a config function:
        func(cfg *oldfritter.Config) {
            cfg.CustomInsightsEvents.Enabled = false
        },
    )
    // If an application could not be created then err will reveal why.
    if err != nil {
        fmt.Println("unable to create New Relic Application", err)
    }
    // Now use the app to instrument everything!
}
```

Now start your application, and within minutes it will appear in the New Relic
UI.  Your application in New Relic won't contain much data (until we complete
the steps below!), but you will already be able to see a
[Go runtime](https://docs.oldfritter.com/docs/agents/go-agent/features/go-runtime-page-troubleshoot-performance-problems)
page that shows goroutine counts, garbage collection, memory, and CPU usage.

## Step 2: Instrument Requests Using Transactions

[Transactions](https://godoc.org/github.com/oldfritter/go-agent/v3/oldfritter/#Transaction) are
used to time inbound requests and background tasks.  Use them to see your
application's throughput and response time.  The instrumentation strategy
depends on the framework you're using:

#### Standard HTTP Library

If you are using the standard library `http` package, use
[WrapHandle](https://godoc.org/github.com/oldfritter/go-agent/v3/oldfritter/#WrapHandle) and
[WrapHandleFunc](https://godoc.org/github.com/oldfritter/go-agent/v3/oldfritter/#WrapHandleFunc).
As an example, the following code:

```go
http.HandleFunc("/users", usersHandler)
```
Can be instrumented like this:
```go
http.HandleFunc(oldfritter.WrapHandleFunc(app, "/users", usersHandler))
```

[Full Example Application](./v3/examples/server/main.go)

#### Popular Web Framework

If you are using a popular framework, then there may be an integration package
designed to instrument it.  [List of New Relic Go agent integration packages](./README.md#integrations).

#### Manual Transactions

If you aren't using the `http` standard library package or an
integration package supported framework, you can create transactions
directly using the application's `StartTransaction` method:

```go
func myHandler(rw http.ResponseWriter, req *http.Request) {
    txn := h.App.StartTransaction("myHandler")
    defer txn.End()
    // Setting the response writer and request is optional. If you don't
    // set the request, the transaction is considered a background task.
    txn.SetWebRequestHTTP(req)
    // Use the ResponseWriter returned in place of the previous ResponseWriter
    rw = txn.SetWebResponse(rw)
    rw.Write(data)
}
```

Be sure to use a limited set of unique names to ensure that transactions are
grouped usefully.  Don't use dynamic URLs!

[More information about transactions](GUIDE.md#transactions)

## Step 3: Instrument Segments

Segments show you where the time in your transactions is being spent.  There are
four types of segments:
[Segment](https://godoc.org/github.com/oldfritter/go-agent/v3/oldfritter/#Segment),
[ExternalSegment](https://godoc.org/github.com/oldfritter/go-agent/v3/oldfritter/#ExternalSegment),
[DatastoreSegment](https://godoc.org/github.com/oldfritter/go-agent/v3/oldfritter/#DatastoreSegment),
and
[MessageProducerSegment](https://godoc.org/github.com/oldfritter/go-agent/v3/oldfritter/#MessageProducerSegment).

Creating a segment requires access to the transaction.  You can pass the
transaction around your functions inside
a [context.Context](https://golang.org/pkg/context/#Context) (preferred), or as an explicit transaction
parameter of the function.  Functions
[FromContext](https://godoc.org/github.com/oldfritter/go-agent/v3/oldfritter/#FromContext)
and [NewContext](https://godoc.org/github.com/oldfritter/go-agent/v3/oldfritter/#NewContext) make it
easy to store and retrieve the transaction from a context.

You may not even need to add the transaction to the context:
[WrapHandle](https://godoc.org/github.com/oldfritter/go-agent/v3/oldfritter/#WrapHandle) and
[WrapHandleFunc](https://godoc.org/github.com/oldfritter/go-agent/v3/oldfritter/#WrapHandleFunc)
add the transaction to the request's context automatically.

```go
func instrumentMe(ctx context.Context) {
    txn := oldfritter.FromContext(ctx)
    segment := txn.StartSegment("instrumentMe")
    time.Sleep(1 * time.Second)
    segment.End()
}

func myHandler(w http.ResponseWriter, r *http.Request) {
    instrumentMe(r.Context())
}

func main() {
    app, _ := oldfritter.NewApplication(
        oldfritter.ConfigAppName("appName"),
        oldfritter.ConfigLicense("__license__"),
    )
    http.HandleFunc(oldfritter.WrapHandleFunc(app, "/handler", myHandler))
}
```

[More information about segments](GUIDE.md#segments)

## Extra Credit

Read our [GUIDE.md](GUIDE.md) and the
[godocs](https://godoc.org/github.com/oldfritter/go-agent/v3/oldfritter) to learn more about
what else you can do with the Go Agent.
