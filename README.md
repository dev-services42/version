# Version for your CLI programs

A simple Go package to manage version information in your CLI applications using build-time flags.

## How it works

This package uses Go's `-ldflags` to inject version information at build time. The version is stored in a package-level
variable that can be set during compilation.

The `GetVersion()` function returns the version using this priority:

1. Version set via `-ldflags` at build time
2. Version from `debug.ReadBuildInfo()` (when using Go modules)
3. `"unknown-local"` as fallback for local development

## Usage

### 1. Import the package

```go
package main

import "github.com/dev-services42/version"

func main() {
	fmt.Printf("Version: %s\n", version.GetVersion())
}
```

### 2. Build with version information

Use `-ldflags` to set the version at build time:

```bash
go build -ldflags "-X github.com/dev-services42/version.version=v1.2.3"
```
