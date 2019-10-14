# Go codegen minimal prototype

This is a small prototype to show how Photon Go (a generated database client) can help the developer by analyzing complex queries and generating types.

## Usage

Write your query in `example/example.go`:

```go
var result MyQuery1
err = client.Post.Select.Name("MyQuery1").Fields(
  Post.Likes.Sum(),
).GroupBy(
  Post.Title.Group(),
).Into(&result).Exec(ctx)
if err != nil {
  panic(err)
}
```

To generate the structs (i.e. `MyQuery1`), run `go run .` in the project root. This generates `./example/photon/structs_gen.go` with structs for your queries.

When you add or remove query parameters (i.e. `Post.Count()`), you have to re-run `go run .`. 

Later, this will be a separate cli tool and/or built-in in the Prisma cli as `prisma generate`.
