# psql

It's like https://github.com/jackc/pgx-gofrs-uuid but for `github.com/google/uuid` library. It allows you to
use `uuid.UUID` and `[]uuid.UUID` with `QueryExecModeExec` and `QueryExecModeSimpleProtocol` mode. 

## Usage

```go
    import pgxuuid "github.com/Timosha/pgx-google-uuid"

    pgx, err := pgxpool.ParseConfig(connString)
	pgx.AfterConnect = func(ctx context.Context, t testing.TB, conn *pgx.Conn) {
		pgxuuid.Register(conn.TypeMap())
	}
```
