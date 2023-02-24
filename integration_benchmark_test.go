// Code generated by erb. DO NOT EDIT.

package uuid_test

import (
	"context"
	"testing"

	"github.com/gofrs/uuid/v5"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func BenchmarkQueryDecode_PG_UUID_to_Go_pgtype_UUID_1_rows_1_columns(b *testing.B) {
	defaultConnTestRunner.RunTest(context.Background(), b, func(ctx context.Context, _ testing.TB, conn *pgx.Conn) {
		b.ResetTimer()
		var v [1]pgtype.UUID
		for i := 0; i < b.N; i++ {
			_, err := conn.QueryFunc(
				ctx,
				`select gen_random_uuid() from generate_series(1, 1) n`,
				nil,
				[]interface{}{&v[0]},
				func(pgx.QueryFuncRow) error { return nil },
			)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkQueryDecode_PG_UUID_to_Go_pgtype_UUID_1_rows_10_columns(b *testing.B) {
	defaultConnTestRunner.RunTest(context.Background(), b, func(ctx context.Context, _ testing.TB, conn *pgx.Conn) {
		b.ResetTimer()
		var v [10]pgtype.UUID
		for i := 0; i < b.N; i++ {
			_, err := conn.QueryFunc(
				ctx,
				`select gen_random_uuid(), gen_random_uuid(), gen_random_uuid(), gen_random_uuid(), gen_random_uuid(), gen_random_uuid(), gen_random_uuid(), gen_random_uuid(), gen_random_uuid(), gen_random_uuid() from generate_series(1, 1) n`,
				nil,
				[]interface{}{&v[0], &v[1], &v[2], &v[3], &v[4], &v[5], &v[6], &v[7], &v[8], &v[9]},
				func(pgx.QueryFuncRow) error { return nil },
			)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkQueryDecode_PG_UUID_to_Go_pgtype_UUID_10_rows_1_columns(b *testing.B) {
	defaultConnTestRunner.RunTest(context.Background(), b, func(ctx context.Context, _ testing.TB, conn *pgx.Conn) {
		b.ResetTimer()
		var v [1]pgtype.UUID
		for i := 0; i < b.N; i++ {
			_, err := conn.QueryFunc(
				ctx,
				`select gen_random_uuid() from generate_series(1, 10) n`,
				nil,
				[]interface{}{&v[0]},
				func(pgx.QueryFuncRow) error { return nil },
			)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkQueryDecode_PG_UUID_to_Go_pgtype_UUID_100_rows_10_columns(b *testing.B) {
	defaultConnTestRunner.RunTest(context.Background(), b, func(ctx context.Context, _ testing.TB, conn *pgx.Conn) {
		b.ResetTimer()
		var v [10]pgtype.UUID
		for i := 0; i < b.N; i++ {
			_, err := conn.QueryFunc(
				ctx,
				`select gen_random_uuid(), gen_random_uuid(), gen_random_uuid(), gen_random_uuid(), gen_random_uuid(), gen_random_uuid(), gen_random_uuid(), gen_random_uuid(), gen_random_uuid(), gen_random_uuid() from generate_series(1, 100) n`,
				nil,
				[]interface{}{&v[0], &v[1], &v[2], &v[3], &v[4], &v[5], &v[6], &v[7], &v[8], &v[9]},
				func(pgx.QueryFuncRow) error { return nil },
			)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkQueryDecode_PG_UUID_to_Go_uuid_UUID_1_rows_1_columns(b *testing.B) {
	defaultConnTestRunner.RunTest(context.Background(), b, func(ctx context.Context, _ testing.TB, conn *pgx.Conn) {
		b.ResetTimer()
		var v [1]uuid.UUID
		for i := 0; i < b.N; i++ {
			_, err := conn.QueryFunc(
				ctx,
				`select gen_random_uuid() from generate_series(1, 1) n`,
				nil,
				[]interface{}{&v[0]},
				func(pgx.QueryFuncRow) error { return nil },
			)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkQueryDecode_PG_UUID_to_Go_uuid_UUID_1_rows_10_columns(b *testing.B) {
	defaultConnTestRunner.RunTest(context.Background(), b, func(ctx context.Context, _ testing.TB, conn *pgx.Conn) {
		b.ResetTimer()
		var v [10]uuid.UUID
		for i := 0; i < b.N; i++ {
			_, err := conn.QueryFunc(
				ctx,
				`select gen_random_uuid(), gen_random_uuid(), gen_random_uuid(), gen_random_uuid(), gen_random_uuid(), gen_random_uuid(), gen_random_uuid(), gen_random_uuid(), gen_random_uuid(), gen_random_uuid() from generate_series(1, 1) n`,
				nil,
				[]interface{}{&v[0], &v[1], &v[2], &v[3], &v[4], &v[5], &v[6], &v[7], &v[8], &v[9]},
				func(pgx.QueryFuncRow) error { return nil },
			)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkQueryDecode_PG_UUID_to_Go_uuid_UUID_10_rows_1_columns(b *testing.B) {
	defaultConnTestRunner.RunTest(context.Background(), b, func(ctx context.Context, _ testing.TB, conn *pgx.Conn) {
		b.ResetTimer()
		var v [1]uuid.UUID
		for i := 0; i < b.N; i++ {
			_, err := conn.QueryFunc(
				ctx,
				`select gen_random_uuid() from generate_series(1, 10) n`,
				nil,
				[]interface{}{&v[0]},
				func(pgx.QueryFuncRow) error { return nil },
			)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkQueryDecode_PG_UUID_to_Go_uuid_UUID_100_rows_10_columns(b *testing.B) {
	defaultConnTestRunner.RunTest(context.Background(), b, func(ctx context.Context, _ testing.TB, conn *pgx.Conn) {
		b.ResetTimer()
		var v [10]uuid.UUID
		for i := 0; i < b.N; i++ {
			_, err := conn.QueryFunc(
				ctx,
				`select gen_random_uuid(), gen_random_uuid(), gen_random_uuid(), gen_random_uuid(), gen_random_uuid(), gen_random_uuid(), gen_random_uuid(), gen_random_uuid(), gen_random_uuid(), gen_random_uuid() from generate_series(1, 100) n`,
				nil,
				[]interface{}{&v[0], &v[1], &v[2], &v[3], &v[4], &v[5], &v[6], &v[7], &v[8], &v[9]},
				func(pgx.QueryFuncRow) error { return nil },
			)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkQueryDecode_PG_UUID_to_Go_uuid_NullUUID_1_rows_1_columns(b *testing.B) {
	defaultConnTestRunner.RunTest(context.Background(), b, func(ctx context.Context, _ testing.TB, conn *pgx.Conn) {
		b.ResetTimer()
		var v [1]uuid.NullUUID
		for i := 0; i < b.N; i++ {
			_, err := conn.QueryFunc(
				ctx,
				`select gen_random_uuid() from generate_series(1, 1) n`,
				nil,
				[]interface{}{&v[0]},
				func(pgx.QueryFuncRow) error { return nil },
			)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkQueryDecode_PG_UUID_to_Go_uuid_NullUUID_1_rows_10_columns(b *testing.B) {
	defaultConnTestRunner.RunTest(context.Background(), b, func(ctx context.Context, _ testing.TB, conn *pgx.Conn) {
		b.ResetTimer()
		var v [10]uuid.NullUUID
		for i := 0; i < b.N; i++ {
			_, err := conn.QueryFunc(
				ctx,
				`select gen_random_uuid(), gen_random_uuid(), gen_random_uuid(), gen_random_uuid(), gen_random_uuid(), gen_random_uuid(), gen_random_uuid(), gen_random_uuid(), gen_random_uuid(), gen_random_uuid() from generate_series(1, 1) n`,
				nil,
				[]interface{}{&v[0], &v[1], &v[2], &v[3], &v[4], &v[5], &v[6], &v[7], &v[8], &v[9]},
				func(pgx.QueryFuncRow) error { return nil },
			)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkQueryDecode_PG_UUID_to_Go_uuid_NullUUID_10_rows_1_columns(b *testing.B) {
	defaultConnTestRunner.RunTest(context.Background(), b, func(ctx context.Context, _ testing.TB, conn *pgx.Conn) {
		b.ResetTimer()
		var v [1]uuid.NullUUID
		for i := 0; i < b.N; i++ {
			_, err := conn.QueryFunc(
				ctx,
				`select gen_random_uuid() from generate_series(1, 10) n`,
				nil,
				[]interface{}{&v[0]},
				func(pgx.QueryFuncRow) error { return nil },
			)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkQueryDecode_PG_UUID_to_Go_uuid_NullUUID_100_rows_10_columns(b *testing.B) {
	defaultConnTestRunner.RunTest(context.Background(), b, func(ctx context.Context, _ testing.TB, conn *pgx.Conn) {
		b.ResetTimer()
		var v [10]uuid.NullUUID
		for i := 0; i < b.N; i++ {
			_, err := conn.QueryFunc(
				ctx,
				`select gen_random_uuid(), gen_random_uuid(), gen_random_uuid(), gen_random_uuid(), gen_random_uuid(), gen_random_uuid(), gen_random_uuid(), gen_random_uuid(), gen_random_uuid(), gen_random_uuid() from generate_series(1, 100) n`,
				nil,
				[]interface{}{&v[0], &v[1], &v[2], &v[3], &v[4], &v[5], &v[6], &v[7], &v[8], &v[9]},
				func(pgx.QueryFuncRow) error { return nil },
			)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}
