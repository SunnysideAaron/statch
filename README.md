<div align="center">
  <img src="docs/logo/logo.jpg">
</div>

# statch

**S**QL **T**emplate M**atch**ing 

statch is a tool for matching data sources to templates for code generation.

What statch does:
- links data sources to templates
- provides tools to expose SQL schema and queries to the templating package

Originally inspired by [sqlc](https://github.com/sqlc-dev/sqlc). statch does allow
for code generation from SQL queries. But takes a higher level approach.

What separates statch from other code generators:
- statch isn't trying to dictate what code is generated. While a set of templates 
  are provided, it is expected that you will adjust them to generate what you want.
  vs sqlc which provides a master template but then keeps adding more and more
  settings to generate variations in code output. We believe this creates an
  endless cycle of users wanting more and more variations and ways to tweak the
  output. Creating more and more complexity. If you want different output use
  different templates.
- Different templating package. sqlc uses [template](https://pkg.go.dev/text/template).
  I personally find template gross. The code is ugly and difficult to read.
  "." notation is awful. Hiding where data is coming from. Others may like it
  but I had an immediate visceral reaction against it.

Initial Code Goals
- PostreSQL
- code generated for pgx
- which template engine?
- sources
  - sql queries
  - sql schema
  - [hjson](https://hjson.github.io/)

Future Goals
- any template engine
  - probably a stretch
- sqlite
- mysql
- ms sql server
- templates not compiled into statch
- source
  - json
  - yaml?
  - go structs / code?
  - OpenAPI 3.0


can't possibly predict what everyone will want their output code to look like

it's dumb


make bash
  build docker and run bash in container

first time will take awhile to build because of postgress parse dependency

make build  

