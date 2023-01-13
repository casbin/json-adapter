JSON Adapter   
<p align="center">
  <a href="#badge">
    <img alt="semantic-release" src="https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg">
  </a>
  <a href="https://github.com/casbin/json-adapter/actions/workflows/default.yml">
    <img alt="GitHub Workflow Status (branch)" src="https://github.com/casbin/json-adapter/workflows/build/badge.svg?style=flat-square">
  </a>
  <a href="https://github.com/casbin/json-adapter/releases/latest">
    <img alt="GitHub Release" src="https://img.shields.io/github/v/release/casbin/json-adapter.svg">
  </a>
</p>

<p align="center">
  <a href="https://goreportcard.com/report/github.com/casbin/json-adapter">
    <img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/casbin/json-adapter?style=flat-square">
  </a>
  <a href="https://github.com/casbin/json-adapter/blob/master/LICENSE">
    <img src="https://img.shields.io/github/license/casbin/json-adapter?style=flat-square" alt="license">
  </a>
  <a href="https://github.com/casbin/json-adapter/issues">
    <img alt="GitHub issues" src="https://img.shields.io/github/issues/casbin/json-adapter?style=flat-square">
  </a>
  <a href="#">
    <img alt="GitHub stars" src="https://img.shields.io/github/stars/casbin/json-adapter?style=flat-square">
  </a>
  <a href="https://github.com/casbin/json-adapter/network">
    <img alt="GitHub forks" src="https://img.shields.io/github/forks/casbin/json-adapter?style=flat-square">
  </a>
</p>
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
