package main

import (
	"fmt"
	"go/ast"
	"log"
	"os"

	"golang.org/x/tools/go/packages"
)

var l = log.New(os.Stdout, "", log.Ltime)

func main() {
	pkgs, err := packages.Load(&packages.Config{
		Mode: packages.LoadSyntax,
	}, "go-codegen/example/...")

	if err != nil {
		panic(err)
	}

	for _, pkg := range pkgs {
		l.Printf("pkg %+v", pkg.Name)

		for _, file := range pkg.Syntax {
			r := findPhotonQuery(file)
			l.Printf("")
			l.Printf("** end result **")
			l.Printf("r: %+v", r)
			l.Printf("")
			l.Printf("___________")
			generate(r)
		}
	}
}

func generate(results []Result) {
	var s string
	for _, r := range results {
		s += fmt.Sprintf("type %s struct {\n", r.Name)
		for _, arg := range r.Args {
			s += fmt.Sprintf("  %s %s\n", arg.Field, "string")
		}
		s += fmt.Sprintf("}\n")
	}
	log.Printf("struct: \n%s", s)
}

type Result struct {
	Name string
	Args []Arg
}

func findPhotonQuery(node *ast.File) []Result {
	var s []Result

	ast.Inspect(node, func(n ast.Node) bool {
		_, i, ok := findMethod(n, "Exec")

		if !ok {
			return true
		}

		var r Result
		r.Name = getName(i)
		r.Args = findFields(i)
		s = append(s, r)

		return true
	})

	return s
}

func findFields(node *ast.SelectorExpr) []Arg {
	var args []Arg

	log.Printf("finding fields for %s", node.Sel.Name)

	ast.Inspect(node, func(n ast.Node) bool {

		if _, i, ok := findMethod(n, "Fields"); ok {
			args = append(args, extractArguments(i)...)
		}

		if _, i, ok := findMethod(n, "GroupBy"); ok {
			args = append(args, extractArguments(i)...)
		}

		return true
	})

	return args
}

func getName(sel *ast.SelectorExpr) string {
	var str string

	ast.Inspect(sel, func(n ast.Node) bool {
		r, _, ok := findMethod(n, "Name")
		if !ok {
			return true
		}

		arg, ok := r.Args[0].(*ast.BasicLit)
		if !ok {
			return true
		}

		str = arg.Value

		return true
	})

	return str
}

type Arg struct {
	Origin     string
	Field      string
	Collection string
}

func extractArguments(sel *ast.SelectorExpr) []Arg {
	var args []Arg

	ast.Inspect(sel, func(n ast.Node) bool {
		log.Printf("1 %+v", n)
		call, ok := n.(*ast.CallExpr)

		if !ok {
			return true
		}

		l.Printf("extracting arguments of %+v", call.Fun.(*ast.SelectorExpr).Sel.Name)

		for _, arg := range call.Args {
			a, ok := arg.(*ast.CallExpr)
			if !ok {
				continue
			}

			// method can be SelectParent, Group, etc.
			method, ok := a.Fun.(*ast.SelectorExpr)
			if !ok {
				continue
			}

			var field string
			var col string

			switch x := method.X.(type) {
			case *ast.Ident:
				field = x.Name
			case *ast.SelectorExpr:
				field = x.Sel.Name
				f, ok := x.X.(*ast.Ident)
				if !ok {
					continue
				}
				col = f.Name
			}

			args = append(args, Arg{
				Origin:     method.Sel.Name,
				Field:      field,
				Collection: col,
			})

		}

		l.Printf("args len %d", len(args))

		return true
	})

	return args
}
