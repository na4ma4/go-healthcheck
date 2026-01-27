# go-healthcheck

[![Go Reference](https://pkg.go.dev/badge/github.com/na4ma4/go-healthcheck.svg)](https://pkg.go.dev/github.com/na4ma4/go-healthcheck)
[![Go Report Card](https://goreportcard.com/badge/github.com/na4ma4/go-healthcheck)](https://goreportcard.com/report/github.com/na4ma4/go-healthcheck)

A flexible and thread-safe Go package for tracking and reporting application component health and lifecycles. Perfect for microservices, long-running applications, and systems that need to expose health status via HTTP endpoints.

## Features

- **Thread-Safe**: Built with concurrent access in mind using proper mutex locking
- **Lifecycle Tracking**: Monitor components through their entire lifecycle (starting, running, finished, errored)
- **HTTP Handler**: Built-in HTTP handler for easy health endpoint integration
- **Flexible Reporting**: Standard and simplified reporting modes with optional verbose output
- **Event History**: Track lifecycle events with timestamps
- **Status Priority**: Automatic determination of overall health based on component status priorities
- **Protocol Buffers**: Support for protobuf serialization

## Installation

```bash
go get github.com/na4ma4/go-healthcheck
```

## Quick Start

```go
package main

import (
    "net/http"
    "time"
    
    health "github.com/na4ma4/go-healthcheck"
)

func main() {
    // Create a new health checker
    hc := health.NewCore()
    
    // Start tracking a component
    database := hc.Get("database").Start()
    
    // Simulate some work
    time.Sleep(100 * time.Millisecond)
    
    // Mark as finished or report an error
    database.Stop()
    // or: database.Error(err)
    
    // Set up HTTP health endpoint
    http.HandleFunc("/health", health.Handler(hc))
    http.ListenAndServe(":8080", nil)
}
```

## Usage

### Creating a Health Checker

```go
hc := health.NewCore()
```

### Tracking Components

```go
// Start tracking a component
db := hc.Get("database").Start()

// Report an error
if err := connectDatabase(); err != nil {
    db.Error(err)
}

// Stop a component
hc.Stop("database")
// or: db.Stop()
```

### HTTP Health Endpoint

The package provides a ready-to-use HTTP handler:

```go
http.HandleFunc("/health", health.Handler(hc))
```

#### Query Parameters

- `?simple=true` - Use simplified status (green/yellow/red)
- `?verbose=true` - Include start times and lifecycle events
- `?lifecycle=true` - Include lifecycle events

#### Example Responses

**Standard Mode** (`/health`):
```json
{
  "status": "running",
  "services": [
    {
      "name": "database",
      "status": "running"
    },
    {
      "name": "cache",
      "status": "running"
    }
  ]
}
```

**Simple Mode** (`/health?simple=true`):
```json
{
  "status": "green",
  "services": [
    {
      "name": "database",
      "status": "green"
    },
    {
      "name": "cache",
      "status": "green"
    }
  ]
}
```

**Verbose Mode** (`/health?verbose=true`):
```json
{
  "status": "running",
  "services": [
    {
      "name": "database",
      "status": "running",
      "start_time": "2026-01-27T10:15:30.123456789Z",
      "lifecycle": [
        {
          "ts": "2026-01-27T10:15:30.123456789Z",
          "status": "starting"
        },
        {
          "ts": "2026-01-27T10:15:30.234567890Z",
          "status": "running"
        }
      ]
    }
  ]
}
```

### Status Types

#### Standard Status
- `StatusStarting` - Component is initializing
- `StatusRunning` - Component is healthy and running
- `StatusFinished` - Component has completed successfully
- `StatusErrored` - Component encountered an error
- `StatusUnknown` - Component status is unknown

#### Simple Status (for simple mode)
- `ReportStatusGreen` - All components healthy
- `ReportStatusYellow` - Some components starting or finishing
- `ReportStatusRed` - One or more components errored

### Checking Component Status

```go
// Get status of all components
status := hc.Status()
// Returns: map[string]bool{"database": true, "cache": false}

// Check if a specific component is healthy
item := hc.Get("database")
isHealthy := health.StatusIsHealthy(item.Status())
```

### Iterating Over Components

```go
err := hc.Iterate(func(name string, item health.Item) error {
    fmt.Printf("%s: %s (duration: %s)\n", 
        name, 
        item.Status(), 
        item.Duration(),
    )
    return nil
})
```

### Accessing Item Details

```go
item := hc.Get("database")

// Get component details
name := item.Name()
duration := item.Duration()
startTime := item.StartTime()
status := item.Status()
lifecycle := item.Lifecycle()
```

## Advanced Usage

### Custom Health Checks with Goroutines

```go
func monitorDatabase(hc health.Health) {
    db := hc.Get("database").Start()
    
    go func() {
        ticker := time.NewTicker(30 * time.Second)
        defer ticker.Stop()
        
        for range ticker.C {
            if err := pingDatabase(); err != nil {
                db.Error(err)
                return
            }
        }
    }()
}
```

### Protocol Buffers Export

```go
// Export to protobuf
proto := hc.ToProto()
// proto is *CoreProto
```

## Status Priority

The package automatically determines overall health based on status priority (highest to lowest):
1. Errored
2. Unknown
3. Starting
4. Running
5. Finished

When generating reports, the overall status will be the highest priority status among all components.

## Thread Safety

All operations on the `Core` type are thread-safe and can be called from multiple goroutines simultaneously.

## Requirements

- Go 1.23 or higher
- Dependencies:
  - `github.com/google/go-cmp` (testing)
  - `google.golang.org/protobuf` (protobuf support)

## License

See the LICENSE file for details.

## Contributing

Contributions are welcome! Please feel free to submit issues or pull requests.
