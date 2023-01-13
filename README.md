Showcase an M2M relation, but with a custom join table.
The goal here though, is to now use `PascalCasing` for the table names.
Redefining the names doesn't seem to work though.

## Generate Assets

```console
go generate ./ent/
```

## Run

```console
go run .
```

For inspection of the resulting schemata, open the sqlite.db with whatever
tool you want to use.
