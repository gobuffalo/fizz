module github.com/gobuffalo/fizz

go 1.16

require (
	github.com/Masterminds/semver/v3 v3.1.1
	github.com/go-sql-driver/mysql v1.6.0
	github.com/gobuffalo/plush/v4 v4.1.9
	github.com/gobuffalo/pop/v6 v6.0.0
	github.com/jackc/pgx/v4 v4.13.0
	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51
	github.com/stretchr/testify v1.7.0
)

replace github.com/gobuffalo/pop/v6 v6.0.0 => github.com/fasmat/pop/v6 v6.0.0-20211121195140-6d95c111f911
