package main

import (
	"go/ast"
	"log"

	"golang.org/x/tools/go/packages"
)

func main() {
	pkgs, err := packages.Load(&packages.Config{
		Mode: packages.LoadSyntax,
	}, "go-codegen/example/...")

	if err != nil {
		panic(err)
	}

	for _, pkg := range pkgs {
		log.Printf("pkg %+v", pkg.Name)

		for _, file := range pkg.Syntax {
			// log.Printf("key %+v", file)

			findPhotonQuery(file)
		}
	}
}

func findPhotonQuery(node ast.Node) {
	ast.Inspect(node, func(n ast.Node) bool {
		// find photon call
		ret, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}

		i, ok := ret.Fun.(*ast.SelectorExpr)
		if !ok {
			return true
		}

		if i.Sel.Name != "Exec" {
			return true
		}

		log.Printf("found %+v", i.Sel)

		findFields(i)

		return true
	})
}

func findFields(node ast.Node) {
	ast.Inspect(node, func(n ast.Node) bool {
		ret, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}

		i, ok := ret.Fun.(*ast.SelectorExpr)
		if !ok {
			return true
		}

		if i.Sel.Name != "GroupBy" && i.Sel.Name != "Fields" {
			return true
		}

		log.Printf("n %+v", i.Sel.Name)

		return true
	})
}
