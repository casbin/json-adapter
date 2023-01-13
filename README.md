JSON Adapter   
[![Go Report Card](https://goreportcard.com/badge/github.com/casbin/json-adapter)](https://goreportcard.com/report/github.com/casbin/json-adapter) [![Coverage Status](https://coveralls.io/repos/github/casbin/json-adapter/badge.svg?branch=master)](https://coveralls.io/github/casbin/json-adapter?branch=master) [![Godoc](https://godoc.org/github.com/casbin/json-adapter?status.svg)](https://godoc.org/github.com/casbin/json-adapter) [![Release](https://img.shields.io/github/release/casbin/json-adapter.svg)](https://github.com/casbin/json-adapter/releases/latest) [![Gitter](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/casbin/lobby) [![Sourcegraph](https://sourcegraph.com/github.com/casbin/json-adapter/-/badge.svg)](https://sourcegraph.com/github.com/casbin/json-adapter?badge)
====

JSON Adapter is the [JSON (JavaScript Object Notation)](https://www.json.org/) adapter for [Casbin](https://github.com/casbin/casbin). With this library, Casbin can load policy from JSON string or save policy to it.

## Installation

    go get github.com/casbin/json-adapter

## Simple Example

```go
package main

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/json-adapter/v2"
)

func main() {
	// Initialize a JSON adapter and use it in a Casbin enforcer:
	b := []byte{} // b stores Casbin policy in JSON bytes.
	a := jsonadapter.NewAdapter(&b) // Use b as the data source. 
	e, _ := casbin.NewEnforcer("examples/rbac_model.conf", a)
	
	// Load the policy from JSON bytes b.
	e.LoadPolicy()
	
	// Check the permission.
	e.Enforce("alice", "data1", "read")
	
	// Modify the policy.
	// e.AddPolicy(...)
	// e.RemovePolicy(...)
	
	// Save the policy back to JSON bytes b.
	e.SavePolicy()
}
```

## Policy JSON

The following illustrates the expected JSON format for a policy.  The [rbac_policy.json](examples/rbac_policy.json) has the same policy found in [rbac_policy.csv](examples/rbac_policy.csv).

```json
[
  {"PType":"p","V0":"alice","V1":"data1","V2":"read"},
  {"PType":"p","V0":"bob","V1":"data2","V2":"write"},
  {"PType":"p","V0":"data2_admin","V1":"data2","V2":"read"},
  {"PType":"p","V0":"data2_admin","V1":"data2","V2":"write"},
  {"PType":"g","V0":"alice","V1":"data2_admin"}
]
```

## Getting Help

- [Casbin](https://github.com/casbin/casbin)

## License

This project is under Apache 2.0 License. See the [LICENSE](LICENSE) file for the full license text.
