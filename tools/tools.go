// +build tools

package tools

import (
    _ "github.com/golang/protobuf/protoc-gen-go"
    _ "sigs.k8s.io/kind"
)

// https://github.com/golang/go/issues/25922#issuecomment-476341329
