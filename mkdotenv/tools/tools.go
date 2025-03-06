// go:build tools
package tools

import (
    _ "github.com/go-delve/delve/cmd/dlv"    // Debugging tool
    _ "github.com/golangci/golangci-lint/cmd/golangci-lint" // Linter
)
