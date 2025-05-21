module statch

go 1.23.5

require (
	github.com/ajitpratap0/GoSQLX v0.0.0-00010101000000-000000000000
	github.com/gobuffalo/plush/v5 v5.0.4
	github.com/hjson/hjson-go/v4 v4.4.0
	github.com/jackc/pgx/v5 v5.7.5
)

replace github.com/ajitpratap0/GoSQLX => /statch/GoSQLX

require (
	github.com/gobuffalo/flect v1.0.2 // indirect
	github.com/stretchr/testify v1.10.0 // indirect
)
