// Copyright 2017 The casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package jsonadapter

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"testing"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/util"
)

func testGetPolicy(t *testing.T, e *casbin.Enforcer, res [][]string) {
	myRes := e.GetPolicy()
	log.Print("Policy: ", myRes)

	if !util.Array2DEquals(res, myRes) {
		t.Error("Policy: ", myRes, ", supposed to be ", res)
	}
}

func TestAdapter(t *testing.T) {
	b, _ := ioutil.ReadFile(filepath.Join("examples", "rbac_policy.json"))
	a := NewAdapter(&b)
	e, _ := casbin.NewEnforcer("examples/rbac_model.conf", a)

	// Now the JSON Buffer has policy, so we can provide a normal use case.
	// Create an adapter and an enforcer.
	// NewEnforcer() will load the policy automatically.
	a = NewAdapter(&b)
	e, _ = casbin.NewEnforcer("examples/rbac_model.conf", a)
	testGetPolicy(t, e, [][]string{{"alice", "data1", "read"}, {"bob", "data2", "write"}, {"data2_admin", "data2", "read"}, {"data2_admin", "data2", "write"}})
}
