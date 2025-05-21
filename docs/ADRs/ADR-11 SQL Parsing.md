# ADR-11 SQL Parsing

## Status

Accepted, Proposed, Deprecated or Superseded (list DR)

## Context

SQL lexing, parsing, and Abstract Syntax Tree (AST) generation.

## Decision

[GoSQLX](https://github.com/ajitpratap0/GoSQLX)

## Why

Honestly, just rolling the dice. Too new to have a community but I'm liking the
brochure. Other parsers are old or would require copying out of other projects.

## Notes

[Introducing GoSQLX: A SQL Query Parser for Go (Not a Replacement for sqlx)](https://medium.com/@ajitpratapsingh/introducing-gosqlx-a-sql-query-parser-for-go-not-a-replacement-for-sqlx-1cfc2bf52d52)

## Consequences


## Other Possible Options

- [gosqlparser](https://github.com/krasun/gosqlparser)
  - older but nice and simple.
- [GoSQLX](https://github.com/ajitpratap0/GoSQLX)
  - [Introducing GoSQLX: A SQL Query Parser for Go (Not a Replacement for sqlx)](https://medium.com/@ajitpratapsingh/introducing-gosqlx-a-sql-query-parser-for-go-not-a-replacement-for-sqlx-1cfc2bf52d52)
  - New, but I'm liking the brochure.
- [sqlc](https://github.com/sqlc-dev/sqlc)
  - Would have to rip out and copy a lot of code but it is possible. Could
    release as it's own library? The maintainers of that library might not like that...
- [sqlparser](https://github.com/marianogappa/sqlparser)
  - [Let's build a SQL parser in Go!](https://marianogappa.github.io/software/2019/06/05/lets-build-a-sql-parser-in-go/)
  - tutorial on building a simple sql parser.
- [pg_query_go](https://github.com/pganalyze/pg_query_go)
  - Uses the actual PostgreSQL server source to parse SQL queries and return the internal PostgreSQL parse tree

## Not an Option

- [Vitess](https://github.com/vitessio/vitess)
  - not a parser. for MySQL. Probably has a parser in it.
- [vitess-sqlparser](https://github.com/blastrain/vitess-sqlparser)
  - SQL and DDL parser for Go (powered by vitess and TiDB)
  - Old but also lists several other options.

## sqlc Notes

These notes may or may not be true. Written by AI. On the surface seems probable.
Saving in case we decide to use sqlc's parser down the road.

sqlc uses different parsing strategies depending on the database engine:

### Parser Architecture
- Parser interface defined in `internal/compiler/compile.go` (around line 20)
- Engine selection in `internal/compiler/engine.go` (around line 40)

### Engine-specific Parsers
1. **PostgreSQL**:
   - Custom parser implementation in `internal/engine/postgresql/parse.go`
   - Uses PostgreSQL's own parser code to generate an AST
   - Translation layer converts PostgreSQL AST to sqlc's common AST format

2. **SQLite**:
   - ANTLR4-based parser in `internal/engine/sqlite/parse.go`
   - Grammar defined in `internal/engine/sqlite/parser/SQLite*.g4` files
   - Converter in `internal/engine/sqlite/convert.go` transforms ANTLR parse tree to sqlc AST

3. **MySQL/Dolphin**:
   - Uses TiDB's parser in `internal/engine/dolphin/parse.go`
   - Converter transforms TiDB AST to sqlc's common AST format

### Parsing Process
- Schema parsing in `internal/compiler/compile.go:parseCatalog()` (around line 100)
- Query parsing in `internal/compiler/compile.go:parseQueries()` (around line 150)
- AST manipulation utilities in `internal/sql/astutils/`

### Common AST Format
- Common AST nodes defined in `internal/sql/ast/` directory
- Allows sqlc to use a single code generation backend regardless of source database

This architecture allows sqlc to support multiple database engines while maintaining a unified code generation pipeline. The approach of using each database's native parser (or close equivalent) ensures high fidelity SQL parsing.
