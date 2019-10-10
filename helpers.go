package main

import (
	"go/ast"
	"log"
)

func findMethods(node ast.Node, names ...string) (*ast.CallExpr, *ast.SelectorExpr, bool) {
	// find photon call
	ret, ok := node.(*ast.CallExpr)
	if !ok {
		return nil, nil, false
	}

	i, ok := ret.Fun.(*ast.SelectorExpr)
	if !ok {
		log.Printf("2")
		return nil, nil, false
	}

	for _, n := range names {
		if i.Sel.Name == n {
			log.Printf("found method %+v", i.Sel)
			return ret, i, true
		}
	}

	return nil, nil, false
}
