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
			r := findPhotonQuery(file)
			log.Printf("r: %+v", r)
		}
	}
}

type Result struct {
	Name string
}

func findPhotonQuery(node *ast.File) (r Result) {
	ast.Inspect(node, func(n ast.Node) bool {
		_, i, ok := findMethods(n, "Exec")

		if !ok {
			return true
		}

		log.Printf("found %+v", i.Sel)

		r.Name = getName(i)

		findFields(i)

		return true
	})
	return
}

func findFields(node ast.Node) {
	ast.Inspect(node, func(n ast.Node) bool {
		_, i, ok := findMethods(n, "GroupBy", "Fields")
		if !ok {
			return true
		}

		log.Printf("n %+v", i.Sel.Name)

		extractArguments(i)

		return true
	})
}

func getName(sel *ast.SelectorExpr) (str string) {
	ast.Inspect(sel, func(n ast.Node) bool {
		r, _, ok := findMethods(n, "Name")
		if !ok {
			return true
		}

		arg := r.Args[0].(*ast.BasicLit)

		str = arg.Value

		return true
	})
	return
}

func extractArguments(sel *ast.SelectorExpr) {
	ast.Inspect(sel, func(n ast.Node) bool {

		// log.Printf("extracting args %+v", n)

		return true
	})
}
