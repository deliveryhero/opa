// Copyright 2023 The OPA Authors.  All rights reserved.
// Use of this source code is governed by an Apache2
// license that can be found in the LICENSE file.

package rego_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/deliveryhero/opa/ast"
	"github.com/deliveryhero/opa/rego"
	"github.com/deliveryhero/opa/topdown"
	"github.com/deliveryhero/opa/types"
)

// NOTE(sr): These test are here because the only cases where PrepareOption are
// used is outside of the rego package. Testing them within the rego package
// would be less realistic.
func TestPrepareOption(t *testing.T) {
	t.Run("BuiltinFuncs", func(t *testing.T) {
		bi := map[string]*topdown.Builtin{
			"count": {
				Decl: ast.BuiltinMap["count"],
				Func: topdown.GetBuiltin("count"),
			},
		}
		pc := &rego.PrepareConfig{}
		rego.WithBuiltinFuncs(bi)(pc)
		act, exp := pc.BuiltinFuncs(), bi
		if diff := cmp.Diff(exp, act,
			cmpopts.IgnoreUnexported(ast.Builtin{}, types.Function{}),
			cmpopts.IgnoreFields(topdown.Builtin{}, "Func")); diff != "" {
			t.Errorf("unexpected result (-want, +got):\n%s", diff)
		}
	})
}
