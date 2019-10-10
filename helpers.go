package main

import (
	"go/ast"
)

func findMethod(node ast.Node, name string) (*ast.CallExpr, *ast.SelectorExpr, bool) {
	// find photon call
	ret, ok := node.(*ast.CallExpr)
	if !ok {
		return nil, nil, false
	}

	i, ok := ret.Fun.(*ast.SelectorExpr)
	if !ok {
		return nil, nil, false
	}

	if i.Sel.Name == name {
		l.Printf("found method %s", name)
		return ret, i, true
	}

	return nil, nil, false
}
