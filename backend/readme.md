
## Usage

* `go run entgo.io/ent/cmd/ent init <Table>` to generate a new ORM table in `ent/schema/table.go`
* `go generate ./ent` to generate query functions for that table

## File Structure Breakdown

```
.
├── ent
│   ├── schema         # defines ORM objects
│   └── generate.go    # defines `go generate ent` command 
├── graph
│   ├── model                # defines structs that can be used to generate `schema.graphqls`
│   ├── resolver.go          # defines global state
│   ├── schema.graphqls      # schema definition (generated)
│   └── schema.resolvers.go  # defines resolvers
├── go.mod
├── go.sum
├── gqlgen.yml      # configures gqlgen
├── readme.md
└── server.go
```

* Once you have modified the schema, run `bazel run //backend:ent` to configure the rest of the ORM client

