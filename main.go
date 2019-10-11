package main

import (
	"fmt"
	"go/ast"
	"io/ioutil"
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

			str := generate(r)
			err := writeFile(str)
			if err != nil {
				panic(err)
			}
		}
	}
}

func writeFile(s string) error {
	b := []byte(s)
	perm := os.FileMode(0644)

	err := ioutil.WriteFile("./example/photon/structs_gen.go", b, perm)
	if err != nil {
		return err
	}

	return nil
}

func generate(results []Result) string {
	s := `package photon

import (
	"context"
)

`

	for _, r := range results {
		s += fmt.Sprintf("type %s struct {\n", r.Name)
		for _, arg := range r.Args {
			var desc string
			switch arg.Origin {
			case "Sum":
				desc = "is the sum of integer values of selected rows"
			case "Group":
				desc = "was grouped by"
			case "Count":
				desc = "is the total count of results"
			case "Select":
				desc = "was selected"
			}
			s += fmt.Sprintf("  // %s %s.\n", arg.Field, desc)
			s += fmt.Sprintf("  %s %s\n", arg.Field, arg.Type)
		}
		s += fmt.Sprintf("}\n")

		s += fmt.Sprintf(`
// Exec runs the query and returns a result and an error
func (r PostMethodsSelect) Exec(ctx context.Context) ([]%s, error) {
	return []UserQueryA{{
		ID:        "123",
		Title:     "My Post",
		Likes:     "50",
		PostCount: 5,
	}}, nil
}
`, r.Name)
	}
	log.Printf("struct: \n%s", s)

	return s
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

		if r, _, ok := findMethod(n, "Fields"); ok {
			args = append(args, extractArguments(r)...)
		}

		if r, _, ok := findMethod(n, "GroupBy"); ok {
			args = append(args, extractArguments(r)...)
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

	// cut off quotes
	return str[1 : len(str)-1]
}

type Arg struct {
	Origin     string
	Field      string
	Collection string
	Type       string
}

func extractArguments(node *ast.CallExpr) []Arg {
	var args []Arg

	l.Printf("extracting arguments of %+v", node.Fun.(*ast.SelectorExpr).Sel.Name)

	for _, arg := range node.Args {
		a, ok := arg.(*ast.CallExpr)
		if !ok {
			continue
		}

		// method can be Select, Group, etc.
		method, ok := a.Fun.(*ast.SelectorExpr)
		if !ok {
			continue
		}

		var field string
		var col string
		var typ string

		switch x := method.X.(type) {
		case *ast.Ident:
			field = x.Name + "Count"
			typ = "int"
		case *ast.SelectorExpr:
			typ = "string"
			field = x.Sel.Name
			f, ok := x.X.(*ast.Ident)
			if !ok {
				continue
			}
			col = f.Name
		}

		// temporarily hardcoded since we don't have type information yet
		if field == "Likes" {
			typ = "int"
		}

		args = append(args, Arg{
			Origin:     method.Sel.Name,
			Field:      field,
			Type:       typ,
			Collection: col,
		})

	}

	l.Printf("args len %d", len(args))

	return args
}
