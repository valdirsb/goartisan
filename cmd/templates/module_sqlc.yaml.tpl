version: "2"
sql:
  - engine: "mysql"
    schema: "../sqlc/schema.sql"
    queries: "sql/queries"
    gen:
      go:
        package: "db"
        out: "pkg/db"